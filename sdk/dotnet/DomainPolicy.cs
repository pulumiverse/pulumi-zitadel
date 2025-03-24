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
    /// Resource representing the custom domain policy of an organization.
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
    ///     var @default = new Zitadel.DomainPolicy("default", new()
    ///     {
    ///         OrgId = defaultZitadelOrg.Id,
    ///         UserLoginMustBeDomain = false,
    ///         ValidateOrgDomains = true,
    ///         SmtpSenderAddressMatchesInstanceDomain = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// bash The resource can be imported using the ID format `&lt;[org_id]&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/domainPolicy:DomainPolicy imported '123456789012345678'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/domainPolicy:DomainPolicy")]
    public partial class DomainPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        [Output("smtpSenderAddressMatchesInstanceDomain")]
        public Output<bool> SmtpSenderAddressMatchesInstanceDomain { get; private set; } = null!;

        /// <summary>
        /// User login must be domain
        /// </summary>
        [Output("userLoginMustBeDomain")]
        public Output<bool> UserLoginMustBeDomain { get; private set; } = null!;

        /// <summary>
        /// Validate organization domains
        /// </summary>
        [Output("validateOrgDomains")]
        public Output<bool> ValidateOrgDomains { get; private set; } = null!;


        /// <summary>
        /// Create a DomainPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainPolicy(string name, DomainPolicyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/domainPolicy:DomainPolicy", name, args ?? new DomainPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainPolicy(string name, Input<string> id, DomainPolicyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/domainPolicy:DomainPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainPolicy Get(string name, Input<string> id, DomainPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainPolicy(name, id, state, options);
        }
    }

    public sealed class DomainPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("smtpSenderAddressMatchesInstanceDomain", required: true)]
        public Input<bool> SmtpSenderAddressMatchesInstanceDomain { get; set; } = null!;

        /// <summary>
        /// User login must be domain
        /// </summary>
        [Input("userLoginMustBeDomain", required: true)]
        public Input<bool> UserLoginMustBeDomain { get; set; } = null!;

        /// <summary>
        /// Validate organization domains
        /// </summary>
        [Input("validateOrgDomains", required: true)]
        public Input<bool> ValidateOrgDomains { get; set; } = null!;

        public DomainPolicyArgs()
        {
        }
        public static new DomainPolicyArgs Empty => new DomainPolicyArgs();
    }

    public sealed class DomainPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("smtpSenderAddressMatchesInstanceDomain")]
        public Input<bool>? SmtpSenderAddressMatchesInstanceDomain { get; set; }

        /// <summary>
        /// User login must be domain
        /// </summary>
        [Input("userLoginMustBeDomain")]
        public Input<bool>? UserLoginMustBeDomain { get; set; }

        /// <summary>
        /// Validate organization domains
        /// </summary>
        [Input("validateOrgDomains")]
        public Input<bool>? ValidateOrgDomains { get; set; }

        public DomainPolicyState()
        {
        }
        public static new DomainPolicyState Empty => new DomainPolicyState();
    }
}
