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

type FloatingipAssociateV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FloatingipAssociateV2Parameters struct {

	// +kubebuilder:validation:Optional
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// +kubebuilder:validation:Required
	FloatingIP *string `json:"floatingIp" tf:"floating_ip,omitempty"`

	// +kubebuilder:validation:Required
	PortID *string `json:"portId" tf:"port_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// FloatingipAssociateV2Spec defines the desired state of FloatingipAssociateV2
type FloatingipAssociateV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FloatingipAssociateV2Parameters `json:"forProvider"`
}

// FloatingipAssociateV2Status defines the observed state of FloatingipAssociateV2.
type FloatingipAssociateV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FloatingipAssociateV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FloatingipAssociateV2 is the Schema for the FloatingipAssociateV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type FloatingipAssociateV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FloatingipAssociateV2Spec   `json:"spec"`
	Status            FloatingipAssociateV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FloatingipAssociateV2List contains a list of FloatingipAssociateV2s
type FloatingipAssociateV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FloatingipAssociateV2 `json:"items"`
}

// Repository type metadata.
var (
	FloatingipAssociateV2_Kind             = "FloatingipAssociateV2"
	FloatingipAssociateV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FloatingipAssociateV2_Kind}.String()
	FloatingipAssociateV2_KindAPIVersion   = FloatingipAssociateV2_Kind + "." + CRDGroupVersion.String()
	FloatingipAssociateV2_GroupVersionKind = CRDGroupVersion.WithKind(FloatingipAssociateV2_Kind)
)

func init() {
	SchemeBuilder.Register(&FloatingipAssociateV2{}, &FloatingipAssociateV2List{})
}
