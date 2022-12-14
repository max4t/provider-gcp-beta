package compute

import (
	mycommon "github.com/max4t/provider-gcp-beta/config/common"
	"github.com/upbound/provider-gcp/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_compute_subnetwork", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")

		r.References["network"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/compute/v1beta1.Network",
			Extractor: mycommon.PathIDExtractor,
		}
	})

	p.AddResourceConfigurator("google_compute_region_target_tcp_proxy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
		r.References["backend_service"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/compute/v1beta1.RegionBackendService",
			Extractor: common.PathSelfLinkExtractor,
		}
	})

	p.AddResourceConfigurator("google_compute_forwarding_rule", func(r *config.Resource) {
		// Note(donovanmuller): See https://github.com/upbound/upjet/issues/95
		// BackendService is also a valid reference Type
		r.References["backend_service"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/compute/v1beta1.BackendService",
			Extractor: common.PathSelfLinkExtractor,
		}
		r.References["ip_address"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/compute/v1beta1.Address",
			Extractor: mycommon.PathIDExtractor,
		}
		r.References["network"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/compute/v1beta1.Network",
		}
		r.References["subnetwork"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork",
		}
		r.References["target"] = config.Reference{
			Type:      "RegionTargetTCPProxy",
			Extractor: mycommon.PathIDExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

}
