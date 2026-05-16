// Package storage contains resource configuration for network-attached
// block storage volumes.
package storage

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb network drive resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_network_drive", func(r *config.Resource) {
		r.ShortGroup = "storage"
	})
}
