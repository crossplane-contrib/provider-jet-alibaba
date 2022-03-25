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

type AccessGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AccessGroupParameters struct {

	// +kubebuilder:validation:Required
	AccessGroupName *string `json:"accessGroupName" tf:"access_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	NetworkType *string `json:"networkType" tf:"network_type,omitempty"`
}

// AccessGroupSpec defines the desired state of AccessGroup
type AccessGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessGroupParameters `json:"forProvider"`
}

// AccessGroupStatus defines the observed state of AccessGroup.
type AccessGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessGroup is the Schema for the AccessGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type AccessGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessGroupSpec   `json:"spec"`
	Status            AccessGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessGroupList contains a list of AccessGroups
type AccessGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessGroup `json:"items"`
}

// Repository type metadata.
var (
	AccessGroup_Kind             = "AccessGroup"
	AccessGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessGroup_Kind}.String()
	AccessGroup_KindAPIVersion   = AccessGroup_Kind + "." + CRDGroupVersion.String()
	AccessGroup_GroupVersionKind = CRDGroupVersion.WithKind(AccessGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessGroup{}, &AccessGroupList{})
}