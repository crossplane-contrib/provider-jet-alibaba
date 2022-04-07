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

type BackupPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackupPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ArchiveBackupKeepCount *float64 `json:"archiveBackupKeepCount,omitempty" tf:"archive_backup_keep_count,omitempty"`

	// +kubebuilder:validation:Optional
	ArchiveBackupKeepPolicy *string `json:"archiveBackupKeepPolicy,omitempty" tf:"archive_backup_keep_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ArchiveBackupRetentionPeriod *float64 `json:"archiveBackupRetentionPeriod,omitempty" tf:"archive_backup_retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	BackupPeriod []*string `json:"backupPeriod,omitempty" tf:"backup_period,omitempty"`

	// +kubebuilder:validation:Optional
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time,omitempty"`

	// +kubebuilder:validation:Optional
	CompressType *string `json:"compressType,omitempty" tf:"compress_type,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBackupLog *bool `json:"enableBackupLog,omitempty" tf:"enable_backup_log,omitempty"`

	// +kubebuilder:validation:Optional
	HighSpaceUsageProtection *string `json:"highSpaceUsageProtection,omitempty" tf:"high_space_usage_protection,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	LocalLogRetentionHours *float64 `json:"localLogRetentionHours,omitempty" tf:"local_log_retention_hours,omitempty"`

	// +kubebuilder:validation:Optional
	LocalLogRetentionSpace *float64 `json:"localLogRetentionSpace,omitempty" tf:"local_log_retention_space,omitempty"`

	// +kubebuilder:validation:Optional
	LogBackup *bool `json:"logBackup,omitempty" tf:"log_backup,omitempty"`

	// +kubebuilder:validation:Optional
	LogBackupFrequency *string `json:"logBackupFrequency,omitempty" tf:"log_backup_frequency,omitempty"`

	// +kubebuilder:validation:Optional
	LogBackupRetentionPeriod *float64 `json:"logBackupRetentionPeriod,omitempty" tf:"log_backup_retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	LogRetentionPeriod *float64 `json:"logRetentionPeriod,omitempty" tf:"log_retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredBackupPeriod []*string `json:"preferredBackupPeriod,omitempty" tf:"preferred_backup_period,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredBackupTime *string `json:"preferredBackupTime,omitempty" tf:"preferred_backup_time,omitempty"`

	// +kubebuilder:validation:Optional
	ReleasedKeepPolicy *string `json:"releasedKeepPolicy,omitempty" tf:"released_keep_policy,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

// BackupPolicySpec defines the desired state of BackupPolicy
type BackupPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyParameters `json:"forProvider"`
}

// BackupPolicyStatus defines the observed state of BackupPolicy.
type BackupPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicy is the Schema for the BackupPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type BackupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPolicySpec   `json:"spec"`
	Status            BackupPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyList contains a list of BackupPolicys
type BackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicy `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicy_Kind             = "BackupPolicy"
	BackupPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicy_Kind}.String()
	BackupPolicy_KindAPIVersion   = BackupPolicy_Kind + "." + CRDGroupVersion.String()
	BackupPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicy{}, &BackupPolicyList{})
}
