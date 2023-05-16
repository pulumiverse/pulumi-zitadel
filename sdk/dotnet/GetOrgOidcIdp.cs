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
    public static class GetOrgOidcIdp
    {
        /// <summary>
        /// Datasource representing a generic OIDC IdP on the organization.
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
        ///     var orgOidcIdpOrgOidcIdp = Zitadel.GetOrgOidcIdp.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Org.Id,
        ///         IdpId = "177073612581240835",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["orgOidcIdp"] = orgOidcIdpOrgOidcIdp.Apply(getOrgOidcIdpResult =&gt; getOrgOidcIdpResult),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrgOidcIdpResult> InvokeAsync(GetOrgOidcIdpArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrgOidcIdpResult>("zitadel:index/getOrgOidcIdp:getOrgOidcIdp", args ?? new GetOrgOidcIdpArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing a generic OIDC IdP on the organization.
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
        ///     var orgOidcIdpOrgOidcIdp = Zitadel.GetOrgOidcIdp.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Org.Id,
        ///         IdpId = "177073612581240835",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["orgOidcIdp"] = orgOidcIdpOrgOidcIdp.Apply(getOrgOidcIdpResult =&gt; getOrgOidcIdpResult),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrgOidcIdpResult> Invoke(GetOrgOidcIdpInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrgOidcIdpResult>("zitadel:index/getOrgOidcIdp:getOrgOidcIdp", args ?? new GetOrgOidcIdpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgOidcIdpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("idpId", required: true)]
        public string IdpId { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public string OrgId { get; set; } = null!;

        public GetOrgOidcIdpArgs()
        {
        }
        public static new GetOrgOidcIdpArgs Empty => new GetOrgOidcIdpArgs();
    }

    public sealed class GetOrgOidcIdpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("idpId", required: true)]
        public Input<string> IdpId { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        public GetOrgOidcIdpInvokeArgs()
        {
        }
        public static new GetOrgOidcIdpInvokeArgs Empty => new GetOrgOidcIdpInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgOidcIdpResult
    {
        /// <summary>
        /// auto register for users from this idp
        /// </summary>
        public readonly bool AutoRegister;
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// definition which field is mapped to the display name of the user
        /// </summary>
        public readonly string DisplayNameMapping;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string IdpId;
        /// <summary>
        /// the oidc issuer of the identity provider
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// Name of the IDP
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public readonly ImmutableArray<string> Scopes;
        /// <summary>
        /// Some identity providers specify the styling of the button to their login
        /// </summary>
        public readonly string StylingType;
        /// <summary>
        /// definition which field is mapped to the email of the user
        /// </summary>
        public readonly string UsernameMapping;

        [OutputConstructor]
        private GetOrgOidcIdpResult(
            bool autoRegister,

            string clientId,

            string clientSecret,

            string displayNameMapping,

            string id,

            string idpId,

            string issuer,

            string name,

            string orgId,

            ImmutableArray<string> scopes,

            string stylingType,

            string usernameMapping)
        {
            AutoRegister = autoRegister;
            ClientId = clientId;
            ClientSecret = clientSecret;
            DisplayNameMapping = displayNameMapping;
            Id = id;
            IdpId = idpId;
            Issuer = issuer;
            Name = name;
            OrgId = orgId;
            Scopes = scopes;
            StylingType = stylingType;
            UsernameMapping = usernameMapping;
        }
    }
}
