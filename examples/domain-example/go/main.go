package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		demoDomain, err := exoscale.NewDomain(ctx, "demo-domain", &exoscale.DomainArgs{
			Name: pulumi.String("my-domain.tld"),
		})
		if err != nil {
			return err
		}
		ctx.Export("domainName", demoDomain.Name)
		return nil
	})
}
