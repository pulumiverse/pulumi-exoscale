// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// Manage Exoscale [Compute Instances](https://community.exoscale.com/documentation/compute/).
//
// Corresponding data sources: exoscale_compute_instance, exoscale_compute_instance_list.
//
// After the creation, you can retrieve the password of an instance with [Exoscale CLI](https://github.com/exoscale/cli): `exo compute instance reveal-password NAME`.
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
//			myTemplate, err := exoscale.GetTemplate(ctx, &exoscale.GetTemplateArgs{
//				Zone: "ch-gva-2",
//				Name: pulumi.StringRef("Linux Ubuntu 22.04 LTS 64-bit"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = exoscale.NewComputeInstance(ctx, "myInstance", &exoscale.ComputeInstanceArgs{
//				Zone:       pulumi.String("ch-gva-2"),
//				TemplateId: pulumi.String(myTemplate.Id),
//				Type:       pulumi.String("standard.medium"),
//				DiskSize:   pulumi.Int(10),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Please refer to the examples
// directory for complete configuration examples.
//
// ## Import
//
// An existing compute instance may be imported by `<ID>@<zone>`:
//
// ```sh
// $ pulumi import exoscale:index/computeInstance:ComputeInstance \
// ```
//
//	exoscale_compute_instance.my_instance \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
type ComputeInstance struct {
	pulumi.CustomResourceState

	// ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
	AntiAffinityGroupIds pulumi.StringArrayOutput `pulumi:"antiAffinityGroupIds"`
	// A list of exoscale*block*storage_volume (ID) to attach to the instance.
	BlockStorageVolumeIds pulumi.StringArrayOutput `pulumi:"blockStorageVolumeIds"`
	// The instance creation date.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// ❗ A deploy target ID.
	DeployTargetId pulumi.StringPtrOutput `pulumi:"deployTargetId"`
	// Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
	DestroyProtected pulumi.BoolPtrOutput `pulumi:"destroyProtected"`
	// The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// A list of exoscale*elastic*ip (IDs) to attach to the instance.
	ElasticIpIds pulumi.StringArrayOutput `pulumi:"elasticIpIds"`
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ipv6 pulumi.BoolPtrOutput `pulumi:"ipv6"`
	// The instance (main network interface) IPv6 address (if enabled).
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// A map of key/value labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// MAC address
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The compute instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private network interfaces (may be specified multiple times). Structure is documented below.
	NetworkInterfaces ComputeInstanceNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// Whether the instance is private (no public IP addresses; default: false)
	Private pulumi.BoolPtrOutput `pulumi:"private"`
	// A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument instead.
	//
	// Deprecated: Use the networkInterface block instead.
	PrivateNetworkIds pulumi.StringArrayOutput `pulumi:"privateNetworkIds"`
	// The instance (main network interface) IPv4 address.
	PublicIpAddress pulumi.StringOutput `pulumi:"publicIpAddress"`
	// Domain name for reverse DNS record.
	ReverseDns pulumi.StringPtrOutput `pulumi:"reverseDns"`
	// A list of exoscale*security*group (IDs) to attach to the instance.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	//
	// Deprecated: Use sshKeys instead
	SshKey pulumi.StringPtrOutput `pulumi:"sshKey"`
	// The list of exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	SshKeys pulumi.StringArrayOutput `pulumi:"sshKeys"`
	// The instance state (`running` or `stopped`; default: `running`).
	State pulumi.StringOutput `pulumi:"state"`
	// ❗ The getTemplate (ID) to use when creating the instance.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
	Type pulumi.StringOutput `pulumi:"type"`
	// [cloud-init](https://cloudinit.readthedocs.io/) configuration.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewComputeInstance registers a new resource with the given unique name, arguments, and options.
func NewComputeInstance(ctx *pulumi.Context,
	name string, args *ComputeInstanceArgs, opts ...pulumi.ResourceOption) (*ComputeInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskSize == nil {
		return nil, errors.New("invalid value for required argument 'DiskSize'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ComputeInstance
	err := ctx.RegisterResource("exoscale:index/computeInstance:ComputeInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeInstance gets an existing ComputeInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeInstanceState, opts ...pulumi.ResourceOption) (*ComputeInstance, error) {
	var resource ComputeInstance
	err := ctx.ReadResource("exoscale:index/computeInstance:ComputeInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeInstance resources.
type computeInstanceState struct {
	// ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
	AntiAffinityGroupIds []string `pulumi:"antiAffinityGroupIds"`
	// A list of exoscale*block*storage_volume (ID) to attach to the instance.
	BlockStorageVolumeIds []string `pulumi:"blockStorageVolumeIds"`
	// The instance creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// ❗ A deploy target ID.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
	DestroyProtected *bool `pulumi:"destroyProtected"`
	// The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
	DiskSize *int `pulumi:"diskSize"`
	// A list of exoscale*elastic*ip (IDs) to attach to the instance.
	ElasticIpIds []string `pulumi:"elasticIpIds"`
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ipv6 *bool `pulumi:"ipv6"`
	// The instance (main network interface) IPv6 address (if enabled).
	Ipv6Address *string `pulumi:"ipv6Address"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// MAC address
	MacAddress *string `pulumi:"macAddress"`
	// The compute instance name.
	Name *string `pulumi:"name"`
	// Private network interfaces (may be specified multiple times). Structure is documented below.
	NetworkInterfaces []ComputeInstanceNetworkInterface `pulumi:"networkInterfaces"`
	// Whether the instance is private (no public IP addresses; default: false)
	Private *bool `pulumi:"private"`
	// A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument instead.
	//
	// Deprecated: Use the networkInterface block instead.
	PrivateNetworkIds []string `pulumi:"privateNetworkIds"`
	// The instance (main network interface) IPv4 address.
	PublicIpAddress *string `pulumi:"publicIpAddress"`
	// Domain name for reverse DNS record.
	ReverseDns *string `pulumi:"reverseDns"`
	// A list of exoscale*security*group (IDs) to attach to the instance.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	//
	// Deprecated: Use sshKeys instead
	SshKey *string `pulumi:"sshKey"`
	// The list of exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	SshKeys []string `pulumi:"sshKeys"`
	// The instance state (`running` or `stopped`; default: `running`).
	State *string `pulumi:"state"`
	// ❗ The getTemplate (ID) to use when creating the instance.
	TemplateId *string `pulumi:"templateId"`
	// The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
	Type *string `pulumi:"type"`
	// [cloud-init](https://cloudinit.readthedocs.io/) configuration.
	UserData *string `pulumi:"userData"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type ComputeInstanceState struct {
	// ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
	AntiAffinityGroupIds pulumi.StringArrayInput
	// A list of exoscale*block*storage_volume (ID) to attach to the instance.
	BlockStorageVolumeIds pulumi.StringArrayInput
	// The instance creation date.
	CreatedAt pulumi.StringPtrInput
	// ❗ A deploy target ID.
	DeployTargetId pulumi.StringPtrInput
	// Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
	DestroyProtected pulumi.BoolPtrInput
	// The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
	DiskSize pulumi.IntPtrInput
	// A list of exoscale*elastic*ip (IDs) to attach to the instance.
	ElasticIpIds pulumi.StringArrayInput
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ipv6 pulumi.BoolPtrInput
	// The instance (main network interface) IPv6 address (if enabled).
	Ipv6Address pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// MAC address
	MacAddress pulumi.StringPtrInput
	// The compute instance name.
	Name pulumi.StringPtrInput
	// Private network interfaces (may be specified multiple times). Structure is documented below.
	NetworkInterfaces ComputeInstanceNetworkInterfaceArrayInput
	// Whether the instance is private (no public IP addresses; default: false)
	Private pulumi.BoolPtrInput
	// A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument instead.
	//
	// Deprecated: Use the networkInterface block instead.
	PrivateNetworkIds pulumi.StringArrayInput
	// The instance (main network interface) IPv4 address.
	PublicIpAddress pulumi.StringPtrInput
	// Domain name for reverse DNS record.
	ReverseDns pulumi.StringPtrInput
	// A list of exoscale*security*group (IDs) to attach to the instance.
	SecurityGroupIds pulumi.StringArrayInput
	// The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	//
	// Deprecated: Use sshKeys instead
	SshKey pulumi.StringPtrInput
	// The list of exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	SshKeys pulumi.StringArrayInput
	// The instance state (`running` or `stopped`; default: `running`).
	State pulumi.StringPtrInput
	// ❗ The getTemplate (ID) to use when creating the instance.
	TemplateId pulumi.StringPtrInput
	// The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
	Type pulumi.StringPtrInput
	// [cloud-init](https://cloudinit.readthedocs.io/) configuration.
	UserData pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (ComputeInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeInstanceState)(nil)).Elem()
}

type computeInstanceArgs struct {
	// ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
	AntiAffinityGroupIds []string `pulumi:"antiAffinityGroupIds"`
	// A list of exoscale*block*storage_volume (ID) to attach to the instance.
	BlockStorageVolumeIds []string `pulumi:"blockStorageVolumeIds"`
	// ❗ A deploy target ID.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
	DestroyProtected *bool `pulumi:"destroyProtected"`
	// The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
	DiskSize int `pulumi:"diskSize"`
	// A list of exoscale*elastic*ip (IDs) to attach to the instance.
	ElasticIpIds []string `pulumi:"elasticIpIds"`
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ipv6 *bool `pulumi:"ipv6"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The compute instance name.
	Name *string `pulumi:"name"`
	// Private network interfaces (may be specified multiple times). Structure is documented below.
	NetworkInterfaces []ComputeInstanceNetworkInterface `pulumi:"networkInterfaces"`
	// Whether the instance is private (no public IP addresses; default: false)
	Private *bool `pulumi:"private"`
	// Domain name for reverse DNS record.
	ReverseDns *string `pulumi:"reverseDns"`
	// A list of exoscale*security*group (IDs) to attach to the instance.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	//
	// Deprecated: Use sshKeys instead
	SshKey *string `pulumi:"sshKey"`
	// The list of exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	SshKeys []string `pulumi:"sshKeys"`
	// The instance state (`running` or `stopped`; default: `running`).
	State *string `pulumi:"state"`
	// ❗ The getTemplate (ID) to use when creating the instance.
	TemplateId string `pulumi:"templateId"`
	// The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
	Type string `pulumi:"type"`
	// [cloud-init](https://cloudinit.readthedocs.io/) configuration.
	UserData *string `pulumi:"userData"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a ComputeInstance resource.
type ComputeInstanceArgs struct {
	// ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
	AntiAffinityGroupIds pulumi.StringArrayInput
	// A list of exoscale*block*storage_volume (ID) to attach to the instance.
	BlockStorageVolumeIds pulumi.StringArrayInput
	// ❗ A deploy target ID.
	DeployTargetId pulumi.StringPtrInput
	// Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
	DestroyProtected pulumi.BoolPtrInput
	// The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
	DiskSize pulumi.IntInput
	// A list of exoscale*elastic*ip (IDs) to attach to the instance.
	ElasticIpIds pulumi.StringArrayInput
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ipv6 pulumi.BoolPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The compute instance name.
	Name pulumi.StringPtrInput
	// Private network interfaces (may be specified multiple times). Structure is documented below.
	NetworkInterfaces ComputeInstanceNetworkInterfaceArrayInput
	// Whether the instance is private (no public IP addresses; default: false)
	Private pulumi.BoolPtrInput
	// Domain name for reverse DNS record.
	ReverseDns pulumi.StringPtrInput
	// A list of exoscale*security*group (IDs) to attach to the instance.
	SecurityGroupIds pulumi.StringArrayInput
	// The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	//
	// Deprecated: Use sshKeys instead
	SshKey pulumi.StringPtrInput
	// The list of exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
	SshKeys pulumi.StringArrayInput
	// The instance state (`running` or `stopped`; default: `running`).
	State pulumi.StringPtrInput
	// ❗ The getTemplate (ID) to use when creating the instance.
	TemplateId pulumi.StringInput
	// The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
	Type pulumi.StringInput
	// [cloud-init](https://cloudinit.readthedocs.io/) configuration.
	UserData pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (ComputeInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeInstanceArgs)(nil)).Elem()
}

type ComputeInstanceInput interface {
	pulumi.Input

	ToComputeInstanceOutput() ComputeInstanceOutput
	ToComputeInstanceOutputWithContext(ctx context.Context) ComputeInstanceOutput
}

func (*ComputeInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstance)(nil)).Elem()
}

func (i *ComputeInstance) ToComputeInstanceOutput() ComputeInstanceOutput {
	return i.ToComputeInstanceOutputWithContext(context.Background())
}

func (i *ComputeInstance) ToComputeInstanceOutputWithContext(ctx context.Context) ComputeInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceOutput)
}

// ComputeInstanceArrayInput is an input type that accepts ComputeInstanceArray and ComputeInstanceArrayOutput values.
// You can construct a concrete instance of `ComputeInstanceArrayInput` via:
//
//	ComputeInstanceArray{ ComputeInstanceArgs{...} }
type ComputeInstanceArrayInput interface {
	pulumi.Input

	ToComputeInstanceArrayOutput() ComputeInstanceArrayOutput
	ToComputeInstanceArrayOutputWithContext(context.Context) ComputeInstanceArrayOutput
}

type ComputeInstanceArray []ComputeInstanceInput

func (ComputeInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeInstance)(nil)).Elem()
}

func (i ComputeInstanceArray) ToComputeInstanceArrayOutput() ComputeInstanceArrayOutput {
	return i.ToComputeInstanceArrayOutputWithContext(context.Background())
}

func (i ComputeInstanceArray) ToComputeInstanceArrayOutputWithContext(ctx context.Context) ComputeInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceArrayOutput)
}

// ComputeInstanceMapInput is an input type that accepts ComputeInstanceMap and ComputeInstanceMapOutput values.
// You can construct a concrete instance of `ComputeInstanceMapInput` via:
//
//	ComputeInstanceMap{ "key": ComputeInstanceArgs{...} }
type ComputeInstanceMapInput interface {
	pulumi.Input

	ToComputeInstanceMapOutput() ComputeInstanceMapOutput
	ToComputeInstanceMapOutputWithContext(context.Context) ComputeInstanceMapOutput
}

type ComputeInstanceMap map[string]ComputeInstanceInput

func (ComputeInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeInstance)(nil)).Elem()
}

func (i ComputeInstanceMap) ToComputeInstanceMapOutput() ComputeInstanceMapOutput {
	return i.ToComputeInstanceMapOutputWithContext(context.Background())
}

func (i ComputeInstanceMap) ToComputeInstanceMapOutputWithContext(ctx context.Context) ComputeInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceMapOutput)
}

type ComputeInstanceOutput struct{ *pulumi.OutputState }

func (ComputeInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstance)(nil)).Elem()
}

func (o ComputeInstanceOutput) ToComputeInstanceOutput() ComputeInstanceOutput {
	return o
}

func (o ComputeInstanceOutput) ToComputeInstanceOutputWithContext(ctx context.Context) ComputeInstanceOutput {
	return o
}

// ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
func (o ComputeInstanceOutput) AntiAffinityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringArrayOutput { return v.AntiAffinityGroupIds }).(pulumi.StringArrayOutput)
}

// A list of exoscale*block*storage_volume (ID) to attach to the instance.
func (o ComputeInstanceOutput) BlockStorageVolumeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringArrayOutput { return v.BlockStorageVolumeIds }).(pulumi.StringArrayOutput)
}

// The instance creation date.
func (o ComputeInstanceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// ❗ A deploy target ID.
func (o ComputeInstanceOutput) DeployTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringPtrOutput { return v.DeployTargetId }).(pulumi.StringPtrOutput)
}

// Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
func (o ComputeInstanceOutput) DestroyProtected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.BoolPtrOutput { return v.DestroyProtected }).(pulumi.BoolPtrOutput)
}

// The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
func (o ComputeInstanceOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// A list of exoscale*elastic*ip (IDs) to attach to the instance.
func (o ComputeInstanceOutput) ElasticIpIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringArrayOutput { return v.ElasticIpIds }).(pulumi.StringArrayOutput)
}

// Enable IPv6 on the instance (boolean; default: `false`).
func (o ComputeInstanceOutput) Ipv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.BoolPtrOutput { return v.Ipv6 }).(pulumi.BoolPtrOutput)
}

// The instance (main network interface) IPv6 address (if enabled).
func (o ComputeInstanceOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

// A map of key/value labels.
func (o ComputeInstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// MAC address
func (o ComputeInstanceOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// The compute instance name.
func (o ComputeInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private network interfaces (may be specified multiple times). Structure is documented below.
func (o ComputeInstanceOutput) NetworkInterfaces() ComputeInstanceNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v *ComputeInstance) ComputeInstanceNetworkInterfaceArrayOutput { return v.NetworkInterfaces }).(ComputeInstanceNetworkInterfaceArrayOutput)
}

// Whether the instance is private (no public IP addresses; default: false)
func (o ComputeInstanceOutput) Private() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.BoolPtrOutput { return v.Private }).(pulumi.BoolPtrOutput)
}

// A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument instead.
//
// Deprecated: Use the networkInterface block instead.
func (o ComputeInstanceOutput) PrivateNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringArrayOutput { return v.PrivateNetworkIds }).(pulumi.StringArrayOutput)
}

// The instance (main network interface) IPv4 address.
func (o ComputeInstanceOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.PublicIpAddress }).(pulumi.StringOutput)
}

// Domain name for reverse DNS record.
func (o ComputeInstanceOutput) ReverseDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringPtrOutput { return v.ReverseDns }).(pulumi.StringPtrOutput)
}

// A list of exoscale*security*group (IDs) to attach to the instance.
func (o ComputeInstanceOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
//
// Deprecated: Use sshKeys instead
func (o ComputeInstanceOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringPtrOutput { return v.SshKey }).(pulumi.StringPtrOutput)
}

// The list of exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
func (o ComputeInstanceOutput) SshKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringArrayOutput { return v.SshKeys }).(pulumi.StringArrayOutput)
}

// The instance state (`running` or `stopped`; default: `running`).
func (o ComputeInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// ❗ The getTemplate (ID) to use when creating the instance.
func (o ComputeInstanceOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

// The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
func (o ComputeInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// [cloud-init](https://cloudinit.readthedocs.io/) configuration.
func (o ComputeInstanceOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o ComputeInstanceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstance) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type ComputeInstanceArrayOutput struct{ *pulumi.OutputState }

func (ComputeInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeInstance)(nil)).Elem()
}

func (o ComputeInstanceArrayOutput) ToComputeInstanceArrayOutput() ComputeInstanceArrayOutput {
	return o
}

func (o ComputeInstanceArrayOutput) ToComputeInstanceArrayOutputWithContext(ctx context.Context) ComputeInstanceArrayOutput {
	return o
}

func (o ComputeInstanceArrayOutput) Index(i pulumi.IntInput) ComputeInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ComputeInstance {
		return vs[0].([]*ComputeInstance)[vs[1].(int)]
	}).(ComputeInstanceOutput)
}

type ComputeInstanceMapOutput struct{ *pulumi.OutputState }

func (ComputeInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeInstance)(nil)).Elem()
}

func (o ComputeInstanceMapOutput) ToComputeInstanceMapOutput() ComputeInstanceMapOutput {
	return o
}

func (o ComputeInstanceMapOutput) ToComputeInstanceMapOutputWithContext(ctx context.Context) ComputeInstanceMapOutput {
	return o
}

func (o ComputeInstanceMapOutput) MapIndex(k pulumi.StringInput) ComputeInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ComputeInstance {
		return vs[0].(map[string]*ComputeInstance)[vs[1].(string)]
	}).(ComputeInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceInput)(nil)).Elem(), &ComputeInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceArrayInput)(nil)).Elem(), ComputeInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceMapInput)(nil)).Elem(), ComputeInstanceMap{})
	pulumi.RegisterOutputType(ComputeInstanceOutput{})
	pulumi.RegisterOutputType(ComputeInstanceArrayOutput{})
	pulumi.RegisterOutputType(ComputeInstanceMapOutput{})
}
