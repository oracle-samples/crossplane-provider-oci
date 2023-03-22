/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/oracle/provider-oci/config/artifacts"
	"github.com/oracle/provider-oci/config/certificatesmanagement"
	"github.com/oracle/provider-oci/config/containerengine"
	"github.com/oracle/provider-oci/config/core"
	"github.com/oracle/provider-oci/config/dns"
	"github.com/oracle/provider-oci/config/events"
	"github.com/oracle/provider-oci/config/filestorage"
	"github.com/oracle/provider-oci/config/functions"
	"github.com/oracle/provider-oci/config/healthchecks"
	"github.com/oracle/provider-oci/config/identity"
	"github.com/oracle/provider-oci/config/kms"
	"github.com/oracle/provider-oci/config/loadbalancer"
	"github.com/oracle/provider-oci/config/logging"
	"github.com/oracle/provider-oci/config/monitoring"
	"github.com/oracle/provider-oci/config/networkfirewall"
	"github.com/oracle/provider-oci/config/networkloadbalancer"
	"github.com/oracle/provider-oci/config/objectstorage"
	"github.com/oracle/provider-oci/config/ons"
	"github.com/oracle/provider-oci/config/servicemesh"
	"github.com/oracle/provider-oci/config/streaming"
	"github.com/oracle/provider-oci/config/vault"
)

const (
	resourcePrefix = "oci"
	modulePath     = "github.com/oracle/provider-oci"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		objectstorage.Configure,
		identity.Configure,
		core.Configure,
		kms.Configure,
		containerengine.Configure,
		artifacts.Configure,
		ons.Configure,
		networkloadbalancer.Configure,
		dns.Configure,
		healthchecks.Configure,
		functions.Configure,
		networkfirewall.Configure,
		logging.Configure,
		monitoring.Configure,
		loadbalancer.Configure,
		servicemesh.Configure,
		certificatesmanagement.Configure,
		filestorage.Configure,
		events.Configure,
		vault.Configure,
		events.Configure,
		streaming.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
