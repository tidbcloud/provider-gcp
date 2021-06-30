package v1beta1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "vpcpeering.gcp.crossplane.io"
	Version = "v1beta1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// Peering type metadata.
var (
	PeeringKind             = reflect.TypeOf(Peering{}).Name()
	PeeringGroupKind        = schema.GroupKind{Group: Group, Kind: PeeringKind}.String()
	PeeringKindAPIVersion   = PeeringKind + "." + SchemeGroupVersion.String()
	PeeringGroupVersionKind = SchemeGroupVersion.WithKind(PeeringKind)
)

func init() {
	SchemeBuilder.Register(&Peering{}, &PeeringList{})
}
