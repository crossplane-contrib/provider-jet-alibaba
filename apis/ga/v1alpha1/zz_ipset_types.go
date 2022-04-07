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

type IPSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddressList []*string `json:"ipAddressList,omitempty" tf:"ip_address_list,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type IPSetParameters struct {

	// +kubebuilder:validation:Required
	AccelerateRegionID *string `json:"accelerateRegionId" tf:"accelerate_region_id,omitempty"`

	// +kubebuilder:validation:Required
	AcceleratorID *string `json:"acceleratorId" tf:"accelerator_id,omitempty"`

	// +kubebuilder:validation:Optional
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`
}

// IPSetSpec defines the desired state of IPSet
type IPSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPSetParameters `json:"forProvider"`
}

// IPSetStatus defines the observed state of IPSet.
type IPSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPSet is the Schema for the IPSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type IPSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPSetSpec   `json:"spec"`
	Status            IPSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPSetList contains a list of IPSets
type IPSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPSet `json:"items"`
}

// Repository type metadata.
var (
	IPSet_Kind             = "IPSet"
	IPSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPSet_Kind}.String()
	IPSet_KindAPIVersion   = IPSet_Kind + "." + CRDGroupVersion.String()
	IPSet_GroupVersionKind = CRDGroupVersion.WithKind(IPSet_Kind)
)

func init() {
	SchemeBuilder.Register(&IPSet{}, &IPSetList{})
}
