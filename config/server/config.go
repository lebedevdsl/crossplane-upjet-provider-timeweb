// Package server contains resource configuration for cloud servers and
// their attached IPs, disks, and disk backup schedules.
package server

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb cloud server resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_server", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["project_id"] = config.Reference{TerraformName: "twc_project"}
		r.References["floating_ip_id"] = config.Reference{TerraformName: "twc_floating_ip"}
		// Used when cloning a server from an existing one.
		r.References["source_server_id"] = config.Reference{TerraformName: "twc_server"}
		// List-of-IDs reference; upjet generates plural Refs/Selector fields.
		r.References["ssh_keys_ids"] = config.Reference{TerraformName: "twc_ssh_key"}
	})
	p.AddResourceConfigurator("twc_server_disk", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["source_server_id"] = config.Reference{TerraformName: "twc_server"}
	})
	p.AddResourceConfigurator("twc_server_disk_backup_schedule", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["source_server_id"] = config.Reference{TerraformName: "twc_server"}
		r.References["source_server_disk_id"] = config.Reference{TerraformName: "twc_server_disk"}
	})
	// Kind renamed from the upjet default (Ip → ServerIP) for readability
	// and to avoid clashing with twc_floating_ip's Kind.
	p.AddResourceConfigurator("twc_server_ip", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.Kind = "ServerIP"
		r.References["source_server_id"] = config.Reference{TerraformName: "twc_server"}
	})
}
