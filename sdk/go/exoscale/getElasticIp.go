// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// Fetch Exoscale [Elastic IPs (EIP)](https://community.exoscale.com/documentation/compute/eip/) data.
//
// Corresponding resource: exoscale_elastic_ip.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myElasticIp, err := exoscale.LookupElasticIp(ctx, &exoscale.LookupElasticIpArgs{
//				Zone:      "ch-gva-2",
//				IpAddress: pulumi.StringRef("1.2.3.4"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("myElasticIpId", myElasticIp.Id)
//			return nil
//		})
//	}
//
// ```
//
// Please refer to the examples
// directory for complete configuration examples.
func LookupElasticIp(ctx *pulumi.Context, args *LookupElasticIpArgs, opts ...pulumi.InvokeOption) (*LookupElasticIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupElasticIpResult
	err := ctx.Invoke("exoscale:index/getElasticIp:getElasticIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getElasticIp.
type LookupElasticIpArgs struct {
	// The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
	Id *string `pulumi:"id"`
	// The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
	IpAddress *string `pulumi:"ipAddress"`
	// The EIP labels to match (conflicts with `ipAddress` and `id`).
	Labels map[string]string `pulumi:"labels"`
	// The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getElasticIp.
type LookupElasticIpResult struct {
	// The Elastic IP (EIP) address family (`inet4` or `inet6`).
	AddressFamily string `pulumi:"addressFamily"`
	// The Elastic IP (EIP) CIDR.
	Cidr string `pulumi:"cidr"`
	// The Elastic IP (EIP) description.
	Description string `pulumi:"description"`
	// The *managed* EIP healthcheck configuration.
	Healthchecks []GetElasticIpHealthcheck `pulumi:"healthchecks"`
	// The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
	Id *string `pulumi:"id"`
	// The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
	IpAddress *string `pulumi:"ipAddress"`
	// The EIP labels to match (conflicts with `ipAddress` and `id`).
	Labels map[string]string `pulumi:"labels"`
	// Domain name for reverse DNS record.
	ReverseDns string `pulumi:"reverseDns"`
	// The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

func LookupElasticIpOutput(ctx *pulumi.Context, args LookupElasticIpOutputArgs, opts ...pulumi.InvokeOption) LookupElasticIpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupElasticIpResult, error) {
			args := v.(LookupElasticIpArgs)
			r, err := LookupElasticIp(ctx, &args, opts...)
			var s LookupElasticIpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupElasticIpResultOutput)
}

// A collection of arguments for invoking getElasticIp.
type LookupElasticIpOutputArgs struct {
	// The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// The EIP labels to match (conflicts with `ipAddress` and `id`).
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupElasticIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticIpArgs)(nil)).Elem()
}

// A collection of values returned by getElasticIp.
type LookupElasticIpResultOutput struct{ *pulumi.OutputState }

func (LookupElasticIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticIpResult)(nil)).Elem()
}

func (o LookupElasticIpResultOutput) ToLookupElasticIpResultOutput() LookupElasticIpResultOutput {
	return o
}

func (o LookupElasticIpResultOutput) ToLookupElasticIpResultOutputWithContext(ctx context.Context) LookupElasticIpResultOutput {
	return o
}

// The Elastic IP (EIP) address family (`inet4` or `inet6`).
func (o LookupElasticIpResultOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIpResult) string { return v.AddressFamily }).(pulumi.StringOutput)
}

// The Elastic IP (EIP) CIDR.
func (o LookupElasticIpResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIpResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// The Elastic IP (EIP) description.
func (o LookupElasticIpResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIpResult) string { return v.Description }).(pulumi.StringOutput)
}

// The *managed* EIP healthcheck configuration.
func (o LookupElasticIpResultOutput) Healthchecks() GetElasticIpHealthcheckArrayOutput {
	return o.ApplyT(func(v LookupElasticIpResult) []GetElasticIpHealthcheck { return v.Healthchecks }).(GetElasticIpHealthcheckArrayOutput)
}

// The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
func (o LookupElasticIpResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticIpResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
func (o LookupElasticIpResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticIpResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// The EIP labels to match (conflicts with `ipAddress` and `id`).
func (o LookupElasticIpResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupElasticIpResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Domain name for reverse DNS record.
func (o LookupElasticIpResultOutput) ReverseDns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIpResult) string { return v.ReverseDns }).(pulumi.StringOutput)
}

// The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
func (o LookupElasticIpResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIpResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupElasticIpResultOutput{})
}
