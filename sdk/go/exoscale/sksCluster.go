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

// Manage Exoscale [Scalable Kubernetes Service (SKS)](https://community.exoscale.com/documentation/sks/) Clusters.
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
//			mySksCluster, err := exoscale.NewSksCluster(ctx, "mySksCluster", &exoscale.SksClusterArgs{
//				Zone: pulumi.String("ch-gva-2"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("mySksClusterEndpoint", mySksCluster.Endpoint)
//			return nil
//		})
//	}
//
// ```
//
// Next step is to attach exoscale_sks_nodepool(s) to the cluster.
//
// Please refer to the examples
// directory for complete configuration examples.
//
// ## Import
//
// An existing SKS cluster may be imported by `<ID>@<zone>`
//
// ```sh
//
//	$ pulumi import exoscale:index/sksCluster:SksCluster \
//
// ```
//
//	exoscale_sks_cluster.my_sks_cluster \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
type SksCluster struct {
	pulumi.CustomResourceState

	// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
	Addons pulumi.StringArrayOutput `pulumi:"addons"`
	// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
	AggregationCa pulumi.StringOutput `pulumi:"aggregationCa"`
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade pulumi.BoolPtrOutput `pulumi:"autoUpgrade"`
	// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	Cni pulumi.StringPtrOutput `pulumi:"cni"`
	// The CA certificate (in PEM format) for TLS communications between control plane components.
	ControlPlaneCa pulumi.StringOutput `pulumi:"controlPlaneCa"`
	// The cluster creation date.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A free-form text describing the cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The cluster API endpoint.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm pulumi.BoolPtrOutput `pulumi:"exoscaleCcm"`
	// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
	KubeletCa pulumi.StringOutput `pulumi:"kubeletCa"`
	// A map of key/value labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer pulumi.BoolPtrOutput `pulumi:"metricsServer"`
	// The SKS cluster name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of exoscale*sks*nodepool (IDs) attached to the cluster.
	Nodepools pulumi.StringArrayOutput `pulumi:"nodepools"`
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc SksClusterOidcOutput `pulumi:"oidc"`
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel pulumi.StringPtrOutput `pulumi:"serviceLevel"`
	// The cluster state.
	State pulumi.StringOutput `pulumi:"state"`
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version pulumi.StringOutput `pulumi:"version"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewSksCluster registers a new resource with the given unique name, arguments, and options.
func NewSksCluster(ctx *pulumi.Context,
	name string, args *SksClusterArgs, opts ...pulumi.ResourceOption) (*SksCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SksCluster
	err := ctx.RegisterResource("exoscale:index/sksCluster:SksCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSksCluster gets an existing SksCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSksCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SksClusterState, opts ...pulumi.ResourceOption) (*SksCluster, error) {
	var resource SksCluster
	err := ctx.ReadResource("exoscale:index/sksCluster:SksCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SksCluster resources.
type sksClusterState struct {
	// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
	Addons []string `pulumi:"addons"`
	// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
	AggregationCa *string `pulumi:"aggregationCa"`
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade *bool `pulumi:"autoUpgrade"`
	// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	Cni *string `pulumi:"cni"`
	// The CA certificate (in PEM format) for TLS communications between control plane components.
	ControlPlaneCa *string `pulumi:"controlPlaneCa"`
	// The cluster creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// A free-form text describing the cluster.
	Description *string `pulumi:"description"`
	// The cluster API endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm *bool `pulumi:"exoscaleCcm"`
	// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
	KubeletCa *string `pulumi:"kubeletCa"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer *bool `pulumi:"metricsServer"`
	// The SKS cluster name.
	Name *string `pulumi:"name"`
	// The list of exoscale*sks*nodepool (IDs) attached to the cluster.
	Nodepools []string `pulumi:"nodepools"`
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc *SksClusterOidc `pulumi:"oidc"`
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel *string `pulumi:"serviceLevel"`
	// The cluster state.
	State *string `pulumi:"state"`
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version *string `pulumi:"version"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type SksClusterState struct {
	// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
	Addons pulumi.StringArrayInput
	// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
	AggregationCa pulumi.StringPtrInput
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade pulumi.BoolPtrInput
	// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	Cni pulumi.StringPtrInput
	// The CA certificate (in PEM format) for TLS communications between control plane components.
	ControlPlaneCa pulumi.StringPtrInput
	// The cluster creation date.
	CreatedAt pulumi.StringPtrInput
	// A free-form text describing the cluster.
	Description pulumi.StringPtrInput
	// The cluster API endpoint.
	Endpoint pulumi.StringPtrInput
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm pulumi.BoolPtrInput
	// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
	KubeletCa pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer pulumi.BoolPtrInput
	// The SKS cluster name.
	Name pulumi.StringPtrInput
	// The list of exoscale*sks*nodepool (IDs) attached to the cluster.
	Nodepools pulumi.StringArrayInput
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc SksClusterOidcPtrInput
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel pulumi.StringPtrInput
	// The cluster state.
	State pulumi.StringPtrInput
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (SksClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*sksClusterState)(nil)).Elem()
}

type sksClusterArgs struct {
	// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
	Addons []string `pulumi:"addons"`
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade *bool `pulumi:"autoUpgrade"`
	// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	Cni *string `pulumi:"cni"`
	// A free-form text describing the cluster.
	Description *string `pulumi:"description"`
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm *bool `pulumi:"exoscaleCcm"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer *bool `pulumi:"metricsServer"`
	// The SKS cluster name.
	Name *string `pulumi:"name"`
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc *SksClusterOidc `pulumi:"oidc"`
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel *string `pulumi:"serviceLevel"`
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version *string `pulumi:"version"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a SksCluster resource.
type SksClusterArgs struct {
	// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
	Addons pulumi.StringArrayInput
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade pulumi.BoolPtrInput
	// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	Cni pulumi.StringPtrInput
	// A free-form text describing the cluster.
	Description pulumi.StringPtrInput
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm pulumi.BoolPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer pulumi.BoolPtrInput
	// The SKS cluster name.
	Name pulumi.StringPtrInput
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc SksClusterOidcPtrInput
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel pulumi.StringPtrInput
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (SksClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sksClusterArgs)(nil)).Elem()
}

type SksClusterInput interface {
	pulumi.Input

	ToSksClusterOutput() SksClusterOutput
	ToSksClusterOutputWithContext(ctx context.Context) SksClusterOutput
}

func (*SksCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**SksCluster)(nil)).Elem()
}

func (i *SksCluster) ToSksClusterOutput() SksClusterOutput {
	return i.ToSksClusterOutputWithContext(context.Background())
}

func (i *SksCluster) ToSksClusterOutputWithContext(ctx context.Context) SksClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SksClusterOutput)
}

// SksClusterArrayInput is an input type that accepts SksClusterArray and SksClusterArrayOutput values.
// You can construct a concrete instance of `SksClusterArrayInput` via:
//
//	SksClusterArray{ SksClusterArgs{...} }
type SksClusterArrayInput interface {
	pulumi.Input

	ToSksClusterArrayOutput() SksClusterArrayOutput
	ToSksClusterArrayOutputWithContext(context.Context) SksClusterArrayOutput
}

type SksClusterArray []SksClusterInput

func (SksClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SksCluster)(nil)).Elem()
}

func (i SksClusterArray) ToSksClusterArrayOutput() SksClusterArrayOutput {
	return i.ToSksClusterArrayOutputWithContext(context.Background())
}

func (i SksClusterArray) ToSksClusterArrayOutputWithContext(ctx context.Context) SksClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SksClusterArrayOutput)
}

// SksClusterMapInput is an input type that accepts SksClusterMap and SksClusterMapOutput values.
// You can construct a concrete instance of `SksClusterMapInput` via:
//
//	SksClusterMap{ "key": SksClusterArgs{...} }
type SksClusterMapInput interface {
	pulumi.Input

	ToSksClusterMapOutput() SksClusterMapOutput
	ToSksClusterMapOutputWithContext(context.Context) SksClusterMapOutput
}

type SksClusterMap map[string]SksClusterInput

func (SksClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SksCluster)(nil)).Elem()
}

func (i SksClusterMap) ToSksClusterMapOutput() SksClusterMapOutput {
	return i.ToSksClusterMapOutputWithContext(context.Background())
}

func (i SksClusterMap) ToSksClusterMapOutputWithContext(ctx context.Context) SksClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SksClusterMapOutput)
}

type SksClusterOutput struct{ *pulumi.OutputState }

func (SksClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SksCluster)(nil)).Elem()
}

func (o SksClusterOutput) ToSksClusterOutput() SksClusterOutput {
	return o
}

func (o SksClusterOutput) ToSksClusterOutputWithContext(ctx context.Context) SksClusterOutput {
	return o
}

// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
func (o SksClusterOutput) Addons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringArrayOutput { return v.Addons }).(pulumi.StringArrayOutput)
}

// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
func (o SksClusterOutput) AggregationCa() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.AggregationCa }).(pulumi.StringOutput)
}

// Enable automatic upgrading of the control plane version.
func (o SksClusterOutput) AutoUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.BoolPtrOutput { return v.AutoUpgrade }).(pulumi.BoolPtrOutput)
}

// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
func (o SksClusterOutput) Cni() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringPtrOutput { return v.Cni }).(pulumi.StringPtrOutput)
}

// The CA certificate (in PEM format) for TLS communications between control plane components.
func (o SksClusterOutput) ControlPlaneCa() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.ControlPlaneCa }).(pulumi.StringOutput)
}

// The cluster creation date.
func (o SksClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A free-form text describing the cluster.
func (o SksClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The cluster API endpoint.
func (o SksClusterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
func (o SksClusterOutput) ExoscaleCcm() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.BoolPtrOutput { return v.ExoscaleCcm }).(pulumi.BoolPtrOutput)
}

// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
func (o SksClusterOutput) KubeletCa() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.KubeletCa }).(pulumi.StringOutput)
}

// A map of key/value labels.
func (o SksClusterOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
func (o SksClusterOutput) MetricsServer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.BoolPtrOutput { return v.MetricsServer }).(pulumi.BoolPtrOutput)
}

// The SKS cluster name.
func (o SksClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of exoscale*sks*nodepool (IDs) attached to the cluster.
func (o SksClusterOutput) Nodepools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringArrayOutput { return v.Nodepools }).(pulumi.StringArrayOutput)
}

// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
func (o SksClusterOutput) Oidc() SksClusterOidcOutput {
	return o.ApplyT(func(v *SksCluster) SksClusterOidcOutput { return v.Oidc }).(SksClusterOidcOutput)
}

// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
func (o SksClusterOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringPtrOutput { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

// The cluster state.
func (o SksClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
func (o SksClusterOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o SksClusterOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *SksCluster) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type SksClusterArrayOutput struct{ *pulumi.OutputState }

func (SksClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SksCluster)(nil)).Elem()
}

func (o SksClusterArrayOutput) ToSksClusterArrayOutput() SksClusterArrayOutput {
	return o
}

func (o SksClusterArrayOutput) ToSksClusterArrayOutputWithContext(ctx context.Context) SksClusterArrayOutput {
	return o
}

func (o SksClusterArrayOutput) Index(i pulumi.IntInput) SksClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SksCluster {
		return vs[0].([]*SksCluster)[vs[1].(int)]
	}).(SksClusterOutput)
}

type SksClusterMapOutput struct{ *pulumi.OutputState }

func (SksClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SksCluster)(nil)).Elem()
}

func (o SksClusterMapOutput) ToSksClusterMapOutput() SksClusterMapOutput {
	return o
}

func (o SksClusterMapOutput) ToSksClusterMapOutputWithContext(ctx context.Context) SksClusterMapOutput {
	return o
}

func (o SksClusterMapOutput) MapIndex(k pulumi.StringInput) SksClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SksCluster {
		return vs[0].(map[string]*SksCluster)[vs[1].(string)]
	}).(SksClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SksClusterInput)(nil)).Elem(), &SksCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*SksClusterArrayInput)(nil)).Elem(), SksClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SksClusterMapInput)(nil)).Elem(), SksClusterMap{})
	pulumi.RegisterOutputType(SksClusterOutput{})
	pulumi.RegisterOutputType(SksClusterArrayOutput{})
	pulumi.RegisterOutputType(SksClusterMapOutput{})
}