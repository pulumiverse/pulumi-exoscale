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
    public static class GetIAMAPIKey
    {
        /// <summary>
        /// Fetch Exoscale [IAM](https://community.exoscale.com/documentation/iam/) API Key.
        /// 
        /// Corresponding resource: exoscale_iam_role.
        /// </summary>
        public static Task<GetIAMAPIKeyResult> InvokeAsync(GetIAMAPIKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIAMAPIKeyResult>("exoscale:index/getIAMAPIKey:getIAMAPIKey", args ?? new GetIAMAPIKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Fetch Exoscale [IAM](https://community.exoscale.com/documentation/iam/) API Key.
        /// 
        /// Corresponding resource: exoscale_iam_role.
        /// </summary>
        public static Output<GetIAMAPIKeyResult> Invoke(GetIAMAPIKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIAMAPIKeyResult>("exoscale:index/getIAMAPIKey:getIAMAPIKey", args ?? new GetIAMAPIKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIAMAPIKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IAM API Key to match.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("timeouts")]
        public Inputs.GetIAMAPIKeyTimeoutsArgs? Timeouts { get; set; }

        public GetIAMAPIKeyArgs()
        {
        }
        public static new GetIAMAPIKeyArgs Empty => new GetIAMAPIKeyArgs();
    }

    public sealed class GetIAMAPIKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IAM API Key to match.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("timeouts")]
        public Input<Inputs.GetIAMAPIKeyTimeoutsInputArgs>? Timeouts { get; set; }

        public GetIAMAPIKeyInvokeArgs()
        {
        }
        public static new GetIAMAPIKeyInvokeArgs Empty => new GetIAMAPIKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetIAMAPIKeyResult
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IAM API Key to match.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// IAM API Key name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IAM API Key role ID.
        /// </summary>
        public readonly string RoleId;
        public readonly Outputs.GetIAMAPIKeyTimeoutsResult? Timeouts;

        [OutputConstructor]
        private GetIAMAPIKeyResult(
            string id,

            string key,

            string name,

            string roleId,

            Outputs.GetIAMAPIKeyTimeoutsResult? timeouts)
        {
            Id = id;
            Key = key;
            Name = name;
            RoleId = roleId;
            Timeouts = timeouts;
        }
    }
}
