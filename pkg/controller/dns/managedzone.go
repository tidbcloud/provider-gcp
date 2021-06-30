package dns

import (
	"context"
	"fmt"

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
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/crossplane/provider-gcp/apis/dns/v1alpha1"

	gcp "github.com/crossplane/provider-gcp/pkg/clients"
	dns2 "github.com/crossplane/provider-gcp/pkg/clients/dns"
)

const (
	errNotManagedZone    = "managed resource is not of type ManagedZone"
	errNewClient         = "cannot create client"
	errCreateManagedZone = "cannot create ManagedZone"
	errGet               = "failed to get the ManagedZone resource"
	errDelete            = "failed to delete the ManagedZone resource"
	errUnexpectedObject  = "The managed resource is not an ManagedZone resource"
	errKubeUpdate        = "failed to update the ManagedZone custom resource"
)

// SetupManagedZone adds a controller that reconciles ManagedZones.
func SetupManagedZone(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(v1alpha1.ManagedZoneGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&v1alpha1.ManagedZone{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha1.ManagedZoneGroupVersionKind),
			managed.WithExternalConnecter(&connecter{kube: mgr.GetClient()}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithInitializers(managed.NewNameAsExternalName(mgr.GetClient())),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connecter struct {
	kube client.Client
}

// Connect sets up dnsservice client using credentials from the provider
func (c *connecter) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	//cr, ok := mg.(*v1alpha1.ManagedZone)
	//if !ok {
	//	return nil, errors.New(errNotManagedZone)
	//}
	//
	//p := &gcpv1alpha3.Provider{}
	//if err := c.kube.Get(ctx, types.NamespacedName{Name: cr.Spec.ProviderReference.Name}, p); err != nil {
	//	return nil, errors.Wrap(err, errGetProvider)
	//}
	//
	//if p.GetCredentialsSecretReference() == nil {
	//	return nil, errors.New(errProviderSecretRef)
	//}
	//
	//s := &corev1.Secret{}
	//n := types.NamespacedName{Namespace: p.Spec.CredentialsSecretRef.Namespace, Name: p.Spec.CredentialsSecretRef.Name}
	//if err := c.kube.Get(ctx, n, s); err != nil {
	//	return nil, errors.Wrap(err, errGetProviderSecret)
	//}

	projectID, opts, err := gcp.GetAuthInfo(ctx, c.kube, mg)
	if err != nil {
		return nil, err
	}

	dnsService, err := dns.NewService(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	return &managedZoneExternal{Service: dnsService, kube: c.kube, projectID: projectID}, nil
}

type managedZoneExternal struct {
	kube client.Client
	*dns.Service
	projectID string
}

func (e *managedZoneExternal) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.ManagedZone)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotManagedZone)
	}

	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}

	t, err := e.ManagedZones.Get(e.projectID, meta.GetExternalName(cr)).Do()
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(resource.Ignore(gcp.IsErrorNotFound, err), errGet)
	}

	current := cr.Spec.ForProvider.DeepCopy()
	dns2.LateInitialize(&cr.Spec.ForProvider, t)
	if !cmp.Equal(current, &cr.Spec.ForProvider) {
		if err := e.kube.Update(ctx, cr); err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, errKubeUpdate)
		}
	}

	cr.Status.SetConditions(xpv1.Available())
	cr.Status.AtProvider.ID = t.Id
	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: dns2.IsUpToDate(cr.Spec.ForProvider, *t),
	}, nil

}

func (e *managedZoneExternal) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.ManagedZone)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotManagedZone)
	}
	cr.Status.SetConditions(xpv1.Creating())

	resp, err := e.ManagedZones.Create(e.projectID, dns2.GenerateManagedZone(e.projectID, cr.Spec.ForProvider)).Do()
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateManagedZone)
	}

	meta.SetExternalName(cr, fmt.Sprintf("%d", resp.Id))
	return managed.ExternalCreation{}, errors.Wrap(e.kube.Update(ctx, cr), errKubeUpdate)
}

func (e *managedZoneExternal) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}

func (e *managedZoneExternal) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.ManagedZone)
	if !ok {
		return errors.New(errUnexpectedObject)
	}

	cr.Status.SetConditions(xpv1.Deleting())

	phzID := meta.GetExternalName(cr)
	resp, err := e.ResourceRecordSets.List(e.projectID, phzID).Do()
	if err == nil {
		for _, rs := range resp.Rrsets {
			_, _ = e.Changes.Create(e.projectID, phzID, &dns.Change{
				Deletions: []*dns.ResourceRecordSet{rs},
			}).Context(ctx).Do()
		}
	}

	err = e.ManagedZones.Delete(e.projectID, phzID).Do()
	return errors.Wrap(resource.Ignore(gcp.IsErrorNotFound, err), errDelete)
}
