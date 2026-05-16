// Package dns contains resource configuration for DNS resource records.
package dns

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb DNS resources. The Kind is renamed from the
// upjet default (DnsRr) for readability.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_dns_rr", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.Kind = "DnsRecord"
	})
}
