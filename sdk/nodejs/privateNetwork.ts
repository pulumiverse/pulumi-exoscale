// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * An existing private network may be imported by `<ID>@<zone>`console
 *
 * ```sh
 *  $ pulumi import exoscale:index/privateNetwork:PrivateNetwork \
 * ```
 *
 *  exoscale_private_network.my_private_network \
 *
 *  f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
 */
export class PrivateNetwork extends pulumi.CustomResource {
    /**
     * Get an existing PrivateNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateNetworkState, opts?: pulumi.CustomResourceOptions): PrivateNetwork {
        return new PrivateNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/privateNetwork:PrivateNetwork';

    /**
     * Returns true if the given object is an instance of PrivateNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateNetwork.__pulumiType;
    }

    /**
     * A free-form text describing the network.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly endIp!: pulumi.Output<string | undefined>;
    /**
     * The private network name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network mask defining the IPv4 network allowed for static leases.
     */
    public readonly netmask!: pulumi.Output<string | undefined>;
    /**
     * /`endIp` - (Required) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
     */
    public readonly startIp!: pulumi.Output<string | undefined>;
    /**
     * The Exoscale [Zone][zone] name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a PrivateNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateNetworkArgs | PrivateNetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateNetworkState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endIp"] = state ? state.endIp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["netmask"] = state ? state.netmask : undefined;
            resourceInputs["startIp"] = state ? state.startIp : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as PrivateNetworkArgs | undefined;
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endIp"] = args ? args.endIp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["netmask"] = args ? args.netmask : undefined;
            resourceInputs["startIp"] = args ? args.startIp : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateNetwork resources.
 */
export interface PrivateNetworkState {
    /**
     * A free-form text describing the network.
     */
    description?: pulumi.Input<string>;
    endIp?: pulumi.Input<string>;
    /**
     * The private network name.
     */
    name?: pulumi.Input<string>;
    /**
     * The network mask defining the IPv4 network allowed for static leases.
     */
    netmask?: pulumi.Input<string>;
    /**
     * /`endIp` - (Required) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
     */
    startIp?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateNetwork resource.
 */
export interface PrivateNetworkArgs {
    /**
     * A free-form text describing the network.
     */
    description?: pulumi.Input<string>;
    endIp?: pulumi.Input<string>;
    /**
     * The private network name.
     */
    name?: pulumi.Input<string>;
    /**
     * The network mask defining the IPv4 network allowed for static leases.
     */
    netmask?: pulumi.Input<string>;
    /**
     * /`endIp` - (Required) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
     */
    startIp?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: pulumi.Input<string>;
}
