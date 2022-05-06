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

type ServiceV2Observation struct {
	ExternalV4IP *string `json:"externalV4Ip,omitempty" tf:"external_v4_ip,omitempty"`

	ExternalV6IP *string `json:"externalV6Ip,omitempty" tf:"external_v6_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ServiceV2Parameters struct {

	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	RouterID *string `json:"routerId" tf:"router_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// ServiceV2Spec defines the desired state of ServiceV2
type ServiceV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceV2Parameters `json:"forProvider"`
}

// ServiceV2Status defines the observed state of ServiceV2.
type ServiceV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceV2 is the Schema for the ServiceV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type ServiceV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceV2Spec   `json:"spec"`
	Status            ServiceV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceV2List contains a list of ServiceV2s
type ServiceV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceV2 `json:"items"`
}

// Repository type metadata.
var (
	ServiceV2_Kind             = "ServiceV2"
	ServiceV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceV2_Kind}.String()
	ServiceV2_KindAPIVersion   = ServiceV2_Kind + "." + CRDGroupVersion.String()
	ServiceV2_GroupVersionKind = CRDGroupVersion.WithKind(ServiceV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceV2{}, &ServiceV2List{})
}
