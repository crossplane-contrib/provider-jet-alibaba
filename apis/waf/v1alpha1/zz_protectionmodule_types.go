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

type ProtectionModuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProtectionModuleParameters struct {

	// +kubebuilder:validation:Required
	DefenseType *string `json:"defenseType" tf:"defense_type,omitempty"`

	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Required
	Mode *float64 `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`
}

// ProtectionModuleSpec defines the desired state of ProtectionModule
type ProtectionModuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtectionModuleParameters `json:"forProvider"`
}

// ProtectionModuleStatus defines the observed state of ProtectionModule.
type ProtectionModuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtectionModuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectionModule is the Schema for the ProtectionModules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ProtectionModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtectionModuleSpec   `json:"spec"`
	Status            ProtectionModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectionModuleList contains a list of ProtectionModules
type ProtectionModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtectionModule `json:"items"`
}

// Repository type metadata.
var (
	ProtectionModule_Kind             = "ProtectionModule"
	ProtectionModule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtectionModule_Kind}.String()
	ProtectionModule_KindAPIVersion   = ProtectionModule_Kind + "." + CRDGroupVersion.String()
	ProtectionModule_GroupVersionKind = CRDGroupVersion.WithKind(ProtectionModule_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtectionModule{}, &ProtectionModuleList{})
}