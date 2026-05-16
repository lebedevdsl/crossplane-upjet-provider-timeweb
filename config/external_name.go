package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider. Every Timeweb resource has its identifier assigned by the API
// at create time, so config.IdentifierFromProvider is the right default.
var ExternalNameConfigs = map[string]config.ExternalName{
	"twc_vpc":         config.IdentifierFromProvider,
	"twc_floating_ip": config.IdentifierFromProvider,

	"twc_firewall":      config.IdentifierFromProvider,
	"twc_firewall_rule": config.IdentifierFromProvider,

	"twc_k8s_cluster":    config.IdentifierFromProvider,
	"twc_k8s_node_group": config.IdentifierFromProvider,

	"twc_server":                      config.IdentifierFromProvider,
	"twc_server_ip":                   config.IdentifierFromProvider,
	"twc_server_disk":                 config.IdentifierFromProvider,
	"twc_server_disk_backup_schedule": config.IdentifierFromProvider,

	"twc_s3_bucket":           config.IdentifierFromProvider,
	"twc_s3_bucket_subdomain": config.IdentifierFromProvider,

	"twc_network_drive": config.IdentifierFromProvider,

	"twc_ssh_key": config.IdentifierFromProvider,

	"twc_dns_rr": config.IdentifierFromProvider,

	"twc_project": config.IdentifierFromProvider,

	"twc_database_cluster":         config.IdentifierFromProvider,
	"twc_database_instance":        config.IdentifierFromProvider,
	"twc_database_user":            config.IdentifierFromProvider,
	"twc_database_backup":          config.IdentifierFromProvider,
	"twc_database_backup_schedule": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
