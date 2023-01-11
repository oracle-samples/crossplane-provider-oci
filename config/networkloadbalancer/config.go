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

package networkloadbalancer

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_network_load_balancer_network_load_balancer", func(r *config.Resource) {
		// r.ShortGroup = "compartment"
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("oci_network_load_balancer_backend_set", func(r *config.Resource) {
		// r.ShortGroup = "compartment"
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["network_load_balancer_id"] = config.Reference{
			Type: "LoadBalancerNetworkLoadBalancer",
		}
	})

	p.AddResourceConfigurator("oci_network_load_balancer_backend", func(r *config.Resource) {
		// r.ShortGroup = "compartment"
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["network_load_balancer_id"] = config.Reference{
			Type: "LoadBalancerNetworkLoadBalancer",
		}
	})

	p.AddResourceConfigurator("oci_network_load_balancer_network_load_balancers_backend_sets_unified", func(r *config.Resource) {
		// r.ShortGroup = "compartment"
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["network_load_balancer_id"] = config.Reference{
			Type: "LoadBalancerNetworkLoadBalancer",
		}
	})

	p.AddResourceConfigurator("oci_network_load_balancer_listener", func(r *config.Resource) {
		// r.ShortGroup = "compartment"
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["network_load_balancer_id"] = config.Reference{
			Type: "LoadBalancerNetworkLoadBalancer",
		}
	})
}
