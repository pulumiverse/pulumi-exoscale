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

type SksKubeconfig struct {
	pulumi.CustomResourceState

	// ❗ The parent exoscale*sks*cluster ID.
	ClusterId           pulumi.StringOutput `pulumi:"clusterId"`
	EarlyRenewalSeconds pulumi.IntPtrOutput `pulumi:"earlyRenewalSeconds"`
	// ❗ Group names in the generated Kubeconfig. The certificate present in the Kubeconfig will have these roles set in the Organization field.
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// The generated Kubeconfig (YAML content).
	Kubeconfig      pulumi.StringOutput `pulumi:"kubeconfig"`
	ReadyForRenewal pulumi.BoolOutput   `pulumi:"readyForRenewal"`
	// ❗ The Time-to-Live of the Kubeconfig, after which it will expire / become invalid (seconds; default: 2592000 = 30 days).
	TtlSeconds pulumi.Float64PtrOutput `pulumi:"ttlSeconds"`
	// ❗ User name in the generated Kubeconfig. The certificate present in the Kubeconfig will also have this name set for the CN field.
	User pulumi.StringOutput `pulumi:"user"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewSksKubeconfig registers a new resource with the given unique name, arguments, and options.
func NewSksKubeconfig(ctx *pulumi.Context,
	name string, args *SksKubeconfigArgs, opts ...pulumi.ResourceOption) (*SksKubeconfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Groups == nil {
		return nil, errors.New("invalid value for required argument 'Groups'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"kubeconfig",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SksKubeconfig
	err := ctx.RegisterResource("exoscale:index/sksKubeconfig:SksKubeconfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSksKubeconfig gets an existing SksKubeconfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSksKubeconfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SksKubeconfigState, opts ...pulumi.ResourceOption) (*SksKubeconfig, error) {
	var resource SksKubeconfig
	err := ctx.ReadResource("exoscale:index/sksKubeconfig:SksKubeconfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SksKubeconfig resources.
type sksKubeconfigState struct {
	// ❗ The parent exoscale*sks*cluster ID.
	ClusterId           *string `pulumi:"clusterId"`
	EarlyRenewalSeconds *int    `pulumi:"earlyRenewalSeconds"`
	// ❗ Group names in the generated Kubeconfig. The certificate present in the Kubeconfig will have these roles set in the Organization field.
	Groups []string `pulumi:"groups"`
	// The generated Kubeconfig (YAML content).
	Kubeconfig      *string `pulumi:"kubeconfig"`
	ReadyForRenewal *bool   `pulumi:"readyForRenewal"`
	// ❗ The Time-to-Live of the Kubeconfig, after which it will expire / become invalid (seconds; default: 2592000 = 30 days).
	TtlSeconds *float64 `pulumi:"ttlSeconds"`
	// ❗ User name in the generated Kubeconfig. The certificate present in the Kubeconfig will also have this name set for the CN field.
	User *string `pulumi:"user"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type SksKubeconfigState struct {
	// ❗ The parent exoscale*sks*cluster ID.
	ClusterId           pulumi.StringPtrInput
	EarlyRenewalSeconds pulumi.IntPtrInput
	// ❗ Group names in the generated Kubeconfig. The certificate present in the Kubeconfig will have these roles set in the Organization field.
	Groups pulumi.StringArrayInput
	// The generated Kubeconfig (YAML content).
	Kubeconfig      pulumi.StringPtrInput
	ReadyForRenewal pulumi.BoolPtrInput
	// ❗ The Time-to-Live of the Kubeconfig, after which it will expire / become invalid (seconds; default: 2592000 = 30 days).
	TtlSeconds pulumi.Float64PtrInput
	// ❗ User name in the generated Kubeconfig. The certificate present in the Kubeconfig will also have this name set for the CN field.
	User pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (SksKubeconfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*sksKubeconfigState)(nil)).Elem()
}

type sksKubeconfigArgs struct {
	// ❗ The parent exoscale*sks*cluster ID.
	ClusterId           string `pulumi:"clusterId"`
	EarlyRenewalSeconds *int   `pulumi:"earlyRenewalSeconds"`
	// ❗ Group names in the generated Kubeconfig. The certificate present in the Kubeconfig will have these roles set in the Organization field.
	Groups []string `pulumi:"groups"`
	// ❗ The Time-to-Live of the Kubeconfig, after which it will expire / become invalid (seconds; default: 2592000 = 30 days).
	TtlSeconds *float64 `pulumi:"ttlSeconds"`
	// ❗ User name in the generated Kubeconfig. The certificate present in the Kubeconfig will also have this name set for the CN field.
	User string `pulumi:"user"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a SksKubeconfig resource.
type SksKubeconfigArgs struct {
	// ❗ The parent exoscale*sks*cluster ID.
	ClusterId           pulumi.StringInput
	EarlyRenewalSeconds pulumi.IntPtrInput
	// ❗ Group names in the generated Kubeconfig. The certificate present in the Kubeconfig will have these roles set in the Organization field.
	Groups pulumi.StringArrayInput
	// ❗ The Time-to-Live of the Kubeconfig, after which it will expire / become invalid (seconds; default: 2592000 = 30 days).
	TtlSeconds pulumi.Float64PtrInput
	// ❗ User name in the generated Kubeconfig. The certificate present in the Kubeconfig will also have this name set for the CN field.
	User pulumi.StringInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (SksKubeconfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sksKubeconfigArgs)(nil)).Elem()
}

type SksKubeconfigInput interface {
	pulumi.Input

	ToSksKubeconfigOutput() SksKubeconfigOutput
	ToSksKubeconfigOutputWithContext(ctx context.Context) SksKubeconfigOutput
}

func (*SksKubeconfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SksKubeconfig)(nil)).Elem()
}

func (i *SksKubeconfig) ToSksKubeconfigOutput() SksKubeconfigOutput {
	return i.ToSksKubeconfigOutputWithContext(context.Background())
}

func (i *SksKubeconfig) ToSksKubeconfigOutputWithContext(ctx context.Context) SksKubeconfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SksKubeconfigOutput)
}

// SksKubeconfigArrayInput is an input type that accepts SksKubeconfigArray and SksKubeconfigArrayOutput values.
// You can construct a concrete instance of `SksKubeconfigArrayInput` via:
//
//	SksKubeconfigArray{ SksKubeconfigArgs{...} }
type SksKubeconfigArrayInput interface {
	pulumi.Input

	ToSksKubeconfigArrayOutput() SksKubeconfigArrayOutput
	ToSksKubeconfigArrayOutputWithContext(context.Context) SksKubeconfigArrayOutput
}

type SksKubeconfigArray []SksKubeconfigInput

func (SksKubeconfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SksKubeconfig)(nil)).Elem()
}

func (i SksKubeconfigArray) ToSksKubeconfigArrayOutput() SksKubeconfigArrayOutput {
	return i.ToSksKubeconfigArrayOutputWithContext(context.Background())
}

func (i SksKubeconfigArray) ToSksKubeconfigArrayOutputWithContext(ctx context.Context) SksKubeconfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SksKubeconfigArrayOutput)
}

// SksKubeconfigMapInput is an input type that accepts SksKubeconfigMap and SksKubeconfigMapOutput values.
// You can construct a concrete instance of `SksKubeconfigMapInput` via:
//
//	SksKubeconfigMap{ "key": SksKubeconfigArgs{...} }
type SksKubeconfigMapInput interface {
	pulumi.Input

	ToSksKubeconfigMapOutput() SksKubeconfigMapOutput
	ToSksKubeconfigMapOutputWithContext(context.Context) SksKubeconfigMapOutput
}

type SksKubeconfigMap map[string]SksKubeconfigInput

func (SksKubeconfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SksKubeconfig)(nil)).Elem()
}

func (i SksKubeconfigMap) ToSksKubeconfigMapOutput() SksKubeconfigMapOutput {
	return i.ToSksKubeconfigMapOutputWithContext(context.Background())
}

func (i SksKubeconfigMap) ToSksKubeconfigMapOutputWithContext(ctx context.Context) SksKubeconfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SksKubeconfigMapOutput)
}

type SksKubeconfigOutput struct{ *pulumi.OutputState }

func (SksKubeconfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SksKubeconfig)(nil)).Elem()
}

func (o SksKubeconfigOutput) ToSksKubeconfigOutput() SksKubeconfigOutput {
	return o
}

func (o SksKubeconfigOutput) ToSksKubeconfigOutputWithContext(ctx context.Context) SksKubeconfigOutput {
	return o
}

// ❗ The parent exoscale*sks*cluster ID.
func (o SksKubeconfigOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *SksKubeconfig) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o SksKubeconfigOutput) EarlyRenewalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SksKubeconfig) pulumi.IntPtrOutput { return v.EarlyRenewalSeconds }).(pulumi.IntPtrOutput)
}

// ❗ Group names in the generated Kubeconfig. The certificate present in the Kubeconfig will have these roles set in the Organization field.
func (o SksKubeconfigOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SksKubeconfig) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

// The generated Kubeconfig (YAML content).
func (o SksKubeconfigOutput) Kubeconfig() pulumi.StringOutput {
	return o.ApplyT(func(v *SksKubeconfig) pulumi.StringOutput { return v.Kubeconfig }).(pulumi.StringOutput)
}

func (o SksKubeconfigOutput) ReadyForRenewal() pulumi.BoolOutput {
	return o.ApplyT(func(v *SksKubeconfig) pulumi.BoolOutput { return v.ReadyForRenewal }).(pulumi.BoolOutput)
}

// ❗ The Time-to-Live of the Kubeconfig, after which it will expire / become invalid (seconds; default: 2592000 = 30 days).
func (o SksKubeconfigOutput) TtlSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SksKubeconfig) pulumi.Float64PtrOutput { return v.TtlSeconds }).(pulumi.Float64PtrOutput)
}

// ❗ User name in the generated Kubeconfig. The certificate present in the Kubeconfig will also have this name set for the CN field.
func (o SksKubeconfigOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *SksKubeconfig) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o SksKubeconfigOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *SksKubeconfig) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type SksKubeconfigArrayOutput struct{ *pulumi.OutputState }

func (SksKubeconfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SksKubeconfig)(nil)).Elem()
}

func (o SksKubeconfigArrayOutput) ToSksKubeconfigArrayOutput() SksKubeconfigArrayOutput {
	return o
}

func (o SksKubeconfigArrayOutput) ToSksKubeconfigArrayOutputWithContext(ctx context.Context) SksKubeconfigArrayOutput {
	return o
}

func (o SksKubeconfigArrayOutput) Index(i pulumi.IntInput) SksKubeconfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SksKubeconfig {
		return vs[0].([]*SksKubeconfig)[vs[1].(int)]
	}).(SksKubeconfigOutput)
}

type SksKubeconfigMapOutput struct{ *pulumi.OutputState }

func (SksKubeconfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SksKubeconfig)(nil)).Elem()
}

func (o SksKubeconfigMapOutput) ToSksKubeconfigMapOutput() SksKubeconfigMapOutput {
	return o
}

func (o SksKubeconfigMapOutput) ToSksKubeconfigMapOutputWithContext(ctx context.Context) SksKubeconfigMapOutput {
	return o
}

func (o SksKubeconfigMapOutput) MapIndex(k pulumi.StringInput) SksKubeconfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SksKubeconfig {
		return vs[0].(map[string]*SksKubeconfig)[vs[1].(string)]
	}).(SksKubeconfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SksKubeconfigInput)(nil)).Elem(), &SksKubeconfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*SksKubeconfigArrayInput)(nil)).Elem(), SksKubeconfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SksKubeconfigMapInput)(nil)).Elem(), SksKubeconfigMap{})
	pulumi.RegisterOutputType(SksKubeconfigOutput{})
	pulumi.RegisterOutputType(SksKubeconfigArrayOutput{})
	pulumi.RegisterOutputType(SksKubeconfigMapOutput{})
}
