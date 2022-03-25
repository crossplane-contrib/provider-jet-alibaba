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

type FirewallControlPolicyOrderObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FirewallControlPolicyOrderParameters struct {

	// +kubebuilder:validation:Required
	ACLUUID *string `json:"aclUuid" tf:"acl_uuid,omitempty"`

	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

// FirewallControlPolicyOrderSpec defines the desired state of FirewallControlPolicyOrder
type FirewallControlPolicyOrderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallControlPolicyOrderParameters `json:"forProvider"`
}

// FirewallControlPolicyOrderStatus defines the observed state of FirewallControlPolicyOrder.
type FirewallControlPolicyOrderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallControlPolicyOrderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallControlPolicyOrder is the Schema for the FirewallControlPolicyOrders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type FirewallControlPolicyOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallControlPolicyOrderSpec   `json:"spec"`
	Status            FirewallControlPolicyOrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallControlPolicyOrderList contains a list of FirewallControlPolicyOrders
type FirewallControlPolicyOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallControlPolicyOrder `json:"items"`
}

// Repository type metadata.
var (
	FirewallControlPolicyOrder_Kind             = "FirewallControlPolicyOrder"
	FirewallControlPolicyOrder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallControlPolicyOrder_Kind}.String()
	FirewallControlPolicyOrder_KindAPIVersion   = FirewallControlPolicyOrder_Kind + "." + CRDGroupVersion.String()
	FirewallControlPolicyOrder_GroupVersionKind = CRDGroupVersion.WithKind(FirewallControlPolicyOrder_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallControlPolicyOrder{}, &FirewallControlPolicyOrderList{})
}
