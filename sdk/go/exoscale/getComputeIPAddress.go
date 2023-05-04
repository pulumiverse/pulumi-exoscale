// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// !> **WARNING:** This data source is **DEPRECATED** and will be removed in the next major version. Please use ElasticIP instead.
func GetComputeIPAddress(ctx *pulumi.Context, args *GetComputeIPAddressArgs, opts ...pulumi.InvokeOption) (*GetComputeIPAddressResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetComputeIPAddressResult
	err := ctx.Invoke("exoscale:index/getComputeIPAddress:getComputeIPAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeIPAddress.
type GetComputeIPAddressArgs struct {
	// The EIP description to match.
	Description *string `pulumi:"description"`
	// The Elastic IP (EIP) ID to match.
	Id *string `pulumi:"id"`
	// The EIP IPv4 address to match.
	IpAddress *string `pulumi:"ipAddress"`
	// The EIP tags to match.
	Tags map[string]string `pulumi:"tags"`
	// The Exoscale Zone name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getComputeIPAddress.
type GetComputeIPAddressResult struct {
	// The EIP description to match.
	Description *string `pulumi:"description"`
	// The Elastic IP (EIP) ID to match.
	Id *string `pulumi:"id"`
	// The EIP IPv4 address to match.
	IpAddress *string `pulumi:"ipAddress"`
	// The EIP tags to match.
	Tags map[string]string `pulumi:"tags"`
	// The Exoscale Zone name.
	Zone string `pulumi:"zone"`
}

func GetComputeIPAddressOutput(ctx *pulumi.Context, args GetComputeIPAddressOutputArgs, opts ...pulumi.InvokeOption) GetComputeIPAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetComputeIPAddressResult, error) {
			args := v.(GetComputeIPAddressArgs)
			r, err := GetComputeIPAddress(ctx, &args, opts...)
			var s GetComputeIPAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetComputeIPAddressResultOutput)
}

// A collection of arguments for invoking getComputeIPAddress.
type GetComputeIPAddressOutputArgs struct {
	// The EIP description to match.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The Elastic IP (EIP) ID to match.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The EIP IPv4 address to match.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// The EIP tags to match.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The Exoscale Zone name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (GetComputeIPAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeIPAddressArgs)(nil)).Elem()
}

// A collection of values returned by getComputeIPAddress.
type GetComputeIPAddressResultOutput struct{ *pulumi.OutputState }

func (GetComputeIPAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeIPAddressResult)(nil)).Elem()
}

func (o GetComputeIPAddressResultOutput) ToGetComputeIPAddressResultOutput() GetComputeIPAddressResultOutput {
	return o
}

func (o GetComputeIPAddressResultOutput) ToGetComputeIPAddressResultOutputWithContext(ctx context.Context) GetComputeIPAddressResultOutput {
	return o
}

// The EIP description to match.
func (o GetComputeIPAddressResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeIPAddressResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Elastic IP (EIP) ID to match.
func (o GetComputeIPAddressResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeIPAddressResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The EIP IPv4 address to match.
func (o GetComputeIPAddressResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeIPAddressResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// The EIP tags to match.
func (o GetComputeIPAddressResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetComputeIPAddressResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The Exoscale Zone name.
func (o GetComputeIPAddressResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetComputeIPAddressResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetComputeIPAddressResultOutput{})
}
