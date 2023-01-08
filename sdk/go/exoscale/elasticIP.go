// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage Exoscale [Elastic IPs (EIP)](https://community.exoscale.com/documentation/compute/eip/).
//
// Corresponding data source: exoscale_elastic_ip.
//
// ## Import
//
// An existing Elastic IP (EIP) may be imported by `<ID>@<zone>`console
//
// ```sh
//
//	$ pulumi import exoscale:index/elasticIP:ElasticIP \
//
// ```
//
//	exoscale_elastic_ip.my_elastic_ip \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
type ElasticIP struct {
	pulumi.CustomResourceState

	// The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// The Elastic IP (EIP) CIDR.
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// A free-form text describing the Elastic IP (EIP).
	Description pulumi.StringOutput `pulumi:"description"`
	// Healthcheck configuration for *managed* EIPs. Structure is documented below.
	Healthcheck ElasticIPHealthcheckOutput `pulumi:"healthcheck"`
	// The Elastic IP (EIP) IPv4 or IPv6 address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// A map of key/value labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Domain name for reverse DNS record.
	ReverseDns pulumi.StringPtrOutput `pulumi:"reverseDns"`
	// The Exoscale [Zone][zone] name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewElasticIP registers a new resource with the given unique name, arguments, and options.
func NewElasticIP(ctx *pulumi.Context,
	name string, args *ElasticIPArgs, opts ...pulumi.ResourceOption) (*ElasticIP, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ElasticIP
	err := ctx.RegisterResource("exoscale:index/elasticIP:ElasticIP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticIP gets an existing ElasticIP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticIPState, opts ...pulumi.ResourceOption) (*ElasticIP, error) {
	var resource ElasticIP
	err := ctx.ReadResource("exoscale:index/elasticIP:ElasticIP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticIP resources.
type elasticIPState struct {
	// The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
	AddressFamily *string `pulumi:"addressFamily"`
	// The Elastic IP (EIP) CIDR.
	Cidr *string `pulumi:"cidr"`
	// A free-form text describing the Elastic IP (EIP).
	Description *string `pulumi:"description"`
	// Healthcheck configuration for *managed* EIPs. Structure is documented below.
	Healthcheck *ElasticIPHealthcheck `pulumi:"healthcheck"`
	// The Elastic IP (EIP) IPv4 or IPv6 address.
	IpAddress *string `pulumi:"ipAddress"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// Domain name for reverse DNS record.
	ReverseDns *string `pulumi:"reverseDns"`
	// The Exoscale [Zone][zone] name.
	Zone *string `pulumi:"zone"`
}

type ElasticIPState struct {
	// The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
	AddressFamily pulumi.StringPtrInput
	// The Elastic IP (EIP) CIDR.
	Cidr pulumi.StringPtrInput
	// A free-form text describing the Elastic IP (EIP).
	Description pulumi.StringPtrInput
	// Healthcheck configuration for *managed* EIPs. Structure is documented below.
	Healthcheck ElasticIPHealthcheckPtrInput
	// The Elastic IP (EIP) IPv4 or IPv6 address.
	IpAddress pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// Domain name for reverse DNS record.
	ReverseDns pulumi.StringPtrInput
	// The Exoscale [Zone][zone] name.
	Zone pulumi.StringPtrInput
}

func (ElasticIPState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticIPState)(nil)).Elem()
}

type elasticIPArgs struct {
	// The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
	AddressFamily *string `pulumi:"addressFamily"`
	// A free-form text describing the Elastic IP (EIP).
	Description *string `pulumi:"description"`
	// Healthcheck configuration for *managed* EIPs. Structure is documented below.
	Healthcheck *ElasticIPHealthcheck `pulumi:"healthcheck"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// Domain name for reverse DNS record.
	ReverseDns *string `pulumi:"reverseDns"`
	// The Exoscale [Zone][zone] name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a ElasticIP resource.
type ElasticIPArgs struct {
	// The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
	AddressFamily pulumi.StringPtrInput
	// A free-form text describing the Elastic IP (EIP).
	Description pulumi.StringPtrInput
	// Healthcheck configuration for *managed* EIPs. Structure is documented below.
	Healthcheck ElasticIPHealthcheckPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// Domain name for reverse DNS record.
	ReverseDns pulumi.StringPtrInput
	// The Exoscale [Zone][zone] name.
	Zone pulumi.StringInput
}

func (ElasticIPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticIPArgs)(nil)).Elem()
}

type ElasticIPInput interface {
	pulumi.Input

	ToElasticIPOutput() ElasticIPOutput
	ToElasticIPOutputWithContext(ctx context.Context) ElasticIPOutput
}

func (*ElasticIP) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticIP)(nil)).Elem()
}

func (i *ElasticIP) ToElasticIPOutput() ElasticIPOutput {
	return i.ToElasticIPOutputWithContext(context.Background())
}

func (i *ElasticIP) ToElasticIPOutputWithContext(ctx context.Context) ElasticIPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticIPOutput)
}

// ElasticIPArrayInput is an input type that accepts ElasticIPArray and ElasticIPArrayOutput values.
// You can construct a concrete instance of `ElasticIPArrayInput` via:
//
//	ElasticIPArray{ ElasticIPArgs{...} }
type ElasticIPArrayInput interface {
	pulumi.Input

	ToElasticIPArrayOutput() ElasticIPArrayOutput
	ToElasticIPArrayOutputWithContext(context.Context) ElasticIPArrayOutput
}

type ElasticIPArray []ElasticIPInput

func (ElasticIPArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticIP)(nil)).Elem()
}

func (i ElasticIPArray) ToElasticIPArrayOutput() ElasticIPArrayOutput {
	return i.ToElasticIPArrayOutputWithContext(context.Background())
}

func (i ElasticIPArray) ToElasticIPArrayOutputWithContext(ctx context.Context) ElasticIPArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticIPArrayOutput)
}

// ElasticIPMapInput is an input type that accepts ElasticIPMap and ElasticIPMapOutput values.
// You can construct a concrete instance of `ElasticIPMapInput` via:
//
//	ElasticIPMap{ "key": ElasticIPArgs{...} }
type ElasticIPMapInput interface {
	pulumi.Input

	ToElasticIPMapOutput() ElasticIPMapOutput
	ToElasticIPMapOutputWithContext(context.Context) ElasticIPMapOutput
}

type ElasticIPMap map[string]ElasticIPInput

func (ElasticIPMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticIP)(nil)).Elem()
}

func (i ElasticIPMap) ToElasticIPMapOutput() ElasticIPMapOutput {
	return i.ToElasticIPMapOutputWithContext(context.Background())
}

func (i ElasticIPMap) ToElasticIPMapOutputWithContext(ctx context.Context) ElasticIPMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticIPMapOutput)
}

type ElasticIPOutput struct{ *pulumi.OutputState }

func (ElasticIPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticIP)(nil)).Elem()
}

func (o ElasticIPOutput) ToElasticIPOutput() ElasticIPOutput {
	return o
}

func (o ElasticIPOutput) ToElasticIPOutputWithContext(ctx context.Context) ElasticIPOutput {
	return o
}

// The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
func (o ElasticIPOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticIP) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// The Elastic IP (EIP) CIDR.
func (o ElasticIPOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticIP) pulumi.StringOutput { return v.Cidr }).(pulumi.StringOutput)
}

// A free-form text describing the Elastic IP (EIP).
func (o ElasticIPOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticIP) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Healthcheck configuration for *managed* EIPs. Structure is documented below.
func (o ElasticIPOutput) Healthcheck() ElasticIPHealthcheckOutput {
	return o.ApplyT(func(v *ElasticIP) ElasticIPHealthcheckOutput { return v.Healthcheck }).(ElasticIPHealthcheckOutput)
}

// The Elastic IP (EIP) IPv4 or IPv6 address.
func (o ElasticIPOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticIP) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// A map of key/value labels.
func (o ElasticIPOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ElasticIP) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Domain name for reverse DNS record.
func (o ElasticIPOutput) ReverseDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticIP) pulumi.StringPtrOutput { return v.ReverseDns }).(pulumi.StringPtrOutput)
}

// The Exoscale [Zone][zone] name.
func (o ElasticIPOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticIP) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type ElasticIPArrayOutput struct{ *pulumi.OutputState }

func (ElasticIPArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticIP)(nil)).Elem()
}

func (o ElasticIPArrayOutput) ToElasticIPArrayOutput() ElasticIPArrayOutput {
	return o
}

func (o ElasticIPArrayOutput) ToElasticIPArrayOutputWithContext(ctx context.Context) ElasticIPArrayOutput {
	return o
}

func (o ElasticIPArrayOutput) Index(i pulumi.IntInput) ElasticIPOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElasticIP {
		return vs[0].([]*ElasticIP)[vs[1].(int)]
	}).(ElasticIPOutput)
}

type ElasticIPMapOutput struct{ *pulumi.OutputState }

func (ElasticIPMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticIP)(nil)).Elem()
}

func (o ElasticIPMapOutput) ToElasticIPMapOutput() ElasticIPMapOutput {
	return o
}

func (o ElasticIPMapOutput) ToElasticIPMapOutputWithContext(ctx context.Context) ElasticIPMapOutput {
	return o
}

func (o ElasticIPMapOutput) MapIndex(k pulumi.StringInput) ElasticIPOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElasticIP {
		return vs[0].(map[string]*ElasticIP)[vs[1].(string)]
	}).(ElasticIPOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticIPInput)(nil)).Elem(), &ElasticIP{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticIPArrayInput)(nil)).Elem(), ElasticIPArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticIPMapInput)(nil)).Elem(), ElasticIPMap{})
	pulumi.RegisterOutputType(ElasticIPOutput{})
	pulumi.RegisterOutputType(ElasticIPArrayOutput{})
	pulumi.RegisterOutputType(ElasticIPMapOutput{})
}
