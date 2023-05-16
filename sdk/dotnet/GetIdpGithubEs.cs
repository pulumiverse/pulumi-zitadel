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
    public static class GetIdpGithubEs
    {
        /// <summary>
        /// Datasource representing a GitHub Enterprise IDP on the instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var githubEs = Zitadel.GetIdpGithubEs.Invoke(new()
        ///     {
        ///         Id = "177073614158299139",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIdpGithubEsResult> InvokeAsync(GetIdpGithubEsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdpGithubEsResult>("zitadel:index/getIdpGithubEs:getIdpGithubEs", args ?? new GetIdpGithubEsArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing a GitHub Enterprise IDP on the instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var githubEs = Zitadel.GetIdpGithubEs.Invoke(new()
        ///     {
        ///         Id = "177073614158299139",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIdpGithubEsResult> Invoke(GetIdpGithubEsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIdpGithubEsResult>("zitadel:index/getIdpGithubEs:getIdpGithubEs", args ?? new GetIdpGithubEsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdpGithubEsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetIdpGithubEsArgs()
        {
        }
        public static new GetIdpGithubEsArgs Empty => new GetIdpGithubEsArgs();
    }

    public sealed class GetIdpGithubEsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetIdpGithubEsInvokeArgs()
        {
        }
        public static new GetIdpGithubEsInvokeArgs Empty => new GetIdpGithubEsInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdpGithubEsResult
    {
        /// <summary>
        /// the providers authorization endpoint
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
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public readonly ImmutableArray<string> Scopes;
        /// <summary>
        /// the providers token endpoint
        /// </summary>
        public readonly string TokenEndpoint;
        /// <summary>
        /// the providers user endpoint
        /// </summary>
        public readonly string UserEndpoint;

        [OutputConstructor]
        private GetIdpGithubEsResult(
            string authorizationEndpoint,

            string clientId,

            string clientSecret,

            string id,

            bool isAutoCreation,

            bool isAutoUpdate,

            bool isCreationAllowed,

            bool isLinkingAllowed,

            string name,

            ImmutableArray<string> scopes,

            string tokenEndpoint,

            string userEndpoint)
        {
            AuthorizationEndpoint = authorizationEndpoint;
            ClientId = clientId;
            ClientSecret = clientSecret;
            Id = id;
            IsAutoCreation = isAutoCreation;
            IsAutoUpdate = isAutoUpdate;
            IsCreationAllowed = isCreationAllowed;
            IsLinkingAllowed = isLinkingAllowed;
            Name = name;
            Scopes = scopes;
            TokenEndpoint = tokenEndpoint;
            UserEndpoint = userEndpoint;
        }
    }
}
