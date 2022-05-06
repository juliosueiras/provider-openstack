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

type AddressscopeV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AddressscopeV2Parameters struct {

	// +kubebuilder:validation:Optional
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`
}

// AddressscopeV2Spec defines the desired state of AddressscopeV2
type AddressscopeV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddressscopeV2Parameters `json:"forProvider"`
}

// AddressscopeV2Status defines the observed state of AddressscopeV2.
type AddressscopeV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddressscopeV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AddressscopeV2 is the Schema for the AddressscopeV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type AddressscopeV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddressscopeV2Spec   `json:"spec"`
	Status            AddressscopeV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddressscopeV2List contains a list of AddressscopeV2s
type AddressscopeV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddressscopeV2 `json:"items"`
}

// Repository type metadata.
var (
	AddressscopeV2_Kind             = "AddressscopeV2"
	AddressscopeV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddressscopeV2_Kind}.String()
	AddressscopeV2_KindAPIVersion   = AddressscopeV2_Kind + "." + CRDGroupVersion.String()
	AddressscopeV2_GroupVersionKind = CRDGroupVersion.WithKind(AddressscopeV2_Kind)
)

func init() {
	SchemeBuilder.Register(&AddressscopeV2{}, &AddressscopeV2List{})
}