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
    /// !&gt; **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use exoscale.ComputeInstance `network_interface` block instead.
    /// </summary>
    [ExoscaleResourceType("exoscale:index/nIC:NIC")]
    public partial class NIC : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The compute instance ID.
        /// </summary>
        [Output("computeId")]
        public Output<string> ComputeId { get; private set; } = null!;

        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The NIC MAC address.
        /// </summary>
        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        [Output("netmask")]
        public Output<string> Netmask { get; private set; } = null!;

        /// <summary>
        /// The private network ID.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;


        /// <summary>
        /// Create a NIC resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NIC(string name, NICArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/nIC:NIC", name, args ?? new NICArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NIC(string name, Input<string> id, NICState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/nIC:NIC", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NIC resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NIC Get(string name, Input<string> id, NICState? state = null, CustomResourceOptions? options = null)
        {
            return new NIC(name, id, state, options);
        }
    }

    public sealed class NICArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compute instance ID.
        /// </summary>
        [Input("computeId", required: true)]
        public Input<string> ComputeId { get; set; } = null!;

        /// <summary>
        /// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The private network ID.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        public NICArgs()
        {
        }
        public static new NICArgs Empty => new NICArgs();
    }

    public sealed class NICState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compute instance ID.
        /// </summary>
        [Input("computeId")]
        public Input<string>? ComputeId { get; set; }

        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The NIC MAC address.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        [Input("netmask")]
        public Input<string>? Netmask { get; set; }

        /// <summary>
        /// The private network ID.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        public NICState()
        {
        }
        public static new NICState Empty => new NICState();
    }
}
