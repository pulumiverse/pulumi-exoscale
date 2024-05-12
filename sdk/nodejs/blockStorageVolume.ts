// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage [Exoscale Block Storage](https://community.exoscale.com/documentation/block-storage/) Volume.
 *
 * Block Storage offers persistent externally attached volumes for your workloads.
 */
export class BlockStorageVolume extends pulumi.CustomResource {
    /**
     * Get an existing BlockStorageVolume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BlockStorageVolumeState, opts?: pulumi.CustomResourceOptions): BlockStorageVolume {
        return new BlockStorageVolume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/blockStorageVolume:BlockStorageVolume';

    /**
     * Returns true if the given object is an instance of BlockStorageVolume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlockStorageVolume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlockStorageVolume.__pulumiType;
    }

    /**
     * Volume block size.
     */
    public /*out*/ readonly blocksize!: pulumi.Output<number>;
    /**
     * Volume creation date.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * ❗ Resource labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * ❗ Volume name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
     */
    public readonly size!: pulumi.Output<number | undefined>;
    /**
     * Block storage snapshot to use when creating a volume. Read-only after creation.
     */
    public readonly snapshotTarget!: pulumi.Output<outputs.BlockStorageVolumeSnapshotTarget | undefined>;
    /**
     * Volume state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    public readonly timeouts!: pulumi.Output<outputs.BlockStorageVolumeTimeouts | undefined>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a BlockStorageVolume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlockStorageVolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BlockStorageVolumeArgs | BlockStorageVolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BlockStorageVolumeState | undefined;
            resourceInputs["blocksize"] = state ? state.blocksize : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotTarget"] = state ? state.snapshotTarget : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["timeouts"] = state ? state.timeouts : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as BlockStorageVolumeArgs | undefined;
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["snapshotTarget"] = args ? args.snapshotTarget : undefined;
            resourceInputs["timeouts"] = args ? args.timeouts : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["blocksize"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BlockStorageVolume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BlockStorageVolume resources.
 */
export interface BlockStorageVolumeState {
    /**
     * Volume block size.
     */
    blocksize?: pulumi.Input<number>;
    /**
     * Volume creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * ❗ Resource labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ❗ Volume name.
     */
    name?: pulumi.Input<string>;
    /**
     * Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
     */
    size?: pulumi.Input<number>;
    /**
     * Block storage snapshot to use when creating a volume. Read-only after creation.
     */
    snapshotTarget?: pulumi.Input<inputs.BlockStorageVolumeSnapshotTarget>;
    /**
     * Volume state.
     */
    state?: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.BlockStorageVolumeTimeouts>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BlockStorageVolume resource.
 */
export interface BlockStorageVolumeArgs {
    /**
     * ❗ Resource labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ❗ Volume name.
     */
    name?: pulumi.Input<string>;
    /**
     * Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
     */
    size?: pulumi.Input<number>;
    /**
     * Block storage snapshot to use when creating a volume. Read-only after creation.
     */
    snapshotTarget?: pulumi.Input<inputs.BlockStorageVolumeSnapshotTarget>;
    timeouts?: pulumi.Input<inputs.BlockStorageVolumeTimeouts>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}