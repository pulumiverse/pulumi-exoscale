// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Outputs
{

    [OutputType]
    public sealed class InstancePoolInstance
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The instance (main network interface) IPv6 address.
        /// </summary>
        public readonly string? Ipv6Address;
        /// <summary>
        /// The instance name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The instance (main network interface) IPv4 address.
        /// </summary>
        public readonly string? PublicIpAddress;

        [OutputConstructor]
        private InstancePoolInstance(
            string? id,

            string? ipv6Address,

            string? name,

            string? publicIpAddress)
        {
            Id = id;
            Ipv6Address = ipv6Address;
            Name = name;
            PublicIpAddress = publicIpAddress;
        }
    }
}
