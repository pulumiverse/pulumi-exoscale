// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class DbaasMysql
    {
        /// <summary>
        /// A custom administrator account password (may only be set at creation time).
        /// </summary>
        public readonly string? AdminPassword;
        /// <summary>
        /// A custom administrator account username (may only be set at creation time).
        /// </summary>
        public readonly string? AdminUsername;
        /// <summary>
        /// The automated backup schedule (`HH:MM`).
        /// </summary>
        public readonly string? BackupSchedule;
        /// <summary>
        /// A list of CIDR blocks to allow incoming connections from.
        /// </summary>
        public readonly ImmutableArray<string> IpFilters;
        /// <summary>
        /// MySQL configuration settings in JSON format (`exo dbaas type show mysql --settings=mysql` for reference).
        /// </summary>
        public readonly string? MysqlSettings;
        /// <summary>
        /// MySQL major version (`exo dbaas type show mysql` for reference; may only be set at creation time).
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private DbaasMysql(
            string? adminPassword,

            string? adminUsername,

            string? backupSchedule,

            ImmutableArray<string> ipFilters,

            string? mysqlSettings,

            string? version)
        {
            AdminPassword = adminPassword;
            AdminUsername = adminUsername;
            BackupSchedule = backupSchedule;
            IpFilters = ipFilters;
            MysqlSettings = mysqlSettings;
            Version = version;
        }
    }
}
