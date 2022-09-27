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
    public sealed class GetElasticIPHealthcheckResult
    {
        /// <summary>
        /// The healthcheck interval in seconds.
        /// </summary>
        public readonly int Interval;
        /// <summary>
        /// The healthcheck mode.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The healthcheck target port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The number of failed healthcheck attempts before considering the target unhealthy.
        /// </summary>
        public readonly int StrikesFail;
        /// <summary>
        /// The number of successful healthcheck attempts before considering the target healthy.
        /// </summary>
        public readonly int StrikesOk;
        /// <summary>
        /// The time in seconds before considering a healthcheck probing failed.
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// Disable TLS certificate verification for healthcheck in `https` mode.
        /// </summary>
        public readonly bool TlsSkipVerify;
        /// <summary>
        /// The healthcheck server name to present with SNI in `https` mode.
        /// </summary>
        public readonly string TlsSni;
        /// <summary>
        /// The healthcheck URI.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private GetElasticIPHealthcheckResult(
            int interval,

            string mode,

            int port,

            int strikesFail,

            int strikesOk,

            int timeout,

            bool tlsSkipVerify,

            string tlsSni,

            string uri)
        {
            Interval = interval;
            Mode = mode;
            Port = port;
            StrikesFail = strikesFail;
            StrikesOk = strikesOk;
            Timeout = timeout;
            TlsSkipVerify = tlsSkipVerify;
            TlsSni = tlsSni;
            Uri = uri;
        }
    }
}
