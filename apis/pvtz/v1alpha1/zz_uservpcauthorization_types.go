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

type UserVPCAuthorizationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UserVPCAuthorizationParameters struct {

	// +kubebuilder:validation:Optional
	AuthChannel *string `json:"authChannel,omitempty" tf:"auth_channel,omitempty"`

	// +kubebuilder:validation:Optional
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	// +kubebuilder:validation:Required
	AuthorizedUserID *string `json:"authorizedUserId" tf:"authorized_user_id,omitempty"`
}

// UserVPCAuthorizationSpec defines the desired state of UserVPCAuthorization
type UserVPCAuthorizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserVPCAuthorizationParameters `json:"forProvider"`
}

// UserVPCAuthorizationStatus defines the observed state of UserVPCAuthorization.
type UserVPCAuthorizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserVPCAuthorizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserVPCAuthorization is the Schema for the UserVPCAuthorizations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type UserVPCAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserVPCAuthorizationSpec   `json:"spec"`
	Status            UserVPCAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserVPCAuthorizationList contains a list of UserVPCAuthorizations
type UserVPCAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserVPCAuthorization `json:"items"`
}

// Repository type metadata.
var (
	UserVPCAuthorization_Kind             = "UserVPCAuthorization"
	UserVPCAuthorization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserVPCAuthorization_Kind}.String()
	UserVPCAuthorization_KindAPIVersion   = UserVPCAuthorization_Kind + "." + CRDGroupVersion.String()
	UserVPCAuthorization_GroupVersionKind = CRDGroupVersion.WithKind(UserVPCAuthorization_Kind)
)

func init() {
	SchemeBuilder.Register(&UserVPCAuthorization{}, &UserVPCAuthorizationList{})
}
