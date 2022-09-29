package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cluster, err := exoscale.NewSKSCluster(ctx, "sks-cluster", &exoscale.SKSClusterArgs{
			Name: pulumi.String("my-sks-cluster"),
			Zone: pulumi.String("ch-gva-2"),
		})
		if err != nil {
			return err
		}
		ctx.Export("endpoint", cluster.Endpoint)
		return nil
	})
}
