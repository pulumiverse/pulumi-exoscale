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

// Manage [Exoscale Block Storage](https://community.exoscale.com/documentation/block-storage/) Volume Snapshot.
//
// Block Storage offers persistent externally attached volumes for your workloads.
type BlockStorageVolumeSnapshot struct {
	pulumi.CustomResourceState

	// Snapshot creation date.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Resource labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Volume snapshot name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Snapshot size in GB.
	Size pulumi.IntOutput `pulumi:"size"`
	// Snapshot state.
	State    pulumi.StringOutput                         `pulumi:"state"`
	Timeouts BlockStorageVolumeSnapshotTimeoutsPtrOutput `pulumi:"timeouts"`
	// Volume from which to create a snapshot.
	Volume BlockStorageVolumeSnapshotVolumeOutput `pulumi:"volume"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewBlockStorageVolumeSnapshot registers a new resource with the given unique name, arguments, and options.
func NewBlockStorageVolumeSnapshot(ctx *pulumi.Context,
	name string, args *BlockStorageVolumeSnapshotArgs, opts ...pulumi.ResourceOption) (*BlockStorageVolumeSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Volume == nil {
		return nil, errors.New("invalid value for required argument 'Volume'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BlockStorageVolumeSnapshot
	err := ctx.RegisterResource("exoscale:index/blockStorageVolumeSnapshot:BlockStorageVolumeSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlockStorageVolumeSnapshot gets an existing BlockStorageVolumeSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlockStorageVolumeSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlockStorageVolumeSnapshotState, opts ...pulumi.ResourceOption) (*BlockStorageVolumeSnapshot, error) {
	var resource BlockStorageVolumeSnapshot
	err := ctx.ReadResource("exoscale:index/blockStorageVolumeSnapshot:BlockStorageVolumeSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlockStorageVolumeSnapshot resources.
type blockStorageVolumeSnapshotState struct {
	// Snapshot creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// Resource labels.
	Labels map[string]string `pulumi:"labels"`
	// Volume snapshot name.
	Name *string `pulumi:"name"`
	// Snapshot size in GB.
	Size *int `pulumi:"size"`
	// Snapshot state.
	State    *string                             `pulumi:"state"`
	Timeouts *BlockStorageVolumeSnapshotTimeouts `pulumi:"timeouts"`
	// Volume from which to create a snapshot.
	Volume *BlockStorageVolumeSnapshotVolume `pulumi:"volume"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type BlockStorageVolumeSnapshotState struct {
	// Snapshot creation date.
	CreatedAt pulumi.StringPtrInput
	// Resource labels.
	Labels pulumi.StringMapInput
	// Volume snapshot name.
	Name pulumi.StringPtrInput
	// Snapshot size in GB.
	Size pulumi.IntPtrInput
	// Snapshot state.
	State    pulumi.StringPtrInput
	Timeouts BlockStorageVolumeSnapshotTimeoutsPtrInput
	// Volume from which to create a snapshot.
	Volume BlockStorageVolumeSnapshotVolumePtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (BlockStorageVolumeSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*blockStorageVolumeSnapshotState)(nil)).Elem()
}

type blockStorageVolumeSnapshotArgs struct {
	// Resource labels.
	Labels map[string]string `pulumi:"labels"`
	// Volume snapshot name.
	Name     *string                             `pulumi:"name"`
	Timeouts *BlockStorageVolumeSnapshotTimeouts `pulumi:"timeouts"`
	// Volume from which to create a snapshot.
	Volume BlockStorageVolumeSnapshotVolume `pulumi:"volume"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a BlockStorageVolumeSnapshot resource.
type BlockStorageVolumeSnapshotArgs struct {
	// Resource labels.
	Labels pulumi.StringMapInput
	// Volume snapshot name.
	Name     pulumi.StringPtrInput
	Timeouts BlockStorageVolumeSnapshotTimeoutsPtrInput
	// Volume from which to create a snapshot.
	Volume BlockStorageVolumeSnapshotVolumeInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (BlockStorageVolumeSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blockStorageVolumeSnapshotArgs)(nil)).Elem()
}

type BlockStorageVolumeSnapshotInput interface {
	pulumi.Input

	ToBlockStorageVolumeSnapshotOutput() BlockStorageVolumeSnapshotOutput
	ToBlockStorageVolumeSnapshotOutputWithContext(ctx context.Context) BlockStorageVolumeSnapshotOutput
}

func (*BlockStorageVolumeSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockStorageVolumeSnapshot)(nil)).Elem()
}

func (i *BlockStorageVolumeSnapshot) ToBlockStorageVolumeSnapshotOutput() BlockStorageVolumeSnapshotOutput {
	return i.ToBlockStorageVolumeSnapshotOutputWithContext(context.Background())
}

func (i *BlockStorageVolumeSnapshot) ToBlockStorageVolumeSnapshotOutputWithContext(ctx context.Context) BlockStorageVolumeSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockStorageVolumeSnapshotOutput)
}

// BlockStorageVolumeSnapshotArrayInput is an input type that accepts BlockStorageVolumeSnapshotArray and BlockStorageVolumeSnapshotArrayOutput values.
// You can construct a concrete instance of `BlockStorageVolumeSnapshotArrayInput` via:
//
//	BlockStorageVolumeSnapshotArray{ BlockStorageVolumeSnapshotArgs{...} }
type BlockStorageVolumeSnapshotArrayInput interface {
	pulumi.Input

	ToBlockStorageVolumeSnapshotArrayOutput() BlockStorageVolumeSnapshotArrayOutput
	ToBlockStorageVolumeSnapshotArrayOutputWithContext(context.Context) BlockStorageVolumeSnapshotArrayOutput
}

type BlockStorageVolumeSnapshotArray []BlockStorageVolumeSnapshotInput

func (BlockStorageVolumeSnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockStorageVolumeSnapshot)(nil)).Elem()
}

func (i BlockStorageVolumeSnapshotArray) ToBlockStorageVolumeSnapshotArrayOutput() BlockStorageVolumeSnapshotArrayOutput {
	return i.ToBlockStorageVolumeSnapshotArrayOutputWithContext(context.Background())
}

func (i BlockStorageVolumeSnapshotArray) ToBlockStorageVolumeSnapshotArrayOutputWithContext(ctx context.Context) BlockStorageVolumeSnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockStorageVolumeSnapshotArrayOutput)
}

// BlockStorageVolumeSnapshotMapInput is an input type that accepts BlockStorageVolumeSnapshotMap and BlockStorageVolumeSnapshotMapOutput values.
// You can construct a concrete instance of `BlockStorageVolumeSnapshotMapInput` via:
//
//	BlockStorageVolumeSnapshotMap{ "key": BlockStorageVolumeSnapshotArgs{...} }
type BlockStorageVolumeSnapshotMapInput interface {
	pulumi.Input

	ToBlockStorageVolumeSnapshotMapOutput() BlockStorageVolumeSnapshotMapOutput
	ToBlockStorageVolumeSnapshotMapOutputWithContext(context.Context) BlockStorageVolumeSnapshotMapOutput
}

type BlockStorageVolumeSnapshotMap map[string]BlockStorageVolumeSnapshotInput

func (BlockStorageVolumeSnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockStorageVolumeSnapshot)(nil)).Elem()
}

func (i BlockStorageVolumeSnapshotMap) ToBlockStorageVolumeSnapshotMapOutput() BlockStorageVolumeSnapshotMapOutput {
	return i.ToBlockStorageVolumeSnapshotMapOutputWithContext(context.Background())
}

func (i BlockStorageVolumeSnapshotMap) ToBlockStorageVolumeSnapshotMapOutputWithContext(ctx context.Context) BlockStorageVolumeSnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockStorageVolumeSnapshotMapOutput)
}

type BlockStorageVolumeSnapshotOutput struct{ *pulumi.OutputState }

func (BlockStorageVolumeSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockStorageVolumeSnapshot)(nil)).Elem()
}

func (o BlockStorageVolumeSnapshotOutput) ToBlockStorageVolumeSnapshotOutput() BlockStorageVolumeSnapshotOutput {
	return o
}

func (o BlockStorageVolumeSnapshotOutput) ToBlockStorageVolumeSnapshotOutputWithContext(ctx context.Context) BlockStorageVolumeSnapshotOutput {
	return o
}

// Snapshot creation date.
func (o BlockStorageVolumeSnapshotOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorageVolumeSnapshot) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Resource labels.
func (o BlockStorageVolumeSnapshotOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BlockStorageVolumeSnapshot) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Volume snapshot name.
func (o BlockStorageVolumeSnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorageVolumeSnapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Snapshot size in GB.
func (o BlockStorageVolumeSnapshotOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *BlockStorageVolumeSnapshot) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Snapshot state.
func (o BlockStorageVolumeSnapshotOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorageVolumeSnapshot) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o BlockStorageVolumeSnapshotOutput) Timeouts() BlockStorageVolumeSnapshotTimeoutsPtrOutput {
	return o.ApplyT(func(v *BlockStorageVolumeSnapshot) BlockStorageVolumeSnapshotTimeoutsPtrOutput { return v.Timeouts }).(BlockStorageVolumeSnapshotTimeoutsPtrOutput)
}

// Volume from which to create a snapshot.
func (o BlockStorageVolumeSnapshotOutput) Volume() BlockStorageVolumeSnapshotVolumeOutput {
	return o.ApplyT(func(v *BlockStorageVolumeSnapshot) BlockStorageVolumeSnapshotVolumeOutput { return v.Volume }).(BlockStorageVolumeSnapshotVolumeOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o BlockStorageVolumeSnapshotOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorageVolumeSnapshot) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type BlockStorageVolumeSnapshotArrayOutput struct{ *pulumi.OutputState }

func (BlockStorageVolumeSnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockStorageVolumeSnapshot)(nil)).Elem()
}

func (o BlockStorageVolumeSnapshotArrayOutput) ToBlockStorageVolumeSnapshotArrayOutput() BlockStorageVolumeSnapshotArrayOutput {
	return o
}

func (o BlockStorageVolumeSnapshotArrayOutput) ToBlockStorageVolumeSnapshotArrayOutputWithContext(ctx context.Context) BlockStorageVolumeSnapshotArrayOutput {
	return o
}

func (o BlockStorageVolumeSnapshotArrayOutput) Index(i pulumi.IntInput) BlockStorageVolumeSnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BlockStorageVolumeSnapshot {
		return vs[0].([]*BlockStorageVolumeSnapshot)[vs[1].(int)]
	}).(BlockStorageVolumeSnapshotOutput)
}

type BlockStorageVolumeSnapshotMapOutput struct{ *pulumi.OutputState }

func (BlockStorageVolumeSnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockStorageVolumeSnapshot)(nil)).Elem()
}

func (o BlockStorageVolumeSnapshotMapOutput) ToBlockStorageVolumeSnapshotMapOutput() BlockStorageVolumeSnapshotMapOutput {
	return o
}

func (o BlockStorageVolumeSnapshotMapOutput) ToBlockStorageVolumeSnapshotMapOutputWithContext(ctx context.Context) BlockStorageVolumeSnapshotMapOutput {
	return o
}

func (o BlockStorageVolumeSnapshotMapOutput) MapIndex(k pulumi.StringInput) BlockStorageVolumeSnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BlockStorageVolumeSnapshot {
		return vs[0].(map[string]*BlockStorageVolumeSnapshot)[vs[1].(string)]
	}).(BlockStorageVolumeSnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BlockStorageVolumeSnapshotInput)(nil)).Elem(), &BlockStorageVolumeSnapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockStorageVolumeSnapshotArrayInput)(nil)).Elem(), BlockStorageVolumeSnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockStorageVolumeSnapshotMapInput)(nil)).Elem(), BlockStorageVolumeSnapshotMap{})
	pulumi.RegisterOutputType(BlockStorageVolumeSnapshotOutput{})
	pulumi.RegisterOutputType(BlockStorageVolumeSnapshotArrayOutput{})
	pulumi.RegisterOutputType(BlockStorageVolumeSnapshotMapOutput{})
}
