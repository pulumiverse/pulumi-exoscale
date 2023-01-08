// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTemplate(ctx *pulumi.Context, args *GetTemplateArgs, opts ...pulumi.InvokeOption) (*GetTemplateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTemplateResult
	err := ctx.Invoke("exoscale:index/getTemplate:getTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplate.
type GetTemplateArgs struct {
	// The compute instance template ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The template name to match (conflicts with `id`) (when multiple templates have the same name, the newest one will be returned).
	Name *string `pulumi:"name"`
	// A template category filter (default: `public`); among:
	Visibility *string `pulumi:"visibility"`
	// The Exoscale [Zone][zone] name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getTemplate.
type GetTemplateResult struct {
	// Username to use to log into a compute instance based on this template
	DefaultUser string  `pulumi:"defaultUser"`
	Id          *string `pulumi:"id"`
	Name        *string `pulumi:"name"`
	Visibility  *string `pulumi:"visibility"`
	Zone        string  `pulumi:"zone"`
}

func GetTemplateOutput(ctx *pulumi.Context, args GetTemplateOutputArgs, opts ...pulumi.InvokeOption) GetTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTemplateResult, error) {
			args := v.(GetTemplateArgs)
			r, err := GetTemplate(ctx, &args, opts...)
			var s GetTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTemplateResultOutput)
}

// A collection of arguments for invoking getTemplate.
type GetTemplateOutputArgs struct {
	// The compute instance template ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The template name to match (conflicts with `id`) (when multiple templates have the same name, the newest one will be returned).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A template category filter (default: `public`); among:
	Visibility pulumi.StringPtrInput `pulumi:"visibility"`
	// The Exoscale [Zone][zone] name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (GetTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getTemplate.
type GetTemplateResultOutput struct{ *pulumi.OutputState }

func (GetTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplateResult)(nil)).Elem()
}

func (o GetTemplateResultOutput) ToGetTemplateResultOutput() GetTemplateResultOutput {
	return o
}

func (o GetTemplateResultOutput) ToGetTemplateResultOutputWithContext(ctx context.Context) GetTemplateResultOutput {
	return o
}

// Username to use to log into a compute instance based on this template
func (o GetTemplateResultOutput) DefaultUser() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplateResult) string { return v.DefaultUser }).(pulumi.StringOutput)
}

func (o GetTemplateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GetTemplateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetTemplateResultOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplateResult) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

func (o GetTemplateResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplateResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTemplateResultOutput{})
}
