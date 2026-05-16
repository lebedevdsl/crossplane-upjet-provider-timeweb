package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/database"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/dns"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/firewall"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/k8s"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/lb"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/project"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/s3"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/server"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/ssh"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/storage"
	"github.com/lebedevdsl/crossplane-upjet-provider-timeweb/config/vpc"
)

const (
	resourcePrefix = "twc"
	modulePath     = "github.com/lebedevdsl/crossplane-upjet-provider-timeweb"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("timeweb.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range resourceConfigurators() {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// resourceConfigurators returns the per-group Configure functions applied to
// every provider instance (cluster-scoped and namespaced) so the resource
// customizations stay in sync between the two.
func resourceConfigurators() []func(*ujconfig.Provider) {
	return []func(*ujconfig.Provider){
		database.Configure,
		dns.Configure,
		firewall.Configure,
		k8s.Configure,
		lb.Configure,
		project.Configure,
		s3.Configure,
		server.Configure,
		ssh.Configure,
		storage.Configure,
		vpc.Configure,
	}
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("timeweb.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range resourceConfigurators() {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
