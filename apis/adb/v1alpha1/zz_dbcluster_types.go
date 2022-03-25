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

type DBClusterObservation struct {
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DBClusterParameters struct {

	// +kubebuilder:validation:Optional
	AutoRenewPeriod *float64 `json:"autoRenewPeriod,omitempty" tf:"auto_renew_period,omitempty"`

	// +kubebuilder:validation:Optional
	ComputeResource *string `json:"computeResource,omitempty" tf:"compute_resource,omitempty"`

	// +kubebuilder:validation:Required
	DBClusterCategory *string `json:"dbClusterCategory" tf:"db_cluster_category,omitempty"`

	// +kubebuilder:validation:Optional
	DBClusterClass *string `json:"dbClusterClass,omitempty" tf:"db_cluster_class,omitempty"`

	// +kubebuilder:validation:Optional
	DBClusterVersion *string `json:"dbClusterVersion,omitempty" tf:"db_cluster_version,omitempty"`

	// +kubebuilder:validation:Optional
	DBNodeClass *string `json:"dbNodeClass,omitempty" tf:"db_node_class,omitempty"`

	// +kubebuilder:validation:Optional
	DBNodeCount *float64 `json:"dbNodeCount,omitempty" tf:"db_node_count,omitempty"`

	// +kubebuilder:validation:Optional
	DBNodeStorage *float64 `json:"dbNodeStorage,omitempty" tf:"db_node_storage,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticIoResource *float64 `json:"elasticIoResource,omitempty" tf:"elastic_io_resource,omitempty"`

	// +kubebuilder:validation:Optional
	MaintainTime *string `json:"maintainTime,omitempty" tf:"maintain_time,omitempty"`

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	ModifyType *string `json:"modifyType,omitempty" tf:"modify_type,omitempty"`

	// +kubebuilder:validation:Optional
	PayType *string `json:"payType,omitempty" tf:"pay_type,omitempty"`

	// +kubebuilder:validation:Optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	RenewalStatus *string `json:"renewalStatus,omitempty" tf:"renewal_status,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityIps []*string `json:"securityIps,omitempty" tf:"security_ips,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

// DBClusterSpec defines the desired state of DBCluster
type DBClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DBClusterParameters `json:"forProvider"`
}

// DBClusterStatus defines the observed state of DBCluster.
type DBClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DBClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DBCluster is the Schema for the DBClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type DBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBClusterSpec   `json:"spec"`
	Status            DBClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DBClusterList contains a list of DBClusters
type DBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBCluster `json:"items"`
}

// Repository type metadata.
var (
	DBCluster_Kind             = "DBCluster"
	DBCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DBCluster_Kind}.String()
	DBCluster_KindAPIVersion   = DBCluster_Kind + "." + CRDGroupVersion.String()
	DBCluster_GroupVersionKind = CRDGroupVersion.WithKind(DBCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&DBCluster{}, &DBClusterList{})
}