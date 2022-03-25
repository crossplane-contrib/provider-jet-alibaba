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

type ReadonlyInstanceObservation struct {
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`
}

type ReadonlyInstanceParameters struct {

	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// +kubebuilder:validation:Optional
	CAType *string `json:"caType,omitempty" tf:"ca_type,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCACert *string `json:"clientCaCert,omitempty" tf:"client_ca_cert,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCAEnabled *float64 `json:"clientCaEnabled,omitempty" tf:"client_ca_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCertRevocationList *string `json:"clientCertRevocationList,omitempty" tf:"client_cert_revocation_list,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCrlEnabled *float64 `json:"clientCrlEnabled,omitempty" tf:"client_crl_enabled,omitempty"`

	// +kubebuilder:validation:Required
	EngineVersion *string `json:"engineVersion" tf:"engine_version,omitempty"`

	// +kubebuilder:validation:Optional
	ForceRestart *bool `json:"forceRestart,omitempty" tf:"force_restart,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// +kubebuilder:validation:Required
	InstanceStorage *float64 `json:"instanceStorage" tf:"instance_storage,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	MasterDBInstanceID *string `json:"masterDbInstanceId" tf:"master_db_instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters []ReadonlyInstanceParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicationACL *string `json:"replicationAcl,omitempty" tf:"replication_acl,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SSLEnabled *float64 `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ServerCert *string `json:"serverCert,omitempty" tf:"server_cert,omitempty"`

	// +kubebuilder:validation:Optional
	ServerKey *string `json:"serverKey,omitempty" tf:"server_key,omitempty"`

	// +kubebuilder:validation:Optional
	SwitchTime *string `json:"switchTime,omitempty" tf:"switch_time,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TargetMinorVersion *string `json:"targetMinorVersion,omitempty" tf:"target_minor_version,omitempty"`

	// +kubebuilder:validation:Optional
	UpgradeDBInstanceKernelVersion *bool `json:"upgradeDbInstanceKernelVersion,omitempty" tf:"upgrade_db_instance_kernel_version,omitempty"`

	// +kubebuilder:validation:Optional
	UpgradeTime *string `json:"upgradeTime,omitempty" tf:"upgrade_time,omitempty"`

	// +kubebuilder:validation:Optional
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ReadonlyInstanceParametersObservation struct {
}

type ReadonlyInstanceParametersParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// ReadonlyInstanceSpec defines the desired state of ReadonlyInstance
type ReadonlyInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReadonlyInstanceParameters `json:"forProvider"`
}

// ReadonlyInstanceStatus defines the observed state of ReadonlyInstance.
type ReadonlyInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReadonlyInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReadonlyInstance is the Schema for the ReadonlyInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ReadonlyInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReadonlyInstanceSpec   `json:"spec"`
	Status            ReadonlyInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReadonlyInstanceList contains a list of ReadonlyInstances
type ReadonlyInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReadonlyInstance `json:"items"`
}

// Repository type metadata.
var (
	ReadonlyInstance_Kind             = "ReadonlyInstance"
	ReadonlyInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReadonlyInstance_Kind}.String()
	ReadonlyInstance_KindAPIVersion   = ReadonlyInstance_Kind + "." + CRDGroupVersion.String()
	ReadonlyInstance_GroupVersionKind = CRDGroupVersion.WithKind(ReadonlyInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ReadonlyInstance{}, &ReadonlyInstanceList{})
}