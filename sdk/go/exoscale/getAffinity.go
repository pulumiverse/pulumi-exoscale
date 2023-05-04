// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// !> **WARNING:** This data source is DEPRECATED and will be removed in the next major version. Please use AntiAffinityGroup instead.
func LookupAffinity(ctx *pulumi.Context, args *LookupAffinityArgs, opts ...pulumi.InvokeOption) (*LookupAffinityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAffinityResult
	err := ctx.Invoke("exoscale:index/getAffinity:getAffinity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAffinity.
type LookupAffinityArgs struct {
	// The anti-affinity group ID to match (conflicts with `name`)
	Id *string `pulumi:"id"`
	// The group name to match (conflicts with `id`)
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAffinity.
type LookupAffinityResult struct {
	// The anti-affinity group ID to match (conflicts with `name`)
	Id *string `pulumi:"id"`
	// The group name to match (conflicts with `id`)
	Name *string `pulumi:"name"`
}

func LookupAffinityOutput(ctx *pulumi.Context, args LookupAffinityOutputArgs, opts ...pulumi.InvokeOption) LookupAffinityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAffinityResult, error) {
			args := v.(LookupAffinityArgs)
			r, err := LookupAffinity(ctx, &args, opts...)
			var s LookupAffinityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAffinityResultOutput)
}

// A collection of arguments for invoking getAffinity.
type LookupAffinityOutputArgs struct {
	// The anti-affinity group ID to match (conflicts with `name`)
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The group name to match (conflicts with `id`)
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAffinityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAffinityArgs)(nil)).Elem()
}

// A collection of values returned by getAffinity.
type LookupAffinityResultOutput struct{ *pulumi.OutputState }

func (LookupAffinityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAffinityResult)(nil)).Elem()
}

func (o LookupAffinityResultOutput) ToLookupAffinityResultOutput() LookupAffinityResultOutput {
	return o
}

func (o LookupAffinityResultOutput) ToLookupAffinityResultOutputWithContext(ctx context.Context) LookupAffinityResultOutput {
	return o
}

// The anti-affinity group ID to match (conflicts with `name`)
func (o LookupAffinityResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAffinityResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The group name to match (conflicts with `id`)
func (o LookupAffinityResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAffinityResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAffinityResultOutput{})
}
