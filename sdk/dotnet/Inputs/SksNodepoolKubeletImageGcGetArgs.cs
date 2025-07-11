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

    public sealed class SksNodepoolKubeletImageGcGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The percent of disk usage after which image garbage collection is always run
        /// </summary>
        [Input("highThreshold")]
        public Input<int>? HighThreshold { get; set; }

        /// <summary>
        /// The percent of disk usage before which image garbage collection is never run
        /// </summary>
        [Input("lowThreshold")]
        public Input<int>? LowThreshold { get; set; }

        /// <summary>
        /// The minimum age for an unused image before it is garbage collected (k8s duration format, eg. 1h)
        /// </summary>
        [Input("minAge")]
        public Input<string>? MinAge { get; set; }

        public SksNodepoolKubeletImageGcGetArgs()
        {
        }
        public static new SksNodepoolKubeletImageGcGetArgs Empty => new SksNodepoolKubeletImageGcGetArgs();
    }
}
