// Package lb contains resource configuration for Timeweb load balancers
// and their rules.
package lb

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb load balancer resources.
func Configure(p *config.Provider) {
	// Kind renamed from the upjet default (Lb → LoadBalancer) for readability.
	p.AddResourceConfigurator("twc_lb", func(r *config.Resource) {
		r.ShortGroup = "lb"
		r.Kind = "LoadBalancer"
		r.References["project_id"] = config.Reference{TerraformName: "twc_project"}
		r.References["floating_ip_id"] = config.Reference{TerraformName: "twc_floating_ip"}
	})
	p.AddResourceConfigurator("twc_lb_rule", func(r *config.Resource) {
		r.ShortGroup = "lb"
		r.Kind = "LoadBalancerRule"
		r.References["lb_id"] = config.Reference{TerraformName: "twc_lb"}
	})
}
