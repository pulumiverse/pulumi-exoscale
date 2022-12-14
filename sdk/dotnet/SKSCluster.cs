// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale
{
    /// <summary>
    /// ## Import
    /// 
    /// An existing SKS cluster may be imported by `&lt;ID&gt;@&lt;zone&gt;`console
    /// 
    /// ```sh
    ///  $ pulumi import exoscale:index/sKSCluster:SKSCluster \
    /// ```
    /// 
    ///  exoscale_sks_cluster.my_sks_cluster \
    /// 
    ///  f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
    /// </summary>
    [ExoscaleResourceType("exoscale:index/sKSCluster:SKSCluster")]
    public partial class SKSCluster : global::Pulumi.CustomResource
    {
        [Output("addons")]
        public Output<ImmutableArray<string>> Addons { get; private set; } = null!;

        /// <summary>
        /// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
        /// </summary>
        [Output("aggregationCa")]
        public Output<string> AggregationCa { get; private set; } = null!;

        /// <summary>
        /// Enable automatic upgrading of the control plane version (boolean; default: `false`).
        /// </summary>
        [Output("autoUpgrade")]
        public Output<bool?> AutoUpgrade { get; private set; } = null!;

        [Output("cni")]
        public Output<string?> Cni { get; private set; } = null!;

        /// <summary>
        /// The CA certificate (in PEM format) for TLS communications between control plane components.
        /// </summary>
        [Output("controlPlaneCa")]
        public Output<string> ControlPlaneCa { get; private set; } = null!;

        /// <summary>
        /// The cluster creation date.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A free-form text describing the cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The cluster API endpoint.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Deploy the Exoscale [Cloud Controller Manager][ccm] in the control plane (boolean; default: `true`; may only be set at creation time).
        /// </summary>
        [Output("exoscaleCcm")]
        public Output<bool?> ExoscaleCcm { get; private set; } = null!;

        /// <summary>
        /// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
        /// </summary>
        [Output("kubeletCa")]
        public Output<string> KubeletCa { get; private set; } = null!;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Deploy the [Kubernetes Metrics Server][ms] in the control plane (boolean; default: `true`; may only be set at creation time).
        /// </summary>
        [Output("metricsServer")]
        public Output<bool?> MetricsServer { get; private set; } = null!;

        /// <summary>
        /// The SKS cluster name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of exoscale.SKSNodepool (IDs) attached to the cluster.
        /// </summary>
        [Output("nodepools")]
        public Output<ImmutableArray<string>> Nodepools { get; private set; } = null!;

        /// <summary>
        /// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
        /// </summary>
        [Output("oidc")]
        public Output<Outputs.SKSClusterOidc> Oidc { get; private set; } = null!;

        /// <summary>
        /// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
        /// </summary>
        [Output("serviceLevel")]
        public Output<string?> ServiceLevel { get; private set; } = null!;

        /// <summary>
        /// The cluster state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The Exoscale [Zone][zone] name.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a SKSCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SKSCluster(string name, SKSClusterArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/sKSCluster:SKSCluster", name, args ?? new SKSClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SKSCluster(string name, Input<string> id, SKSClusterState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/sKSCluster:SKSCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-exoscale",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SKSCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SKSCluster Get(string name, Input<string> id, SKSClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new SKSCluster(name, id, state, options);
        }
    }

    public sealed class SKSClusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<string>? _addons;
        [Obsolete(@"This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.")]
        public InputList<string> Addons
        {
            get => _addons ?? (_addons = new InputList<string>());
            set => _addons = value;
        }

        /// <summary>
        /// Enable automatic upgrading of the control plane version (boolean; default: `false`).
        /// </summary>
        [Input("autoUpgrade")]
        public Input<bool>? AutoUpgrade { get; set; }

        [Input("cni")]
        public Input<string>? Cni { get; set; }

        /// <summary>
        /// A free-form text describing the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Deploy the Exoscale [Cloud Controller Manager][ccm] in the control plane (boolean; default: `true`; may only be set at creation time).
        /// </summary>
        [Input("exoscaleCcm")]
        public Input<bool>? ExoscaleCcm { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Deploy the [Kubernetes Metrics Server][ms] in the control plane (boolean; default: `true`; may only be set at creation time).
        /// </summary>
        [Input("metricsServer")]
        public Input<bool>? MetricsServer { get; set; }

        /// <summary>
        /// The SKS cluster name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
        /// </summary>
        [Input("oidc")]
        public Input<Inputs.SKSClusterOidcArgs>? Oidc { get; set; }

        /// <summary>
        /// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
        /// </summary>
        [Input("serviceLevel")]
        public Input<string>? ServiceLevel { get; set; }

        /// <summary>
        /// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The Exoscale [Zone][zone] name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public SKSClusterArgs()
        {
        }
        public static new SKSClusterArgs Empty => new SKSClusterArgs();
    }

    public sealed class SKSClusterState : global::Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<string>? _addons;
        [Obsolete(@"This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.")]
        public InputList<string> Addons
        {
            get => _addons ?? (_addons = new InputList<string>());
            set => _addons = value;
        }

        /// <summary>
        /// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
        /// </summary>
        [Input("aggregationCa")]
        public Input<string>? AggregationCa { get; set; }

        /// <summary>
        /// Enable automatic upgrading of the control plane version (boolean; default: `false`).
        /// </summary>
        [Input("autoUpgrade")]
        public Input<bool>? AutoUpgrade { get; set; }

        [Input("cni")]
        public Input<string>? Cni { get; set; }

        /// <summary>
        /// The CA certificate (in PEM format) for TLS communications between control plane components.
        /// </summary>
        [Input("controlPlaneCa")]
        public Input<string>? ControlPlaneCa { get; set; }

        /// <summary>
        /// The cluster creation date.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A free-form text describing the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The cluster API endpoint.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Deploy the Exoscale [Cloud Controller Manager][ccm] in the control plane (boolean; default: `true`; may only be set at creation time).
        /// </summary>
        [Input("exoscaleCcm")]
        public Input<bool>? ExoscaleCcm { get; set; }

        /// <summary>
        /// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
        /// </summary>
        [Input("kubeletCa")]
        public Input<string>? KubeletCa { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Deploy the [Kubernetes Metrics Server][ms] in the control plane (boolean; default: `true`; may only be set at creation time).
        /// </summary>
        [Input("metricsServer")]
        public Input<bool>? MetricsServer { get; set; }

        /// <summary>
        /// The SKS cluster name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodepools")]
        private InputList<string>? _nodepools;

        /// <summary>
        /// The list of exoscale.SKSNodepool (IDs) attached to the cluster.
        /// </summary>
        public InputList<string> Nodepools
        {
            get => _nodepools ?? (_nodepools = new InputList<string>());
            set => _nodepools = value;
        }

        /// <summary>
        /// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
        /// </summary>
        [Input("oidc")]
        public Input<Inputs.SKSClusterOidcGetArgs>? Oidc { get; set; }

        /// <summary>
        /// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
        /// </summary>
        [Input("serviceLevel")]
        public Input<string>? ServiceLevel { get; set; }

        /// <summary>
        /// The cluster state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The Exoscale [Zone][zone] name.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public SKSClusterState()
        {
        }
        public static new SKSClusterState Empty => new SKSClusterState();
    }
}
