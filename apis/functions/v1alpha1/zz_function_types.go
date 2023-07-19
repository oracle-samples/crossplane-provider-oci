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

type FunctionObservation struct {

	// The OCID of the compartment that contains the function.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// The OCID of the function.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The base https invoke URL to set on a client in order to invoke a function. This URL will never change over the lifetime of the function and can be cached.
	InvokeEndpoint *string `json:"invokeEndpoint,omitempty" tf:"invoke_endpoint,omitempty"`

	// The current state of the function.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The time the function was created, expressed in RFC 3339 timestamp format.  Example: 2018-09-12T22:47:12.613Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The time the function was updated, expressed in RFC 3339 timestamp format.  Example: 2018-09-12T22:47:12.613Z
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type FunctionParameters struct {

	// The OCID of the application this function belongs to.
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// (Updatable) Function configuration. These values are passed on to the function as environment variables, this overrides application configuration values. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: {"MY_FUNCTION_CONFIG": "ConfVal"}
	// +kubebuilder:validation:Optional
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// The display name of the function. The display name must be unique within the application containing the function. Avoid entering confidential information.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The qualified name of the Docker image to use in the function, including the image tag. The image should be in the Oracle Cloud Infrastructure Registry that is in the same region as the function itself. This field must be updated if image_digest is updated. Example: phx.ocir.io/ten/functions/function:0.0.1
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// (Updatable) The image digest for the version of the image that will be pulled when invoking this function. If no value is specified, the digest currently associated with the image in the Oracle Cloud Infrastructure Registry will be used. This field must be updated if image is updated. Example: sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7
	// +kubebuilder:validation:Optional
	ImageDigest *string `json:"imageDigest,omitempty" tf:"image_digest,omitempty"`

	// (Updatable) Maximum usable memory for the function (MiB).
	// +kubebuilder:validation:Required
	MemoryInMbs *string `json:"memoryInMbs" tf:"memory_in_mbs,omitempty"`

	// (Updatable) Define the strategy for provisioned concurrency for the function.
	// +kubebuilder:validation:Optional
	ProvisionedConcurrencyConfig []ProvisionedConcurrencyConfigParameters `json:"provisionedConcurrencyConfig,omitempty" tf:"provisioned_concurrency_config,omitempty"`

	// The source details for the Function. The function can be created from various sources.
	// +kubebuilder:validation:Optional
	SourceDetails []SourceDetailsParameters `json:"sourceDetails,omitempty" tf:"source_details,omitempty"`

	// (Updatable) Timeout for executions of the function. Value in seconds.
	// +kubebuilder:validation:Optional
	TimeoutInSeconds *float64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`

	// (Updatable) Define the tracing configuration for a function.
	// +kubebuilder:validation:Optional
	TraceConfig []FunctionTraceConfigParameters `json:"traceConfig,omitempty" tf:"trace_config,omitempty"`
}

type FunctionTraceConfigObservation struct {
}

type FunctionTraceConfigParameters struct {

	// (Updatable) Define if tracing is enabled for the resource.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type ProvisionedConcurrencyConfigObservation struct {
}

type ProvisionedConcurrencyConfigParameters struct {

	// (Updatable)
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// (Updatable) The strategy for provisioned concurrency to be used.
	// +kubebuilder:validation:Required
	Strategy *string `json:"strategy" tf:"strategy,omitempty"`
}

type SourceDetailsObservation struct {
}

type SourceDetailsParameters struct {

	// The OCID of the PbfListing this function is sourced from.
	// +kubebuilder:validation:Required
	PbfListingID *string `json:"pbfListingId" tf:"pbf_listing_id,omitempty"`

	// Type of the Function Source. Possible values: PRE_BUILT_FUNCTIONS.
	// +kubebuilder:validation:Required
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API. Provides the Function resource in Oracle Cloud Infrastructure Functions service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
