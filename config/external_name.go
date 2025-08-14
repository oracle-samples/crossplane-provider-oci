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
	"github.com/oracle/provider-oci/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"oci_objectstorage_bucket":                        config.IdentifierFromProvider,
	"oci_objectstorage_object":                        config.IdentifierFromProvider,
	"oci_objectstorage_object_lifecycle_policy":       config.IdentifierFromProvider,
	"oci_artifacts_container_repository":              config.IdentifierFromProvider,
	"oci_artifacts_repository":                        config.IdentifierFromProvider,
	"oci_identity_compartment":                        config.IdentifierFromProvider,
	"oci_identity_group":                              config.IdentifierFromProvider,
	"oci_identity_policy":                             config.IdentifierFromProvider,
	"oci_identity_tag":                                config.IdentifierFromProvider,
	"oci_identity_tag_namespace":                      config.IdentifierFromProvider,
	"oci_identity_authentication_policy":              config.IdentifierFromProvider,
	"oci_identity_identity_provider":                  config.IdentifierFromProvider,
	"oci_core_instance":                               config.IdentifierFromProvider,
	"oci_core_image":                                  config.IdentifierFromProvider,
	"oci_core_vcn":                                    config.IdentifierFromProvider,
	"oci_core_vnic_attachment":                        config.IdentifierFromProvider,
	"oci_core_subnet":                                 config.IdentifierFromProvider,
	"oci_core_dedicated_vm_host":                      config.IdentifierFromProvider,
	"oci_core_dhcp_options":                           config.IdentifierFromProvider,
	"oci_core_vlan":                                   config.IdentifierFromProvider,
	"oci_core_internet_gateway":                       config.IdentifierFromProvider,
	"oci_core_nat_gateway":                            config.IdentifierFromProvider,
	"oci_core_service_gateway":                        config.IdentifierFromProvider,
	"oci_core_network_security_group":                 config.IdentifierFromProvider,
	"oci_core_network_security_group_security_rule":   config.IdentifierFromProvider,
	"oci_core_route_table":                            config.IdentifierFromProvider,
	"oci_core_security_list":                          config.IdentifierFromProvider,
	"oci_core_drg":                                    config.IdentifierFromProvider,
	"oci_core_drg_attachment":                         config.IdentifierFromProvider,
	"oci_core_drg_attachment_management":              config.IdentifierFromProvider,
	"oci_core_drg_attachments_list":                   config.IdentifierFromProvider,
	"oci_core_drg_route_distribution":                 config.IdentifierFromProvider,
	"oci_core_drg_route_distribution_statement":       config.IdentifierFromProvider,
	"oci_core_drg_route_table":                        config.IdentifierFromProvider,
	"oci_core_drg_route_table_route_rule":             config.IdentifierFromProvider,
	"oci_core_private_ip":                             config.IdentifierFromProvider,
	"oci_core_public_ip":                              config.IdentifierFromProvider,
	"oci_core_public_ip_pool":                         config.IdentifierFromProvider,
	"oci_core_public_ip_pool_capacity":                config.IdentifierFromProvider,
	"oci_core_remote_peering_connection":              config.IdentifierFromProvider,
	"oci_core_volume":                                 config.IdentifierFromProvider,
	"oci_core_volume_backup":                          config.IdentifierFromProvider,
	"oci_core_volume_attachment":                      config.IdentifierFromProvider,
	"oci_core_volume_backup_policy":                   config.IdentifierFromProvider,
	"oci_core_volume_backup_policy_assignment":        config.IdentifierFromProvider,
	"oci_core_volume_group":                           config.IdentifierFromProvider,
	"oci_core_volume_group_backup":                    config.IdentifierFromProvider,
	"oci_core_capture_filter":                         config.IdentifierFromProvider,
	"oci_kms_key":                                     config.IdentifierFromProvider,
	"oci_kms_key_version":                             config.IdentifierFromProvider,
	"oci_kms_vault":                                   config.IdentifierFromProvider,
	"oci_containerengine_cluster":                     config.IdentifierFromProvider,
	"oci_containerengine_node_pool":                   config.IdentifierFromProvider,
	"oci_artifacts_container_configuration":           config.IdentifierFromProvider,
	"oci_artifacts_generic_artifact":                  config.IdentifierFromProvider,
	"oci_ons_notification_topic":                      config.IdentifierFromProvider,
	"oci_ons_subscription":                            config.IdentifierFromProvider,
	"oci_network_load_balancer_network_load_balancer": config.IdentifierFromProvider,
	"oci_network_load_balancer_backend_set":           config.IdentifierFromProvider,
	"oci_network_load_balancer_backend":               config.IdentifierFromProvider,
	"oci_network_load_balancer_network_load_balancers_backend_sets_unified": config.IdentifierFromProvider,
	"oci_network_load_balancer_listener":                                    config.IdentifierFromProvider,
	"oci_dns_record":                                                        config.IdentifierFromProvider,
	"oci_dns_resolver":                                                      config.IdentifierFromProvider,
	"oci_dns_rrset":                                                         config.IdentifierFromProvider,
	"oci_dns_resolver_endpoint":                                             config.IdentifierFromProvider,
	"oci_dns_steering_policy":                                               config.IdentifierFromProvider,
	"oci_dns_steering_policy_attachment":                                    config.IdentifierFromProvider,
	"oci_dns_tsig_key":                                                      config.IdentifierFromProvider,
	"oci_dns_zone":                                                          config.IdentifierFromProvider,
	"oci_dns_view":                                                          config.IdentifierFromProvider,
	"oci_monitoring_alarm_history_collection":                               config.IdentifierFromProvider,
	"oci_monitoring_alarm":                                                  config.IdentifierFromProvider,
	"oci_health_checks_http_monitor":                                        config.IdentifierFromProvider,
	"oci_health_checks_ping_monitor":                                        config.IdentifierFromProvider,
	"oci_functions_application":                                             config.IdentifierFromProvider,
	"oci_functions_function":                                                config.IdentifierFromProvider,
	"oci_functions_invoke_function":                                         config.IdentifierFromProvider,
	"oci_network_firewall_network_firewall_policy":                          config.IdentifierFromProvider,
	"oci_network_firewall_network_firewall":                                 config.IdentifierFromProvider,
	"oci_logging_log_group":                                                 config.IdentifierFromProvider,
	"oci_logging_log":                                                       config.IdentifierFromProvider,
	"oci_logging_log_saved_search":                                          config.IdentifierFromProvider,
	"oci_logging_unified_agent_configuration":                               config.IdentifierFromProvider,
	"oci_load_balancer_load_balancer":                                       config.IdentifierFromProvider,
	"oci_load_balancer_hostname":                                            config.IdentifierFromProvider,
	"oci_load_balancer_backend_set":                                         config.IdentifierFromProvider,
	"oci_load_balancer_rule_set":                                            config.IdentifierFromProvider,
	"oci_load_balancer_backend":                                             config.IdentifierFromProvider,
	"oci_load_balancer_certificate":                                         config.IdentifierFromProvider,
	"oci_load_balancer_load_balancer_routing_policy":                        config.IdentifierFromProvider,
	"oci_load_balancer_path_route_set":                                      config.IdentifierFromProvider,
	"oci_load_balancer_listener":                                            config.IdentifierFromProvider,
	"oci_load_balancer_ssl_cipher_suite":                                    config.IdentifierFromProvider,
	"oci_core_instance_configuration":                                       config.IdentifierFromProvider,
	"oci_core_cpe":                                                          config.IdentifierFromProvider,
	"oci_core_ipsec":                                                        config.IdentifierFromProvider,
	"oci_core_local_peering_gateway":                                        config.IdentifierFromProvider,
	"oci_core_byoip_range":                                                  config.IdentifierFromProvider,
	"oci_certificates_management_certificate_authority":                     config.IdentifierFromProvider,
	"oci_file_storage_file_system":                                          config.IdentifierFromProvider,
	"oci_file_storage_mount_target":                                         config.IdentifierFromProvider,
	"oci_file_storage_export_set":                                           config.IdentifierFromProvider,
	"oci_file_storage_export":                                               config.IdentifierFromProvider,
	"oci_file_storage_replication":                                          config.IdentifierFromProvider,
	"oci_file_storage_snapshot":                                             config.IdentifierFromProvider,
	"oci_events_rule":                                                       config.IdentifierFromProvider,
	"oci_streaming_stream":                                                  config.IdentifierFromProvider,
	"oci_streaming_stream_pool":                                             config.IdentifierFromProvider,
	"oci_streaming_connect_harness":                                         config.IdentifierFromProvider,
	"oci_vault_secret":                                                      config.IdentifierFromProvider,
	"oci_core_app_catalog_listing_resource_version_agreement":               config.IdentifierFromProvider,
	"oci_core_app_catalog_subscription":                                     config.IdentifierFromProvider,
	"oci_core_boot_volume":                                                  config.IdentifierFromProvider,
	"oci_core_boot_volume_backup":                                           config.IdentifierFromProvider,
	"oci_core_cluster_network":                                              config.IdentifierFromProvider,
	"oci_core_compute_capacity_reservation":                                 config.IdentifierFromProvider,
	"oci_core_compute_cluster":                                              config.IdentifierFromProvider,
	"oci_core_compute_image_capability_schema":                              config.IdentifierFromProvider,
	"oci_core_console_history":                                              config.IdentifierFromProvider,
	"oci_core_cross_connect":                                                config.IdentifierFromProvider,
	"oci_core_cross_connect_group":                                          config.IdentifierFromProvider,
	"oci_core_instance_console_connection":                                  config.IdentifierFromProvider,
	"oci_core_instance_pool":                                                config.IdentifierFromProvider,
	"oci_core_instance_pool_instance":                                       config.IdentifierFromProvider,
	"oci_core_ipsec_connection_tunnel_management":                           config.IdentifierFromProvider,
	"oci_core_ipv6":                                                         config.IdentifierFromProvider,
	"oci_core_route_table_attachment":                                       config.IdentifierFromProvider,
	"oci_core_shape_management":                                             config.IdentifierFromProvider,
	"oci_core_virtual_circuit":                                              config.IdentifierFromProvider,
	"oci_core_vtap":                                                         config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
			r.Version = common.VersionAlpha1
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
