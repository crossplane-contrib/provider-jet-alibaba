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
func (in *BandwidthPackage) DeepCopyInto(out *BandwidthPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackage.
func (in *BandwidthPackage) DeepCopy() *BandwidthPackage {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachment) DeepCopyInto(out *BandwidthPackageAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachment.
func (in *BandwidthPackageAttachment) DeepCopy() *BandwidthPackageAttachment {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthPackageAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentList) DeepCopyInto(out *BandwidthPackageAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BandwidthPackageAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentList.
func (in *BandwidthPackageAttachmentList) DeepCopy() *BandwidthPackageAttachmentList {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthPackageAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentObservation) DeepCopyInto(out *BandwidthPackageAttachmentObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentObservation.
func (in *BandwidthPackageAttachmentObservation) DeepCopy() *BandwidthPackageAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentParameters) DeepCopyInto(out *BandwidthPackageAttachmentParameters) {
	*out = *in
	if in.BandwidthPackageID != nil {
		in, out := &in.BandwidthPackageID, &out.BandwidthPackageID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentParameters.
func (in *BandwidthPackageAttachmentParameters) DeepCopy() *BandwidthPackageAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentSpec) DeepCopyInto(out *BandwidthPackageAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentSpec.
func (in *BandwidthPackageAttachmentSpec) DeepCopy() *BandwidthPackageAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentStatus) DeepCopyInto(out *BandwidthPackageAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentStatus.
func (in *BandwidthPackageAttachmentStatus) DeepCopy() *BandwidthPackageAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageList) DeepCopyInto(out *BandwidthPackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BandwidthPackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageList.
func (in *BandwidthPackageList) DeepCopy() *BandwidthPackageList {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthPackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageObservation) DeepCopyInto(out *BandwidthPackageObservation) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageObservation.
func (in *BandwidthPackageObservation) DeepCopy() *BandwidthPackageObservation {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageParameters) DeepCopyInto(out *BandwidthPackageParameters) {
	*out = *in
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(string)
		**out = **in
	}
	if in.BandwidthPackageName != nil {
		in, out := &in.BandwidthPackageName, &out.BandwidthPackageName
		*out = new(string)
		**out = **in
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(string)
		**out = **in
	}
	if in.InternetChargeType != nil {
		in, out := &in.InternetChargeType, &out.InternetChargeType
		*out = new(string)
		**out = **in
	}
	if in.Isp != nil {
		in, out := &in.Isp, &out.Isp
		*out = new(string)
		**out = **in
	}
	if in.Ratio != nil {
		in, out := &in.Ratio, &out.Ratio
		*out = new(float64)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageParameters.
func (in *BandwidthPackageParameters) DeepCopy() *BandwidthPackageParameters {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageSpec) DeepCopyInto(out *BandwidthPackageSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageSpec.
func (in *BandwidthPackageSpec) DeepCopy() *BandwidthPackageSpec {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageStatus) DeepCopyInto(out *BandwidthPackageStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageStatus.
func (in *BandwidthPackageStatus) DeepCopy() *BandwidthPackageStatus {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageStatus)
	in.DeepCopyInto(out)
	return out
}
