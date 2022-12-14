// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Exoscale CloudStack API endpoint (by default: https://api.exoscale.com/v1)
func GetComputeEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:computeEndpoint")
}

// CloudStack ini configuration filename (by default: cloudstack.ini)
func GetConfig(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:config")
}

// Deprecated: Does nothing
func GetDelay(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "exoscale:delay")
}

// Exoscale DNS API endpoint (by default: https://api.exoscale.com/dns)
func GetDnsEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:dnsEndpoint")
}
func GetEnvironment(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:environment")
}

// Defines if the user-data of compute instances should be gzipped (by default: true)
func GetGzipUserData(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "exoscale:gzipUserData")
}

// Exoscale API key
func GetKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:key")
}

// Deprecated: Use region instead
func GetProfile(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:profile")
}

// CloudStack ini configuration section name (by default: cloudstack)
func GetRegion(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:region")
}

// Exoscale API secret
func GetSecret(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:secret")
}

// Timeout in seconds for waiting on compute resources to become available (by default: 300)
func GetTimeout(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "exoscale:timeout")
}

// Deprecated: Use key instead
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "exoscale:token")
}
