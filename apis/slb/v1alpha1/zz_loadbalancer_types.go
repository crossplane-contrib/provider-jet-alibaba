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

type LoadBalancerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoadBalancerParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	AddressIPVersion *string `json:"addressIpVersion,omitempty" tf:"address_ip_version,omitempty"`

	// +kubebuilder:validation:Optional
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// +kubebuilder:validation:Optional
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteProtection *string `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// +kubebuilder:validation:Optional
	Internet *bool `json:"internet,omitempty" tf:"internet,omitempty"`

	// +kubebuilder:validation:Optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerSpec *string `json:"loadBalancerSpec,omitempty" tf:"load_balancer_spec,omitempty"`

	// +kubebuilder:validation:Optional
	MasterZoneID *string `json:"masterZoneId,omitempty" tf:"master_zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	ModificationProtectionReason *string `json:"modificationProtectionReason,omitempty" tf:"modification_protection_reason,omitempty"`

	// +kubebuilder:validation:Optional
	ModificationProtectionStatus *string `json:"modificationProtectionStatus,omitempty" tf:"modification_protection_status,omitempty"`

	// +kubebuilder:validation:Optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SlaveZoneID *string `json:"slaveZoneId,omitempty" tf:"slave_zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`
}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerParameters `json:"forProvider"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer.
type LoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancer is the Schema for the LoadBalancers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerSpec   `json:"spec"`
	Status            LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancer_Kind             = "LoadBalancer"
	LoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancer_Kind}.String()
	LoadBalancer_KindAPIVersion   = LoadBalancer_Kind + "." + CRDGroupVersion.String()
	LoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
