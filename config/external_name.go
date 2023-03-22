/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

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
	"oci_core_volume":                                  config.IdentifierFromProvider,
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
	"oci_service_mesh_access_policy":                                        config.IdentifierFromProvider,
	"oci_service_mesh_ingress_gateway":                                      config.IdentifierFromProvider,
	"oci_service_mesh_ingress_gateway_route_table":                          config.IdentifierFromProvider,
	"oci_service_mesh_mesh":                                                 config.IdentifierFromProvider,
	"oci_service_mesh_virtual_deployment":                                   config.IdentifierFromProvider,
	"oci_service_mesh_virtual_service":                                      config.IdentifierFromProvider,
	"oci_service_mesh_virtual_service_route_table":                          config.IdentifierFromProvider,
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
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
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
