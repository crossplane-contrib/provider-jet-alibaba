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

type BackendServerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackendServerParameters struct {

	// +kubebuilder:validation:Optional
	BackendServers []BackendServersParameters `json:"backendServers,omitempty" tf:"backend_servers,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteProtectionValidation *bool `json:"deleteProtectionValidation,omitempty" tf:"delete_protection_validation,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancerID *string `json:"loadBalancerId" tf:"load_balancer_id,omitempty"`
}

type BackendServersObservation struct {
}

type BackendServersParameters struct {

	// +kubebuilder:validation:Required
	ServerID *string `json:"serverId" tf:"server_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServerIP *string `json:"serverIp,omitempty" tf:"server_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// BackendServerSpec defines the desired state of BackendServer
type BackendServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendServerParameters `json:"forProvider"`
}

// BackendServerStatus defines the observed state of BackendServer.
type BackendServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackendServer is the Schema for the BackendServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type BackendServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackendServerSpec   `json:"spec"`
	Status            BackendServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendServerList contains a list of BackendServers
type BackendServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendServer `json:"items"`
}

// Repository type metadata.
var (
	BackendServer_Kind             = "BackendServer"
	BackendServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendServer_Kind}.String()
	BackendServer_KindAPIVersion   = BackendServer_Kind + "." + CRDGroupVersion.String()
	BackendServer_GroupVersionKind = CRDGroupVersion.WithKind(BackendServer_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendServer{}, &BackendServerList{})
}
