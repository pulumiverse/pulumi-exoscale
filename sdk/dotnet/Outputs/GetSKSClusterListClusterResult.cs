// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Outputs
{

    [OutputType]
    public sealed class GetSKSClusterListClusterResult
    {
        public readonly ImmutableArray<string> Addons;
        public readonly string AggregationCa;
        public readonly bool? AutoUpgrade;
        public readonly string? Cni;
        public readonly string ControlPlaneCa;
        public readonly string CreatedAt;
        public readonly string? Description;
        public readonly string Endpoint;
        public readonly bool? ExoscaleCcm;
        public readonly string? Id;
        public readonly string KubeletCa;
        public readonly ImmutableDictionary<string, string>? Labels;
        public readonly bool? MetricsServer;
        public readonly string? Name;
        public readonly ImmutableArray<string> Nodepools;
        public readonly Outputs.GetSKSClusterListClusterOidcResult Oidc;
        public readonly string? ServiceLevel;
        public readonly string State;
        public readonly string Version;
        public readonly string Zone;

        [OutputConstructor]
        private GetSKSClusterListClusterResult(
            ImmutableArray<string> addons,

            string aggregationCa,

            bool? autoUpgrade,

            string? cni,

            string controlPlaneCa,

            string createdAt,

            string? description,

            string endpoint,

            bool? exoscaleCcm,

            string? id,

            string kubeletCa,

            ImmutableDictionary<string, string>? labels,

            bool? metricsServer,

            string? name,

            ImmutableArray<string> nodepools,

            Outputs.GetSKSClusterListClusterOidcResult oidc,

            string? serviceLevel,

            string state,

            string version,

            string zone)
        {
            Addons = addons;
            AggregationCa = aggregationCa;
            AutoUpgrade = autoUpgrade;
            Cni = cni;
            ControlPlaneCa = controlPlaneCa;
            CreatedAt = createdAt;
            Description = description;
            Endpoint = endpoint;
            ExoscaleCcm = exoscaleCcm;
            Id = id;
            KubeletCa = kubeletCa;
            Labels = labels;
            MetricsServer = metricsServer;
            Name = name;
            Nodepools = nodepools;
            Oidc = oidc;
            ServiceLevel = serviceLevel;
            State = state;
            Version = version;
            Zone = zone;
        }
    }
}
