// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Security Groups](https://community.exoscale.com/documentation/compute/security-groups/).
 *
 * Corresponding data source: exoscale_security_group.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const mySecurityGroup = new exoscale.SecurityGroup("mySecurityGroup", {});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Next step is to attach exoscale_security_group_rule(s) to the group.
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing security group may be imported by `<ID>`:
 *
 * ```sh
 * $ pulumi import exoscale:index/securityGroup:SecurityGroup \
 * ```
 *
 *   exoscale_security_group.my_security_group \
 *
 *   f81d4fae-7dec-11d0-a765-00a0c91e6bf6
 */
export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupState, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/securityGroup:SecurityGroup';

    /**
     * Returns true if the given object is an instance of SecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroup.__pulumiType;
    }

    /**
     * ❗ A free-form text describing the group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of external network sources, in [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notatio) notation.
     */
    public readonly externalSources!: pulumi.Output<string[] | undefined>;
    /**
     * ❗ The security group name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupArgs | SecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["externalSources"] = state ? state.externalSources : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["externalSources"] = args ? args.externalSources : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * ❗ A free-form text describing the group.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of external network sources, in [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notatio) notation.
     */
    externalSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ❗ The security group name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * ❗ A free-form text describing the group.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of external network sources, in [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notatio) notation.
     */
    externalSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ❗ The security group name.
     */
    name?: pulumi.Input<string>;
}
