// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getNLB(args: GetNLBArgs, opts?: pulumi.InvokeOptions): Promise<GetNLBResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getNLB:getNLB", {
        "id": args.id,
        "name": args.name,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getNLB.
 */
export interface GetNLBArgs {
    /**
     * The Network Load Balancers (NLB) ID to match (conflicts with `name`).
     */
    id?: string;
    /**
     * The NLB name to match (conflicts with `id`).
     */
    name?: string;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: string;
}

/**
 * A collection of values returned by getNLB.
 */
export interface GetNLBResult {
    /**
     * The NLB creation date.
     */
    readonly createdAt: string;
    /**
     * The Network Load Balancers (NLB) description.
     */
    readonly description: string;
    readonly id?: string;
    /**
     * The NLB public IPv4 address.
     */
    readonly ipAddress: string;
    readonly name?: string;
    /**
     * The current NLB state.
     */
    readonly state: string;
    readonly zone: string;
}
export function getNLBOutput(args: GetNLBOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNLBResult> {
    return pulumi.output(args).apply((a: any) => getNLB(a, opts))
}

/**
 * A collection of arguments for invoking getNLB.
 */
export interface GetNLBOutputArgs {
    /**
     * The Network Load Balancers (NLB) ID to match (conflicts with `name`).
     */
    id?: pulumi.Input<string>;
    /**
     * The NLB name to match (conflicts with `id`).
     */
    name?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: pulumi.Input<string>;
}
