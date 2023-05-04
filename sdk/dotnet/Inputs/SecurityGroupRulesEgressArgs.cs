// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Inputs
{

    public sealed class SecurityGroupRulesEgressArgs : global::Pulumi.ResourceArgs
    {
        [Input("cidrLists")]
        private InputList<string>? _cidrLists;
        public InputList<string> CidrLists
        {
            get => _cidrLists ?? (_cidrLists = new InputList<string>());
            set => _cidrLists = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("icmpCode")]
        public Input<int>? IcmpCode { get; set; }

        [Input("icmpType")]
        public Input<int>? IcmpType { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("ports")]
        private InputList<string>? _ports;
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("userSecurityGroupLists")]
        private InputList<string>? _userSecurityGroupLists;
        public InputList<string> UserSecurityGroupLists
        {
            get => _userSecurityGroupLists ?? (_userSecurityGroupLists = new InputList<string>());
            set => _userSecurityGroupLists = value;
        }

        public SecurityGroupRulesEgressArgs()
        {
        }
        public static new SecurityGroupRulesEgressArgs Empty => new SecurityGroupRulesEgressArgs();
    }
}
