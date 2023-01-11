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

package functions

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_functions_application", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["compartment_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment",
		}
		r.References["subnet_ids"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.Subnet",
		}
		r.References["network_security_group_ids"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.NetworkSecurityGroup",
		}
		r.References["kms_key_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-oci/apis/kms/v1alpha1.Key",
		}
		r.References["domain_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-oci/apis/dns/v1alpha1.Domain",
		}
	})

	p.AddResourceConfigurator("oci_functions_function", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["application_id"] = config.Reference{
			Type: "Application",
		}
	})

	p.AddResourceConfigurator("oci_functions_invoke_function", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["function_id"] = config.Reference{
			Type: "Function",
		}
	})
}
