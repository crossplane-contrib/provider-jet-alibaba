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

type ConstantParametersObservation struct {
}

type ConstantParametersParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	In *string `json:"in" tf:"in,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FcServiceConfigObservation struct {
}

type FcServiceConfigParameters struct {

	// +kubebuilder:validation:Optional
	ArnRole *string `json:"arnRole,omitempty" tf:"arn_role,omitempty"`

	// +kubebuilder:validation:Required
	FunctionName *string `json:"functionName" tf:"function_name,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Required
	Timeout *float64 `json:"timeout" tf:"timeout,omitempty"`
}

type GatewayAPIObservation struct {
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GatewayAPIParameters struct {

	// +kubebuilder:validation:Required
	AuthType *string `json:"authType" tf:"auth_type,omitempty"`

	// +kubebuilder:validation:Optional
	ConstantParameters []ConstantParametersParameters `json:"constantParameters,omitempty" tf:"constant_parameters,omitempty"`

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FcServiceConfig []FcServiceConfigParameters `json:"fcServiceConfig,omitempty" tf:"fc_service_config,omitempty"`

	// +kubebuilder:validation:Optional
	ForceNonceCheck *bool `json:"forceNonceCheck,omitempty" tf:"force_nonce_check,omitempty"`

	// +kubebuilder:validation:Required
	GroupID *string `json:"groupId" tf:"group_id,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPServiceConfig []HTTPServiceConfigParameters `json:"httpServiceConfig,omitempty" tf:"http_service_config,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPVPCServiceConfig []HTTPVPCServiceConfigParameters `json:"httpVpcServiceConfig,omitempty" tf:"http_vpc_service_config,omitempty"`

	// +kubebuilder:validation:Optional
	MockServiceConfig []MockServiceConfigParameters `json:"mockServiceConfig,omitempty" tf:"mock_service_config,omitempty"`

	// +kubebuilder:validation:Required
	RequestConfig []RequestConfigParameters `json:"requestConfig" tf:"request_config,omitempty"`

	// +kubebuilder:validation:Optional
	RequestParameters []RequestParametersParameters `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// +kubebuilder:validation:Required
	ServiceType *string `json:"serviceType" tf:"service_type,omitempty"`

	// +kubebuilder:validation:Optional
	StageNames []*string `json:"stageNames,omitempty" tf:"stage_names,omitempty"`

	// +kubebuilder:validation:Optional
	SystemParameters []SystemParametersParameters `json:"systemParameters,omitempty" tf:"system_parameters,omitempty"`
}

type HTTPServiceConfigObservation struct {
}

type HTTPServiceConfigParameters struct {

	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	AoneName *string `json:"aoneName,omitempty" tf:"aone_name,omitempty"`

	// +kubebuilder:validation:Required
	Method *string `json:"method" tf:"method,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	Timeout *float64 `json:"timeout" tf:"timeout,omitempty"`
}

type HTTPVPCServiceConfigObservation struct {
}

type HTTPVPCServiceConfigParameters struct {

	// +kubebuilder:validation:Optional
	AoneName *string `json:"aoneName,omitempty" tf:"aone_name,omitempty"`

	// +kubebuilder:validation:Required
	Method *string `json:"method" tf:"method,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	Timeout *float64 `json:"timeout" tf:"timeout,omitempty"`
}

type MockServiceConfigObservation struct {
}

type MockServiceConfigParameters struct {

	// +kubebuilder:validation:Optional
	AoneName *string `json:"aoneName,omitempty" tf:"aone_name,omitempty"`

	// +kubebuilder:validation:Required
	Result *string `json:"result" tf:"result,omitempty"`
}

type RequestConfigObservation struct {
}

type RequestConfigParameters struct {

	// +kubebuilder:validation:Optional
	BodyFormat *string `json:"bodyFormat,omitempty" tf:"body_format,omitempty"`

	// +kubebuilder:validation:Required
	Method *string `json:"method" tf:"method,omitempty"`

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type RequestParametersObservation struct {
}

type RequestParametersParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	In *string `json:"in" tf:"in,omitempty"`

	// +kubebuilder:validation:Required
	InService *string `json:"inService" tf:"in_service,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NameService *string `json:"nameService" tf:"name_service,omitempty"`

	// +kubebuilder:validation:Required
	Required *string `json:"required" tf:"required,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SystemParametersObservation struct {
}

type SystemParametersParameters struct {

	// +kubebuilder:validation:Required
	In *string `json:"in" tf:"in,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NameService *string `json:"nameService" tf:"name_service,omitempty"`
}

// GatewayAPISpec defines the desired state of GatewayAPI
type GatewayAPISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayAPIParameters `json:"forProvider"`
}

// GatewayAPIStatus defines the observed state of GatewayAPI.
type GatewayAPIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayAPIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAPI is the Schema for the GatewayAPIs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type GatewayAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayAPISpec   `json:"spec"`
	Status            GatewayAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAPIList contains a list of GatewayAPIs
type GatewayAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayAPI `json:"items"`
}

// Repository type metadata.
var (
	GatewayAPI_Kind             = "GatewayAPI"
	GatewayAPI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayAPI_Kind}.String()
	GatewayAPI_KindAPIVersion   = GatewayAPI_Kind + "." + CRDGroupVersion.String()
	GatewayAPI_GroupVersionKind = CRDGroupVersion.WithKind(GatewayAPI_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayAPI{}, &GatewayAPIList{})
}