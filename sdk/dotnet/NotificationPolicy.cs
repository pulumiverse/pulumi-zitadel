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
    /// Resource representing the custom notification policy of an organization.
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
    ///     var @default = new Zitadel.NotificationPolicy("default", new()
    ///     {
    ///         OrgId = defaultZitadelOrg.Id,
    ///         PasswordChange = false,
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
    ///  $ pulumi import zitadel:index/notificationPolicy:NotificationPolicy imported '123456789012345678'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/notificationPolicy:NotificationPolicy")]
    public partial class NotificationPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// Send notification if a user changes his password
        /// </summary>
        [Output("passwordChange")]
        public Output<bool> PasswordChange { get; private set; } = null!;


        /// <summary>
        /// Create a NotificationPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotificationPolicy(string name, NotificationPolicyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/notificationPolicy:NotificationPolicy", name, args ?? new NotificationPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotificationPolicy(string name, Input<string> id, NotificationPolicyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/notificationPolicy:NotificationPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NotificationPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotificationPolicy Get(string name, Input<string> id, NotificationPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new NotificationPolicy(name, id, state, options);
        }
    }

    public sealed class NotificationPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Send notification if a user changes his password
        /// </summary>
        [Input("passwordChange", required: true)]
        public Input<bool> PasswordChange { get; set; } = null!;

        public NotificationPolicyArgs()
        {
        }
        public static new NotificationPolicyArgs Empty => new NotificationPolicyArgs();
    }

    public sealed class NotificationPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Send notification if a user changes his password
        /// </summary>
        [Input("passwordChange")]
        public Input<bool>? PasswordChange { get; set; }

        public NotificationPolicyState()
        {
        }
        public static new NotificationPolicyState Empty => new NotificationPolicyState();
    }
}
