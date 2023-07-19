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

package events

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_events_rule", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = config.IdentifierFromProvider

		r.References["compartment_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment",
		}

		r.References["actions.actions.function_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/functions/v1alpha1.Function",
		}

		r.References["actions.actions.stream_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/streaming/v1alpha1.Stream",
		}

		r.References["actions.actions.topic_id"] = config.Reference{
			Type: "github.com/oracle/provider-oci/apis/ons/v1alpha1.NotificationTopic",
		}
	})
}
