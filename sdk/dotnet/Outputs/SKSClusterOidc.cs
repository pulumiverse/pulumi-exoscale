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
    public sealed class SKSClusterOidc
    {
        public readonly string ClientId;
        public readonly string? GroupsClaim;
        public readonly string? GroupsPrefix;
        public readonly string IssuerUrl;
        public readonly ImmutableDictionary<string, string>? RequiredClaim;
        public readonly string? UsernameClaim;
        public readonly string? UsernamePrefix;

        [OutputConstructor]
        private SKSClusterOidc(
            string clientId,

            string? groupsClaim,

            string? groupsPrefix,

            string issuerUrl,

            ImmutableDictionary<string, string>? requiredClaim,

            string? usernameClaim,

            string? usernamePrefix)
        {
            ClientId = clientId;
            GroupsClaim = groupsClaim;
            GroupsPrefix = groupsPrefix;
            IssuerUrl = issuerUrl;
            RequiredClaim = requiredClaim;
            UsernameClaim = usernameClaim;
            UsernamePrefix = usernamePrefix;
        }
    }
}
