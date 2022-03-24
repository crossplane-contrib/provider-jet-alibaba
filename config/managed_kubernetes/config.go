package managed_kubernetes

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_cs_managed_kubernetes", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "managedk8s"
	})
}
