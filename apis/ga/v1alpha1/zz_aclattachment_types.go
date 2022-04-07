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

type ACLAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ACLAttachmentParameters struct {

	// +kubebuilder:validation:Required
	ACLID *string `json:"aclId" tf:"acl_id,omitempty"`

	// +kubebuilder:validation:Required
	ACLType *string `json:"aclType" tf:"acl_type,omitempty"`

	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// +kubebuilder:validation:Required
	ListenerID *string `json:"listenerId" tf:"listener_id,omitempty"`
}

// ACLAttachmentSpec defines the desired state of ACLAttachment
type ACLAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ACLAttachmentParameters `json:"forProvider"`
}

// ACLAttachmentStatus defines the observed state of ACLAttachment.
type ACLAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ACLAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ACLAttachment is the Schema for the ACLAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ACLAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ACLAttachmentSpec   `json:"spec"`
	Status            ACLAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ACLAttachmentList contains a list of ACLAttachments
type ACLAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ACLAttachment `json:"items"`
}

// Repository type metadata.
var (
	ACLAttachment_Kind             = "ACLAttachment"
	ACLAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ACLAttachment_Kind}.String()
	ACLAttachment_KindAPIVersion   = ACLAttachment_Kind + "." + CRDGroupVersion.String()
	ACLAttachment_GroupVersionKind = CRDGroupVersion.WithKind(ACLAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&ACLAttachment{}, &ACLAttachmentList{})
}
