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
    /// <summary>
    /// Manage Exoscale [IAM](https://community.exoscale.com/documentation/iam/) Role.
    /// </summary>
    [ExoscaleResourceType("exoscale:index/iamRole:IamRole")]
    public partial class IamRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A free-form text describing the IAM Role
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Defines if IAM Role Policy is editable or not.
        /// </summary>
        [Output("editable")]
        public Output<bool> Editable { get; private set; } = null!;

        /// <summary>
        /// IAM Role labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// ❗Name of IAM Role.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// IAM Role permissions.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// IAM Policy.
        /// </summary>
        [Output("policy")]
        public Output<Outputs.IamRolePolicy> Policy { get; private set; } = null!;

        [Output("timeouts")]
        public Output<Outputs.IamRoleTimeouts?> Timeouts { get; private set; } = null!;


        /// <summary>
        /// Create a IamRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamRole(string name, IamRoleArgs? args = null, CustomResourceOptions? options = null)
            : base("exoscale:index/iamRole:IamRole", name, args ?? new IamRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamRole(string name, Input<string> id, IamRoleState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/iamRole:IamRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamRole Get(string name, Input<string> id, IamRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new IamRole(name, id, state, options);
        }
    }

    public sealed class IamRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-form text describing the IAM Role
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Defines if IAM Role Policy is editable or not.
        /// </summary>
        [Input("editable")]
        public Input<bool>? Editable { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// IAM Role labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// ❗Name of IAM Role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// IAM Role permissions.
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// IAM Policy.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.IamRolePolicyArgs>? Policy { get; set; }

        [Input("timeouts")]
        public Input<Inputs.IamRoleTimeoutsArgs>? Timeouts { get; set; }

        public IamRoleArgs()
        {
        }
        public static new IamRoleArgs Empty => new IamRoleArgs();
    }

    public sealed class IamRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-form text describing the IAM Role
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Defines if IAM Role Policy is editable or not.
        /// </summary>
        [Input("editable")]
        public Input<bool>? Editable { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// IAM Role labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// ❗Name of IAM Role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// IAM Role permissions.
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// IAM Policy.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.IamRolePolicyGetArgs>? Policy { get; set; }

        [Input("timeouts")]
        public Input<Inputs.IamRoleTimeoutsGetArgs>? Timeouts { get; set; }

        public IamRoleState()
        {
        }
        public static new IamRoleState Empty => new IamRoleState();
    }
}