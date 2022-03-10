package vpcpeering

import (
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	gcp "github.com/crossplane/provider-gcp/pkg/clients"
	"github.com/crossplane/provider-gcp/pkg/clients/peering"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1"
)

// Error strings.
const (
	errNewClient     = "cannot create new Compute Service"
	errNotPeering    = "managed resource is not a Peering"
	errListPeering   = "cannot list external Peering resources"
	errCreatePeering = "cannot create external Peering resource"
	errDeletePeering = "cannot delete external Peering resource"
	// When GCP deletes a non-existing peering, the returned status code is 400, and the error is `There is no peering..., bad request`
	gcpErrDeletePeering = "There is no peering"
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

	project, network, err := parsePeerNetwork(peer.Spec.ForProvider.PeerNetwork)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, err
	}

	o := peering.Observation{Peering: findPeering(peer.Spec.ForProvider.PeerNetwork, r.Peerings)}
	// If it is not the same project peering, if the request peering is empty, it can be considered that the resource does not exist
	if o.Peering == nil && project != peer.Spec.ForProvider.Project {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	if project == peer.Spec.ForProvider.Project {
		r2, err := e.compute.Networks.Get(project, network).Do()
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(resource.Ignore(gcp.IsErrorNotFound, err), errListPeering)
		}
		requesterPeerNetwork := newPeerNetwork(peer.Spec.ForProvider.Project, peer.Spec.ForProvider.Network)
		peering := findPeering(requesterPeerNetwork, r2.Peerings)
		// If it is the same project peering, if the request peering and accepter peering is empty, it can be considered that the resource does not exist.
		// Because we need ensure all peering will removed from GCP cloud when CRD deleted
		if meta.WasDeleted(peer) {
			// It is considered non-exist only if requester and accepter peering is empty
			if peering == nil && o.Peering == nil {
				return managed.ExternalObservation{ResourceExists: false}, nil
			}
		} else {
			// It is considered non-exist only if requester or accepter peering is empty
			if peering == nil || o.Peering == nil {
				return managed.ExternalObservation{ResourceExists: false}, nil
			}
		}

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
	if err != nil && !gcp.IsErrorAlreadyExists(err) {
		return managed.ExternalCreation{}, err
	}

	project, network, err := parsePeerNetwork(peer.Spec.ForProvider.PeerNetwork)
	if err != nil {
		return managed.ExternalCreation{}, err
	}
	if project == peer.Spec.ForProvider.Project {
		_, err := e.compute.Networks.AddPeering(project, network, &compute.NetworksAddPeeringRequest{
			Name:             peer.Spec.ForProvider.Name,
			PeerNetwork:      newPeerNetwork(peer.Spec.ForProvider.Project, peer.Spec.ForProvider.Network),
			AutoCreateRoutes: true,
		}).Do()
		if err != nil && !gcp.IsErrorAlreadyExists(err) {
			return managed.ExternalCreation{}, err
		}
	}
	return managed.ExternalCreation{}, errors.Wrap(resource.Ignore(gcp.IsErrorAlreadyExists, err), errCreatePeering)
}

func findPeering(peerNetwork string, peerings []*compute.NetworkPeering) *compute.NetworkPeering {
	for _, p := range peerings {
		if p.Network == peerNetwork {
			return p
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
	if err != nil && !gcp.IsErrorNotFound(err) && !isErrorNoPeering(err) {
		return err
	}

	project, network, err := parsePeerNetwork(peer.Spec.ForProvider.PeerNetwork)
	if err != nil {
		return err
	}

	if project == peer.Spec.ForProvider.Project {
		_, err = e.compute.Networks.RemovePeering(project, network, &compute.NetworksRemovePeeringRequest{
			Name: peer.Spec.ForProvider.Name,
		}).Do()
		if err != nil && !gcp.IsErrorNotFound(err) && !isErrorNoPeering(err) {
			return err
		}
	}

	return errors.Wrap(resource.Ignore(gcp.IsErrorNotFound, err), errDeletePeering)
}

func parsePeerNetwork(peerNetwork string) (project string, Network string, err error) {
	res := strings.TrimPrefix(peerNetwork, "https://www.googleapis.com/compute/v1/projects/")
	item := strings.Split(res, "/")
	if len(item) < 2 {
		return "", "", fmt.Errorf("invalid peer network url")
	}

	return item[0], item[len(item)-1], nil
}
func newPeerNetwork(project string, network string) string {
	return fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/networks/%s", project, network)
}

func isErrorNoPeering(err error) bool {
	if strings.Contains(err.Error(), gcpErrDeletePeering) {
		return true
	}
	return false
}
