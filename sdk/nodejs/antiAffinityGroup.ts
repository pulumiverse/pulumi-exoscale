// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Anti-Affinity Groups](https://community.exoscale.com/documentation/compute/anti-affinity-groups/).
 *
 * Corresponding data source: exoscale_anti_affinity_group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myAntiAffinityGroup = new exoscale.AntiAffinityGroup("myAntiAffinityGroup", {description: "Prevent compute instances to run on the same host"});
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing anti-affinity group may be imported by `<ID>`:
 *
 * ```sh
 * $ pulumi import exoscale:index/antiAffinityGroup:AntiAffinityGroup \
 * ```
 *
 *   exoscale_anti_affinity_group.my_anti_affinity_group \
 *
 *   f81d4fae-7dec-11d0-a765-00a0c91e6bf6
 */
export class AntiAffinityGroup extends pulumi.CustomResource {
    /**
     * Get an existing AntiAffinityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AntiAffinityGroupState, opts?: pulumi.CustomResourceOptions): AntiAffinityGroup {
        return new AntiAffinityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/antiAffinityGroup:AntiAffinityGroup';

    /**
     * Returns true if the given object is an instance of AntiAffinityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AntiAffinityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AntiAffinityGroup.__pulumiType;
    }

    /**
     * ❗ A free-form text describing the group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ❗ The anti-affinity group name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a AntiAffinityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AntiAffinityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AntiAffinityGroupArgs | AntiAffinityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AntiAffinityGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as AntiAffinityGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AntiAffinityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AntiAffinityGroup resources.
 */
export interface AntiAffinityGroupState {
    /**
     * ❗ A free-form text describing the group.
     */
    description?: pulumi.Input<string>;
    /**
     * ❗ The anti-affinity group name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AntiAffinityGroup resource.
 */
export interface AntiAffinityGroupArgs {
    /**
     * ❗ A free-form text describing the group.
     */
    description?: pulumi.Input<string>;
    /**
     * ❗ The anti-affinity group name.
     */
    name?: pulumi.Input<string>;
}
