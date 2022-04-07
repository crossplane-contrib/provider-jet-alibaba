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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfigObservation) DeepCopyInto(out *AuthConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfigObservation.
func (in *AuthConfigObservation) DeepCopy() *AuthConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AuthConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfigParameters) DeepCopyInto(out *AuthConfigParameters) {
	*out = *in
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(string)
		**out = **in
	}
	if in.MasterKey != nil {
		in, out := &in.MasterKey, &out.MasterKey
		*out = new(string)
		**out = **in
	}
	if in.SlaveKey != nil {
		in, out := &in.SlaveKey, &out.SlaveKey
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfigParameters.
func (in *AuthConfigParameters) DeepCopy() *AuthConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AuthConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheConfigObservation) DeepCopyInto(out *CacheConfigObservation) {
	*out = *in
	if in.CacheID != nil {
		in, out := &in.CacheID, &out.CacheID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheConfigObservation.
func (in *CacheConfigObservation) DeepCopy() *CacheConfigObservation {
	if in == nil {
		return nil
	}
	out := new(CacheConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheConfigParameters) DeepCopyInto(out *CacheConfigParameters) {
	*out = *in
	if in.CacheContent != nil {
		in, out := &in.CacheContent, &out.CacheContent
		*out = new(string)
		**out = **in
	}
	if in.CacheType != nil {
		in, out := &in.CacheType, &out.CacheType
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheConfigParameters.
func (in *CacheConfigParameters) DeepCopy() *CacheConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CacheConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateConfigObservation) DeepCopyInto(out *CertificateConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateConfigObservation.
func (in *CertificateConfigObservation) DeepCopy() *CertificateConfigObservation {
	if in == nil {
		return nil
	}
	out := new(CertificateConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateConfigParameters) DeepCopyInto(out *CertificateConfigParameters) {
	*out = *in
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServerCertificateSecretRef != nil {
		in, out := &in.ServerCertificateSecretRef, &out.ServerCertificateSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServerCertificateStatus != nil {
		in, out := &in.ServerCertificateStatus, &out.ServerCertificateStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateConfigParameters.
func (in *CertificateConfigParameters) DeepCopy() *CertificateConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CertificateConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfig) DeepCopyInto(out *DomainConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfig.
func (in *DomainConfig) DeepCopy() *DomainConfig {
	if in == nil {
		return nil
	}
	out := new(DomainConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigList) DeepCopyInto(out *DomainConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigList.
func (in *DomainConfigList) DeepCopy() *DomainConfigList {
	if in == nil {
		return nil
	}
	out := new(DomainConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigObservation) DeepCopyInto(out *DomainConfigObservation) {
	*out = *in
	if in.ConfigID != nil {
		in, out := &in.ConfigID, &out.ConfigID
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigObservation.
func (in *DomainConfigObservation) DeepCopy() *DomainConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DomainConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigParameters) DeepCopyInto(out *DomainConfigParameters) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.FunctionArgs != nil {
		in, out := &in.FunctionArgs, &out.FunctionArgs
		*out = make([]FunctionArgsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigParameters.
func (in *DomainConfigParameters) DeepCopy() *DomainConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DomainConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigSpec) DeepCopyInto(out *DomainConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigSpec.
func (in *DomainConfigSpec) DeepCopy() *DomainConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DomainConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigStatus) DeepCopyInto(out *DomainConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigStatus.
func (in *DomainConfigStatus) DeepCopy() *DomainConfigStatus {
	if in == nil {
		return nil
	}
	out := new(DomainConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNew) DeepCopyInto(out *DomainNew) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNew.
func (in *DomainNew) DeepCopy() *DomainNew {
	if in == nil {
		return nil
	}
	out := new(DomainNew)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainNew) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewCertificateConfigObservation) DeepCopyInto(out *DomainNewCertificateConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewCertificateConfigObservation.
func (in *DomainNewCertificateConfigObservation) DeepCopy() *DomainNewCertificateConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DomainNewCertificateConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewCertificateConfigParameters) DeepCopyInto(out *DomainNewCertificateConfigParameters) {
	*out = *in
	if in.CertName != nil {
		in, out := &in.CertName, &out.CertName
		*out = new(string)
		**out = **in
	}
	if in.CertType != nil {
		in, out := &in.CertType, &out.CertType
		*out = new(string)
		**out = **in
	}
	if in.ForceSet != nil {
		in, out := &in.ForceSet, &out.ForceSet
		*out = new(string)
		**out = **in
	}
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServerCertificateSecretRef != nil {
		in, out := &in.ServerCertificateSecretRef, &out.ServerCertificateSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServerCertificateStatus != nil {
		in, out := &in.ServerCertificateStatus, &out.ServerCertificateStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewCertificateConfigParameters.
func (in *DomainNewCertificateConfigParameters) DeepCopy() *DomainNewCertificateConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DomainNewCertificateConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewList) DeepCopyInto(out *DomainNewList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainNew, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewList.
func (in *DomainNewList) DeepCopy() *DomainNewList {
	if in == nil {
		return nil
	}
	out := new(DomainNewList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainNewList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewObservation) DeepCopyInto(out *DomainNewObservation) {
	*out = *in
	if in.Cname != nil {
		in, out := &in.Cname, &out.Cname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewObservation.
func (in *DomainNewObservation) DeepCopy() *DomainNewObservation {
	if in == nil {
		return nil
	}
	out := new(DomainNewObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewParameters) DeepCopyInto(out *DomainNewParameters) {
	*out = *in
	if in.CdnType != nil {
		in, out := &in.CdnType, &out.CdnType
		*out = new(string)
		**out = **in
	}
	if in.CertificateConfig != nil {
		in, out := &in.CertificateConfig, &out.CertificateConfig
		*out = make([]DomainNewCertificateConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]SourcesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewParameters.
func (in *DomainNewParameters) DeepCopy() *DomainNewParameters {
	if in == nil {
		return nil
	}
	out := new(DomainNewParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewSpec) DeepCopyInto(out *DomainNewSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewSpec.
func (in *DomainNewSpec) DeepCopy() *DomainNewSpec {
	if in == nil {
		return nil
	}
	out := new(DomainNewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewStatus) DeepCopyInto(out *DomainNewStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewStatus.
func (in *DomainNewStatus) DeepCopy() *DomainNewStatus {
	if in == nil {
		return nil
	}
	out := new(DomainNewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainObservation) DeepCopyInto(out *DomainObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainObservation.
func (in *DomainObservation) DeepCopy() *DomainObservation {
	if in == nil {
		return nil
	}
	out := new(DomainObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainParameters) DeepCopyInto(out *DomainParameters) {
	*out = *in
	if in.AuthConfig != nil {
		in, out := &in.AuthConfig, &out.AuthConfig
		*out = make([]AuthConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BlockIps != nil {
		in, out := &in.BlockIps, &out.BlockIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CacheConfig != nil {
		in, out := &in.CacheConfig, &out.CacheConfig
		*out = make([]CacheConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CdnType != nil {
		in, out := &in.CdnType, &out.CdnType
		*out = new(string)
		**out = **in
	}
	if in.CertificateConfig != nil {
		in, out := &in.CertificateConfig, &out.CertificateConfig
		*out = make([]CertificateConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.HTTPHeaderConfig != nil {
		in, out := &in.HTTPHeaderConfig, &out.HTTPHeaderConfig
		*out = make([]HTTPHeaderConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OptimizeEnable != nil {
		in, out := &in.OptimizeEnable, &out.OptimizeEnable
		*out = new(string)
		**out = **in
	}
	if in.Page404Config != nil {
		in, out := &in.Page404Config, &out.Page404Config
		*out = make([]Page404ConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PageCompressEnable != nil {
		in, out := &in.PageCompressEnable, &out.PageCompressEnable
		*out = new(string)
		**out = **in
	}
	if in.ParameterFilterConfig != nil {
		in, out := &in.ParameterFilterConfig, &out.ParameterFilterConfig
		*out = make([]ParameterFilterConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RangeEnable != nil {
		in, out := &in.RangeEnable, &out.RangeEnable
		*out = new(string)
		**out = **in
	}
	if in.ReferConfig != nil {
		in, out := &in.ReferConfig, &out.ReferConfig
		*out = make([]ReferConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.SourcePort != nil {
		in, out := &in.SourcePort, &out.SourcePort
		*out = new(float64)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VideoSeekEnable != nil {
		in, out := &in.VideoSeekEnable, &out.VideoSeekEnable
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainParameters.
func (in *DomainParameters) DeepCopy() *DomainParameters {
	if in == nil {
		return nil
	}
	out := new(DomainParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionArgsObservation) DeepCopyInto(out *FunctionArgsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionArgsObservation.
func (in *FunctionArgsObservation) DeepCopy() *FunctionArgsObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionArgsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionArgsParameters) DeepCopyInto(out *FunctionArgsParameters) {
	*out = *in
	if in.ArgName != nil {
		in, out := &in.ArgName, &out.ArgName
		*out = new(string)
		**out = **in
	}
	if in.ArgValue != nil {
		in, out := &in.ArgValue, &out.ArgValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionArgsParameters.
func (in *FunctionArgsParameters) DeepCopy() *FunctionArgsParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionArgsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPHeaderConfigObservation) DeepCopyInto(out *HTTPHeaderConfigObservation) {
	*out = *in
	if in.HeaderID != nil {
		in, out := &in.HeaderID, &out.HeaderID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeaderConfigObservation.
func (in *HTTPHeaderConfigObservation) DeepCopy() *HTTPHeaderConfigObservation {
	if in == nil {
		return nil
	}
	out := new(HTTPHeaderConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPHeaderConfigParameters) DeepCopyInto(out *HTTPHeaderConfigParameters) {
	*out = *in
	if in.HeaderKey != nil {
		in, out := &in.HeaderKey, &out.HeaderKey
		*out = new(string)
		**out = **in
	}
	if in.HeaderValue != nil {
		in, out := &in.HeaderValue, &out.HeaderValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeaderConfigParameters.
func (in *HTTPHeaderConfigParameters) DeepCopy() *HTTPHeaderConfigParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPHeaderConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Page404ConfigObservation) DeepCopyInto(out *Page404ConfigObservation) {
	*out = *in
	if in.ErrorCode != nil {
		in, out := &in.ErrorCode, &out.ErrorCode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Page404ConfigObservation.
func (in *Page404ConfigObservation) DeepCopy() *Page404ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(Page404ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Page404ConfigParameters) DeepCopyInto(out *Page404ConfigParameters) {
	*out = *in
	if in.CustomPageURL != nil {
		in, out := &in.CustomPageURL, &out.CustomPageURL
		*out = new(string)
		**out = **in
	}
	if in.PageType != nil {
		in, out := &in.PageType, &out.PageType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Page404ConfigParameters.
func (in *Page404ConfigParameters) DeepCopy() *Page404ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(Page404ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterFilterConfigObservation) DeepCopyInto(out *ParameterFilterConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterFilterConfigObservation.
func (in *ParameterFilterConfigObservation) DeepCopy() *ParameterFilterConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ParameterFilterConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterFilterConfigParameters) DeepCopyInto(out *ParameterFilterConfigParameters) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(string)
		**out = **in
	}
	if in.HashKeyArgs != nil {
		in, out := &in.HashKeyArgs, &out.HashKeyArgs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterFilterConfigParameters.
func (in *ParameterFilterConfigParameters) DeepCopy() *ParameterFilterConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ParameterFilterConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDelivery) DeepCopyInto(out *RealTimeLogDelivery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDelivery.
func (in *RealTimeLogDelivery) DeepCopy() *RealTimeLogDelivery {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDelivery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RealTimeLogDelivery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliveryList) DeepCopyInto(out *RealTimeLogDeliveryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RealTimeLogDelivery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliveryList.
func (in *RealTimeLogDeliveryList) DeepCopy() *RealTimeLogDeliveryList {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliveryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RealTimeLogDeliveryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliveryObservation) DeepCopyInto(out *RealTimeLogDeliveryObservation) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliveryObservation.
func (in *RealTimeLogDeliveryObservation) DeepCopy() *RealTimeLogDeliveryObservation {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliveryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliveryParameters) DeepCopyInto(out *RealTimeLogDeliveryParameters) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Logstore != nil {
		in, out := &in.Logstore, &out.Logstore
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.SlsRegion != nil {
		in, out := &in.SlsRegion, &out.SlsRegion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliveryParameters.
func (in *RealTimeLogDeliveryParameters) DeepCopy() *RealTimeLogDeliveryParameters {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliveryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliverySpec) DeepCopyInto(out *RealTimeLogDeliverySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliverySpec.
func (in *RealTimeLogDeliverySpec) DeepCopy() *RealTimeLogDeliverySpec {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliveryStatus) DeepCopyInto(out *RealTimeLogDeliveryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliveryStatus.
func (in *RealTimeLogDeliveryStatus) DeepCopy() *RealTimeLogDeliveryStatus {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferConfigObservation) DeepCopyInto(out *ReferConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferConfigObservation.
func (in *ReferConfigObservation) DeepCopy() *ReferConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ReferConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferConfigParameters) DeepCopyInto(out *ReferConfigParameters) {
	*out = *in
	if in.AllowEmpty != nil {
		in, out := &in.AllowEmpty, &out.AllowEmpty
		*out = new(string)
		**out = **in
	}
	if in.ReferList != nil {
		in, out := &in.ReferList, &out.ReferList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ReferType != nil {
		in, out := &in.ReferType, &out.ReferType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferConfigParameters.
func (in *ReferConfigParameters) DeepCopy() *ReferConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ReferConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcesObservation) DeepCopyInto(out *SourcesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcesObservation.
func (in *SourcesObservation) DeepCopy() *SourcesObservation {
	if in == nil {
		return nil
	}
	out := new(SourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcesParameters) DeepCopyInto(out *SourcesParameters) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcesParameters.
func (in *SourcesParameters) DeepCopy() *SourcesParameters {
	if in == nil {
		return nil
	}
	out := new(SourcesParameters)
	in.DeepCopyInto(out)
	return out
}
