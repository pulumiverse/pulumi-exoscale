// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale
{
    public static class GetComputeInstanceList
    {
        public static Task<GetComputeInstanceListResult> InvokeAsync(GetComputeInstanceListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetComputeInstanceListResult>("exoscale:index/getComputeInstanceList:getComputeInstanceList", args ?? new GetComputeInstanceListArgs(), options.WithDefaults());

        public static Output<GetComputeInstanceListResult> Invoke(GetComputeInstanceListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetComputeInstanceListResult>("exoscale:index/getComputeInstanceList:getComputeInstanceList", args ?? new GetComputeInstanceListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetComputeInstanceListArgs : global::Pulumi.InvokeArgs
    {
        [Input("createdAt")]
        public string? CreatedAt { get; set; }

        [Input("deployTargetId")]
        public string? DeployTargetId { get; set; }

        [Input("diskSize")]
        public int? DiskSize { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("ipv6")]
        public bool? Ipv6 { get; set; }

        [Input("ipv6Address")]
        public string? Ipv6Address { get; set; }

        [Input("labels")]
        private Dictionary<string, string>? _labels;
        public Dictionary<string, string> Labels
        {
            get => _labels ?? (_labels = new Dictionary<string, string>());
            set => _labels = value;
        }

        [Input("managerId")]
        public string? ManagerId { get; set; }

        [Input("managerType")]
        public string? ManagerType { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("publicIpAddress")]
        public string? PublicIpAddress { get; set; }

        [Input("reverseDns")]
        public string? ReverseDns { get; set; }

        [Input("sshKey")]
        public string? SshKey { get; set; }

        [Input("state")]
        public string? State { get; set; }

        [Input("templateId")]
        public string? TemplateId { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        [Input("userData")]
        public string? UserData { get; set; }

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetComputeInstanceListArgs()
        {
        }
        public static new GetComputeInstanceListArgs Empty => new GetComputeInstanceListArgs();
    }

    public sealed class GetComputeInstanceListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("managerId")]
        public Input<string>? ManagerId { get; set; }

        [Input("managerType")]
        public Input<string>? ManagerType { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("publicIpAddress")]
        public Input<string>? PublicIpAddress { get; set; }

        [Input("reverseDns")]
        public Input<string>? ReverseDns { get; set; }

        [Input("sshKey")]
        public Input<string>? SshKey { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetComputeInstanceListInvokeArgs()
        {
        }
        public static new GetComputeInstanceListInvokeArgs Empty => new GetComputeInstanceListInvokeArgs();
    }


    [OutputType]
    public sealed class GetComputeInstanceListResult
    {
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? DeployTargetId;
        /// <summary>
        /// Match against this int
        /// </summary>
        public readonly int? DiskSize;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The list of exoscale*compute*instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceListInstanceResult> Instances;
        /// <summary>
        /// Match against this bool
        /// </summary>
        public readonly bool? Ipv6;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? Ipv6Address;
        /// <summary>
        /// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? ManagerId;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? ManagerType;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? PublicIpAddress;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? ReverseDns;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? SshKey;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? TemplateId;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? UserData;
        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetComputeInstanceListResult(
            string? createdAt,

            string? deployTargetId,

            int? diskSize,

            string? id,

            ImmutableArray<Outputs.GetComputeInstanceListInstanceResult> instances,

            bool? ipv6,

            string? ipv6Address,

            ImmutableDictionary<string, string>? labels,

            string? managerId,

            string? managerType,

            string? name,

            string? publicIpAddress,

            string? reverseDns,

            string? sshKey,

            string? state,

            string? templateId,

            string? type,

            string? userData,

            string zone)
        {
            CreatedAt = createdAt;
            DeployTargetId = deployTargetId;
            DiskSize = diskSize;
            Id = id;
            Instances = instances;
            Ipv6 = ipv6;
            Ipv6Address = ipv6Address;
            Labels = labels;
            ManagerId = managerId;
            ManagerType = managerType;
            Name = name;
            PublicIpAddress = publicIpAddress;
            ReverseDns = reverseDns;
            SshKey = sshKey;
            State = state;
            TemplateId = templateId;
            Type = type;
            UserData = userData;
            Zone = zone;
        }
    }
}
