/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=Secret;InjectedIdentity;Environment;Filesystem
	Source xpv1.CredentialsSource `json:"source"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a Oracle Cloud Infrastructure (OCI) provider.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,oci}
// +kubebuilder:subresource:status
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,oci}
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// GetProviderConfigReference of this ProviderConfigUsage.
func (pc *ProviderConfigUsage) GetProviderConfigReference() xpv1.Reference {
	return pc.ProviderConfigReference
}

// GetResourceReference of this ProviderConfigUsage.
func (pc *ProviderConfigUsage) GetResourceReference() xpv1.TypedReference {
	return pc.ResourceReference
}

// SetProviderConfigReference of this ProviderConfigUsage.
func (pc *ProviderConfigUsage) SetProviderConfigReference(r xpv1.Reference) {
	pc.ProviderConfigReference = r
}

// SetResourceReference of this ProviderConfigUsage.
func (pc *ProviderConfigUsage) SetResourceReference(r xpv1.TypedReference) {
	pc.ResourceReference = r
}

// ProviderConfigGroupKind is the GroupKind for ProviderConfig
var ProviderConfigGroupKind = schema.GroupKind{
	Group: Group,
	Kind:  "ProviderConfig",
}.String()

// ProviderConfigGroupVersionKind is the GroupVersionKind for ProviderConfig  
var ProviderConfigGroupVersionKind = schema.GroupVersionKind{
	Group:   Group,
	Version: Version,
	Kind:    "ProviderConfig",
}

// ProviderConfigUsageListGroupVersionKind is the GroupVersionKind for ProviderConfigUsageList
var ProviderConfigUsageListGroupVersionKind = schema.GroupVersionKind{
	Group:   Group,
	Version: Version,
	Kind:    "ProviderConfigUsageList",
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage.
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}