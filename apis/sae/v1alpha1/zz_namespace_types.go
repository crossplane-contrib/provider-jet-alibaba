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

type NamespaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NamespaceParameters struct {

	// +kubebuilder:validation:Optional
	NamespaceDescription *string `json:"namespaceDescription,omitempty" tf:"namespace_description,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceID *string `json:"namespaceId" tf:"namespace_id,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceName *string `json:"namespaceName" tf:"namespace_name,omitempty"`
}

// NamespaceSpec defines the desired state of Namespace
type NamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamespaceParameters `json:"forProvider"`
}

// NamespaceStatus defines the observed state of Namespace.
type NamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Namespace is the Schema for the Namespaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type Namespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespaceSpec   `json:"spec"`
	Status            NamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceList contains a list of Namespaces
type NamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Namespace `json:"items"`
}

// Repository type metadata.
var (
	Namespace_Kind             = "Namespace"
	Namespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Namespace_Kind}.String()
	Namespace_KindAPIVersion   = Namespace_Kind + "." + CRDGroupVersion.String()
	Namespace_GroupVersionKind = CRDGroupVersion.WithKind(Namespace_Kind)
)

func init() {
	SchemeBuilder.Register(&Namespace{}, &NamespaceList{})
}
