// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumiverse.Exoscale
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("exoscale");

        private static readonly __Value<string?> _computeEndpoint = new __Value<string?>(() => __config.Get("computeEndpoint"));
        /// <summary>
        /// Exoscale CloudStack API endpoint (by default: https://api.exoscale.com/v1)
        /// </summary>
        public static string? ComputeEndpoint
        {
            get => _computeEndpoint.Get();
            set => _computeEndpoint.Set(value);
        }

        private static readonly __Value<string?> _config = new __Value<string?>(() => __config.Get("config"));
        /// <summary>
        /// CloudStack ini configuration filename (by default: cloudstack.ini)
        /// </summary>
        public static string? CloudStackIniConfig
        {
            get => _config.Get();
            set => _config.Set(value);
        }

        private static readonly __Value<int?> _delay = new __Value<int?>(() => __config.GetInt32("delay"));
        public static int? Delay
        {
            get => _delay.Get();
            set => _delay.Set(value);
        }

        private static readonly __Value<string?> _dnsEndpoint = new __Value<string?>(() => __config.Get("dnsEndpoint"));
        /// <summary>
        /// Exoscale DNS API endpoint (by default: https://api.exoscale.com/dns)
        /// </summary>
        public static string? DnsEndpoint
        {
            get => _dnsEndpoint.Get();
            set => _dnsEndpoint.Set(value);
        }

        private static readonly __Value<string?> _environment = new __Value<string?>(() => __config.Get("environment"));
        public static string? Environment
        {
            get => _environment.Get();
            set => _environment.Set(value);
        }

        private static readonly __Value<bool?> _gzipUserData = new __Value<bool?>(() => __config.GetBoolean("gzipUserData"));
        /// <summary>
        /// Defines if the user-data of compute instances should be gzipped (by default: true)
        /// </summary>
        public static bool? GzipUserData
        {
            get => _gzipUserData.Get();
            set => _gzipUserData.Set(value);
        }

        private static readonly __Value<string?> _key = new __Value<string?>(() => __config.Get("key"));
        /// <summary>
        /// Exoscale API key
        /// </summary>
        public static string? Key
        {
            get => _key.Get();
            set => _key.Set(value);
        }

        private static readonly __Value<string?> _profile = new __Value<string?>(() => __config.Get("profile"));
        public static string? Profile
        {
            get => _profile.Get();
            set => _profile.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region"));
        /// <summary>
        /// CloudStack ini configuration section name (by default: cloudstack)
        /// </summary>
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _secret = new __Value<string?>(() => __config.Get("secret"));
        /// <summary>
        /// Exoscale API secret
        /// </summary>
        public static string? Secret
        {
            get => _secret.Get();
            set => _secret.Set(value);
        }

        private static readonly __Value<int?> _timeout = new __Value<int?>(() => __config.GetInt32("timeout"));
        /// <summary>
        /// Timeout in seconds for waiting on compute resources to become available (by default: 300)
        /// </summary>
        public static int? Timeout
        {
            get => _timeout.Get();
            set => _timeout.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

    }
}
