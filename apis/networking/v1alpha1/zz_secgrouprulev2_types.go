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

type SecgroupRuleV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecgroupRuleV2Parameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// +kubebuilder:validation:Required
	Ethertype *string `json:"ethertype" tf:"ethertype,omitempty"`

	// +kubebuilder:validation:Optional
	PortRangeMax *float64 `json:"portRangeMax,omitempty" tf:"port_range_max,omitempty"`

	// +kubebuilder:validation:Optional
	PortRangeMin *float64 `json:"portRangeMin,omitempty" tf:"port_range_min,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteGroupID *string `json:"remoteGroupId,omitempty" tf:"remote_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteIPPrefix *string `json:"remoteIpPrefix,omitempty" tf:"remote_ip_prefix,omitempty"`

	// +kubebuilder:validation:Required
	SecurityGroupID *string `json:"securityGroupId" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// SecgroupRuleV2Spec defines the desired state of SecgroupRuleV2
type SecgroupRuleV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecgroupRuleV2Parameters `json:"forProvider"`
}

// SecgroupRuleV2Status defines the observed state of SecgroupRuleV2.
type SecgroupRuleV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecgroupRuleV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecgroupRuleV2 is the Schema for the SecgroupRuleV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type SecgroupRuleV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecgroupRuleV2Spec   `json:"spec"`
	Status            SecgroupRuleV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecgroupRuleV2List contains a list of SecgroupRuleV2s
type SecgroupRuleV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecgroupRuleV2 `json:"items"`
}

// Repository type metadata.
var (
	SecgroupRuleV2_Kind             = "SecgroupRuleV2"
	SecgroupRuleV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecgroupRuleV2_Kind}.String()
	SecgroupRuleV2_KindAPIVersion   = SecgroupRuleV2_Kind + "." + CRDGroupVersion.String()
	SecgroupRuleV2_GroupVersionKind = CRDGroupVersion.WithKind(SecgroupRuleV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SecgroupRuleV2{}, &SecgroupRuleV2List{})
}
