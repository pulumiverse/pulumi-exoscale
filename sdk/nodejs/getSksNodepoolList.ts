// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSksNodepoolList(args: GetSksNodepoolListArgs, opts?: pulumi.InvokeOptions): Promise<GetSksNodepoolListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getSksNodepoolList:getSksNodepoolList", {
        "clusterId": args.clusterId,
        "createdAt": args.createdAt,
        "deployTargetId": args.deployTargetId,
        "description": args.description,
        "diskSize": args.diskSize,
        "id": args.id,
        "instancePoolId": args.instancePoolId,
        "instancePrefix": args.instancePrefix,
        "instanceType": args.instanceType,
        "labels": args.labels,
        "name": args.name,
        "size": args.size,
        "state": args.state,
        "storageLvm": args.storageLvm,
        "taints": args.taints,
        "templateId": args.templateId,
        "version": args.version,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSksNodepoolList.
 */
export interface GetSksNodepoolListArgs {
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    clusterId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    createdAt?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    deployTargetId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    description?: string;
    /**
     * Match against this int
     */
    diskSize?: number;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    id?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    instancePoolId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    instancePrefix?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    instanceType?: string;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    labels?: {[key: string]: string};
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    name?: string;
    /**
     * Match against this int
     */
    size?: number;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    state?: string;
    /**
     * Match against this bool
     */
    storageLvm?: boolean;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    taints?: {[key: string]: string};
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    templateId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    version?: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

/**
 * A collection of values returned by getSksNodepoolList.
 */
export interface GetSksNodepoolListResult {
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly clusterId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly createdAt?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly deployTargetId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly description?: string;
    /**
     * Match against this int
     */
    readonly diskSize?: number;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly id?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly instancePoolId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly instancePrefix?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly instanceType?: string;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    readonly labels?: {[key: string]: string};
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly name?: string;
    readonly nodepools: outputs.GetSksNodepoolListNodepool[];
    /**
     * Match against this int
     */
    readonly size?: number;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly state?: string;
    /**
     * Match against this bool
     */
    readonly storageLvm?: boolean;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    readonly taints?: {[key: string]: string};
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly templateId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly version?: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    readonly zone: string;
}
export function getSksNodepoolListOutput(args: GetSksNodepoolListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSksNodepoolListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getSksNodepoolList:getSksNodepoolList", {
        "clusterId": args.clusterId,
        "createdAt": args.createdAt,
        "deployTargetId": args.deployTargetId,
        "description": args.description,
        "diskSize": args.diskSize,
        "id": args.id,
        "instancePoolId": args.instancePoolId,
        "instancePrefix": args.instancePrefix,
        "instanceType": args.instanceType,
        "labels": args.labels,
        "name": args.name,
        "size": args.size,
        "state": args.state,
        "storageLvm": args.storageLvm,
        "taints": args.taints,
        "templateId": args.templateId,
        "version": args.version,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSksNodepoolList.
 */
export interface GetSksNodepoolListOutputArgs {
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    description?: pulumi.Input<string>;
    /**
     * Match against this int
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    id?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    instancePoolId?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    instancePrefix?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    name?: pulumi.Input<string>;
    /**
     * Match against this int
     */
    size?: pulumi.Input<number>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    state?: pulumi.Input<string>;
    /**
     * Match against this bool
     */
    storageLvm?: pulumi.Input<boolean>;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    taints?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    templateId?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    version?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
