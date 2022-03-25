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

type TransitRouterPeerAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TransitRouterAttachmentID *string `json:"transitRouterAttachmentId,omitempty" tf:"transit_router_attachment_id,omitempty"`
}

type TransitRouterPeerAttachmentParameters struct {

	// +kubebuilder:validation:Optional
	AutoPublishRouteEnabled *bool `json:"autoPublishRouteEnabled,omitempty" tf:"auto_publish_route_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// +kubebuilder:validation:Optional
	BandwidthType *string `json:"bandwidthType,omitempty" tf:"bandwidth_type,omitempty"`

	// +kubebuilder:validation:Optional
	CenBandwidthPackageID *string `json:"cenBandwidthPackageId,omitempty" tf:"cen_bandwidth_package_id,omitempty"`

	// +kubebuilder:validation:Required
	CenID *string `json:"cenId" tf:"cen_id,omitempty"`

	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// +kubebuilder:validation:Required
	PeerTransitRouterID *string `json:"peerTransitRouterId" tf:"peer_transit_router_id,omitempty"`

	// +kubebuilder:validation:Required
	PeerTransitRouterRegionID *string `json:"peerTransitRouterRegionId" tf:"peer_transit_router_region_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// +kubebuilder:validation:Optional
	RouteTableAssociationEnabled *bool `json:"routeTableAssociationEnabled,omitempty" tf:"route_table_association_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RouteTablePropagationEnabled *bool `json:"routeTablePropagationEnabled,omitempty" tf:"route_table_propagation_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TransitRouterAttachmentDescription *string `json:"transitRouterAttachmentDescription,omitempty" tf:"transit_router_attachment_description,omitempty"`

	// +kubebuilder:validation:Optional
	TransitRouterAttachmentName *string `json:"transitRouterAttachmentName,omitempty" tf:"transit_router_attachment_name,omitempty"`

	// +kubebuilder:validation:Optional
	TransitRouterID *string `json:"transitRouterId,omitempty" tf:"transit_router_id,omitempty"`
}

// TransitRouterPeerAttachmentSpec defines the desired state of TransitRouterPeerAttachment
type TransitRouterPeerAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitRouterPeerAttachmentParameters `json:"forProvider"`
}

// TransitRouterPeerAttachmentStatus defines the observed state of TransitRouterPeerAttachment.
type TransitRouterPeerAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitRouterPeerAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitRouterPeerAttachment is the Schema for the TransitRouterPeerAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type TransitRouterPeerAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitRouterPeerAttachmentSpec   `json:"spec"`
	Status            TransitRouterPeerAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitRouterPeerAttachmentList contains a list of TransitRouterPeerAttachments
type TransitRouterPeerAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitRouterPeerAttachment `json:"items"`
}

// Repository type metadata.
var (
	TransitRouterPeerAttachment_Kind             = "TransitRouterPeerAttachment"
	TransitRouterPeerAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitRouterPeerAttachment_Kind}.String()
	TransitRouterPeerAttachment_KindAPIVersion   = TransitRouterPeerAttachment_Kind + "." + CRDGroupVersion.String()
	TransitRouterPeerAttachment_GroupVersionKind = CRDGroupVersion.WithKind(TransitRouterPeerAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitRouterPeerAttachment{}, &TransitRouterPeerAttachmentList{})
}