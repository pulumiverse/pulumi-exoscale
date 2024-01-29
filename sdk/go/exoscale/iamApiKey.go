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

type IamApiKey struct {
	pulumi.CustomResourceState

	// The IAM API Key to match.
	Key pulumi.StringOutput `pulumi:"key"`
	// ❗ IAM API Key name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ❗ IAM API role ID.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// Secret for the IAM API Key.
	Secret   pulumi.StringOutput        `pulumi:"secret"`
	Timeouts IamApiKeyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewIamApiKey registers a new resource with the given unique name, arguments, and options.
func NewIamApiKey(ctx *pulumi.Context,
	name string, args *IamApiKeyArgs, opts ...pulumi.ResourceOption) (*IamApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamApiKey
	err := ctx.RegisterResource("exoscale:index/iamApiKey:IamApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamApiKey gets an existing IamApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamApiKeyState, opts ...pulumi.ResourceOption) (*IamApiKey, error) {
	var resource IamApiKey
	err := ctx.ReadResource("exoscale:index/iamApiKey:IamApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamApiKey resources.
type iamApiKeyState struct {
	// The IAM API Key to match.
	Key *string `pulumi:"key"`
	// ❗ IAM API Key name.
	Name *string `pulumi:"name"`
	// ❗ IAM API role ID.
	RoleId *string `pulumi:"roleId"`
	// Secret for the IAM API Key.
	Secret   *string            `pulumi:"secret"`
	Timeouts *IamApiKeyTimeouts `pulumi:"timeouts"`
}

type IamApiKeyState struct {
	// The IAM API Key to match.
	Key pulumi.StringPtrInput
	// ❗ IAM API Key name.
	Name pulumi.StringPtrInput
	// ❗ IAM API role ID.
	RoleId pulumi.StringPtrInput
	// Secret for the IAM API Key.
	Secret   pulumi.StringPtrInput
	Timeouts IamApiKeyTimeoutsPtrInput
}

func (IamApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamApiKeyState)(nil)).Elem()
}

type iamApiKeyArgs struct {
	// ❗ IAM API Key name.
	Name *string `pulumi:"name"`
	// ❗ IAM API role ID.
	RoleId   string             `pulumi:"roleId"`
	Timeouts *IamApiKeyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a IamApiKey resource.
type IamApiKeyArgs struct {
	// ❗ IAM API Key name.
	Name pulumi.StringPtrInput
	// ❗ IAM API role ID.
	RoleId   pulumi.StringInput
	Timeouts IamApiKeyTimeoutsPtrInput
}

func (IamApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamApiKeyArgs)(nil)).Elem()
}

type IamApiKeyInput interface {
	pulumi.Input

	ToIamApiKeyOutput() IamApiKeyOutput
	ToIamApiKeyOutputWithContext(ctx context.Context) IamApiKeyOutput
}

func (*IamApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((**IamApiKey)(nil)).Elem()
}

func (i *IamApiKey) ToIamApiKeyOutput() IamApiKeyOutput {
	return i.ToIamApiKeyOutputWithContext(context.Background())
}

func (i *IamApiKey) ToIamApiKeyOutputWithContext(ctx context.Context) IamApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamApiKeyOutput)
}

// IamApiKeyArrayInput is an input type that accepts IamApiKeyArray and IamApiKeyArrayOutput values.
// You can construct a concrete instance of `IamApiKeyArrayInput` via:
//
//	IamApiKeyArray{ IamApiKeyArgs{...} }
type IamApiKeyArrayInput interface {
	pulumi.Input

	ToIamApiKeyArrayOutput() IamApiKeyArrayOutput
	ToIamApiKeyArrayOutputWithContext(context.Context) IamApiKeyArrayOutput
}

type IamApiKeyArray []IamApiKeyInput

func (IamApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamApiKey)(nil)).Elem()
}

func (i IamApiKeyArray) ToIamApiKeyArrayOutput() IamApiKeyArrayOutput {
	return i.ToIamApiKeyArrayOutputWithContext(context.Background())
}

func (i IamApiKeyArray) ToIamApiKeyArrayOutputWithContext(ctx context.Context) IamApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamApiKeyArrayOutput)
}

// IamApiKeyMapInput is an input type that accepts IamApiKeyMap and IamApiKeyMapOutput values.
// You can construct a concrete instance of `IamApiKeyMapInput` via:
//
//	IamApiKeyMap{ "key": IamApiKeyArgs{...} }
type IamApiKeyMapInput interface {
	pulumi.Input

	ToIamApiKeyMapOutput() IamApiKeyMapOutput
	ToIamApiKeyMapOutputWithContext(context.Context) IamApiKeyMapOutput
}

type IamApiKeyMap map[string]IamApiKeyInput

func (IamApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamApiKey)(nil)).Elem()
}

func (i IamApiKeyMap) ToIamApiKeyMapOutput() IamApiKeyMapOutput {
	return i.ToIamApiKeyMapOutputWithContext(context.Background())
}

func (i IamApiKeyMap) ToIamApiKeyMapOutputWithContext(ctx context.Context) IamApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamApiKeyMapOutput)
}

type IamApiKeyOutput struct{ *pulumi.OutputState }

func (IamApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamApiKey)(nil)).Elem()
}

func (o IamApiKeyOutput) ToIamApiKeyOutput() IamApiKeyOutput {
	return o
}

func (o IamApiKeyOutput) ToIamApiKeyOutputWithContext(ctx context.Context) IamApiKeyOutput {
	return o
}

// The IAM API Key to match.
func (o IamApiKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *IamApiKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// ❗ IAM API Key name.
func (o IamApiKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IamApiKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ❗ IAM API role ID.
func (o IamApiKeyOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamApiKey) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// Secret for the IAM API Key.
func (o IamApiKeyOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *IamApiKey) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

func (o IamApiKeyOutput) Timeouts() IamApiKeyTimeoutsPtrOutput {
	return o.ApplyT(func(v *IamApiKey) IamApiKeyTimeoutsPtrOutput { return v.Timeouts }).(IamApiKeyTimeoutsPtrOutput)
}

type IamApiKeyArrayOutput struct{ *pulumi.OutputState }

func (IamApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamApiKey)(nil)).Elem()
}

func (o IamApiKeyArrayOutput) ToIamApiKeyArrayOutput() IamApiKeyArrayOutput {
	return o
}

func (o IamApiKeyArrayOutput) ToIamApiKeyArrayOutputWithContext(ctx context.Context) IamApiKeyArrayOutput {
	return o
}

func (o IamApiKeyArrayOutput) Index(i pulumi.IntInput) IamApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamApiKey {
		return vs[0].([]*IamApiKey)[vs[1].(int)]
	}).(IamApiKeyOutput)
}

type IamApiKeyMapOutput struct{ *pulumi.OutputState }

func (IamApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamApiKey)(nil)).Elem()
}

func (o IamApiKeyMapOutput) ToIamApiKeyMapOutput() IamApiKeyMapOutput {
	return o
}

func (o IamApiKeyMapOutput) ToIamApiKeyMapOutputWithContext(ctx context.Context) IamApiKeyMapOutput {
	return o
}

func (o IamApiKeyMapOutput) MapIndex(k pulumi.StringInput) IamApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamApiKey {
		return vs[0].(map[string]*IamApiKey)[vs[1].(string)]
	}).(IamApiKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamApiKeyInput)(nil)).Elem(), &IamApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamApiKeyArrayInput)(nil)).Elem(), IamApiKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamApiKeyMapInput)(nil)).Elem(), IamApiKeyMap{})
	pulumi.RegisterOutputType(IamApiKeyOutput{})
	pulumi.RegisterOutputType(IamApiKeyArrayOutput{})
	pulumi.RegisterOutputType(IamApiKeyMapOutput{})
}
