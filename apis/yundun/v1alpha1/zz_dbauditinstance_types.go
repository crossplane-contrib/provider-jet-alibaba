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

type DbauditInstanceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DbauditInstanceParameters struct {

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Period *float64 `json:"period" tf:"period,omitempty"`

	// +kubebuilder:validation:Required
	PlanCode *string `json:"planCode" tf:"plan_code,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VswitchID *string `json:"vswitchId" tf:"vswitch_id,omitempty"`
}

// DbauditInstanceSpec defines the desired state of DbauditInstance
type DbauditInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DbauditInstanceParameters `json:"forProvider"`
}

// DbauditInstanceStatus defines the observed state of DbauditInstance.
type DbauditInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DbauditInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbauditInstance is the Schema for the DbauditInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type DbauditInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbauditInstanceSpec   `json:"spec"`
	Status            DbauditInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbauditInstanceList contains a list of DbauditInstances
type DbauditInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbauditInstance `json:"items"`
}

// Repository type metadata.
var (
	DbauditInstance_Kind             = "DbauditInstance"
	DbauditInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DbauditInstance_Kind}.String()
	DbauditInstance_KindAPIVersion   = DbauditInstance_Kind + "." + CRDGroupVersion.String()
	DbauditInstance_GroupVersionKind = CRDGroupVersion.WithKind(DbauditInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&DbauditInstance{}, &DbauditInstanceList{})
}