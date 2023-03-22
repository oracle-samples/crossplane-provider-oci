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

type DrgAllAttachmentsObservation struct {

	// The Oracle-assigned ID of the DRG attachment
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DrgAllAttachmentsParameters struct {
}

type DrgAttachmentsListObservation struct {

	// The list of drg_attachments.
	DrgAllAttachments []DrgAllAttachmentsObservation `json:"drgAllAttachments,omitempty" tf:"drg_all_attachments,omitempty"`

	// The Oracle-assigned ID of the DRG attachment
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DrgAttachmentsListParameters struct {

	// The type for the network resource attached to the DRG.
	// +kubebuilder:validation:Optional
	AttachmentType *string `json:"attachmentType,omitempty" tf:"attachment_type,omitempty"`

	// The OCID of the DRG.
	// +crossplane:generate:reference:type=Drg
	// +kubebuilder:validation:Optional
	DrgID *string `json:"drgId,omitempty" tf:"drg_id,omitempty"`

	// Reference to a Drg to populate drgId.
	// +kubebuilder:validation:Optional
	DrgIDRef *v1.Reference `json:"drgIdRef,omitempty" tf:"-"`

	// Selector for a Drg to populate drgId.
	// +kubebuilder:validation:Optional
	DrgIDSelector *v1.Selector `json:"drgIdSelector,omitempty" tf:"-"`

	// Whether the DRG attachment lives in a different tenancy than the DRG.
	// +kubebuilder:validation:Optional
	IsCrossTenancy *bool `json:"isCrossTenancy,omitempty" tf:"is_cross_tenancy,omitempty"`
}

// DrgAttachmentsListSpec defines the desired state of DrgAttachmentsList
type DrgAttachmentsListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DrgAttachmentsListParameters `json:"forProvider"`
}

// DrgAttachmentsListStatus defines the observed state of DrgAttachmentsList.
type DrgAttachmentsListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DrgAttachmentsListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DrgAttachmentsList is the Schema for the DrgAttachmentsLists API. Provides the Drg Attachments List resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type DrgAttachmentsList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DrgAttachmentsListSpec   `json:"spec"`
	Status            DrgAttachmentsListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DrgAttachmentsListList contains a list of DrgAttachmentsLists
type DrgAttachmentsListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DrgAttachmentsList `json:"items"`
}

// Repository type metadata.
var (
	DrgAttachmentsList_Kind             = "DrgAttachmentsList"
	DrgAttachmentsList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DrgAttachmentsList_Kind}.String()
	DrgAttachmentsList_KindAPIVersion   = DrgAttachmentsList_Kind + "." + CRDGroupVersion.String()
	DrgAttachmentsList_GroupVersionKind = CRDGroupVersion.WithKind(DrgAttachmentsList_Kind)
)

func init() {
	SchemeBuilder.Register(&DrgAttachmentsList{}, &DrgAttachmentsListList{})
}
