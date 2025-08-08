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

package objectstorage

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_objectstorage_bucket", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "oci_kms_key",
		}
	})

	p.AddResourceConfigurator("oci_objectstorage_object", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
		r.References["opc_sse_kms_key_id"] = config.Reference{
			TerraformName: "oci_kms_key",
		}
	})

	p.AddResourceConfigurator("oci_objectstorage_object_lifecycle_policy", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider
	})
}
