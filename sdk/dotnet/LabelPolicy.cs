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
    /// Resource representing the custom label policy of an organization.
    /// </summary>
    [ZitadelResourceType("zitadel:index/labelPolicy:LabelPolicy")]
    public partial class LabelPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// hex value for background color
        /// </summary>
        [Output("backgroundColor")]
        public Output<string> BackgroundColor { get; private set; } = null!;

        /// <summary>
        /// hex value for background color dark theme
        /// </summary>
        [Output("backgroundColorDark")]
        public Output<string> BackgroundColorDark { get; private set; } = null!;

        /// <summary>
        /// disable watermark
        /// </summary>
        [Output("disableWatermark")]
        public Output<bool> DisableWatermark { get; private set; } = null!;

        /// <summary>
        /// hex value for font color
        /// </summary>
        [Output("fontColor")]
        public Output<string> FontColor { get; private set; } = null!;

        /// <summary>
        /// hex value for font color dark theme
        /// </summary>
        [Output("fontColorDark")]
        public Output<string> FontColorDark { get; private set; } = null!;

        [Output("fontHash")]
        public Output<string?> FontHash { get; private set; } = null!;

        [Output("fontPath")]
        public Output<string?> FontPath { get; private set; } = null!;

        [Output("fontUrl")]
        public Output<string> FontUrl { get; private set; } = null!;

        /// <summary>
        /// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://zitadel.com/docs/apis/openidoauth/scopes#reserved-scopes
        /// </summary>
        [Output("hideLoginNameSuffix")]
        public Output<bool> HideLoginNameSuffix { get; private set; } = null!;

        [Output("iconDarkHash")]
        public Output<string?> IconDarkHash { get; private set; } = null!;

        [Output("iconDarkPath")]
        public Output<string?> IconDarkPath { get; private set; } = null!;

        [Output("iconHash")]
        public Output<string?> IconHash { get; private set; } = null!;

        [Output("iconPath")]
        public Output<string?> IconPath { get; private set; } = null!;

        [Output("iconUrl")]
        public Output<string> IconUrl { get; private set; } = null!;

        [Output("iconUrlDark")]
        public Output<string> IconUrlDark { get; private set; } = null!;

        [Output("logoDarkHash")]
        public Output<string?> LogoDarkHash { get; private set; } = null!;

        [Output("logoDarkPath")]
        public Output<string?> LogoDarkPath { get; private set; } = null!;

        [Output("logoHash")]
        public Output<string?> LogoHash { get; private set; } = null!;

        [Output("logoPath")]
        public Output<string?> LogoPath { get; private set; } = null!;

        [Output("logoUrl")]
        public Output<string> LogoUrl { get; private set; } = null!;

        [Output("logoUrlDark")]
        public Output<string> LogoUrlDark { get; private set; } = null!;

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// hex value for primary color
        /// </summary>
        [Output("primaryColor")]
        public Output<string> PrimaryColor { get; private set; } = null!;

        /// <summary>
        /// hex value for primary color dark theme
        /// </summary>
        [Output("primaryColorDark")]
        public Output<string> PrimaryColorDark { get; private set; } = null!;

        /// <summary>
        /// set the label policy active after creating/updating
        /// </summary>
        [Output("setActive")]
        public Output<bool?> SetActive { get; private set; } = null!;

        /// <summary>
        /// hex value for warn color
        /// </summary>
        [Output("warnColor")]
        public Output<string> WarnColor { get; private set; } = null!;

        /// <summary>
        /// hex value for warn color dark theme
        /// </summary>
        [Output("warnColorDark")]
        public Output<string> WarnColorDark { get; private set; } = null!;


        /// <summary>
        /// Create a LabelPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LabelPolicy(string name, LabelPolicyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/labelPolicy:LabelPolicy", name, args ?? new LabelPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LabelPolicy(string name, Input<string> id, LabelPolicyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/labelPolicy:LabelPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LabelPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LabelPolicy Get(string name, Input<string> id, LabelPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new LabelPolicy(name, id, state, options);
        }
    }

    public sealed class LabelPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// hex value for background color
        /// </summary>
        [Input("backgroundColor", required: true)]
        public Input<string> BackgroundColor { get; set; } = null!;

        /// <summary>
        /// hex value for background color dark theme
        /// </summary>
        [Input("backgroundColorDark", required: true)]
        public Input<string> BackgroundColorDark { get; set; } = null!;

        /// <summary>
        /// disable watermark
        /// </summary>
        [Input("disableWatermark", required: true)]
        public Input<bool> DisableWatermark { get; set; } = null!;

        /// <summary>
        /// hex value for font color
        /// </summary>
        [Input("fontColor", required: true)]
        public Input<string> FontColor { get; set; } = null!;

        /// <summary>
        /// hex value for font color dark theme
        /// </summary>
        [Input("fontColorDark", required: true)]
        public Input<string> FontColorDark { get; set; } = null!;

        [Input("fontHash")]
        public Input<string>? FontHash { get; set; }

        [Input("fontPath")]
        public Input<string>? FontPath { get; set; }

        /// <summary>
        /// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://zitadel.com/docs/apis/openidoauth/scopes#reserved-scopes
        /// </summary>
        [Input("hideLoginNameSuffix", required: true)]
        public Input<bool> HideLoginNameSuffix { get; set; } = null!;

        [Input("iconDarkHash")]
        public Input<string>? IconDarkHash { get; set; }

        [Input("iconDarkPath")]
        public Input<string>? IconDarkPath { get; set; }

        [Input("iconHash")]
        public Input<string>? IconHash { get; set; }

        [Input("iconPath")]
        public Input<string>? IconPath { get; set; }

        [Input("logoDarkHash")]
        public Input<string>? LogoDarkHash { get; set; }

        [Input("logoDarkPath")]
        public Input<string>? LogoDarkPath { get; set; }

        [Input("logoHash")]
        public Input<string>? LogoHash { get; set; }

        [Input("logoPath")]
        public Input<string>? LogoPath { get; set; }

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// hex value for primary color
        /// </summary>
        [Input("primaryColor", required: true)]
        public Input<string> PrimaryColor { get; set; } = null!;

        /// <summary>
        /// hex value for primary color dark theme
        /// </summary>
        [Input("primaryColorDark", required: true)]
        public Input<string> PrimaryColorDark { get; set; } = null!;

        /// <summary>
        /// set the label policy active after creating/updating
        /// </summary>
        [Input("setActive")]
        public Input<bool>? SetActive { get; set; }

        /// <summary>
        /// hex value for warn color
        /// </summary>
        [Input("warnColor", required: true)]
        public Input<string> WarnColor { get; set; } = null!;

        /// <summary>
        /// hex value for warn color dark theme
        /// </summary>
        [Input("warnColorDark", required: true)]
        public Input<string> WarnColorDark { get; set; } = null!;

        public LabelPolicyArgs()
        {
        }
        public static new LabelPolicyArgs Empty => new LabelPolicyArgs();
    }

    public sealed class LabelPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// hex value for background color
        /// </summary>
        [Input("backgroundColor")]
        public Input<string>? BackgroundColor { get; set; }

        /// <summary>
        /// hex value for background color dark theme
        /// </summary>
        [Input("backgroundColorDark")]
        public Input<string>? BackgroundColorDark { get; set; }

        /// <summary>
        /// disable watermark
        /// </summary>
        [Input("disableWatermark")]
        public Input<bool>? DisableWatermark { get; set; }

        /// <summary>
        /// hex value for font color
        /// </summary>
        [Input("fontColor")]
        public Input<string>? FontColor { get; set; }

        /// <summary>
        /// hex value for font color dark theme
        /// </summary>
        [Input("fontColorDark")]
        public Input<string>? FontColorDark { get; set; }

        [Input("fontHash")]
        public Input<string>? FontHash { get; set; }

        [Input("fontPath")]
        public Input<string>? FontPath { get; set; }

        [Input("fontUrl")]
        public Input<string>? FontUrl { get; set; }

        /// <summary>
        /// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://zitadel.com/docs/apis/openidoauth/scopes#reserved-scopes
        /// </summary>
        [Input("hideLoginNameSuffix")]
        public Input<bool>? HideLoginNameSuffix { get; set; }

        [Input("iconDarkHash")]
        public Input<string>? IconDarkHash { get; set; }

        [Input("iconDarkPath")]
        public Input<string>? IconDarkPath { get; set; }

        [Input("iconHash")]
        public Input<string>? IconHash { get; set; }

        [Input("iconPath")]
        public Input<string>? IconPath { get; set; }

        [Input("iconUrl")]
        public Input<string>? IconUrl { get; set; }

        [Input("iconUrlDark")]
        public Input<string>? IconUrlDark { get; set; }

        [Input("logoDarkHash")]
        public Input<string>? LogoDarkHash { get; set; }

        [Input("logoDarkPath")]
        public Input<string>? LogoDarkPath { get; set; }

        [Input("logoHash")]
        public Input<string>? LogoHash { get; set; }

        [Input("logoPath")]
        public Input<string>? LogoPath { get; set; }

        [Input("logoUrl")]
        public Input<string>? LogoUrl { get; set; }

        [Input("logoUrlDark")]
        public Input<string>? LogoUrlDark { get; set; }

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// hex value for primary color
        /// </summary>
        [Input("primaryColor")]
        public Input<string>? PrimaryColor { get; set; }

        /// <summary>
        /// hex value for primary color dark theme
        /// </summary>
        [Input("primaryColorDark")]
        public Input<string>? PrimaryColorDark { get; set; }

        /// <summary>
        /// set the label policy active after creating/updating
        /// </summary>
        [Input("setActive")]
        public Input<bool>? SetActive { get; set; }

        /// <summary>
        /// hex value for warn color
        /// </summary>
        [Input("warnColor")]
        public Input<string>? WarnColor { get; set; }

        /// <summary>
        /// hex value for warn color dark theme
        /// </summary>
        [Input("warnColorDark")]
        public Input<string>? WarnColorDark { get; set; }

        public LabelPolicyState()
        {
        }
        public static new LabelPolicyState Empty => new LabelPolicyState();
    }
}
