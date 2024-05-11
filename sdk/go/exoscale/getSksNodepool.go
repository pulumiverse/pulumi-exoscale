// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

func LookupSksNodepool(ctx *pulumi.Context, args *LookupSksNodepoolArgs, opts ...pulumi.InvokeOption) (*LookupSksNodepoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSksNodepoolResult
	err := ctx.Invoke("exoscale:index/getSksNodepool:getSksNodepool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSksNodepool.
type LookupSksNodepoolArgs struct {
	// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
	AntiAffinityGroupIds []string `pulumi:"antiAffinityGroupIds"`
	ClusterId            string   `pulumi:"clusterId"`
	// The pool creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// A deploy target ID.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// A free-form text describing the pool.
	Description *string `pulumi:"description"`
	// The managed instances disk size (GiB; default: `50`).
	DiskSize *int `pulumi:"diskSize"`
	// The ID of this resource.
	Id *string `pulumi:"id"`
	// The underlying exoscale*instance*pool ID.
	InstancePoolId *string `pulumi:"instancePoolId"`
	// The string used to prefix the managed instances name (default `pool`).
	InstancePrefix *string `pulumi:"instancePrefix"`
	// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
	InstanceType *string `pulumi:"instanceType"`
	// Configuration for this nodepool's kubelet image garbage collector
	KubeletImageGcs []GetSksNodepoolKubeletImageGc `pulumi:"kubeletImageGcs"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	Name   *string           `pulumi:"name"`
	// A list of exoscale*private*network (IDs) to be attached to the managed instances.
	PrivateNetworkIds []string `pulumi:"privateNetworkIds"`
	// A list of exoscale*security*group (IDs) to be attached to the managed instances.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	Size             *int     `pulumi:"size"`
	// The current pool state.
	State *string `pulumi:"state"`
	// Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
	StorageLvm *bool `pulumi:"storageLvm"`
	// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) ('taints = { \n\n = "\n\n:\n\n" }').
	Taints map[string]string `pulumi:"taints"`
	// The managed instances template ID.
	TemplateId *string `pulumi:"templateId"`
	// The managed instances version.
	Version *string `pulumi:"version"`
	Zone    string  `pulumi:"zone"`
}

// A collection of values returned by getSksNodepool.
type LookupSksNodepoolResult struct {
	// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
	AntiAffinityGroupIds []string `pulumi:"antiAffinityGroupIds"`
	ClusterId            string   `pulumi:"clusterId"`
	// The pool creation date.
	CreatedAt string `pulumi:"createdAt"`
	// A deploy target ID.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// A free-form text describing the pool.
	Description *string `pulumi:"description"`
	// The managed instances disk size (GiB; default: `50`).
	DiskSize *int `pulumi:"diskSize"`
	// The ID of this resource.
	Id *string `pulumi:"id"`
	// The underlying exoscale*instance*pool ID.
	InstancePoolId string `pulumi:"instancePoolId"`
	// The string used to prefix the managed instances name (default `pool`).
	InstancePrefix *string `pulumi:"instancePrefix"`
	// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
	InstanceType *string `pulumi:"instanceType"`
	// Configuration for this nodepool's kubelet image garbage collector
	KubeletImageGcs []GetSksNodepoolKubeletImageGc `pulumi:"kubeletImageGcs"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	Name   *string           `pulumi:"name"`
	// A list of exoscale*private*network (IDs) to be attached to the managed instances.
	PrivateNetworkIds []string `pulumi:"privateNetworkIds"`
	// A list of exoscale*security*group (IDs) to be attached to the managed instances.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	Size             *int     `pulumi:"size"`
	// The current pool state.
	State string `pulumi:"state"`
	// Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
	StorageLvm *bool `pulumi:"storageLvm"`
	// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) ('taints = { \n\n = "\n\n:\n\n" }').
	Taints map[string]string `pulumi:"taints"`
	// The managed instances template ID.
	TemplateId string `pulumi:"templateId"`
	// The managed instances version.
	Version string `pulumi:"version"`
	Zone    string `pulumi:"zone"`
}

func LookupSksNodepoolOutput(ctx *pulumi.Context, args LookupSksNodepoolOutputArgs, opts ...pulumi.InvokeOption) LookupSksNodepoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSksNodepoolResult, error) {
			args := v.(LookupSksNodepoolArgs)
			r, err := LookupSksNodepool(ctx, &args, opts...)
			var s LookupSksNodepoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSksNodepoolResultOutput)
}

// A collection of arguments for invoking getSksNodepool.
type LookupSksNodepoolOutputArgs struct {
	// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
	AntiAffinityGroupIds pulumi.StringArrayInput `pulumi:"antiAffinityGroupIds"`
	ClusterId            pulumi.StringInput      `pulumi:"clusterId"`
	// The pool creation date.
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// A deploy target ID.
	DeployTargetId pulumi.StringPtrInput `pulumi:"deployTargetId"`
	// A free-form text describing the pool.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The managed instances disk size (GiB; default: `50`).
	DiskSize pulumi.IntPtrInput `pulumi:"diskSize"`
	// The ID of this resource.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The underlying exoscale*instance*pool ID.
	InstancePoolId pulumi.StringPtrInput `pulumi:"instancePoolId"`
	// The string used to prefix the managed instances name (default `pool`).
	InstancePrefix pulumi.StringPtrInput `pulumi:"instancePrefix"`
	// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// Configuration for this nodepool's kubelet image garbage collector
	KubeletImageGcs GetSksNodepoolKubeletImageGcArrayInput `pulumi:"kubeletImageGcs"`
	// A map of key/value labels.
	Labels pulumi.StringMapInput `pulumi:"labels"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	// A list of exoscale*private*network (IDs) to be attached to the managed instances.
	PrivateNetworkIds pulumi.StringArrayInput `pulumi:"privateNetworkIds"`
	// A list of exoscale*security*group (IDs) to be attached to the managed instances.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	Size             pulumi.IntPtrInput      `pulumi:"size"`
	// The current pool state.
	State pulumi.StringPtrInput `pulumi:"state"`
	// Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
	StorageLvm pulumi.BoolPtrInput `pulumi:"storageLvm"`
	// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) ('taints = { \n\n = "\n\n:\n\n" }').
	Taints pulumi.StringMapInput `pulumi:"taints"`
	// The managed instances template ID.
	TemplateId pulumi.StringPtrInput `pulumi:"templateId"`
	// The managed instances version.
	Version pulumi.StringPtrInput `pulumi:"version"`
	Zone    pulumi.StringInput    `pulumi:"zone"`
}

func (LookupSksNodepoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSksNodepoolArgs)(nil)).Elem()
}

// A collection of values returned by getSksNodepool.
type LookupSksNodepoolResultOutput struct{ *pulumi.OutputState }

func (LookupSksNodepoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSksNodepoolResult)(nil)).Elem()
}

func (o LookupSksNodepoolResultOutput) ToLookupSksNodepoolResultOutput() LookupSksNodepoolResultOutput {
	return o
}

func (o LookupSksNodepoolResultOutput) ToLookupSksNodepoolResultOutputWithContext(ctx context.Context) LookupSksNodepoolResultOutput {
	return o
}

// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
func (o LookupSksNodepoolResultOutput) AntiAffinityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) []string { return v.AntiAffinityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupSksNodepoolResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The pool creation date.
func (o LookupSksNodepoolResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A deploy target ID.
func (o LookupSksNodepoolResultOutput) DeployTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *string { return v.DeployTargetId }).(pulumi.StringPtrOutput)
}

// A free-form text describing the pool.
func (o LookupSksNodepoolResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The managed instances disk size (GiB; default: `50`).
func (o LookupSksNodepoolResultOutput) DiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *int { return v.DiskSize }).(pulumi.IntPtrOutput)
}

// The ID of this resource.
func (o LookupSksNodepoolResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The underlying exoscale*instance*pool ID.
func (o LookupSksNodepoolResultOutput) InstancePoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) string { return v.InstancePoolId }).(pulumi.StringOutput)
}

// The string used to prefix the managed instances name (default `pool`).
func (o LookupSksNodepoolResultOutput) InstancePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *string { return v.InstancePrefix }).(pulumi.StringPtrOutput)
}

// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
func (o LookupSksNodepoolResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// Configuration for this nodepool's kubelet image garbage collector
func (o LookupSksNodepoolResultOutput) KubeletImageGcs() GetSksNodepoolKubeletImageGcArrayOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) []GetSksNodepoolKubeletImageGc { return v.KubeletImageGcs }).(GetSksNodepoolKubeletImageGcArrayOutput)
}

// A map of key/value labels.
func (o LookupSksNodepoolResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupSksNodepoolResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of exoscale*private*network (IDs) to be attached to the managed instances.
func (o LookupSksNodepoolResultOutput) PrivateNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) []string { return v.PrivateNetworkIds }).(pulumi.StringArrayOutput)
}

// A list of exoscale*security*group (IDs) to be attached to the managed instances.
func (o LookupSksNodepoolResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupSksNodepoolResultOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *int { return v.Size }).(pulumi.IntPtrOutput)
}

// The current pool state.
func (o LookupSksNodepoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) string { return v.State }).(pulumi.StringOutput)
}

// Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
func (o LookupSksNodepoolResultOutput) StorageLvm() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) *bool { return v.StorageLvm }).(pulumi.BoolPtrOutput)
}

// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) ('taints = { \n\n = "\n\n:\n\n" }').
func (o LookupSksNodepoolResultOutput) Taints() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) map[string]string { return v.Taints }).(pulumi.StringMapOutput)
}

// The managed instances template ID.
func (o LookupSksNodepoolResultOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) string { return v.TemplateId }).(pulumi.StringOutput)
}

// The managed instances version.
func (o LookupSksNodepoolResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupSksNodepoolResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSksNodepoolResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSksNodepoolResultOutput{})
}
