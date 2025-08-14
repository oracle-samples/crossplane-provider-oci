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

package dns

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_dns_record", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_dns_resolver", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_dns_rrset", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_dns_resolver_endpoint", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["resolver_id"] = config.Reference{
			TerraformName: "oci_dns_resolver",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
		r.References["nsg_ids"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}
	})

	p.AddResourceConfigurator("oci_dns_steering_policy", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["health_check_monitor_id"] = config.Reference{
			TerraformName: "oci_health_checks_http_monitor",
		}
	})

	p.AddResourceConfigurator("oci_dns_steering_policy_attachment", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["steering_policy_id"] = config.Reference{
			TerraformName: "oci_dns_steering_policy",
		}
		r.References["zone_id"] = config.Reference{
			TerraformName: "oci_dns_zone",
		}
	})

	p.AddResourceConfigurator("oci_dns_tsig_key", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_dns_view", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_dns_zone", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["view_id"] = config.Reference{
			TerraformName: "oci_dns_view",
		}
		r.References["tsig_key_id"] = config.Reference{
			TerraformName: "oci_dns_tsig_key",
		}
	})
}
