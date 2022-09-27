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
    public static class GetComputeIPAddress
    {
        /// <summary>
        /// !&gt; **WARNING:** This data source is **DEPRECATED** and will be removed in the next major version. Please use exoscale.ElasticIP instead.
        /// </summary>
        public static Task<GetComputeIPAddressResult> InvokeAsync(GetComputeIPAddressArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetComputeIPAddressResult>("exoscale:index/getComputeIPAddress:getComputeIPAddress", args ?? new GetComputeIPAddressArgs(), options.WithDefaults());

        /// <summary>
        /// !&gt; **WARNING:** This data source is **DEPRECATED** and will be removed in the next major version. Please use exoscale.ElasticIP instead.
        /// </summary>
        public static Output<GetComputeIPAddressResult> Invoke(GetComputeIPAddressInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetComputeIPAddressResult>("exoscale:index/getComputeIPAddress:getComputeIPAddress", args ?? new GetComputeIPAddressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetComputeIPAddressArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The EIP description to match.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The Elastic IP (EIP) ID to match.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The EIP IPv4 address to match.
        /// </summary>
        [Input("ipAddress")]
        public string? IpAddress { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// The EIP tags to match.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Exoscale Zone name.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetComputeIPAddressArgs()
        {
        }
        public static new GetComputeIPAddressArgs Empty => new GetComputeIPAddressArgs();
    }

    public sealed class GetComputeIPAddressInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The EIP description to match.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Elastic IP (EIP) ID to match.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The EIP IPv4 address to match.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The EIP tags to match.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Exoscale Zone name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetComputeIPAddressInvokeArgs()
        {
        }
        public static new GetComputeIPAddressInvokeArgs Empty => new GetComputeIPAddressInvokeArgs();
    }


    [OutputType]
    public sealed class GetComputeIPAddressResult
    {
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? IpAddress;
        public readonly ImmutableDictionary<string, string>? Tags;
        public readonly string Zone;

        [OutputConstructor]
        private GetComputeIPAddressResult(
            string? description,

            string? id,

            string? ipAddress,

            ImmutableDictionary<string, string>? tags,

            string zone)
        {
            Description = description;
            Id = id;
            IpAddress = ipAddress;
            Tags = tags;
            Zone = zone;
        }
    }
}
