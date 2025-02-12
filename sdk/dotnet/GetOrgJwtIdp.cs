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
    public static class GetOrgJwtIdp
    {
        /// <summary>
        /// Datasource representing a generic JWT IdP on the organization.
        /// </summary>
        public static Task<GetOrgJwtIdpResult> InvokeAsync(GetOrgJwtIdpArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrgJwtIdpResult>("zitadel:index/getOrgJwtIdp:getOrgJwtIdp", args ?? new GetOrgJwtIdpArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing a generic JWT IdP on the organization.
        /// </summary>
        public static Output<GetOrgJwtIdpResult> Invoke(GetOrgJwtIdpInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrgJwtIdpResult>("zitadel:index/getOrgJwtIdp:getOrgJwtIdp", args ?? new GetOrgJwtIdpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgJwtIdpArgs : global::Pulumi.InvokeArgs
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

        public GetOrgJwtIdpArgs()
        {
        }
        public static new GetOrgJwtIdpArgs Empty => new GetOrgJwtIdpArgs();
    }

    public sealed class GetOrgJwtIdpInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetOrgJwtIdpInvokeArgs()
        {
        }
        public static new GetOrgJwtIdpInvokeArgs Empty => new GetOrgJwtIdpInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgJwtIdpResult
    {
        /// <summary>
        /// auto register for users from this idp
        /// </summary>
        public readonly bool AutoRegister;
        /// <summary>
        /// the name of the header where the JWT is sent in, default is authorization
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string IdpId;
        /// <summary>
        /// the issuer of the jwt (for validation)
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// the endpoint where the jwt can be extracted
        /// </summary>
        public readonly string JwtEndpoint;
        /// <summary>
        /// the endpoint to the key (JWK) which are used to sign the JWT with
        /// </summary>
        public readonly string KeysEndpoint;
        /// <summary>
        /// Name of the IDP
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// Some identity providers specify the styling of the button to their login
        /// </summary>
        public readonly string StylingType;

        [OutputConstructor]
        private GetOrgJwtIdpResult(
            bool autoRegister,

            string headerName,

            string id,

            string idpId,

            string issuer,

            string jwtEndpoint,

            string keysEndpoint,

            string name,

            string orgId,

            string stylingType)
        {
            AutoRegister = autoRegister;
            HeaderName = headerName;
            Id = id;
            IdpId = idpId;
            Issuer = issuer;
            JwtEndpoint = jwtEndpoint;
            KeysEndpoint = keysEndpoint;
            Name = name;
            OrgId = orgId;
            StylingType = stylingType;
        }
    }
}
