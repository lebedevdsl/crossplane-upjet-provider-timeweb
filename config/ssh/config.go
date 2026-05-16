// Package ssh contains resource configuration for SSH keys.
package ssh

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb SSH key resources. The Kind is renamed from
// the upjet default (SshKey) to follow Go/Kubernetes convention for
// initialisms.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_ssh_key", func(r *config.Resource) {
		r.ShortGroup = "ssh"
		r.Kind = "SSHKey"
	})
}
