// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale
{
    public static class GetDatabaseUri
    {
        /// <summary>
        /// ## Example Usage
        /// </summary>
        public static Task<GetDatabaseUriResult> InvokeAsync(GetDatabaseUriArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseUriResult>("exoscale:index/getDatabaseUri:getDatabaseUri", args ?? new GetDatabaseUriArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// </summary>
        public static Output<GetDatabaseUriResult> Invoke(GetDatabaseUriInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseUriResult>("exoscale:index/getDatabaseUri:getDatabaseUri", args ?? new GetDatabaseUriInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// </summary>
        public static Output<GetDatabaseUriResult> Invoke(GetDatabaseUriInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseUriResult>("exoscale:index/getDatabaseUri:getDatabaseUri", args ?? new GetDatabaseUriInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseUriArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of database service to match.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("timeouts")]
        public Inputs.GetDatabaseUriTimeoutsArgs? Timeouts { get; set; }

        /// <summary>
        /// The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `valkey`, `grafana`).
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        /// <summary>
        /// The Exoscale Zone name.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetDatabaseUriArgs()
        {
        }
        public static new GetDatabaseUriArgs Empty => new GetDatabaseUriArgs();
    }

    public sealed class GetDatabaseUriInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of database service to match.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("timeouts")]
        public Input<Inputs.GetDatabaseUriTimeoutsInputArgs>? Timeouts { get; set; }

        /// <summary>
        /// The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `valkey`, `grafana`).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The Exoscale Zone name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetDatabaseUriInvokeArgs()
        {
        }
        public static new GetDatabaseUriInvokeArgs Empty => new GetDatabaseUriInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseUriResult
    {
        /// <summary>
        /// Default database name
        /// </summary>
        public readonly string DbName;
        /// <summary>
        /// Database service hostname
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of database service to match.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Admin user password
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Database service port
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Database service connection schema
        /// </summary>
        public readonly string Schema;
        public readonly Outputs.GetDatabaseUriTimeoutsResult? Timeouts;
        /// <summary>
        /// The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `valkey`, `grafana`).
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Database service connection URI.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// Admin user username
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// The Exoscale Zone name.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetDatabaseUriResult(
            string dbName,

            string host,

            string id,

            string name,

            string password,

            int port,

            string schema,

            Outputs.GetDatabaseUriTimeoutsResult? timeouts,

            string type,

            string uri,

            string username,

            string zone)
        {
            DbName = dbName;
            Host = host;
            Id = id;
            Name = name;
            Password = password;
            Port = port;
            Schema = schema;
            Timeouts = timeouts;
            Type = type;
            Uri = uri;
            Username = username;
            Zone = zone;
        }
    }
}
