// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Zitadel
{
    /// <summary>
    /// Resource representing the custom privacy policy of an organization.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zitadel = Pulumiverse.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Zitadel.PrivacyPolicy("default", new()
    ///     {
    ///         OrgId = defaultZitadelOrg.Id,
    ///         TosLink = "https://example.com/tos",
    ///         PrivacyLink = "https://example.com/privacy",
    ///         HelpLink = "https://example.com/help",
    ///         SupportEmail = "support@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// bash The resource can be imported using the ID format `&lt;[org_id]&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/privacyPolicy:PrivacyPolicy imported '123456789012345678'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/privacyPolicy:PrivacyPolicy")]
    public partial class PrivacyPolicy : global::Pulumi.CustomResource
    {
        [Output("helpLink")]
        public Output<string?> HelpLink { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        [Output("privacyLink")]
        public Output<string?> PrivacyLink { get; private set; } = null!;

        [Output("supportEmail")]
        public Output<string?> SupportEmail { get; private set; } = null!;

        [Output("tosLink")]
        public Output<string?> TosLink { get; private set; } = null!;


        /// <summary>
        /// Create a PrivacyPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivacyPolicy(string name, PrivacyPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("zitadel:index/privacyPolicy:PrivacyPolicy", name, args ?? new PrivacyPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivacyPolicy(string name, Input<string> id, PrivacyPolicyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/privacyPolicy:PrivacyPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivacyPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivacyPolicy Get(string name, Input<string> id, PrivacyPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivacyPolicy(name, id, state, options);
        }
    }

    public sealed class PrivacyPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("helpLink")]
        public Input<string>? HelpLink { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("privacyLink")]
        public Input<string>? PrivacyLink { get; set; }

        [Input("supportEmail")]
        public Input<string>? SupportEmail { get; set; }

        [Input("tosLink")]
        public Input<string>? TosLink { get; set; }

        public PrivacyPolicyArgs()
        {
        }
        public static new PrivacyPolicyArgs Empty => new PrivacyPolicyArgs();
    }

    public sealed class PrivacyPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("helpLink")]
        public Input<string>? HelpLink { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("privacyLink")]
        public Input<string>? PrivacyLink { get; set; }

        [Input("supportEmail")]
        public Input<string>? SupportEmail { get; set; }

        [Input("tosLink")]
        public Input<string>? TosLink { get; set; }

        public PrivacyPolicyState()
        {
        }
        public static new PrivacyPolicyState Empty => new PrivacyPolicyState();
    }
}
