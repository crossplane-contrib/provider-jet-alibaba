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

type TrafficMirrorFilterIngressRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TrafficMirrorFilterIngressRuleID *string `json:"trafficMirrorFilterIngressRuleId,omitempty" tf:"traffic_mirror_filter_ingress_rule_id,omitempty"`
}

type TrafficMirrorFilterIngressRuleParameters struct {

	// +kubebuilder:validation:Required
	DestinationCidrBlock *string `json:"destinationCidrBlock" tf:"destination_cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	RuleAction *string `json:"ruleAction" tf:"rule_action,omitempty"`

	// +kubebuilder:validation:Required
	SourceCidrBlock *string `json:"sourceCidrBlock" tf:"source_cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// +kubebuilder:validation:Required
	TrafficMirrorFilterID *string `json:"trafficMirrorFilterId" tf:"traffic_mirror_filter_id,omitempty"`
}

// TrafficMirrorFilterIngressRuleSpec defines the desired state of TrafficMirrorFilterIngressRule
type TrafficMirrorFilterIngressRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficMirrorFilterIngressRuleParameters `json:"forProvider"`
}

// TrafficMirrorFilterIngressRuleStatus defines the observed state of TrafficMirrorFilterIngressRule.
type TrafficMirrorFilterIngressRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficMirrorFilterIngressRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorFilterIngressRule is the Schema for the TrafficMirrorFilterIngressRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type TrafficMirrorFilterIngressRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficMirrorFilterIngressRuleSpec   `json:"spec"`
	Status            TrafficMirrorFilterIngressRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorFilterIngressRuleList contains a list of TrafficMirrorFilterIngressRules
type TrafficMirrorFilterIngressRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficMirrorFilterIngressRule `json:"items"`
}

// Repository type metadata.
var (
	TrafficMirrorFilterIngressRule_Kind             = "TrafficMirrorFilterIngressRule"
	TrafficMirrorFilterIngressRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficMirrorFilterIngressRule_Kind}.String()
	TrafficMirrorFilterIngressRule_KindAPIVersion   = TrafficMirrorFilterIngressRule_Kind + "." + CRDGroupVersion.String()
	TrafficMirrorFilterIngressRule_GroupVersionKind = CRDGroupVersion.WithKind(TrafficMirrorFilterIngressRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficMirrorFilterIngressRule{}, &TrafficMirrorFilterIngressRuleList{})
}
