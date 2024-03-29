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
    /// Manage Exoscale [Anti-Affinity Groups](https://community.exoscale.com/documentation/compute/anti-affinity-groups/).
    /// 
    /// Corresponding data source: exoscale_anti_affinity_group.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Exoscale = Pulumiverse.Exoscale;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myAntiAffinityGroup = new Exoscale.AntiAffinityGroup("myAntiAffinityGroup", new()
    ///     {
    ///         Description = "Prevent compute instances to run on the same host",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// Please refer to the examples
    /// directory for complete configuration examples.
    /// 
    /// ## Import
    /// 
    /// An existing anti-affinity group may be imported by `&lt;ID&gt;`:
    /// 
    /// ```sh
    /// $ pulumi import exoscale:index/antiAffinityGroup:AntiAffinityGroup \
    /// ```
    /// 
    ///   exoscale_anti_affinity_group.my_anti_affinity_group \
    /// 
    ///   f81d4fae-7dec-11d0-a765-00a0c91e6bf6
    /// </summary>
    [ExoscaleResourceType("exoscale:index/antiAffinityGroup:AntiAffinityGroup")]
    public partial class AntiAffinityGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ❗ A free-form text describing the group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ❗ The anti-affinity group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a AntiAffinityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AntiAffinityGroup(string name, AntiAffinityGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("exoscale:index/antiAffinityGroup:AntiAffinityGroup", name, args ?? new AntiAffinityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AntiAffinityGroup(string name, Input<string> id, AntiAffinityGroupState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/antiAffinityGroup:AntiAffinityGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AntiAffinityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AntiAffinityGroup Get(string name, Input<string> id, AntiAffinityGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AntiAffinityGroup(name, id, state, options);
        }
    }

    public sealed class AntiAffinityGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ❗ A free-form text describing the group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ❗ The anti-affinity group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AntiAffinityGroupArgs()
        {
        }
        public static new AntiAffinityGroupArgs Empty => new AntiAffinityGroupArgs();
    }

    public sealed class AntiAffinityGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ❗ A free-form text describing the group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ❗ The anti-affinity group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AntiAffinityGroupState()
        {
        }
        public static new AntiAffinityGroupState Empty => new AntiAffinityGroupState();
    }
}
