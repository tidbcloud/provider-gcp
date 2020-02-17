// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalAddress) DeepCopyInto(out *GlobalAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalAddress.
func (in *GlobalAddress) DeepCopy() *GlobalAddress {
	if in == nil {
		return nil
	}
	out := new(GlobalAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalAddressList) DeepCopyInto(out *GlobalAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalAddressList.
func (in *GlobalAddressList) DeepCopy() *GlobalAddressList {
	if in == nil {
		return nil
	}
	out := new(GlobalAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalAddressNameReferencer) DeepCopyInto(out *GlobalAddressNameReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalAddressNameReferencer.
func (in *GlobalAddressNameReferencer) DeepCopy() *GlobalAddressNameReferencer {
	if in == nil {
		return nil
	}
	out := new(GlobalAddressNameReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalAddressObservation) DeepCopyInto(out *GlobalAddressObservation) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalAddressObservation.
func (in *GlobalAddressObservation) DeepCopy() *GlobalAddressObservation {
	if in == nil {
		return nil
	}
	out := new(GlobalAddressObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalAddressParameters) DeepCopyInto(out *GlobalAddressParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.AddressType != nil {
		in, out := &in.AddressType, &out.AddressType
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(NetworkURIReferencerForGlobalAddress)
		**out = **in
	}
	if in.PrefixLength != nil {
		in, out := &in.PrefixLength, &out.PrefixLength
		*out = new(int64)
		**out = **in
	}
	if in.Purpose != nil {
		in, out := &in.Purpose, &out.Purpose
		*out = new(string)
		**out = **in
	}
	if in.Subnetwork != nil {
		in, out := &in.Subnetwork, &out.Subnetwork
		*out = new(string)
		**out = **in
	}
	if in.SubnetworkRef != nil {
		in, out := &in.SubnetworkRef, &out.SubnetworkRef
		*out = new(SubnetworkURIReferencerForGlobalAddress)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalAddressParameters.
func (in *GlobalAddressParameters) DeepCopy() *GlobalAddressParameters {
	if in == nil {
		return nil
	}
	out := new(GlobalAddressParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalAddressSpec) DeepCopyInto(out *GlobalAddressSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalAddressSpec.
func (in *GlobalAddressSpec) DeepCopy() *GlobalAddressSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalAddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalAddressStatus) DeepCopyInto(out *GlobalAddressStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalAddressStatus.
func (in *GlobalAddressStatus) DeepCopy() *GlobalAddressStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalAddressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Network) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkList) DeepCopyInto(out *NetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Network, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkList.
func (in *NetworkList) DeepCopy() *NetworkList {
	if in == nil {
		return nil
	}
	out := new(NetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkObservation) DeepCopyInto(out *NetworkObservation) {
	*out = *in
	if in.Peerings != nil {
		in, out := &in.Peerings, &out.Peerings
		*out = make([]*NetworkPeering, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NetworkPeering)
				**out = **in
			}
		}
	}
	if in.Subnetworks != nil {
		in, out := &in.Subnetworks, &out.Subnetworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkObservation.
func (in *NetworkObservation) DeepCopy() *NetworkObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParameters) DeepCopyInto(out *NetworkParameters) {
	*out = *in
	if in.AutoCreateSubnetworks != nil {
		in, out := &in.AutoCreateSubnetworks, &out.AutoCreateSubnetworks
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.RoutingConfig != nil {
		in, out := &in.RoutingConfig, &out.RoutingConfig
		*out = new(NetworkRoutingConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParameters.
func (in *NetworkParameters) DeepCopy() *NetworkParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPeering) DeepCopyInto(out *NetworkPeering) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPeering.
func (in *NetworkPeering) DeepCopy() *NetworkPeering {
	if in == nil {
		return nil
	}
	out := new(NetworkPeering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRoutingConfig) DeepCopyInto(out *NetworkRoutingConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRoutingConfig.
func (in *NetworkRoutingConfig) DeepCopy() *NetworkRoutingConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkRoutingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkURIReferencer) DeepCopyInto(out *NetworkURIReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkURIReferencer.
func (in *NetworkURIReferencer) DeepCopy() *NetworkURIReferencer {
	if in == nil {
		return nil
	}
	out := new(NetworkURIReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkURIReferencerForGlobalAddress) DeepCopyInto(out *NetworkURIReferencerForGlobalAddress) {
	*out = *in
	out.NetworkURIReferencer = in.NetworkURIReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkURIReferencerForGlobalAddress.
func (in *NetworkURIReferencerForGlobalAddress) DeepCopy() *NetworkURIReferencerForGlobalAddress {
	if in == nil {
		return nil
	}
	out := new(NetworkURIReferencerForGlobalAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkURIReferencerForSubnetwork) DeepCopyInto(out *NetworkURIReferencerForSubnetwork) {
	*out = *in
	out.NetworkURIReferencer = in.NetworkURIReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkURIReferencerForSubnetwork.
func (in *NetworkURIReferencerForSubnetwork) DeepCopy() *NetworkURIReferencerForSubnetwork {
	if in == nil {
		return nil
	}
	out := new(NetworkURIReferencerForSubnetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnetwork) DeepCopyInto(out *Subnetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnetwork.
func (in *Subnetwork) DeepCopy() *Subnetwork {
	if in == nil {
		return nil
	}
	out := new(Subnetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subnetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetworkList) DeepCopyInto(out *SubnetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subnetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetworkList.
func (in *SubnetworkList) DeepCopy() *SubnetworkList {
	if in == nil {
		return nil
	}
	out := new(SubnetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetworkObservation) DeepCopyInto(out *SubnetworkObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetworkObservation.
func (in *SubnetworkObservation) DeepCopy() *SubnetworkObservation {
	if in == nil {
		return nil
	}
	out := new(SubnetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetworkParameters) DeepCopyInto(out *SubnetworkParameters) {
	*out = *in
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(NetworkURIReferencerForSubnetwork)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnableFlowLogs != nil {
		in, out := &in.EnableFlowLogs, &out.EnableFlowLogs
		*out = new(bool)
		**out = **in
	}
	if in.PrivateIPGoogleAccess != nil {
		in, out := &in.PrivateIPGoogleAccess, &out.PrivateIPGoogleAccess
		*out = new(bool)
		**out = **in
	}
	if in.SecondaryIPRanges != nil {
		in, out := &in.SecondaryIPRanges, &out.SecondaryIPRanges
		*out = make([]*SubnetworkSecondaryRange, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SubnetworkSecondaryRange)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetworkParameters.
func (in *SubnetworkParameters) DeepCopy() *SubnetworkParameters {
	if in == nil {
		return nil
	}
	out := new(SubnetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetworkSecondaryRange) DeepCopyInto(out *SubnetworkSecondaryRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetworkSecondaryRange.
func (in *SubnetworkSecondaryRange) DeepCopy() *SubnetworkSecondaryRange {
	if in == nil {
		return nil
	}
	out := new(SubnetworkSecondaryRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetworkSpec) DeepCopyInto(out *SubnetworkSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetworkSpec.
func (in *SubnetworkSpec) DeepCopy() *SubnetworkSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetworkStatus) DeepCopyInto(out *SubnetworkStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetworkStatus.
func (in *SubnetworkStatus) DeepCopy() *SubnetworkStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetworkURIReferencer) DeepCopyInto(out *SubnetworkURIReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetworkURIReferencer.
func (in *SubnetworkURIReferencer) DeepCopy() *SubnetworkURIReferencer {
	if in == nil {
		return nil
	}
	out := new(SubnetworkURIReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetworkURIReferencerForGlobalAddress) DeepCopyInto(out *SubnetworkURIReferencerForGlobalAddress) {
	*out = *in
	out.SubnetworkURIReferencer = in.SubnetworkURIReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetworkURIReferencerForGlobalAddress.
func (in *SubnetworkURIReferencerForGlobalAddress) DeepCopy() *SubnetworkURIReferencerForGlobalAddress {
	if in == nil {
		return nil
	}
	out := new(SubnetworkURIReferencerForGlobalAddress)
	in.DeepCopyInto(out)
	return out
}