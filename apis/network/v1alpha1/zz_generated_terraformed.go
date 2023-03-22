/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this FirewallNetworkFirewall
func (mg *FirewallNetworkFirewall) GetTerraformResourceType() string {
	return "oci_network_firewall_network_firewall"
}

// GetConnectionDetailsMapping for this FirewallNetworkFirewall
func (tr *FirewallNetworkFirewall) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FirewallNetworkFirewall
func (tr *FirewallNetworkFirewall) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FirewallNetworkFirewall
func (tr *FirewallNetworkFirewall) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FirewallNetworkFirewall
func (tr *FirewallNetworkFirewall) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FirewallNetworkFirewall
func (tr *FirewallNetworkFirewall) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FirewallNetworkFirewall
func (tr *FirewallNetworkFirewall) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this FirewallNetworkFirewall using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FirewallNetworkFirewall) LateInitialize(attrs []byte) (bool, error) {
	params := &FirewallNetworkFirewallParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FirewallNetworkFirewall) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this FirewallNetworkFirewallPolicy
func (mg *FirewallNetworkFirewallPolicy) GetTerraformResourceType() string {
	return "oci_network_firewall_network_firewall_policy"
}

// GetConnectionDetailsMapping for this FirewallNetworkFirewallPolicy
func (tr *FirewallNetworkFirewallPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FirewallNetworkFirewallPolicy
func (tr *FirewallNetworkFirewallPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FirewallNetworkFirewallPolicy
func (tr *FirewallNetworkFirewallPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FirewallNetworkFirewallPolicy
func (tr *FirewallNetworkFirewallPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FirewallNetworkFirewallPolicy
func (tr *FirewallNetworkFirewallPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FirewallNetworkFirewallPolicy
func (tr *FirewallNetworkFirewallPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this FirewallNetworkFirewallPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FirewallNetworkFirewallPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &FirewallNetworkFirewallPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FirewallNetworkFirewallPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LoadBalancerBackend
func (mg *LoadBalancerBackend) GetTerraformResourceType() string {
	return "oci_network_load_balancer_backend"
}

// GetConnectionDetailsMapping for this LoadBalancerBackend
func (tr *LoadBalancerBackend) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LoadBalancerBackend
func (tr *LoadBalancerBackend) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LoadBalancerBackend
func (tr *LoadBalancerBackend) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LoadBalancerBackend
func (tr *LoadBalancerBackend) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LoadBalancerBackend
func (tr *LoadBalancerBackend) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LoadBalancerBackend
func (tr *LoadBalancerBackend) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LoadBalancerBackend using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LoadBalancerBackend) LateInitialize(attrs []byte) (bool, error) {
	params := &LoadBalancerBackendParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LoadBalancerBackend) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LoadBalancerBackendSet
func (mg *LoadBalancerBackendSet) GetTerraformResourceType() string {
	return "oci_network_load_balancer_backend_set"
}

// GetConnectionDetailsMapping for this LoadBalancerBackendSet
func (tr *LoadBalancerBackendSet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LoadBalancerBackendSet
func (tr *LoadBalancerBackendSet) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LoadBalancerBackendSet
func (tr *LoadBalancerBackendSet) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LoadBalancerBackendSet
func (tr *LoadBalancerBackendSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LoadBalancerBackendSet
func (tr *LoadBalancerBackendSet) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LoadBalancerBackendSet
func (tr *LoadBalancerBackendSet) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LoadBalancerBackendSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LoadBalancerBackendSet) LateInitialize(attrs []byte) (bool, error) {
	params := &LoadBalancerBackendSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LoadBalancerBackendSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LoadBalancerListener
func (mg *LoadBalancerListener) GetTerraformResourceType() string {
	return "oci_network_load_balancer_listener"
}

// GetConnectionDetailsMapping for this LoadBalancerListener
func (tr *LoadBalancerListener) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LoadBalancerListener
func (tr *LoadBalancerListener) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LoadBalancerListener
func (tr *LoadBalancerListener) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LoadBalancerListener
func (tr *LoadBalancerListener) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LoadBalancerListener
func (tr *LoadBalancerListener) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LoadBalancerListener
func (tr *LoadBalancerListener) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LoadBalancerListener using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LoadBalancerListener) LateInitialize(attrs []byte) (bool, error) {
	params := &LoadBalancerListenerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LoadBalancerListener) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LoadBalancerNetworkLoadBalancer
func (mg *LoadBalancerNetworkLoadBalancer) GetTerraformResourceType() string {
	return "oci_network_load_balancer_network_load_balancer"
}

// GetConnectionDetailsMapping for this LoadBalancerNetworkLoadBalancer
func (tr *LoadBalancerNetworkLoadBalancer) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LoadBalancerNetworkLoadBalancer
func (tr *LoadBalancerNetworkLoadBalancer) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LoadBalancerNetworkLoadBalancer
func (tr *LoadBalancerNetworkLoadBalancer) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LoadBalancerNetworkLoadBalancer
func (tr *LoadBalancerNetworkLoadBalancer) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LoadBalancerNetworkLoadBalancer
func (tr *LoadBalancerNetworkLoadBalancer) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LoadBalancerNetworkLoadBalancer
func (tr *LoadBalancerNetworkLoadBalancer) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LoadBalancerNetworkLoadBalancer using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LoadBalancerNetworkLoadBalancer) LateInitialize(attrs []byte) (bool, error) {
	params := &LoadBalancerNetworkLoadBalancerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LoadBalancerNetworkLoadBalancer) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LoadBalancerNetworkLoadBalancersBackendSetsUnified
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetTerraformResourceType() string {
	return "oci_network_load_balancer_network_load_balancers_backend_sets_unified"
}

// GetConnectionDetailsMapping for this LoadBalancerNetworkLoadBalancersBackendSetsUnified
func (tr *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LoadBalancerNetworkLoadBalancersBackendSetsUnified
func (tr *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LoadBalancerNetworkLoadBalancersBackendSetsUnified
func (tr *LoadBalancerNetworkLoadBalancersBackendSetsUnified) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LoadBalancerNetworkLoadBalancersBackendSetsUnified
func (tr *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LoadBalancerNetworkLoadBalancersBackendSetsUnified
func (tr *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LoadBalancerNetworkLoadBalancersBackendSetsUnified
func (tr *LoadBalancerNetworkLoadBalancersBackendSetsUnified) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LoadBalancerNetworkLoadBalancersBackendSetsUnified using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LoadBalancerNetworkLoadBalancersBackendSetsUnified) LateInitialize(attrs []byte) (bool, error) {
	params := &LoadBalancerNetworkLoadBalancersBackendSetsUnifiedParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetTerraformSchemaVersion() int {
	return 0
}
