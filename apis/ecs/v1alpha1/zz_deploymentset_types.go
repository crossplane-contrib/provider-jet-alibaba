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

type DeploymentSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DeploymentSetParameters struct {

	// +kubebuilder:validation:Optional
	DeploymentSetName *string `json:"deploymentSetName,omitempty" tf:"deployment_set_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	Granularity *string `json:"granularity,omitempty" tf:"granularity,omitempty"`

	// +kubebuilder:validation:Optional
	OnUnableToRedeployFailedInstance *string `json:"onUnableToRedeployFailedInstance,omitempty" tf:"on_unable_to_redeploy_failed_instance,omitempty"`

	// +kubebuilder:validation:Optional
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

// DeploymentSetSpec defines the desired state of DeploymentSet
type DeploymentSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentSetParameters `json:"forProvider"`
}

// DeploymentSetStatus defines the observed state of DeploymentSet.
type DeploymentSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentSet is the Schema for the DeploymentSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type DeploymentSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSetSpec   `json:"spec"`
	Status            DeploymentSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentSetList contains a list of DeploymentSets
type DeploymentSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentSet `json:"items"`
}

// Repository type metadata.
var (
	DeploymentSet_Kind             = "DeploymentSet"
	DeploymentSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeploymentSet_Kind}.String()
	DeploymentSet_KindAPIVersion   = DeploymentSet_Kind + "." + CRDGroupVersion.String()
	DeploymentSet_GroupVersionKind = CRDGroupVersion.WithKind(DeploymentSet_Kind)
)

func init() {
	SchemeBuilder.Register(&DeploymentSet{}, &DeploymentSetList{})
}