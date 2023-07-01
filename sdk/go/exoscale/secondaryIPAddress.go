// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use ComputeInstance `elasticIpIds` list instead.
type SecondaryIPAddress struct {
	pulumi.CustomResourceState

	// ❗ The compute instance ID.
	ComputeId pulumi.StringOutput `pulumi:"computeId"`
	// ❗ The Elastic IP (EIP) address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The network (ID) the compute instance NIC is attached to.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// The network interface (NIC) ID.
	NicId pulumi.StringOutput `pulumi:"nicId"`
}

// NewSecondaryIPAddress registers a new resource with the given unique name, arguments, and options.
func NewSecondaryIPAddress(ctx *pulumi.Context,
	name string, args *SecondaryIPAddressArgs, opts ...pulumi.ResourceOption) (*SecondaryIPAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeId'")
	}
	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SecondaryIPAddress
	err := ctx.RegisterResource("exoscale:index/secondaryIPAddress:SecondaryIPAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecondaryIPAddress gets an existing SecondaryIPAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecondaryIPAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecondaryIPAddressState, opts ...pulumi.ResourceOption) (*SecondaryIPAddress, error) {
	var resource SecondaryIPAddress
	err := ctx.ReadResource("exoscale:index/secondaryIPAddress:SecondaryIPAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecondaryIPAddress resources.
type secondaryIPAddressState struct {
	// ❗ The compute instance ID.
	ComputeId *string `pulumi:"computeId"`
	// ❗ The Elastic IP (EIP) address.
	IpAddress *string `pulumi:"ipAddress"`
	// The network (ID) the compute instance NIC is attached to.
	NetworkId *string `pulumi:"networkId"`
	// The network interface (NIC) ID.
	NicId *string `pulumi:"nicId"`
}

type SecondaryIPAddressState struct {
	// ❗ The compute instance ID.
	ComputeId pulumi.StringPtrInput
	// ❗ The Elastic IP (EIP) address.
	IpAddress pulumi.StringPtrInput
	// The network (ID) the compute instance NIC is attached to.
	NetworkId pulumi.StringPtrInput
	// The network interface (NIC) ID.
	NicId pulumi.StringPtrInput
}

func (SecondaryIPAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*secondaryIPAddressState)(nil)).Elem()
}

type secondaryIPAddressArgs struct {
	// ❗ The compute instance ID.
	ComputeId string `pulumi:"computeId"`
	// ❗ The Elastic IP (EIP) address.
	IpAddress string `pulumi:"ipAddress"`
}

// The set of arguments for constructing a SecondaryIPAddress resource.
type SecondaryIPAddressArgs struct {
	// ❗ The compute instance ID.
	ComputeId pulumi.StringInput
	// ❗ The Elastic IP (EIP) address.
	IpAddress pulumi.StringInput
}

func (SecondaryIPAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secondaryIPAddressArgs)(nil)).Elem()
}

type SecondaryIPAddressInput interface {
	pulumi.Input

	ToSecondaryIPAddressOutput() SecondaryIPAddressOutput
	ToSecondaryIPAddressOutputWithContext(ctx context.Context) SecondaryIPAddressOutput
}

func (*SecondaryIPAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**SecondaryIPAddress)(nil)).Elem()
}

func (i *SecondaryIPAddress) ToSecondaryIPAddressOutput() SecondaryIPAddressOutput {
	return i.ToSecondaryIPAddressOutputWithContext(context.Background())
}

func (i *SecondaryIPAddress) ToSecondaryIPAddressOutputWithContext(ctx context.Context) SecondaryIPAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecondaryIPAddressOutput)
}

// SecondaryIPAddressArrayInput is an input type that accepts SecondaryIPAddressArray and SecondaryIPAddressArrayOutput values.
// You can construct a concrete instance of `SecondaryIPAddressArrayInput` via:
//
//	SecondaryIPAddressArray{ SecondaryIPAddressArgs{...} }
type SecondaryIPAddressArrayInput interface {
	pulumi.Input

	ToSecondaryIPAddressArrayOutput() SecondaryIPAddressArrayOutput
	ToSecondaryIPAddressArrayOutputWithContext(context.Context) SecondaryIPAddressArrayOutput
}

type SecondaryIPAddressArray []SecondaryIPAddressInput

func (SecondaryIPAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecondaryIPAddress)(nil)).Elem()
}

func (i SecondaryIPAddressArray) ToSecondaryIPAddressArrayOutput() SecondaryIPAddressArrayOutput {
	return i.ToSecondaryIPAddressArrayOutputWithContext(context.Background())
}

func (i SecondaryIPAddressArray) ToSecondaryIPAddressArrayOutputWithContext(ctx context.Context) SecondaryIPAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecondaryIPAddressArrayOutput)
}

// SecondaryIPAddressMapInput is an input type that accepts SecondaryIPAddressMap and SecondaryIPAddressMapOutput values.
// You can construct a concrete instance of `SecondaryIPAddressMapInput` via:
//
//	SecondaryIPAddressMap{ "key": SecondaryIPAddressArgs{...} }
type SecondaryIPAddressMapInput interface {
	pulumi.Input

	ToSecondaryIPAddressMapOutput() SecondaryIPAddressMapOutput
	ToSecondaryIPAddressMapOutputWithContext(context.Context) SecondaryIPAddressMapOutput
}

type SecondaryIPAddressMap map[string]SecondaryIPAddressInput

func (SecondaryIPAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecondaryIPAddress)(nil)).Elem()
}

func (i SecondaryIPAddressMap) ToSecondaryIPAddressMapOutput() SecondaryIPAddressMapOutput {
	return i.ToSecondaryIPAddressMapOutputWithContext(context.Background())
}

func (i SecondaryIPAddressMap) ToSecondaryIPAddressMapOutputWithContext(ctx context.Context) SecondaryIPAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecondaryIPAddressMapOutput)
}

type SecondaryIPAddressOutput struct{ *pulumi.OutputState }

func (SecondaryIPAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecondaryIPAddress)(nil)).Elem()
}

func (o SecondaryIPAddressOutput) ToSecondaryIPAddressOutput() SecondaryIPAddressOutput {
	return o
}

func (o SecondaryIPAddressOutput) ToSecondaryIPAddressOutputWithContext(ctx context.Context) SecondaryIPAddressOutput {
	return o
}

// ❗ The compute instance ID.
func (o SecondaryIPAddressOutput) ComputeId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecondaryIPAddress) pulumi.StringOutput { return v.ComputeId }).(pulumi.StringOutput)
}

// ❗ The Elastic IP (EIP) address.
func (o SecondaryIPAddressOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SecondaryIPAddress) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The network (ID) the compute instance NIC is attached to.
func (o SecondaryIPAddressOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecondaryIPAddress) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// The network interface (NIC) ID.
func (o SecondaryIPAddressOutput) NicId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecondaryIPAddress) pulumi.StringOutput { return v.NicId }).(pulumi.StringOutput)
}

type SecondaryIPAddressArrayOutput struct{ *pulumi.OutputState }

func (SecondaryIPAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecondaryIPAddress)(nil)).Elem()
}

func (o SecondaryIPAddressArrayOutput) ToSecondaryIPAddressArrayOutput() SecondaryIPAddressArrayOutput {
	return o
}

func (o SecondaryIPAddressArrayOutput) ToSecondaryIPAddressArrayOutputWithContext(ctx context.Context) SecondaryIPAddressArrayOutput {
	return o
}

func (o SecondaryIPAddressArrayOutput) Index(i pulumi.IntInput) SecondaryIPAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecondaryIPAddress {
		return vs[0].([]*SecondaryIPAddress)[vs[1].(int)]
	}).(SecondaryIPAddressOutput)
}

type SecondaryIPAddressMapOutput struct{ *pulumi.OutputState }

func (SecondaryIPAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecondaryIPAddress)(nil)).Elem()
}

func (o SecondaryIPAddressMapOutput) ToSecondaryIPAddressMapOutput() SecondaryIPAddressMapOutput {
	return o
}

func (o SecondaryIPAddressMapOutput) ToSecondaryIPAddressMapOutputWithContext(ctx context.Context) SecondaryIPAddressMapOutput {
	return o
}

func (o SecondaryIPAddressMapOutput) MapIndex(k pulumi.StringInput) SecondaryIPAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecondaryIPAddress {
		return vs[0].(map[string]*SecondaryIPAddress)[vs[1].(string)]
	}).(SecondaryIPAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecondaryIPAddressInput)(nil)).Elem(), &SecondaryIPAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecondaryIPAddressArrayInput)(nil)).Elem(), SecondaryIPAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecondaryIPAddressMapInput)(nil)).Elem(), SecondaryIPAddressMap{})
	pulumi.RegisterOutputType(SecondaryIPAddressOutput{})
	pulumi.RegisterOutputType(SecondaryIPAddressArrayOutput{})
	pulumi.RegisterOutputType(SecondaryIPAddressMapOutput{})
}
