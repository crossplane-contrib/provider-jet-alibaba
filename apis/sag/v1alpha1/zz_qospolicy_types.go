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

type QosPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QosPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DestCidr *string `json:"destCidr" tf:"dest_cidr,omitempty"`

	// +kubebuilder:validation:Required
	DestPortRange *string `json:"destPortRange" tf:"dest_port_range,omitempty"`

	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	QosID *string `json:"qosId" tf:"qos_id,omitempty"`

	// +kubebuilder:validation:Required
	SourceCidr *string `json:"sourceCidr" tf:"source_cidr,omitempty"`

	// +kubebuilder:validation:Required
	SourcePortRange *string `json:"sourcePortRange" tf:"source_port_range,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

// QosPolicySpec defines the desired state of QosPolicy
type QosPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QosPolicyParameters `json:"forProvider"`
}

// QosPolicyStatus defines the observed state of QosPolicy.
type QosPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QosPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QosPolicy is the Schema for the QosPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type QosPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QosPolicySpec   `json:"spec"`
	Status            QosPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QosPolicyList contains a list of QosPolicys
type QosPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QosPolicy `json:"items"`
}

// Repository type metadata.
var (
	QosPolicy_Kind             = "QosPolicy"
	QosPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QosPolicy_Kind}.String()
	QosPolicy_KindAPIVersion   = QosPolicy_Kind + "." + CRDGroupVersion.String()
	QosPolicy_GroupVersionKind = CRDGroupVersion.WithKind(QosPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&QosPolicy{}, &QosPolicyList{})
}