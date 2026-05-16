// Package vpc contains resource configuration for the VPC group:
// VPCs and floating IPs.
package vpc

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb VPC group resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_vpc", func(r *config.Resource) {
		r.ShortGroup = "vpc"
	})
	// Kind renamed from the upjet default (Ip → FloatingIP) for readability
	// and to avoid clashing with twc_server_ip's Kind.
	p.AddResourceConfigurator("twc_floating_ip", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "FloatingIP"
	})
}
