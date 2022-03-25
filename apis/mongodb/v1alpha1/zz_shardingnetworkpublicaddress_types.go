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

type ShardingNetworkPublicAddressNetworkAddressObservation struct {
}

type ShardingNetworkPublicAddressNetworkAddressParameters struct {

	// +kubebuilder:validation:Required
	ExpiredTime *string `json:"expiredTime" tf:"expired_time,omitempty"`

	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Required
	NetworkAddress *string `json:"networkAddress" tf:"network_address,omitempty"`

	// +kubebuilder:validation:Required
	NetworkType *string `json:"networkType" tf:"network_type,omitempty"`

	// +kubebuilder:validation:Required
	NodeID *string `json:"nodeId" tf:"node_id,omitempty"`

	// +kubebuilder:validation:Required
	NodeType *string `json:"nodeType" tf:"node_type,omitempty"`

	// +kubebuilder:validation:Required
	Port *string `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Required
	VswitchID *string `json:"vswitchId" tf:"vswitch_id,omitempty"`
}

type ShardingNetworkPublicAddressObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NetworkAddress []ShardingNetworkPublicAddressNetworkAddressObservation `json:"networkAddress,omitempty" tf:"network_address,omitempty"`
}

type ShardingNetworkPublicAddressParameters struct {

	// +kubebuilder:validation:Required
	DBInstanceID *string `json:"dbInstanceId" tf:"db_instance_id,omitempty"`

	// +kubebuilder:validation:Required
	NodeID *string `json:"nodeId" tf:"node_id,omitempty"`
}

// ShardingNetworkPublicAddressSpec defines the desired state of ShardingNetworkPublicAddress
type ShardingNetworkPublicAddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShardingNetworkPublicAddressParameters `json:"forProvider"`
}

// ShardingNetworkPublicAddressStatus defines the observed state of ShardingNetworkPublicAddress.
type ShardingNetworkPublicAddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShardingNetworkPublicAddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ShardingNetworkPublicAddress is the Schema for the ShardingNetworkPublicAddresss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ShardingNetworkPublicAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShardingNetworkPublicAddressSpec   `json:"spec"`
	Status            ShardingNetworkPublicAddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShardingNetworkPublicAddressList contains a list of ShardingNetworkPublicAddresss
type ShardingNetworkPublicAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShardingNetworkPublicAddress `json:"items"`
}

// Repository type metadata.
var (
	ShardingNetworkPublicAddress_Kind             = "ShardingNetworkPublicAddress"
	ShardingNetworkPublicAddress_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ShardingNetworkPublicAddress_Kind}.String()
	ShardingNetworkPublicAddress_KindAPIVersion   = ShardingNetworkPublicAddress_Kind + "." + CRDGroupVersion.String()
	ShardingNetworkPublicAddress_GroupVersionKind = CRDGroupVersion.WithKind(ShardingNetworkPublicAddress_Kind)
)

func init() {
	SchemeBuilder.Register(&ShardingNetworkPublicAddress{}, &ShardingNetworkPublicAddressList{})
}
