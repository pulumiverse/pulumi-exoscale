// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Scalable Kubernetes Service (SKS)](https://community.exoscale.com/product/compute/containers/) Clusters.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const mySksCluster = new exoscale.SksCluster("mySksCluster", {zone: "ch-gva-2"});
 * export const mySksClusterEndpoint = mySksCluster.endpoint;
 * ```
 *
 * Next step is to attach exoscale_sks_nodepool(s) to the cluster.
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing SKS cluster may be imported by `<ID>@<zone>`:
 *
 * ```sh
 * $ pulumi import exoscale:index/sksCluster:SksCluster \ 
 * ```
 *
 *   exoscale_sks_cluster.my_sks_cluster \
 *
 *   f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
 */
export class SksCluster extends pulumi.CustomResource {
    /**
     * Get an existing SksCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SksClusterState, opts?: pulumi.CustomResourceOptions): SksCluster {
        return new SksCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/sksCluster:SksCluster';

    /**
     * Returns true if the given object is an instance of SksCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SksCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SksCluster.__pulumiType;
    }

    /**
     * @deprecated This attribute has been replaced by `exoscaleCcm`/`metricsServer` attributes, it will be removed in a future release.
     */
    public readonly addons!: pulumi.Output<string[]>;
    /**
     * The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
     */
    public /*out*/ readonly aggregationCa!: pulumi.Output<string>;
    /**
     * Enable automatic upgrading of the control plane version.
     */
    public readonly autoUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
     */
    public readonly cni!: pulumi.Output<string | undefined>;
    /**
     * The CA certificate (in PEM format) for TLS communications between control plane components.
     */
    public /*out*/ readonly controlPlaneCa!: pulumi.Output<string>;
    /**
     * The cluster creation date.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A free-form text describing the cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ❗ Indicates whether to deploy the Kubernetes network proxy. (may only be set at creation time)
     */
    public readonly enableKubeProxy!: pulumi.Output<boolean>;
    /**
     * The cluster API endpoint.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    public readonly exoscaleCcm!: pulumi.Output<boolean | undefined>;
    /**
     * Deploy the Exoscale [Container Storage Interface](https://github.com/exoscale/exoscale-csi-driver/) on worker nodes (boolean; default: `false`; requires the CCM to be enabled).
     */
    public readonly exoscaleCsi!: pulumi.Output<boolean | undefined>;
    /**
     * Feature gates options for the cluster.
     */
    public readonly featureGates!: pulumi.Output<string[] | undefined>;
    /**
     * The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
     */
    public /*out*/ readonly kubeletCa!: pulumi.Output<string>;
    /**
     * A map of key/value labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    public readonly metricsServer!: pulumi.Output<boolean | undefined>;
    /**
     * The SKS cluster name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of exoscale*sks*nodepool (IDs) attached to the cluster.
     */
    public /*out*/ readonly nodepools!: pulumi.Output<string[]>;
    /**
     * An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
     */
    public readonly oidc!: pulumi.Output<outputs.SksClusterOidc>;
    /**
     * The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
     */
    public readonly serviceLevel!: pulumi.Output<string | undefined>;
    /**
     * The cluster state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a SksCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SksClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SksClusterArgs | SksClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SksClusterState | undefined;
            resourceInputs["addons"] = state ? state.addons : undefined;
            resourceInputs["aggregationCa"] = state ? state.aggregationCa : undefined;
            resourceInputs["autoUpgrade"] = state ? state.autoUpgrade : undefined;
            resourceInputs["cni"] = state ? state.cni : undefined;
            resourceInputs["controlPlaneCa"] = state ? state.controlPlaneCa : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableKubeProxy"] = state ? state.enableKubeProxy : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["exoscaleCcm"] = state ? state.exoscaleCcm : undefined;
            resourceInputs["exoscaleCsi"] = state ? state.exoscaleCsi : undefined;
            resourceInputs["featureGates"] = state ? state.featureGates : undefined;
            resourceInputs["kubeletCa"] = state ? state.kubeletCa : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["metricsServer"] = state ? state.metricsServer : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodepools"] = state ? state.nodepools : undefined;
            resourceInputs["oidc"] = state ? state.oidc : undefined;
            resourceInputs["serviceLevel"] = state ? state.serviceLevel : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as SksClusterArgs | undefined;
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["addons"] = args ? args.addons : undefined;
            resourceInputs["autoUpgrade"] = args ? args.autoUpgrade : undefined;
            resourceInputs["cni"] = args ? args.cni : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableKubeProxy"] = args ? args.enableKubeProxy : undefined;
            resourceInputs["exoscaleCcm"] = args ? args.exoscaleCcm : undefined;
            resourceInputs["exoscaleCsi"] = args ? args.exoscaleCsi : undefined;
            resourceInputs["featureGates"] = args ? args.featureGates : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["metricsServer"] = args ? args.metricsServer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oidc"] = args ? args.oidc : undefined;
            resourceInputs["serviceLevel"] = args ? args.serviceLevel : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["aggregationCa"] = undefined /*out*/;
            resourceInputs["controlPlaneCa"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["kubeletCa"] = undefined /*out*/;
            resourceInputs["nodepools"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SksCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SksCluster resources.
 */
export interface SksClusterState {
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
    /**
     * The SKS cluster name.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of exoscale*sks*nodepool (IDs) attached to the cluster.
     */
    nodepools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
     */
    oidc?: pulumi.Input<inputs.SksClusterOidc>;
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
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SksCluster resource.
 */
export interface SksClusterArgs {
    /**
     * @deprecated This attribute has been replaced by `exoscaleCcm`/`metricsServer` attributes, it will be removed in a future release.
     */
    addons?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable automatic upgrading of the control plane version.
     */
    autoUpgrade?: pulumi.Input<boolean>;
    /**
     * The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
     */
    cni?: pulumi.Input<string>;
    /**
     * A free-form text describing the cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * ❗ Indicates whether to deploy the Kubernetes network proxy. (may only be set at creation time)
     */
    enableKubeProxy?: pulumi.Input<boolean>;
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
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
     */
    metricsServer?: pulumi.Input<boolean>;
    /**
     * The SKS cluster name.
     */
    name?: pulumi.Input<string>;
    /**
     * An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
     */
    oidc?: pulumi.Input<inputs.SksClusterOidc>;
    /**
     * The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
     */
    serviceLevel?: pulumi.Input<string>;
    /**
     * The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
     */
    version?: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
