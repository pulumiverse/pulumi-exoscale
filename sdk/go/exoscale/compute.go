// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use ComputeInstance instead.
type Compute struct {
	pulumi.CustomResourceState

	// A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinityGroups`).
	AffinityGroupIds pulumi.StringArrayOutput `pulumi:"affinityGroupIds"`
	// A list of anti-affinity groups (names; at creation time only; conflicts with `affinityGroupIds`).
	AffinityGroups pulumi.StringArrayOutput `pulumi:"affinityGroups"`
	// The instance disk size (GiB; at least `10`).
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen ("-") characters; it can be changed to any character during a later update. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Gateway     pulumi.StringOutput `pulumi:"gateway"`
	// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Enable IPv4 on the instance (only supported value is `true`).
	Ip4 pulumi.BoolPtrOutput `pulumi:"ip4"`
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ip6 pulumi.BoolPtrOutput `pulumi:"ip6"`
	// The instance (main network interface) IPv6 address (if enabled).
	Ip6Address pulumi.StringOutput `pulumi:"ip6Address"`
	Ip6Cidr    pulumi.StringOutput `pulumi:"ip6Cidr"`
	// The instance (main network interface) IPv4 address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The SSH keypair (name) to authorize in the instance.
	KeyPair pulumi.StringPtrOutput `pulumi:"keyPair"`
	// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
	Keyboard pulumi.StringPtrOutput `pulumi:"keyboard"`
	// (Deprecated) The instance hostname. Please use the `hostname` argument instead.
	//
	// Deprecated: use `hostname` attribute instead
	Name pulumi.StringOutput `pulumi:"name"`
	// The instance initial password and/or encrypted password.
	Password pulumi.StringOutput `pulumi:"password"`
	// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
	ReverseDns pulumi.StringPtrOutput `pulumi:"reverseDns"`
	// A list of security groups (IDs; conflicts with `securityGroups`).
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A list of security groups (names; conflicts with `securityGroupIds`).
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
	Size pulumi.StringPtrOutput `pulumi:"size"`
	// The instance state (`Running` or `Stopped`; default: `Running`)
	State pulumi.StringOutput `pulumi:"state"`
	// A map of tags (key/value). To remove all tags, set `tags = {}`.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `templateId` attribute instead.
	Template pulumi.StringOutput `pulumi:"template"`
	// The compute instance template (ID). Usage of the `getComputeTemplate` data source is recommended.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// cloud-init configuration (no need to base64-encode or gzip it as the provider will take care of it).
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// was the cloud-init configuration base64 encoded
	UserDataBase64 pulumi.BoolOutput `pulumi:"userDataBase64"`
	// The user to use to connect to the instance. If you've referenced a *custom template* in the resource, use the `getComputeTemplate` data source `username` attribute instead.
	//
	// Deprecated: broken, use `compute_template` data source `username` attribute
	Username pulumi.StringOutput `pulumi:"username"`
	// The Exoscale Zone name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewCompute registers a new resource with the given unique name, arguments, and options.
func NewCompute(ctx *pulumi.Context,
	name string, args *ComputeArgs, opts ...pulumi.ResourceOption) (*Compute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskSize == nil {
		return nil, errors.New("invalid value for required argument 'DiskSize'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Compute
	err := ctx.RegisterResource("exoscale:index/compute:Compute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompute gets an existing Compute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeState, opts ...pulumi.ResourceOption) (*Compute, error) {
	var resource Compute
	err := ctx.ReadResource("exoscale:index/compute:Compute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Compute resources.
type computeState struct {
	// A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinityGroups`).
	AffinityGroupIds []string `pulumi:"affinityGroupIds"`
	// A list of anti-affinity groups (names; at creation time only; conflicts with `affinityGroupIds`).
	AffinityGroups []string `pulumi:"affinityGroups"`
	// The instance disk size (GiB; at least `10`).
	DiskSize *int `pulumi:"diskSize"`
	// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen ("-") characters; it can be changed to any character during a later update. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically.
	DisplayName *string `pulumi:"displayName"`
	Gateway     *string `pulumi:"gateway"`
	// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
	Hostname *string `pulumi:"hostname"`
	// Enable IPv4 on the instance (only supported value is `true`).
	Ip4 *bool `pulumi:"ip4"`
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ip6 *bool `pulumi:"ip6"`
	// The instance (main network interface) IPv6 address (if enabled).
	Ip6Address *string `pulumi:"ip6Address"`
	Ip6Cidr    *string `pulumi:"ip6Cidr"`
	// The instance (main network interface) IPv4 address.
	IpAddress *string `pulumi:"ipAddress"`
	// The SSH keypair (name) to authorize in the instance.
	KeyPair *string `pulumi:"keyPair"`
	// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
	Keyboard *string `pulumi:"keyboard"`
	// (Deprecated) The instance hostname. Please use the `hostname` argument instead.
	//
	// Deprecated: use `hostname` attribute instead
	Name *string `pulumi:"name"`
	// The instance initial password and/or encrypted password.
	Password *string `pulumi:"password"`
	// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
	ReverseDns *string `pulumi:"reverseDns"`
	// A list of security groups (IDs; conflicts with `securityGroups`).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A list of security groups (names; conflicts with `securityGroupIds`).
	SecurityGroups []string `pulumi:"securityGroups"`
	// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
	Size *string `pulumi:"size"`
	// The instance state (`Running` or `Stopped`; default: `Running`)
	State *string `pulumi:"state"`
	// A map of tags (key/value). To remove all tags, set `tags = {}`.
	Tags map[string]string `pulumi:"tags"`
	// The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `templateId` attribute instead.
	Template *string `pulumi:"template"`
	// The compute instance template (ID). Usage of the `getComputeTemplate` data source is recommended.
	TemplateId *string `pulumi:"templateId"`
	// cloud-init configuration (no need to base64-encode or gzip it as the provider will take care of it).
	UserData *string `pulumi:"userData"`
	// was the cloud-init configuration base64 encoded
	UserDataBase64 *bool `pulumi:"userDataBase64"`
	// The user to use to connect to the instance. If you've referenced a *custom template* in the resource, use the `getComputeTemplate` data source `username` attribute instead.
	//
	// Deprecated: broken, use `compute_template` data source `username` attribute
	Username *string `pulumi:"username"`
	// The Exoscale Zone name.
	Zone *string `pulumi:"zone"`
}

type ComputeState struct {
	// A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinityGroups`).
	AffinityGroupIds pulumi.StringArrayInput
	// A list of anti-affinity groups (names; at creation time only; conflicts with `affinityGroupIds`).
	AffinityGroups pulumi.StringArrayInput
	// The instance disk size (GiB; at least `10`).
	DiskSize pulumi.IntPtrInput
	// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen ("-") characters; it can be changed to any character during a later update. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically.
	DisplayName pulumi.StringPtrInput
	Gateway     pulumi.StringPtrInput
	// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
	Hostname pulumi.StringPtrInput
	// Enable IPv4 on the instance (only supported value is `true`).
	Ip4 pulumi.BoolPtrInput
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ip6 pulumi.BoolPtrInput
	// The instance (main network interface) IPv6 address (if enabled).
	Ip6Address pulumi.StringPtrInput
	Ip6Cidr    pulumi.StringPtrInput
	// The instance (main network interface) IPv4 address.
	IpAddress pulumi.StringPtrInput
	// The SSH keypair (name) to authorize in the instance.
	KeyPair pulumi.StringPtrInput
	// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
	Keyboard pulumi.StringPtrInput
	// (Deprecated) The instance hostname. Please use the `hostname` argument instead.
	//
	// Deprecated: use `hostname` attribute instead
	Name pulumi.StringPtrInput
	// The instance initial password and/or encrypted password.
	Password pulumi.StringPtrInput
	// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
	ReverseDns pulumi.StringPtrInput
	// A list of security groups (IDs; conflicts with `securityGroups`).
	SecurityGroupIds pulumi.StringArrayInput
	// A list of security groups (names; conflicts with `securityGroupIds`).
	SecurityGroups pulumi.StringArrayInput
	// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
	Size pulumi.StringPtrInput
	// The instance state (`Running` or `Stopped`; default: `Running`)
	State pulumi.StringPtrInput
	// A map of tags (key/value). To remove all tags, set `tags = {}`.
	Tags pulumi.StringMapInput
	// The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `templateId` attribute instead.
	Template pulumi.StringPtrInput
	// The compute instance template (ID). Usage of the `getComputeTemplate` data source is recommended.
	TemplateId pulumi.StringPtrInput
	// cloud-init configuration (no need to base64-encode or gzip it as the provider will take care of it).
	UserData pulumi.StringPtrInput
	// was the cloud-init configuration base64 encoded
	UserDataBase64 pulumi.BoolPtrInput
	// The user to use to connect to the instance. If you've referenced a *custom template* in the resource, use the `getComputeTemplate` data source `username` attribute instead.
	//
	// Deprecated: broken, use `compute_template` data source `username` attribute
	Username pulumi.StringPtrInput
	// The Exoscale Zone name.
	Zone pulumi.StringPtrInput
}

func (ComputeState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeState)(nil)).Elem()
}

type computeArgs struct {
	// A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinityGroups`).
	AffinityGroupIds []string `pulumi:"affinityGroupIds"`
	// A list of anti-affinity groups (names; at creation time only; conflicts with `affinityGroupIds`).
	AffinityGroups []string `pulumi:"affinityGroups"`
	// The instance disk size (GiB; at least `10`).
	DiskSize int `pulumi:"diskSize"`
	// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen ("-") characters; it can be changed to any character during a later update. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically.
	DisplayName *string `pulumi:"displayName"`
	// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
	Hostname *string `pulumi:"hostname"`
	// Enable IPv4 on the instance (only supported value is `true`).
	Ip4 *bool `pulumi:"ip4"`
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ip6 *bool `pulumi:"ip6"`
	// The SSH keypair (name) to authorize in the instance.
	KeyPair *string `pulumi:"keyPair"`
	// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
	Keyboard *string `pulumi:"keyboard"`
	// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
	ReverseDns *string `pulumi:"reverseDns"`
	// A list of security groups (IDs; conflicts with `securityGroups`).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A list of security groups (names; conflicts with `securityGroupIds`).
	SecurityGroups []string `pulumi:"securityGroups"`
	// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
	Size *string `pulumi:"size"`
	// The instance state (`Running` or `Stopped`; default: `Running`)
	State *string `pulumi:"state"`
	// A map of tags (key/value). To remove all tags, set `tags = {}`.
	Tags map[string]string `pulumi:"tags"`
	// The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `templateId` attribute instead.
	Template *string `pulumi:"template"`
	// The compute instance template (ID). Usage of the `getComputeTemplate` data source is recommended.
	TemplateId *string `pulumi:"templateId"`
	// cloud-init configuration (no need to base64-encode or gzip it as the provider will take care of it).
	UserData *string `pulumi:"userData"`
	// The Exoscale Zone name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Compute resource.
type ComputeArgs struct {
	// A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinityGroups`).
	AffinityGroupIds pulumi.StringArrayInput
	// A list of anti-affinity groups (names; at creation time only; conflicts with `affinityGroupIds`).
	AffinityGroups pulumi.StringArrayInput
	// The instance disk size (GiB; at least `10`).
	DiskSize pulumi.IntInput
	// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen ("-") characters; it can be changed to any character during a later update. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically.
	DisplayName pulumi.StringPtrInput
	// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
	Hostname pulumi.StringPtrInput
	// Enable IPv4 on the instance (only supported value is `true`).
	Ip4 pulumi.BoolPtrInput
	// Enable IPv6 on the instance (boolean; default: `false`).
	Ip6 pulumi.BoolPtrInput
	// The SSH keypair (name) to authorize in the instance.
	KeyPair pulumi.StringPtrInput
	// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
	Keyboard pulumi.StringPtrInput
	// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
	ReverseDns pulumi.StringPtrInput
	// A list of security groups (IDs; conflicts with `securityGroups`).
	SecurityGroupIds pulumi.StringArrayInput
	// A list of security groups (names; conflicts with `securityGroupIds`).
	SecurityGroups pulumi.StringArrayInput
	// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
	Size pulumi.StringPtrInput
	// The instance state (`Running` or `Stopped`; default: `Running`)
	State pulumi.StringPtrInput
	// A map of tags (key/value). To remove all tags, set `tags = {}`.
	Tags pulumi.StringMapInput
	// The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `templateId` attribute instead.
	Template pulumi.StringPtrInput
	// The compute instance template (ID). Usage of the `getComputeTemplate` data source is recommended.
	TemplateId pulumi.StringPtrInput
	// cloud-init configuration (no need to base64-encode or gzip it as the provider will take care of it).
	UserData pulumi.StringPtrInput
	// The Exoscale Zone name.
	Zone pulumi.StringInput
}

func (ComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeArgs)(nil)).Elem()
}

type ComputeInput interface {
	pulumi.Input

	ToComputeOutput() ComputeOutput
	ToComputeOutputWithContext(ctx context.Context) ComputeOutput
}

func (*Compute) ElementType() reflect.Type {
	return reflect.TypeOf((**Compute)(nil)).Elem()
}

func (i *Compute) ToComputeOutput() ComputeOutput {
	return i.ToComputeOutputWithContext(context.Background())
}

func (i *Compute) ToComputeOutputWithContext(ctx context.Context) ComputeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeOutput)
}

// ComputeArrayInput is an input type that accepts ComputeArray and ComputeArrayOutput values.
// You can construct a concrete instance of `ComputeArrayInput` via:
//
//	ComputeArray{ ComputeArgs{...} }
type ComputeArrayInput interface {
	pulumi.Input

	ToComputeArrayOutput() ComputeArrayOutput
	ToComputeArrayOutputWithContext(context.Context) ComputeArrayOutput
}

type ComputeArray []ComputeInput

func (ComputeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Compute)(nil)).Elem()
}

func (i ComputeArray) ToComputeArrayOutput() ComputeArrayOutput {
	return i.ToComputeArrayOutputWithContext(context.Background())
}

func (i ComputeArray) ToComputeArrayOutputWithContext(ctx context.Context) ComputeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeArrayOutput)
}

// ComputeMapInput is an input type that accepts ComputeMap and ComputeMapOutput values.
// You can construct a concrete instance of `ComputeMapInput` via:
//
//	ComputeMap{ "key": ComputeArgs{...} }
type ComputeMapInput interface {
	pulumi.Input

	ToComputeMapOutput() ComputeMapOutput
	ToComputeMapOutputWithContext(context.Context) ComputeMapOutput
}

type ComputeMap map[string]ComputeInput

func (ComputeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Compute)(nil)).Elem()
}

func (i ComputeMap) ToComputeMapOutput() ComputeMapOutput {
	return i.ToComputeMapOutputWithContext(context.Background())
}

func (i ComputeMap) ToComputeMapOutputWithContext(ctx context.Context) ComputeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeMapOutput)
}

type ComputeOutput struct{ *pulumi.OutputState }

func (ComputeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Compute)(nil)).Elem()
}

func (o ComputeOutput) ToComputeOutput() ComputeOutput {
	return o
}

func (o ComputeOutput) ToComputeOutputWithContext(ctx context.Context) ComputeOutput {
	return o
}

// A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinityGroups`).
func (o ComputeOutput) AffinityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringArrayOutput { return v.AffinityGroupIds }).(pulumi.StringArrayOutput)
}

// A list of anti-affinity groups (names; at creation time only; conflicts with `affinityGroupIds`).
func (o ComputeOutput) AffinityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringArrayOutput { return v.AffinityGroups }).(pulumi.StringArrayOutput)
}

// The instance disk size (GiB; at least `10`).
func (o ComputeOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Compute) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen ("-") characters; it can be changed to any character during a later update. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically.
func (o ComputeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ComputeOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `displayName` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
func (o ComputeOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Enable IPv4 on the instance (only supported value is `true`).
func (o ComputeOutput) Ip4() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.BoolPtrOutput { return v.Ip4 }).(pulumi.BoolPtrOutput)
}

// Enable IPv6 on the instance (boolean; default: `false`).
func (o ComputeOutput) Ip6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.BoolPtrOutput { return v.Ip6 }).(pulumi.BoolPtrOutput)
}

// The instance (main network interface) IPv6 address (if enabled).
func (o ComputeOutput) Ip6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Ip6Address }).(pulumi.StringOutput)
}

func (o ComputeOutput) Ip6Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Ip6Cidr }).(pulumi.StringOutput)
}

// The instance (main network interface) IPv4 address.
func (o ComputeOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The SSH keypair (name) to authorize in the instance.
func (o ComputeOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringPtrOutput { return v.KeyPair }).(pulumi.StringPtrOutput)
}

// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
func (o ComputeOutput) Keyboard() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringPtrOutput { return v.Keyboard }).(pulumi.StringPtrOutput)
}

// (Deprecated) The instance hostname. Please use the `hostname` argument instead.
//
// Deprecated: use `hostname` attribute instead
func (o ComputeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The instance initial password and/or encrypted password.
func (o ComputeOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
func (o ComputeOutput) ReverseDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringPtrOutput { return v.ReverseDns }).(pulumi.StringPtrOutput)
}

// A list of security groups (IDs; conflicts with `securityGroups`).
func (o ComputeOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// A list of security groups (names; conflicts with `securityGroupIds`).
func (o ComputeOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
func (o ComputeOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringPtrOutput { return v.Size }).(pulumi.StringPtrOutput)
}

// The instance state (`Running` or `Stopped`; default: `Running`)
func (o ComputeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A map of tags (key/value). To remove all tags, set `tags = {}`.
func (o ComputeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `templateId` attribute instead.
func (o ComputeOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Template }).(pulumi.StringOutput)
}

// The compute instance template (ID). Usage of the `getComputeTemplate` data source is recommended.
func (o ComputeOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

// cloud-init configuration (no need to base64-encode or gzip it as the provider will take care of it).
func (o ComputeOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// was the cloud-init configuration base64 encoded
func (o ComputeOutput) UserDataBase64() pulumi.BoolOutput {
	return o.ApplyT(func(v *Compute) pulumi.BoolOutput { return v.UserDataBase64 }).(pulumi.BoolOutput)
}

// The user to use to connect to the instance. If you've referenced a *custom template* in the resource, use the `getComputeTemplate` data source `username` attribute instead.
//
// Deprecated: broken, use `compute_template` data source `username` attribute
func (o ComputeOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// The Exoscale Zone name.
func (o ComputeOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type ComputeArrayOutput struct{ *pulumi.OutputState }

func (ComputeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Compute)(nil)).Elem()
}

func (o ComputeArrayOutput) ToComputeArrayOutput() ComputeArrayOutput {
	return o
}

func (o ComputeArrayOutput) ToComputeArrayOutputWithContext(ctx context.Context) ComputeArrayOutput {
	return o
}

func (o ComputeArrayOutput) Index(i pulumi.IntInput) ComputeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Compute {
		return vs[0].([]*Compute)[vs[1].(int)]
	}).(ComputeOutput)
}

type ComputeMapOutput struct{ *pulumi.OutputState }

func (ComputeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Compute)(nil)).Elem()
}

func (o ComputeMapOutput) ToComputeMapOutput() ComputeMapOutput {
	return o
}

func (o ComputeMapOutput) ToComputeMapOutputWithContext(ctx context.Context) ComputeMapOutput {
	return o
}

func (o ComputeMapOutput) MapIndex(k pulumi.StringInput) ComputeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Compute {
		return vs[0].(map[string]*Compute)[vs[1].(string)]
	}).(ComputeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInput)(nil)).Elem(), &Compute{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeArrayInput)(nil)).Elem(), ComputeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeMapInput)(nil)).Elem(), ComputeMap{})
	pulumi.RegisterOutputType(ComputeOutput{})
	pulumi.RegisterOutputType(ComputeArrayOutput{})
	pulumi.RegisterOutputType(ComputeMapOutput{})
}
