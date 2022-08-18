//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The BMCGO Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalNode) DeepCopyInto(out *BareMetalNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalNode.
func (in *BareMetalNode) DeepCopy() *BareMetalNode {
	if in == nil {
		return nil
	}
	out := new(BareMetalNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalNodeList) DeepCopyInto(out *BareMetalNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BareMetalNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalNodeList.
func (in *BareMetalNodeList) DeepCopy() *BareMetalNodeList {
	if in == nil {
		return nil
	}
	out := new(BareMetalNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalNodeSpec) DeepCopyInto(out *BareMetalNodeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalNodeSpec.
func (in *BareMetalNodeSpec) DeepCopy() *BareMetalNodeSpec {
	if in == nil {
		return nil
	}
	out := new(BareMetalNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalNodeStatus) DeepCopyInto(out *BareMetalNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalNodeStatus.
func (in *BareMetalNodeStatus) DeepCopy() *BareMetalNodeStatus {
	if in == nil {
		return nil
	}
	out := new(BareMetalNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredSystem) DeepCopyInto(out *DiscoveredSystem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredSystem.
func (in *DiscoveredSystem) DeepCopy() *DiscoveredSystem {
	if in == nil {
		return nil
	}
	out := new(DiscoveredSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPMIEndpoint) DeepCopyInto(out *IPMIEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPMIEndpoint.
func (in *IPMIEndpoint) DeepCopy() *IPMIEndpoint {
	if in == nil {
		return nil
	}
	out := new(IPMIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPMIEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPMIEndpointList) DeepCopyInto(out *IPMIEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPMIEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPMIEndpointList.
func (in *IPMIEndpointList) DeepCopy() *IPMIEndpointList {
	if in == nil {
		return nil
	}
	out := new(IPMIEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPMIEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPMIEndpointSpec) DeepCopyInto(out *IPMIEndpointSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPMIEndpointSpec.
func (in *IPMIEndpointSpec) DeepCopy() *IPMIEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(IPMIEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPMIEndpointStatus) DeepCopyInto(out *IPMIEndpointStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPMIEndpointStatus.
func (in *IPMIEndpointStatus) DeepCopy() *IPMIEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(IPMIEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedfishEndpoint) DeepCopyInto(out *RedfishEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedfishEndpoint.
func (in *RedfishEndpoint) DeepCopy() *RedfishEndpoint {
	if in == nil {
		return nil
	}
	out := new(RedfishEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedfishEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedfishEndpointList) DeepCopyInto(out *RedfishEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedfishEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedfishEndpointList.
func (in *RedfishEndpointList) DeepCopy() *RedfishEndpointList {
	if in == nil {
		return nil
	}
	out := new(RedfishEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedfishEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedfishEndpointSpec) DeepCopyInto(out *RedfishEndpointSpec) {
	*out = *in
	out.BasicAuth = in.BasicAuth
	out.RedfishSessionAuth = in.RedfishSessionAuth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedfishEndpointSpec.
func (in *RedfishEndpointSpec) DeepCopy() *RedfishEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(RedfishEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedfishEndpointStatus) DeepCopyInto(out *RedfishEndpointStatus) {
	*out = *in
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
	if in.DiscoveredSystems != nil {
		in, out := &in.DiscoveredSystems, &out.DiscoveredSystems
		*out = make([]DiscoveredSystem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedfishEndpointStatus.
func (in *RedfishEndpointStatus) DeepCopy() *RedfishEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(RedfishEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedfishSessionAuth) DeepCopyInto(out *RedfishSessionAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedfishSessionAuth.
func (in *RedfishSessionAuth) DeepCopy() *RedfishSessionAuth {
	if in == nil {
		return nil
	}
	out := new(RedfishSessionAuth)
	in.DeepCopyInto(out)
	return out
}
