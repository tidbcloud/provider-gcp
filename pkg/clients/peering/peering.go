package peering

import (
	"github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1"
	"github.com/google/go-cmp/cmp"
	compute "google.golang.org/api/compute/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VPC Network peering states.
const (
	PeeringStateActive   = "ACTIVE"
	PeeringStateInactive = "INACTIVE"
)

func IsUpToDate(p v1beta1.PeeringParameters, observed *compute.NetworkPeering) bool {
	if observed == nil {
		return false
	}
	return cmp.Equal(p.Name, observed.Name)
}

type Observation struct {
	Peering *compute.NetworkPeering
}

func UpdateStatus(s *v1beta1.PeeringStatus, o Observation) {
	if o.Peering == nil {
		s.SetConditions(xpv1.Unavailable())
		return
	}

	s.AtProvider.Peering = o.Peering.Network
	switch o.Peering.State {
	case PeeringStateActive:
		s.SetConditions(xpv1.Available())
	case PeeringStateInactive:
		s.SetConditions(xpv1.Unavailable())
	}
}
