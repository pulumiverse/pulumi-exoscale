// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getElasticIP(args: GetElasticIPArgs, opts?: pulumi.InvokeOptions): Promise<GetElasticIPResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getElasticIP:getElasticIP", {
        "id": args.id,
        "ipAddress": args.ipAddress,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getElasticIP.
 */
export interface GetElasticIPArgs {
    /**
     * The Elastic IP (EIP) ID to match (conflicts with `ipAddress`).
     */
    id?: string;
    /**
     * The EIP IPv4 or IPv6 address to match (conflicts with `id`).
     */
    ipAddress?: string;
    /**
     * The Exocale [Zone][zone] name.
     */
    zone: string;
}

/**
 * A collection of values returned by getElasticIP.
 */
export interface GetElasticIPResult {
    /**
     * The Elastic IP (EIP) address family (`inet4` or `inet6`).
     */
    readonly addressFamily: string;
    /**
     * The Elastic IP (EIP) CIDR.
     */
    readonly cidr: string;
    /**
     * The Elastic IP (EIP) description.
     */
    readonly description: string;
    /**
     * (Block) The *managed* EIP healthcheck configuration. Structure is documented below.
     */
    readonly healthchecks: outputs.GetElasticIPHealthcheck[];
    readonly id?: string;
    readonly ipAddress?: string;
    /**
     * A map of key/value labels.
     */
    readonly labels: {[key: string]: string};
    /**
     * Domain name for reverse DNS record.
     */
    readonly reverseDns: string;
    readonly zone: string;
}
export function getElasticIPOutput(args: GetElasticIPOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetElasticIPResult> {
    return pulumi.output(args).apply((a: any) => getElasticIP(a, opts))
}

/**
 * A collection of arguments for invoking getElasticIP.
 */
export interface GetElasticIPOutputArgs {
    /**
     * The Elastic IP (EIP) ID to match (conflicts with `ipAddress`).
     */
    id?: pulumi.Input<string>;
    /**
     * The EIP IPv4 or IPv6 address to match (conflicts with `id`).
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The Exocale [Zone][zone] name.
     */
    zone: pulumi.Input<string>;
}
