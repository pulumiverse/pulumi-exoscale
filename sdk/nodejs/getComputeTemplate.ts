// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getComputeTemplate(args: GetComputeTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getComputeTemplate:getComputeTemplate", {
        "filter": args.filter,
        "id": args.id,
        "name": args.name,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeTemplate.
 */
export interface GetComputeTemplateArgs {
    /**
     * A template category filter (default: `featured`); among:
     */
    filter?: string;
    /**
     * The compute instance template ID to match (conflicts with `name`).
     */
    id?: string;
    /**
     * The template name to match (conflicts with `id`).
     */
    name?: string;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: string;
}

/**
 * A collection of values returned by getComputeTemplate.
 */
export interface GetComputeTemplateResult {
    readonly filter?: string;
    readonly id?: string;
    readonly name?: string;
    /**
     * Username to use to log into a compute instance based on this template
     */
    readonly username: string;
    readonly zone: string;
}
export function getComputeTemplateOutput(args: GetComputeTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeTemplateResult> {
    return pulumi.output(args).apply((a: any) => getComputeTemplate(a, opts))
}

/**
 * A collection of arguments for invoking getComputeTemplate.
 */
export interface GetComputeTemplateOutputArgs {
    /**
     * A template category filter (default: `featured`); among:
     */
    filter?: pulumi.Input<string>;
    /**
     * The compute instance template ID to match (conflicts with `name`).
     */
    id?: pulumi.Input<string>;
    /**
     * The template name to match (conflicts with `id`).
     */
    name?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: pulumi.Input<string>;
}
