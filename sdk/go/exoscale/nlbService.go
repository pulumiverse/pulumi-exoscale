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

// Manage Exoscale [Network Load Balancer (NLB)](https://community.exoscale.com/product/networking/nlb/) Services.
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
//			myNlb, err := exoscale.NewNlb(ctx, "myNlb", &exoscale.NlbArgs{
//				Zone: pulumi.String("ch-gva-2"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = exoscale.NewNlbService(ctx, "myNlbService", &exoscale.NlbServiceArgs{
//				NlbId:          myNlb.ID(),
//				Zone:           myNlb.Zone,
//				InstancePoolId: pulumi.Any(exoscale_instance_pool.My_instance_pool.Id),
//				Protocol:       pulumi.String("tcp"),
//				Port:           pulumi.Int(443),
//				TargetPort:     pulumi.Int(8443),
//				Strategy:       pulumi.String("round-robin"),
//				Healthchecks: exoscale.NlbServiceHealthcheckArray{
//					&exoscale.NlbServiceHealthcheckArgs{
//						Mode:     pulumi.String("https"),
//						Port:     pulumi.Int(8443),
//						Uri:      pulumi.String("/healthz"),
//						TlsSni:   pulumi.String("example.net"),
//						Interval: pulumi.Int(5),
//						Timeout:  pulumi.Int(3),
//						Retries:  pulumi.Int(1),
//					},
//				},
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
// An existing NLB service may be imported by `<nlb-ID>/<service-ID>@<zone>`:
//
// ```sh
// $ pulumi import exoscale:index/nlbService:NlbService \
// ```
//
//	exoscale_nlb_service.my_nlb_service \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524@ch-gva-2
type NlbService struct {
	pulumi.CustomResourceState

	// A free-form text describing the NLB service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The service health checking configuration.
	Healthchecks NlbServiceHealthcheckArrayOutput `pulumi:"healthchecks"`
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId pulumi.StringOutput `pulumi:"instancePoolId"`
	// The NLB service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ❗ The parent Nlb ID.
	NlbId pulumi.StringOutput `pulumi:"nlbId"`
	// The healthcheck port.
	Port pulumi.IntOutput `pulumi:"port"`
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	State    pulumi.StringOutput    `pulumi:"state"`
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy pulumi.StringPtrOutput `pulumi:"strategy"`
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort pulumi.IntOutput `pulumi:"targetPort"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNlbService registers a new resource with the given unique name, arguments, and options.
func NewNlbService(ctx *pulumi.Context,
	name string, args *NlbServiceArgs, opts ...pulumi.ResourceOption) (*NlbService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Healthchecks == nil {
		return nil, errors.New("invalid value for required argument 'Healthchecks'")
	}
	if args.InstancePoolId == nil {
		return nil, errors.New("invalid value for required argument 'InstancePoolId'")
	}
	if args.NlbId == nil {
		return nil, errors.New("invalid value for required argument 'NlbId'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.TargetPort == nil {
		return nil, errors.New("invalid value for required argument 'TargetPort'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NlbService
	err := ctx.RegisterResource("exoscale:index/nlbService:NlbService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNlbService gets an existing NlbService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNlbService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NlbServiceState, opts ...pulumi.ResourceOption) (*NlbService, error) {
	var resource NlbService
	err := ctx.ReadResource("exoscale:index/nlbService:NlbService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NlbService resources.
type nlbServiceState struct {
	// A free-form text describing the NLB service.
	Description *string `pulumi:"description"`
	// The service health checking configuration.
	Healthchecks []NlbServiceHealthcheck `pulumi:"healthchecks"`
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId *string `pulumi:"instancePoolId"`
	// The NLB service name.
	Name *string `pulumi:"name"`
	// ❗ The parent Nlb ID.
	NlbId *string `pulumi:"nlbId"`
	// The healthcheck port.
	Port *int `pulumi:"port"`
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol *string `pulumi:"protocol"`
	State    *string `pulumi:"state"`
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy *string `pulumi:"strategy"`
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort *int `pulumi:"targetPort"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type NlbServiceState struct {
	// A free-form text describing the NLB service.
	Description pulumi.StringPtrInput
	// The service health checking configuration.
	Healthchecks NlbServiceHealthcheckArrayInput
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId pulumi.StringPtrInput
	// The NLB service name.
	Name pulumi.StringPtrInput
	// ❗ The parent Nlb ID.
	NlbId pulumi.StringPtrInput
	// The healthcheck port.
	Port pulumi.IntPtrInput
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol pulumi.StringPtrInput
	State    pulumi.StringPtrInput
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy pulumi.StringPtrInput
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort pulumi.IntPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (NlbServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*nlbServiceState)(nil)).Elem()
}

type nlbServiceArgs struct {
	// A free-form text describing the NLB service.
	Description *string `pulumi:"description"`
	// The service health checking configuration.
	Healthchecks []NlbServiceHealthcheck `pulumi:"healthchecks"`
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId string `pulumi:"instancePoolId"`
	// The NLB service name.
	Name *string `pulumi:"name"`
	// ❗ The parent Nlb ID.
	NlbId string `pulumi:"nlbId"`
	// The healthcheck port.
	Port int `pulumi:"port"`
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol *string `pulumi:"protocol"`
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy *string `pulumi:"strategy"`
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort int `pulumi:"targetPort"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a NlbService resource.
type NlbServiceArgs struct {
	// A free-form text describing the NLB service.
	Description pulumi.StringPtrInput
	// The service health checking configuration.
	Healthchecks NlbServiceHealthcheckArrayInput
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId pulumi.StringInput
	// The NLB service name.
	Name pulumi.StringPtrInput
	// ❗ The parent Nlb ID.
	NlbId pulumi.StringInput
	// The healthcheck port.
	Port pulumi.IntInput
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol pulumi.StringPtrInput
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy pulumi.StringPtrInput
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort pulumi.IntInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (NlbServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nlbServiceArgs)(nil)).Elem()
}

type NlbServiceInput interface {
	pulumi.Input

	ToNlbServiceOutput() NlbServiceOutput
	ToNlbServiceOutputWithContext(ctx context.Context) NlbServiceOutput
}

func (*NlbService) ElementType() reflect.Type {
	return reflect.TypeOf((**NlbService)(nil)).Elem()
}

func (i *NlbService) ToNlbServiceOutput() NlbServiceOutput {
	return i.ToNlbServiceOutputWithContext(context.Background())
}

func (i *NlbService) ToNlbServiceOutputWithContext(ctx context.Context) NlbServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NlbServiceOutput)
}

// NlbServiceArrayInput is an input type that accepts NlbServiceArray and NlbServiceArrayOutput values.
// You can construct a concrete instance of `NlbServiceArrayInput` via:
//
//	NlbServiceArray{ NlbServiceArgs{...} }
type NlbServiceArrayInput interface {
	pulumi.Input

	ToNlbServiceArrayOutput() NlbServiceArrayOutput
	ToNlbServiceArrayOutputWithContext(context.Context) NlbServiceArrayOutput
}

type NlbServiceArray []NlbServiceInput

func (NlbServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NlbService)(nil)).Elem()
}

func (i NlbServiceArray) ToNlbServiceArrayOutput() NlbServiceArrayOutput {
	return i.ToNlbServiceArrayOutputWithContext(context.Background())
}

func (i NlbServiceArray) ToNlbServiceArrayOutputWithContext(ctx context.Context) NlbServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NlbServiceArrayOutput)
}

// NlbServiceMapInput is an input type that accepts NlbServiceMap and NlbServiceMapOutput values.
// You can construct a concrete instance of `NlbServiceMapInput` via:
//
//	NlbServiceMap{ "key": NlbServiceArgs{...} }
type NlbServiceMapInput interface {
	pulumi.Input

	ToNlbServiceMapOutput() NlbServiceMapOutput
	ToNlbServiceMapOutputWithContext(context.Context) NlbServiceMapOutput
}

type NlbServiceMap map[string]NlbServiceInput

func (NlbServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NlbService)(nil)).Elem()
}

func (i NlbServiceMap) ToNlbServiceMapOutput() NlbServiceMapOutput {
	return i.ToNlbServiceMapOutputWithContext(context.Background())
}

func (i NlbServiceMap) ToNlbServiceMapOutputWithContext(ctx context.Context) NlbServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NlbServiceMapOutput)
}

type NlbServiceOutput struct{ *pulumi.OutputState }

func (NlbServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NlbService)(nil)).Elem()
}

func (o NlbServiceOutput) ToNlbServiceOutput() NlbServiceOutput {
	return o
}

func (o NlbServiceOutput) ToNlbServiceOutputWithContext(ctx context.Context) NlbServiceOutput {
	return o
}

// A free-form text describing the NLB service.
func (o NlbServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NlbService) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The service health checking configuration.
func (o NlbServiceOutput) Healthchecks() NlbServiceHealthcheckArrayOutput {
	return o.ApplyT(func(v *NlbService) NlbServiceHealthcheckArrayOutput { return v.Healthchecks }).(NlbServiceHealthcheckArrayOutput)
}

// ❗ The exoscale*instance*pool (ID) to forward traffic to.
func (o NlbServiceOutput) InstancePoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *NlbService) pulumi.StringOutput { return v.InstancePoolId }).(pulumi.StringOutput)
}

// The NLB service name.
func (o NlbServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NlbService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ❗ The parent Nlb ID.
func (o NlbServiceOutput) NlbId() pulumi.StringOutput {
	return o.ApplyT(func(v *NlbService) pulumi.StringOutput { return v.NlbId }).(pulumi.StringOutput)
}

// The healthcheck port.
func (o NlbServiceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *NlbService) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The protocol (`tcp`|`udp`; default: `tcp`).
func (o NlbServiceOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NlbService) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o NlbServiceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NlbService) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
func (o NlbServiceOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NlbService) pulumi.StringPtrOutput { return v.Strategy }).(pulumi.StringPtrOutput)
}

// The (TCP/UDP) port to forward traffic to (on target instance pool members).
func (o NlbServiceOutput) TargetPort() pulumi.IntOutput {
	return o.ApplyT(func(v *NlbService) pulumi.IntOutput { return v.TargetPort }).(pulumi.IntOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o NlbServiceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *NlbService) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type NlbServiceArrayOutput struct{ *pulumi.OutputState }

func (NlbServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NlbService)(nil)).Elem()
}

func (o NlbServiceArrayOutput) ToNlbServiceArrayOutput() NlbServiceArrayOutput {
	return o
}

func (o NlbServiceArrayOutput) ToNlbServiceArrayOutputWithContext(ctx context.Context) NlbServiceArrayOutput {
	return o
}

func (o NlbServiceArrayOutput) Index(i pulumi.IntInput) NlbServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NlbService {
		return vs[0].([]*NlbService)[vs[1].(int)]
	}).(NlbServiceOutput)
}

type NlbServiceMapOutput struct{ *pulumi.OutputState }

func (NlbServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NlbService)(nil)).Elem()
}

func (o NlbServiceMapOutput) ToNlbServiceMapOutput() NlbServiceMapOutput {
	return o
}

func (o NlbServiceMapOutput) ToNlbServiceMapOutputWithContext(ctx context.Context) NlbServiceMapOutput {
	return o
}

func (o NlbServiceMapOutput) MapIndex(k pulumi.StringInput) NlbServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NlbService {
		return vs[0].(map[string]*NlbService)[vs[1].(string)]
	}).(NlbServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NlbServiceInput)(nil)).Elem(), &NlbService{})
	pulumi.RegisterInputType(reflect.TypeOf((*NlbServiceArrayInput)(nil)).Elem(), NlbServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NlbServiceMapInput)(nil)).Elem(), NlbServiceMap{})
	pulumi.RegisterOutputType(NlbServiceOutput{})
	pulumi.RegisterOutputType(NlbServiceArrayOutput{})
	pulumi.RegisterOutputType(NlbServiceMapOutput{})
}
