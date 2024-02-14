// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSksClusterList(args: GetSksClusterListArgs, opts?: pulumi.InvokeOptions): Promise<GetSksClusterListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getSksClusterList:getSksClusterList", {
        "aggregationCa": args.aggregationCa,
        "autoUpgrade": args.autoUpgrade,
        "cni": args.cni,
        "controlPlaneCa": args.controlPlaneCa,
        "createdAt": args.createdAt,
        "description": args.description,
        "endpoint": args.endpoint,
        "exoscaleCcm": args.exoscaleCcm,
        "id": args.id,
        "kubeletCa": args.kubeletCa,
        "labels": args.labels,
        "metricsServer": args.metricsServer,
        "name": args.name,
        "serviceLevel": args.serviceLevel,
        "state": args.state,
        "version": args.version,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSksClusterList.
 */
export interface GetSksClusterListArgs {
    aggregationCa?: string;
    autoUpgrade?: boolean;
    cni?: string;
    controlPlaneCa?: string;
    createdAt?: string;
    description?: string;
    endpoint?: string;
    exoscaleCcm?: boolean;
    id?: string;
    kubeletCa?: string;
    labels?: {[key: string]: string};
    metricsServer?: boolean;
    name?: string;
    serviceLevel?: string;
    state?: string;
    version?: string;
    zone: string;
}

/**
 * A collection of values returned by getSksClusterList.
 */
export interface GetSksClusterListResult {
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly aggregationCa?: string;
    /**
     * Match against this bool
     */
    readonly autoUpgrade?: boolean;
    readonly clusters: outputs.GetSksClusterListCluster[];
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly cni?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly controlPlaneCa?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly createdAt?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly description?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly endpoint?: string;
    /**
     * Match against this bool
     */
    readonly exoscaleCcm?: boolean;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly id?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly kubeletCa?: string;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    readonly labels?: {[key: string]: string};
    /**
     * Match against this bool
     */
    readonly metricsServer?: boolean;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly name?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly serviceLevel?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly state?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly version?: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    readonly zone: string;
}
export function getSksClusterListOutput(args: GetSksClusterListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSksClusterListResult> {
    return pulumi.output(args).apply((a: any) => getSksClusterList(a, opts))
}

/**
 * A collection of arguments for invoking getSksClusterList.
 */
export interface GetSksClusterListOutputArgs {
    aggregationCa?: pulumi.Input<string>;
    autoUpgrade?: pulumi.Input<boolean>;
    cni?: pulumi.Input<string>;
    controlPlaneCa?: pulumi.Input<string>;
    createdAt?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    endpoint?: pulumi.Input<string>;
    exoscaleCcm?: pulumi.Input<boolean>;
    id?: pulumi.Input<string>;
    kubeletCa?: pulumi.Input<string>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    metricsServer?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    serviceLevel?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    version?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}