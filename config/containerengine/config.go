/*
Copyright (c) 2022, Oracle and/or its affiliates

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
package containerengine

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_containerengine_cluster", func(r *config.Resource) {
		// r.ShortGroup = "compartment"
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["vcn_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/core/v1alpha1.Vcn",
		}

		r.References["endpoint_config.nsg_ids"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/core/v1alpha1.NetworkSecurityGroup",
		}

		r.References["endpoint_config.subnet_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet",
		}

		r.References["options.service_lb_subnet_ids"] = config.Reference{
			Type:              "github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet",
			RefFieldName:      "ServiceLBSubnetIdsRef",
			SelectorFieldName: "ServiceLBSubnetIDSelector",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("oci_containerengine_node_pool", func(r *config.Resource) {
		// r.ShortGroup = "compartment"
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}

		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}

		r.References["node_config_details.nsg_ids"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/core/v1alpha1.NetworkSecurityGroup",
		}

		r.References["node_config_details.placement_configs.subnet_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet",
		}

		r.LateInitializer = config.LateInitializer{
			// These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/crossplane/terrajet/issues/107
			IgnoredFields: []string{
				"subnet_ids",
				"node_image_name",
				"node_image_id",
			},
		}

		r.UseAsync = true
	})

}
