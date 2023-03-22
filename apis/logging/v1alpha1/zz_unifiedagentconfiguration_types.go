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

type DestinationObservation struct {
}

type DestinationParameters struct {

	// (Updatable) The OCID of the resource.
	// +crossplane:generate:reference:type=Log
	// +kubebuilder:validation:Optional
	LogObjectID *string `json:"logObjectId,omitempty" tf:"log_object_id,omitempty"`

	// Reference to a Log to populate logObjectId.
	// +kubebuilder:validation:Optional
	LogObjectIDRef *v1.Reference `json:"logObjectIdRef,omitempty" tf:"-"`

	// Selector for a Log to populate logObjectId.
	// +kubebuilder:validation:Optional
	LogObjectIDSelector *v1.Selector `json:"logObjectIdSelector,omitempty" tf:"-"`
}

type GroupAssociationObservation struct {
}

type GroupAssociationParameters struct {

	// (Updatable) list of group/dynamic group ids associated with this configuration.
	// +kubebuilder:validation:Optional
	GroupList []*string `json:"groupList,omitempty" tf:"group_list,omitempty"`
}

type ParserObservation struct {
}

type ParserParameters struct {

	// (Applicable when parser_type=CSV | TSV) (Updatable)
	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// (Applicable when parser_type=REGEXP) (Updatable)
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) Specify time field for the event time. If the event doesn't have this field, the current time is used.
	// +kubebuilder:validation:Optional
	FieldTimeKey *string `json:"fieldTimeKey,omitempty" tf:"field_time_key,omitempty"`

	// (Applicable when parser_type=MULTILINE) (Updatable)
	// +kubebuilder:validation:Optional
	Format []*string `json:"format,omitempty" tf:"format,omitempty"`

	// (Applicable when parser_type=MULTILINE) (Updatable)
	// +kubebuilder:validation:Optional
	FormatFirstline *string `json:"formatFirstline,omitempty" tf:"format_firstline,omitempty"`

	// (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable)
	// +kubebuilder:validation:Optional
	GrokFailureKey *string `json:"grokFailureKey,omitempty" tf:"grok_failure_key,omitempty"`

	// (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable)
	// +kubebuilder:validation:Optional
	GrokNameKey *string `json:"grokNameKey,omitempty" tf:"grok_name_key,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) If true, use Fluent::EventTime.now(current time) as a timestamp when time_key is specified.
	// +kubebuilder:validation:Optional
	IsEstimateCurrentEvent *bool `json:"isEstimateCurrentEvent,omitempty" tf:"is_estimate_current_event,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) If true, keep time field in the record.
	// +kubebuilder:validation:Optional
	IsKeepTimeKey *bool `json:"isKeepTimeKey,omitempty" tf:"is_keep_time_key,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) If true, an empty string field is replaced with nil.
	// +kubebuilder:validation:Optional
	IsNullEmptyString *bool `json:"isNullEmptyString,omitempty" tf:"is_null_empty_string,omitempty"`

	// (Applicable when parser_type=SYSLOG) (Updatable)
	// +kubebuilder:validation:Optional
	IsSupportColonlessIdent *bool `json:"isSupportColonlessIdent,omitempty" tf:"is_support_colonless_ident,omitempty"`

	// (Applicable when parser_type=SYSLOG) (Updatable)
	// +kubebuilder:validation:Optional
	IsWithPriority *bool `json:"isWithPriority,omitempty" tf:"is_with_priority,omitempty"`

	// (Applicable when parser_type=CSV | TSV) (Updatable)
	// +kubebuilder:validation:Optional
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`

	// (Applicable when parser_type=SYSLOG) (Updatable)
	// +kubebuilder:validation:Optional
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format,omitempty"`

	// (Applicable when parser_type=NONE) (Updatable)
	// +kubebuilder:validation:Optional
	MessageKey *string `json:"messageKey,omitempty" tf:"message_key,omitempty"`

	// (Applicable when parser_type=MULTILINE_GROK) (Updatable)
	// +kubebuilder:validation:Optional
	MultiLineStartRegexp *string `json:"multiLineStartRegexp,omitempty" tf:"multi_line_start_regexp,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) Specify the null value pattern.
	// +kubebuilder:validation:Optional
	NullValuePattern *string `json:"nullValuePattern,omitempty" tf:"null_value_pattern,omitempty"`

	// (Updatable) Type of fluent parser.
	// +kubebuilder:validation:Required
	ParserType *string `json:"parserType" tf:"parser_type,omitempty"`

	// (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable)
	// +kubebuilder:validation:Optional
	Patterns []PatternsParameters `json:"patterns,omitempty" tf:"patterns,omitempty"`

	// (Applicable when parser_type=SYSLOG) (Updatable)
	// +kubebuilder:validation:Optional
	Rfc5424TimeFormat *string `json:"rfc5424timeFormat,omitempty" tf:"rfc5424time_format,omitempty"`

	// (Applicable when parser_type=SYSLOG) (Updatable)
	// +kubebuilder:validation:Optional
	SyslogParserType *string `json:"syslogParserType,omitempty" tf:"syslog_parser_type,omitempty"`

	// (Applicable when parser_type=JSON | REGEXP | SYSLOG) (Updatable)
	// +kubebuilder:validation:Optional
	TimeFormat *string `json:"timeFormat,omitempty" tf:"time_format,omitempty"`

	// (Applicable when parser_type=JSON) (Updatable)
	// +kubebuilder:validation:Optional
	TimeType *string `json:"timeType,omitempty" tf:"time_type,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
	// +kubebuilder:validation:Optional
	TimeoutInMilliseconds *float64 `json:"timeoutInMilliseconds,omitempty" tf:"timeout_in_milliseconds,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) Specify types for converting a field into another type.
	// +kubebuilder:validation:Optional
	Types map[string]*string `json:"types,omitempty" tf:"types,omitempty"`
}

type PatternsObservation struct {
}

type PatternsParameters struct {

	// (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Process value using the specified format. This is available only when time_type is a string.
	// +kubebuilder:validation:Optional
	FieldTimeFormat *string `json:"fieldTimeFormat,omitempty" tf:"field_time_format,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) Specify time field for the event time. If the event doesn't have this field, the current time is used.
	// +kubebuilder:validation:Optional
	FieldTimeKey *string `json:"fieldTimeKey,omitempty" tf:"field_time_key,omitempty"`

	// (Applicable when parser_type=GROK | MULTILINE_GROK) (Updatable) Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
	// +kubebuilder:validation:Optional
	FieldTimeZone *string `json:"fieldTimeZone,omitempty" tf:"field_time_zone,omitempty"`

	// (Updatable) unique name for the source
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The grok pattern.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type ServiceConfigurationObservation struct {
}

type ServiceConfigurationParameters struct {

	// (Updatable) Type of Unified Agent service configuration.
	// +kubebuilder:validation:Required
	ConfigurationType *string `json:"configurationType" tf:"configuration_type,omitempty"`

	// (Updatable) Logging destination object.
	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// (Updatable)
	// +kubebuilder:validation:Required
	Sources []SourcesParameters `json:"sources" tf:"sources,omitempty"`
}

type SourcesObservation struct {
}

type SourcesParameters struct {

	// (Applicable when source_type=WINDOWS_EVENT_LOG) (Updatable)
	// +kubebuilder:validation:Optional
	Channels []*string `json:"channels,omitempty" tf:"channels,omitempty"`

	// (Updatable) unique name for the source
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable) source parser object.
	// +kubebuilder:validation:Optional
	Parser []ParserParameters `json:"parser,omitempty" tf:"parser,omitempty"`

	// (Applicable when source_type=LOG_TAIL) (Updatable)
	// +kubebuilder:validation:Optional
	Paths []*string `json:"paths,omitempty" tf:"paths,omitempty"`

	// (Updatable) Unified schema logging source type.
	// +kubebuilder:validation:Required
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

type UnifiedAgentConfigurationObservation struct {

	// State of unified agent service configuration.
	ConfigurationState *string `json:"configurationState,omitempty" tf:"configuration_state,omitempty"`

	// The OCID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The pipeline state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Time the resource was created.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// Time the resource was last modified.
	TimeLastModified *string `json:"timeLastModified,omitempty" tf:"time_last_modified,omitempty"`
}

type UnifiedAgentConfigurationParameters struct {

	// (Updatable) The OCID of the compartment that the resource belongs to.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Description for this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Groups using the configuration.
	// +kubebuilder:validation:Optional
	GroupAssociation []GroupAssociationParameters `json:"groupAssociation,omitempty" tf:"group_association,omitempty"`

	// (Updatable) Whether or not this resource is currently enabled.
	// +kubebuilder:validation:Required
	IsEnabled *bool `json:"isEnabled" tf:"is_enabled,omitempty"`

	// (Updatable) Top level Unified Agent service configuration object.
	// +kubebuilder:validation:Required
	ServiceConfiguration []ServiceConfigurationParameters `json:"serviceConfiguration" tf:"service_configuration,omitempty"`
}

// UnifiedAgentConfigurationSpec defines the desired state of UnifiedAgentConfiguration
type UnifiedAgentConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UnifiedAgentConfigurationParameters `json:"forProvider"`
}

// UnifiedAgentConfigurationStatus defines the observed state of UnifiedAgentConfiguration.
type UnifiedAgentConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UnifiedAgentConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UnifiedAgentConfiguration is the Schema for the UnifiedAgentConfigurations API. Provides the Unified Agent Configuration resource in Oracle Cloud Infrastructure Logging service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type UnifiedAgentConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UnifiedAgentConfigurationSpec   `json:"spec"`
	Status            UnifiedAgentConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UnifiedAgentConfigurationList contains a list of UnifiedAgentConfigurations
type UnifiedAgentConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UnifiedAgentConfiguration `json:"items"`
}

// Repository type metadata.
var (
	UnifiedAgentConfiguration_Kind             = "UnifiedAgentConfiguration"
	UnifiedAgentConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UnifiedAgentConfiguration_Kind}.String()
	UnifiedAgentConfiguration_KindAPIVersion   = UnifiedAgentConfiguration_Kind + "." + CRDGroupVersion.String()
	UnifiedAgentConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(UnifiedAgentConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&UnifiedAgentConfiguration{}, &UnifiedAgentConfigurationList{})
}
