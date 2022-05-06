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

type QosBandwidthLimitRuleV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QosBandwidthLimitRuleV2Parameters struct {

	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// +kubebuilder:validation:Optional
	MaxBurstKbps *float64 `json:"maxBurstKbps,omitempty" tf:"max_burst_kbps,omitempty"`

	// +kubebuilder:validation:Required
	MaxKbps *float64 `json:"maxKbps" tf:"max_kbps,omitempty"`

	// +kubebuilder:validation:Required
	QosPolicyID *string `json:"qosPolicyId" tf:"qos_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// QosBandwidthLimitRuleV2Spec defines the desired state of QosBandwidthLimitRuleV2
type QosBandwidthLimitRuleV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QosBandwidthLimitRuleV2Parameters `json:"forProvider"`
}

// QosBandwidthLimitRuleV2Status defines the observed state of QosBandwidthLimitRuleV2.
type QosBandwidthLimitRuleV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QosBandwidthLimitRuleV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QosBandwidthLimitRuleV2 is the Schema for the QosBandwidthLimitRuleV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type QosBandwidthLimitRuleV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QosBandwidthLimitRuleV2Spec   `json:"spec"`
	Status            QosBandwidthLimitRuleV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QosBandwidthLimitRuleV2List contains a list of QosBandwidthLimitRuleV2s
type QosBandwidthLimitRuleV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QosBandwidthLimitRuleV2 `json:"items"`
}

// Repository type metadata.
var (
	QosBandwidthLimitRuleV2_Kind             = "QosBandwidthLimitRuleV2"
	QosBandwidthLimitRuleV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QosBandwidthLimitRuleV2_Kind}.String()
	QosBandwidthLimitRuleV2_KindAPIVersion   = QosBandwidthLimitRuleV2_Kind + "." + CRDGroupVersion.String()
	QosBandwidthLimitRuleV2_GroupVersionKind = CRDGroupVersion.WithKind(QosBandwidthLimitRuleV2_Kind)
)

func init() {
	SchemeBuilder.Register(&QosBandwidthLimitRuleV2{}, &QosBandwidthLimitRuleV2List{})
}
