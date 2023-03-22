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

type DrgRouteDistributionStatementObservation struct {

	// The Oracle-assigned ID of the route distribution statement.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DrgRouteDistributionStatementParameters struct {

	// Accept: import/export the route "as is"
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// The OCID of the route distribution.
	// +crossplane:generate:reference:type=DrgRouteDistribution
	// +kubebuilder:validation:Optional
	DrgRouteDistributionID *string `json:"drgRouteDistributionId,omitempty" tf:"drg_route_distribution_id,omitempty"`

	// Reference to a DrgRouteDistribution to populate drgRouteDistributionId.
	// +kubebuilder:validation:Optional
	DrgRouteDistributionIDRef *v1.Reference `json:"drgRouteDistributionIdRef,omitempty" tf:"-"`

	// Selector for a DrgRouteDistribution to populate drgRouteDistributionId.
	// +kubebuilder:validation:Optional
	DrgRouteDistributionIDSelector *v1.Selector `json:"drgRouteDistributionIdSelector,omitempty" tf:"-"`

	// (Updatable) The action is applied only if all of the match criteria are met. MATCH_ALL match type implies any input is considered a match.
	// * attachment_type -  The type of the network resource to be included in this match. A match for a network type implies that all DRG attachments of that type insert routes into the table.
	// * drg_attachment_id -  The OCID of the DRG attachment.
	// * match_type -  (Updatable) The type of the match criteria for a route distribution statement.
	// +kubebuilder:validation:Required
	MatchCriteria []MatchCriteriaParameters `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// (Updatable) This field is used to specify the priority of each statement in a route distribution. The priority will be represented as a number between 0 and 65535 where a lower number indicates a higher priority. When a route is processed, statements are applied in the order defined by their priority. The first matching rule dictates the action that will be taken on the route.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`
}

type MatchCriteriaObservation struct {
}

type MatchCriteriaParameters struct {

	// The type of the network resource to be included in this match. A match for a network type implies that all DRG attachments of that type insert routes into the table.
	// +kubebuilder:validation:Optional
	AttachmentType *string `json:"attachmentType,omitempty" tf:"attachment_type,omitempty"`

	// The OCID of the DRG attachment.
	// +kubebuilder:validation:Optional
	DrgAttachmentID *string `json:"drgAttachmentId,omitempty" tf:"drg_attachment_id,omitempty"`

	// The type of the match criteria for a route distribution statement. Can take three values- DRG_ATTACHMENT_TYPE, DRG_ATTACHMENT_ID and MATCH_ALL.
	// +kubebuilder:validation:Optional
	MatchType *string `json:"matchType,omitempty" tf:"match_type,omitempty"`
}

// DrgRouteDistributionStatementSpec defines the desired state of DrgRouteDistributionStatement
type DrgRouteDistributionStatementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DrgRouteDistributionStatementParameters `json:"forProvider"`
}

// DrgRouteDistributionStatementStatus defines the observed state of DrgRouteDistributionStatement.
type DrgRouteDistributionStatementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DrgRouteDistributionStatementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DrgRouteDistributionStatement is the Schema for the DrgRouteDistributionStatements API. Provides the Drg Route Distribution Statement resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type DrgRouteDistributionStatement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DrgRouteDistributionStatementSpec   `json:"spec"`
	Status            DrgRouteDistributionStatementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DrgRouteDistributionStatementList contains a list of DrgRouteDistributionStatements
type DrgRouteDistributionStatementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DrgRouteDistributionStatement `json:"items"`
}

// Repository type metadata.
var (
	DrgRouteDistributionStatement_Kind             = "DrgRouteDistributionStatement"
	DrgRouteDistributionStatement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DrgRouteDistributionStatement_Kind}.String()
	DrgRouteDistributionStatement_KindAPIVersion   = DrgRouteDistributionStatement_Kind + "." + CRDGroupVersion.String()
	DrgRouteDistributionStatement_GroupVersionKind = CRDGroupVersion.WithKind(DrgRouteDistributionStatement_Kind)
)

func init() {
	SchemeBuilder.Register(&DrgRouteDistributionStatement{}, &DrgRouteDistributionStatementList{})
}
