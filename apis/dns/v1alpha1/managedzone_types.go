package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ManagedZoneParameters defines parameters for a ManagedZone.
type ManagedZoneParameters struct {
	// Name: User assigned name for this resource. Must be unique within the
	// project. The name must be 1-63 characters long, must begin with a
	// letter, end with a letter or digit, and only contain lowercase
	// letters, digits or dashes.
	Name string `json:"name,omitempty"`

	// Labels are used as additional metadata on ManagedZone.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Description: A mutable string of at most 1024 characters associated
	// with this resource for the user's convenience. Has no effect on the
	// managed zone's function.
	// +optional
	Description string `json:"description,omitempty"`

	// Network: the VPC network to bind to.
	Network string `json:"network,omitempty"`

	// DNSName: The DNS name of this managed zone, for instance
	// "example.com.".
	DNSName string `json:"dnsName,omitempty"`

	// Visibility: The zone's visibility: public zones are exposed to the
	// Internet, while private zones are visible only to Virtual Private
	// Cloud resources.
	//
	// Possible values:
	//   "private"
	//   "public"
	// +optional
	Visibility string `json:"visibility,omitempty"`
}

// ManagedZoneObservation is used to show the observed state of the
// ManagedZone resource on GCP. All fields in this structure should only
// be populated from GCP responses; any changes made to the k8s resource outside
// of the crossplane gcp controller will be ignored and overwritten.
type ManagedZoneObservation struct {
	// ID that gcp dns assigned to the managed zone when you created
	// it.
	ID uint64 `json:"id,omitempty"`
}

// ManagedZoneSpec defines the desired state of a
// ManagedZone.
type ManagedZoneSpec struct {
	xpv1.ResourceSpec `json:",inline"`

	ForProvider ManagedZoneParameters `json:"forProvider"`
}

// ManagedZoneStatus represents the observed state of a
// ManagedZone.
type ManagedZoneStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ManagedZoneObservation `json:"atProvider,omitempty"`
}

// ManagedZone is a managed resource that represents a Google IAM Service Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="ID",type="string",JSONPath=".status.atProvider.id"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ManagedZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedZoneSpec   `json:"spec"`
	Status ManagedZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedZoneList contains a list of ManagedZone types
type ManagedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedZone `json:"items"`
}
