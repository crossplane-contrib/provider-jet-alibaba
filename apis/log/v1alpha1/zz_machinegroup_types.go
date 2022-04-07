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

type MachineGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MachineGroupParameters struct {

	// +kubebuilder:validation:Required
	IdentifyList []*string `json:"identifyList" tf:"identify_list,omitempty"`

	// +kubebuilder:validation:Optional
	IdentifyType *string `json:"identifyType,omitempty" tf:"identify_type,omitempty"`

	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

// MachineGroupSpec defines the desired state of MachineGroup
type MachineGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MachineGroupParameters `json:"forProvider"`
}

// MachineGroupStatus defines the observed state of MachineGroup.
type MachineGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MachineGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MachineGroup is the Schema for the MachineGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type MachineGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MachineGroupSpec   `json:"spec"`
	Status            MachineGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MachineGroupList contains a list of MachineGroups
type MachineGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachineGroup `json:"items"`
}

// Repository type metadata.
var (
	MachineGroup_Kind             = "MachineGroup"
	MachineGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MachineGroup_Kind}.String()
	MachineGroup_KindAPIVersion   = MachineGroup_Kind + "." + CRDGroupVersion.String()
	MachineGroup_GroupVersionKind = CRDGroupVersion.WithKind(MachineGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&MachineGroup{}, &MachineGroupList{})
}
