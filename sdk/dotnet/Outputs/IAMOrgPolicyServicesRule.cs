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
    public sealed class IAMOrgPolicyServicesRule
    {
        /// <summary>
        /// IAM policy rule action (`allow` or `deny`).
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// IAM policy rule expression.
        /// </summary>
        public readonly string? Expression;
        /// <summary>
        /// List of resources that IAM policy rule applies to.
        /// </summary>
        public readonly ImmutableArray<string> Resources;

        [OutputConstructor]
        private IAMOrgPolicyServicesRule(
            string? action,

            string? expression,

            ImmutableArray<string> resources)
        {
            Action = action;
            Expression = expression;
            Resources = resources;
        }
    }
}
