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

type GroupRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GroupRuleParameters struct {

	// +kubebuilder:validation:Optional
	CidrIP *string `json:"cidrIp,omitempty" tf:"cidr_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// +kubebuilder:validation:Optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// +kubebuilder:validation:Optional
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	SecurityGroupID *string `json:"securityGroupId" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceGroupOwnerAccount *string `json:"sourceGroupOwnerAccount,omitempty" tf:"source_group_owner_account,omitempty"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupID *string `json:"sourceSecurityGroupId,omitempty" tf:"source_security_group_id,omitempty"`

	// Type of rule, ingress (inbound) or egress (outbound).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// GroupRuleSpec defines the desired state of GroupRule
type GroupRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupRuleParameters `json:"forProvider"`
}

// GroupRuleStatus defines the observed state of GroupRule.
type GroupRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupRule is the Schema for the GroupRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type GroupRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupRuleSpec   `json:"spec"`
	Status            GroupRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupRuleList contains a list of GroupRules
type GroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupRule `json:"items"`
}

// Repository type metadata.
var (
	GroupRule_Kind             = "GroupRule"
	GroupRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupRule_Kind}.String()
	GroupRule_KindAPIVersion   = GroupRule_Kind + "." + CRDGroupVersion.String()
	GroupRule_GroupVersionKind = CRDGroupVersion.WithKind(GroupRule_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupRule{}, &GroupRuleList{})
}