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

type ScalingRuleObservation struct {
	Ari *string `json:"ari,omitempty" tf:"ari,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScalingRuleParameters struct {

	// +kubebuilder:validation:Optional
	AdjustmentType *string `json:"adjustmentType,omitempty" tf:"adjustment_type,omitempty"`

	// +kubebuilder:validation:Optional
	AdjustmentValue *float64 `json:"adjustmentValue,omitempty" tf:"adjustment_value,omitempty"`

	// +kubebuilder:validation:Optional
	Cooldown *float64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`

	// +kubebuilder:validation:Optional
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup,omitempty" tf:"estimated_instance_warmup,omitempty"`

	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Required
	ScalingGroupID *string `json:"scalingGroupId" tf:"scaling_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ScalingRuleName *string `json:"scalingRuleName,omitempty" tf:"scaling_rule_name,omitempty"`

	// +kubebuilder:validation:Optional
	ScalingRuleType *string `json:"scalingRuleType,omitempty" tf:"scaling_rule_type,omitempty"`

	// +kubebuilder:validation:Optional
	StepAdjustment []StepAdjustmentParameters `json:"stepAdjustment,omitempty" tf:"step_adjustment,omitempty"`

	// +kubebuilder:validation:Optional
	TargetValue *float64 `json:"targetValue,omitempty" tf:"target_value,omitempty"`
}

type StepAdjustmentObservation struct {
}

type StepAdjustmentParameters struct {

	// +kubebuilder:validation:Optional
	MetricIntervalLowerBound *string `json:"metricIntervalLowerBound,omitempty" tf:"metric_interval_lower_bound,omitempty"`

	// +kubebuilder:validation:Optional
	MetricIntervalUpperBound *string `json:"metricIntervalUpperBound,omitempty" tf:"metric_interval_upper_bound,omitempty"`

	// +kubebuilder:validation:Optional
	ScalingAdjustment *float64 `json:"scalingAdjustment,omitempty" tf:"scaling_adjustment,omitempty"`
}

// ScalingRuleSpec defines the desired state of ScalingRule
type ScalingRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScalingRuleParameters `json:"forProvider"`
}

// ScalingRuleStatus defines the observed state of ScalingRule.
type ScalingRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScalingRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScalingRule is the Schema for the ScalingRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ScalingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScalingRuleSpec   `json:"spec"`
	Status            ScalingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScalingRuleList contains a list of ScalingRules
type ScalingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScalingRule `json:"items"`
}

// Repository type metadata.
var (
	ScalingRule_Kind             = "ScalingRule"
	ScalingRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScalingRule_Kind}.String()
	ScalingRule_KindAPIVersion   = ScalingRule_Kind + "." + CRDGroupVersion.String()
	ScalingRule_GroupVersionKind = CRDGroupVersion.WithKind(ScalingRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ScalingRule{}, &ScalingRuleList{})
}
