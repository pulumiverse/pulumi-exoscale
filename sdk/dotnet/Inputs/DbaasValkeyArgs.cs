// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Inputs
{

    public sealed class DbaasValkeyArgs : global::Pulumi.ResourceArgs
    {
        [Input("ipFilters")]
        private InputList<string>? _ipFilters;

        /// <summary>
        /// A list of CIDR blocks to allow incoming connections from.
        /// </summary>
        public InputList<string> IpFilters
        {
            get => _ipFilters ?? (_ipFilters = new InputList<string>());
            set => _ipFilters = value;
        }

        /// <summary>
        /// Valkey configuration settings in JSON format (`exo dbaas type show valkey --settings=valkey` for reference).
        /// </summary>
        [Input("valkeySettings")]
        public Input<string>? ValkeySettings { get; set; }

        public DbaasValkeyArgs()
        {
        }
        public static new DbaasValkeyArgs Empty => new DbaasValkeyArgs();
    }
}
