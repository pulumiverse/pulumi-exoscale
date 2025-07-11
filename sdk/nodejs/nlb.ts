// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Network Load Balancers (NLB)](https://community.exoscale.com/product/networking/nlb/).
 *
 * Corresponding data source: exoscale_nlb.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myNlb = new exoscale.Nlb("myNlb", {zone: "ch-gva-2"});
 * ```
 *
 * Next step is to attach exoscale_nlb_service(s) to the NLB.
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing network load balancer (NLB) may be imported by `<ID>@<zone>`:
 *
 * console
 *
 * ```sh
 * $ pulumi import exoscale:index/nlb:Nlb \ 
 * ```
 *
 *   exoscale_nlb.my_nlb \
 *
 *   f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
 */
export class Nlb extends pulumi.CustomResource {
    /**
     * Get an existing Nlb resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NlbState, opts?: pulumi.CustomResourceOptions): Nlb {
        return new Nlb(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/nlb:Nlb';

    /**
     * Returns true if the given object is an instance of Nlb.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Nlb {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Nlb.__pulumiType;
    }

    /**
     * The NLB creation date.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A free-form text describing the NLB.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The NLB IPv4 address.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * A map of key/value labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The network load balancer (NLB) name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of the exoscale*nlb*service (names).
     */
    public /*out*/ readonly services!: pulumi.Output<string[]>;
    /**
     * The current NLB state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Nlb resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NlbArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NlbArgs | NlbState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NlbState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NlbArgs | undefined;
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["services"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Nlb.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Nlb resources.
 */
export interface NlbState {
    /**
     * The NLB creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A free-form text describing the NLB.
     */
    description?: pulumi.Input<string>;
    /**
     * The NLB IPv4 address.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The network load balancer (NLB) name.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of the exoscale*nlb*service (names).
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The current NLB state.
     */
    state?: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Nlb resource.
 */
export interface NlbArgs {
    /**
     * A free-form text describing the NLB.
     */
    description?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The network load balancer (NLB) name.
     */
    name?: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
