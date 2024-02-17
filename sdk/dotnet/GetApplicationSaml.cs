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
    public static class GetApplicationSaml
    {
        /// <summary>
        /// Datasource representing a SAML application belonging to a project, with all configuration possibilities.
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
        ///     var @default = Zitadel.GetApplicationSaml.Invoke(new()
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
        public static Task<GetApplicationSamlResult> InvokeAsync(GetApplicationSamlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationSamlResult>("zitadel:index/getApplicationSaml:getApplicationSaml", args ?? new GetApplicationSamlArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing a SAML application belonging to a project, with all configuration possibilities.
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
        ///     var @default = Zitadel.GetApplicationSaml.Invoke(new()
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
        public static Output<GetApplicationSamlResult> Invoke(GetApplicationSamlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationSamlResult>("zitadel:index/getApplicationSaml:getApplicationSaml", args ?? new GetApplicationSamlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationSamlArgs : global::Pulumi.InvokeArgs
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

        public GetApplicationSamlArgs()
        {
        }
        public static new GetApplicationSamlArgs Empty => new GetApplicationSamlArgs();
    }

    public sealed class GetApplicationSamlInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetApplicationSamlInvokeArgs()
        {
        }
        public static new GetApplicationSamlInvokeArgs Empty => new GetApplicationSamlInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationSamlResult
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Metadata as XML file
        /// </summary>
        public readonly string MetadataXml;
        /// <summary>
        /// Name of the application
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// ID of the project
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private GetApplicationSamlResult(
            string appId,

            string id,

            string metadataXml,

            string name,

            string? orgId,

            string projectId)
        {
            AppId = appId;
            Id = id;
            MetadataXml = metadataXml;
            Name = name;
            OrgId = orgId;
            ProjectId = projectId;
        }
    }
}
