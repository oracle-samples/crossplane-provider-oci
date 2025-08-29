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

package loadbalancer

import (
	"github.com/oracle/provider-oci/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

const (
	GroupName string = "loadbalancer"

	KindLoadBalancer   string = "LoadBalancer"
	KindBackendSet     string = "BackendSet"
	KindLBHostname     string = "LBHostname"
	KindRuleSet        string = "RuleSet"
	KindBackend        string = "Backend"
	KindCertificate    string = "Certificate"
	KindRoutingPolicy  string = "RoutingPolicy"
	KindPathRouteSet   string = "PathRouteSet"
	KindListener       string = "Listener"
	KindSSLCipherSuite string = "SSLCipherSuite"

	ParamLoadBalancerID string = "load_balancer_id"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_load_balancer_load_balancer", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindLoadBalancer
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["subnet_ids"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}

	})
	p.AddResourceConfigurator("oci_load_balancer_hostname", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindLBHostname
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = GetLBHostnameIDFunc
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_backend_set", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindBackendSet
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = GetBackendSetIDFunc
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_rule_set", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindRuleSet
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = GetRuleSetIDFunc
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_backend", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindBackend
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = GetBackendIDFunc
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
		r.References["backendset_name"] = config.Reference{
			TerraformName: "oci_load_balancer_backend_set",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_certificate", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindCertificate
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = GetCertificateIDFunc
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_load_balancer_routing_policy", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindRoutingPolicy
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = GetRoutingPolicyIDFunc
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
		r.References["rules.actions.backend_set_name"] = config.Reference{
			TerraformName: "oci_load_balancer_backend_set",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_path_route_set", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindPathRouteSet
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = GetPathRouteSetIDFunc
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
		r.References["path_routes.backend_set_name"] = config.Reference{
			TerraformName: "oci_load_balancer_backend_set",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_listener", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindListener
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = GetPathRouteSetIDFunc
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
		r.References["default_backend_set_name"] = config.Reference{
			TerraformName: "oci_load_balancer_backend_set",
		}
		r.References["hostname_names"] = config.Reference{
			TerraformName: "oci_load_balancer_hostname",
		}
		r.References["path_route_set_name"] = config.Reference{
			TerraformName: "oci_load_balancer_path_route_set",
		}
		r.References["routing_policy_name"] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer_routing_policy",
		}
		r.References["rule_set_names"] = config.Reference{
			TerraformName: "oci_load_balancer_rule_set",
		}
	})
	p.AddResourceConfigurator("oci_load_balancer_ssl_cipher_suite", func(r *config.Resource) {
		r.ShortGroup = GroupName
		r.Kind = KindSSLCipherSuite
		r.References[ParamLoadBalancerID] = config.Reference{
			TerraformName: "oci_load_balancer_load_balancer",
		}
	})
}
