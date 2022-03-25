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

type RouterInterfaceObservation struct {
	AccessPointID *string `json:"accessPointId,omitempty" tf:"access_point_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OppositeRouterInterfaceID *string `json:"oppositeRouterInterfaceId,omitempty" tf:"opposite_interface_id,omitempty"`

	OppositeRouterInterfaceOwnerID *string `json:"oppositeRouterInterfaceOwnerId,omitempty" tf:"opposite_interface_owner_id,omitempty"`

	OppositeRouterID *string `json:"oppositeRouterId,omitempty" tf:"opposite_router_id,omitempty"`

	OppositeRouterType *string `json:"oppositeRouterType,omitempty" tf:"opposite_router_type,omitempty"`
}

type RouterInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckSourceIP *string `json:"healthCheckSourceIp,omitempty" tf:"health_check_source_ip,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckTargetIP *string `json:"healthCheckTargetIp,omitempty" tf:"health_check_target_ip,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// +kubebuilder:validation:Optional
	OppositeAccessPointID *string `json:"oppositeAccessPointId,omitempty" tf:"opposite_access_point_id,omitempty"`

	// +kubebuilder:validation:Required
	OppositeRegion *string `json:"oppositeRegion" tf:"opposite_region,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	RouterID *string `json:"routerId" tf:"router_id,omitempty"`

	// +kubebuilder:validation:Required
	RouterType *string `json:"routerType" tf:"router_type,omitempty"`

	// +kubebuilder:validation:Optional
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`
}

// RouterInterfaceSpec defines the desired state of RouterInterface
type RouterInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterInterfaceParameters `json:"forProvider"`
}

// RouterInterfaceStatus defines the observed state of RouterInterface.
type RouterInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouterInterface is the Schema for the RouterInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type RouterInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterInterfaceSpec   `json:"spec"`
	Status            RouterInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterInterfaceList contains a list of RouterInterfaces
type RouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterInterface `json:"items"`
}

// Repository type metadata.
var (
	RouterInterface_Kind             = "RouterInterface"
	RouterInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterInterface_Kind}.String()
	RouterInterface_KindAPIVersion   = RouterInterface_Kind + "." + CRDGroupVersion.String()
	RouterInterface_GroupVersionKind = CRDGroupVersion.WithKind(RouterInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterInterface{}, &RouterInterfaceList{})
}
