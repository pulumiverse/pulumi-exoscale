// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * An existing compute instance may be imported by `<ID>@<zone>`console
 *
 * ```sh
 *  $ pulumi import exoscale:index/computeInstance:ComputeInstance \
 * ```
 *
 *  exoscale_compute_instance.my_instance \
 *
 *  f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
 */
export class ComputeInstance extends pulumi.CustomResource {
    /**
     * Get an existing ComputeInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeInstanceState, opts?: pulumi.CustomResourceOptions): ComputeInstance {
        return new ComputeInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/computeInstance:ComputeInstance';

    /**
     * Returns true if the given object is an instance of ComputeInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeInstance.__pulumiType;
    }

    /**
     * A list of exoscale.AntiAffinityGroup (IDs) to attach to the instance (may only be set at creation time).
     */
    public readonly antiAffinityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The instance creation date.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A deploy target ID.
     */
    public readonly deployTargetId!: pulumi.Output<string | undefined>;
    /**
     * The instance disk size (GiB; at least `10`). **WARNING**: updating this attribute stops/restarts the instance.
     */
    public readonly diskSize!: pulumi.Output<number>;
    /**
     * A list of exoscale.ElasticIP (IDs) to attach to the instance.
     */
    public readonly elasticIpIds!: pulumi.Output<string[] | undefined>;
    /**
     * Enable IPv6 on the instance (boolean; default: `false`).
     */
    public readonly ipv6!: pulumi.Output<boolean | undefined>;
    /**
     * The instance (main network interface) IPv6 address (if enabled).
     */
    public /*out*/ readonly ipv6Address!: pulumi.Output<string>;
    /**
     * A map of key/value labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The compute instance name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Private network interfaces (may be specified multiple times). Structure is documented below.
     */
    public readonly networkInterfaces!: pulumi.Output<outputs.ComputeInstanceNetworkInterface[] | undefined>;
    /**
     * (Deprecated) A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument instead.
     *
     * @deprecated Use the network_interface block instead.
     */
    public /*out*/ readonly privateNetworkIds!: pulumi.Output<string[]>;
    /**
     * The instance (main network interface) IPv4 address.
     */
    public /*out*/ readonly publicIpAddress!: pulumi.Output<string>;
    /**
     * A list of exoscale.SecurityGroup (IDs) to attach to the instance.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The exoscale.SSHKey (name) to authorize in the instance (may only be set at creation time).
     */
    public readonly sshKey!: pulumi.Output<string | undefined>;
    /**
     * The instance state (`running` or `stopped`; default: `running`).
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The exoscale.getComputeTemplate (ID) to use when creating the instance.
     */
    public readonly templateId!: pulumi.Output<string>;
    /**
     * The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI][cli] - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * [cloud-init][cloud-init] configuration (no need to base64-encode or gzip it as the provider will take care of it).
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * The Exoscale [Zone][zone] name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a ComputeInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeInstanceArgs | ComputeInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeInstanceState | undefined;
            resourceInputs["antiAffinityGroupIds"] = state ? state.antiAffinityGroupIds : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deployTargetId"] = state ? state.deployTargetId : undefined;
            resourceInputs["diskSize"] = state ? state.diskSize : undefined;
            resourceInputs["elasticIpIds"] = state ? state.elasticIpIds : undefined;
            resourceInputs["ipv6"] = state ? state.ipv6 : undefined;
            resourceInputs["ipv6Address"] = state ? state.ipv6Address : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            resourceInputs["privateNetworkIds"] = state ? state.privateNetworkIds : undefined;
            resourceInputs["publicIpAddress"] = state ? state.publicIpAddress : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["sshKey"] = state ? state.sshKey : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["templateId"] = state ? state.templateId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ComputeInstanceArgs | undefined;
            if ((!args || args.templateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["antiAffinityGroupIds"] = args ? args.antiAffinityGroupIds : undefined;
            resourceInputs["deployTargetId"] = args ? args.deployTargetId : undefined;
            resourceInputs["diskSize"] = args ? args.diskSize : undefined;
            resourceInputs["elasticIpIds"] = args ? args.elasticIpIds : undefined;
            resourceInputs["ipv6"] = args ? args.ipv6 : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["sshKey"] = args ? args.sshKey : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["ipv6Address"] = undefined /*out*/;
            resourceInputs["privateNetworkIds"] = undefined /*out*/;
            resourceInputs["publicIpAddress"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComputeInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeInstance resources.
 */
export interface ComputeInstanceState {
    /**
     * A list of exoscale.AntiAffinityGroup (IDs) to attach to the instance (may only be set at creation time).
     */
    antiAffinityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A deploy target ID.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * The instance disk size (GiB; at least `10`). **WARNING**: updating this attribute stops/restarts the instance.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * A list of exoscale.ElasticIP (IDs) to attach to the instance.
     */
    elasticIpIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable IPv6 on the instance (boolean; default: `false`).
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * The instance (main network interface) IPv6 address (if enabled).
     */
    ipv6Address?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The compute instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * Private network interfaces (may be specified multiple times). Structure is documented below.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.ComputeInstanceNetworkInterface>[]>;
    /**
     * (Deprecated) A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument instead.
     *
     * @deprecated Use the network_interface block instead.
     */
    privateNetworkIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance (main network interface) IPv4 address.
     */
    publicIpAddress?: pulumi.Input<string>;
    /**
     * A list of exoscale.SecurityGroup (IDs) to attach to the instance.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The exoscale.SSHKey (name) to authorize in the instance (may only be set at creation time).
     */
    sshKey?: pulumi.Input<string>;
    /**
     * The instance state (`running` or `stopped`; default: `running`).
     */
    state?: pulumi.Input<string>;
    /**
     * The exoscale.getComputeTemplate (ID) to use when creating the instance.
     */
    templateId?: pulumi.Input<string>;
    /**
     * The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI][cli] - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
     */
    type?: pulumi.Input<string>;
    /**
     * [cloud-init][cloud-init] configuration (no need to base64-encode or gzip it as the provider will take care of it).
     */
    userData?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeInstance resource.
 */
export interface ComputeInstanceArgs {
    /**
     * A list of exoscale.AntiAffinityGroup (IDs) to attach to the instance (may only be set at creation time).
     */
    antiAffinityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A deploy target ID.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * The instance disk size (GiB; at least `10`). **WARNING**: updating this attribute stops/restarts the instance.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * A list of exoscale.ElasticIP (IDs) to attach to the instance.
     */
    elasticIpIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable IPv6 on the instance (boolean; default: `false`).
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The compute instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * Private network interfaces (may be specified multiple times). Structure is documented below.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.ComputeInstanceNetworkInterface>[]>;
    /**
     * A list of exoscale.SecurityGroup (IDs) to attach to the instance.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The exoscale.SSHKey (name) to authorize in the instance (may only be set at creation time).
     */
    sshKey?: pulumi.Input<string>;
    /**
     * The instance state (`running` or `stopped`; default: `running`).
     */
    state?: pulumi.Input<string>;
    /**
     * The exoscale.getComputeTemplate (ID) to use when creating the instance.
     */
    templateId: pulumi.Input<string>;
    /**
     * The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI][cli] - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
     */
    type: pulumi.Input<string>;
    /**
     * [cloud-init][cloud-init] configuration (no need to base64-encode or gzip it as the provider will take care of it).
     */
    userData?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: pulumi.Input<string>;
}
