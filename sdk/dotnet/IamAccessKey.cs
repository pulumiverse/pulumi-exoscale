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
    ///     var mySosAccessKey = new Exoscale.IamAccessKey("mySosAccessKey", new()
    ///     {
    ///         Operations = new[]
    ///         {
    ///             "get-sos-object",
    ///             "list-sos-bucket",
    ///         },
    ///         Resources = new[]
    ///         {
    ///             "sos/bucket:my-bucket",
    ///         },
    ///     });
    /// 
    ///     var mySksAccessKey = new Exoscale.IamAccessKey("mySksAccessKey", new()
    ///     {
    ///         Tags = new[]
    ///         {
    ///             "sks",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Please refer to the examples
    /// directory for complete configuration examples.
    /// 
    /// &gt; **NOTE:** You can retrieve the list of available operations and tags using the [Exoscale CLI](https://github.com/exoscale/cli/): `exo iam access-key list-operations`.
    /// </summary>
    [ExoscaleResourceType("exoscale:index/iamAccessKey:IamAccessKey")]
    public partial class IamAccessKey : global::Pulumi.CustomResource
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
        /// Create a IamAccessKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamAccessKey(string name, IamAccessKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("exoscale:index/iamAccessKey:IamAccessKey", name, args ?? new IamAccessKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamAccessKey(string name, Input<string> id, IamAccessKeyState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/iamAccessKey:IamAccessKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamAccessKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamAccessKey Get(string name, Input<string> id, IamAccessKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new IamAccessKey(name, id, state, options);
        }
    }

    public sealed class IamAccessKeyArgs : global::Pulumi.ResourceArgs
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

        public IamAccessKeyArgs()
        {
        }
        public static new IamAccessKeyArgs Empty => new IamAccessKeyArgs();
    }

    public sealed class IamAccessKeyState : global::Pulumi.ResourceArgs
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

        public IamAccessKeyState()
        {
        }
        public static new IamAccessKeyState Empty => new IamAccessKeyState();
    }
}
