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
    public sealed class ElasticIpHealthcheck
    {
        /// <summary>
        /// The healthcheck interval (seconds; must be between `5` and `300`; default: `10`).
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// The healthcheck mode (`tcp`, `http` or `https`; may only be set at creation time).
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The healthcheck target port (must be between `1` and `65535`).
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The number of failed healthcheck attempts before considering the target unhealthy (must be between `1` and `20`; default: `2`).
        /// </summary>
        public readonly int? StrikesFail;
        /// <summary>
        /// The number of successful healthcheck attempts before considering the target healthy (must be between `1` and `20`; default: `3`).
        /// </summary>
        public readonly int? StrikesOk;
        /// <summary>
        /// The time before considering a healthcheck probing failed (seconds; must be between `2` and `60`; default: `3`).
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// Disable TLS certificate verification for healthcheck in `https` mode (boolean; default: `false`).
        /// </summary>
        public readonly bool? TlsSkipVerify;
        /// <summary>
        /// The healthcheck server name to present with SNI in `https` mode.
        /// </summary>
        public readonly string? TlsSni;
        /// <summary>
        /// The healthcheck target URI (required in `http(s)` modes).
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private ElasticIpHealthcheck(
            int? interval,

            string mode,

            int port,

            int? strikesFail,

            int? strikesOk,

            int? timeout,

            bool? tlsSkipVerify,

            string? tlsSni,

            string? uri)
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