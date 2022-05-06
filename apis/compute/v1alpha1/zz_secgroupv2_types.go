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

type RuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleParameters struct {

	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// +kubebuilder:validation:Optional
	FromGroupID *string `json:"fromGroupId,omitempty" tf:"from_group_id,omitempty"`

	// +kubebuilder:validation:Required
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self,omitempty"`

	// +kubebuilder:validation:Required
	ToPort *float64 `json:"toPort" tf:"to_port,omitempty"`
}

type SecgroupV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecgroupV2Parameters struct {

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

// SecgroupV2Spec defines the desired state of SecgroupV2
type SecgroupV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecgroupV2Parameters `json:"forProvider"`
}

// SecgroupV2Status defines the observed state of SecgroupV2.
type SecgroupV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecgroupV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecgroupV2 is the Schema for the SecgroupV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type SecgroupV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecgroupV2Spec   `json:"spec"`
	Status            SecgroupV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecgroupV2List contains a list of SecgroupV2s
type SecgroupV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecgroupV2 `json:"items"`
}

// Repository type metadata.
var (
	SecgroupV2_Kind             = "SecgroupV2"
	SecgroupV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecgroupV2_Kind}.String()
	SecgroupV2_KindAPIVersion   = SecgroupV2_Kind + "." + CRDGroupVersion.String()
	SecgroupV2_GroupVersionKind = CRDGroupVersion.WithKind(SecgroupV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SecgroupV2{}, &SecgroupV2List{})
}
