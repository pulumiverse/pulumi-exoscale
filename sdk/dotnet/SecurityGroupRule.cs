// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale
{
    /// <summary>
    /// Manage Exoscale [Security Group](https://community.exoscale.com/product/networking/security-group/) Rules.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Exoscale = Pulumiverse.Exoscale;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mySecurityGroup = new Exoscale.SecurityGroup("mySecurityGroup");
    /// 
    ///     var mySecurityGroupRule = new Exoscale.SecurityGroupRule("mySecurityGroupRule", new()
    ///     {
    ///         SecurityGroupId = mySecurityGroup.Id,
    ///         Type = "INGRESS",
    ///         Protocol = "TCP",
    ///         Cidr = "0.0.0.0/0",
    ///         StartPort = 80,
    ///         EndPort = 80,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Please refer to the examples
    /// directory for complete configuration examples.
    /// 
    /// ## Import
    /// 
    /// An existing security group rule may be imported by `&lt;security-group-ID&gt;/&lt;security-group-rule-ID&gt;`:
    /// 
    /// ```sh
    /// $ pulumi import exoscale:index/securityGroupRule:SecurityGroupRule \ 
    /// ```
    /// 
    ///   exoscale_security_group_rule.my_security_group_rule \
    /// 
    ///   f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524
    /// </summary>
    [ExoscaleResourceType("exoscale:index/securityGroupRule:SecurityGroupRule")]
    public partial class SecurityGroupRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)) to match (conflicts with `public_security_group`/`user_security_group`/`user_security_group_id`).
        /// </summary>
        [Output("cidr")]
        public Output<string?> Cidr { get; private set; } = null!;

        /// <summary>
        /// ❗ A free-form text describing the security group rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ❗ A `TCP`/`UDP` port range to match.
        /// </summary>
        [Output("endPort")]
        public Output<int?> EndPort { get; private set; } = null!;

        /// <summary>
        /// ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
        /// </summary>
        [Output("icmpCode")]
        public Output<int?> IcmpCode { get; private set; } = null!;

        /// <summary>
        /// ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
        /// </summary>
        [Output("icmpType")]
        public Output<int?> IcmpType { get; private set; } = null!;

        /// <summary>
        /// ❗ The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination public security group name to match (conflicts with `cidr`/`user_security_group`/`user_security_group_id`).
        /// </summary>
        [Output("publicSecurityGroup")]
        public Output<string> PublicSecurityGroup { get; private set; } = null!;

        /// <summary>
        /// ❗ The parent security group name. Please use the `security_group_id` argument along the exoscale*security*group data source instead.
        /// </summary>
        [Output("securityGroup")]
        public Output<string> SecurityGroup { get; private set; } = null!;

        /// <summary>
        /// ❗ The parent exoscale*security*group ID.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// ❗ A `TCP`/`UDP` port range to match.
        /// </summary>
        [Output("startPort")]
        public Output<int?> StartPort { get; private set; } = null!;

        /// <summary>
        /// ❗ The traffic direction to match (`INGRESS` or `EGRESS`).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`public_security_group`/`user_security_group_id`). Please use the `user_security_group_id` argument along the exoscale*security*group data source instead.
        /// </summary>
        [Output("userSecurityGroup")]
        public Output<string> UserSecurityGroup { get; private set; } = null!;

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`public_security_group`/`user_security_group)`).
        /// </summary>
        [Output("userSecurityGroupId")]
        public Output<string?> UserSecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupRule(string name, SecurityGroupRuleArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/securityGroupRule:SecurityGroupRule", name, args ?? new SecurityGroupRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupRule(string name, Input<string> id, SecurityGroupRuleState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/securityGroupRule:SecurityGroupRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-exoscale",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupRule Get(string name, Input<string> id, SecurityGroupRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroupRule(name, id, state, options);
        }
    }

    public sealed class SecurityGroupRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)) to match (conflicts with `public_security_group`/`user_security_group`/`user_security_group_id`).
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// ❗ A free-form text describing the security group rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ❗ A `TCP`/`UDP` port range to match.
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
        /// </summary>
        [Input("icmpCode")]
        public Input<int>? IcmpCode { get; set; }

        /// <summary>
        /// ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
        /// </summary>
        [Input("icmpType")]
        public Input<int>? IcmpType { get; set; }

        /// <summary>
        /// ❗ The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination public security group name to match (conflicts with `cidr`/`user_security_group`/`user_security_group_id`).
        /// </summary>
        [Input("publicSecurityGroup")]
        public Input<string>? PublicSecurityGroup { get; set; }

        /// <summary>
        /// ❗ The parent security group name. Please use the `security_group_id` argument along the exoscale*security*group data source instead.
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        /// <summary>
        /// ❗ The parent exoscale*security*group ID.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// ❗ A `TCP`/`UDP` port range to match.
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// ❗ The traffic direction to match (`INGRESS` or `EGRESS`).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`public_security_group`/`user_security_group_id`). Please use the `user_security_group_id` argument along the exoscale*security*group data source instead.
        /// </summary>
        [Input("userSecurityGroup")]
        public Input<string>? UserSecurityGroup { get; set; }

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`public_security_group`/`user_security_group)`).
        /// </summary>
        [Input("userSecurityGroupId")]
        public Input<string>? UserSecurityGroupId { get; set; }

        public SecurityGroupRuleArgs()
        {
        }
        public static new SecurityGroupRuleArgs Empty => new SecurityGroupRuleArgs();
    }

    public sealed class SecurityGroupRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)) to match (conflicts with `public_security_group`/`user_security_group`/`user_security_group_id`).
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// ❗ A free-form text describing the security group rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ❗ A `TCP`/`UDP` port range to match.
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
        /// </summary>
        [Input("icmpCode")]
        public Input<int>? IcmpCode { get; set; }

        /// <summary>
        /// ❗ An ICMP/ICMPv6 [type/code](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages) to match.
        /// </summary>
        [Input("icmpType")]
        public Input<int>? IcmpType { get; set; }

        /// <summary>
        /// ❗ The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination public security group name to match (conflicts with `cidr`/`user_security_group`/`user_security_group_id`).
        /// </summary>
        [Input("publicSecurityGroup")]
        public Input<string>? PublicSecurityGroup { get; set; }

        /// <summary>
        /// ❗ The parent security group name. Please use the `security_group_id` argument along the exoscale*security*group data source instead.
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        /// <summary>
        /// ❗ The parent exoscale*security*group ID.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// ❗ A `TCP`/`UDP` port range to match.
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// ❗ The traffic direction to match (`INGRESS` or `EGRESS`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`public_security_group`/`user_security_group_id`). Please use the `user_security_group_id` argument along the exoscale*security*group data source instead.
        /// </summary>
        [Input("userSecurityGroup")]
        public Input<string>? UserSecurityGroup { get; set; }

        /// <summary>
        /// ❗ An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`public_security_group`/`user_security_group)`).
        /// </summary>
        [Input("userSecurityGroupId")]
        public Input<string>? UserSecurityGroupId { get; set; }

        public SecurityGroupRuleState()
        {
        }
        public static new SecurityGroupRuleState Empty => new SecurityGroupRuleState();
    }
}
