// Package s3 contains resource configuration for S3-compatible buckets
// and their public subdomain bindings.
package s3

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure customizes Timeweb S3 resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("twc_s3_bucket", func(r *config.Resource) {
		r.ShortGroup = "s3"
		r.References["project_id"] = config.Reference{TerraformName: "twc_project"}
	})
	p.AddResourceConfigurator("twc_s3_bucket_subdomain", func(r *config.Resource) {
		r.ShortGroup = "s3"
		r.References["bucket_id"] = config.Reference{TerraformName: "twc_s3_bucket"}
	})
}
