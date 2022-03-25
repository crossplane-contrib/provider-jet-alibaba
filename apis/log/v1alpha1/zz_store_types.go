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

type EncryptConfObservation struct {
}

type EncryptConfParameters struct {

	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptType *string `json:"encryptType,omitempty" tf:"encrypt_type,omitempty"`

	// +kubebuilder:validation:Optional
	UserCmkInfo []UserCmkInfoParameters `json:"userCmkInfo,omitempty" tf:"user_cmk_info,omitempty"`
}

type ShardsObservation struct {
}

type ShardsParameters struct {

	// +kubebuilder:validation:Required
	BeginKey *string `json:"beginKey" tf:"begin_key,omitempty"`

	// +kubebuilder:validation:Required
	EndKey *string `json:"endKey" tf:"end_key,omitempty"`

	// +kubebuilder:validation:Required
	ID *float64 `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type StoreObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Shards []ShardsObservation `json:"shards,omitempty" tf:"shards,omitempty"`
}

type StoreParameters struct {

	// +kubebuilder:validation:Optional
	AppendMeta *bool `json:"appendMeta,omitempty" tf:"append_meta,omitempty"`

	// +kubebuilder:validation:Optional
	AutoSplit *bool `json:"autoSplit,omitempty" tf:"auto_split,omitempty"`

	// +kubebuilder:validation:Optional
	EnableWebTracking *bool `json:"enableWebTracking,omitempty" tf:"enable_web_tracking,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptConf []EncryptConfParameters `json:"encryptConf,omitempty" tf:"encrypt_conf,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSplitShardCount *float64 `json:"maxSplitShardCount,omitempty" tf:"max_split_shard_count,omitempty"`

	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	ShardCount *float64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`
}

type UserCmkInfoObservation struct {
}

type UserCmkInfoParameters struct {

	// +kubebuilder:validation:Required
	Arn *string `json:"arn" tf:"arn,omitempty"`

	// +kubebuilder:validation:Required
	CmkKeyID *string `json:"cmkKeyId" tf:"cmk_key_id,omitempty"`

	// +kubebuilder:validation:Required
	RegionID *string `json:"regionId" tf:"region_id,omitempty"`
}

// StoreSpec defines the desired state of Store
type StoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StoreParameters `json:"forProvider"`
}

// StoreStatus defines the observed state of Store.
type StoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Store is the Schema for the Stores API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type Store struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoreSpec   `json:"spec"`
	Status            StoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoreList contains a list of Stores
type StoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Store `json:"items"`
}

// Repository type metadata.
var (
	Store_Kind             = "Store"
	Store_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Store_Kind}.String()
	Store_KindAPIVersion   = Store_Kind + "." + CRDGroupVersion.String()
	Store_GroupVersionKind = CRDGroupVersion.WithKind(Store_Kind)
)

func init() {
	SchemeBuilder.Register(&Store{}, &StoreList{})
}
