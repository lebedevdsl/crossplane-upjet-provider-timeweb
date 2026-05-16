// Package vpc contains resource configuration for the VPC group:
// VPCs, floating IPs, and virtual routers.
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
	// Router glues VPCs and floating IPs together. The nested
	// `networks[].network_name`, `ips[].ip`, and `ips[].nat[].network_name`
	// paths reference other VPC-group MRs. `ips[].ip` extracts the FIP's
	// assigned IPv4 from its status; the others use the default extractor
	// (external-name == Timeweb id).
	//
	// `parent_services[].id` would naturally reference twc_k8s_cluster,
	// but twc_k8s_cluster already references twc_vpc via `network_id` —
	// adding the reverse-direction reference here would create an
	// `apis/vpc ↔ apis/k8s` Go import cycle (vpc resolvers would import
	// k8s types, and vice-versa). Keep the cluster→vpc direction (more
	// useful day-to-day) and fill `parent_services[0].id` manually after
	// the K8s cluster is Ready.
	p.AddResourceConfigurator("twc_router", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.References["project_id"] = config.Reference{TerraformName: "twc_project"}
		r.References["networks.network_name"] = config.Reference{
			TerraformName: "twc_vpc",
		}
		r.References["ips.ip"] = config.Reference{
			TerraformName: "twc_floating_ip",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("ip", true)`,
		}
		r.References["ips.nat.network_name"] = config.Reference{
			TerraformName: "twc_vpc",
		}
	})
}
