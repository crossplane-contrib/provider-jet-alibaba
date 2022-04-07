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
func (in *BootstrapActionObservation) DeepCopyInto(out *BootstrapActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapActionObservation.
func (in *BootstrapActionObservation) DeepCopy() *BootstrapActionObservation {
	if in == nil {
		return nil
	}
	out := new(BootstrapActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapActionParameters) DeepCopyInto(out *BootstrapActionParameters) {
	*out = *in
	if in.Arg != nil {
		in, out := &in.Arg, &out.Arg
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapActionParameters.
func (in *BootstrapActionParameters) DeepCopy() *BootstrapActionParameters {
	if in == nil {
		return nil
	}
	out := new(BootstrapActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.BootstrapAction != nil {
		in, out := &in.BootstrapAction, &out.BootstrapAction
		*out = make([]BootstrapActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.ClusterType != nil {
		in, out := &in.ClusterType, &out.ClusterType
		*out = new(string)
		**out = **in
	}
	if in.DepositType != nil {
		in, out := &in.DepositType, &out.DepositType
		*out = new(string)
		**out = **in
	}
	if in.EasEnable != nil {
		in, out := &in.EasEnable, &out.EasEnable
		*out = new(bool)
		**out = **in
	}
	if in.EmrVer != nil {
		in, out := &in.EmrVer, &out.EmrVer
		*out = new(string)
		**out = **in
	}
	if in.HighAvailabilityEnable != nil {
		in, out := &in.HighAvailabilityEnable, &out.HighAvailabilityEnable
		*out = new(bool)
		**out = **in
	}
	if in.HostGroup != nil {
		in, out := &in.HostGroup, &out.HostGroup
		*out = make([]HostGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsOpenPublicIP != nil {
		in, out := &in.IsOpenPublicIP, &out.IsOpenPublicIP
		*out = new(bool)
		**out = **in
	}
	if in.KeyPairName != nil {
		in, out := &in.KeyPairName, &out.KeyPairName
		*out = new(string)
		**out = **in
	}
	if in.MasterPwd != nil {
		in, out := &in.MasterPwd, &out.MasterPwd
		*out = new(string)
		**out = **in
	}
	if in.OptionSoftwareList != nil {
		in, out := &in.OptionSoftwareList, &out.OptionSoftwareList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.RelatedClusterID != nil {
		in, out := &in.RelatedClusterID, &out.RelatedClusterID
		*out = new(string)
		**out = **in
	}
	if in.SSHEnable != nil {
		in, out := &in.SSHEnable, &out.SSHEnable
		*out = new(bool)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.UseLocalMetadb != nil {
		in, out := &in.UseLocalMetadb, &out.UseLocalMetadb
		*out = new(bool)
		**out = **in
	}
	if in.UserDefinedEmrEcsRole != nil {
		in, out := &in.UserDefinedEmrEcsRole, &out.UserDefinedEmrEcsRole
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostGroupObservation) DeepCopyInto(out *HostGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostGroupObservation.
func (in *HostGroupObservation) DeepCopy() *HostGroupObservation {
	if in == nil {
		return nil
	}
	out := new(HostGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostGroupParameters) DeepCopyInto(out *HostGroupParameters) {
	*out = *in
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.DiskCapacity != nil {
		in, out := &in.DiskCapacity, &out.DiskCapacity
		*out = new(string)
		**out = **in
	}
	if in.DiskCount != nil {
		in, out := &in.DiskCount, &out.DiskCount
		*out = new(string)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.GpuDriver != nil {
		in, out := &in.GpuDriver, &out.GpuDriver
		*out = new(string)
		**out = **in
	}
	if in.HostGroupName != nil {
		in, out := &in.HostGroupName, &out.HostGroupName
		*out = new(string)
		**out = **in
	}
	if in.HostGroupType != nil {
		in, out := &in.HostGroupType, &out.HostGroupType
		*out = new(string)
		**out = **in
	}
	if in.InstanceList != nil {
		in, out := &in.InstanceList, &out.InstanceList
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.SysDiskCapacity != nil {
		in, out := &in.SysDiskCapacity, &out.SysDiskCapacity
		*out = new(string)
		**out = **in
	}
	if in.SysDiskType != nil {
		in, out := &in.SysDiskType, &out.SysDiskType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostGroupParameters.
func (in *HostGroupParameters) DeepCopy() *HostGroupParameters {
	if in == nil {
		return nil
	}
	out := new(HostGroupParameters)
	in.DeepCopyInto(out)
	return out
}