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

package loadbalancer

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_load_balancer_load_balancer", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["subnet_ids"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet",
		}

	})
	p.AddResourceConfigurator("oci_load_balancer_hostname", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_backend_set", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_rule_set", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_backend", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
		r.References["backendset_name"] = config.Reference{
			Type: "BalancerBackendSet",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_certificate", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_load_balancer_routing_policy", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
		r.References["rules.actions.backend_set_name"] = config.Reference{
			Type: "BalancerBackendSet",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_path_route_set", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
		r.References["path_routes.backend_set_name"] = config.Reference{
			Type: "BalancerBackendSet",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_listener", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
		r.References["default_backend_set_name"] = config.Reference{
			Type: "BalancerBackendSet",
		}
		r.References["hostname_names"] = config.Reference{
			Type: "BalancerHostname",
		}
		r.References["path_route_set_name"] = config.Reference{
			Type: "BalancerPathRouteSet",
		}
		r.References["routing_policy_name"] = config.Reference{
			Type: "BalancerLoadBalancerRoutingPolicy",
		}
		r.References["rule_set_names"] = config.Reference{
			Type: "BalancerRuleSet",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_ssl_cipher_suite", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["load_balancer_id"] = config.Reference{
			Type: "BalancerLoadBalancer",
		}
	})
}
