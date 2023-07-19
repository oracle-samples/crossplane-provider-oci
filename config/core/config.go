/*
 * Copyright (c) 2022, 2023 Oracle and/or its affiliates
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

package core

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_core_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		// r.ShortGroup = "compute"

		// This resource reside inside a compartment. By defining it as a reference to Compartment
		// object, we can build cross resource referencing.
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
			// RefFieldName:      "compartmentIdRef",
			// SelectorFieldName: "compartmentIdSelector",
		}

		r.References["create_vnic_details.subnet_id"] = config.Reference{
			Type: "Subnet",
		}

		r.References["create_vnic_details.nsg_ids"] = config.Reference{
			Type: "NetworkSecurityGroup",
		}

		r.References["create_vnic_details.vlan_id"] = config.Reference{
			Type: "Vlan",
		}

		r.References["dedicated_vm_host_id"] = config.Reference{
			Type: "DedicatedVMHost",
		}

		r.References["source_details.source_id"] = config.Reference{
			Type: "Image",
		}

		// r.LateInitializer = config.LateInitializer{
		// 	// NOTE(): These are ignored because they conflict with each other.
		// 	// See the following for more details: https://github.com/crossplane/terrajet/issues/107
		// 	IgnoredFields: []string{
		// 		"compartment_id",
		//         "dedicated_vm_host_id"

		// 	},

	})

	p.AddResourceConfigurator("oci_core_image", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		// r.ShortGroup = "compute"

	})

	p.AddResourceConfigurator("oci_core_vcn", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		// r.ShortGroup = "compute"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
			// SelectorFieldName: "compartmentIdSelector",
		}

	})

	p.AddResourceConfigurator("oci_core_subnet", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
			// RefFieldName:      "vcnIdRef",
			// SelectorFieldName: "vcnIdSelector",
		}
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}

		r.References["security_list_ids"] = config.Reference{
			Type:              "SecurityList", // how to add array
			RefFieldName:      "SecurityListIDRefs",
			SelectorFieldName: "SecurityListIDSelector",
		}
		r.References["dhcp_options_id"] = config.Reference{
			Type: "DHCPOptions", // how to add array
		}

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

	})

	p.AddResourceConfigurator("oci_core_dedicated_vm_host", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

	})

	p.AddResourceConfigurator("oci_core_internet_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}

	})

	p.AddResourceConfigurator("oci_core_nat_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}

	})

	p.AddResourceConfigurator("oci_core_service_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_network_security_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}

		// TODO: these need more modification to work (likely custom Selectors/Extractors?)
		// r.References["source"] = config.Reference{
		// 	Type: "NetworkSecurityGroup",
		// }

		// r.References["destination"] = config.Reference{
		// 	Type: "NetworkSecurityGroup",
		// }

	})

	p.AddResourceConfigurator("oci_core_network_security_group_security_rule", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["network_security_group_id"] = config.Reference{
			Type: "NetworkSecurityGroup",
		}

	})

	p.AddResourceConfigurator("oci_core_route_table", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}

		r.References["route_rules.network_entity_id"] = config.Reference{
			Type: "NATGateway",
		}

		// r.References["route_rules.network_entity_id"] = config.Reference{
		// 	Type: "InternetGateway",
		// }

	})

	p.AddResourceConfigurator("oci_core_security_list", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}

	})

	p.AddResourceConfigurator("oci_core_dhcp_options", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_drg", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["drg_id"] = config.Reference{
			Type: "Drg",
		}
		r.References["drg_route_table_id"] = config.Reference{
			Type: "DrgRouteTable",
		}
		r.References["network_details.route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
		r.References["network_details.id"] = config.Reference{
			Type: "Vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_attachment_management", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["network_id"] = config.Reference{
			Type: "Vcn",
		}
		r.References["drg_id"] = config.Reference{
			Type: "Drg",
		}
		r.References["drg_route_table_id"] = config.Reference{
			Type: "DrgRouteTable",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_attachments_list", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["drg_id"] = config.Reference{
			Type: "Drg",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_route_distribution", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["drg_id"] = config.Reference{
			Type: "Drg",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_route_distribution_statement", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["drg_route_distribution_id"] = config.Reference{
			Type: "DrgRouteDistribution",
		}
		r.References["drg_attachment_id"] = config.Reference{
			Type: "DrgAttachment",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_route_table", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["drg_id"] = config.Reference{
			Type: "Drg",
		}
		r.References["import_drg_route_distribution_id"] = config.Reference{
			Type: "DrgRouteDistribution",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_route_table_route_rule", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["drg_route_table_id"] = config.Reference{
			Type: "DrgRouteTable",
		}
		r.References["next_hop_drg_attachment_id"] = config.Reference{
			Type: "DrgAttachment",
		}
	})

	p.AddResourceConfigurator("oci_core_private_ip", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["vlan_id"] = config.Reference{
			Type: "Vlan",
		}
		r.References["vnic_id"] = config.Reference{
			Type: "VnicAttachment",
		}
	})

	p.AddResourceConfigurator("oci_core_public_ip", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["private_ip_id"] = config.Reference{
			Type: "PrivateIP",
		}
		r.References["public_ip_pool_id"] = config.Reference{
			Type: "PublicIPPool",
		}
	})

	p.AddResourceConfigurator("oci_core_public_ip_pool", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_public_ip_pool_capacity", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["public_ip_pool_id"] = config.Reference{
			Type: "PublicIPPool",
		}
		// Data Source seems not supported at this moment for ByoipId
	})

	p.AddResourceConfigurator("oci_core_remote_peering_connection", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["drg_id"] = config.Reference{
			Type: "Drg",
		}
		r.References["peer_drg_id"] = config.Reference{
			Type: "RemotePeeringConnection",
		}
	})

	p.AddResourceConfigurator("oci_core_instance_configuration", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

	})

	p.AddResourceConfigurator("oci_core_cpe", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

	})

	p.AddResourceConfigurator("oci_core_capture_filter", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_ipsec", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["cpe_id"] = config.Reference{
			Type: "Cpe",
		}

		r.References["drg_id"] = config.Reference{
			Type: "Drg",
		}

	})

	p.AddResourceConfigurator("oci_core_local_peering_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}

	})

	p.AddResourceConfigurator("oci_core_byoip_range", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"
		r.Kind = "ByoipRange"
		r.References["byoip_range_id"] = config.Reference{
			Type: "ByoipRange",
		}
	})

	p.AddResourceConfigurator("oci_core_vlan", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["vcn_id"] = config.Reference{
			Type: "Vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_volume", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_backup", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["volume_id"] = config.Reference{
			Type: "Volume",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_backup_policy", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_backup_policy_assignment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["asset_id"] = config.Reference{
			Type: "Volume",
		}
		r.References["policy_id"] = config.Reference{
			Type: "VolumeBackupPolicy",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_group_backup", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["volume_group_id"] = config.Reference{
			Type: "VolumeGroup",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["source_details.volume_ids"] = config.Reference{
			Type: "Volume",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
		r.References["volume_id"] = config.Reference{
			Type: "Volume",
		}
	})
	p.AddResourceConfigurator("oci_core_vnic_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// we need to override the default group that terrajet generated for
		r.Version = "v1alpha1"
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
	})

}
