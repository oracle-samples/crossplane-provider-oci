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

package servicemesh

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_service_mesh_access_policy", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "AccessPolicy"
		r.ShortGroup = "servicemesh"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["mesh_id"] = config.Reference{
			Type: "Mesh",
		}
		r.References["rules.destination.ingress_gateway_id"] = config.Reference{
			Type: "IngressGateway",
		}
		r.References["rules.destination.vitual_service_id"] = config.Reference{
			Type: "VirtualService",
		}
	})

	p.AddResourceConfigurator("oci_service_mesh_ingress_gateway", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "IngressGateway"
		r.ShortGroup = "servicemesh"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["mesh_id"] = config.Reference{
			Type: "Mesh",
		}
	})

	p.AddResourceConfigurator("oci_service_mesh_ingress_gateway_route_table", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "IngressGatewayRouteTable"
		r.ShortGroup = "servicemesh"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["ingress_gateway_id"] = config.Reference{
			Type: "IngressGateway",
		}
		r.References["route_rules.destinations.virtual_service_id"] = config.Reference{
			Type: "VirtualService",
		}
	})

	p.AddResourceConfigurator("oci_service_mesh_mesh", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "Mesh"
		r.ShortGroup = "servicemesh"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
	})

	p.AddResourceConfigurator("oci_service_mesh_virtual_deployment", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "VirtualDeployment"
		r.ShortGroup = "servicemesh"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["virtual_service_id"] = config.Reference{
			Type: "VirtualService",
		}
	})

	p.AddResourceConfigurator("oci_service_mesh_virtual_service", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "VirtualService"
		r.ShortGroup = "servicemesh"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["mesh_id"] = config.Reference{
			Type: "Mesh",
		}
	})

	p.AddResourceConfigurator("oci_service_mesh_virtual_service_route_table", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "VirtualServiceRouteTable"
		r.ShortGroup = "servicemesh"

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["route_rules.destinations.virtual_deployment_id"] = config.Reference{
			Type: "VirtualDeployment",
		}
		r.References["virtual_service_id"] = config.Reference{
			Type: "VirtualService",
		}
	})
}
