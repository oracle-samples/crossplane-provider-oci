/*
Copyright 2021 The Crossplane Authors.

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

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/terrajet/pkg/terraform"

	"github.com/crossplane-contrib/provider-jet-oci/apis/v1alpha1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal oci credentials as JSON"
	keyTenancy_ocid         = "tenancy_ocid"
	keyUser_ocid            = "user_ocid"
	keyPrivate_key          = "private_key"
	keyPrivate_key_path     = "private_key_path"
	keyPrivate_key_password = "private_key_password"
	keyFingerprint          = "fingerprint"
	keyRegion               = "region"
	keyConfig_file_profile  = "config_file_profile"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		ociCreds := map[string]string{}
		if err := json.Unmarshal(data, &ociCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set environment variables for sensitive provider configuration
		// Deprecated: In shared gRPC mode we do not support injecting
		// credentials via the environment variables. You should specify
		// credentials via the Terraform main.tf.json instead.
		/*ps.Env = []string{
			fmt.Sprintf("%s=%s", "HASHICUPS_USERNAME", ociCreds["username"]),
			fmt.Sprintf("%s=%s", "HASHICUPS_PASSWORD", ociCreds["password"]),
		}*/
		// set credentials in Terraform provider configuration
		ps.Configuration = map[string]interface{}{
			"tenancy_ocid":     ociCreds["tenancy_ocid"],
			"user_ocid":        ociCreds["user_ocid"],
			"private_key":      ociCreds["private_key"],
			"private_key_path": ociCreds["private_key_path"],
			"fingerprint":      ociCreds["fingerprint"],
			"region":           ociCreds["region"],
		}
		return ps, nil
	}
}
