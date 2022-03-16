package vpc

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_vpc", func(r *config.Resource) {
		r.ShortGroup = "vpc"
	})
}
