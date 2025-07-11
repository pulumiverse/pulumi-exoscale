// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage [Exoscale Block Storage](https://community.exoscale.com/product/storage/block-storage/) Volume Snapshot.
 *
 * Block Storage offers persistent externally attached volumes for your workloads.
 */
export class BlockStorageVolumeSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing BlockStorageVolumeSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BlockStorageVolumeSnapshotState, opts?: pulumi.CustomResourceOptions): BlockStorageVolumeSnapshot {
        return new BlockStorageVolumeSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/blockStorageVolumeSnapshot:BlockStorageVolumeSnapshot';

    /**
     * Returns true if the given object is an instance of BlockStorageVolumeSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlockStorageVolumeSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlockStorageVolumeSnapshot.__pulumiType;
    }

    /**
     * Snapshot creation date.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Resource labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Volume snapshot name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Snapshot size in GB.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Snapshot state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    public readonly timeouts!: pulumi.Output<outputs.BlockStorageVolumeSnapshotTimeouts | undefined>;
    /**
     * Volume from which to create a snapshot.
     */
    public readonly volume!: pulumi.Output<outputs.BlockStorageVolumeSnapshotVolume>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a BlockStorageVolumeSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlockStorageVolumeSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BlockStorageVolumeSnapshotArgs | BlockStorageVolumeSnapshotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BlockStorageVolumeSnapshotState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["timeouts"] = state ? state.timeouts : undefined;
            resourceInputs["volume"] = state ? state.volume : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as BlockStorageVolumeSnapshotArgs | undefined;
            if ((!args || args.volume === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volume'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["timeouts"] = args ? args.timeouts : undefined;
            resourceInputs["volume"] = args ? args.volume : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BlockStorageVolumeSnapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BlockStorageVolumeSnapshot resources.
 */
export interface BlockStorageVolumeSnapshotState {
    /**
     * Snapshot creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Resource labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Volume snapshot name.
     */
    name?: pulumi.Input<string>;
    /**
     * Snapshot size in GB.
     */
    size?: pulumi.Input<number>;
    /**
     * Snapshot state.
     */
    state?: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.BlockStorageVolumeSnapshotTimeouts>;
    /**
     * Volume from which to create a snapshot.
     */
    volume?: pulumi.Input<inputs.BlockStorageVolumeSnapshotVolume>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BlockStorageVolumeSnapshot resource.
 */
export interface BlockStorageVolumeSnapshotArgs {
    /**
     * Resource labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Volume snapshot name.
     */
    name?: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.BlockStorageVolumeSnapshotTimeouts>;
    /**
     * Volume from which to create a snapshot.
     */
    volume: pulumi.Input<inputs.BlockStorageVolumeSnapshotVolume>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
