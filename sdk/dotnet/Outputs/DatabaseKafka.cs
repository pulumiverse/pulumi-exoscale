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
    public sealed class DatabaseKafka
    {
        /// <summary>
        /// Enable certificate-based authentication method.
        /// </summary>
        public readonly bool? EnableCertAuth;
        /// <summary>
        /// Enable Kafka Connect.
        /// </summary>
        public readonly bool? EnableKafkaConnect;
        /// <summary>
        /// Enable Kafka REST.
        /// </summary>
        public readonly bool? EnableKafkaRest;
        /// <summary>
        /// Enable SASL-based authentication method.
        /// </summary>
        public readonly bool? EnableSaslAuth;
        /// <summary>
        /// Enable Schema Registry.
        /// </summary>
        public readonly bool? EnableSchemaRegistry;
        /// <summary>
        /// A list of CIDR blocks to allow incoming connections from.
        /// </summary>
        public readonly ImmutableArray<string> IpFilters;
        /// <summary>
        /// Kafka Connect configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka-connect` for reference).
        /// </summary>
        public readonly string? KafkaConnectSettings;
        /// <summary>
        /// Kafka REST configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka-rest` for reference).
        /// </summary>
        public readonly string? KafkaRestSettings;
        /// <summary>
        /// Kafka configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka` for reference).
        /// </summary>
        public readonly string? KafkaSettings;
        /// <summary>
        /// Schema Registry configuration settings in JSON format (`exo dbaas type show kafka --settings=schema-registry` for reference)
        /// </summary>
        public readonly string? SchemaRegistrySettings;
        /// <summary>
        /// Kafka major version (`exo dbaas type show kafka` for reference; may only be set at creation time).
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private DatabaseKafka(
            bool? enableCertAuth,

            bool? enableKafkaConnect,

            bool? enableKafkaRest,

            bool? enableSaslAuth,

            bool? enableSchemaRegistry,

            ImmutableArray<string> ipFilters,

            string? kafkaConnectSettings,

            string? kafkaRestSettings,

            string? kafkaSettings,

            string? schemaRegistrySettings,

            string? version)
        {
            EnableCertAuth = enableCertAuth;
            EnableKafkaConnect = enableKafkaConnect;
            EnableKafkaRest = enableKafkaRest;
            EnableSaslAuth = enableSaslAuth;
            EnableSchemaRegistry = enableSchemaRegistry;
            IpFilters = ipFilters;
            KafkaConnectSettings = kafkaConnectSettings;
            KafkaRestSettings = kafkaRestSettings;
            KafkaSettings = kafkaSettings;
            SchemaRegistrySettings = schemaRegistrySettings;
            Version = version;
        }
    }
}
