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

type ACLObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ACLParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EgressACLEntries []EgressACLEntriesParameters `json:"egressAclEntries,omitempty" tf:"egress_acl_entries,omitempty"`

	// +kubebuilder:validation:Optional
	IngressACLEntries []IngressACLEntriesParameters `json:"ingressAclEntries,omitempty" tf:"ingress_acl_entries,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkACLName *string `json:"networkAclName,omitempty" tf:"network_acl_name,omitempty"`

	// +kubebuilder:validation:Optional
	Resources []ResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

type EgressACLEntriesObservation struct {
}

type EgressACLEntriesParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationCidrIP *string `json:"destinationCidrIp,omitempty" tf:"destination_cidr_ip,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkACLEntryName *string `json:"networkAclEntryName,omitempty" tf:"network_acl_entry_name,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type IngressACLEntriesObservation struct {
}

type IngressACLEntriesParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkACLEntryName *string `json:"networkAclEntryName,omitempty" tf:"network_acl_entry_name,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SourceCidrIP *string `json:"sourceCidrIp,omitempty" tf:"source_cidr_ip,omitempty"`
}

type ResourcesObservation struct {
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

// ACLSpec defines the desired state of ACL
type ACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ACLParameters `json:"forProvider"`
}

// ACLStatus defines the observed state of ACL.
type ACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ACL is the Schema for the ACLs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ACLSpec   `json:"spec"`
	Status            ACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ACLList contains a list of ACLs
type ACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ACL `json:"items"`
}

// Repository type metadata.
var (
	ACL_Kind             = "ACL"
	ACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ACL_Kind}.String()
	ACL_KindAPIVersion   = ACL_Kind + "." + CRDGroupVersion.String()
	ACL_GroupVersionKind = CRDGroupVersion.WithKind(ACL_Kind)
)

func init() {
	SchemeBuilder.Register(&ACL{}, &ACLList{})
}
