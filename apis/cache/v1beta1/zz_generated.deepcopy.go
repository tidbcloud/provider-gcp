//go:build !ignore_autogenerated
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
func (in *CloudMemorystoreInstance) DeepCopyInto(out *CloudMemorystoreInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstance.
func (in *CloudMemorystoreInstance) DeepCopy() *CloudMemorystoreInstance {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudMemorystoreInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMemorystoreInstanceList) DeepCopyInto(out *CloudMemorystoreInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudMemorystoreInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstanceList.
func (in *CloudMemorystoreInstanceList) DeepCopy() *CloudMemorystoreInstanceList {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudMemorystoreInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMemorystoreInstanceObservation) DeepCopyInto(out *CloudMemorystoreInstanceObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstanceObservation.
func (in *CloudMemorystoreInstanceObservation) DeepCopy() *CloudMemorystoreInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMemorystoreInstanceParameters) DeepCopyInto(out *CloudMemorystoreInstanceParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LocationID != nil {
		in, out := &in.LocationID, &out.LocationID
		*out = new(string)
		**out = **in
	}
	if in.AlternativeLocationID != nil {
		in, out := &in.AlternativeLocationID, &out.AlternativeLocationID
		*out = new(string)
		**out = **in
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReservedIPRange != nil {
		in, out := &in.ReservedIPRange, &out.ReservedIPRange
		*out = new(string)
		**out = **in
	}
	if in.RedisConfigs != nil {
		in, out := &in.RedisConfigs, &out.RedisConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AuthorizedNetwork != nil {
		in, out := &in.AuthorizedNetwork, &out.AuthorizedNetwork
		*out = new(string)
		**out = **in
	}
	if in.ConnectMode != nil {
		in, out := &in.ConnectMode, &out.ConnectMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstanceParameters.
func (in *CloudMemorystoreInstanceParameters) DeepCopy() *CloudMemorystoreInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMemorystoreInstanceSpec) DeepCopyInto(out *CloudMemorystoreInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstanceSpec.
func (in *CloudMemorystoreInstanceSpec) DeepCopy() *CloudMemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMemorystoreInstanceStatus) DeepCopyInto(out *CloudMemorystoreInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstanceStatus.
func (in *CloudMemorystoreInstanceStatus) DeepCopy() *CloudMemorystoreInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
