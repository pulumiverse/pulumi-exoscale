// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSksCluster(args: GetSksClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetSksClusterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getSksCluster:getSksCluster", {
        "addons": args.addons,
        "aggregationCa": args.aggregationCa,
        "autoUpgrade": args.autoUpgrade,
        "cni": args.cni,
        "controlPlaneCa": args.controlPlaneCa,
        "createdAt": args.createdAt,
        "description": args.description,
        "enableKubeProxy": args.enableKubeProxy,
        "endpoint": args.endpoint,
        "exoscaleCcm": args.exoscaleCcm,
        "exoscaleCsi": args.exoscaleCsi,
        "featureGates": args.featureGates,
        "id": args.id,
        "kubeletCa": args.kubeletCa,
        "labels": args.labels,
        "metricsServer": args.metricsServer,
        "name": args.name,
        "nodepools": args.nodepools,
        "oidc": args.oidc,
        "serviceLevel": args.serviceLevel,
        "state": args.state,
        "version": args.version,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSksCluster.
 */
export interface GetSksClusterArgs {
    /**
     * @deprecated This attribute has been replaced by `exoscaleCcm`/`metricsServer` attributes, it will be removed in a future release.
     */
    addons?: string[];
    /**
     * The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
     */
    aggregationCa?: string;
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
    controlPlaneCa?: string;
    /**
     * The cluster creation date.
     */
    createdAt?: string;
    /**
     * A free-form text describing the cluster.
     */
    description?: string;
    /**
     * ❗ Indicates whether to deploy the Kubernetes network proxy. (may only be set at creation time)
     */
    enableKubeProxy?: boolean;
    /**
     * The cluster API endpoint.
     */
    endpoint?: string;
    /**
     * Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    exoscaleCcm?: boolean;
    /**
     * Deploy the Exoscale [Container Storage Interface](https://github.com/exoscale/exoscale-csi-driver/) on worker nodes (boolean; default: `false`; requires the CCM to be enabled).
     */
    exoscaleCsi?: boolean;
    /**
     * Feature gates options for the cluster.
     */
    featureGates?: string[];
    /**
     * The ID of this resource.
     */
    id?: string;
    /**
     * The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
     */
    kubeletCa?: string;
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
     * The list of exoscale*sks*nodepool (IDs) attached to the cluster.
     */
    nodepools?: string[];
    /**
     * An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
     */
    oidc?: inputs.GetSksClusterOidc;
    /**
     * The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
     */
    serviceLevel?: string;
    /**
     * The cluster state.
     */
    state?: string;
    /**
     * The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
     */
    version?: string;
    zone: string;
}

/**
 * A collection of values returned by getSksCluster.
 */
export interface GetSksClusterResult {
    /**
     * @deprecated This attribute has been replaced by `exoscaleCcm`/`metricsServer` attributes, it will be removed in a future release.
     */
    readonly addons: string[];
    /**
     * The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
     */
    readonly aggregationCa: string;
    /**
     * Enable automatic upgrading of the control plane version.
     */
    readonly autoUpgrade?: boolean;
    /**
     * The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
     */
    readonly cni?: string;
    /**
     * The CA certificate (in PEM format) for TLS communications between control plane components.
     */
    readonly controlPlaneCa: string;
    /**
     * The cluster creation date.
     */
    readonly createdAt: string;
    /**
     * A free-form text describing the cluster.
     */
    readonly description?: string;
    /**
     * ❗ Indicates whether to deploy the Kubernetes network proxy. (may only be set at creation time)
     */
    readonly enableKubeProxy: boolean;
    /**
     * The cluster API endpoint.
     */
    readonly endpoint: string;
    /**
     * Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    readonly exoscaleCcm?: boolean;
    /**
     * Deploy the Exoscale [Container Storage Interface](https://github.com/exoscale/exoscale-csi-driver/) on worker nodes (boolean; default: `false`; requires the CCM to be enabled).
     */
    readonly exoscaleCsi?: boolean;
    /**
     * Feature gates options for the cluster.
     */
    readonly featureGates?: string[];
    /**
     * The ID of this resource.
     */
    readonly id?: string;
    /**
     * The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
     */
    readonly kubeletCa: string;
    /**
     * A map of key/value labels.
     */
    readonly labels?: {[key: string]: string};
    /**
     * Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    readonly metricsServer?: boolean;
    readonly name?: string;
    /**
     * The list of exoscale*sks*nodepool (IDs) attached to the cluster.
     */
    readonly nodepools: string[];
    /**
     * An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
     */
    readonly oidc: outputs.GetSksClusterOidc;
    /**
     * The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
     */
    readonly serviceLevel?: string;
    /**
     * The cluster state.
     */
    readonly state: string;
    /**
     * The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
     */
    readonly version: string;
    readonly zone: string;
}
export function getSksClusterOutput(args: GetSksClusterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSksClusterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getSksCluster:getSksCluster", {
        "addons": args.addons,
        "aggregationCa": args.aggregationCa,
        "autoUpgrade": args.autoUpgrade,
        "cni": args.cni,
        "controlPlaneCa": args.controlPlaneCa,
        "createdAt": args.createdAt,
        "description": args.description,
        "enableKubeProxy": args.enableKubeProxy,
        "endpoint": args.endpoint,
        "exoscaleCcm": args.exoscaleCcm,
        "exoscaleCsi": args.exoscaleCsi,
        "featureGates": args.featureGates,
        "id": args.id,
        "kubeletCa": args.kubeletCa,
        "labels": args.labels,
        "metricsServer": args.metricsServer,
        "name": args.name,
        "nodepools": args.nodepools,
        "oidc": args.oidc,
        "serviceLevel": args.serviceLevel,
        "state": args.state,
        "version": args.version,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSksCluster.
 */
export interface GetSksClusterOutputArgs {
    /**
     * @deprecated This attribute has been replaced by `exoscaleCcm`/`metricsServer` attributes, it will be removed in a future release.
     */
    addons?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
     */
    aggregationCa?: pulumi.Input<string>;
    /**
     * Enable automatic upgrading of the control plane version.
     */
    autoUpgrade?: pulumi.Input<boolean>;
    /**
     * The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
     */
    cni?: pulumi.Input<string>;
    /**
     * The CA certificate (in PEM format) for TLS communications between control plane components.
     */
    controlPlaneCa?: pulumi.Input<string>;
    /**
     * The cluster creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A free-form text describing the cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * ❗ Indicates whether to deploy the Kubernetes network proxy. (may only be set at creation time)
     */
    enableKubeProxy?: pulumi.Input<boolean>;
    /**
     * The cluster API endpoint.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    exoscaleCcm?: pulumi.Input<boolean>;
    /**
     * Deploy the Exoscale [Container Storage Interface](https://github.com/exoscale/exoscale-csi-driver/) on worker nodes (boolean; default: `false`; requires the CCM to be enabled).
     */
    exoscaleCsi?: pulumi.Input<boolean>;
    /**
     * Feature gates options for the cluster.
     */
    featureGates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of this resource.
     */
    id?: pulumi.Input<string>;
    /**
     * The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
     */
    kubeletCa?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    metricsServer?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * The list of exoscale*sks*nodepool (IDs) attached to the cluster.
     */
    nodepools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
     */
    oidc?: pulumi.Input<inputs.GetSksClusterOidcArgs>;
    /**
     * The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
     */
    serviceLevel?: pulumi.Input<string>;
    /**
     * The cluster state.
     */
    state?: pulumi.Input<string>;
    /**
     * The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
     */
    version?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
