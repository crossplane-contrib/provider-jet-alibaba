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

type ExecutionObservation struct {
	Counters *string `json:"counters,omitempty" tf:"counters,omitempty"`

	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	ExecutedBy *string `json:"executedBy,omitempty" tf:"executed_by,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsParent *bool `json:"isParent,omitempty" tf:"is_parent,omitempty"`

	Outputs *string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	RAMRole *string `json:"ramRole,omitempty" tf:"ram_role,omitempty"`

	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty" tf:"status_message,omitempty"`

	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	UpdateDate *string `json:"updateDate,omitempty" tf:"update_date,omitempty"`
}

type ExecutionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	LoopMode *string `json:"loopMode,omitempty" tf:"loop_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	ParentExecutionID *string `json:"parentExecutionId,omitempty" tf:"parent_execution_id,omitempty"`

	// +kubebuilder:validation:Optional
	SafetyCheck *string `json:"safetyCheck,omitempty" tf:"safety_check,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// +kubebuilder:validation:Required
	TemplateName *string `json:"templateName" tf:"template_name,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateVersion *string `json:"templateVersion,omitempty" tf:"template_version,omitempty"`
}

// ExecutionSpec defines the desired state of Execution
type ExecutionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExecutionParameters `json:"forProvider"`
}

// ExecutionStatus defines the observed state of Execution.
type ExecutionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExecutionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Execution is the Schema for the Executions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type Execution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExecutionSpec   `json:"spec"`
	Status            ExecutionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExecutionList contains a list of Executions
type ExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Execution `json:"items"`
}

// Repository type metadata.
var (
	Execution_Kind             = "Execution"
	Execution_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Execution_Kind}.String()
	Execution_KindAPIVersion   = Execution_Kind + "." + CRDGroupVersion.String()
	Execution_GroupVersionKind = CRDGroupVersion.WithKind(Execution_Kind)
)

func init() {
	SchemeBuilder.Register(&Execution{}, &ExecutionList{})
}