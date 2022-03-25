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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SAMLIdentityProviderConfigurationObservation struct {
}

type SAMLIdentityProviderConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	EncodedMetadataDocumentSecretRef *v1.SecretKeySelector `json:"encodedMetadataDocumentSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SsoStatus *string `json:"ssoStatus,omitempty" tf:"sso_status,omitempty"`
}

type SsoDirectoryObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SsoDirectoryParameters struct {

	// +kubebuilder:validation:Optional
	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name,omitempty"`

	// +kubebuilder:validation:Optional
	MfaAuthenticationStatus *string `json:"mfaAuthenticationStatus,omitempty" tf:"mfa_authentication_status,omitempty"`

	// +kubebuilder:validation:Optional
	SAMLIdentityProviderConfiguration []SAMLIdentityProviderConfigurationParameters `json:"samlIdentityProviderConfiguration,omitempty" tf:"saml_identity_provider_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	ScimSynchronizationStatus *string `json:"scimSynchronizationStatus,omitempty" tf:"scim_synchronization_status,omitempty"`
}

// SsoDirectorySpec defines the desired state of SsoDirectory
type SsoDirectorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SsoDirectoryParameters `json:"forProvider"`
}

// SsoDirectoryStatus defines the observed state of SsoDirectory.
type SsoDirectoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SsoDirectoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsoDirectory is the Schema for the SsoDirectorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type SsoDirectory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsoDirectorySpec   `json:"spec"`
	Status            SsoDirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsoDirectoryList contains a list of SsoDirectorys
type SsoDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsoDirectory `json:"items"`
}

// Repository type metadata.
var (
	SsoDirectory_Kind             = "SsoDirectory"
	SsoDirectory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SsoDirectory_Kind}.String()
	SsoDirectory_KindAPIVersion   = SsoDirectory_Kind + "." + CRDGroupVersion.String()
	SsoDirectory_GroupVersionKind = CRDGroupVersion.WithKind(SsoDirectory_Kind)
)

func init() {
	SchemeBuilder.Register(&SsoDirectory{}, &SsoDirectoryList{})
}
