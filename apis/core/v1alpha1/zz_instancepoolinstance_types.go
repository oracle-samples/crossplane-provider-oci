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

type InstancePoolInstanceObservation struct {

	// The availability domain the instance is running in.
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	// The OCID of the compartment that contains the instance.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The fault domain the instance is running in.
	FaultDomain *string `json:"faultDomain,omitempty" tf:"fault_domain,omitempty"`

	// The OCID of the instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The OCID of the instance configuration used to create the instance.
	InstanceConfigurationID *string `json:"instanceConfigurationId,omitempty" tf:"instance_configuration_id,omitempty"`

	// The load balancer backends that are configured for the instance pool instance.
	LoadBalancerBackends []LoadBalancerBackendsObservation `json:"loadBalancerBackends,omitempty" tf:"load_balancer_backends,omitempty"`

	// The region that contains the availability domain the instance is running in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.
	Shape *string `json:"shape,omitempty" tf:"shape,omitempty"`

	// The lifecycle state of the instance. Refer to lifecycleState in the Instance resource.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the instance pool instance was created, in the format defined by RFC3339. Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type InstancePoolInstanceParameters struct {

	// +kubebuilder:validation:Optional
	AutoTerminateInstanceOnDelete *bool `json:"autoTerminateInstanceOnDelete,omitempty" tf:"auto_terminate_instance_on_delete,omitempty"`

	// +kubebuilder:validation:Optional
	DecrementSizeOnDelete *bool `json:"decrementSizeOnDelete,omitempty" tf:"decrement_size_on_delete,omitempty"`

	// The OCID of the instance.
	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// The OCID of the instance pool.
	// +kubebuilder:validation:Required
	InstancePoolID *string `json:"instancePoolId" tf:"instance_pool_id,omitempty"`
}

type LoadBalancerBackendsObservation struct {

	// The health of the backend as observed by the load balancer.
	BackendHealthStatus *string `json:"backendHealthStatus,omitempty" tf:"backend_health_status,omitempty"`

	// The name of the backend in the backend set.
	BackendName *string `json:"backendName,omitempty" tf:"backend_name,omitempty"`

	// The name of the backend set on the load balancer.
	BackendSetName *string `json:"backendSetName,omitempty" tf:"backend_set_name,omitempty"`

	// The OCID of the load balancer attached to the instance pool.
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// The lifecycle state of the instance. Refer to lifecycleState in the Instance resource.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type LoadBalancerBackendsParameters struct {
}

// InstancePoolInstanceSpec defines the desired state of InstancePoolInstance
type InstancePoolInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstancePoolInstanceParameters `json:"forProvider"`
}

// InstancePoolInstanceStatus defines the observed state of InstancePoolInstance.
type InstancePoolInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstancePoolInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstancePoolInstance is the Schema for the InstancePoolInstances API. Provides the Instance Pool Instance resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type InstancePoolInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstancePoolInstanceSpec   `json:"spec"`
	Status            InstancePoolInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstancePoolInstanceList contains a list of InstancePoolInstances
type InstancePoolInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstancePoolInstance `json:"items"`
}

// Repository type metadata.
var (
	InstancePoolInstance_Kind             = "InstancePoolInstance"
	InstancePoolInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstancePoolInstance_Kind}.String()
	InstancePoolInstance_KindAPIVersion   = InstancePoolInstance_Kind + "." + CRDGroupVersion.String()
	InstancePoolInstance_GroupVersionKind = CRDGroupVersion.WithKind(InstancePoolInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&InstancePoolInstance{}, &InstancePoolInstanceList{})
}