/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-oci/config/artifacts"
	"github.com/crossplane-contrib/provider-jet-oci/config/certificatesmanagement"
	"github.com/crossplane-contrib/provider-jet-oci/config/containerengine"
	"github.com/crossplane-contrib/provider-jet-oci/config/core"
	"github.com/crossplane-contrib/provider-jet-oci/config/dns"
	"github.com/crossplane-contrib/provider-jet-oci/config/events"
	"github.com/crossplane-contrib/provider-jet-oci/config/filestorage"
	"github.com/crossplane-contrib/provider-jet-oci/config/functions"
	"github.com/crossplane-contrib/provider-jet-oci/config/healthchecks"
	"github.com/crossplane-contrib/provider-jet-oci/config/identity"
	"github.com/crossplane-contrib/provider-jet-oci/config/kms"
	"github.com/crossplane-contrib/provider-jet-oci/config/loadbalancer"
	"github.com/crossplane-contrib/provider-jet-oci/config/logging"
	"github.com/crossplane-contrib/provider-jet-oci/config/monitoring"
	"github.com/crossplane-contrib/provider-jet-oci/config/networkfirewall"
	"github.com/crossplane-contrib/provider-jet-oci/config/networkloadbalancer"
	"github.com/crossplane-contrib/provider-jet-oci/config/objectstorage"
	"github.com/crossplane-contrib/provider-jet-oci/config/ons"
	"github.com/crossplane-contrib/provider-jet-oci/config/servicemesh"
	"github.com/crossplane-contrib/provider-jet-oci/config/streaming"
	"github.com/crossplane-contrib/provider-jet-oci/config/vault"
)

const (
	resourcePrefix = "oci"
	modulePath     = "github.com/crossplane-contrib/provider-jet-oci"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList([]string{
			"oci_artifacts_container_repository$",
			"oci_artifacts_repository$",
			"oci_objectstorage_bucket$",
			"oci_objectstorage_object$",
			"oci_objectstorage_object_lifecycle_policy$",
			"oci_identity_compartment$",
			"oci_identity_group$",
			"oci_identity_policy$",
			"oci_identity_tag$",
			"oci_identity_tag_namespace$",
			"oci_identity_authentication_policy$",
			"oci_identity_identity_provider$",
			"oci_core_instance$",
			"oci_core_image$",
			"oci_core_vcn$",
			"oci_core_vnic_attachment$",
			"oci_core_subnet$",
			"oci_core_dedicated_vm_host$",
			"oci_core_dhcp_options$",
			"oci_core_vlan$",
			"oci_core_internet_gateway$",
			"oci_core_nat_gateway$",
			"oci_core_network_security_group$",
			"oci_core_network_security_group_security_rule$",
			"oci_core_route_table$",
			"oci_core_security_list$",
			"oci_core_drg$",
			"oci_core_drg_attachment$",
			"oci_core_drg_attachment_management$",
			"oci_core_drg_attachments_list$",
			"oci_core_drg_route_distribution$",
			"oci_core_drg_route_distribution_statement$",
			"oci_core_drg_route_table$",
			"oci_core_drg_route_table_route_rule$",
			"oci_core_private_ip$",
			"oci_core_public_ip$",
			"oci_core_public_ip_pool$",
			"oci_core_public_ip_pool_capacity$",
			"oci_core_remote_peering_connection$",
			"oci_core_volume$",
			"oci_core_volume_backup$",
			"oci_core_volume_attachment$",
			"oci_core_volume_backup_policy$",
			"oci_core_volume_backup_policy_assignment$",
			"oci_core_volume_group$",
			"oci_core_volume_group_backup$",
			"oci_core_capture_filter$",
			"oci_kms_key$",
			"oci_kms_key_version$",
			"oci_kms_vault$",
			"oci_containerengine_cluster$",
			"oci_containerengine_node_pool$",
			"oci_artifacts_container_configuration$",
			"oci_artifacts_generic_artifact$",
			"oci_ons_notification_topic$",
			"oci_ons_subscription$",
			"oci_network_load_balancer_network_load_balancer$",
			"oci_network_load_balancer_backend_set$",
			"oci_network_load_balancer_backend$",
			"oci_network_load_balancer_network_load_balancers_backend_sets_unified$",
			"oci_network_load_balancer_listener$",
			"oci_dns_record$",
			"oci_dns_resolver$",
			"oci_dns_rrset$",
			"oci_dns_resolver_endpoint$",
			"oci_dns_steering_policy$",
			"oci_dns_steering_policy_attachment$",
			"oci_dns_tsig_key$",
			"oci_dns_zone$",
			"oci_dns_view$",
			"oci_monitoring_alarm_history_collection$",
			"oci_monitoring_alarm$",
			"oci_health_checks_http_monitor$",
			"oci_health_checks_ping_monitor$",
			"oci_functions_application$",
			"oci_functions_function$",
			"oci_functions_invoke_function$",
			"oci_network_firewall_network_firewall_policy$",
			"oci_network_firewall_network_firewall$",
			"oci_logging_log_group$",
			"oci_logging_log$",
			"oci_logging_log_saved_search$",
			"oci_logging_unified_agent_configuration$",
			"oci_load_balancer_load_balancer$",
			"oci_load_balancer_hostname$",
			"oci_load_balancer_backend_set$",
			"oci_load_balancer_rule_set$",
			"oci_load_balancer_backend$",
			"oci_load_balancer_certificate$",
			"oci_load_balancer_load_balancer_routing_policy$",
			"oci_load_balancer_path_route_set$",
			"oci_load_balancer_listener$",
			"oci_load_balancer_ssl_cipher_suite$",
			"oci_core_instance_configuration$",
			"oci_core_cpe$",
			"oci_core_ipsec$",
			"oci_core_local_peering_gateway$",
			"oci_core_byoip_range$",
			"oci_service_mesh_access_policy$",
			"oci_service_mesh_ingress_gateway$",
			"oci_service_mesh_ingress_gateway_route_table$",
			"oci_service_mesh_mesh$",
			"oci_service_mesh_virtual_deployment$",
			"oci_service_mesh_virtual_service$",
			"oci_service_mesh_virtual_service_route_table$",
			"oci_certificates_management_certificate_authority$",
			"oci_file_storage_file_system$",
			"oci_file_storage_mount_target$",
			"oci_file_storage_export_set$",
			"oci_file_storage_export$",
			"oci_file_storage_replication$",
			"oci_file_storage_snapshot$",
			"oci_events_rule$",
			"oci_streaming_stream$",
			"oci_streaming_stream_pool$",
			"oci_streaming_connect_harness$",
			"oci_vault_secret$",
		}))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		objectstorage.Configure,
		identity.Configure,
		core.Configure,
		kms.Configure,
		containerengine.Configure,
		artifacts.Configure,
		ons.Configure,
		networkloadbalancer.Configure,
		dns.Configure,
		healthchecks.Configure,
		functions.Configure,
		networkfirewall.Configure,
		logging.Configure,
		monitoring.Configure,
		loadbalancer.Configure,
		servicemesh.Configure,
		certificatesmanagement.Configure,
		filestorage.Configure,
		events.Configure,
		vault.Configure,
		events.Configure,
		streaming.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
