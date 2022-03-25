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

type StorageGatewayGatewayObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type StorageGatewayGatewayParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayClass *string `json:"gatewayClass,omitempty" tf:"gateway_class,omitempty"`

	// +kubebuilder:validation:Required
	GatewayName *string `json:"gatewayName" tf:"gateway_name,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkBandwidth *float64 `json:"publicNetworkBandwidth,omitempty" tf:"public_network_bandwidth,omitempty"`

	// +kubebuilder:validation:Optional
	ReasonDetail *string `json:"reasonDetail,omitempty" tf:"reason_detail,omitempty"`

	// +kubebuilder:validation:Optional
	ReasonType *string `json:"reasonType,omitempty" tf:"reason_type,omitempty"`

	// +kubebuilder:validation:Optional
	ReleaseAfterExpiration *bool `json:"releaseAfterExpiration,omitempty" tf:"release_after_expiration,omitempty"`

	// +kubebuilder:validation:Required
	StorageBundleID *string `json:"storageBundleId" tf:"storage_bundle_id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`
}

// StorageGatewayGatewaySpec defines the desired state of StorageGatewayGateway
type StorageGatewayGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageGatewayGatewayParameters `json:"forProvider"`
}

// StorageGatewayGatewayStatus defines the observed state of StorageGatewayGateway.
type StorageGatewayGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageGatewayGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageGatewayGateway is the Schema for the StorageGatewayGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type StorageGatewayGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageGatewayGatewaySpec   `json:"spec"`
	Status            StorageGatewayGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageGatewayGatewayList contains a list of StorageGatewayGateways
type StorageGatewayGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageGatewayGateway `json:"items"`
}

// Repository type metadata.
var (
	StorageGatewayGateway_Kind             = "StorageGatewayGateway"
	StorageGatewayGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageGatewayGateway_Kind}.String()
	StorageGatewayGateway_KindAPIVersion   = StorageGatewayGateway_Kind + "." + CRDGroupVersion.String()
	StorageGatewayGateway_GroupVersionKind = CRDGroupVersion.WithKind(StorageGatewayGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageGatewayGateway{}, &StorageGatewayGatewayList{})
}