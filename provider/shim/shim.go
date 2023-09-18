package shim

import (
	p "github.com/exoscale/terraform-provider-exoscale/pkg/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProvider() provider.Provider {
	exo := p.New()
	return exo()
}
