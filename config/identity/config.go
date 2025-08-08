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

package identity

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_identity_compartment", func(r *config.Resource) {
		// r.ShortGroup = "compartment"
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

	})

	p.AddResourceConfigurator("oci_identity_group", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("oci_identity_authentication_policy", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_identity_tag_namespace", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})
	p.AddResourceConfigurator("oci_identity_policy", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_identity_tag", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["tag_namespace_id"] = config.Reference{
			TerraformName: "oci_identity_tag_namespace",
		}
	})

	p.AddResourceConfigurator("oci_identity_identity_provider", func(r *config.Resource) {
		r.Version = "v1alpha1"
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})
}
