// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Vavsab.Zitadel
{
    /// <summary>
    /// Resource representing the membership of a user on an organization, defined with the given role.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zitadel = Vavsab.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var orgMember = new Zitadel.OrgMember("orgMember", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         UserId = zitadel_human_user.Human_user.Id,
    ///         Roles = new[]
    ///         {
    ///             "ORG_OWNER",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/orgMember:OrgMember")]
    public partial class OrgMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// List of roles granted
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// ID of the user
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a OrgMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrgMember(string name, OrgMemberArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/orgMember:OrgMember", name, args ?? new OrgMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrgMember(string name, Input<string> id, OrgMemberState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/orgMember:OrgMember", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrgMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrgMember Get(string name, Input<string> id, OrgMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new OrgMember(name, id, state, options);
        }
    }

    public sealed class OrgMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles granted
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public OrgMemberArgs()
        {
        }
        public static new OrgMemberArgs Empty => new OrgMemberArgs();
    }

    public sealed class OrgMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles granted
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public OrgMemberState()
        {
        }
        public static new OrgMemberState Empty => new OrgMemberState();
    }
}
