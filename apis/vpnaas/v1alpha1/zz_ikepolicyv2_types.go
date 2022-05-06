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

type IkePolicyV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IkePolicyV2Parameters struct {

	// +kubebuilder:validation:Optional
	AuthAlgorithm *string `json:"authAlgorithm,omitempty" tf:"auth_algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	IkeVersion *string `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`

	// +kubebuilder:validation:Optional
	Lifetime []LifetimeParameters `json:"lifetime,omitempty" tf:"lifetime,omitempty"`

	// +kubebuilder:validation:Optional
	Pfs *string `json:"pfs,omitempty" tf:"pfs,omitempty"`

	// +kubebuilder:validation:Optional
	Phase1NegotiationMode *string `json:"phase1NegotiationMode,omitempty" tf:"phase1_negotiation_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type LifetimeObservation struct {
}

type LifetimeParameters struct {

	// +kubebuilder:validation:Optional
	Units *string `json:"units,omitempty" tf:"units,omitempty"`

	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

// IkePolicyV2Spec defines the desired state of IkePolicyV2
type IkePolicyV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IkePolicyV2Parameters `json:"forProvider"`
}

// IkePolicyV2Status defines the observed state of IkePolicyV2.
type IkePolicyV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IkePolicyV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IkePolicyV2 is the Schema for the IkePolicyV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type IkePolicyV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IkePolicyV2Spec   `json:"spec"`
	Status            IkePolicyV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IkePolicyV2List contains a list of IkePolicyV2s
type IkePolicyV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IkePolicyV2 `json:"items"`
}

// Repository type metadata.
var (
	IkePolicyV2_Kind             = "IkePolicyV2"
	IkePolicyV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IkePolicyV2_Kind}.String()
	IkePolicyV2_KindAPIVersion   = IkePolicyV2_Kind + "." + CRDGroupVersion.String()
	IkePolicyV2_GroupVersionKind = CRDGroupVersion.WithKind(IkePolicyV2_Kind)
)

func init() {
	SchemeBuilder.Register(&IkePolicyV2{}, &IkePolicyV2List{})
}
