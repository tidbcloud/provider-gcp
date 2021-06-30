package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "dns.gcp.crossplane.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// ManagedZone type metadata.
var (
	ManagedZoneKind             = reflect.TypeOf(ManagedZone{}).Name()
	ManagedZoneGroupKind        = schema.GroupKind{Group: Group, Kind: ManagedZoneKind}.String()
	ManagedZoneKindAPIVersion   = ManagedZoneKind + "." + SchemeGroupVersion.String()
	ManagedZoneGroupVersionKind = SchemeGroupVersion.WithKind(ManagedZoneKind)
)

// ResourceRecordSet type metadata.
var (
	ResourceRecordSetKind             = reflect.TypeOf(ResourceRecordSet{}).Name()
	ResourceRecordSetGroupKind        = schema.GroupKind{Group: Group, Kind: ResourceRecordSetKind}.String()
	ResourceRecordSetKindAPIVersion   = ResourceRecordSetKind + "." + SchemeGroupVersion.String()
	ResourceRecordSetGroupVersionKind = SchemeGroupVersion.WithKind(ResourceRecordSetKind)
)

func init() {
	SchemeBuilder.Register(&ManagedZone{}, &ManagedZoneList{})
	SchemeBuilder.Register(&ResourceRecordSet{}, &ResourceRecordSetList{})
}
