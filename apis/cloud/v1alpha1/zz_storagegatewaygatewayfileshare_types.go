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

type StorageGatewayGatewayFileShareObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`
}

type StorageGatewayGatewayFileShareParameters struct {

	// +kubebuilder:validation:Optional
	AccessBasedEnumeration *bool `json:"accessBasedEnumeration,omitempty" tf:"access_based_enumeration,omitempty"`

	// +kubebuilder:validation:Optional
	BackendLimit *float64 `json:"backendLimit,omitempty" tf:"backend_limit,omitempty"`

	// +kubebuilder:validation:Optional
	Browsable *bool `json:"browsable,omitempty" tf:"browsable,omitempty"`

	// +kubebuilder:validation:Optional
	BypassCacheRead *bool `json:"bypassCacheRead,omitempty" tf:"bypass_cache_read,omitempty"`

	// +kubebuilder:validation:Optional
	CacheMode *string `json:"cacheMode,omitempty" tf:"cache_mode,omitempty"`

	// +kubebuilder:validation:Optional
	DirectIo *bool `json:"directIo,omitempty" tf:"direct_io,omitempty"`

	// +kubebuilder:validation:Optional
	DownloadLimit *float64 `json:"downloadLimit,omitempty" tf:"download_limit,omitempty"`

	// +kubebuilder:validation:Optional
	FastReclaim *bool `json:"fastReclaim,omitempty" tf:"fast_reclaim,omitempty"`

	// +kubebuilder:validation:Optional
	FeLimit *float64 `json:"feLimit,omitempty" tf:"fe_limit,omitempty"`

	// +kubebuilder:validation:Required
	GatewayFileShareName *string `json:"gatewayFileShareName" tf:"gateway_file_share_name,omitempty"`

	// +kubebuilder:validation:Required
	GatewayID *string `json:"gatewayId" tf:"gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreDelete *bool `json:"ignoreDelete,omitempty" tf:"ignore_delete,omitempty"`

	// +kubebuilder:validation:Optional
	InPlace *bool `json:"inPlace,omitempty" tf:"in_place,omitempty"`

	// +kubebuilder:validation:Optional
	LagPeriod *float64 `json:"lagPeriod,omitempty" tf:"lag_period,omitempty"`

	// +kubebuilder:validation:Required
	LocalPath *string `json:"localPath" tf:"local_path,omitempty"`

	// +kubebuilder:validation:Optional
	NFSV4Optimization *bool `json:"nfsV4Optimization,omitempty" tf:"nfs_v4_optimization,omitempty"`

	// +kubebuilder:validation:Required
	OssBucketName *string `json:"ossBucketName" tf:"oss_bucket_name,omitempty"`

	// +kubebuilder:validation:Optional
	OssBucketSSL *bool `json:"ossBucketSsl,omitempty" tf:"oss_bucket_ssl,omitempty"`

	// +kubebuilder:validation:Required
	OssEndpoint *string `json:"ossEndpoint" tf:"oss_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	PartialSyncPaths *string `json:"partialSyncPaths,omitempty" tf:"partial_sync_paths,omitempty"`

	// +kubebuilder:validation:Optional
	PathPrefix *string `json:"pathPrefix,omitempty" tf:"path_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	PollingInterval *float64 `json:"pollingInterval,omitempty" tf:"polling_interval,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteSync *bool `json:"remoteSync,omitempty" tf:"remote_sync,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteSyncDownload *bool `json:"remoteSyncDownload,omitempty" tf:"remote_sync_download,omitempty"`

	// +kubebuilder:validation:Optional
	RoClientList *string `json:"roClientList,omitempty" tf:"ro_client_list,omitempty"`

	// +kubebuilder:validation:Optional
	RoUserList *string `json:"roUserList,omitempty" tf:"ro_user_list,omitempty"`

	// +kubebuilder:validation:Optional
	RwClientList *string `json:"rwClientList,omitempty" tf:"rw_client_list,omitempty"`

	// +kubebuilder:validation:Optional
	RwUserList *string `json:"rwUserList,omitempty" tf:"rw_user_list,omitempty"`

	// +kubebuilder:validation:Optional
	Squash *string `json:"squash,omitempty" tf:"squash,omitempty"`

	// +kubebuilder:validation:Optional
	SupportArchive *bool `json:"supportArchive,omitempty" tf:"support_archive,omitempty"`

	// +kubebuilder:validation:Optional
	TransferAcceleration *bool `json:"transferAcceleration,omitempty" tf:"transfer_acceleration,omitempty"`

	// +kubebuilder:validation:Optional
	WindowsACL *bool `json:"windowsAcl,omitempty" tf:"windows_acl,omitempty"`
}

// StorageGatewayGatewayFileShareSpec defines the desired state of StorageGatewayGatewayFileShare
type StorageGatewayGatewayFileShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageGatewayGatewayFileShareParameters `json:"forProvider"`
}

// StorageGatewayGatewayFileShareStatus defines the observed state of StorageGatewayGatewayFileShare.
type StorageGatewayGatewayFileShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageGatewayGatewayFileShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageGatewayGatewayFileShare is the Schema for the StorageGatewayGatewayFileShares API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type StorageGatewayGatewayFileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageGatewayGatewayFileShareSpec   `json:"spec"`
	Status            StorageGatewayGatewayFileShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageGatewayGatewayFileShareList contains a list of StorageGatewayGatewayFileShares
type StorageGatewayGatewayFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageGatewayGatewayFileShare `json:"items"`
}

// Repository type metadata.
var (
	StorageGatewayGatewayFileShare_Kind             = "StorageGatewayGatewayFileShare"
	StorageGatewayGatewayFileShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageGatewayGatewayFileShare_Kind}.String()
	StorageGatewayGatewayFileShare_KindAPIVersion   = StorageGatewayGatewayFileShare_Kind + "." + CRDGroupVersion.String()
	StorageGatewayGatewayFileShare_GroupVersionKind = CRDGroupVersion.WithKind(StorageGatewayGatewayFileShare_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageGatewayGatewayFileShare{}, &StorageGatewayGatewayFileShareList{})
}