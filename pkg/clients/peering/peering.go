package peering

import (
	"github.com/google/go-cmp/cmp"

	"github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1"

	compute "google.golang.org/api/compute/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VPC Network peering states.
const (
	PeeringStateActive   = "ACTIVE"
	PeeringStateInactive = "INACTIVE"
)

// IsUpToDate check whether the resource is date
func IsUpToDate(p v1beta1.PeeringParameters, observed *compute.NetworkPeering) bool {
	return cmp.Equal(p.Name, observed.Name)
}

// Observation observate peering state
type Observation struct {
	Peering *compute.NetworkPeering
}

// UpdateStatus update peering status
func UpdateStatus(s *v1beta1.PeeringStatus, o Observation) {
	s.AtProvider.Peering = o.Peering.Network

	if o.Peering == nil {
		s.SetConditions(xpv1.Unavailable())
		return
	}

	switch o.Peering.State {
	case PeeringStateActive:
		s.SetConditions(xpv1.Available())
	case PeeringStateInactive:
		s.SetConditions(xpv1.Unavailable())
	}
}
