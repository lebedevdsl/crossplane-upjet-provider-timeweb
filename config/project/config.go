// Package project contains resource configuration for Timeweb projects
// (logical groupings of cloud resources).
package project

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb project resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_project", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
}
