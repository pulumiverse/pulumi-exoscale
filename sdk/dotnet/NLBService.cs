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
    /// An existing NLB service may be imported by `&lt;nlb-ID&gt;/&lt;service-ID&gt;@&lt;zone&gt;`
    /// 
    /// ```sh
    ///  $ pulumi import exoscale:index/nLBService:NLBService \
    /// ```
    /// 
    ///  exoscale_nlb_service.my_nlb_service \
    /// 
    ///  f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524@ch-gva-2
    /// </summary>
    [ExoscaleResourceType("exoscale:index/nLBService:NLBService")]
    public partial class NLBService : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A free-form text describing the NLB service.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The service health checking configuration.
        /// </summary>
        [Output("healthchecks")]
        public Output<ImmutableArray<Outputs.NLBServiceHealthcheck>> Healthchecks { get; private set; } = null!;

        /// <summary>
        /// ❗ The exoscale*instance*pool (ID) to forward traffic to.
        /// </summary>
        [Output("instancePoolId")]
        public Output<string> InstancePoolId { get; private set; } = null!;

        /// <summary>
        /// The NLB service name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ❗ The parent exoscale.NLB ID.
        /// </summary>
        [Output("nlbId")]
        public Output<string> NlbId { get; private set; } = null!;

        /// <summary>
        /// The NLB service (TCP/UDP) port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The protocol (`tcp`|`udp`; default: `tcp`).
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
        /// </summary>
        [Output("strategy")]
        public Output<string?> Strategy { get; private set; } = null!;

        /// <summary>
        /// The (TCP/UDP) port to forward traffic to (on target instance pool members).
        /// </summary>
        [Output("targetPort")]
        public Output<int> TargetPort { get; private set; } = null!;

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a NLBService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NLBService(string name, NLBServiceArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/nLBService:NLBService", name, args ?? new NLBServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NLBService(string name, Input<string> id, NLBServiceState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/nLBService:NLBService", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NLBService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NLBService Get(string name, Input<string> id, NLBServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new NLBService(name, id, state, options);
        }
    }

    public sealed class NLBServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-form text describing the NLB service.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("healthchecks", required: true)]
        private InputList<Inputs.NLBServiceHealthcheckArgs>? _healthchecks;

        /// <summary>
        /// The service health checking configuration.
        /// </summary>
        public InputList<Inputs.NLBServiceHealthcheckArgs> Healthchecks
        {
            get => _healthchecks ?? (_healthchecks = new InputList<Inputs.NLBServiceHealthcheckArgs>());
            set => _healthchecks = value;
        }

        /// <summary>
        /// ❗ The exoscale*instance*pool (ID) to forward traffic to.
        /// </summary>
        [Input("instancePoolId", required: true)]
        public Input<string> InstancePoolId { get; set; } = null!;

        /// <summary>
        /// The NLB service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ❗ The parent exoscale.NLB ID.
        /// </summary>
        [Input("nlbId", required: true)]
        public Input<string> NlbId { get; set; } = null!;

        /// <summary>
        /// The NLB service (TCP/UDP) port.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol (`tcp`|`udp`; default: `tcp`).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        /// <summary>
        /// The (TCP/UDP) port to forward traffic to (on target instance pool members).
        /// </summary>
        [Input("targetPort", required: true)]
        public Input<int> TargetPort { get; set; } = null!;

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public NLBServiceArgs()
        {
        }
        public static new NLBServiceArgs Empty => new NLBServiceArgs();
    }

    public sealed class NLBServiceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-form text describing the NLB service.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("healthchecks")]
        private InputList<Inputs.NLBServiceHealthcheckGetArgs>? _healthchecks;

        /// <summary>
        /// The service health checking configuration.
        /// </summary>
        public InputList<Inputs.NLBServiceHealthcheckGetArgs> Healthchecks
        {
            get => _healthchecks ?? (_healthchecks = new InputList<Inputs.NLBServiceHealthcheckGetArgs>());
            set => _healthchecks = value;
        }

        /// <summary>
        /// ❗ The exoscale*instance*pool (ID) to forward traffic to.
        /// </summary>
        [Input("instancePoolId")]
        public Input<string>? InstancePoolId { get; set; }

        /// <summary>
        /// The NLB service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ❗ The parent exoscale.NLB ID.
        /// </summary>
        [Input("nlbId")]
        public Input<string>? NlbId { get; set; }

        /// <summary>
        /// The NLB service (TCP/UDP) port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol (`tcp`|`udp`; default: `tcp`).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        /// <summary>
        /// The (TCP/UDP) port to forward traffic to (on target instance pool members).
        /// </summary>
        [Input("targetPort")]
        public Input<int>? TargetPort { get; set; }

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public NLBServiceState()
        {
        }
        public static new NLBServiceState Empty => new NLBServiceState();
    }
}
