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
    /// Resource representing the custom password complexity policy of an organization.
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
    ///     var @default = new Zitadel.PasswordComplexityPolicy("default", new()
    ///     {
    ///         OrgId = data.Zitadel_org.Default.Id,
    ///         MinLength = 8,
    ///         HasUppercase = true,
    ///         HasLowercase = true,
    ///         HasNumber = true,
    ///         HasSymbol = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/passwordComplexityPolicy:PasswordComplexityPolicy")]
    public partial class PasswordComplexityPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// defines if the password MUST contain a lower case letter
        /// </summary>
        [Output("hasLowercase")]
        public Output<bool> HasLowercase { get; private set; } = null!;

        /// <summary>
        /// defines if the password MUST contain a number
        /// </summary>
        [Output("hasNumber")]
        public Output<bool> HasNumber { get; private set; } = null!;

        /// <summary>
        /// defines if the password MUST contain a symbol. E.g. "$"
        /// </summary>
        [Output("hasSymbol")]
        public Output<bool> HasSymbol { get; private set; } = null!;

        /// <summary>
        /// defines if the password MUST contain an upper case letter
        /// </summary>
        [Output("hasUppercase")]
        public Output<bool> HasUppercase { get; private set; } = null!;

        /// <summary>
        /// Minimal length for the password
        /// </summary>
        [Output("minLength")]
        public Output<int> MinLength { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;


        /// <summary>
        /// Create a PasswordComplexityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PasswordComplexityPolicy(string name, PasswordComplexityPolicyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/passwordComplexityPolicy:PasswordComplexityPolicy", name, args ?? new PasswordComplexityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PasswordComplexityPolicy(string name, Input<string> id, PasswordComplexityPolicyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/passwordComplexityPolicy:PasswordComplexityPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PasswordComplexityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PasswordComplexityPolicy Get(string name, Input<string> id, PasswordComplexityPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new PasswordComplexityPolicy(name, id, state, options);
        }
    }

    public sealed class PasswordComplexityPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// defines if the password MUST contain a lower case letter
        /// </summary>
        [Input("hasLowercase", required: true)]
        public Input<bool> HasLowercase { get; set; } = null!;

        /// <summary>
        /// defines if the password MUST contain a number
        /// </summary>
        [Input("hasNumber", required: true)]
        public Input<bool> HasNumber { get; set; } = null!;

        /// <summary>
        /// defines if the password MUST contain a symbol. E.g. "$"
        /// </summary>
        [Input("hasSymbol", required: true)]
        public Input<bool> HasSymbol { get; set; } = null!;

        /// <summary>
        /// defines if the password MUST contain an upper case letter
        /// </summary>
        [Input("hasUppercase", required: true)]
        public Input<bool> HasUppercase { get; set; } = null!;

        /// <summary>
        /// Minimal length for the password
        /// </summary>
        [Input("minLength", required: true)]
        public Input<int> MinLength { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public PasswordComplexityPolicyArgs()
        {
        }
        public static new PasswordComplexityPolicyArgs Empty => new PasswordComplexityPolicyArgs();
    }

    public sealed class PasswordComplexityPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// defines if the password MUST contain a lower case letter
        /// </summary>
        [Input("hasLowercase")]
        public Input<bool>? HasLowercase { get; set; }

        /// <summary>
        /// defines if the password MUST contain a number
        /// </summary>
        [Input("hasNumber")]
        public Input<bool>? HasNumber { get; set; }

        /// <summary>
        /// defines if the password MUST contain a symbol. E.g. "$"
        /// </summary>
        [Input("hasSymbol")]
        public Input<bool>? HasSymbol { get; set; }

        /// <summary>
        /// defines if the password MUST contain an upper case letter
        /// </summary>
        [Input("hasUppercase")]
        public Input<bool>? HasUppercase { get; set; }

        /// <summary>
        /// Minimal length for the password
        /// </summary>
        [Input("minLength")]
        public Input<int>? MinLength { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public PasswordComplexityPolicyState()
        {
        }
        public static new PasswordComplexityPolicyState Empty => new PasswordComplexityPolicyState();
    }
}
