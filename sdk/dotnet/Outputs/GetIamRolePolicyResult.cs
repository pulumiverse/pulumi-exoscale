// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class GetIamRolePolicyResult
    {
        /// <summary>
        /// Default service strategy (`allow` or `deny`).
        /// </summary>
        public readonly string DefaultServiceStrategy;
        /// <summary>
        /// IAM policy services.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.GetIamRolePolicyServicesResult> Services;

        [OutputConstructor]
        private GetIamRolePolicyResult(
            string defaultServiceStrategy,

            ImmutableDictionary<string, Outputs.GetIamRolePolicyServicesResult> services)
        {
            DefaultServiceStrategy = defaultServiceStrategy;
            Services = services;
        }
    }
}
