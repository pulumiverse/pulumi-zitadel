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
    /// Resource representing a GitLab Self Hosted IdP on the organization.
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
    ///     var gitlabSelfHosted = new Zitadel.OrgIdpGitlabSelfHosted("gitlabSelfHosted", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         ClientId = "15765e...",
    ///         ClientSecret = "*****abcxyz",
    ///         Scopes = new[]
    ///         {
    ///             "openid",
    ///             "profile",
    ///             "email",
    ///         },
    ///         Issuer = "https://my.issuer",
    ///         IsLinkingAllowed = false,
    ///         IsCreationAllowed = true,
    ///         IsAutoCreation = false,
    ///         IsAutoUpdate = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/orgIdpGitlabSelfHosted:OrgIdpGitlabSelfHosted")]
    public partial class OrgIdpGitlabSelfHosted : global::Pulumi.CustomResource
    {
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Output("isAutoCreation")]
        public Output<bool> IsAutoCreation { get; private set; } = null!;

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Output("isAutoUpdate")]
        public Output<bool> IsAutoUpdate { get; private set; } = null!;

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Output("isCreationAllowed")]
        public Output<bool> IsCreationAllowed { get; private set; } = null!;

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Output("isLinkingAllowed")]
        public Output<bool> IsLinkingAllowed { get; private set; } = null!;

        /// <summary>
        /// the providers issuer
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;


        /// <summary>
        /// Create a OrgIdpGitlabSelfHosted resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrgIdpGitlabSelfHosted(string name, OrgIdpGitlabSelfHostedArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/orgIdpGitlabSelfHosted:OrgIdpGitlabSelfHosted", name, args ?? new OrgIdpGitlabSelfHostedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrgIdpGitlabSelfHosted(string name, Input<string> id, OrgIdpGitlabSelfHostedState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/orgIdpGitlabSelfHosted:OrgIdpGitlabSelfHosted", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrgIdpGitlabSelfHosted resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrgIdpGitlabSelfHosted Get(string name, Input<string> id, OrgIdpGitlabSelfHostedState? state = null, CustomResourceOptions? options = null)
        {
            return new OrgIdpGitlabSelfHosted(name, id, state, options);
        }
    }

    public sealed class OrgIdpGitlabSelfHostedArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Input("isAutoCreation", required: true)]
        public Input<bool> IsAutoCreation { get; set; } = null!;

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Input("isAutoUpdate", required: true)]
        public Input<bool> IsAutoUpdate { get; set; } = null!;

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Input("isCreationAllowed", required: true)]
        public Input<bool> IsCreationAllowed { get; set; } = null!;

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Input("isLinkingAllowed", required: true)]
        public Input<bool> IsLinkingAllowed { get; set; } = null!;

        /// <summary>
        /// the providers issuer
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public OrgIdpGitlabSelfHostedArgs()
        {
        }
        public static new OrgIdpGitlabSelfHostedArgs Empty => new OrgIdpGitlabSelfHostedArgs();
    }

    public sealed class OrgIdpGitlabSelfHostedState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Input("isAutoCreation")]
        public Input<bool>? IsAutoCreation { get; set; }

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Input("isAutoUpdate")]
        public Input<bool>? IsAutoUpdate { get; set; }

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Input("isCreationAllowed")]
        public Input<bool>? IsCreationAllowed { get; set; }

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Input("isLinkingAllowed")]
        public Input<bool>? IsLinkingAllowed { get; set; }

        /// <summary>
        /// the providers issuer
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public OrgIdpGitlabSelfHostedState()
        {
        }
        public static new OrgIdpGitlabSelfHostedState Empty => new OrgIdpGitlabSelfHostedState();
    }
}
