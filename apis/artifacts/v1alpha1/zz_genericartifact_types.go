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

type GenericArtifactObservation struct {

	// A user-defined path to describe the location of an artifact. Slashes do not create a directory structure, but you can use slashes to organize the repository. An artifact path does not include an artifact version.  Example: project01/my-web-app/artifact-abc
	ArtifactPath *string `json:"artifactPath,omitempty" tf:"artifact_path,omitempty"`

	// The OCID of the repository's compartment.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// The artifact name with the format of <artifact-path>:<artifact-version>. The artifact name is truncated to a maximum length of 255.  Example: project01/my-web-app/artifact-abc:1.0.0
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The OCID of the artifact.  Example: ocid1.genericartifact.oc1..exampleuniqueID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The OCID of the repository.
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// The SHA256 digest for the artifact. When you upload an artifact to the repository, a SHA256 digest is calculated and added to the artifact properties.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`

	// The size of the artifact in bytes.
	SizeInBytes *string `json:"sizeInBytes,omitempty" tf:"size_in_bytes,omitempty"`

	// The current state of the artifact.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// An RFC 3339 timestamp indicating when the repository was created.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// A user-defined string to describe the artifact version.  Example: 1.1.0 or 1.2-beta-2
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type GenericArtifactParameters struct {

	// The OCID of the artifact.  Example: ocid1.genericartifact.oc1..exampleuniqueID
	// +kubebuilder:validation:Required
	ArtifactID *string `json:"artifactId" tf:"artifact_id,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`
}

// GenericArtifactSpec defines the desired state of GenericArtifact
type GenericArtifactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GenericArtifactParameters `json:"forProvider"`
}

// GenericArtifactStatus defines the observed state of GenericArtifact.
type GenericArtifactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GenericArtifactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GenericArtifact is the Schema for the GenericArtifacts API. Provides the Generic Artifact resource in Oracle Cloud Infrastructure Artifacts service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type GenericArtifact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GenericArtifactSpec   `json:"spec"`
	Status            GenericArtifactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GenericArtifactList contains a list of GenericArtifacts
type GenericArtifactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GenericArtifact `json:"items"`
}

// Repository type metadata.
var (
	GenericArtifact_Kind             = "GenericArtifact"
	GenericArtifact_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GenericArtifact_Kind}.String()
	GenericArtifact_KindAPIVersion   = GenericArtifact_Kind + "." + CRDGroupVersion.String()
	GenericArtifact_GroupVersionKind = CRDGroupVersion.WithKind(GenericArtifact_Kind)
)

func init() {
	SchemeBuilder.Register(&GenericArtifact{}, &GenericArtifactList{})
}
