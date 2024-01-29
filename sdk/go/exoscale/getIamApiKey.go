// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// Fetch Exoscale [IAM](https://community.exoscale.com/documentation/iam/) API Key.
//
// Corresponding resource: exoscale_iam_role.
func LookupIamApiKey(ctx *pulumi.Context, args *LookupIamApiKeyArgs, opts ...pulumi.InvokeOption) (*LookupIamApiKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIamApiKeyResult
	err := ctx.Invoke("exoscale:index/getIamApiKey:getIamApiKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamApiKey.
type LookupIamApiKeyArgs struct {
	// The IAM API Key to match.
	Key      string                `pulumi:"key"`
	Timeouts *GetIamApiKeyTimeouts `pulumi:"timeouts"`
}

// A collection of values returned by getIamApiKey.
type LookupIamApiKeyResult struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// The IAM API Key to match.
	Key string `pulumi:"key"`
	// IAM API Key name.
	Name string `pulumi:"name"`
	// IAM API Key role ID.
	RoleId   string                `pulumi:"roleId"`
	Timeouts *GetIamApiKeyTimeouts `pulumi:"timeouts"`
}

func LookupIamApiKeyOutput(ctx *pulumi.Context, args LookupIamApiKeyOutputArgs, opts ...pulumi.InvokeOption) LookupIamApiKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIamApiKeyResult, error) {
			args := v.(LookupIamApiKeyArgs)
			r, err := LookupIamApiKey(ctx, &args, opts...)
			var s LookupIamApiKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIamApiKeyResultOutput)
}

// A collection of arguments for invoking getIamApiKey.
type LookupIamApiKeyOutputArgs struct {
	// The IAM API Key to match.
	Key      pulumi.StringInput           `pulumi:"key"`
	Timeouts GetIamApiKeyTimeoutsPtrInput `pulumi:"timeouts"`
}

func (LookupIamApiKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamApiKeyArgs)(nil)).Elem()
}

// A collection of values returned by getIamApiKey.
type LookupIamApiKeyResultOutput struct{ *pulumi.OutputState }

func (LookupIamApiKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamApiKeyResult)(nil)).Elem()
}

func (o LookupIamApiKeyResultOutput) ToLookupIamApiKeyResultOutput() LookupIamApiKeyResultOutput {
	return o
}

func (o LookupIamApiKeyResultOutput) ToLookupIamApiKeyResultOutputWithContext(ctx context.Context) LookupIamApiKeyResultOutput {
	return o
}

// The ID of this resource.
func (o LookupIamApiKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamApiKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IAM API Key to match.
func (o LookupIamApiKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamApiKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

// IAM API Key name.
func (o LookupIamApiKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamApiKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

// IAM API Key role ID.
func (o LookupIamApiKeyResultOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamApiKeyResult) string { return v.RoleId }).(pulumi.StringOutput)
}

func (o LookupIamApiKeyResultOutput) Timeouts() GetIamApiKeyTimeoutsPtrOutput {
	return o.ApplyT(func(v LookupIamApiKeyResult) *GetIamApiKeyTimeouts { return v.Timeouts }).(GetIamApiKeyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIamApiKeyResultOutput{})
}
