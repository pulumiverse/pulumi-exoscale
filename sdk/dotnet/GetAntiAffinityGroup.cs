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
    public static class GetAntiAffinityGroup
    {
        public static Task<GetAntiAffinityGroupResult> InvokeAsync(GetAntiAffinityGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAntiAffinityGroupResult>("exoscale:index/getAntiAffinityGroup:getAntiAffinityGroup", args ?? new GetAntiAffinityGroupArgs(), options.WithDefaults());

        public static Output<GetAntiAffinityGroupResult> Invoke(GetAntiAffinityGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAntiAffinityGroupResult>("exoscale:index/getAntiAffinityGroup:getAntiAffinityGroup", args ?? new GetAntiAffinityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAntiAffinityGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The anti-affinity group ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The group name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetAntiAffinityGroupArgs()
        {
        }
        public static new GetAntiAffinityGroupArgs Empty => new GetAntiAffinityGroupArgs();
    }

    public sealed class GetAntiAffinityGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The anti-affinity group ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The group name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetAntiAffinityGroupInvokeArgs()
        {
        }
        public static new GetAntiAffinityGroupInvokeArgs Empty => new GetAntiAffinityGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetAntiAffinityGroupResult
    {
        public readonly string? Id;
        /// <summary>
        /// The list of attached exoscale.ComputeInstance (IDs).
        /// </summary>
        public readonly ImmutableArray<string> Instances;
        public readonly string? Name;

        [OutputConstructor]
        private GetAntiAffinityGroupResult(
            string? id,

            ImmutableArray<string> instances,

            string? name)
        {
            Id = id;
            Instances = instances;
            Name = name;
        }
    }
}
