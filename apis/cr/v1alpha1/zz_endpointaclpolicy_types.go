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

type EndpointACLPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointACLPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	EndpointType *string `json:"endpointType" tf:"endpoint_type,omitempty"`

	// +kubebuilder:validation:Required
	Entry *string `json:"entry" tf:"entry,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	ModuleName *string `json:"moduleName,omitempty" tf:"module_name,omitempty"`
}

// EndpointACLPolicySpec defines the desired state of EndpointACLPolicy
type EndpointACLPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointACLPolicyParameters `json:"forProvider"`
}

// EndpointACLPolicyStatus defines the observed state of EndpointACLPolicy.
type EndpointACLPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointACLPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointACLPolicy is the Schema for the EndpointACLPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type EndpointACLPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointACLPolicySpec   `json:"spec"`
	Status            EndpointACLPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointACLPolicyList contains a list of EndpointACLPolicys
type EndpointACLPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointACLPolicy `json:"items"`
}

// Repository type metadata.
var (
	EndpointACLPolicy_Kind             = "EndpointACLPolicy"
	EndpointACLPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointACLPolicy_Kind}.String()
	EndpointACLPolicy_KindAPIVersion   = EndpointACLPolicy_Kind + "." + CRDGroupVersion.String()
	EndpointACLPolicy_GroupVersionKind = CRDGroupVersion.WithKind(EndpointACLPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointACLPolicy{}, &EndpointACLPolicyList{})
}