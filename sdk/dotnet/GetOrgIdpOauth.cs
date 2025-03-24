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
    public static class GetOrgIdpOauth
    {
        /// <summary>
        /// Datasource representing a generic OAuth2 IDP of the organization.
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
        ///     var @default = Zitadel.GetOrgIdpOauth.Invoke(new()
        ///     {
        ///         OrgId = defaultZitadelOrg.Id,
        ///         Id = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrgIdpOauthResult> InvokeAsync(GetOrgIdpOauthArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrgIdpOauthResult>("zitadel:index/getOrgIdpOauth:getOrgIdpOauth", args ?? new GetOrgIdpOauthArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing a generic OAuth2 IDP of the organization.
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
        ///     var @default = Zitadel.GetOrgIdpOauth.Invoke(new()
        ///     {
        ///         OrgId = defaultZitadelOrg.Id,
        ///         Id = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrgIdpOauthResult> Invoke(GetOrgIdpOauthInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrgIdpOauthResult>("zitadel:index/getOrgIdpOauth:getOrgIdpOauth", args ?? new GetOrgIdpOauthInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgIdpOauthArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        public GetOrgIdpOauthArgs()
        {
        }
        public static new GetOrgIdpOauthArgs Empty => new GetOrgIdpOauthArgs();
    }

    public sealed class GetOrgIdpOauthInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public GetOrgIdpOauthInvokeArgs()
        {
        }
        public static new GetOrgIdpOauthInvokeArgs Empty => new GetOrgIdpOauthInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgIdpOauthResult
    {
        /// <summary>
        /// The authorization endpoint
        /// </summary>
        public readonly string AuthorizationEndpoint;
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id attribute
        /// </summary>
        public readonly string IdAttribute;
        /// <summary>
        /// enabled if a new account in ZITADEL are created automatically on login with an external account
        /// </summary>
        public readonly bool IsAutoCreation;
        /// <summary>
        /// enabled if a the ZITADEL account fields are updated automatically on each login
        /// </summary>
        public readonly bool IsAutoUpdate;
        /// <summary>
        /// enabled if users are able to create a new account in ZITADEL when using an external account
        /// </summary>
        public readonly bool IsCreationAllowed;
        /// <summary>
        /// enabled if users are able to link an existing ZITADEL user with an external account
        /// </summary>
        public readonly bool IsLinkingAllowed;
        /// <summary>
        /// Name of the IDP
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public readonly ImmutableArray<string> Scopes;
        /// <summary>
        /// The token endpoint
        /// </summary>
        public readonly string TokenEndpoint;
        /// <summary>
        /// The user endpoint
        /// </summary>
        public readonly string UserEndpoint;

        [OutputConstructor]
        private GetOrgIdpOauthResult(
            string authorizationEndpoint,

            string clientId,

            string clientSecret,

            string id,

            string idAttribute,

            bool isAutoCreation,

            bool isAutoUpdate,

            bool isCreationAllowed,

            bool isLinkingAllowed,

            string name,

            string? orgId,

            ImmutableArray<string> scopes,

            string tokenEndpoint,

            string userEndpoint)
        {
            AuthorizationEndpoint = authorizationEndpoint;
            ClientId = clientId;
            ClientSecret = clientSecret;
            Id = id;
            IdAttribute = idAttribute;
            IsAutoCreation = isAutoCreation;
            IsAutoUpdate = isAutoUpdate;
            IsCreationAllowed = isCreationAllowed;
            IsLinkingAllowed = isLinkingAllowed;
            Name = name;
            OrgId = orgId;
            Scopes = scopes;
            TokenEndpoint = tokenEndpoint;
            UserEndpoint = userEndpoint;
        }
    }
}
