// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Compute Instances](https://community.exoscale.com/documentation/compute/).
 *
 * Corresponding data sources: exoscale_compute_instance, exoscale_compute_instance_list.
 *
 * After the creation, you can retrieve the password of an instance with [Exoscale CLI](https://github.com/exoscale/cli): `exo compute instance reveal-password NAME`.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myTemplate = exoscale.getTemplate({
 *     zone: "ch-gva-2",
 *     name: "Linux Ubuntu 22.04 LTS 64-bit",
 * });
 * const myInstance = new exoscale.ComputeInstance("myInstance", {
 *     zone: "ch-gva-2",
 *     templateId: myTemplate.then(myTemplate => myTemplate.id),
 *     type: "standard.medium",
 *     diskSize: 10,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing compute instance may be imported by `<ID>@<zone>`:
 *
 * ```sh
 * $ pulumi import exoscale:index/computeInstance:ComputeInstance \
 * ```
 *
 *   exoscale_compute_instance.my_instance \
 *
 *   f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
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
     * ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
     */
    public readonly antiAffinityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The instance creation date.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * ❗ A deploy target ID.
     */
    public readonly deployTargetId!: pulumi.Output<string | undefined>;
    /**
     * Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
     */
    public readonly destroyProtected!: pulumi.Output<boolean | undefined>;
    /**
     * The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
     */
    public readonly diskSize!: pulumi.Output<number>;
    /**
     * A list of exoscale*elastic*ip (IDs) to attach to the instance.
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
     * Whether the instance is private (no public IP addresses; default: false)
     */
    public readonly private!: pulumi.Output<boolean | undefined>;
    /**
     * A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument instead.
     *
     * @deprecated Use the networkInterface block instead.
     */
    public /*out*/ readonly privateNetworkIds!: pulumi.Output<string[]>;
    /**
     * The instance (main network interface) IPv4 address.
     */
    public /*out*/ readonly publicIpAddress!: pulumi.Output<string>;
    /**
     * Domain name for reverse DNS record.
     */
    public readonly reverseDns!: pulumi.Output<string | undefined>;
    /**
     * A list of exoscale*security*group (IDs) to attach to the instance.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
     */
    public readonly sshKey!: pulumi.Output<string | undefined>;
    /**
     * The instance state (`running` or `stopped`; default: `running`).
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * ❗ The exoscale.getTemplate (ID) to use when creating the instance.
     */
    public readonly templateId!: pulumi.Output<string>;
    /**
     * The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * [cloud-init](https://cloudinit.readthedocs.io/) configuration.
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
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
            resourceInputs["destroyProtected"] = state ? state.destroyProtected : undefined;
            resourceInputs["diskSize"] = state ? state.diskSize : undefined;
            resourceInputs["elasticIpIds"] = state ? state.elasticIpIds : undefined;
            resourceInputs["ipv6"] = state ? state.ipv6 : undefined;
            resourceInputs["ipv6Address"] = state ? state.ipv6Address : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            resourceInputs["private"] = state ? state.private : undefined;
            resourceInputs["privateNetworkIds"] = state ? state.privateNetworkIds : undefined;
            resourceInputs["publicIpAddress"] = state ? state.publicIpAddress : undefined;
            resourceInputs["reverseDns"] = state ? state.reverseDns : undefined;
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
            resourceInputs["destroyProtected"] = args ? args.destroyProtected : undefined;
            resourceInputs["diskSize"] = args ? args.diskSize : undefined;
            resourceInputs["elasticIpIds"] = args ? args.elasticIpIds : undefined;
            resourceInputs["ipv6"] = args ? args.ipv6 : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            resourceInputs["private"] = args ? args.private : undefined;
            resourceInputs["reverseDns"] = args ? args.reverseDns : undefined;
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
     * ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
     */
    antiAffinityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * ❗ A deploy target ID.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
     */
    destroyProtected?: pulumi.Input<boolean>;
    /**
     * The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * A list of exoscale*elastic*ip (IDs) to attach to the instance.
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
     * Whether the instance is private (no public IP addresses; default: false)
     */
    private?: pulumi.Input<boolean>;
    /**
     * A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument instead.
     *
     * @deprecated Use the networkInterface block instead.
     */
    privateNetworkIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance (main network interface) IPv4 address.
     */
    publicIpAddress?: pulumi.Input<string>;
    /**
     * Domain name for reverse DNS record.
     */
    reverseDns?: pulumi.Input<string>;
    /**
     * A list of exoscale*security*group (IDs) to attach to the instance.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
     */
    sshKey?: pulumi.Input<string>;
    /**
     * The instance state (`running` or `stopped`; default: `running`).
     */
    state?: pulumi.Input<string>;
    /**
     * ❗ The exoscale.getTemplate (ID) to use when creating the instance.
     */
    templateId?: pulumi.Input<string>;
    /**
     * The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
     */
    type?: pulumi.Input<string>;
    /**
     * [cloud-init](https://cloudinit.readthedocs.io/) configuration.
     */
    userData?: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeInstance resource.
 */
export interface ComputeInstanceArgs {
    /**
     * ❗ A list of exoscale*anti*affinity_group (IDs) to attach to the instance (may only be set at creation time).
     */
    antiAffinityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ❗ A deploy target ID.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * Mark the instance as protected, the Exoscale API will refuse to delete the instance until the protection is removed (boolean; default: `false`).
     */
    destroyProtected?: pulumi.Input<boolean>;
    /**
     * The instance disk size (GiB; at least `10`). Can not be decreased after creation. **WARNING**: updating this attribute stops/restarts the instance.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * A list of exoscale*elastic*ip (IDs) to attach to the instance.
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
     * Whether the instance is private (no public IP addresses; default: false)
     */
    private?: pulumi.Input<boolean>;
    /**
     * Domain name for reverse DNS record.
     */
    reverseDns?: pulumi.Input<string>;
    /**
     * A list of exoscale*security*group (IDs) to attach to the instance.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The exoscale*ssh*key (name) to authorize in the instance (may only be set at creation time).
     */
    sshKey?: pulumi.Input<string>;
    /**
     * The instance state (`running` or `stopped`; default: `running`).
     */
    state?: pulumi.Input<string>;
    /**
     * ❗ The exoscale.getTemplate (ID) to use when creating the instance.
     */
    templateId: pulumi.Input<string>;
    /**
     * The instance type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts the instance.
     */
    type: pulumi.Input<string>;
    /**
     * [cloud-init](https://cloudinit.readthedocs.io/) configuration.
     */
    userData?: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
