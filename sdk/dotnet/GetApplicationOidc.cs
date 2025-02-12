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
    public static class GetApplicationOidc
    {
        /// <summary>
        /// Datasource representing an OIDC application belonging to a project, with all configuration possibilities.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Zitadel.GetApplicationOidc.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Default.Id,
        ///         ProjectId = data.Zitadel_project.Default.Id,
        ///         AppId = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationOidcResult> InvokeAsync(GetApplicationOidcArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationOidcResult>("zitadel:index/getApplicationOidc:getApplicationOidc", args ?? new GetApplicationOidcArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing an OIDC application belonging to a project, with all configuration possibilities.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Zitadel.GetApplicationOidc.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Default.Id,
        ///         ProjectId = data.Zitadel_project.Default.Id,
        ///         AppId = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApplicationOidcResult> Invoke(GetApplicationOidcInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationOidcResult>("zitadel:index/getApplicationOidc:getApplicationOidc", args ?? new GetApplicationOidcInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationOidcArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetApplicationOidcArgs()
        {
        }
        public static new GetApplicationOidcArgs Empty => new GetApplicationOidcArgs();
    }

    public sealed class GetApplicationOidcInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

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

        public GetApplicationOidcInvokeArgs()
        {
        }
        public static new GetApplicationOidcInvokeArgs Empty => new GetApplicationOidcInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationOidcResult
    {
        /// <summary>
        /// Access token role assertion
        /// </summary>
        public readonly bool AccessTokenRoleAssertion;
        /// <summary>
        /// Access token type
        /// </summary>
        public readonly string AccessTokenType;
        /// <summary>
        /// Additional origins
        /// </summary>
        public readonly ImmutableArray<string> AdditionalOrigins;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// App type
        /// </summary>
        public readonly string AppType;
        /// <summary>
        /// Auth method type
        /// </summary>
        public readonly string AuthMethodType;
        public readonly string ClientId;
        /// <summary>
        /// Clockskew
        /// </summary>
        public readonly string ClockSkew;
        /// <summary>
        /// Dev mode
        /// </summary>
        public readonly bool DevMode;
        /// <summary>
        /// Grant types
        /// </summary>
        public readonly ImmutableArray<string> GrantTypes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID token role assertion
        /// </summary>
        public readonly bool IdTokenRoleAssertion;
        /// <summary>
        /// Token userinfo assertion
        /// </summary>
        public readonly bool IdTokenUserinfoAssertion;
        /// <summary>
        /// Name of the application
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// Post logout redirect URIs
        /// </summary>
        public readonly ImmutableArray<string> PostLogoutRedirectUris;
        /// <summary>
        /// ID of the project
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// RedirectURIs
        /// </summary>
        public readonly ImmutableArray<string> RedirectUris;
        /// <summary>
        /// Response type
        /// </summary>
        public readonly ImmutableArray<string> ResponseTypes;
        /// <summary>
        /// Version
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetApplicationOidcResult(
            bool accessTokenRoleAssertion,

            string accessTokenType,

            ImmutableArray<string> additionalOrigins,

            string appId,

            string appType,

            string authMethodType,

            string clientId,

            string clockSkew,

            bool devMode,

            ImmutableArray<string> grantTypes,

            string id,

            bool idTokenRoleAssertion,

            bool idTokenUserinfoAssertion,

            string name,

            string? orgId,

            ImmutableArray<string> postLogoutRedirectUris,

            string projectId,

            ImmutableArray<string> redirectUris,

            ImmutableArray<string> responseTypes,

            string version)
        {
            AccessTokenRoleAssertion = accessTokenRoleAssertion;
            AccessTokenType = accessTokenType;
            AdditionalOrigins = additionalOrigins;
            AppId = appId;
            AppType = appType;
            AuthMethodType = authMethodType;
            ClientId = clientId;
            ClockSkew = clockSkew;
            DevMode = devMode;
            GrantTypes = grantTypes;
            Id = id;
            IdTokenRoleAssertion = idTokenRoleAssertion;
            IdTokenUserinfoAssertion = idTokenUserinfoAssertion;
            Name = name;
            OrgId = orgId;
            PostLogoutRedirectUris = postLogoutRedirectUris;
            ProjectId = projectId;
            RedirectUris = redirectUris;
            ResponseTypes = responseTypes;
            Version = version;
        }
    }
}
