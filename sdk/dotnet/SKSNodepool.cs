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
    /// An existing SKS node pool may be imported by `&lt;cluster-ID&gt;/&lt;nodepool-ID&gt;@&lt;zone&gt;`
    /// 
    /// ```sh
    ///  $ pulumi import exoscale:index/sKSNodepool:SKSNodepool \
    /// ```
    /// 
    ///  exoscale_sks_nodepool.my_sks_nodepool \
    /// 
    ///  f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524@ch-gva-2
    /// </summary>
    [ExoscaleResourceType("exoscale:index/sKSNodepool:SKSNodepool")]
    public partial class SKSNodepool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
        /// </summary>
        [Output("antiAffinityGroupIds")]
        public Output<ImmutableArray<string>> AntiAffinityGroupIds { get; private set; } = null!;

        /// <summary>
        /// ❗ The parent exoscale*sks*cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The pool creation date.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Output("deployTargetId")]
        public Output<string?> DeployTargetId { get; private set; } = null!;

        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The managed instances disk size (GiB; default: `50`).
        /// </summary>
        [Output("diskSize")]
        public Output<int?> DiskSize { get; private set; } = null!;

        /// <summary>
        /// The underlying exoscale*instance*pool ID.
        /// </summary>
        [Output("instancePoolId")]
        public Output<string> InstancePoolId { get; private set; } = null!;

        /// <summary>
        /// The string used to prefix the managed instances name (default `pool`).
        /// </summary>
        [Output("instancePrefix")]
        public Output<string?> InstancePrefix { get; private set; } = null!;

        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The SKS node pool name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of exoscale*private*network (IDs) to be attached to the managed instances.
        /// </summary>
        [Output("privateNetworkIds")]
        public Output<ImmutableArray<string>> PrivateNetworkIds { get; private set; } = null!;

        /// <summary>
        /// A list of exoscale*security*group (IDs) to be attached to the managed instances.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The current pool state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
        /// </summary>
        [Output("storageLvm")]
        public Output<bool?> StorageLvm { get; private set; } = null!;

        /// <summary>
        /// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) (`&lt;value&gt;:&lt;effect&gt;`).
        /// </summary>
        [Output("taints")]
        public Output<ImmutableDictionary<string, string>?> Taints { get; private set; } = null!;

        /// <summary>
        /// The managed instances template ID.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// The managed instances version.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a SKSNodepool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SKSNodepool(string name, SKSNodepoolArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/sKSNodepool:SKSNodepool", name, args ?? new SKSNodepoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SKSNodepool(string name, Input<string> id, SKSNodepoolState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/sKSNodepool:SKSNodepool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SKSNodepool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SKSNodepool Get(string name, Input<string> id, SKSNodepoolState? state = null, CustomResourceOptions? options = null)
        {
            return new SKSNodepool(name, id, state, options);
        }
    }

    public sealed class SKSNodepoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("antiAffinityGroupIds")]
        private InputList<string>? _antiAffinityGroupIds;

        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> AntiAffinityGroupIds
        {
            get => _antiAffinityGroupIds ?? (_antiAffinityGroupIds = new InputList<string>());
            set => _antiAffinityGroupIds = value;
        }

        /// <summary>
        /// ❗ The parent exoscale*sks*cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The managed instances disk size (GiB; default: `50`).
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The string used to prefix the managed instances name (default `pool`).
        /// </summary>
        [Input("instancePrefix")]
        public Input<string>? InstancePrefix { get; set; }

        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

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
        /// The SKS node pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateNetworkIds")]
        private InputList<string>? _privateNetworkIds;

        /// <summary>
        /// A list of exoscale*private*network (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> PrivateNetworkIds
        {
            get => _privateNetworkIds ?? (_privateNetworkIds = new InputList<string>());
            set => _privateNetworkIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of exoscale*security*group (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
        /// </summary>
        [Input("storageLvm")]
        public Input<bool>? StorageLvm { get; set; }

        [Input("taints")]
        private InputMap<string>? _taints;

        /// <summary>
        /// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) (`&lt;value&gt;:&lt;effect&gt;`).
        /// </summary>
        public InputMap<string> Taints
        {
            get => _taints ?? (_taints = new InputMap<string>());
            set => _taints = value;
        }

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public SKSNodepoolArgs()
        {
        }
        public static new SKSNodepoolArgs Empty => new SKSNodepoolArgs();
    }

    public sealed class SKSNodepoolState : global::Pulumi.ResourceArgs
    {
        [Input("antiAffinityGroupIds")]
        private InputList<string>? _antiAffinityGroupIds;

        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> AntiAffinityGroupIds
        {
            get => _antiAffinityGroupIds ?? (_antiAffinityGroupIds = new InputList<string>());
            set => _antiAffinityGroupIds = value;
        }

        /// <summary>
        /// ❗ The parent exoscale*sks*cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The pool creation date.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The managed instances disk size (GiB; default: `50`).
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The underlying exoscale*instance*pool ID.
        /// </summary>
        [Input("instancePoolId")]
        public Input<string>? InstancePoolId { get; set; }

        /// <summary>
        /// The string used to prefix the managed instances name (default `pool`).
        /// </summary>
        [Input("instancePrefix")]
        public Input<string>? InstancePrefix { get; set; }

        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

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
        /// The SKS node pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateNetworkIds")]
        private InputList<string>? _privateNetworkIds;

        /// <summary>
        /// A list of exoscale*private*network (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> PrivateNetworkIds
        {
            get => _privateNetworkIds ?? (_privateNetworkIds = new InputList<string>());
            set => _privateNetworkIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of exoscale*security*group (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The current pool state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
        /// </summary>
        [Input("storageLvm")]
        public Input<bool>? StorageLvm { get; set; }

        [Input("taints")]
        private InputMap<string>? _taints;

        /// <summary>
        /// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) (`&lt;value&gt;:&lt;effect&gt;`).
        /// </summary>
        public InputMap<string> Taints
        {
            get => _taints ?? (_taints = new InputMap<string>());
            set => _taints = value;
        }

        /// <summary>
        /// The managed instances template ID.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// The managed instances version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public SKSNodepoolState()
        {
        }
        public static new SKSNodepoolState Empty => new SKSNodepoolState();
    }
}
