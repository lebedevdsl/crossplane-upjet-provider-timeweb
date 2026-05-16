// Package database contains resource configuration for managed database
// clusters (the modern /api/v1/databases family: cluster + instance + user
// + backup + backup schedule). The legacy /api/v1/dbs single-engine family
// is intentionally not wrapped.
package database

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb managed database resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_database_cluster", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.References["project_id"] = config.Reference{TerraformName: "twc_project"}
	})
	// Every member of a cluster (instances, users, backups, backup schedules)
	// links back to its parent via cluster_id.
	for _, name := range []string{
		"twc_database_instance",
		"twc_database_user",
		"twc_database_backup",
		"twc_database_backup_schedule",
	} {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			r.ShortGroup = "database"
			r.References["cluster_id"] = config.Reference{TerraformName: "twc_database_cluster"}
		})
	}
}
