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

type RouteEntryObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteEntryParameters struct {

	// +kubebuilder:validation:Required
	NextHop *string `json:"nextHop" tf:"next_hop,omitempty"`

	// +kubebuilder:validation:Required
	PublishVPC *bool `json:"publishVpc" tf:"publish_vpc,omitempty"`

	// +kubebuilder:validation:Required
	RouteDest *string `json:"routeDest" tf:"route_dest,omitempty"`

	// +kubebuilder:validation:Required
	VPNGatewayID *string `json:"vpnGatewayId" tf:"vpn_gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// RouteEntrySpec defines the desired state of RouteEntry
type RouteEntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteEntryParameters `json:"forProvider"`
}

// RouteEntryStatus defines the observed state of RouteEntry.
type RouteEntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteEntryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteEntry is the Schema for the RouteEntrys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type RouteEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteEntrySpec   `json:"spec"`
	Status            RouteEntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteEntryList contains a list of RouteEntrys
type RouteEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteEntry `json:"items"`
}

// Repository type metadata.
var (
	RouteEntry_Kind             = "RouteEntry"
	RouteEntry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteEntry_Kind}.String()
	RouteEntry_KindAPIVersion   = RouteEntry_Kind + "." + CRDGroupVersion.String()
	RouteEntry_GroupVersionKind = CRDGroupVersion.WithKind(RouteEntry_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteEntry{}, &RouteEntryList{})
}
