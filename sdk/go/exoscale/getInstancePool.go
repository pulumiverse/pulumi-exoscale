// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// Fetch Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/) data.
//
// Corresponding resource: exoscale_instance_pool.
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
//			myInstancePool, err := exoscale.LookupInstancePool(ctx, &exoscale.LookupInstancePoolArgs{
//				Zone: "ch-gva-2",
//				Name: pulumi.StringRef("my-instance-pool"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("myInstancePoolId", myInstancePool.Id)
//			return nil
//		})
//	}
//
// ```
//
// Please refer to the examples
// directory for complete configuration examples.
func LookupInstancePool(ctx *pulumi.Context, args *LookupInstancePoolArgs, opts ...pulumi.InvokeOption) (*LookupInstancePoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstancePoolResult
	err := ctx.Invoke("exoscale:index/getInstancePool:getInstancePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstancePool.
type LookupInstancePoolArgs struct {
	// The instance pool ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The pool name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getInstancePool.
type LookupInstancePoolResult struct {
	// The list of attached exoscale*anti*affinity_group (IDs).
	AffinityGroupIds []string `pulumi:"affinityGroupIds"`
	// The deploy target ID.
	DeployTargetId string `pulumi:"deployTargetId"`
	// The instance pool description.
	Description string `pulumi:"description"`
	// The managed instances disk size.
	DiskSize int `pulumi:"diskSize"`
	// The list of attached exoscale*elastic*ip (IDs).
	ElasticIpIds []string `pulumi:"elasticIpIds"`
	// The instance pool ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The string used to prefix the managed instances name.
	InstancePrefix string `pulumi:"instancePrefix"`
	// The managed instances type.
	InstanceType string `pulumi:"instanceType"`
	// The list of managed instances. Structure is documented below.
	Instances []GetInstancePoolInstance `pulumi:"instances"`
	// Whether IPv6 is enabled on managed instances.
	Ipv6 bool `pulumi:"ipv6"`
	// The exoscale*ssh*key (name) authorized on the managed instances.
	KeyPair string `pulumi:"keyPair"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The pool name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The list of attached exoscale*private*network (IDs).
	NetworkIds []string `pulumi:"networkIds"`
	// The list of attached exoscale*security*group (IDs).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The number managed instances.
	Size int `pulumi:"size"`
	// The pool state.
	State string `pulumi:"state"`
	// The managed instances getTemplate ID.
	TemplateId string `pulumi:"templateId"`
	// [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
	UserData string `pulumi:"userData"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

func LookupInstancePoolOutput(ctx *pulumi.Context, args LookupInstancePoolOutputArgs, opts ...pulumi.InvokeOption) LookupInstancePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstancePoolResult, error) {
			args := v.(LookupInstancePoolArgs)
			r, err := LookupInstancePool(ctx, &args, opts...)
			var s LookupInstancePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstancePoolResultOutput)
}

// A collection of arguments for invoking getInstancePool.
type LookupInstancePoolOutputArgs struct {
	// The instance pool ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// A map of key/value labels.
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// The pool name to match (conflicts with `id`).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupInstancePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstancePoolArgs)(nil)).Elem()
}

// A collection of values returned by getInstancePool.
type LookupInstancePoolResultOutput struct{ *pulumi.OutputState }

func (LookupInstancePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstancePoolResult)(nil)).Elem()
}

func (o LookupInstancePoolResultOutput) ToLookupInstancePoolResultOutput() LookupInstancePoolResultOutput {
	return o
}

func (o LookupInstancePoolResultOutput) ToLookupInstancePoolResultOutputWithContext(ctx context.Context) LookupInstancePoolResultOutput {
	return o
}

// The list of attached exoscale*anti*affinity_group (IDs).
func (o LookupInstancePoolResultOutput) AffinityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) []string { return v.AffinityGroupIds }).(pulumi.StringArrayOutput)
}

// The deploy target ID.
func (o LookupInstancePoolResultOutput) DeployTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.DeployTargetId }).(pulumi.StringOutput)
}

// The instance pool description.
func (o LookupInstancePoolResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.Description }).(pulumi.StringOutput)
}

// The managed instances disk size.
func (o LookupInstancePoolResultOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) int { return v.DiskSize }).(pulumi.IntOutput)
}

// The list of attached exoscale*elastic*ip (IDs).
func (o LookupInstancePoolResultOutput) ElasticIpIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) []string { return v.ElasticIpIds }).(pulumi.StringArrayOutput)
}

// The instance pool ID to match (conflicts with `name`).
func (o LookupInstancePoolResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The string used to prefix the managed instances name.
func (o LookupInstancePoolResultOutput) InstancePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.InstancePrefix }).(pulumi.StringOutput)
}

// The managed instances type.
func (o LookupInstancePoolResultOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.InstanceType }).(pulumi.StringOutput)
}

// The list of managed instances. Structure is documented below.
func (o LookupInstancePoolResultOutput) Instances() GetInstancePoolInstanceArrayOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) []GetInstancePoolInstance { return v.Instances }).(GetInstancePoolInstanceArrayOutput)
}

// Whether IPv6 is enabled on managed instances.
func (o LookupInstancePoolResultOutput) Ipv6() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) bool { return v.Ipv6 }).(pulumi.BoolOutput)
}

// The exoscale*ssh*key (name) authorized on the managed instances.
func (o LookupInstancePoolResultOutput) KeyPair() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.KeyPair }).(pulumi.StringOutput)
}

// A map of key/value labels.
func (o LookupInstancePoolResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The pool name to match (conflicts with `id`).
func (o LookupInstancePoolResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The list of attached exoscale*private*network (IDs).
func (o LookupInstancePoolResultOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) []string { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

// The list of attached exoscale*security*group (IDs).
func (o LookupInstancePoolResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The number managed instances.
func (o LookupInstancePoolResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) int { return v.Size }).(pulumi.IntOutput)
}

// The pool state.
func (o LookupInstancePoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.State }).(pulumi.StringOutput)
}

// The managed instances getTemplate ID.
func (o LookupInstancePoolResultOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.TemplateId }).(pulumi.StringOutput)
}

// [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
func (o LookupInstancePoolResultOutput) UserData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.UserData }).(pulumi.StringOutput)
}

// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o LookupInstancePoolResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstancePoolResultOutput{})
}
