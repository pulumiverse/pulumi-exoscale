// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

func GetComputeInstanceList(ctx *pulumi.Context, args *GetComputeInstanceListArgs, opts ...pulumi.InvokeOption) (*GetComputeInstanceListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetComputeInstanceListResult
	err := ctx.Invoke("exoscale:index/getComputeInstanceList:getComputeInstanceList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeInstanceList.
type GetComputeInstanceListArgs struct {
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	CreatedAt *string `pulumi:"createdAt"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// Match against this int
	DiskSize *int `pulumi:"diskSize"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Id *string `pulumi:"id"`
	// Match against this bool
	Ipv6 *bool `pulumi:"ipv6"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Ipv6Address *string `pulumi:"ipv6Address"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Labels map[string]string `pulumi:"labels"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ManagerId *string `pulumi:"managerId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ManagerType *string `pulumi:"managerType"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Name *string `pulumi:"name"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	PublicIpAddress *string `pulumi:"publicIpAddress"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ReverseDns *string `pulumi:"reverseDns"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	SshKey *string `pulumi:"sshKey"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	State *string `pulumi:"state"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	TemplateId *string `pulumi:"templateId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Type *string `pulumi:"type"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	UserData *string `pulumi:"userData"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getComputeInstanceList.
type GetComputeInstanceListResult struct {
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	CreatedAt *string `pulumi:"createdAt"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// Match against this int
	DiskSize *int `pulumi:"diskSize"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Id *string `pulumi:"id"`
	// The list of exoscale*compute*instance.
	Instances []GetComputeInstanceListInstance `pulumi:"instances"`
	// Match against this bool
	Ipv6 *bool `pulumi:"ipv6"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Ipv6Address *string `pulumi:"ipv6Address"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Labels map[string]string `pulumi:"labels"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ManagerId *string `pulumi:"managerId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ManagerType *string `pulumi:"managerType"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Name *string `pulumi:"name"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	PublicIpAddress *string `pulumi:"publicIpAddress"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ReverseDns *string `pulumi:"reverseDns"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	SshKey *string `pulumi:"sshKey"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	State *string `pulumi:"state"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	TemplateId *string `pulumi:"templateId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Type *string `pulumi:"type"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	UserData *string `pulumi:"userData"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Zone *string `pulumi:"zone"`
}

func GetComputeInstanceListOutput(ctx *pulumi.Context, args GetComputeInstanceListOutputArgs, opts ...pulumi.InvokeOption) GetComputeInstanceListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetComputeInstanceListResult, error) {
			args := v.(GetComputeInstanceListArgs)
			r, err := GetComputeInstanceList(ctx, &args, opts...)
			var s GetComputeInstanceListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetComputeInstanceListResultOutput)
}

// A collection of arguments for invoking getComputeInstanceList.
type GetComputeInstanceListOutputArgs struct {
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	DeployTargetId pulumi.StringPtrInput `pulumi:"deployTargetId"`
	// Match against this int
	DiskSize pulumi.IntPtrInput `pulumi:"diskSize"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Match against this bool
	Ipv6 pulumi.BoolPtrInput `pulumi:"ipv6"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Ipv6Address pulumi.StringPtrInput `pulumi:"ipv6Address"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ManagerId pulumi.StringPtrInput `pulumi:"managerId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ManagerType pulumi.StringPtrInput `pulumi:"managerType"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	PublicIpAddress pulumi.StringPtrInput `pulumi:"publicIpAddress"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ReverseDns pulumi.StringPtrInput `pulumi:"reverseDns"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	SshKey pulumi.StringPtrInput `pulumi:"sshKey"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	State pulumi.StringPtrInput `pulumi:"state"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	TemplateId pulumi.StringPtrInput `pulumi:"templateId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	UserData pulumi.StringPtrInput `pulumi:"userData"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetComputeInstanceListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeInstanceListArgs)(nil)).Elem()
}

// A collection of values returned by getComputeInstanceList.
type GetComputeInstanceListResultOutput struct{ *pulumi.OutputState }

func (GetComputeInstanceListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeInstanceListResult)(nil)).Elem()
}

func (o GetComputeInstanceListResultOutput) ToGetComputeInstanceListResultOutput() GetComputeInstanceListResultOutput {
	return o
}

func (o GetComputeInstanceListResultOutput) ToGetComputeInstanceListResultOutputWithContext(ctx context.Context) GetComputeInstanceListResultOutput {
	return o
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) DeployTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.DeployTargetId }).(pulumi.StringPtrOutput)
}

// Match against this int
func (o GetComputeInstanceListResultOutput) DiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *int { return v.DiskSize }).(pulumi.IntPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The list of exoscale*compute*instance.
func (o GetComputeInstanceListResultOutput) Instances() GetComputeInstanceListInstanceArrayOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) []GetComputeInstanceListInstance { return v.Instances }).(GetComputeInstanceListInstanceArrayOutput)
}

// Match against this bool
func (o GetComputeInstanceListResultOutput) Ipv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *bool { return v.Ipv6 }).(pulumi.BoolPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) Ipv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.Ipv6Address }).(pulumi.StringPtrOutput)
}

// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
func (o GetComputeInstanceListResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) ManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.ManagerId }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) ManagerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.ManagerType }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) ReverseDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.ReverseDns }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.SshKey }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) TemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.TemplateId }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.UserData }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetComputeInstanceListResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComputeInstanceListResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetComputeInstanceListResultOutput{})
}
