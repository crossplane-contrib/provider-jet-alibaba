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

type AlertContactObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AlertContactParameters struct {

	// +kubebuilder:validation:Optional
	AlertContactName *string `json:"alertContactName,omitempty" tf:"alert_contact_name,omitempty"`

	// +kubebuilder:validation:Optional
	DingRobotWebhookURL *string `json:"dingRobotWebhookUrl,omitempty" tf:"ding_robot_webhook_url,omitempty"`

	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	PhoneNum *string `json:"phoneNum,omitempty" tf:"phone_num,omitempty"`

	// +kubebuilder:validation:Optional
	SystemNoc *bool `json:"systemNoc,omitempty" tf:"system_noc,omitempty"`
}

// AlertContactSpec defines the desired state of AlertContact
type AlertContactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertContactParameters `json:"forProvider"`
}

// AlertContactStatus defines the observed state of AlertContact.
type AlertContactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertContactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlertContact is the Schema for the AlertContacts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type AlertContact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertContactSpec   `json:"spec"`
	Status            AlertContactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertContactList contains a list of AlertContacts
type AlertContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertContact `json:"items"`
}

// Repository type metadata.
var (
	AlertContact_Kind             = "AlertContact"
	AlertContact_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlertContact_Kind}.String()
	AlertContact_KindAPIVersion   = AlertContact_Kind + "." + CRDGroupVersion.String()
	AlertContact_GroupVersionKind = CRDGroupVersion.WithKind(AlertContact_Kind)
)

func init() {
	SchemeBuilder.Register(&AlertContact{}, &AlertContactList{})
}
