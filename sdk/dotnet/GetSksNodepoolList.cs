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
    public static class GetSksNodepoolList
    {
        public static Task<GetSksNodepoolListResult> InvokeAsync(GetSksNodepoolListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSksNodepoolListResult>("exoscale:index/getSksNodepoolList:getSksNodepoolList", args ?? new GetSksNodepoolListArgs(), options.WithDefaults());

        public static Output<GetSksNodepoolListResult> Invoke(GetSksNodepoolListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSksNodepoolListResult>("exoscale:index/getSksNodepoolList:getSksNodepoolList", args ?? new GetSksNodepoolListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSksNodepoolListArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("createdAt")]
        public string? CreatedAt { get; set; }

        [Input("deployTargetId")]
        public string? DeployTargetId { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("diskSize")]
        public int? DiskSize { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("instancePoolId")]
        public string? InstancePoolId { get; set; }

        [Input("instancePrefix")]
        public string? InstancePrefix { get; set; }

        [Input("instanceType")]
        public string? InstanceType { get; set; }

        [Input("labels")]
        private Dictionary<string, string>? _labels;
        public Dictionary<string, string> Labels
        {
            get => _labels ?? (_labels = new Dictionary<string, string>());
            set => _labels = value;
        }

        [Input("name")]
        public string? Name { get; set; }

        [Input("size")]
        public int? Size { get; set; }

        [Input("state")]
        public string? State { get; set; }

        [Input("storageLvm")]
        public bool? StorageLvm { get; set; }

        [Input("taints")]
        private Dictionary<string, string>? _taints;
        public Dictionary<string, string> Taints
        {
            get => _taints ?? (_taints = new Dictionary<string, string>());
            set => _taints = value;
        }

        [Input("templateId")]
        public string? TemplateId { get; set; }

        [Input("version")]
        public string? Version { get; set; }

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetSksNodepoolListArgs()
        {
        }
        public static new GetSksNodepoolListArgs Empty => new GetSksNodepoolListArgs();
    }

    public sealed class GetSksNodepoolListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("instancePoolId")]
        public Input<string>? InstancePoolId { get; set; }

        [Input("instancePrefix")]
        public Input<string>? InstancePrefix { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("size")]
        public Input<int>? Size { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("storageLvm")]
        public Input<bool>? StorageLvm { get; set; }

        [Input("taints")]
        private InputMap<string>? _taints;
        public InputMap<string> Taints
        {
            get => _taints ?? (_taints = new InputMap<string>());
            set => _taints = value;
        }

        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetSksNodepoolListInvokeArgs()
        {
        }
        public static new GetSksNodepoolListInvokeArgs Empty => new GetSksNodepoolListInvokeArgs();
    }


    [OutputType]
    public sealed class GetSksNodepoolListResult
    {
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? ClusterId;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? DeployTargetId;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Match against this int
        /// </summary>
        public readonly int? DiskSize;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? InstancePoolId;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? InstancePrefix;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GetSksNodepoolListNodepoolResult> Nodepools;
        /// <summary>
        /// Match against this int
        /// </summary>
        public readonly int? Size;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Match against this bool
        /// </summary>
        public readonly bool? StorageLvm;
        /// <summary>
        /// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Taints;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? TemplateId;
        /// <summary>
        /// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        /// </summary>
        public readonly string? Version;
        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetSksNodepoolListResult(
            string? clusterId,

            string? createdAt,

            string? deployTargetId,

            string? description,

            int? diskSize,

            string? id,

            string? instancePoolId,

            string? instancePrefix,

            string? instanceType,

            ImmutableDictionary<string, string>? labels,

            string? name,

            ImmutableArray<Outputs.GetSksNodepoolListNodepoolResult> nodepools,

            int? size,

            string? state,

            bool? storageLvm,

            ImmutableDictionary<string, string>? taints,

            string? templateId,

            string? version,

            string zone)
        {
            ClusterId = clusterId;
            CreatedAt = createdAt;
            DeployTargetId = deployTargetId;
            Description = description;
            DiskSize = diskSize;
            Id = id;
            InstancePoolId = instancePoolId;
            InstancePrefix = instancePrefix;
            InstanceType = instanceType;
            Labels = labels;
            Name = name;
            Nodepools = nodepools;
            Size = size;
            State = state;
            StorageLvm = storageLvm;
            Taints = taints;
            TemplateId = templateId;
            Version = version;
            Zone = zone;
        }
    }
}
