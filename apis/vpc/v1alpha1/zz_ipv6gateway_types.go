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

type IPv6GatewayObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type IPv6GatewayParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6GatewayName *string `json:"ipv6GatewayName,omitempty" tf:"ipv6_gateway_name,omitempty"`

	// +kubebuilder:validation:Optional
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

// IPv6GatewaySpec defines the desired state of IPv6Gateway
type IPv6GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPv6GatewayParameters `json:"forProvider"`
}

// IPv6GatewayStatus defines the observed state of IPv6Gateway.
type IPv6GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPv6GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPv6Gateway is the Schema for the IPv6Gateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type IPv6Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPv6GatewaySpec   `json:"spec"`
	Status            IPv6GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPv6GatewayList contains a list of IPv6Gateways
type IPv6GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPv6Gateway `json:"items"`
}

// Repository type metadata.
var (
	IPv6Gateway_Kind             = "IPv6Gateway"
	IPv6Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPv6Gateway_Kind}.String()
	IPv6Gateway_KindAPIVersion   = IPv6Gateway_Kind + "." + CRDGroupVersion.String()
	IPv6Gateway_GroupVersionKind = CRDGroupVersion.WithKind(IPv6Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&IPv6Gateway{}, &IPv6GatewayList{})
}