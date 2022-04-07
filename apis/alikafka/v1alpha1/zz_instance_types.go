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

type InstanceObservation struct {
	EndPoint *string `json:"endPoint,omitempty" tf:"end_point,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type InstanceParameters struct {

	// +kubebuilder:validation:Optional
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Required
	DeployType *float64 `json:"deployType" tf:"deploy_type,omitempty"`

	// +kubebuilder:validation:Required
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Required
	DiskType *float64 `json:"diskType" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	EIPMax *float64 `json:"eipMax,omitempty" tf:"eip_max,omitempty"`

	// +kubebuilder:validation:Required
	IoMax *float64 `json:"ioMax" tf:"io_max,omitempty"`

	// +kubebuilder:validation:Optional
	PaidType *string `json:"paidType,omitempty" tf:"paid_type,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroup *string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceVersion *string `json:"serviceVersion,omitempty" tf:"service_version,omitempty"`

	// +kubebuilder:validation:Optional
	SpecType *string `json:"specType,omitempty" tf:"spec_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TopicQuota *float64 `json:"topicQuota" tf:"topic_quota,omitempty"`

	// +kubebuilder:validation:Required
	VswitchID *string `json:"vswitchId" tf:"vswitch_id,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
