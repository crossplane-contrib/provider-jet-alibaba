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

type CenterGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CenterGroupParameters struct {

	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`
}

// CenterGroupSpec defines the desired state of CenterGroup
type CenterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CenterGroupParameters `json:"forProvider"`
}

// CenterGroupStatus defines the observed state of CenterGroup.
type CenterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CenterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CenterGroup is the Schema for the CenterGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type CenterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CenterGroupSpec   `json:"spec"`
	Status            CenterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CenterGroupList contains a list of CenterGroups
type CenterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CenterGroup `json:"items"`
}

// Repository type metadata.
var (
	CenterGroup_Kind             = "CenterGroup"
	CenterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CenterGroup_Kind}.String()
	CenterGroup_KindAPIVersion   = CenterGroup_Kind + "." + CRDGroupVersion.String()
	CenterGroup_GroupVersionKind = CRDGroupVersion.WithKind(CenterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&CenterGroup{}, &CenterGroupList{})
}
