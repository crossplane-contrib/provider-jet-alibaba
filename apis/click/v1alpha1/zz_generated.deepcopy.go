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
func (in *DBClusterAccessWhiteListObservation) DeepCopyInto(out *DBClusterAccessWhiteListObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBClusterAccessWhiteListObservation.
func (in *DBClusterAccessWhiteListObservation) DeepCopy() *DBClusterAccessWhiteListObservation {
	if in == nil {
		return nil
	}
	out := new(DBClusterAccessWhiteListObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBClusterAccessWhiteListParameters) DeepCopyInto(out *DBClusterAccessWhiteListParameters) {
	*out = *in
	if in.DBClusterIPArrayAttribute != nil {
		in, out := &in.DBClusterIPArrayAttribute, &out.DBClusterIPArrayAttribute
		*out = new(string)
		**out = **in
	}
	if in.DBClusterIPArrayName != nil {
		in, out := &in.DBClusterIPArrayName, &out.DBClusterIPArrayName
		*out = new(string)
		**out = **in
	}
	if in.SecurityIPList != nil {
		in, out := &in.SecurityIPList, &out.SecurityIPList
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBClusterAccessWhiteListParameters.
func (in *DBClusterAccessWhiteListParameters) DeepCopy() *DBClusterAccessWhiteListParameters {
	if in == nil {
		return nil
	}
	out := new(DBClusterAccessWhiteListParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseAccount) DeepCopyInto(out *HouseAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseAccount.
func (in *HouseAccount) DeepCopy() *HouseAccount {
	if in == nil {
		return nil
	}
	out := new(HouseAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HouseAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseAccountList) DeepCopyInto(out *HouseAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HouseAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseAccountList.
func (in *HouseAccountList) DeepCopy() *HouseAccountList {
	if in == nil {
		return nil
	}
	out := new(HouseAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HouseAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseAccountObservation) DeepCopyInto(out *HouseAccountObservation) {
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseAccountObservation.
func (in *HouseAccountObservation) DeepCopy() *HouseAccountObservation {
	if in == nil {
		return nil
	}
	out := new(HouseAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseAccountParameters) DeepCopyInto(out *HouseAccountParameters) {
	*out = *in
	if in.AccountDescription != nil {
		in, out := &in.AccountDescription, &out.AccountDescription
		*out = new(string)
		**out = **in
	}
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.AccountPassword != nil {
		in, out := &in.AccountPassword, &out.AccountPassword
		*out = new(string)
		**out = **in
	}
	if in.DBClusterID != nil {
		in, out := &in.DBClusterID, &out.DBClusterID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseAccountParameters.
func (in *HouseAccountParameters) DeepCopy() *HouseAccountParameters {
	if in == nil {
		return nil
	}
	out := new(HouseAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseAccountSpec) DeepCopyInto(out *HouseAccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseAccountSpec.
func (in *HouseAccountSpec) DeepCopy() *HouseAccountSpec {
	if in == nil {
		return nil
	}
	out := new(HouseAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseAccountStatus) DeepCopyInto(out *HouseAccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseAccountStatus.
func (in *HouseAccountStatus) DeepCopy() *HouseAccountStatus {
	if in == nil {
		return nil
	}
	out := new(HouseAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseBackupPolicy) DeepCopyInto(out *HouseBackupPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseBackupPolicy.
func (in *HouseBackupPolicy) DeepCopy() *HouseBackupPolicy {
	if in == nil {
		return nil
	}
	out := new(HouseBackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HouseBackupPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseBackupPolicyList) DeepCopyInto(out *HouseBackupPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HouseBackupPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseBackupPolicyList.
func (in *HouseBackupPolicyList) DeepCopy() *HouseBackupPolicyList {
	if in == nil {
		return nil
	}
	out := new(HouseBackupPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HouseBackupPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseBackupPolicyObservation) DeepCopyInto(out *HouseBackupPolicyObservation) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseBackupPolicyObservation.
func (in *HouseBackupPolicyObservation) DeepCopy() *HouseBackupPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(HouseBackupPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseBackupPolicyParameters) DeepCopyInto(out *HouseBackupPolicyParameters) {
	*out = *in
	if in.BackupRetentionPeriod != nil {
		in, out := &in.BackupRetentionPeriod, &out.BackupRetentionPeriod
		*out = new(float64)
		**out = **in
	}
	if in.DBClusterID != nil {
		in, out := &in.DBClusterID, &out.DBClusterID
		*out = new(string)
		**out = **in
	}
	if in.PreferredBackupPeriod != nil {
		in, out := &in.PreferredBackupPeriod, &out.PreferredBackupPeriod
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreferredBackupTime != nil {
		in, out := &in.PreferredBackupTime, &out.PreferredBackupTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseBackupPolicyParameters.
func (in *HouseBackupPolicyParameters) DeepCopy() *HouseBackupPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(HouseBackupPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseBackupPolicySpec) DeepCopyInto(out *HouseBackupPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseBackupPolicySpec.
func (in *HouseBackupPolicySpec) DeepCopy() *HouseBackupPolicySpec {
	if in == nil {
		return nil
	}
	out := new(HouseBackupPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseBackupPolicyStatus) DeepCopyInto(out *HouseBackupPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseBackupPolicyStatus.
func (in *HouseBackupPolicyStatus) DeepCopy() *HouseBackupPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(HouseBackupPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseDBCluster) DeepCopyInto(out *HouseDBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseDBCluster.
func (in *HouseDBCluster) DeepCopy() *HouseDBCluster {
	if in == nil {
		return nil
	}
	out := new(HouseDBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HouseDBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseDBClusterList) DeepCopyInto(out *HouseDBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HouseDBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseDBClusterList.
func (in *HouseDBClusterList) DeepCopy() *HouseDBClusterList {
	if in == nil {
		return nil
	}
	out := new(HouseDBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HouseDBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseDBClusterObservation) DeepCopyInto(out *HouseDBClusterObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseDBClusterObservation.
func (in *HouseDBClusterObservation) DeepCopy() *HouseDBClusterObservation {
	if in == nil {
		return nil
	}
	out := new(HouseDBClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseDBClusterParameters) DeepCopyInto(out *HouseDBClusterParameters) {
	*out = *in
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.DBClusterAccessWhiteList != nil {
		in, out := &in.DBClusterAccessWhiteList, &out.DBClusterAccessWhiteList
		*out = make([]DBClusterAccessWhiteListParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DBClusterClass != nil {
		in, out := &in.DBClusterClass, &out.DBClusterClass
		*out = new(string)
		**out = **in
	}
	if in.DBClusterDescription != nil {
		in, out := &in.DBClusterDescription, &out.DBClusterDescription
		*out = new(string)
		**out = **in
	}
	if in.DBClusterNetworkType != nil {
		in, out := &in.DBClusterNetworkType, &out.DBClusterNetworkType
		*out = new(string)
		**out = **in
	}
	if in.DBClusterVersion != nil {
		in, out := &in.DBClusterVersion, &out.DBClusterVersion
		*out = new(string)
		**out = **in
	}
	if in.DBNodeGroupCount != nil {
		in, out := &in.DBNodeGroupCount, &out.DBNodeGroupCount
		*out = new(float64)
		**out = **in
	}
	if in.DBNodeStorage != nil {
		in, out := &in.DBNodeStorage, &out.DBNodeStorage
		*out = new(string)
		**out = **in
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = new(string)
		**out = **in
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.MaintainTime != nil {
		in, out := &in.MaintainTime, &out.MaintainTime
		*out = new(string)
		**out = **in
	}
	if in.PaymentType != nil {
		in, out := &in.PaymentType, &out.PaymentType
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.UsedTime != nil {
		in, out := &in.UsedTime, &out.UsedTime
		*out = new(string)
		**out = **in
	}
	if in.VswitchID != nil {
		in, out := &in.VswitchID, &out.VswitchID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseDBClusterParameters.
func (in *HouseDBClusterParameters) DeepCopy() *HouseDBClusterParameters {
	if in == nil {
		return nil
	}
	out := new(HouseDBClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseDBClusterSpec) DeepCopyInto(out *HouseDBClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseDBClusterSpec.
func (in *HouseDBClusterSpec) DeepCopy() *HouseDBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(HouseDBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HouseDBClusterStatus) DeepCopyInto(out *HouseDBClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HouseDBClusterStatus.
func (in *HouseDBClusterStatus) DeepCopy() *HouseDBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(HouseDBClusterStatus)
	in.DeepCopyInto(out)
	return out
}
