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

type VolumeAttachV3Observation struct {
	DriverVolumeType *string `json:"driverVolumeType,omitempty" tf:"driver_volume_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MountPointBase *string `json:"mountPointBase,omitempty" tf:"mount_point_base,omitempty"`
}

type VolumeAttachV3Parameters struct {

	// +kubebuilder:validation:Optional
	AttachMode *string `json:"attachMode,omitempty" tf:"attach_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	Initiator *string `json:"initiator,omitempty" tf:"initiator,omitempty"`

	// +kubebuilder:validation:Optional
	Multipath *bool `json:"multipath,omitempty" tf:"multipath,omitempty"`

	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// +kubebuilder:validation:Optional
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	VolumeID *string `json:"volumeId" tf:"volume_id,omitempty"`

	// +kubebuilder:validation:Optional
	Wwnn *string `json:"wwnn,omitempty" tf:"wwnn,omitempty"`

	// +kubebuilder:validation:Optional
	Wwpn []*string `json:"wwpn,omitempty" tf:"wwpn,omitempty"`
}

// VolumeAttachV3Spec defines the desired state of VolumeAttachV3
type VolumeAttachV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeAttachV3Parameters `json:"forProvider"`
}

// VolumeAttachV3Status defines the observed state of VolumeAttachV3.
type VolumeAttachV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeAttachV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachV3 is the Schema for the VolumeAttachV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type VolumeAttachV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeAttachV3Spec   `json:"spec"`
	Status            VolumeAttachV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachV3List contains a list of VolumeAttachV3s
type VolumeAttachV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeAttachV3 `json:"items"`
}

// Repository type metadata.
var (
	VolumeAttachV3_Kind             = "VolumeAttachV3"
	VolumeAttachV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeAttachV3_Kind}.String()
	VolumeAttachV3_KindAPIVersion   = VolumeAttachV3_Kind + "." + CRDGroupVersion.String()
	VolumeAttachV3_GroupVersionKind = CRDGroupVersion.WithKind(VolumeAttachV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeAttachV3{}, &VolumeAttachV3List{})
}