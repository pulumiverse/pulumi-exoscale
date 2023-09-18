// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Inputs
{

    public sealed class DatabaseOpensearchGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OpenSearch Dashboards settings
        /// </summary>
        [Input("dashboards")]
        public Input<Inputs.DatabaseOpensearchDashboardsGetArgs>? Dashboards { get; set; }

        /// <summary>
        /// ❗ Service name
        /// </summary>
        [Input("forkFromService")]
        public Input<string>? ForkFromService { get; set; }

        [Input("indexPatterns")]
        private InputList<Inputs.DatabaseOpensearchIndexPatternGetArgs>? _indexPatterns;

        /// <summary>
        /// (can be used multiple times) Allows you to create glob style patterns and set a max number of indexes matching this pattern you want to keep. Creating indexes exceeding this value will cause the oldest one to get deleted. You could for example create a pattern looking like 'logs.?' and then create index logs.1, logs.2 etc, it will delete logs.1 once you create logs.6. Do note 'logs.?' does not apply to logs.10. Note: Setting max*index*count to 0 will do nothing and the pattern gets ignored.
        /// </summary>
        public InputList<Inputs.DatabaseOpensearchIndexPatternGetArgs> IndexPatterns
        {
            get => _indexPatterns ?? (_indexPatterns = new InputList<Inputs.DatabaseOpensearchIndexPatternGetArgs>());
            set => _indexPatterns = value;
        }

        /// <summary>
        /// Template settings for all new indexes
        /// </summary>
        [Input("indexTemplate")]
        public Input<Inputs.DatabaseOpensearchIndexTemplateGetArgs>? IndexTemplate { get; set; }

        [Input("ipFilters")]
        private InputList<string>? _ipFilters;

        /// <summary>
        /// Allow incoming connections from this list of CIDR address block, e.g. `["10.20.0.0/16"]
        /// </summary>
        public InputList<string> IpFilters
        {
            get => _ipFilters ?? (_ipFilters = new InputList<string>());
            set => _ipFilters = value;
        }

        /// <summary>
        /// Aiven automation resets index.refresh_interval to default value for every index to be sure that indices are always visible to search. If it doesn't fit your case, you can disable this by setting up this flag to true.
        /// </summary>
        [Input("keepIndexRefreshInterval")]
        public Input<bool>? KeepIndexRefreshInterval { get; set; }

        /// <summary>
        /// Maximum number of indexes to keep (Minimum value is `0`)
        /// </summary>
        [Input("maxIndexCount")]
        public Input<int>? MaxIndexCount { get; set; }

        /// <summary>
        /// ❗ Name of a backup to recover from
        /// </summary>
        [Input("recoveryBackupName")]
        public Input<string>? RecoveryBackupName { get; set; }

        /// <summary>
        /// OpenSearch-specific settings, in json. e.g.`jsonencode({thread_pool_search_size: 64})`. Use `exo x get-dbaas-settings-opensearch` to get a list of available settings.
        /// </summary>
        [Input("settings")]
        public Input<string>? Settings { get; set; }

        /// <summary>
        /// ❗ OpenSearch major version (`exo dbaas type show opensearch` for reference)
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public DatabaseOpensearchGetArgs()
        {
        }
        public static new DatabaseOpensearchGetArgs Empty => new DatabaseOpensearchGetArgs();
    }
}
