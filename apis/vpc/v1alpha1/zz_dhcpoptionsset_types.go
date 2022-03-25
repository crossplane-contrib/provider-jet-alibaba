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

type AssociateVpcsObservation struct {
	AssociateStatus *string `json:"associateStatus,omitempty" tf:"associate_status,omitempty"`
}

type AssociateVpcsParameters struct {

	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DHCPOptionsSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DHCPOptionsSetParameters struct {

	// +kubebuilder:validation:Optional
	AssociateVpcs []AssociateVpcsParameters `json:"associateVpcs,omitempty" tf:"associate_vpcs,omitempty"`

	// +kubebuilder:validation:Optional
	DHCPOptionsSetDescription *string `json:"dhcpOptionsSetDescription,omitempty" tf:"dhcp_options_set_description,omitempty"`

	// +kubebuilder:validation:Optional
	DHCPOptionsSetName *string `json:"dhcpOptionsSetName,omitempty" tf:"dhcp_options_set_name,omitempty"`

	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	DomainNameServers *string `json:"domainNameServers,omitempty" tf:"domain_name_servers,omitempty"`

	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`
}

// DHCPOptionsSetSpec defines the desired state of DHCPOptionsSet
type DHCPOptionsSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DHCPOptionsSetParameters `json:"forProvider"`
}

// DHCPOptionsSetStatus defines the observed state of DHCPOptionsSet.
type DHCPOptionsSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DHCPOptionsSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DHCPOptionsSet is the Schema for the DHCPOptionsSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type DHCPOptionsSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DHCPOptionsSetSpec   `json:"spec"`
	Status            DHCPOptionsSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DHCPOptionsSetList contains a list of DHCPOptionsSets
type DHCPOptionsSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DHCPOptionsSet `json:"items"`
}

// Repository type metadata.
var (
	DHCPOptionsSet_Kind             = "DHCPOptionsSet"
	DHCPOptionsSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DHCPOptionsSet_Kind}.String()
	DHCPOptionsSet_KindAPIVersion   = DHCPOptionsSet_Kind + "." + CRDGroupVersion.String()
	DHCPOptionsSet_GroupVersionKind = CRDGroupVersion.WithKind(DHCPOptionsSet_Kind)
)

func init() {
	SchemeBuilder.Register(&DHCPOptionsSet{}, &DHCPOptionsSetList{})
}
