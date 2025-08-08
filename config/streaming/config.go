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

package streaming

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_streaming_stream", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["stream_pool_id"] = config.Reference{
			TerraformName: "oci_streaming_stream_pool",
		}
	})

	p.AddResourceConfigurator("oci_streaming_stream_pool", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["custom_encryption_key.kms_key_id"] = config.Reference{
			TerraformName: "oci_kms_key",
		}
		r.References["private_endpoint_settings.nsg_ids"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}
		r.References["private_endpoint_settings.subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
		r.References["private_endpoint_settings.private_endpoint_ip"] = config.Reference{
			TerraformName: "oci_core_private_ip",
		}
	})

	p.AddResourceConfigurator("oci_streaming_connect_harness", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["stream_pool_id"] = config.Reference{
			TerraformName: "oci_streaming_stream_pool",
		}
	})
}
