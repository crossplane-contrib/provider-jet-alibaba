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

type ConnectPhysicalConnectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConnectPhysicalConnectionParameters struct {

	// +kubebuilder:validation:Required
	AccessPointID *string `json:"accessPointId" tf:"access_point_id,omitempty"`

	// +kubebuilder:validation:Optional
	Bandwidth *string `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// +kubebuilder:validation:Optional
	CircuitCode *string `json:"circuitCode,omitempty" tf:"circuit_code,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	LineOperator *string `json:"lineOperator" tf:"line_operator,omitempty"`

	// +kubebuilder:validation:Required
	PeerLocation *string `json:"peerLocation" tf:"peer_location,omitempty"`

	// +kubebuilder:validation:Optional
	PhysicalConnectionName *string `json:"physicalConnectionName,omitempty" tf:"physical_connection_name,omitempty"`

	// +kubebuilder:validation:Optional
	PortType *string `json:"portType,omitempty" tf:"port_type,omitempty"`

	// +kubebuilder:validation:Optional
	RedundantPhysicalConnectionID *string `json:"redundantPhysicalConnectionId,omitempty" tf:"redundant_physical_connection_id,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ConnectPhysicalConnectionSpec defines the desired state of ConnectPhysicalConnection
type ConnectPhysicalConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectPhysicalConnectionParameters `json:"forProvider"`
}

// ConnectPhysicalConnectionStatus defines the observed state of ConnectPhysicalConnection.
type ConnectPhysicalConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectPhysicalConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectPhysicalConnection is the Schema for the ConnectPhysicalConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ConnectPhysicalConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectPhysicalConnectionSpec   `json:"spec"`
	Status            ConnectPhysicalConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectPhysicalConnectionList contains a list of ConnectPhysicalConnections
type ConnectPhysicalConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectPhysicalConnection `json:"items"`
}

// Repository type metadata.
var (
	ConnectPhysicalConnection_Kind             = "ConnectPhysicalConnection"
	ConnectPhysicalConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectPhysicalConnection_Kind}.String()
	ConnectPhysicalConnection_KindAPIVersion   = ConnectPhysicalConnection_Kind + "." + CRDGroupVersion.String()
	ConnectPhysicalConnection_GroupVersionKind = CRDGroupVersion.WithKind(ConnectPhysicalConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectPhysicalConnection{}, &ConnectPhysicalConnectionList{})
}
