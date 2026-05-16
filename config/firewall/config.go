// Package firewall contains resource configuration for firewalls and their rules.
package firewall

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb firewall resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_firewall", func(r *config.Resource) {
		r.ShortGroup = "firewall"
	})
	p.AddResourceConfigurator("twc_firewall_rule", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.References["firewall_id"] = config.Reference{TerraformName: "twc_firewall"}
	})
}
