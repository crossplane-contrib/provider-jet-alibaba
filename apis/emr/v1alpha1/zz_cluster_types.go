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

type BootstrapActionObservation struct {
}

type BootstrapActionParameters struct {

	// +kubebuilder:validation:Optional
	Arg *string `json:"arg,omitempty" tf:"arg,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type ClusterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	BootstrapAction []BootstrapActionParameters `json:"bootstrapAction,omitempty" tf:"bootstrap_action,omitempty"`

	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// +kubebuilder:validation:Required
	ClusterType *string `json:"clusterType" tf:"cluster_type,omitempty"`

	// +kubebuilder:validation:Optional
	DepositType *string `json:"depositType,omitempty" tf:"deposit_type,omitempty"`

	// +kubebuilder:validation:Optional
	EasEnable *bool `json:"easEnable,omitempty" tf:"eas_enable,omitempty"`

	// +kubebuilder:validation:Required
	EmrVer *string `json:"emrVer" tf:"emr_ver,omitempty"`

	// +kubebuilder:validation:Optional
	HighAvailabilityEnable *bool `json:"highAvailabilityEnable,omitempty" tf:"high_availability_enable,omitempty"`

	// +kubebuilder:validation:Optional
	HostGroup []HostGroupParameters `json:"hostGroup,omitempty" tf:"host_group,omitempty"`

	// +kubebuilder:validation:Optional
	IsOpenPublicIP *bool `json:"isOpenPublicIp,omitempty" tf:"is_open_public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`

	// +kubebuilder:validation:Optional
	MasterPwd *string `json:"masterPwd,omitempty" tf:"master_pwd,omitempty"`

	// +kubebuilder:validation:Optional
	OptionSoftwareList []*string `json:"optionSoftwareList,omitempty" tf:"option_software_list,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	RelatedClusterID *string `json:"relatedClusterId,omitempty" tf:"related_cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	SSHEnable *bool `json:"sshEnable,omitempty" tf:"ssh_enable,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UseLocalMetadb *bool `json:"useLocalMetadb,omitempty" tf:"use_local_metadb,omitempty"`

	// +kubebuilder:validation:Optional
	UserDefinedEmrEcsRole *string `json:"userDefinedEmrEcsRole,omitempty" tf:"user_defined_emr_ecs_role,omitempty"`

	// +kubebuilder:validation:Optional
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`

	// +kubebuilder:validation:Required
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

type HostGroupObservation struct {
}

type HostGroupParameters struct {

	// +kubebuilder:validation:Optional
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// +kubebuilder:validation:Optional
	DiskCapacity *string `json:"diskCapacity,omitempty" tf:"disk_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	DiskCount *string `json:"diskCount,omitempty" tf:"disk_count,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	GpuDriver *string `json:"gpuDriver,omitempty" tf:"gpu_driver,omitempty"`

	// +kubebuilder:validation:Optional
	HostGroupName *string `json:"hostGroupName,omitempty" tf:"host_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	HostGroupType *string `json:"hostGroupType,omitempty" tf:"host_group_type,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceList *string `json:"instanceList,omitempty" tf:"instance_list,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	NodeCount *string `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	SysDiskCapacity *string `json:"sysDiskCapacity,omitempty" tf:"sys_disk_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	SysDiskType *string `json:"sysDiskType,omitempty" tf:"sys_disk_type,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}