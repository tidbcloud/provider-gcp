package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PeeringParameters struct {
	Project string `json:"project"`
	
	Network string `json:"network"`

	// Name: Name of the peering, which should conform to RFC1035.
	Name string `json:"name,omitempty"`

	// PeerNetwork: URL of the peer network. It can be either full URL or
	// partial URL. The peer network may belong to a different project. If
	// the partial URL does not contain project, it is assumed that the peer
	// network is in the same project as the current network.
	PeerNetwork string `json:"peerNetwork,omitempty"`

	// AutoCreateRoutes: This field will be deprecated soon. Use
	// exchange_subnet_routes in network_peering instead. Indicates whether
	// full mesh connectivity is created and managed automatically between
	// peered networks. Currently this field should always be true since
	// Google Compute Engine will automatically create and manage subnetwork
	// routes between two networks when peering state is ACTIVE.
	AutoCreateRoutes bool `json:"autoCreateRoutes,omitempty"`

}

// A Peering is a managed resource that represents a Google Cloud Service
// Networking Peering.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Peering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PeeringSpec   `json:"spec"`
	Status PeeringStatus `json:"status,omitempty"`
}

// PeeringObservation is used to show the observed state of the Peering.
type PeeringObservation struct {
	// Peering: The name of the VPC Network Peering Peering that was created
	// by the service producer.
	Peering string `json:"peering,omitempty"`

	// Service: The name of the peering service that's associated with this
	// Peering, in the following format: `services/{service name}`.
	Service string `json:"service,omitempty"`
}

// A PeeringSpec defines the desired state of a Peering.
type PeeringSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PeeringParameters `json:"forProvider"`
}

// A PeeringStatus represents the observed state of a Peering.
type PeeringStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringList contains a list of Peering.
type PeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Peering `json:"items"`
}
