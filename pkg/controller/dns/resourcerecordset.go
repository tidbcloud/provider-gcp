package dns

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"google.golang.org/api/dns/v1"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/crossplane/provider-gcp/apis/dns/v1alpha1"
	gcp "github.com/crossplane/provider-gcp/pkg/clients"
	dns2 "github.com/crossplane/provider-gcp/pkg/clients/dns"
)

const (
	errNotResourceRecordSet        = "managed resource is not of type ResourceRecordSet"
	errCreateResourceRecordSet     = "cannot create ResourceRecordSet"
	errGetResourceRecordSet        = "failed to get the ResourceRecordSet resource"
	errDeleteResourceRecordSet     = "failed to delete the ResourceRecordSet resource"
	errKubeUpdateResourceRecordSet = "failed to update the ResourceRecordSet custom resource"
)

// SetupResourceRecordSet adds a controller that reconciles Topics.
func SetupResourceRecordSet(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(v1alpha1.ResourceRecordSetGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&v1alpha1.ResourceRecordSet{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha1.ResourceRecordSetGroupVersionKind),
			managed.WithExternalConnecter(&rconnector{kube: mgr.GetClient()}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithInitializers(managed.NewNameAsExternalName(mgr.GetClient())),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type rconnector struct {
	kube client.Client
}

func (c *rconnector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	projectID, opts, err := gcp.GetAuthInfo(ctx, c.kube, mg)
	if err != nil {
		return nil, err
	}

	dnsService, err := dns.NewService(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	return &resourceRecordSetExternal{Service: dnsService, kube: c.kube, projectID: projectID}, nil
}

type resourceRecordSetExternal struct {
	kube client.Client
	*dns.Service
	projectID string
}

func (e *resourceRecordSetExternal) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.ResourceRecordSet)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotResourceRecordSet)
	}

	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}

	resp, err := e.ResourceRecordSets.List(e.projectID, *cr.Spec.ForProvider.ZoneID).Name(dns2.AppendDot(meta.GetExternalName(cr))).Do()
	if err != nil {
		// Either there is err and retry. Or Resource does not exist.
		return managed.ExternalObservation{
			ResourceExists: false,
		}, errors.Wrap(resource.Ignore(gcp.IsErrorNotFound, err), errGetResourceRecordSet)
	}

	if resp.Rrsets == nil || len(resp.Rrsets) == 0 {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}

	res := resp.Rrsets[0]

	current := cr.Spec.ForProvider.DeepCopy()
	dns2.LateInitializeResourceRecordSet(&cr.Spec.ForProvider, res)
	if !cmp.Equal(current, &cr.Spec.ForProvider) {
		if err := e.kube.Update(ctx, cr); err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, errKubeUpdateResourceRecordSet)
		}
	}

	cr.Status.SetConditions(xpv1.Available())
	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: dns2.IsUpToDateResourceRecordSet(cr.Spec.ForProvider, *res),
	}, nil
}

func (e *resourceRecordSetExternal) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.ResourceRecordSet)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotResourceRecordSet)
	}
	cr.SetConditions(xpv1.Creating())

	_, err := e.Changes.Create(e.projectID, *cr.Spec.ForProvider.ZoneID, &dns.Change{
		Additions: []*dns.ResourceRecordSet{dns2.GenerateResourceRecordSet(meta.GetExternalName(cr), cr.Spec.ForProvider)},
	}).Do()

	return managed.ExternalCreation{}, errors.Wrap(err, errCreateResourceRecordSet)
}

func (e *resourceRecordSetExternal) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}

func (e *resourceRecordSetExternal) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.ResourceRecordSet)
	if !ok {
		return errors.New(errNotResourceRecordSet)
	}

	cr.Status.SetConditions(xpv1.Deleting())

	_, err := e.Changes.Create(e.projectID, *cr.Spec.ForProvider.ZoneID, &dns.Change{
		Deletions: []*dns.ResourceRecordSet{
			dns2.GenerateResourceRecordSet(meta.GetExternalName(cr), cr.Spec.ForProvider),
		},
	}).Do()

	return errors.Wrap(resource.Ignore(gcp.IsErrorNotFound, err), errDeleteResourceRecordSet)
}
