// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    createdAt?: string;
    deployTargetId?: string;
    diskSize?: number;
    id?: string;
    ipv6?: boolean;
    ipv6Address?: string;
    labels?: {[key: string]: string};
    managerId?: string;
    managerType?: string;
    name?: string;
    publicIpAddress?: string;
    reverseDns?: string;
    sshKey?: string;
    state?: string;
    templateId?: string;
    type?: string;
    userData?: string;
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
export function getComputeInstanceListOutput(args: GetComputeInstanceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeInstanceListResult> {
    return pulumi.output(args).apply((a: any) => getComputeInstanceList(a, opts))
}

/**
 * A collection of arguments for invoking getComputeInstanceList.
 */
export interface GetComputeInstanceListOutputArgs {
    createdAt?: pulumi.Input<string>;
    deployTargetId?: pulumi.Input<string>;
    diskSize?: pulumi.Input<number>;
    id?: pulumi.Input<string>;
    ipv6?: pulumi.Input<boolean>;
    ipv6Address?: pulumi.Input<string>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    managerId?: pulumi.Input<string>;
    managerType?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    publicIpAddress?: pulumi.Input<string>;
    reverseDns?: pulumi.Input<string>;
    sshKey?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    templateId?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userData?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
