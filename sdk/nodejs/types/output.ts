// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface DatabaseGrafana {
    /**
     * Grafana configuration settings in JSON format (`exo dbaas type show grafana --settings=grafana` for reference).
     */
    grafanaSettings: string;
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
}

export interface DatabaseKafka {
    /**
     * Enable certificate-based authentication method.
     */
    enableCertAuth: boolean;
    /**
     * Enable Kafka Connect.
     */
    enableKafkaConnect: boolean;
    /**
     * Enable Kafka REST.
     */
    enableKafkaRest: boolean;
    /**
     * Enable SASL-based authentication method.
     */
    enableSaslAuth: boolean;
    /**
     * Enable Schema Registry.
     */
    enableSchemaRegistry: boolean;
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
    /**
     * Kafka Connect configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka-connect` for reference).
     */
    kafkaConnectSettings: string;
    /**
     * Kafka REST configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka-rest` for reference).
     */
    kafkaRestSettings: string;
    /**
     * Kafka configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka` for reference).
     */
    kafkaSettings: string;
    /**
     * Schema Registry configuration settings in JSON format (`exo dbaas type show kafka --settings=schema-registry` for reference)
     */
    schemaRegistrySettings: string;
    /**
     * Kafka major version (`exo dbaas type show kafka` for reference; may only be set at creation time).
     */
    version: string;
}

export interface DatabaseMysql {
    /**
     * A custom administrator account password (may only be set at creation time).
     */
    adminPassword?: string;
    /**
     * A custom administrator account username (may only be set at creation time).
     */
    adminUsername?: string;
    /**
     * The automated backup schedule (`HH:MM`).
     */
    backupSchedule: string;
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
    /**
     * MySQL configuration settings in JSON format (`exo dbaas type show mysql --settings=mysql` for reference).
     */
    mysqlSettings: string;
    /**
     * MySQL major version (`exo dbaas type show mysql` for reference; may only be set at creation time).
     */
    version: string;
}

export interface DatabaseOpensearch {
    /**
     * OpenSearch Dashboards settings
     */
    dashboards?: outputs.DatabaseOpensearchDashboards;
    /**
     * ❗ Service name
     */
    forkFromService?: string;
    /**
     * (can be used multiple times) Allows you to create glob style patterns and set a max number of indexes matching this pattern you want to keep. Creating indexes exceeding this value will cause the oldest one to get deleted. You could for example create a pattern looking like 'logs.?' and then create index logs.1, logs.2 etc, it will delete logs.1 once you create logs.6. Do note 'logs.?' does not apply to logs.10. Note: Setting max*index*count to 0 will do nothing and the pattern gets ignored.
     */
    indexPatterns?: outputs.DatabaseOpensearchIndexPattern[];
    /**
     * Template settings for all new indexes
     */
    indexTemplate?: outputs.DatabaseOpensearchIndexTemplate;
    /**
     * Allow incoming connections from this list of CIDR address block, e.g. `["10.20.0.0/16"]
     */
    ipFilters: string[];
    /**
     * Aiven automation resets index.refresh_interval to default value for every index to be sure that indices are always visible to search. If it doesn't fit your case, you can disable this by setting up this flag to true.
     */
    keepIndexRefreshInterval?: boolean;
    /**
     * Maximum number of indexes to keep (Minimum value is `0`)
     */
    maxIndexCount?: number;
    /**
     * ❗ Name of a backup to recover from
     */
    recoveryBackupName?: string;
    /**
     * OpenSearch-specific settings, in json. e.g.`jsonencode({thread_pool_search_size: 64})`. Use `exo x get-dbaas-settings-opensearch` to get a list of available settings.
     */
    settings: string;
    /**
     * ❗ OpenSearch major version (`exo dbaas type show opensearch` for reference)
     */
    version: string;
}

export interface DatabaseOpensearchDashboards {
    enabled?: boolean;
    maxOldSpaceSize?: number;
    requestTimeout?: number;
}

export interface DatabaseOpensearchIndexPattern {
    maxIndexCount?: number;
    pattern?: string;
    sortingAlgorithm?: string;
}

export interface DatabaseOpensearchIndexTemplate {
    mappingNestedObjectsLimit?: number;
    numberOfReplicas?: number;
    numberOfShards?: number;
}

export interface DatabasePg {
    /**
     * A custom administrator account password (may only be set at creation time).
     */
    adminPassword?: string;
    /**
     * A custom administrator account username (may only be set at creation time).
     */
    adminUsername?: string;
    /**
     * The automated backup schedule (`HH:MM`).
     */
    backupSchedule: string;
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
    /**
     * PostgreSQL configuration settings in JSON format (`exo dbaas type show pg --settings=pg` for reference).
     */
    pgSettings: string;
    /**
     * PgBouncer configuration settings in JSON format (`exo dbaas type show pg --settings=pgbouncer` for reference).
     */
    pgbouncerSettings: string;
    /**
     * pglookout configuration settings in JSON format (`exo dbaas type show pg --settings=pglookout` for reference).
     */
    pglookoutSettings: string;
    /**
     * PostgreSQL major version (`exo dbaas type show pg` for reference; may only be set at creation time).
     */
    version: string;
}

export interface DatabaseRedis {
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
    /**
     * Redis configuration settings in JSON format (`exo dbaas type show redis --settings=redis` for reference).
     */
    redisSettings: string;
}

export interface DatabaseTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
     */
    create?: string;
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
     */
    delete?: string;
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
     */
    update?: string;
}

export interface GetDatabaseURITimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface GetNLBServiceListService {
    /**
     * NLB service description.
     */
    description: string;
    healthcheck: outputs.GetNLBServiceListServiceHealthcheck;
    /**
     * NLB service ID.
     */
    id: string;
    /**
     * The exoscale*instance*pool (ID) to forward traffic to.
     */
    instancePoolId: string;
    /**
     * NLB Service name.
     */
    name: string;
    /**
     * Port exposed on the NLB's public IP.
     */
    port: number;
    /**
     * Network traffic protocol.
     */
    protocol: string;
    /**
     * NLB Service State.
     */
    state: string;
    /**
     * The strategy (`round-robin`|`source-hash`).
     */
    strategy: string;
    /**
     * Port on which the network traffic will be forwarded to on the receiving instance.
     */
    targetPort: number;
}

export interface GetNLBServiceListServiceHealthcheck {
    interval: number;
    mode: string;
    port: number;
    retries: number;
    timeout: number;
    tlsSni: string;
    uri: string;
}

export interface GetNLBServiceListTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

