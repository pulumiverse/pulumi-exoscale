// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SSHKeypair struct {
	pulumi.CustomResourceState

	// The SSH keypair unique identifier.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The SSH keypair name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The SSH *private* key generated if no public key was provided.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// A SSH *public* key that will be authorized in compute instances. If not provided, an SSH keypair (including the *private* key) is generated and saved locally (see the `privateKey` attribute).
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
}

// NewSSHKeypair registers a new resource with the given unique name, arguments, and options.
func NewSSHKeypair(ctx *pulumi.Context,
	name string, args *SSHKeypairArgs, opts ...pulumi.ResourceOption) (*SSHKeypair, error) {
	if args == nil {
		args = &SSHKeypairArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SSHKeypair
	err := ctx.RegisterResource("exoscale:index/sSHKeypair:SSHKeypair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSSHKeypair gets an existing SSHKeypair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSSHKeypair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SSHKeypairState, opts ...pulumi.ResourceOption) (*SSHKeypair, error) {
	var resource SSHKeypair
	err := ctx.ReadResource("exoscale:index/sSHKeypair:SSHKeypair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SSHKeypair resources.
type sshkeypairState struct {
	// The SSH keypair unique identifier.
	Fingerprint *string `pulumi:"fingerprint"`
	// The SSH keypair name.
	Name *string `pulumi:"name"`
	// The SSH *private* key generated if no public key was provided.
	PrivateKey *string `pulumi:"privateKey"`
	// A SSH *public* key that will be authorized in compute instances. If not provided, an SSH keypair (including the *private* key) is generated and saved locally (see the `privateKey` attribute).
	PublicKey *string `pulumi:"publicKey"`
}

type SSHKeypairState struct {
	// The SSH keypair unique identifier.
	Fingerprint pulumi.StringPtrInput
	// The SSH keypair name.
	Name pulumi.StringPtrInput
	// The SSH *private* key generated if no public key was provided.
	PrivateKey pulumi.StringPtrInput
	// A SSH *public* key that will be authorized in compute instances. If not provided, an SSH keypair (including the *private* key) is generated and saved locally (see the `privateKey` attribute).
	PublicKey pulumi.StringPtrInput
}

func (SSHKeypairState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshkeypairState)(nil)).Elem()
}

type sshkeypairArgs struct {
	// The SSH keypair name.
	Name *string `pulumi:"name"`
	// A SSH *public* key that will be authorized in compute instances. If not provided, an SSH keypair (including the *private* key) is generated and saved locally (see the `privateKey` attribute).
	PublicKey *string `pulumi:"publicKey"`
}

// The set of arguments for constructing a SSHKeypair resource.
type SSHKeypairArgs struct {
	// The SSH keypair name.
	Name pulumi.StringPtrInput
	// A SSH *public* key that will be authorized in compute instances. If not provided, an SSH keypair (including the *private* key) is generated and saved locally (see the `privateKey` attribute).
	PublicKey pulumi.StringPtrInput
}

func (SSHKeypairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshkeypairArgs)(nil)).Elem()
}

type SSHKeypairInput interface {
	pulumi.Input

	ToSSHKeypairOutput() SSHKeypairOutput
	ToSSHKeypairOutputWithContext(ctx context.Context) SSHKeypairOutput
}

func (*SSHKeypair) ElementType() reflect.Type {
	return reflect.TypeOf((**SSHKeypair)(nil)).Elem()
}

func (i *SSHKeypair) ToSSHKeypairOutput() SSHKeypairOutput {
	return i.ToSSHKeypairOutputWithContext(context.Background())
}

func (i *SSHKeypair) ToSSHKeypairOutputWithContext(ctx context.Context) SSHKeypairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeypairOutput)
}

// SSHKeypairArrayInput is an input type that accepts SSHKeypairArray and SSHKeypairArrayOutput values.
// You can construct a concrete instance of `SSHKeypairArrayInput` via:
//
//	SSHKeypairArray{ SSHKeypairArgs{...} }
type SSHKeypairArrayInput interface {
	pulumi.Input

	ToSSHKeypairArrayOutput() SSHKeypairArrayOutput
	ToSSHKeypairArrayOutputWithContext(context.Context) SSHKeypairArrayOutput
}

type SSHKeypairArray []SSHKeypairInput

func (SSHKeypairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSHKeypair)(nil)).Elem()
}

func (i SSHKeypairArray) ToSSHKeypairArrayOutput() SSHKeypairArrayOutput {
	return i.ToSSHKeypairArrayOutputWithContext(context.Background())
}

func (i SSHKeypairArray) ToSSHKeypairArrayOutputWithContext(ctx context.Context) SSHKeypairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeypairArrayOutput)
}

// SSHKeypairMapInput is an input type that accepts SSHKeypairMap and SSHKeypairMapOutput values.
// You can construct a concrete instance of `SSHKeypairMapInput` via:
//
//	SSHKeypairMap{ "key": SSHKeypairArgs{...} }
type SSHKeypairMapInput interface {
	pulumi.Input

	ToSSHKeypairMapOutput() SSHKeypairMapOutput
	ToSSHKeypairMapOutputWithContext(context.Context) SSHKeypairMapOutput
}

type SSHKeypairMap map[string]SSHKeypairInput

func (SSHKeypairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSHKeypair)(nil)).Elem()
}

func (i SSHKeypairMap) ToSSHKeypairMapOutput() SSHKeypairMapOutput {
	return i.ToSSHKeypairMapOutputWithContext(context.Background())
}

func (i SSHKeypairMap) ToSSHKeypairMapOutputWithContext(ctx context.Context) SSHKeypairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeypairMapOutput)
}

type SSHKeypairOutput struct{ *pulumi.OutputState }

func (SSHKeypairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SSHKeypair)(nil)).Elem()
}

func (o SSHKeypairOutput) ToSSHKeypairOutput() SSHKeypairOutput {
	return o
}

func (o SSHKeypairOutput) ToSSHKeypairOutputWithContext(ctx context.Context) SSHKeypairOutput {
	return o
}

// The SSH keypair unique identifier.
func (o SSHKeypairOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKeypair) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The SSH keypair name.
func (o SSHKeypairOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKeypair) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The SSH *private* key generated if no public key was provided.
func (o SSHKeypairOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKeypair) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// A SSH *public* key that will be authorized in compute instances. If not provided, an SSH keypair (including the *private* key) is generated and saved locally (see the `privateKey` attribute).
func (o SSHKeypairOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SSHKeypair) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

type SSHKeypairArrayOutput struct{ *pulumi.OutputState }

func (SSHKeypairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSHKeypair)(nil)).Elem()
}

func (o SSHKeypairArrayOutput) ToSSHKeypairArrayOutput() SSHKeypairArrayOutput {
	return o
}

func (o SSHKeypairArrayOutput) ToSSHKeypairArrayOutputWithContext(ctx context.Context) SSHKeypairArrayOutput {
	return o
}

func (o SSHKeypairArrayOutput) Index(i pulumi.IntInput) SSHKeypairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SSHKeypair {
		return vs[0].([]*SSHKeypair)[vs[1].(int)]
	}).(SSHKeypairOutput)
}

type SSHKeypairMapOutput struct{ *pulumi.OutputState }

func (SSHKeypairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSHKeypair)(nil)).Elem()
}

func (o SSHKeypairMapOutput) ToSSHKeypairMapOutput() SSHKeypairMapOutput {
	return o
}

func (o SSHKeypairMapOutput) ToSSHKeypairMapOutputWithContext(ctx context.Context) SSHKeypairMapOutput {
	return o
}

func (o SSHKeypairMapOutput) MapIndex(k pulumi.StringInput) SSHKeypairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SSHKeypair {
		return vs[0].(map[string]*SSHKeypair)[vs[1].(string)]
	}).(SSHKeypairOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeypairInput)(nil)).Elem(), &SSHKeypair{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeypairArrayInput)(nil)).Elem(), SSHKeypairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeypairMapInput)(nil)).Elem(), SSHKeypairMap{})
	pulumi.RegisterOutputType(SSHKeypairOutput{})
	pulumi.RegisterOutputType(SSHKeypairArrayOutput{})
	pulumi.RegisterOutputType(SSHKeypairMapOutput{})
}
