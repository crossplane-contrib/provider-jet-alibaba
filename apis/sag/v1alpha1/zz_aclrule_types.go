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

type ACLRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ACLRuleParameters struct {

	// +kubebuilder:validation:Required
	ACLID *string `json:"aclId" tf:"acl_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DestCidr *string `json:"destCidr" tf:"dest_cidr,omitempty"`

	// +kubebuilder:validation:Required
	DestPortRange *string `json:"destPortRange" tf:"dest_port_range,omitempty"`

	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	SourceCidr *string `json:"sourceCidr" tf:"source_cidr,omitempty"`

	// +kubebuilder:validation:Required
	SourcePortRange *string `json:"sourcePortRange" tf:"source_port_range,omitempty"`
}

// ACLRuleSpec defines the desired state of ACLRule
type ACLRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ACLRuleParameters `json:"forProvider"`
}

// ACLRuleStatus defines the observed state of ACLRule.
type ACLRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ACLRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ACLRule is the Schema for the ACLRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ACLRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ACLRuleSpec   `json:"spec"`
	Status            ACLRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ACLRuleList contains a list of ACLRules
type ACLRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ACLRule `json:"items"`
}

// Repository type metadata.
var (
	ACLRule_Kind             = "ACLRule"
	ACLRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ACLRule_Kind}.String()
	ACLRule_KindAPIVersion   = ACLRule_Kind + "." + CRDGroupVersion.String()
	ACLRule_GroupVersionKind = CRDGroupVersion.WithKind(ACLRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ACLRule{}, &ACLRuleList{})
}