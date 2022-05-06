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

type ImageV2Observation struct {
	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	File *string `json:"file,omitempty" tf:"file,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	SizeBytes *float64 `json:"sizeBytes,omitempty" tf:"size_bytes,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	UpdateAt *string `json:"updateAt,omitempty" tf:"update_at,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ImageV2Parameters struct {

	// +kubebuilder:validation:Required
	ContainerFormat *string `json:"containerFormat" tf:"container_format,omitempty"`

	// +kubebuilder:validation:Required
	DiskFormat *string `json:"diskFormat" tf:"disk_format,omitempty"`

	// +kubebuilder:validation:Optional
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// +kubebuilder:validation:Optional
	ImageCachePath *string `json:"imageCachePath,omitempty" tf:"image_cache_path,omitempty"`

	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	ImageSourcePasswordSecretRef *v1.SecretKeySelector `json:"imageSourcePasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ImageSourceURL *string `json:"imageSourceUrl,omitempty" tf:"image_source_url,omitempty"`

	// +kubebuilder:validation:Optional
	ImageSourceUsername *string `json:"imageSourceUsername,omitempty" tf:"image_source_username,omitempty"`

	// +kubebuilder:validation:Optional
	LocalFilePath *string `json:"localFilePath,omitempty" tf:"local_file_path,omitempty"`

	// +kubebuilder:validation:Optional
	MinDiskGb *float64 `json:"minDiskGb,omitempty" tf:"min_disk_gb,omitempty"`

	// +kubebuilder:validation:Optional
	MinRAMMb *float64 `json:"minRamMb,omitempty" tf:"min_ram_mb,omitempty"`

	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// +kubebuilder:validation:Optional
	Protected *bool `json:"protected,omitempty" tf:"protected,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VerifyChecksum *bool `json:"verifyChecksum,omitempty" tf:"verify_checksum,omitempty"`

	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// +kubebuilder:validation:Optional
	WebDownload *bool `json:"webDownload,omitempty" tf:"web_download,omitempty"`
}

// ImageV2Spec defines the desired state of ImageV2
type ImageV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageV2Parameters `json:"forProvider"`
}

// ImageV2Status defines the observed state of ImageV2.
type ImageV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImageV2 is the Schema for the ImageV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type ImageV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageV2Spec   `json:"spec"`
	Status            ImageV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageV2List contains a list of ImageV2s
type ImageV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageV2 `json:"items"`
}

// Repository type metadata.
var (
	ImageV2_Kind             = "ImageV2"
	ImageV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageV2_Kind}.String()
	ImageV2_KindAPIVersion   = ImageV2_Kind + "." + CRDGroupVersion.String()
	ImageV2_GroupVersionKind = CRDGroupVersion.WithKind(ImageV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageV2{}, &ImageV2List{})
}
