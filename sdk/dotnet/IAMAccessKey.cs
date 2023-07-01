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
    [ExoscaleResourceType("exoscale:index/iAMAccessKey:IAMAccessKey")]
    public partial class IAMAccessKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IAM access key (identifier).
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// ❗ The IAM access key name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ❗ A list of API operations to restrict the key to.
        /// </summary>
        [Output("operations")]
        public Output<ImmutableArray<string>> Operations { get; private set; } = null!;

        /// <summary>
        /// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`&lt;domain&gt;/&lt;type&gt;:&lt;name&gt;`).
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<string>> Resources { get; private set; } = null!;

        /// <summary>
        /// The key secret.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        /// <summary>
        /// ❗ A list of tags to restrict the key to.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tagsOperations")]
        public Output<ImmutableArray<string>> TagsOperations { get; private set; } = null!;


        /// <summary>
        /// Create a IAMAccessKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMAccessKey(string name, IAMAccessKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("exoscale:index/iAMAccessKey:IAMAccessKey", name, args ?? new IAMAccessKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IAMAccessKey(string name, Input<string> id, IAMAccessKeyState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/iAMAccessKey:IAMAccessKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-exoscale",
                AdditionalSecretOutputs =
                {
                    "key",
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IAMAccessKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IAMAccessKey Get(string name, Input<string> id, IAMAccessKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new IAMAccessKey(name, id, state, options);
        }
    }

    public sealed class IAMAccessKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ❗ The IAM access key name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operations")]
        private InputList<string>? _operations;

        /// <summary>
        /// ❗ A list of API operations to restrict the key to.
        /// </summary>
        public InputList<string> Operations
        {
            get => _operations ?? (_operations = new InputList<string>());
            set => _operations = value;
        }

        [Input("resources")]
        private InputList<string>? _resources;

        /// <summary>
        /// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`&lt;domain&gt;/&lt;type&gt;:&lt;name&gt;`).
        /// </summary>
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// ❗ A list of tags to restrict the key to.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public IAMAccessKeyArgs()
        {
        }
        public static new IAMAccessKeyArgs Empty => new IAMAccessKeyArgs();
    }

    public sealed class IAMAccessKeyState : global::Pulumi.ResourceArgs
    {
        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// The IAM access key (identifier).
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// ❗ The IAM access key name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operations")]
        private InputList<string>? _operations;

        /// <summary>
        /// ❗ A list of API operations to restrict the key to.
        /// </summary>
        public InputList<string> Operations
        {
            get => _operations ?? (_operations = new InputList<string>());
            set => _operations = value;
        }

        [Input("resources")]
        private InputList<string>? _resources;

        /// <summary>
        /// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`&lt;domain&gt;/&lt;type&gt;:&lt;name&gt;`).
        /// </summary>
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// The key secret.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// ❗ A list of tags to restrict the key to.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tagsOperations")]
        private InputList<string>? _tagsOperations;
        public InputList<string> TagsOperations
        {
            get => _tagsOperations ?? (_tagsOperations = new InputList<string>());
            set => _tagsOperations = value;
        }

        public IAMAccessKeyState()
        {
        }
        public static new IAMAccessKeyState Empty => new IAMAccessKeyState();
    }
}
