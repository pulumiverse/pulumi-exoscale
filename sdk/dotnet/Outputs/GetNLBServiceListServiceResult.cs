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
    public sealed class GetNLBServiceListServiceResult
    {
        /// <summary>
        /// NLB service description.
        /// </summary>
        public readonly string Description;
        public readonly Outputs.GetNLBServiceListServiceHealthcheckResult Healthcheck;
        /// <summary>
        /// NLB service ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The exoscale*instance*pool (ID) to forward traffic to.
        /// </summary>
        public readonly string InstancePoolId;
        /// <summary>
        /// NLB Service name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Port exposed on the NLB's public IP.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Network traffic protocol.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// NLB Service State.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The strategy (`round-robin`|`source-hash`).
        /// </summary>
        public readonly string Strategy;
        /// <summary>
        /// Port on which the network traffic will be forwarded to on the receiving instance.
        /// </summary>
        public readonly int TargetPort;

        [OutputConstructor]
        private GetNLBServiceListServiceResult(
            string description,

            Outputs.GetNLBServiceListServiceHealthcheckResult healthcheck,

            string id,

            string instancePoolId,

            string name,

            int port,

            string protocol,

            string state,

            string strategy,

            int targetPort)
        {
            Description = description;
            Healthcheck = healthcheck;
            Id = id;
            InstancePoolId = instancePoolId;
            Name = name;
            Port = port;
            Protocol = protocol;
            State = state;
            Strategy = strategy;
            TargetPort = targetPort;
        }
    }
}
