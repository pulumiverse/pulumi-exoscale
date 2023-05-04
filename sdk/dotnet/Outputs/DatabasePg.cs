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
    public sealed class DatabasePg
    {
        public readonly string? AdminPassword;
        public readonly string? AdminUsername;
        public readonly string? BackupSchedule;
        public readonly ImmutableArray<string> IpFilters;
        public readonly string? PgSettings;
        public readonly string? PgbouncerSettings;
        public readonly string? PglookoutSettings;
        public readonly string? Version;

        [OutputConstructor]
        private DatabasePg(
            string? adminPassword,

            string? adminUsername,

            string? backupSchedule,

            ImmutableArray<string> ipFilters,

            string? pgSettings,

            string? pgbouncerSettings,

            string? pglookoutSettings,

            string? version)
        {
            AdminPassword = adminPassword;
            AdminUsername = adminUsername;
            BackupSchedule = backupSchedule;
            IpFilters = ipFilters;
            PgSettings = pgSettings;
            PgbouncerSettings = pgbouncerSettings;
            PglookoutSettings = pglookoutSettings;
            Version = version;
        }
    }
}
