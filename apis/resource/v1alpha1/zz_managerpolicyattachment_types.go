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

type ManagerPolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagerPolicyAttachmentParameters struct {

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Required
	PolicyType *string `json:"policyType" tf:"policy_type,omitempty"`

	// +kubebuilder:validation:Required
	PrincipalName *string `json:"principalName" tf:"principal_name,omitempty"`

	// +kubebuilder:validation:Required
	PrincipalType *string `json:"principalType" tf:"principal_type,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupID *string `json:"resourceGroupId" tf:"resource_group_id,omitempty"`
}

// ManagerPolicyAttachmentSpec defines the desired state of ManagerPolicyAttachment
type ManagerPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagerPolicyAttachmentParameters `json:"forProvider"`
}

// ManagerPolicyAttachmentStatus defines the observed state of ManagerPolicyAttachment.
type ManagerPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagerPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerPolicyAttachment is the Schema for the ManagerPolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ManagerPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerPolicyAttachmentSpec   `json:"spec"`
	Status            ManagerPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerPolicyAttachmentList contains a list of ManagerPolicyAttachments
type ManagerPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagerPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	ManagerPolicyAttachment_Kind             = "ManagerPolicyAttachment"
	ManagerPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagerPolicyAttachment_Kind}.String()
	ManagerPolicyAttachment_KindAPIVersion   = ManagerPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	ManagerPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(ManagerPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagerPolicyAttachment{}, &ManagerPolicyAttachmentList{})
}