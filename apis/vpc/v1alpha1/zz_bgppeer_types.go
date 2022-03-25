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

type BGPPeerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BGPPeerParameters struct {

	// +kubebuilder:validation:Required
	BGPGroupID *string `json:"bgpGroupId" tf:"bgp_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	BfdMultiHop *float64 `json:"bfdMultiHop,omitempty" tf:"bfd_multi_hop,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBfd *bool `json:"enableBfd,omitempty" tf:"enable_bfd,omitempty"`

	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// +kubebuilder:validation:Optional
	PeerIPAddress *string `json:"peerIpAddress,omitempty" tf:"peer_ip_address,omitempty"`
}

// BGPPeerSpec defines the desired state of BGPPeer
type BGPPeerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BGPPeerParameters `json:"forProvider"`
}

// BGPPeerStatus defines the observed state of BGPPeer.
type BGPPeerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BGPPeerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BGPPeer is the Schema for the BGPPeers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type BGPPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BGPPeerSpec   `json:"spec"`
	Status            BGPPeerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BGPPeerList contains a list of BGPPeers
type BGPPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BGPPeer `json:"items"`
}

// Repository type metadata.
var (
	BGPPeer_Kind             = "BGPPeer"
	BGPPeer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BGPPeer_Kind}.String()
	BGPPeer_KindAPIVersion   = BGPPeer_Kind + "." + CRDGroupVersion.String()
	BGPPeer_GroupVersionKind = CRDGroupVersion.WithKind(BGPPeer_Kind)
)

func init() {
	SchemeBuilder.Register(&BGPPeer{}, &BGPPeerList{})
}
