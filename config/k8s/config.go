// Package k8s contains resource configuration for managed Kubernetes
// clusters and their node groups.
package k8s

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb Kubernetes resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_k8s_cluster", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.References["network_id"] = config.Reference{TerraformName: "twc_vpc"}
		r.References["project_id"] = config.Reference{TerraformName: "twc_project"}
	})
	p.AddResourceConfigurator("twc_k8s_node_group", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.References["cluster_id"] = config.Reference{TerraformName: "twc_k8s_cluster"}
	})
}
