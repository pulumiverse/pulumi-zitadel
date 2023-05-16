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
    /// Resource representing the grant of a project to a different organization, also containing the available roles which can be given to the members of the projectgrant.
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
    ///     var projectGrant = new Zitadel.ProjectGrant("projectGrant", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         ProjectId = zitadel_project.Project.Id,
    ///         GrantedOrgId = zitadel_org.Grantedorg.Id,
    ///         RoleKeys = new[]
    ///         {
    ///             zitadel_project_role.Project_role.Role_key,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/projectGrant:ProjectGrant")]
    public partial class ProjectGrant : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the organization granted the project
        /// </summary>
        [Output("grantedOrgId")]
        public Output<string> GrantedOrgId { get; private set; } = null!;

        /// <summary>
        /// ID of the organization which owns the resource
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// List of roles granted
        /// </summary>
        [Output("roleKeys")]
        public Output<ImmutableArray<string>> RoleKeys { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectGrant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectGrant(string name, ProjectGrantArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/projectGrant:ProjectGrant", name, args ?? new ProjectGrantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectGrant(string name, Input<string> id, ProjectGrantState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/projectGrant:ProjectGrant", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectGrant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectGrant Get(string name, Input<string> id, ProjectGrantState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectGrant(name, id, state, options);
        }
    }

    public sealed class ProjectGrantArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization granted the project
        /// </summary>
        [Input("grantedOrgId", required: true)]
        public Input<string> GrantedOrgId { get; set; } = null!;

        /// <summary>
        /// ID of the organization which owns the resource
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("roleKeys")]
        private InputList<string>? _roleKeys;

        /// <summary>
        /// List of roles granted
        /// </summary>
        public InputList<string> RoleKeys
        {
            get => _roleKeys ?? (_roleKeys = new InputList<string>());
            set => _roleKeys = value;
        }

        public ProjectGrantArgs()
        {
        }
        public static new ProjectGrantArgs Empty => new ProjectGrantArgs();
    }

    public sealed class ProjectGrantState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization granted the project
        /// </summary>
        [Input("grantedOrgId")]
        public Input<string>? GrantedOrgId { get; set; }

        /// <summary>
        /// ID of the organization which owns the resource
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("roleKeys")]
        private InputList<string>? _roleKeys;

        /// <summary>
        /// List of roles granted
        /// </summary>
        public InputList<string> RoleKeys
        {
            get => _roleKeys ?? (_roleKeys = new InputList<string>());
            set => _roleKeys = value;
        }

        public ProjectGrantState()
        {
        }
        public static new ProjectGrantState Empty => new ProjectGrantState();
    }
}
