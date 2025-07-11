// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * List Exoscale [Compute Instances](https://community.exoscale.com/documentation/compute/).
 *
 * Corresponding resource: exoscale_compute_instance.
 */
export function getComputeInstanceList(args: GetComputeInstanceListArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeInstanceListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getComputeInstanceList:getComputeInstanceList", {
        "createdAt": args.createdAt,
        "deployTargetId": args.deployTargetId,
        "diskSize": args.diskSize,
        "id": args.id,
        "ipv6": args.ipv6,
        "ipv6Address": args.ipv6Address,
        "labels": args.labels,
        "managerId": args.managerId,
        "managerType": args.managerType,
        "name": args.name,
        "publicIpAddress": args.publicIpAddress,
        "reverseDns": args.reverseDns,
        "sshKey": args.sshKey,
        "state": args.state,
        "templateId": args.templateId,
        "type": args.type,
        "userData": args.userData,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeInstanceList.
 */
export interface GetComputeInstanceListArgs {
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    createdAt?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    deployTargetId?: string;
    /**
     * Match against this int
     */
    diskSize?: number;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    id?: string;
    /**
     * Match against this bool
     */
    ipv6?: boolean;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    ipv6Address?: string;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    labels?: {[key: string]: string};
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    managerId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    managerType?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    name?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    publicIpAddress?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    reverseDns?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    sshKey?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    state?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    templateId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    type?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    userData?: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

/**
 * A collection of values returned by getComputeInstanceList.
 */
export interface GetComputeInstanceListResult {
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly createdAt?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly deployTargetId?: string;
    /**
     * Match against this int
     */
    readonly diskSize?: number;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly id?: string;
    /**
     * The list of exoscale*compute*instance.
     */
    readonly instances: outputs.GetComputeInstanceListInstance[];
    /**
     * Match against this bool
     */
    readonly ipv6?: boolean;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly ipv6Address?: string;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    readonly labels?: {[key: string]: string};
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly managerId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly managerType?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly name?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly publicIpAddress?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly reverseDns?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly sshKey?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly state?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly templateId?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly type?: string;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    readonly userData?: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    readonly zone: string;
}
/**
 * List Exoscale [Compute Instances](https://community.exoscale.com/documentation/compute/).
 *
 * Corresponding resource: exoscale_compute_instance.
 */
export function getComputeInstanceListOutput(args: GetComputeInstanceListOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetComputeInstanceListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getComputeInstanceList:getComputeInstanceList", {
        "createdAt": args.createdAt,
        "deployTargetId": args.deployTargetId,
        "diskSize": args.diskSize,
        "id": args.id,
        "ipv6": args.ipv6,
        "ipv6Address": args.ipv6Address,
        "labels": args.labels,
        "managerId": args.managerId,
        "managerType": args.managerType,
        "name": args.name,
        "publicIpAddress": args.publicIpAddress,
        "reverseDns": args.reverseDns,
        "sshKey": args.sshKey,
        "state": args.state,
        "templateId": args.templateId,
        "type": args.type,
        "userData": args.userData,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeInstanceList.
 */
export interface GetComputeInstanceListOutputArgs {
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * Match against this int
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    id?: pulumi.Input<string>;
    /**
     * Match against this bool
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    ipv6Address?: pulumi.Input<string>;
    /**
     * Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    managerId?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    managerType?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    name?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    publicIpAddress?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    reverseDns?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    sshKey?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    state?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    templateId?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    type?: pulumi.Input<string>;
    /**
     * Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
     */
    userData?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
