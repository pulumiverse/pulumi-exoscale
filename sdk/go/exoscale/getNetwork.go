// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// !> **WARNING:** This data source is **DEPRECATED** and will be removed in the next major version. Please use PrivateNetwork instead.
func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNetworkResult
	err := ctx.Invoke("exoscale:index/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	// The private network ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The network name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The Exoscale Zone name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// The private network description.
	Description string  `pulumi:"description"`
	EndIp       string  `pulumi:"endIp"`
	Id          *string `pulumi:"id"`
	Name        *string `pulumi:"name"`
	// The network mask defining the IPv4 network allowed for static leases.
	// * `startIp`/`endIp` - The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	Netmask string `pulumi:"netmask"`
	StartIp string `pulumi:"startIp"`
	Zone    string `pulumi:"zone"`
}

func LookupNetworkOutput(ctx *pulumi.Context, args LookupNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkResult, error) {
			args := v.(LookupNetworkArgs)
			r, err := LookupNetwork(ctx, &args, opts...)
			var s LookupNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkOutputArgs struct {
	// The private network ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The network name to match (conflicts with `id`).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The Exoscale Zone name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type LookupNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkResult)(nil)).Elem()
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutput() LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutputWithContext(ctx context.Context) LookupNetworkResultOutput {
	return o
}

// The private network description.
func (o LookupNetworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) EndIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.EndIp }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The network mask defining the IPv4 network allowed for static leases.
// * `startIp`/`endIp` - The first/last IPv4 addresses used by the DHCP service for dynamic leases.
func (o LookupNetworkResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Netmask }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) StartIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.StartIp }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkResultOutput{})
}
