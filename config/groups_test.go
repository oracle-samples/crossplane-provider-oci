/*
 * Copyright (c) 2023 Oracle and/or its affiliates
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"testing"

	"github.com/crossplane/upjet/pkg/config"
)

// TestGroupMapCompleteness verifies that all resources in ExternalNameConfigs
// have corresponding entries in GroupMap
func TestGroupMapCompleteness(t *testing.T) {
	missing := []string{}
	
	for resourceName := range ExternalNameConfigs {
		if _, exists := GroupMap[resourceName]; !exists {
			missing = append(missing, resourceName)
		}
	}
	
	if len(missing) > 0 {
		t.Errorf("GroupMap is missing entries for %d resources: %v", len(missing), missing)
	}
}

// TestGroupMapValidity verifies that all GroupMap entries return valid group and kind values
func TestGroupMapValidity(t *testing.T) {
	for resourceName, calculator := range GroupMap {
		group, kind := calculator(resourceName)
		
		// Verify group is not empty
		if group == "" {
			t.Errorf("Resource %s has empty group", resourceName)
		}
		
		// Verify kind is not empty
		if kind == "" {
			t.Errorf("Resource %s has empty kind", resourceName)
		}
		
		// Verify group follows naming convention (lowercase, no underscores)
		for _, char := range group {
			if char == '_' {
				t.Errorf("Resource %s has invalid group name %s (contains underscore)", resourceName, group)
			}
		}
		
		// Verify kind follows naming convention (PascalCase)
		if len(kind) > 0 && kind[0] < 'A' || kind[0] > 'Z' {
			t.Errorf("Resource %s has invalid kind name %s (should start with uppercase)", resourceName, kind)
		}
	}
}

// TestServiceGroupings verifies that resources are grouped into expected services
func TestServiceGroupings(t *testing.T) {
	expectedGroups := map[string][]string{
		"compute": {
			"oci_core_instance",
			"oci_core_image",
			"oci_core_dedicated_vm_host",
			"oci_core_console_history",
			"oci_core_instance_configuration",
			"oci_core_instance_console_connection",
			"oci_core_instance_pool",
			"oci_core_instance_pool_instance",
			"oci_core_cluster_network",
			"oci_core_compute_capacity_reservation",
			"oci_core_compute_cluster",
			"oci_core_compute_image_capability_schema",
			"oci_core_shape_management",
			"oci_core_app_catalog_listing_resource_version_agreement",
			"oci_core_app_catalog_subscription",
		},
		"networking": {
			"oci_core_vcn",
			"oci_core_subnet",
			"oci_core_vnic_attachment",
			"oci_core_dhcp_options",
			"oci_core_vlan",
			"oci_core_internet_gateway",
			"oci_core_nat_gateway",
			"oci_core_service_gateway",
			"oci_core_network_security_group",
			"oci_core_network_security_group_security_rule",
			"oci_core_route_table",
			"oci_core_security_list",
			"oci_core_private_ip",
			"oci_core_public_ip",
			"oci_core_public_ip_pool",
			"oci_core_public_ip_pool_capacity",
			"oci_core_byoip_range",
			"oci_core_local_peering_gateway",
			"oci_core_remote_peering_connection",
			"oci_core_route_table_attachment",
			"oci_core_ipv6",
		},
		"blockstorage": {
			"oci_core_volume",
			"oci_core_volume_backup",
			"oci_core_volume_attachment",
			"oci_core_volume_backup_policy",
			"oci_core_volume_backup_policy_assignment",
			"oci_core_volume_group",
			"oci_core_volume_group_backup",
			"oci_core_boot_volume",
			"oci_core_boot_volume_backup",
		},
		"networkconnectivity": {
			"oci_core_drg",
			"oci_core_drg_attachment",
			"oci_core_drg_attachment_management",
			"oci_core_drg_attachments_list",
			"oci_core_drg_route_distribution",
			"oci_core_drg_route_distribution_statement",
			"oci_core_drg_route_table",
			"oci_core_drg_route_table_route_rule",
			"oci_core_cross_connect",
			"oci_core_cross_connect_group",
			"oci_core_virtual_circuit",
			"oci_core_cpe",
			"oci_core_ipsec",
			"oci_core_ipsec_connection_tunnel_management",
		},
	}
	
	for expectedGroup, resources := range expectedGroups {
		for _, resourceName := range resources {
			if calculator, exists := GroupMap[resourceName]; exists {
				group, _ := calculator(resourceName)
				if group != expectedGroup {
					t.Errorf("Resource %s expected to be in group %s, but got %s", resourceName, expectedGroup, group)
				}
			} else {
				t.Errorf("Resource %s not found in GroupMap", resourceName)
			}
		}
	}
}

// TestGroupKindOverrides verifies that the GroupKindOverrides function works correctly
func TestGroupKindOverrides(t *testing.T) {
	// Create a test resource configuration
	testResource := &config.Resource{
		Name: "oci_core_instance",
	}
	
	// Apply the GroupKindOverrides function
	override := GroupKindOverrides()
	override(testResource)
	
	// Verify the override was applied
	if testResource.ShortGroup != "compute" {
		t.Errorf("Expected ShortGroup to be 'compute', got '%s'", testResource.ShortGroup)
	}
	
	if testResource.Kind != "Instance" {
		t.Errorf("Expected Kind to be 'Instance', got '%s'", testResource.Kind)
	}
}

// TestAllResourcesHaveGroupMapping verifies each resource from ExternalNameConfigs
// has a valid group mapping and no resources are missed
func TestAllResourcesHaveGroupMapping(t *testing.T) {
	totalResourcesInExternalName := len(ExternalNameConfigs)
	totalResourcesInGroupMap := len(GroupMap)
	
	if totalResourcesInExternalName != totalResourcesInGroupMap {
		t.Errorf("Mismatch in resource count. ExternalNameConfigs has %d resources, GroupMap has %d resources", 
			totalResourcesInExternalName, totalResourcesInGroupMap)
	}
	
	// Verify every resource in ExternalNameConfigs is in GroupMap
	for resourceName := range ExternalNameConfigs {
		if _, exists := GroupMap[resourceName]; !exists {
			t.Errorf("Resource %s exists in ExternalNameConfigs but not in GroupMap", resourceName)
		}
	}
	
	// Verify every resource in GroupMap exists in ExternalNameConfigs
	for resourceName := range GroupMap {
		if _, exists := ExternalNameConfigs[resourceName]; !exists {
			t.Errorf("Resource %s exists in GroupMap but not in ExternalNameConfigs", resourceName)
		}
	}
}

// TestSpecificServiceMappings tests specific service mappings to ensure correctness
func TestSpecificServiceMappings(t *testing.T) {
	testCases := []struct {
		resourceName    string
		expectedGroup   string
		expectedKind    string
	}{
		{"oci_core_instance", "compute", "Instance"},
		{"oci_core_vcn", "networking", "Vcn"},
		{"oci_core_volume", "blockstorage", "Volume"},
		{"oci_core_drg", "networkconnectivity", "Drg"},
		{"oci_identity_compartment", "identity", "Compartment"},
		{"oci_containerengine_cluster", "containerengine", "Cluster"},
		{"oci_objectstorage_bucket", "objectstorage", "Bucket"},
		{"oci_kms_key", "kms", "Key"},
		{"oci_dns_zone", "dns", "Zone"},
		{"oci_load_balancer_load_balancer", "loadbalancer", "LoadBalancer"},
		{"oci_network_load_balancer_network_load_balancer", "networkloadbalancer", "NetworkLoadBalancer"},
	}
	
	for _, tc := range testCases {
		if calculator, exists := GroupMap[tc.resourceName]; exists {
			group, kind := calculator(tc.resourceName)
			if group != tc.expectedGroup {
				t.Errorf("Resource %s: expected group %s, got %s", tc.resourceName, tc.expectedGroup, group)
			}
			if kind != tc.expectedKind {
				t.Errorf("Resource %s: expected kind %s, got %s", tc.resourceName, tc.expectedKind, kind)
			}
		} else {
			t.Errorf("Resource %s not found in GroupMap", tc.resourceName)
		}
	}
}