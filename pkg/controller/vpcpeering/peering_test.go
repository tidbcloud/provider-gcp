package vpcpeering

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1"
	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/gomega"
)

func TestDelete(t *testing.T) {
	g := NewGomegaWithT(t)
	type args struct {
		ctx context.Context
		mg  resource.Managed
	}

	cases := map[string]struct {
		handler   http.Handler
		reason    string
		args      args
		want      error
		callCount int
	}{
		"NotAPeering": {
			reason: "We should return an error if the supplied managed resource is not a peering",
			args: args{
				mg: nil,
			},
			want: errors.New(errNotPeering),
		},
		"DeleteNonExistingPeering": {
			reason: "We should successful when no peering",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project:     "test-project",
							Network:     "test-network",
							Name:        "test-peering-name",
							PeerNetwork: newPeerNetwork("other-project", "other-network"),
						},
					},
				},
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				g.Expect(r.URL.Path).Should(Equal("/projects/test-project/global/networks/test-network/removePeering"))
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(&compute.Operation{HttpErrorMessage: gcpErrDeletePeering})
			}),
		},
		"SuccessDeleteNotSameProject": {
			reason: "We should call delete requester network's peering and return successful",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project:     "test-project",
							Network:     "test-network",
							Name:        "test-peering-name",
							PeerNetwork: newPeerNetwork("other-project", "other-network"),
						},
					},
				},
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				g.Expect(r.URL.Path).Should(Equal("/projects/test-project/global/networks/test-network/removePeering"))
				w.WriteHeader(http.StatusNoContent)
			}),
		},
		"SuccessDeleteSameProject": {
			reason: "We should call delete requester and accepter network's peering and return successful",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project:     "test-project",
							Network:     "test-network",
							Name:        "test-peering-name",
							PeerNetwork: newPeerNetwork("test-project", "other-network"),
						},
					},
				},
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(http.StatusNoContent)
			}),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(tc.handler)
			defer server.Close()
			s, _ := compute.NewService(context.Background(), option.WithEndpoint(server.URL), option.WithoutAuthentication())
			e := &external{projectID: "test", compute: s}
			err := e.Delete(tc.args.ctx, tc.args.mg)
			if err != nil {
				if diff := cmp.Diff(tc.want.Error(), err.Error()); diff != "" {
					t.Errorf("\n%s\ne.Create(...): -want error, +got error:\n%s\n", tc.reason, diff)
				}
			}
		})
	}
}

func TestObserve(t *testing.T) {
	// g := NewGomegaWithT(t)
	type args struct {
		ctx context.Context
		mg  resource.Managed
	}

	cases := map[string]struct {
		handler         http.Handler
		reason          string
		args            args
		want            error
		wantObservation managed.ExternalObservation
		callCount       int
	}{
		"NotAPeering": {
			reason: "We should return an error if the supplied managed resource is not a peering",
			args: args{
				mg: nil,
			},
			want:            errors.New(errNotPeering),
			wantObservation: managed.ExternalObservation{},
		},
		"ObserveResourceNonExist": {
			reason: "Peering not exist in gcp provider",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project:     "test-project",
							Network:     "test-network",
							Name:        "test-peering-name",
							PeerNetwork: newPeerNetwork("other-project", "other-network"),
						},
					},
				},
			},
			wantObservation: managed.ExternalObservation{
				ResourceExists: false,
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(&compute.Network{
					Peerings: []*compute.NetworkPeering{
						{},
					},
				})
			}),
		},
		"ObserveResourcexist": {
			reason: "Peering exist in gcp provider",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project:     "test-project",
							Network:     "test-network",
							Name:        "test-peering-name",
							PeerNetwork: newPeerNetwork("other-project", "other-network"),
						},
					},
				},
			},
			wantObservation: managed.ExternalObservation{
				ResourceExists: true,
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(&compute.Network{
					Peerings: []*compute.NetworkPeering{
						{
							Network: newPeerNetwork("other-project", "other-network"),
						},
					},
				})
			}),
		},
		"ObserveResourceUpdateToDate": {
			reason: "Peering exist in gcp provider",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project:     "test-project",
							Network:     "test-network",
							Name:        "test-peering-name",
							PeerNetwork: newPeerNetwork("other-project", "other-network"),
						},
					},
				},
			},
			wantObservation: managed.ExternalObservation{
				ResourceExists:   true,
				ResourceUpToDate: true,
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(&compute.Network{
					Peerings: []*compute.NetworkPeering{
						{
							Network: newPeerNetwork("other-project", "other-network"),
							Name:    "test-peering-name",
						},
					},
				})
			}),
		},
		"ObserveResourceUpdateToDateSameProject": {
			reason: "Request and accepter peering exist in gcp provider",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project: "test-project",
							Network: "test-network",
							Name:    "test-peering-name",
							// same project
							PeerNetwork: newPeerNetwork("test-project", "other-network"),
						},
					},
				},
			},
			wantObservation: managed.ExternalObservation{
				ResourceExists:   true,
				ResourceUpToDate: true,
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(&compute.Network{
					Peerings: []*compute.NetworkPeering{
						{
							Network: newPeerNetwork("test-project", "other-network"),
							Name:    "test-peering-name",
						},
						{
							Network: newPeerNetwork("test-project", "test-network"),
							Name:    "test-peering-name",
						},
					},
				})
			}),
		},
		"ObserveResourceNonExistInSameProjectWhenStatusDeleting": {
			reason: "If status is deleteing, it is considered non-exist only if requester and accepter peering is empty",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
						// object is in deleting, we need delete all peering
						DeletionTimestamp: &v1.Time{Time: time.Now()},
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project: "test-project",
							Network: "test-network",
							Name:    "test-peering-name",
							// same project
							PeerNetwork: newPeerNetwork("test-project", "other-network"),
						},
					},
				},
			},
			wantObservation: managed.ExternalObservation{
				ResourceExists: true,
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(&compute.Network{
					Peerings: []*compute.NetworkPeering{
						{
							// still have one peering not deleted
							Network: newPeerNetwork("test-project", "test-network"),
							Name:    "test-peering-name",
						},
					},
				})
			}),
		},
		"ObserveResourceNonExistSameProject": {
			reason: "It is considered exist only if requester and accepter peering exist in provider",
			args: args{
				mg: &v1beta1.Peering{
					ObjectMeta: v1.ObjectMeta{
						Name: "test-peering",
					},
					Spec: v1beta1.PeeringSpec{
						ForProvider: v1beta1.PeeringParameters{
							Project: "test-project",
							Network: "test-network",
							Name:    "test-peering-name",
							// same project
							PeerNetwork: newPeerNetwork("test-project", "other-network"),
						},
					},
				},
			},
			wantObservation: managed.ExternalObservation{
				ResourceExists: false,
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(&compute.Network{
					Peerings: []*compute.NetworkPeering{
						{
							// only have one peering, need create another
							Network: newPeerNetwork("test-project", "test-network"),
							Name:    "test-peering-name",
						},
					},
				})
			}),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(tc.handler)
			defer server.Close()
			s, _ := compute.NewService(context.Background(), option.WithEndpoint(server.URL), option.WithoutAuthentication())
			e := &external{projectID: "test", compute: s}
			ob, err := e.Observe(tc.args.ctx, tc.args.mg)
			if err != nil {
				if diff := cmp.Diff(tc.want.Error(), err.Error()); diff != "" {
					t.Errorf("\n%s\ne.Create(...): -want error, +got error:\n%s\n", tc.reason, diff)
				}
			}
			if diff := cmp.Diff(tc.wantObservation, ob); diff != "" {
				t.Errorf("\n%s\ne.Create(...): -want error, +got error:\n%s\n", tc.reason, diff)
			}
		})
	}
}
