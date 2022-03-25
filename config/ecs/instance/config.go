package ecs

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_instance", func(r *config.Resource) {
		r.ShortGroup = "ecs"
	})
}
