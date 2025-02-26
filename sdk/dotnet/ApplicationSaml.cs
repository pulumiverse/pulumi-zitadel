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
    /// Resource representing a SAML application belonging to a project, with all configuration possibilities.
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
    ///     var @default = new Zitadel.ApplicationSaml("default", new()
    ///     {
    ///         OrgId = defaultZitadelOrg.Id,
    ///         ProjectId = defaultZitadelProject.Id,
    ///         Name = "applicationapi",
    ///         MetadataXml = @"&lt;?xml version=""1.0""?&gt;
    /// &lt;md:EntityDescriptor xmlns:md=""urn:oasis:names:tc:SAML:2.0:metadata""
    ///                      validUntil=""2024-01-26T17:48:38Z""
    ///                      cacheDuration=""PT604800S""
    ///                      entityID=""http://example.com/saml/metadata""&gt;
    ///     &lt;md:SPSSODescriptor AuthnRequestsSigned=""false"" WantAssertionsSigned=""false"" protocolSupportEnumeration=""urn:oasis:names:tc:SAML:2.0:protocol""&gt;
    ///         &lt;md:NameIDFormat&gt;urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified&lt;/md:NameIDFormat&gt;
    ///         &lt;md:AssertionConsumerService Binding=""urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST""
    ///                                      Location=""http://example.com/saml/cas""
    ///                                      index=""1"" /&gt;
    ///         
    ///     &lt;/md:SPSSODescriptor&gt;
    /// &lt;/md:EntityDescriptor&gt;",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform The resource can be imported using the ID format `&lt;id:project_id[:org_id]&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/applicationSaml:ApplicationSaml imported '123456789012345678:123456789012345678:123456789012345678'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/applicationSaml:ApplicationSaml")]
    public partial class ApplicationSaml : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Metadata as XML file
        /// </summary>
        [Output("metadataXml")]
        public Output<string> MetadataXml { get; private set; } = null!;

        /// <summary>
        /// Name of the application
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationSaml resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationSaml(string name, ApplicationSamlArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/applicationSaml:ApplicationSaml", name, args ?? new ApplicationSamlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationSaml(string name, Input<string> id, ApplicationSamlState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/applicationSaml:ApplicationSaml", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "metadataXml",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationSaml resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationSaml Get(string name, Input<string> id, ApplicationSamlState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationSaml(name, id, state, options);
        }
    }

    public sealed class ApplicationSamlArgs : global::Pulumi.ResourceArgs
    {
        [Input("metadataXml", required: true)]
        private Input<string>? _metadataXml;

        /// <summary>
        /// Metadata as XML file
        /// </summary>
        public Input<string>? MetadataXml
        {
            get => _metadataXml;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _metadataXml = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of the application
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public ApplicationSamlArgs()
        {
        }
        public static new ApplicationSamlArgs Empty => new ApplicationSamlArgs();
    }

    public sealed class ApplicationSamlState : global::Pulumi.ResourceArgs
    {
        [Input("metadataXml")]
        private Input<string>? _metadataXml;

        /// <summary>
        /// Metadata as XML file
        /// </summary>
        public Input<string>? MetadataXml
        {
            get => _metadataXml;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _metadataXml = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of the application
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public ApplicationSamlState()
        {
        }
        public static new ApplicationSamlState Empty => new ApplicationSamlState();
    }
}
