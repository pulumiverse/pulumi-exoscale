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
    public sealed class DatabaseOpensearchIndexTemplate
    {
        /// <summary>
        /// The maximum number of nested JSON objects that a single document can contain across all nested types. This limit helps to prevent out of memory errors when a document contains too many nested objects. (Default is 10000. Minimum value is `0`, maximum value is `100000`.)
        /// </summary>
        public readonly int? MappingNestedObjectsLimit;
        /// <summary>
        /// The number of replicas each primary shard has. (Minimum value is `0`, maximum value is `29`)
        /// </summary>
        public readonly int? NumberOfReplicas;
        /// <summary>
        /// The number of primary shards that an index should have. (Minimum value is `1`, maximum value is `1024`.)
        /// </summary>
        public readonly int? NumberOfShards;

        [OutputConstructor]
        private DatabaseOpensearchIndexTemplate(
            int? mappingNestedObjectsLimit,

            int? numberOfReplicas,

            int? numberOfShards)
        {
            MappingNestedObjectsLimit = mappingNestedObjectsLimit;
            NumberOfReplicas = numberOfReplicas;
            NumberOfShards = numberOfShards;
        }
    }
}
