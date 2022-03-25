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

type FieldSearchObservation struct {
}

type FieldSearchParameters struct {

	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// +kubebuilder:validation:Optional
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAnalytics *bool `json:"enableAnalytics,omitempty" tf:"enable_analytics,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeChinese *bool `json:"includeChinese,omitempty" tf:"include_chinese,omitempty"`

	// +kubebuilder:validation:Optional
	JSONKeys []JSONKeysParameters `json:"jsonKeys,omitempty" tf:"json_keys,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FullTextObservation struct {
}

type FullTextParameters struct {

	// +kubebuilder:validation:Optional
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeChinese *bool `json:"includeChinese,omitempty" tf:"include_chinese,omitempty"`

	// +kubebuilder:validation:Optional
	Token *string `json:"token,omitempty" tf:"token,omitempty"`
}

type JSONKeysObservation struct {
}

type JSONKeysParameters struct {

	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// +kubebuilder:validation:Optional
	DocValue *bool `json:"docValue,omitempty" tf:"doc_value,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StoreIndexObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StoreIndexParameters struct {

	// +kubebuilder:validation:Optional
	FieldSearch []FieldSearchParameters `json:"fieldSearch,omitempty" tf:"field_search,omitempty"`

	// +kubebuilder:validation:Optional
	FullText []FullTextParameters `json:"fullText,omitempty" tf:"full_text,omitempty"`

	// +kubebuilder:validation:Required
	Logstore *string `json:"logstore" tf:"logstore,omitempty"`

	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`
}

// StoreIndexSpec defines the desired state of StoreIndex
type StoreIndexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StoreIndexParameters `json:"forProvider"`
}

// StoreIndexStatus defines the observed state of StoreIndex.
type StoreIndexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StoreIndexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoreIndex is the Schema for the StoreIndexs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type StoreIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoreIndexSpec   `json:"spec"`
	Status            StoreIndexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoreIndexList contains a list of StoreIndexs
type StoreIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoreIndex `json:"items"`
}

// Repository type metadata.
var (
	StoreIndex_Kind             = "StoreIndex"
	StoreIndex_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StoreIndex_Kind}.String()
	StoreIndex_KindAPIVersion   = StoreIndex_Kind + "." + CRDGroupVersion.String()
	StoreIndex_GroupVersionKind = CRDGroupVersion.WithKind(StoreIndex_Kind)
)

func init() {
	SchemeBuilder.Register(&StoreIndex{}, &StoreIndexList{})
}