// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Security Group](https://community.exoscale.com/documentation/compute/security-groups/) Rules.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const mySecurityGroup = new exoscale.SecurityGroup("mySecurityGroup", {});
 * const mySecurityGroupRule = new exoscale.SecurityGroupRule("mySecurityGroupRule", {
 *     securityGroupId: mySecurityGroup.id,
 *     type: "INGRESS",
 *     protocol: "TCP",
 *     cidr: "0.0.0.0/0",
 *     startPort: 80,
 *     endPort: 80,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing security group rule may be imported by `<security-group-ID>/<security-group-rule-ID>`:
 *
 * ```sh
 * $ pulumi import exoscale:index/securityGroupRule:SecurityGroupRule \
 * ```
 *
 *   exoscale_security_group_rule.my_security_group_rule \
 *
 *   f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524
 */
export class SecurityGroupRule extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions): SecurityGroupRule {
        return new SecurityGroupRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/securityGroupRule:SecurityGroupRule';

    /**
     * Returns true if the given object is an instance of SecurityGroupRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroupRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroupRule.__pulumiType;
    }

    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)) to match (conflicts with `publicSecurityGroup`/`userSecurityGroup`/`userSecurityGroupId`).
     */
    public readonly cidr!: pulumi.Output<string | undefined>;
    /**
     * ❗ A free-form text describing the security group rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ❗ A `TCP`/`UDP` port range to match.
     */
    public readonly endPort!: pulumi.Output<number | undefined>;
    /**
     * ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
     */
    public readonly icmpCode!: pulumi.Output<number | undefined>;
    /**
     * ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
     */
    public readonly icmpType!: pulumi.Output<number | undefined>;
    /**
     * ❗ The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination public security group name to match (conflicts with `cidr`/`userSecurityGroup`/`userSecurityGroupId`).
     */
    public readonly publicSecurityGroup!: pulumi.Output<string>;
    /**
     * ❗ The parent security group name. Please use the `securityGroupId` argument along the exoscale*security*group data source instead.
     *
     * @deprecated Deprecated in favor of `securityGroupId`
     */
    public readonly securityGroup!: pulumi.Output<string>;
    /**
     * ❗ The parent exoscale*security*group ID.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * ❗ A `TCP`/`UDP` port range to match.
     */
    public readonly startPort!: pulumi.Output<number | undefined>;
    /**
     * ❗ The traffic direction to match (`INGRESS` or `EGRESS`).
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`publicSecurityGroup`/`userSecurityGroupId`). Please use the `userSecurityGroupId` argument along the exoscale*security*group data source instead.
     *
     * @deprecated Deprecated in favor of `userSecurityGroupId`
     */
    public readonly userSecurityGroup!: pulumi.Output<string>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`publicSecurityGroup`/`user_security_group)`).
     */
    public readonly userSecurityGroupId!: pulumi.Output<string | undefined>;

    /**
     * Create a SecurityGroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupRuleArgs | SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupRuleState | undefined;
            resourceInputs["cidr"] = state ? state.cidr : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endPort"] = state ? state.endPort : undefined;
            resourceInputs["icmpCode"] = state ? state.icmpCode : undefined;
            resourceInputs["icmpType"] = state ? state.icmpType : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["publicSecurityGroup"] = state ? state.publicSecurityGroup : undefined;
            resourceInputs["securityGroup"] = state ? state.securityGroup : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["startPort"] = state ? state.startPort : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userSecurityGroup"] = state ? state.userSecurityGroup : undefined;
            resourceInputs["userSecurityGroupId"] = state ? state.userSecurityGroupId : undefined;
        } else {
            const args = argsOrState as SecurityGroupRuleArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["cidr"] = args ? args.cidr : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endPort"] = args ? args.endPort : undefined;
            resourceInputs["icmpCode"] = args ? args.icmpCode : undefined;
            resourceInputs["icmpType"] = args ? args.icmpType : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["publicSecurityGroup"] = args ? args.publicSecurityGroup : undefined;
            resourceInputs["securityGroup"] = args ? args.securityGroup : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["startPort"] = args ? args.startPort : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userSecurityGroup"] = args ? args.userSecurityGroup : undefined;
            resourceInputs["userSecurityGroupId"] = args ? args.userSecurityGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroupRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroupRule resources.
 */
export interface SecurityGroupRuleState {
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)) to match (conflicts with `publicSecurityGroup`/`userSecurityGroup`/`userSecurityGroupId`).
     */
    cidr?: pulumi.Input<string>;
    /**
     * ❗ A free-form text describing the security group rule.
     */
    description?: pulumi.Input<string>;
    /**
     * ❗ A `TCP`/`UDP` port range to match.
     */
    endPort?: pulumi.Input<number>;
    /**
     * ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
     */
    icmpCode?: pulumi.Input<number>;
    /**
     * ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
     */
    icmpType?: pulumi.Input<number>;
    /**
     * ❗ The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
     */
    protocol?: pulumi.Input<string>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination public security group name to match (conflicts with `cidr`/`userSecurityGroup`/`userSecurityGroupId`).
     */
    publicSecurityGroup?: pulumi.Input<string>;
    /**
     * ❗ The parent security group name. Please use the `securityGroupId` argument along the exoscale*security*group data source instead.
     *
     * @deprecated Deprecated in favor of `securityGroupId`
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * ❗ The parent exoscale*security*group ID.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * ❗ A `TCP`/`UDP` port range to match.
     */
    startPort?: pulumi.Input<number>;
    /**
     * ❗ The traffic direction to match (`INGRESS` or `EGRESS`).
     */
    type?: pulumi.Input<string>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`publicSecurityGroup`/`userSecurityGroupId`). Please use the `userSecurityGroupId` argument along the exoscale*security*group data source instead.
     *
     * @deprecated Deprecated in favor of `userSecurityGroupId`
     */
    userSecurityGroup?: pulumi.Input<string>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`publicSecurityGroup`/`user_security_group)`).
     */
    userSecurityGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroupRule resource.
 */
export interface SecurityGroupRuleArgs {
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)) to match (conflicts with `publicSecurityGroup`/`userSecurityGroup`/`userSecurityGroupId`).
     */
    cidr?: pulumi.Input<string>;
    /**
     * ❗ A free-form text describing the security group rule.
     */
    description?: pulumi.Input<string>;
    /**
     * ❗ A `TCP`/`UDP` port range to match.
     */
    endPort?: pulumi.Input<number>;
    /**
     * ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
     */
    icmpCode?: pulumi.Input<number>;
    /**
     * ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
     */
    icmpType?: pulumi.Input<number>;
    /**
     * ❗ The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
     */
    protocol?: pulumi.Input<string>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination public security group name to match (conflicts with `cidr`/`userSecurityGroup`/`userSecurityGroupId`).
     */
    publicSecurityGroup?: pulumi.Input<string>;
    /**
     * ❗ The parent security group name. Please use the `securityGroupId` argument along the exoscale*security*group data source instead.
     *
     * @deprecated Deprecated in favor of `securityGroupId`
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * ❗ The parent exoscale*security*group ID.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * ❗ A `TCP`/`UDP` port range to match.
     */
    startPort?: pulumi.Input<number>;
    /**
     * ❗ The traffic direction to match (`INGRESS` or `EGRESS`).
     */
    type: pulumi.Input<string>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`publicSecurityGroup`/`userSecurityGroupId`). Please use the `userSecurityGroupId` argument along the exoscale*security*group data source instead.
     *
     * @deprecated Deprecated in favor of `userSecurityGroupId`
     */
    userSecurityGroup?: pulumi.Input<string>;
    /**
     * ❗ An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`publicSecurityGroup`/`user_security_group)`).
     */
    userSecurityGroupId?: pulumi.Input<string>;
}
