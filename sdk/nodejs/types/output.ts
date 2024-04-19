// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ComputeInstanceNetworkInterface {
    /**
     * The IPv4 address to request as static DHCP lease if the network interface is attached to a *managed* private network.
     */
    ipAddress: string;
    /**
     * The exoscale*private*network (ID) to attach to the instance.
     */
    networkId: string;
}

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
    /**
     * Enable or disable OpenSearch Dashboards (default: true).
     */
    enabled?: boolean;
    /**
     * Limits the maximum amount of memory (in MiB) the OpenSearch Dashboards process can use. This sets the max*old*space_size option of the nodejs running the OpenSearch Dashboards. Note: the memory reserved by OpenSearch Dashboards is not available for OpenSearch. (default: 128).
     */
    maxOldSpaceSize?: number;
    /**
     * Timeout in milliseconds for requests made by OpenSearch Dashboards towards OpenSearch (default: 30000)
     */
    requestTimeout?: number;
}

export interface DatabaseOpensearchIndexPattern {
    /**
     * Maximum number of indexes to keep before deleting the oldest one (Minimum value is `0`)
     */
    maxIndexCount?: number;
    /**
     * fnmatch pattern
     */
    pattern?: string;
    /**
     * `alphabetical` or `creationDate`.
     */
    sortingAlgorithm?: string;
}

export interface DatabaseOpensearchIndexTemplate {
    /**
     * The maximum number of nested JSON objects that a single document can contain across all nested types. This limit helps to prevent out of memory errors when a document contains too many nested objects. (Default is 10000. Minimum value is `0`, maximum value is `100000`.)
     */
    mappingNestedObjectsLimit?: number;
    /**
     * The number of replicas each primary shard has. (Minimum value is `0`, maximum value is `29`)
     */
    numberOfReplicas?: number;
    /**
     * The number of primary shards that an index should have. (Minimum value is `1`, maximum value is `1024`.)
     */
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

export interface ElasticIpHealthcheck {
    /**
     * The healthcheck interval (seconds; must be between `5` and `300`; default: `10`).
     */
    interval?: number;
    /**
     * The healthcheck mode (`tcp`, `http` or `https`; may only be set at creation time).
     */
    mode: string;
    /**
     * The healthcheck target port (must be between `1` and `65535`).
     */
    port: number;
    /**
     * The number of failed healthcheck attempts before considering the target unhealthy (must be between `1` and `20`; default: `2`).
     */
    strikesFail?: number;
    /**
     * The number of successful healthcheck attempts before considering the target healthy (must be between `1` and `20`; default: `3`).
     */
    strikesOk?: number;
    /**
     * The time before considering a healthcheck probing failed (seconds; must be between `2` and `60`; default: `3`).
     */
    timeout?: number;
    /**
     * Disable TLS certificate verification for healthcheck in `https` mode (boolean; default: `false`).
     */
    tlsSkipVerify?: boolean;
    /**
     * The healthcheck server name to present with SNI in `https` mode.
     */
    tlsSni?: string;
    /**
     * The healthcheck target URI (required in `http(s)` modes).
     */
    uri?: string;
}

export interface GetComputeInstanceListInstance {
    /**
     * The list of attached exoscale.AntiAffinityGroup (IDs).
     */
    antiAffinityGroupIds: string[];
    /**
     * The compute instance creation date.
     */
    createdAt: string;
    /**
     * A deploy target ID.
     */
    deployTargetId: string;
    /**
     * The instance disk size (GiB).
     */
    diskSize: number;
    /**
     * The list of attached exoscale.ElasticIp (IDs).
     */
    elasticIpIds: string[];
    /**
     * The compute instance ID to match (conflicts with `name`).
     */
    id?: string;
    /**
     * Whether IPv6 is enabled on the instance.
     */
    ipv6: boolean;
    /**
     * The instance (main network interface) IPv6 address (if enabled).
     */
    ipv6Address: string;
    /**
     * A map of key/value labels.
     */
    labels: {[key: string]: string};
    /**
     * The instance manager ID, if any.
     */
    managerId: string;
    /**
     * The instance manager type (instance pool, SKS node pool, etc.), if any.
     */
    managerType: string;
    /**
     * The instance name to match (conflicts with `id`).
     */
    name?: string;
    /**
     * The list of attached exoscale.PrivateNetwork (IDs).
     */
    privateNetworkIds: string[];
    /**
     * The instance (main network interface) IPv4 address.
     */
    publicIpAddress: string;
    /**
     * Domain name for reverse DNS record.
     */
    reverseDns: string;
    /**
     * The list of attached exoscale.SecurityGroup (IDs).
     */
    securityGroupIds: string[];
    /**
     * The exoscale.SshKey (name) authorized on the instance.
     */
    sshKey: string;
    /**
     * The instance state.
     */
    state: string;
    /**
     * The instance exoscale.getTemplate ID.
     */
    templateId: string;
    /**
     * The instance type.
     */
    type: string;
    /**
     * The instance [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
     */
    userData: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

export interface GetDatabaseUriTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface GetDomainRecordFilter {
    /**
     * A regular expression to match the record content.
     */
    contentRegex?: string;
    /**
     * The record ID to match.
     */
    id?: string;
    /**
     * The domain record name to match.
     */
    name?: string;
    /**
     * The record type to match.
     */
    recordType?: string;
}

export interface GetDomainRecordRecord {
    /**
     * Content of the Record
     */
    content?: string;
    /**
     * Domain of the Record
     */
    domain?: string;
    /**
     * ID of the Record
     */
    id?: string;
    /**
     * Name of the Record
     */
    name?: string;
    /**
     * Priority of the Record
     */
    prio?: number;
    /**
     * Type of the Record
     */
    recordType?: string;
    /**
     * TTL of the Record
     */
    ttl?: number;
}

export interface GetElasticIpHealthcheck {
    /**
     * The healthcheck interval in seconds.
     */
    interval: number;
    /**
     * The healthcheck mode.
     */
    mode: string;
    /**
     * The healthcheck target port.
     */
    port: number;
    /**
     * The number of failed healthcheck attempts before considering the target unhealthy.
     */
    strikesFail: number;
    /**
     * The number of successful healthcheck attempts before considering the target healthy.
     */
    strikesOk: number;
    /**
     * The time in seconds before considering a healthcheck probing failed.
     */
    timeout: number;
    /**
     * Disable TLS certificate verification for healthcheck in `https` mode.
     */
    tlsSkipVerify: boolean;
    /**
     * The healthcheck server name to present with SNI in `https` mode.
     */
    tlsSni: string;
    /**
     * The healthcheck URI.
     */
    uri: string;
}

export interface GetIamApiKeyTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface GetIamOrgPolicyServices {
    /**
     * List of IAM service rules (if type is `rules`).
     */
    rules: outputs.GetIamOrgPolicyServicesRule[];
    /**
     * Service type (`rules`, `allow`, or `deny`).
     */
    type: string;
}

export interface GetIamOrgPolicyServicesRule {
    /**
     * IAM policy rule action (`allow` or `deny`).
     */
    action: string;
    /**
     * IAM policy rule expression.
     */
    expression: string;
    /**
     * @deprecated This field is no longer suported.
     */
    resources: string[];
}

export interface GetIamOrgPolicyTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface GetIamRolePolicy {
    /**
     * Default service strategy (`allow` or `deny`).
     */
    defaultServiceStrategy: string;
    /**
     * IAM policy services.
     */
    services: {[key: string]: outputs.GetIamRolePolicyServices};
}

export interface GetIamRolePolicyServices {
    /**
     * List of IAM service rules (if type is `rules`).
     */
    rules: outputs.GetIamRolePolicyServicesRule[];
    /**
     * Service type (`rules`, `allow`, or `deny`).
     */
    type: string;
}

export interface GetIamRolePolicyServicesRule {
    /**
     * IAM policy rule action (`allow` or `deny`).
     */
    action: string;
    /**
     * IAM policy rule expression.
     */
    expression: string;
    /**
     * @deprecated This field is no longer suported.
     */
    resources: string[];
}

export interface GetIamRoleTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface GetInstancePoolInstance {
    /**
     * The compute instance ID.
     */
    id?: string;
    /**
     * The instance (main network interface) IPv6 address.
     */
    ipv6Address: string;
    /**
     * The instance name.
     */
    name?: string;
    /**
     * The instance (main network interface) IPv4 address.
     */
    publicIpAddress: string;
}

export interface GetInstancePoolListPool {
    /**
     * The list of attached exoscale.AntiAffinityGroup (IDs).
     */
    affinityGroupIds: string[];
    /**
     * The deploy target ID.
     */
    deployTargetId: string;
    /**
     * The instance pool description.
     */
    description: string;
    /**
     * The managed instances disk size.
     */
    diskSize: number;
    /**
     * The list of attached exoscale.ElasticIp (IDs).
     */
    elasticIpIds: string[];
    /**
     * The instance pool ID to match (conflicts with `name`).
     */
    id?: string;
    /**
     * The string used to prefix the managed instances name.
     */
    instancePrefix: string;
    /**
     * The managed instances type.
     */
    instanceType: string;
    /**
     * The list of managed instances. Structure is documented below.
     */
    instances: outputs.GetInstancePoolListPoolInstance[];
    /**
     * Whether IPv6 is enabled on managed instances.
     */
    ipv6: boolean;
    /**
     * The exoscale.SshKey (name) authorized on the managed instances.
     */
    keyPair: string;
    /**
     * A map of key/value labels.
     */
    labels?: {[key: string]: string};
    /**
     * The pool name to match (conflicts with `id`).
     */
    name?: string;
    /**
     * The list of attached exoscale.PrivateNetwork (IDs).
     */
    networkIds: string[];
    /**
     * The list of attached exoscale.SecurityGroup (IDs).
     */
    securityGroupIds: string[];
    /**
     * The number managed instances.
     */
    size: number;
    /**
     * The pool state.
     */
    state: string;
    /**
     * The managed instances exoscale.getTemplate ID.
     */
    templateId: string;
    /**
     * [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
     */
    userData: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

export interface GetInstancePoolListPoolInstance {
    /**
     * The compute instance ID.
     */
    id?: string;
    /**
     * The instance (main network interface) IPv6 address.
     */
    ipv6Address: string;
    /**
     * The instance name.
     */
    name?: string;
    /**
     * The instance (main network interface) IPv4 address.
     */
    publicIpAddress: string;
}

export interface GetNlbServiceListService {
    /**
     * NLB service description.
     */
    description: string;
    healthcheck: outputs.GetNlbServiceListServiceHealthcheck;
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

export interface GetNlbServiceListServiceHealthcheck {
    interval: number;
    mode: string;
    port: number;
    retries: number;
    timeout: number;
    tlsSni: string;
    uri: string;
}

export interface GetNlbServiceListTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface GetSksClusterListCluster {
    /**
     * @deprecated This attribute has been replaced by `exoscaleCcm`/`metricsServer` attributes, it will be removed in a future release.
     */
    addons: string[];
    /**
     * The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
     */
    aggregationCa: string;
    /**
     * Enable automatic upgrading of the control plane version.
     */
    autoUpgrade?: boolean;
    /**
     * The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
     */
    cni?: string;
    /**
     * The CA certificate (in PEM format) for TLS communications between control plane components.
     */
    controlPlaneCa: string;
    /**
     * The cluster creation date.
     */
    createdAt: string;
    /**
     * A free-form text describing the cluster.
     */
    description?: string;
    /**
     * The cluster API endpoint.
     */
    endpoint: string;
    /**
     * Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    exoscaleCcm?: boolean;
    /**
     * Deploy the Exoscale [Container Storage Interface](https://github.com/exoscale/exoscale-csi-driver/) on worker nodes (boolean; default: `false`; may only be set at creation time).
     */
    exoscaleCsi?: boolean;
    id?: string;
    /**
     * The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
     */
    kubeletCa: string;
    /**
     * A map of key/value labels.
     */
    labels?: {[key: string]: string};
    /**
     * Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    metricsServer?: boolean;
    name?: string;
    /**
     * The list of exoscale.SksNodepool (IDs) attached to the cluster.
     */
    nodepools: string[];
    /**
     * An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
     */
    oidc: outputs.GetSksClusterListClusterOidc;
    /**
     * The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
     */
    serviceLevel?: string;
    /**
     * The cluster state.
     */
    state: string;
    /**
     * The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
     */
    version: string;
    zone: string;
}

export interface GetSksClusterListClusterOidc {
    /**
     * The OpenID client ID.
     */
    clientId: string;
    /**
     * An OpenID JWT claim to use as the user's group.
     */
    groupsClaim?: string;
    /**
     * An OpenID prefix prepended to group claims.
     */
    groupsPrefix?: string;
    /**
     * The OpenID provider URL.
     */
    issuerUrl: string;
    /**
     * A map of key/value pairs that describes a required claim in the OpenID Token.
     */
    requiredClaim?: {[key: string]: string};
    /**
     * An OpenID JWT claim to use as the user name.
     */
    usernameClaim?: string;
    /**
     * An OpenID prefix prepended to username claims.
     */
    usernamePrefix?: string;
}

export interface GetSksClusterOidc {
    /**
     * The OpenID client ID.
     */
    clientId: string;
    /**
     * An OpenID JWT claim to use as the user's group.
     */
    groupsClaim?: string;
    /**
     * An OpenID prefix prepended to group claims.
     */
    groupsPrefix?: string;
    /**
     * The OpenID provider URL.
     */
    issuerUrl: string;
    /**
     * A map of key/value pairs that describes a required claim in the OpenID Token.
     */
    requiredClaim?: {[key: string]: string};
    /**
     * An OpenID JWT claim to use as the user name.
     */
    usernameClaim?: string;
    /**
     * An OpenID prefix prepended to username claims.
     */
    usernamePrefix?: string;
}

export interface GetSksNodepoolListNodepool {
    /**
     * A list of exoscale.AntiAffinityGroup (IDs) to be attached to the managed instances.
     */
    antiAffinityGroupIds?: string[];
    clusterId: string;
    /**
     * The pool creation date.
     */
    createdAt: string;
    /**
     * A deploy target ID.
     */
    deployTargetId?: string;
    /**
     * A free-form text describing the pool.
     */
    description?: string;
    /**
     * The managed instances disk size (GiB; default: `50`).
     */
    diskSize?: number;
    id?: string;
    /**
     * The underlying exoscale.InstancePool ID.
     */
    instancePoolId: string;
    /**
     * The string used to prefix the managed instances name (default `pool`).
     */
    instancePrefix?: string;
    /**
     * The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
     */
    instanceType?: string;
    /**
     * A map of key/value labels.
     */
    labels?: {[key: string]: string};
    name?: string;
    /**
     * A list of exoscale.PrivateNetwork (IDs) to be attached to the managed instances.
     */
    privateNetworkIds?: string[];
    /**
     * A list of exoscale.SecurityGroup (IDs) to be attached to the managed instances.
     */
    securityGroupIds?: string[];
    size?: number;
    /**
     * The current pool state.
     */
    state: string;
    /**
     * Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
     */
    storageLvm?: boolean;
    /**
     * A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) ('taints = { <key> = "<value>:<effect>" }').
     */
    taints?: {[key: string]: string};
    /**
     * The managed instances template ID.
     */
    templateId: string;
    /**
     * The managed instances version.
     */
    version: string;
    zone: string;
}

export interface IamApiKeyTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface IamOrgPolicyServices {
    /**
     * List of IAM service rules (if type is `rules`).
     */
    rules: outputs.IamOrgPolicyServicesRule[];
    /**
     * Service type (`rules`, `allow`, or `deny`).
     */
    type: string;
}

export interface IamOrgPolicyServicesRule {
    /**
     * IAM policy rule action (`allow` or `deny`).
     */
    action: string;
    /**
     * IAM policy rule expression.
     */
    expression: string;
    /**
     * @deprecated This field is not suported. Specify resources using CEL expressions.
     */
    resources: string[];
}

export interface IamOrgPolicyTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface IamRolePolicy {
    /**
     * Default service strategy (`allow` or `deny`).
     */
    defaultServiceStrategy: string;
    /**
     * IAM policy services.
     */
    services: {[key: string]: outputs.IamRolePolicyServices};
}

export interface IamRolePolicyServices {
    /**
     * List of IAM service rules (if type is `rules`).
     */
    rules: outputs.IamRolePolicyServicesRule[];
    /**
     * Service type (`rules`, `allow`, or `deny`).
     */
    type: string;
}

export interface IamRolePolicyServicesRule {
    /**
     * IAM policy rule action (`allow` or `deny`).
     */
    action: string;
    /**
     * IAM policy rule expression.
     */
    expression: string;
    /**
     * @deprecated This field is not suported. Specify resources using CEL expressions.
     */
    resources: string[];
}

export interface IamRoleTimeouts {
    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
     */
    read?: string;
}

export interface InstancePoolInstance {
    /**
     * The ID of this resource.
     */
    id?: string;
    /**
     * The instance (main network interface) IPv6 address.
     */
    ipv6Address: string;
    /**
     * The instance name.
     */
    name?: string;
    /**
     * The instance (main network interface) IPv4 address.
     */
    publicIpAddress: string;
}

export interface NlbServiceHealthcheck {
    /**
     * The healthcheck interval in seconds (default: `10`).
     */
    interval?: number;
    /**
     * The healthcheck mode (`tcp`|`http`|`https`; default: `tcp`).
     */
    mode?: string;
    /**
     * The NLB service (TCP/UDP) port.
     */
    port: number;
    /**
     * The healthcheck retries (default: `1`).
     */
    retries?: number;
    /**
     * The healthcheck timeout (seconds; default: `5`).
     */
    timeout?: number;
    /**
     * The healthcheck TLS SNI server name (only if `mode` is `https`).
     */
    tlsSni?: string;
    /**
     * The healthcheck URI (must be set only if `mode` is `http(s)`).
     */
    uri?: string;
}

export interface SksClusterOidc {
    /**
     * The OpenID client ID.
     */
    clientId: string;
    /**
     * An OpenID JWT claim to use as the user's group.
     */
    groupsClaim?: string;
    /**
     * An OpenID prefix prepended to group claims.
     */
    groupsPrefix?: string;
    /**
     * The OpenID provider URL.
     */
    issuerUrl: string;
    /**
     * A map of key/value pairs that describes a required claim in the OpenID Token.
     */
    requiredClaim?: {[key: string]: string};
    /**
     * An OpenID JWT claim to use as the user name.
     */
    usernameClaim?: string;
    /**
     * An OpenID prefix prepended to username claims.
     */
    usernamePrefix?: string;
}

