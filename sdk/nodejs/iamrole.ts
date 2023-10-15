// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [IAM](https://community.exoscale.com/documentation/iam/) Role.
 */
export class IAMRole extends pulumi.CustomResource {
    /**
     * Get an existing IAMRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMRoleState, opts?: pulumi.CustomResourceOptions): IAMRole {
        return new IAMRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/iAMRole:IAMRole';

    /**
     * Returns true if the given object is an instance of IAMRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMRole.__pulumiType;
    }

    /**
     * A free-form text describing the IAM Role
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Defines if IAM Role Policy is editable or not.
     */
    public readonly editable!: pulumi.Output<boolean>;
    /**
     * IAM Role labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of IAM Role.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IAM Role permissions.
     */
    public readonly permissions!: pulumi.Output<string[]>;
    /**
     * IAM Policy.
     */
    public readonly policy!: pulumi.Output<outputs.IAMRolePolicy>;
    public readonly timeouts!: pulumi.Output<outputs.IAMRoleTimeouts | undefined>;

    /**
     * Create a IAMRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IAMRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMRoleArgs | IAMRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IAMRoleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["editable"] = state ? state.editable : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["timeouts"] = state ? state.timeouts : undefined;
        } else {
            const args = argsOrState as IAMRoleArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["editable"] = args ? args.editable : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["timeouts"] = args ? args.timeouts : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IAMRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMRole resources.
 */
export interface IAMRoleState {
    /**
     * A free-form text describing the IAM Role
     */
    description?: pulumi.Input<string>;
    /**
     * Defines if IAM Role Policy is editable or not.
     */
    editable?: pulumi.Input<boolean>;
    /**
     * IAM Role labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of IAM Role.
     */
    name?: pulumi.Input<string>;
    /**
     * IAM Role permissions.
     */
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IAM Policy.
     */
    policy?: pulumi.Input<inputs.IAMRolePolicy>;
    timeouts?: pulumi.Input<inputs.IAMRoleTimeouts>;
}

/**
 * The set of arguments for constructing a IAMRole resource.
 */
export interface IAMRoleArgs {
    /**
     * A free-form text describing the IAM Role
     */
    description?: pulumi.Input<string>;
    /**
     * Defines if IAM Role Policy is editable or not.
     */
    editable?: pulumi.Input<boolean>;
    /**
     * IAM Role labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of IAM Role.
     */
    name?: pulumi.Input<string>;
    /**
     * IAM Role permissions.
     */
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IAM Policy.
     */
    policy?: pulumi.Input<inputs.IAMRolePolicy>;
    timeouts?: pulumi.Input<inputs.IAMRoleTimeouts>;
}
