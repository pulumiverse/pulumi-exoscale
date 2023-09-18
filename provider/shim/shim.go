package shim

import (
	"context"
	"github.com/exoscale/terraform-provider-exoscale/exoscale"
	p "github.com/exoscale/terraform-provider-exoscale/pkg/provider"
	pftfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
)

func ShimmedProvider() shim.Provider {
	exo := p.New()
	return pftfbridge.MuxShimWithPF(
		context.Background(),
		shimv2.NewProvider(exoscale.Provider()),
		exo(),
	)
}
