package compute

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_compute_region_target_tcp_proxy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
}
