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

type JobTemplateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type JobTemplateParameters struct {

	// +kubebuilder:validation:Optional
	ArrayRequest *string `json:"arrayRequest,omitempty" tf:"array_request,omitempty"`

	// +kubebuilder:validation:Optional
	ClockTime *string `json:"clockTime,omitempty" tf:"clock_time,omitempty"`

	// +kubebuilder:validation:Required
	CommandLine *string `json:"commandLine" tf:"command_line,omitempty"`

	// +kubebuilder:validation:Optional
	Gpu *float64 `json:"gpu,omitempty" tf:"gpu,omitempty"`

	// +kubebuilder:validation:Required
	JobTemplateName *string `json:"jobTemplateName" tf:"job_template_name,omitempty"`

	// +kubebuilder:validation:Optional
	Mem *string `json:"mem,omitempty" tf:"mem,omitempty"`

	// +kubebuilder:validation:Optional
	Node *float64 `json:"node,omitempty" tf:"node,omitempty"`

	// +kubebuilder:validation:Optional
	PackagePath *string `json:"packagePath,omitempty" tf:"package_path,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	Queue *string `json:"queue,omitempty" tf:"queue,omitempty"`

	// +kubebuilder:validation:Optional
	ReRunable *bool `json:"reRunable,omitempty" tf:"re_runable,omitempty"`

	// +kubebuilder:validation:Optional
	RunasUser *string `json:"runasUser,omitempty" tf:"runas_user,omitempty"`

	// +kubebuilder:validation:Optional
	StderrRedirectPath *string `json:"stderrRedirectPath,omitempty" tf:"stderr_redirect_path,omitempty"`

	// +kubebuilder:validation:Optional
	StdoutRedirectPath *string `json:"stdoutRedirectPath,omitempty" tf:"stdout_redirect_path,omitempty"`

	// +kubebuilder:validation:Optional
	Task *float64 `json:"task,omitempty" tf:"task,omitempty"`

	// +kubebuilder:validation:Optional
	Thread *float64 `json:"thread,omitempty" tf:"thread,omitempty"`

	// +kubebuilder:validation:Optional
	Variables *string `json:"variables,omitempty" tf:"variables,omitempty"`
}

// JobTemplateSpec defines the desired state of JobTemplate
type JobTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobTemplateParameters `json:"forProvider"`
}

// JobTemplateStatus defines the observed state of JobTemplate.
type JobTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JobTemplate is the Schema for the JobTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type JobTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobTemplateSpec   `json:"spec"`
	Status            JobTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobTemplateList contains a list of JobTemplates
type JobTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobTemplate `json:"items"`
}

// Repository type metadata.
var (
	JobTemplate_Kind             = "JobTemplate"
	JobTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobTemplate_Kind}.String()
	JobTemplate_KindAPIVersion   = JobTemplate_Kind + "." + CRDGroupVersion.String()
	JobTemplate_GroupVersionKind = CRDGroupVersion.WithKind(JobTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&JobTemplate{}, &JobTemplateList{})
}
