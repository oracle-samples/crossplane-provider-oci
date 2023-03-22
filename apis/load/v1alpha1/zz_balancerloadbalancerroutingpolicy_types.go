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

type ActionsObservation struct {
}

type ActionsParameters struct {

	// (Updatable) Name of the backend set the listener will forward the traffic to.  Example: backendSetForImages
	// +crossplane:generate:reference:type=BalancerBackendSet
	// +kubebuilder:validation:Optional
	BackendSetName *string `json:"backendSetName,omitempty" tf:"backend_set_name,omitempty"`

	// Reference to a BalancerBackendSet to populate backendSetName.
	// +kubebuilder:validation:Optional
	BackendSetNameRef *v1.Reference `json:"backendSetNameRef,omitempty" tf:"-"`

	// Selector for a BalancerBackendSet to populate backendSetName.
	// +kubebuilder:validation:Optional
	BackendSetNameSelector *v1.Selector `json:"backendSetNameSelector,omitempty" tf:"-"`

	// The name for this list of routing rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: example_routing_rules
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type BalancerLoadBalancerRoutingPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type BalancerLoadBalancerRoutingPolicyParameters struct {

	// (Updatable) The version of the language in which condition of rules are composed.
	// +kubebuilder:validation:Required
	ConditionLanguageVersion *string `json:"conditionLanguageVersion" tf:"condition_language_version,omitempty"`

	// The OCID of the load balancer to add the routing policy rule list to.
	// +crossplane:generate:reference:type=BalancerLoadBalancer
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a BalancerLoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a BalancerLoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// The name for this list of routing rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: example_routing_rules
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// (Updatable) The list of routing rules.
	// +kubebuilder:validation:Required
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// (Updatable) A list of actions to be applied when conditions of the routing rule are met.
	// +kubebuilder:validation:Required
	Actions []ActionsParameters `json:"actions" tf:"actions,omitempty"`

	// (Updatable) A routing rule to evaluate defined conditions against the incoming HTTP request and perform an action.
	// +kubebuilder:validation:Required
	Condition *string `json:"condition" tf:"condition,omitempty"`

	// The name for this list of routing rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: example_routing_rules
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// BalancerLoadBalancerRoutingPolicySpec defines the desired state of BalancerLoadBalancerRoutingPolicy
type BalancerLoadBalancerRoutingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerLoadBalancerRoutingPolicyParameters `json:"forProvider"`
}

// BalancerLoadBalancerRoutingPolicyStatus defines the observed state of BalancerLoadBalancerRoutingPolicy.
type BalancerLoadBalancerRoutingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerLoadBalancerRoutingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerLoadBalancerRoutingPolicy is the Schema for the BalancerLoadBalancerRoutingPolicys API. Provides the Load Balancer Routing Policy resource in Oracle Cloud Infrastructure Load Balancer service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type BalancerLoadBalancerRoutingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BalancerLoadBalancerRoutingPolicySpec   `json:"spec"`
	Status            BalancerLoadBalancerRoutingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerLoadBalancerRoutingPolicyList contains a list of BalancerLoadBalancerRoutingPolicys
type BalancerLoadBalancerRoutingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerLoadBalancerRoutingPolicy `json:"items"`
}

// Repository type metadata.
var (
	BalancerLoadBalancerRoutingPolicy_Kind             = "BalancerLoadBalancerRoutingPolicy"
	BalancerLoadBalancerRoutingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerLoadBalancerRoutingPolicy_Kind}.String()
	BalancerLoadBalancerRoutingPolicy_KindAPIVersion   = BalancerLoadBalancerRoutingPolicy_Kind + "." + CRDGroupVersion.String()
	BalancerLoadBalancerRoutingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BalancerLoadBalancerRoutingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerLoadBalancerRoutingPolicy{}, &BalancerLoadBalancerRoutingPolicyList{})
}
