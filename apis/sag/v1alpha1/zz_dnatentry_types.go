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

type DnatEntryObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DnatEntryParameters struct {

	// +kubebuilder:validation:Optional
	ExternalIP *string `json:"externalIp,omitempty" tf:"external_ip,omitempty"`

	// +kubebuilder:validation:Required
	ExternalPort *string `json:"externalPort" tf:"external_port,omitempty"`

	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// +kubebuilder:validation:Required
	InternalIP *string `json:"internalIp" tf:"internal_ip,omitempty"`

	// +kubebuilder:validation:Required
	InternalPort *string `json:"internalPort" tf:"internal_port,omitempty"`

	// +kubebuilder:validation:Required
	SagID *string `json:"sagId" tf:"sag_id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// DnatEntrySpec defines the desired state of DnatEntry
type DnatEntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DnatEntryParameters `json:"forProvider"`
}

// DnatEntryStatus defines the observed state of DnatEntry.
type DnatEntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DnatEntryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DnatEntry is the Schema for the DnatEntrys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type DnatEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnatEntrySpec   `json:"spec"`
	Status            DnatEntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DnatEntryList contains a list of DnatEntrys
type DnatEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnatEntry `json:"items"`
}

// Repository type metadata.
var (
	DnatEntry_Kind             = "DnatEntry"
	DnatEntry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DnatEntry_Kind}.String()
	DnatEntry_KindAPIVersion   = DnatEntry_Kind + "." + CRDGroupVersion.String()
	DnatEntry_GroupVersionKind = CRDGroupVersion.WithKind(DnatEntry_Kind)
)

func init() {
	SchemeBuilder.Register(&DnatEntry{}, &DnatEntryList{})
}
