// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Fetch Exoscale [SOS Bucket Policies](https://community.exoscale.com/documentation/storage/bucketpolicy/).
 */
export function getSosBucketPolicy(args: GetSosBucketPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetSosBucketPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getSosBucketPolicy:getSosBucketPolicy", {
        "bucket": args.bucket,
        "timeouts": args.timeouts,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSosBucketPolicy.
 */
export interface GetSosBucketPolicyArgs {
    /**
     * The name of the bucket to which the policy is to be applied.
     */
    bucket: string;
    timeouts?: inputs.GetSosBucketPolicyTimeouts;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

/**
 * A collection of values returned by getSosBucketPolicy.
 */
export interface GetSosBucketPolicyResult {
    /**
     * The name of the bucket to which the policy is to be applied.
     */
    readonly bucket: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The content of the policy
     */
    readonly policy: string;
    readonly timeouts?: outputs.GetSosBucketPolicyTimeouts;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    readonly zone: string;
}
/**
 * Fetch Exoscale [SOS Bucket Policies](https://community.exoscale.com/documentation/storage/bucketpolicy/).
 */
export function getSosBucketPolicyOutput(args: GetSosBucketPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSosBucketPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getSosBucketPolicy:getSosBucketPolicy", {
        "bucket": args.bucket,
        "timeouts": args.timeouts,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSosBucketPolicy.
 */
export interface GetSosBucketPolicyOutputArgs {
    /**
     * The name of the bucket to which the policy is to be applied.
     */
    bucket: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.GetSosBucketPolicyTimeoutsArgs>;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}