// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// Fetch Exoscale [IAM](https://community.exoscale.com/documentation/iam/) Role.
//
// Corresponding resource: exoscale_iam_role.
func LookupIAMRole(ctx *pulumi.Context, args *LookupIAMRoleArgs, opts ...pulumi.InvokeOption) (*LookupIAMRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIAMRoleResult
	err := ctx.Invoke("exoscale:index/getIAMRole:getIAMRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIAMRole.
type LookupIAMRoleArgs struct {
	// The role ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// the role name to match (conflicts with `id`).
	Name     *string             `pulumi:"name"`
	Timeouts *GetIAMRoleTimeouts `pulumi:"timeouts"`
}

// A collection of values returned by getIAMRole.
type LookupIAMRoleResult struct {
	// A free-form text describing the IAM Role
	Description string `pulumi:"description"`
	// Defines if IAM Role Policy is editable or not.
	Editable bool `pulumi:"editable"`
	// The role ID to match (conflicts with `name`).
	Id string `pulumi:"id"`
	// IAM Role labels.
	Labels map[string]string `pulumi:"labels"`
	// the role name to match (conflicts with `id`).
	Name string `pulumi:"name"`
	// IAM Role permissions.
	Permissions []string `pulumi:"permissions"`
	// IAM Policy.
	Policy   GetIAMRolePolicy    `pulumi:"policy"`
	Timeouts *GetIAMRoleTimeouts `pulumi:"timeouts"`
}

func LookupIAMRoleOutput(ctx *pulumi.Context, args LookupIAMRoleOutputArgs, opts ...pulumi.InvokeOption) LookupIAMRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIAMRoleResult, error) {
			args := v.(LookupIAMRoleArgs)
			r, err := LookupIAMRole(ctx, &args, opts...)
			var s LookupIAMRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIAMRoleResultOutput)
}

// A collection of arguments for invoking getIAMRole.
type LookupIAMRoleOutputArgs struct {
	// The role ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// the role name to match (conflicts with `id`).
	Name     pulumi.StringPtrInput      `pulumi:"name"`
	Timeouts GetIAMRoleTimeoutsPtrInput `pulumi:"timeouts"`
}

func (LookupIAMRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIAMRoleArgs)(nil)).Elem()
}

// A collection of values returned by getIAMRole.
type LookupIAMRoleResultOutput struct{ *pulumi.OutputState }

func (LookupIAMRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIAMRoleResult)(nil)).Elem()
}

func (o LookupIAMRoleResultOutput) ToLookupIAMRoleResultOutput() LookupIAMRoleResultOutput {
	return o
}

func (o LookupIAMRoleResultOutput) ToLookupIAMRoleResultOutputWithContext(ctx context.Context) LookupIAMRoleResultOutput {
	return o
}

func (o LookupIAMRoleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIAMRoleResult] {
	return pulumix.Output[LookupIAMRoleResult]{
		OutputState: o.OutputState,
	}
}

// A free-form text describing the IAM Role
func (o LookupIAMRoleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIAMRoleResult) string { return v.Description }).(pulumi.StringOutput)
}

// Defines if IAM Role Policy is editable or not.
func (o LookupIAMRoleResultOutput) Editable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIAMRoleResult) bool { return v.Editable }).(pulumi.BoolOutput)
}

// The role ID to match (conflicts with `name`).
func (o LookupIAMRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIAMRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// IAM Role labels.
func (o LookupIAMRoleResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIAMRoleResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// the role name to match (conflicts with `id`).
func (o LookupIAMRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIAMRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

// IAM Role permissions.
func (o LookupIAMRoleResultOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIAMRoleResult) []string { return v.Permissions }).(pulumi.StringArrayOutput)
}

// IAM Policy.
func (o LookupIAMRoleResultOutput) Policy() GetIAMRolePolicyOutput {
	return o.ApplyT(func(v LookupIAMRoleResult) GetIAMRolePolicy { return v.Policy }).(GetIAMRolePolicyOutput)
}

func (o LookupIAMRoleResultOutput) Timeouts() GetIAMRoleTimeoutsPtrOutput {
	return o.ApplyT(func(v LookupIAMRoleResult) *GetIAMRoleTimeouts { return v.Timeouts }).(GetIAMRoleTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIAMRoleResultOutput{})
}
