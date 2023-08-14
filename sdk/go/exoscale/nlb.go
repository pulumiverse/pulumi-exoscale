// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// ## Import
//
// An existing network load balancer (NLB) may be imported by `<ID>@<zone>`console
//
// ```sh
//
//	$ pulumi import exoscale:index/nLB:NLB \
//
// ```
//
//	exoscale_nlb.my_nlb \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
type NLB struct {
	pulumi.CustomResourceState

	// The NLB creation date.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A free-form text describing the NLB.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The NLB IPv4 address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// A map of key/value labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The network load balancer (NLB) name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of the exoscale*nlb*service (names).
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// The current NLB state.
	State pulumi.StringOutput `pulumi:"state"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNLB registers a new resource with the given unique name, arguments, and options.
func NewNLB(ctx *pulumi.Context,
	name string, args *NLBArgs, opts ...pulumi.ResourceOption) (*NLB, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NLB
	err := ctx.RegisterResource("exoscale:index/nLB:NLB", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNLB gets an existing NLB resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNLB(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NLBState, opts ...pulumi.ResourceOption) (*NLB, error) {
	var resource NLB
	err := ctx.ReadResource("exoscale:index/nLB:NLB", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NLB resources.
type nlbState struct {
	// The NLB creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// A free-form text describing the NLB.
	Description *string `pulumi:"description"`
	// The NLB IPv4 address.
	IpAddress *string `pulumi:"ipAddress"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The network load balancer (NLB) name.
	Name *string `pulumi:"name"`
	// The list of the exoscale*nlb*service (names).
	Services []string `pulumi:"services"`
	// The current NLB state.
	State *string `pulumi:"state"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type NLBState struct {
	// The NLB creation date.
	CreatedAt pulumi.StringPtrInput
	// A free-form text describing the NLB.
	Description pulumi.StringPtrInput
	// The NLB IPv4 address.
	IpAddress pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The network load balancer (NLB) name.
	Name pulumi.StringPtrInput
	// The list of the exoscale*nlb*service (names).
	Services pulumi.StringArrayInput
	// The current NLB state.
	State pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (NLBState) ElementType() reflect.Type {
	return reflect.TypeOf((*nlbState)(nil)).Elem()
}

type nlbArgs struct {
	// A free-form text describing the NLB.
	Description *string `pulumi:"description"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The network load balancer (NLB) name.
	Name *string `pulumi:"name"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a NLB resource.
type NLBArgs struct {
	// A free-form text describing the NLB.
	Description pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The network load balancer (NLB) name.
	Name pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (NLBArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nlbArgs)(nil)).Elem()
}

type NLBInput interface {
	pulumi.Input

	ToNLBOutput() NLBOutput
	ToNLBOutputWithContext(ctx context.Context) NLBOutput
}

func (*NLB) ElementType() reflect.Type {
	return reflect.TypeOf((**NLB)(nil)).Elem()
}

func (i *NLB) ToNLBOutput() NLBOutput {
	return i.ToNLBOutputWithContext(context.Background())
}

func (i *NLB) ToNLBOutputWithContext(ctx context.Context) NLBOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NLBOutput)
}

// NLBArrayInput is an input type that accepts NLBArray and NLBArrayOutput values.
// You can construct a concrete instance of `NLBArrayInput` via:
//
//	NLBArray{ NLBArgs{...} }
type NLBArrayInput interface {
	pulumi.Input

	ToNLBArrayOutput() NLBArrayOutput
	ToNLBArrayOutputWithContext(context.Context) NLBArrayOutput
}

type NLBArray []NLBInput

func (NLBArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NLB)(nil)).Elem()
}

func (i NLBArray) ToNLBArrayOutput() NLBArrayOutput {
	return i.ToNLBArrayOutputWithContext(context.Background())
}

func (i NLBArray) ToNLBArrayOutputWithContext(ctx context.Context) NLBArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NLBArrayOutput)
}

// NLBMapInput is an input type that accepts NLBMap and NLBMapOutput values.
// You can construct a concrete instance of `NLBMapInput` via:
//
//	NLBMap{ "key": NLBArgs{...} }
type NLBMapInput interface {
	pulumi.Input

	ToNLBMapOutput() NLBMapOutput
	ToNLBMapOutputWithContext(context.Context) NLBMapOutput
}

type NLBMap map[string]NLBInput

func (NLBMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NLB)(nil)).Elem()
}

func (i NLBMap) ToNLBMapOutput() NLBMapOutput {
	return i.ToNLBMapOutputWithContext(context.Background())
}

func (i NLBMap) ToNLBMapOutputWithContext(ctx context.Context) NLBMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NLBMapOutput)
}

type NLBOutput struct{ *pulumi.OutputState }

func (NLBOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NLB)(nil)).Elem()
}

func (o NLBOutput) ToNLBOutput() NLBOutput {
	return o
}

func (o NLBOutput) ToNLBOutputWithContext(ctx context.Context) NLBOutput {
	return o
}

// The NLB creation date.
func (o NLBOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NLB) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A free-form text describing the NLB.
func (o NLBOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NLB) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The NLB IPv4 address.
func (o NLBOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NLB) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// A map of key/value labels.
func (o NLBOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NLB) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The network load balancer (NLB) name.
func (o NLBOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NLB) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of the exoscale*nlb*service (names).
func (o NLBOutput) Services() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NLB) pulumi.StringArrayOutput { return v.Services }).(pulumi.StringArrayOutput)
}

// The current NLB state.
func (o NLBOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NLB) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o NLBOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *NLB) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type NLBArrayOutput struct{ *pulumi.OutputState }

func (NLBArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NLB)(nil)).Elem()
}

func (o NLBArrayOutput) ToNLBArrayOutput() NLBArrayOutput {
	return o
}

func (o NLBArrayOutput) ToNLBArrayOutputWithContext(ctx context.Context) NLBArrayOutput {
	return o
}

func (o NLBArrayOutput) Index(i pulumi.IntInput) NLBOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NLB {
		return vs[0].([]*NLB)[vs[1].(int)]
	}).(NLBOutput)
}

type NLBMapOutput struct{ *pulumi.OutputState }

func (NLBMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NLB)(nil)).Elem()
}

func (o NLBMapOutput) ToNLBMapOutput() NLBMapOutput {
	return o
}

func (o NLBMapOutput) ToNLBMapOutputWithContext(ctx context.Context) NLBMapOutput {
	return o
}

func (o NLBMapOutput) MapIndex(k pulumi.StringInput) NLBOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NLB {
		return vs[0].(map[string]*NLB)[vs[1].(string)]
	}).(NLBOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NLBInput)(nil)).Elem(), &NLB{})
	pulumi.RegisterInputType(reflect.TypeOf((*NLBArrayInput)(nil)).Elem(), NLBArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NLBMapInput)(nil)).Elem(), NLBMap{})
	pulumi.RegisterOutputType(NLBOutput{})
	pulumi.RegisterOutputType(NLBArrayOutput{})
	pulumi.RegisterOutputType(NLBMapOutput{})
}
