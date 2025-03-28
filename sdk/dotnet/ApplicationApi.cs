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
    /// Resource representing an API application belonging to a project, with all configuration possibilities.
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
    ///     var @default = new Zitadel.ApplicationApi("default", new()
    ///     {
    ///         OrgId = defaultZitadelOrg.Id,
    ///         ProjectId = defaultZitadelProject.Id,
    ///         Name = "applicationapi",
    ///         AuthMethodType = "API_AUTH_METHOD_TYPE_BASIC",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// bash The resource can be imported using the ID format `&lt;id:project_id[:org_id][:client_id][:client_secret]&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/applicationApi:ApplicationApi imported '123456789012345678:123456789012345678:123456789012345678:123456789012345678@zitadel:JuaDFFeOak5DGE655KCYPSAclSkbMVEJXXuX1lEMBT14eLMSs0A0qhafKX5SA2Df'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/applicationApi:ApplicationApi")]
    public partial class ApplicationApi : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        /// </summary>
        [Output("authMethodType")]
        public Output<string?> AuthMethodType { get; private set; } = null!;

        /// <summary>
        /// generated ID for this config
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// generated secret for this config
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Name of the application
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationApi(string name, ApplicationApiArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/applicationApi:ApplicationApi", name, args ?? new ApplicationApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationApi(string name, Input<string> id, ApplicationApiState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/applicationApi:ApplicationApi", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "clientId",
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationApi Get(string name, Input<string> id, ApplicationApiState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationApi(name, id, state, options);
        }
    }

    public sealed class ApplicationApiArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        /// </summary>
        [Input("authMethodType")]
        public Input<string>? AuthMethodType { get; set; }

        /// <summary>
        /// Name of the application
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public ApplicationApiArgs()
        {
        }
        public static new ApplicationApiArgs Empty => new ApplicationApiArgs();
    }

    public sealed class ApplicationApiState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        /// </summary>
        [Input("authMethodType")]
        public Input<string>? AuthMethodType { get; set; }

        [Input("clientId")]
        private Input<string>? _clientId;

        /// <summary>
        /// generated ID for this config
        /// </summary>
        public Input<string>? ClientId
        {
            get => _clientId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// generated secret for this config
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of the application
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public ApplicationApiState()
        {
        }
        public static new ApplicationApiState Empty => new ApplicationApiState();
    }
}
