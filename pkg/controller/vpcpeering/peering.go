package vpcpeering

import (
	"github.com/pkg/errors"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/provider-gcp/pkg/clients"
	"github.com/crossplane/provider-gcp/pkg/clients/peering"
	"golang.org/x/net/context"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	compute "google.golang.org/api/compute/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)
// Error strings.
const (
	errNewClient        = "cannot create new Compute Service"
	errNotPeering    = "managed resource is not a Peering"
	errListPeering  = "cannot list external Peering resources"
	errCreatePeering = "cannot create external Peering resource"
	errDeletePeering = "cannot delete external Peering resource"
)

func SetupPeering(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(v1beta1.PeeringGroupKind)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&v1beta1.Peering{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1beta1.PeeringGroupVersionKind),
			managed.WithExternalConnecter(&connector{client: mgr.GetClient()}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connector struct {
	client client.Client
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	projectID, opts, err := gcp.GetAuthInfo(ctx, c.client, mg)
	if err != nil {
		return nil, err
	}

	cmp, err := compute.NewService(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	return &external{compute: cmp, projectID: projectID}, errors.Wrap(err, errNewClient)
}

type external struct {
	compute   *compute.Service
	projectID string
}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	peer, ok := mg.(*v1beta1.Peering)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotPeering)
	}

	r, err := e.compute.Networks.Get(peer.Spec.ForProvider.Project, peer.Spec.ForProvider.Network).Do()
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(resource.Ignore(gcp.IsErrorNotFound, err), errListPeering)
	}

	o := peering.Observation{Peering: findPeering(peer.Spec.ForProvider.Name, r.Peerings)}
	if o.Peering == nil {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	eo := managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: peering.IsUpToDate(peer.Spec.ForProvider, o.Peering),
	}

	peering.UpdateStatus(&peer.Status, o)

	return eo, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	peer, ok := mg.(*v1beta1.Peering)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotPeering)
	}

	peer.Status.SetConditions(xpv1.Creating())

	_, err := e.compute.Networks.AddPeering(peer.Spec.ForProvider.Project, peer.Spec.ForProvider.Network, &compute.NetworksAddPeeringRequest{
		Name:             peer.Spec.ForProvider.Name,
		PeerNetwork:      peer.Spec.ForProvider.PeerNetwork,
		AutoCreateRoutes: true,
	}).Do()

	return managed.ExternalCreation{}, errors.Wrap(resource.Ignore(gcp.IsErrorAlreadyExists, err), errCreatePeering)
}

func findPeering(name string, peerings []*compute.NetworkPeering) *compute.NetworkPeering {
	for _, peering := range peerings {
		if peering.Name == name {
			return peering
		}
	}
	return nil
}

func (e *external) Update(ctx context.Context, msg resource.Managed) (managed.ExternalUpdate, error) {
	// TOOD: support update operation
	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
	peer, ok := mg.(*v1beta1.Peering)
	if !ok {
		return errors.New(errNotPeering)
	}

	peer.Status.SetConditions(xpv1.Deleting())

	_, err := e.compute.Networks.RemovePeering(peer.Spec.ForProvider.Project, peer.Spec.ForProvider.Network, &compute.NetworksRemovePeeringRequest{
		Name: peer.Spec.ForProvider.Name,
	}).Do()

	return errors.Wrap(resource.Ignore(gcp.IsErrorNotFound, err), errDeletePeering)
}