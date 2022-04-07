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

type ManagerResourceGroupObservation struct {
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RegionStatuses []RegionStatusesObservation `json:"regionStatuses,omitempty" tf:"region_statuses,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ManagerResourceGroupParameters struct {

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type RegionStatusesObservation struct {
}

type RegionStatusesParameters struct {

	// +kubebuilder:validation:Required
	RegionID *string `json:"regionId" tf:"region_id,omitempty"`

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

// ManagerResourceGroupSpec defines the desired state of ManagerResourceGroup
type ManagerResourceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagerResourceGroupParameters `json:"forProvider"`
}

// ManagerResourceGroupStatus defines the observed state of ManagerResourceGroup.
type ManagerResourceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagerResourceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerResourceGroup is the Schema for the ManagerResourceGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ManagerResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerResourceGroupSpec   `json:"spec"`
	Status            ManagerResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerResourceGroupList contains a list of ManagerResourceGroups
type ManagerResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagerResourceGroup `json:"items"`
}

// Repository type metadata.
var (
	ManagerResourceGroup_Kind             = "ManagerResourceGroup"
	ManagerResourceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagerResourceGroup_Kind}.String()
	ManagerResourceGroup_KindAPIVersion   = ManagerResourceGroup_Kind + "." + CRDGroupVersion.String()
	ManagerResourceGroup_GroupVersionKind = CRDGroupVersion.WithKind(ManagerResourceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagerResourceGroup{}, &ManagerResourceGroupList{})
}
