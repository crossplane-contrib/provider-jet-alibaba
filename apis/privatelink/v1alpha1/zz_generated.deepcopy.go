//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
func (in *VPCEndpoint) DeepCopyInto(out *VPCEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpoint.
func (in *VPCEndpoint) DeepCopy() *VPCEndpoint {
	if in == nil {
		return nil
	}
	out := new(VPCEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointConnection) DeepCopyInto(out *VPCEndpointConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointConnection.
func (in *VPCEndpointConnection) DeepCopy() *VPCEndpointConnection {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointConnectionList) DeepCopyInto(out *VPCEndpointConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCEndpointConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointConnectionList.
func (in *VPCEndpointConnectionList) DeepCopy() *VPCEndpointConnectionList {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointConnectionObservation) DeepCopyInto(out *VPCEndpointConnectionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointConnectionObservation.
func (in *VPCEndpointConnectionObservation) DeepCopy() *VPCEndpointConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointConnectionParameters) DeepCopyInto(out *VPCEndpointConnectionParameters) {
	*out = *in
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(float64)
		**out = **in
	}
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.EndpointID != nil {
		in, out := &in.EndpointID, &out.EndpointID
		*out = new(string)
		**out = **in
	}
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointConnectionParameters.
func (in *VPCEndpointConnectionParameters) DeepCopy() *VPCEndpointConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointConnectionSpec) DeepCopyInto(out *VPCEndpointConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointConnectionSpec.
func (in *VPCEndpointConnectionSpec) DeepCopy() *VPCEndpointConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointConnectionStatus) DeepCopyInto(out *VPCEndpointConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointConnectionStatus.
func (in *VPCEndpointConnectionStatus) DeepCopy() *VPCEndpointConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointList) DeepCopyInto(out *VPCEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointList.
func (in *VPCEndpointList) DeepCopy() *VPCEndpointList {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointObservation) DeepCopyInto(out *VPCEndpointObservation) {
	*out = *in
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(float64)
		**out = **in
	}
	if in.ConnectionStatus != nil {
		in, out := &in.ConnectionStatus, &out.ConnectionStatus
		*out = new(string)
		**out = **in
	}
	if in.EndpointBusinessStatus != nil {
		in, out := &in.EndpointBusinessStatus, &out.EndpointBusinessStatus
		*out = new(string)
		**out = **in
	}
	if in.EndpointDomain != nil {
		in, out := &in.EndpointDomain, &out.EndpointDomain
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointObservation.
func (in *VPCEndpointObservation) DeepCopy() *VPCEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointParameters) DeepCopyInto(out *VPCEndpointParameters) {
	*out = *in
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.EndpointDescription != nil {
		in, out := &in.EndpointDescription, &out.EndpointDescription
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.VPCEndpointName != nil {
		in, out := &in.VPCEndpointName, &out.VPCEndpointName
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointParameters.
func (in *VPCEndpointParameters) DeepCopy() *VPCEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointService) DeepCopyInto(out *VPCEndpointService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointService.
func (in *VPCEndpointService) DeepCopy() *VPCEndpointService {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceList) DeepCopyInto(out *VPCEndpointServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCEndpointService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceList.
func (in *VPCEndpointServiceList) DeepCopy() *VPCEndpointServiceList {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceObservation) DeepCopyInto(out *VPCEndpointServiceObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ServiceBusinessStatus != nil {
		in, out := &in.ServiceBusinessStatus, &out.ServiceBusinessStatus
		*out = new(string)
		**out = **in
	}
	if in.ServiceDomain != nil {
		in, out := &in.ServiceDomain, &out.ServiceDomain
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceObservation.
func (in *VPCEndpointServiceObservation) DeepCopy() *VPCEndpointServiceObservation {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceParameters) DeepCopyInto(out *VPCEndpointServiceParameters) {
	*out = *in
	if in.AutoAcceptConnection != nil {
		in, out := &in.AutoAcceptConnection, &out.AutoAcceptConnection
		*out = new(bool)
		**out = **in
	}
	if in.ConnectBandwidth != nil {
		in, out := &in.ConnectBandwidth, &out.ConnectBandwidth
		*out = new(float64)
		**out = **in
	}
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.Payer != nil {
		in, out := &in.Payer, &out.Payer
		*out = new(string)
		**out = **in
	}
	if in.ServiceDescription != nil {
		in, out := &in.ServiceDescription, &out.ServiceDescription
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceParameters.
func (in *VPCEndpointServiceParameters) DeepCopy() *VPCEndpointServiceParameters {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceResource) DeepCopyInto(out *VPCEndpointServiceResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceResource.
func (in *VPCEndpointServiceResource) DeepCopy() *VPCEndpointServiceResource {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointServiceResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceResourceList) DeepCopyInto(out *VPCEndpointServiceResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCEndpointServiceResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceResourceList.
func (in *VPCEndpointServiceResourceList) DeepCopy() *VPCEndpointServiceResourceList {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointServiceResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceResourceObservation) DeepCopyInto(out *VPCEndpointServiceResourceObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceResourceObservation.
func (in *VPCEndpointServiceResourceObservation) DeepCopy() *VPCEndpointServiceResourceObservation {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceResourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceResourceParameters) DeepCopyInto(out *VPCEndpointServiceResourceParameters) {
	*out = *in
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceResourceParameters.
func (in *VPCEndpointServiceResourceParameters) DeepCopy() *VPCEndpointServiceResourceParameters {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceResourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceResourceSpec) DeepCopyInto(out *VPCEndpointServiceResourceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceResourceSpec.
func (in *VPCEndpointServiceResourceSpec) DeepCopy() *VPCEndpointServiceResourceSpec {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceResourceStatus) DeepCopyInto(out *VPCEndpointServiceResourceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceResourceStatus.
func (in *VPCEndpointServiceResourceStatus) DeepCopy() *VPCEndpointServiceResourceStatus {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceSpec) DeepCopyInto(out *VPCEndpointServiceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceSpec.
func (in *VPCEndpointServiceSpec) DeepCopy() *VPCEndpointServiceSpec {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceStatus) DeepCopyInto(out *VPCEndpointServiceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceStatus.
func (in *VPCEndpointServiceStatus) DeepCopy() *VPCEndpointServiceStatus {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceUser) DeepCopyInto(out *VPCEndpointServiceUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceUser.
func (in *VPCEndpointServiceUser) DeepCopy() *VPCEndpointServiceUser {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointServiceUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceUserList) DeepCopyInto(out *VPCEndpointServiceUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCEndpointServiceUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceUserList.
func (in *VPCEndpointServiceUserList) DeepCopy() *VPCEndpointServiceUserList {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointServiceUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceUserObservation) DeepCopyInto(out *VPCEndpointServiceUserObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceUserObservation.
func (in *VPCEndpointServiceUserObservation) DeepCopy() *VPCEndpointServiceUserObservation {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceUserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceUserParameters) DeepCopyInto(out *VPCEndpointServiceUserParameters) {
	*out = *in
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceUserParameters.
func (in *VPCEndpointServiceUserParameters) DeepCopy() *VPCEndpointServiceUserParameters {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceUserSpec) DeepCopyInto(out *VPCEndpointServiceUserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceUserSpec.
func (in *VPCEndpointServiceUserSpec) DeepCopy() *VPCEndpointServiceUserSpec {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointServiceUserStatus) DeepCopyInto(out *VPCEndpointServiceUserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointServiceUserStatus.
func (in *VPCEndpointServiceUserStatus) DeepCopy() *VPCEndpointServiceUserStatus {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointServiceUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointSpec) DeepCopyInto(out *VPCEndpointSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointSpec.
func (in *VPCEndpointSpec) DeepCopy() *VPCEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointStatus) DeepCopyInto(out *VPCEndpointStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointStatus.
func (in *VPCEndpointStatus) DeepCopy() *VPCEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointZone) DeepCopyInto(out *VPCEndpointZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointZone.
func (in *VPCEndpointZone) DeepCopy() *VPCEndpointZone {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointZoneList) DeepCopyInto(out *VPCEndpointZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCEndpointZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointZoneList.
func (in *VPCEndpointZoneList) DeepCopy() *VPCEndpointZoneList {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCEndpointZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointZoneObservation) DeepCopyInto(out *VPCEndpointZoneObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointZoneObservation.
func (in *VPCEndpointZoneObservation) DeepCopy() *VPCEndpointZoneObservation {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointZoneObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointZoneParameters) DeepCopyInto(out *VPCEndpointZoneParameters) {
	*out = *in
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.EndpointID != nil {
		in, out := &in.EndpointID, &out.EndpointID
		*out = new(string)
		**out = **in
	}
	if in.VswitchID != nil {
		in, out := &in.VswitchID, &out.VswitchID
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointZoneParameters.
func (in *VPCEndpointZoneParameters) DeepCopy() *VPCEndpointZoneParameters {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointZoneParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointZoneSpec) DeepCopyInto(out *VPCEndpointZoneSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointZoneSpec.
func (in *VPCEndpointZoneSpec) DeepCopy() *VPCEndpointZoneSpec {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCEndpointZoneStatus) DeepCopyInto(out *VPCEndpointZoneStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCEndpointZoneStatus.
func (in *VPCEndpointZoneStatus) DeepCopy() *VPCEndpointZoneStatus {
	if in == nil {
		return nil
	}
	out := new(VPCEndpointZoneStatus)
	in.DeepCopyInto(out)
	return out
}