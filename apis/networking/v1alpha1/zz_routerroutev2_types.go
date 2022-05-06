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

type RouterRouteV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouterRouteV2Parameters struct {

	// +kubebuilder:validation:Required
	DestinationCidr *string `json:"destinationCidr" tf:"destination_cidr,omitempty"`

	// +kubebuilder:validation:Required
	NextHop *string `json:"nextHop" tf:"next_hop,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	RouterID *string `json:"routerId" tf:"router_id,omitempty"`
}

// RouterRouteV2Spec defines the desired state of RouterRouteV2
type RouterRouteV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterRouteV2Parameters `json:"forProvider"`
}

// RouterRouteV2Status defines the observed state of RouterRouteV2.
type RouterRouteV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterRouteV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouterRouteV2 is the Schema for the RouterRouteV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type RouterRouteV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterRouteV2Spec   `json:"spec"`
	Status            RouterRouteV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterRouteV2List contains a list of RouterRouteV2s
type RouterRouteV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterRouteV2 `json:"items"`
}

// Repository type metadata.
var (
	RouterRouteV2_Kind             = "RouterRouteV2"
	RouterRouteV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterRouteV2_Kind}.String()
	RouterRouteV2_KindAPIVersion   = RouterRouteV2_Kind + "." + CRDGroupVersion.String()
	RouterRouteV2_GroupVersionKind = CRDGroupVersion.WithKind(RouterRouteV2_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterRouteV2{}, &RouterRouteV2List{})
}
