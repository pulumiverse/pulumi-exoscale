// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDomainRecord(args: GetDomainRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainRecordResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getDomainRecord:getDomainRecord", {
        "domain": args.domain,
        "filter": args.filter,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomainRecord.
 */
export interface GetDomainRecordArgs {
    domain: string;
    /**
     * Filter to apply when looking up domain records.
     */
    filter: inputs.GetDomainRecordFilter;
}

/**
 * A collection of values returned by getDomainRecord.
 */
export interface GetDomainRecordResult {
    /**
     * The exoscale.Domain name to match.
     */
    readonly domain: string;
    /**
     * Filter to apply when looking up domain records.
     */
    readonly filter: outputs.GetDomainRecordFilter;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of matching records. Structure is documented below.
     */
    readonly records: outputs.GetDomainRecordRecord[];
}
export function getDomainRecordOutput(args: GetDomainRecordOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainRecordResult> {
    return pulumi.output(args).apply((a: any) => getDomainRecord(a, opts))
}

/**
 * A collection of arguments for invoking getDomainRecord.
 */
export interface GetDomainRecordOutputArgs {
    domain: pulumi.Input<string>;
    /**
     * Filter to apply when looking up domain records.
     */
    filter: pulumi.Input<inputs.GetDomainRecordFilterArgs>;
}
