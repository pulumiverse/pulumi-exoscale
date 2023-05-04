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
    public static class GetPrivateNetwork
    {
        public static Task<GetPrivateNetworkResult> InvokeAsync(GetPrivateNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateNetworkResult>("exoscale:index/getPrivateNetwork:getPrivateNetwork", args ?? new GetPrivateNetworkArgs(), options.WithDefaults());

        public static Output<GetPrivateNetworkResult> Invoke(GetPrivateNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateNetworkResult>("exoscale:index/getPrivateNetwork:getPrivateNetwork", args ?? new GetPrivateNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The private network description.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The private network ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The network name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetPrivateNetworkArgs()
        {
        }
        public static new GetPrivateNetworkArgs Empty => new GetPrivateNetworkArgs();
    }

    public sealed class GetPrivateNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The private network description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The private network ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The network name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetPrivateNetworkInvokeArgs()
        {
        }
        public static new GetPrivateNetworkInvokeArgs Empty => new GetPrivateNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivateNetworkResult
    {
        /// <summary>
        /// The private network description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        /// </summary>
        public readonly string EndIp;
        /// <summary>
        /// The private network ID to match (conflicts with `name`).
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The network name to match (conflicts with `id`).
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The network mask defining the IPv4 network allowed for static leases.
        /// </summary>
        public readonly string Netmask;
        /// <summary>
        /// The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        /// </summary>
        public readonly string StartIp;
        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetPrivateNetworkResult(
            string? description,

            string endIp,

            string? id,

            string? name,

            string netmask,

            string startIp,

            string zone)
        {
            Description = description;
            EndIp = endIp;
            Id = id;
            Name = name;
            Netmask = netmask;
            StartIp = startIp;
            Zone = zone;
        }
    }
}
