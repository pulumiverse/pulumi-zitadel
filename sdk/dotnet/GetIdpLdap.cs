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
    public static class GetIdpLdap
    {
        /// <summary>
        /// Datasource representing an LDAP IDP on the instance.
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
        ///     var @default = Zitadel.GetIdpLdap.Invoke(new()
        ///     {
        ///         Id = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIdpLdapResult> InvokeAsync(GetIdpLdapArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdpLdapResult>("zitadel:index/getIdpLdap:getIdpLdap", args ?? new GetIdpLdapArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing an LDAP IDP on the instance.
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
        ///     var @default = Zitadel.GetIdpLdap.Invoke(new()
        ///     {
        ///         Id = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIdpLdapResult> Invoke(GetIdpLdapInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIdpLdapResult>("zitadel:index/getIdpLdap:getIdpLdap", args ?? new GetIdpLdapInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdpLdapArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetIdpLdapArgs()
        {
        }
        public static new GetIdpLdapArgs Empty => new GetIdpLdapArgs();
    }

    public sealed class GetIdpLdapInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetIdpLdapInvokeArgs()
        {
        }
        public static new GetIdpLdapInvokeArgs Empty => new GetIdpLdapInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdpLdapResult
    {
        /// <summary>
        /// User attribute for the avatar url
        /// </summary>
        public readonly string AvatarUrlAttribute;
        /// <summary>
        /// Base DN for LDAP connections
        /// </summary>
        public readonly string BaseDn;
        /// <summary>
        /// Bind DN for LDAP connections
        /// </summary>
        public readonly string BindDn;
        /// <summary>
        /// Bind password for LDAP connections
        /// </summary>
        public readonly string BindPassword;
        /// <summary>
        /// User attribute for the display name
        /// </summary>
        public readonly string DisplayNameAttribute;
        /// <summary>
        /// User attribute for the email
        /// </summary>
        public readonly string EmailAttribute;
        /// <summary>
        /// User attribute for the email verified state
        /// </summary>
        public readonly string EmailVerifiedAttribute;
        /// <summary>
        /// User attribute for the first name
        /// </summary>
        public readonly string FirstNameAttribute;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// User attribute for the id
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
        /// User attribute for the last name
        /// </summary>
        public readonly string LastNameAttribute;
        /// <summary>
        /// Name of the IDP
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// User attribute for the nick name
        /// </summary>
        public readonly string NickNameAttribute;
        /// <summary>
        /// User attribute for the phone
        /// </summary>
        public readonly string PhoneAttribute;
        /// <summary>
        /// User attribute for the phone verified state
        /// </summary>
        public readonly string PhoneVerifiedAttribute;
        /// <summary>
        /// User attribute for the preferred language
        /// </summary>
        public readonly string PreferredLanguageAttribute;
        /// <summary>
        /// User attribute for the preferred username
        /// </summary>
        public readonly string PreferredUsernameAttribute;
        /// <summary>
        /// User attribute for the profile
        /// </summary>
        public readonly string ProfileAttribute;
        /// <summary>
        /// Servers to try in order for establishing LDAP connections
        /// </summary>
        public readonly ImmutableArray<string> Servers;
        /// <summary>
        /// Wether to use StartTLS for LDAP connections
        /// </summary>
        public readonly bool StartTls;
        /// <summary>
        /// Timeout for LDAP connections
        /// </summary>
        public readonly string Timeout;
        /// <summary>
        /// User base for LDAP connections
        /// </summary>
        public readonly string UserBase;
        /// <summary>
        /// User filters for LDAP connections
        /// </summary>
        public readonly ImmutableArray<string> UserFilters;
        /// <summary>
        /// User object classes for LDAP connections
        /// </summary>
        public readonly ImmutableArray<string> UserObjectClasses;

        [OutputConstructor]
        private GetIdpLdapResult(
            string avatarUrlAttribute,

            string baseDn,

            string bindDn,

            string bindPassword,

            string displayNameAttribute,

            string emailAttribute,

            string emailVerifiedAttribute,

            string firstNameAttribute,

            string id,

            string idAttribute,

            bool isAutoCreation,

            bool isAutoUpdate,

            bool isCreationAllowed,

            bool isLinkingAllowed,

            string lastNameAttribute,

            string name,

            string nickNameAttribute,

            string phoneAttribute,

            string phoneVerifiedAttribute,

            string preferredLanguageAttribute,

            string preferredUsernameAttribute,

            string profileAttribute,

            ImmutableArray<string> servers,

            bool startTls,

            string timeout,

            string userBase,

            ImmutableArray<string> userFilters,

            ImmutableArray<string> userObjectClasses)
        {
            AvatarUrlAttribute = avatarUrlAttribute;
            BaseDn = baseDn;
            BindDn = bindDn;
            BindPassword = bindPassword;
            DisplayNameAttribute = displayNameAttribute;
            EmailAttribute = emailAttribute;
            EmailVerifiedAttribute = emailVerifiedAttribute;
            FirstNameAttribute = firstNameAttribute;
            Id = id;
            IdAttribute = idAttribute;
            IsAutoCreation = isAutoCreation;
            IsAutoUpdate = isAutoUpdate;
            IsCreationAllowed = isCreationAllowed;
            IsLinkingAllowed = isLinkingAllowed;
            LastNameAttribute = lastNameAttribute;
            Name = name;
            NickNameAttribute = nickNameAttribute;
            PhoneAttribute = phoneAttribute;
            PhoneVerifiedAttribute = phoneVerifiedAttribute;
            PreferredLanguageAttribute = preferredLanguageAttribute;
            PreferredUsernameAttribute = preferredUsernameAttribute;
            ProfileAttribute = profileAttribute;
            Servers = servers;
            StartTls = startTls;
            Timeout = timeout;
            UserBase = userBase;
            UserFilters = userFilters;
            UserObjectClasses = userObjectClasses;
        }
    }
}
