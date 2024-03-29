/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ComputeImageCapabilitySchemaObservation struct {

	// The ocid of the compute global image capability schema
	ComputeGlobalImageCapabilitySchemaID *string `json:"computeGlobalImageCapabilitySchemaId,omitempty" tf:"compute_global_image_capability_schema_id,omitempty"`

	// The compute image capability schema OCID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time the compute image capability schema was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type ComputeImageCapabilitySchemaParameters struct {

	// (Updatable) The OCID of the compartment that contains the resource.
	// +kubebuilder:validation:Required
	CompartmentID *string `json:"compartmentId" tf:"compartment_id,omitempty"`

	// The name of the compute global image capability schema version
	// +kubebuilder:validation:Required
	ComputeGlobalImageCapabilitySchemaVersionName *string `json:"computeGlobalImageCapabilitySchemaVersionName" tf:"compute_global_image_capability_schema_version_name,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The ocid of the image
	// +kubebuilder:validation:Required
	ImageID *string `json:"imageId" tf:"image_id,omitempty"`

	// (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
	// +kubebuilder:validation:Required
	SchemaData map[string]*string `json:"schemaData" tf:"schema_data,omitempty"`
}

// ComputeImageCapabilitySchemaSpec defines the desired state of ComputeImageCapabilitySchema
type ComputeImageCapabilitySchemaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComputeImageCapabilitySchemaParameters `json:"forProvider"`
}

// ComputeImageCapabilitySchemaStatus defines the observed state of ComputeImageCapabilitySchema.
type ComputeImageCapabilitySchemaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComputeImageCapabilitySchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeImageCapabilitySchema is the Schema for the ComputeImageCapabilitySchemas API. Provides the Compute Image Capability Schema resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type ComputeImageCapabilitySchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeImageCapabilitySchemaSpec   `json:"spec"`
	Status            ComputeImageCapabilitySchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeImageCapabilitySchemaList contains a list of ComputeImageCapabilitySchemas
type ComputeImageCapabilitySchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeImageCapabilitySchema `json:"items"`
}

// Repository type metadata.
var (
	ComputeImageCapabilitySchema_Kind             = "ComputeImageCapabilitySchema"
	ComputeImageCapabilitySchema_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ComputeImageCapabilitySchema_Kind}.String()
	ComputeImageCapabilitySchema_KindAPIVersion   = ComputeImageCapabilitySchema_Kind + "." + CRDGroupVersion.String()
	ComputeImageCapabilitySchema_GroupVersionKind = CRDGroupVersion.WithKind(ComputeImageCapabilitySchema_Kind)
)

func init() {
	SchemeBuilder.Register(&ComputeImageCapabilitySchema{}, &ComputeImageCapabilitySchemaList{})
}
