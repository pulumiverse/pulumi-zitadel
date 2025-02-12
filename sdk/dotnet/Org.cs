// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace scoretechnologies.Zitadel
{
    /// <summary>
    /// Resource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zitadel = scoretechnologies.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Zitadel.Org("default");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// bash The resource can be imported using the ID format `&lt;id&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/org:Org imported '123456789012345678'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/org:Org")]
    public partial class Org : global::Pulumi.CustomResource
    {
        /// <summary>
        /// True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
        /// </summary>
        [Output("isDefault")]
        public Output<bool?> IsDefault { get; private set; } = null!;

        /// <summary>
        /// Name of the org
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Primary domain of the org
        /// </summary>
        [Output("primaryDomain")]
        public Output<string> PrimaryDomain { get; private set; } = null!;

        /// <summary>
        /// State of the org
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Org resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Org(string name, OrgArgs? args = null, CustomResourceOptions? options = null)
            : base("zitadel:index/org:Org", name, args ?? new OrgArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Org(string name, Input<string> id, OrgState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/org:Org", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/scoretechnologies",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Org resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Org Get(string name, Input<string> id, OrgState? state = null, CustomResourceOptions? options = null)
        {
            return new Org(name, id, state, options);
        }
    }

    public sealed class OrgArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Name of the org
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public OrgArgs()
        {
        }
        public static new OrgArgs Empty => new OrgArgs();
    }

    public sealed class OrgState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Name of the org
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Primary domain of the org
        /// </summary>
        [Input("primaryDomain")]
        public Input<string>? PrimaryDomain { get; set; }

        /// <summary>
        /// State of the org
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public OrgState()
        {
        }
        public static new OrgState Empty => new OrgState();
    }
}
