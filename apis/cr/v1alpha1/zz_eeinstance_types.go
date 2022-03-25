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

type EeInstanceObservation struct {
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type EeInstanceParameters struct {

	// +kubebuilder:validation:Optional
	CustomOssBucket *string `json:"customOssBucket,omitempty" tf:"custom_oss_bucket,omitempty"`

	// +kubebuilder:validation:Required
	InstanceName *string `json:"instanceName" tf:"instance_name,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	KMSEncryptedPassword *string `json:"kmsEncryptedPassword,omitempty" tf:"kms_encrypted_password,omitempty"`

	// +kubebuilder:validation:Optional
	KMSEncryptionContext map[string]*string `json:"kmsEncryptionContext,omitempty" tf:"kms_encryption_context,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	RenewPeriod *float64 `json:"renewPeriod,omitempty" tf:"renew_period,omitempty"`

	// +kubebuilder:validation:Optional
	RenewalStatus *string `json:"renewalStatus,omitempty" tf:"renewal_status,omitempty"`
}

// EeInstanceSpec defines the desired state of EeInstance
type EeInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EeInstanceParameters `json:"forProvider"`
}

// EeInstanceStatus defines the observed state of EeInstance.
type EeInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EeInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EeInstance is the Schema for the EeInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type EeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EeInstanceSpec   `json:"spec"`
	Status            EeInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EeInstanceList contains a list of EeInstances
type EeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EeInstance `json:"items"`
}

// Repository type metadata.
var (
	EeInstance_Kind             = "EeInstance"
	EeInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EeInstance_Kind}.String()
	EeInstance_KindAPIVersion   = EeInstance_Kind + "." + CRDGroupVersion.String()
	EeInstance_GroupVersionKind = CRDGroupVersion.WithKind(EeInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&EeInstance{}, &EeInstanceList{})
}
