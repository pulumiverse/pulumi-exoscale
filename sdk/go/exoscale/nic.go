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

// !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use ComputeInstance `networkInterface` block instead.
type NIC struct {
	pulumi.CustomResourceState

	// ❗ The compute instance ID.
	ComputeId pulumi.StringOutput `pulumi:"computeId"`
	Gateway   pulumi.StringOutput `pulumi:"gateway"`
	// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The NIC MAC address.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	Netmask    pulumi.StringOutput `pulumi:"netmask"`
	// ❗ The private network ID.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
}

// NewNIC registers a new resource with the given unique name, arguments, and options.
func NewNIC(ctx *pulumi.Context,
	name string, args *NICArgs, opts ...pulumi.ResourceOption) (*NIC, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeId'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NIC
	err := ctx.RegisterResource("exoscale:index/nIC:NIC", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNIC gets an existing NIC resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNIC(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NICState, opts ...pulumi.ResourceOption) (*NIC, error) {
	var resource NIC
	err := ctx.ReadResource("exoscale:index/nIC:NIC", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NIC resources.
type nicState struct {
	// ❗ The compute instance ID.
	ComputeId *string `pulumi:"computeId"`
	Gateway   *string `pulumi:"gateway"`
	// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
	IpAddress *string `pulumi:"ipAddress"`
	// The NIC MAC address.
	MacAddress *string `pulumi:"macAddress"`
	Netmask    *string `pulumi:"netmask"`
	// ❗ The private network ID.
	NetworkId *string `pulumi:"networkId"`
}

type NICState struct {
	// ❗ The compute instance ID.
	ComputeId pulumi.StringPtrInput
	Gateway   pulumi.StringPtrInput
	// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
	IpAddress pulumi.StringPtrInput
	// The NIC MAC address.
	MacAddress pulumi.StringPtrInput
	Netmask    pulumi.StringPtrInput
	// ❗ The private network ID.
	NetworkId pulumi.StringPtrInput
}

func (NICState) ElementType() reflect.Type {
	return reflect.TypeOf((*nicState)(nil)).Elem()
}

type nicArgs struct {
	// ❗ The compute instance ID.
	ComputeId string `pulumi:"computeId"`
	// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
	IpAddress *string `pulumi:"ipAddress"`
	// ❗ The private network ID.
	NetworkId string `pulumi:"networkId"`
}

// The set of arguments for constructing a NIC resource.
type NICArgs struct {
	// ❗ The compute instance ID.
	ComputeId pulumi.StringInput
	// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
	IpAddress pulumi.StringPtrInput
	// ❗ The private network ID.
	NetworkId pulumi.StringInput
}

func (NICArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nicArgs)(nil)).Elem()
}

type NICInput interface {
	pulumi.Input

	ToNICOutput() NICOutput
	ToNICOutputWithContext(ctx context.Context) NICOutput
}

func (*NIC) ElementType() reflect.Type {
	return reflect.TypeOf((**NIC)(nil)).Elem()
}

func (i *NIC) ToNICOutput() NICOutput {
	return i.ToNICOutputWithContext(context.Background())
}

func (i *NIC) ToNICOutputWithContext(ctx context.Context) NICOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NICOutput)
}

// NICArrayInput is an input type that accepts NICArray and NICArrayOutput values.
// You can construct a concrete instance of `NICArrayInput` via:
//
//	NICArray{ NICArgs{...} }
type NICArrayInput interface {
	pulumi.Input

	ToNICArrayOutput() NICArrayOutput
	ToNICArrayOutputWithContext(context.Context) NICArrayOutput
}

type NICArray []NICInput

func (NICArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NIC)(nil)).Elem()
}

func (i NICArray) ToNICArrayOutput() NICArrayOutput {
	return i.ToNICArrayOutputWithContext(context.Background())
}

func (i NICArray) ToNICArrayOutputWithContext(ctx context.Context) NICArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NICArrayOutput)
}

// NICMapInput is an input type that accepts NICMap and NICMapOutput values.
// You can construct a concrete instance of `NICMapInput` via:
//
//	NICMap{ "key": NICArgs{...} }
type NICMapInput interface {
	pulumi.Input

	ToNICMapOutput() NICMapOutput
	ToNICMapOutputWithContext(context.Context) NICMapOutput
}

type NICMap map[string]NICInput

func (NICMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NIC)(nil)).Elem()
}

func (i NICMap) ToNICMapOutput() NICMapOutput {
	return i.ToNICMapOutputWithContext(context.Background())
}

func (i NICMap) ToNICMapOutputWithContext(ctx context.Context) NICMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NICMapOutput)
}

type NICOutput struct{ *pulumi.OutputState }

func (NICOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NIC)(nil)).Elem()
}

func (o NICOutput) ToNICOutput() NICOutput {
	return o
}

func (o NICOutput) ToNICOutputWithContext(ctx context.Context) NICOutput {
	return o
}

// ❗ The compute instance ID.
func (o NICOutput) ComputeId() pulumi.StringOutput {
	return o.ApplyT(func(v *NIC) pulumi.StringOutput { return v.ComputeId }).(pulumi.StringOutput)
}

func (o NICOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NIC) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
func (o NICOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NIC) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The NIC MAC address.
func (o NICOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NIC) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

func (o NICOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v *NIC) pulumi.StringOutput { return v.Netmask }).(pulumi.StringOutput)
}

// ❗ The private network ID.
func (o NICOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *NIC) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

type NICArrayOutput struct{ *pulumi.OutputState }

func (NICArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NIC)(nil)).Elem()
}

func (o NICArrayOutput) ToNICArrayOutput() NICArrayOutput {
	return o
}

func (o NICArrayOutput) ToNICArrayOutputWithContext(ctx context.Context) NICArrayOutput {
	return o
}

func (o NICArrayOutput) Index(i pulumi.IntInput) NICOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NIC {
		return vs[0].([]*NIC)[vs[1].(int)]
	}).(NICOutput)
}

type NICMapOutput struct{ *pulumi.OutputState }

func (NICMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NIC)(nil)).Elem()
}

func (o NICMapOutput) ToNICMapOutput() NICMapOutput {
	return o
}

func (o NICMapOutput) ToNICMapOutputWithContext(ctx context.Context) NICMapOutput {
	return o
}

func (o NICMapOutput) MapIndex(k pulumi.StringInput) NICOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NIC {
		return vs[0].(map[string]*NIC)[vs[1].(string)]
	}).(NICOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NICInput)(nil)).Elem(), &NIC{})
	pulumi.RegisterInputType(reflect.TypeOf((*NICArrayInput)(nil)).Elem(), NICArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NICMapInput)(nil)).Elem(), NICMap{})
	pulumi.RegisterOutputType(NICOutput{})
	pulumi.RegisterOutputType(NICArrayOutput{})
	pulumi.RegisterOutputType(NICMapOutput{})
}
