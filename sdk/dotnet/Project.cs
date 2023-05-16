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
    /// Resource representing the project, which can then be granted to different organizations or users directly, containing different applications.
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
    ///     var project = new Zitadel.Project("project", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         ProjectRoleAssertion = true,
    ///         ProjectRoleCheck = true,
    ///         HasProjectCheck = true,
    ///         PrivateLabelingSetting = "PRIVATE_LABELING_SETTING_ENFORCE_PROJECT_RESOURCE_OWNER_POLICY",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/project:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ZITADEL checks if the org of the user has permission to this project
        /// </summary>
        [Output("hasProjectCheck")]
        public Output<bool?> HasProjectCheck { get; private set; } = null!;

        /// <summary>
        /// Name of the project
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Organization in which the project is located
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Defines from where the private labeling should be triggered, supported values: PRIVATE*LABELING*SETTING*UNSPECIFIED, PRIVATE*LABELING*SETTING*ENFORCE*PROJECT*RESOURCE*OWNER*POLICY, PRIVATE*LABELING*SETTING*ALLOW*LOGIN*USER*RESOURCE*OWNER*POLICY
        /// </summary>
        [Output("privateLabelingSetting")]
        public Output<string?> PrivateLabelingSetting { get; private set; } = null!;

        /// <summary>
        /// describes if roles of user should be added in token
        /// </summary>
        [Output("projectRoleAssertion")]
        public Output<bool?> ProjectRoleAssertion { get; private set; } = null!;

        /// <summary>
        /// ZITADEL checks if the user has at least one on this project
        /// </summary>
        [Output("projectRoleCheck")]
        public Output<bool?> ProjectRoleCheck { get; private set; } = null!;

        /// <summary>
        /// State of the project
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/project:Project", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ZITADEL checks if the org of the user has permission to this project
        /// </summary>
        [Input("hasProjectCheck")]
        public Input<bool>? HasProjectCheck { get; set; }

        /// <summary>
        /// Name of the project
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Organization in which the project is located
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// Defines from where the private labeling should be triggered, supported values: PRIVATE*LABELING*SETTING*UNSPECIFIED, PRIVATE*LABELING*SETTING*ENFORCE*PROJECT*RESOURCE*OWNER*POLICY, PRIVATE*LABELING*SETTING*ALLOW*LOGIN*USER*RESOURCE*OWNER*POLICY
        /// </summary>
        [Input("privateLabelingSetting")]
        public Input<string>? PrivateLabelingSetting { get; set; }

        /// <summary>
        /// describes if roles of user should be added in token
        /// </summary>
        [Input("projectRoleAssertion")]
        public Input<bool>? ProjectRoleAssertion { get; set; }

        /// <summary>
        /// ZITADEL checks if the user has at least one on this project
        /// </summary>
        [Input("projectRoleCheck")]
        public Input<bool>? ProjectRoleCheck { get; set; }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }

    public sealed class ProjectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ZITADEL checks if the org of the user has permission to this project
        /// </summary>
        [Input("hasProjectCheck")]
        public Input<bool>? HasProjectCheck { get; set; }

        /// <summary>
        /// Name of the project
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Organization in which the project is located
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Defines from where the private labeling should be triggered, supported values: PRIVATE*LABELING*SETTING*UNSPECIFIED, PRIVATE*LABELING*SETTING*ENFORCE*PROJECT*RESOURCE*OWNER*POLICY, PRIVATE*LABELING*SETTING*ALLOW*LOGIN*USER*RESOURCE*OWNER*POLICY
        /// </summary>
        [Input("privateLabelingSetting")]
        public Input<string>? PrivateLabelingSetting { get; set; }

        /// <summary>
        /// describes if roles of user should be added in token
        /// </summary>
        [Input("projectRoleAssertion")]
        public Input<bool>? ProjectRoleAssertion { get; set; }

        /// <summary>
        /// ZITADEL checks if the user has at least one on this project
        /// </summary>
        [Input("projectRoleCheck")]
        public Input<bool>? ProjectRoleCheck { get; set; }

        /// <summary>
        /// State of the project
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public ProjectState()
        {
        }
        public static new ProjectState Empty => new ProjectState();
    }
}
