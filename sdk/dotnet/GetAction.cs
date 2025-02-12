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
    public static class GetAction
    {
        /// <summary>
        /// Datasource representing an action belonging to an organization.
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
        ///     var @default = Zitadel.GetAction.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Default.Id,
        ///         ActionId = "123456789012345678",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["action"] = @default,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetActionResult> InvokeAsync(GetActionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActionResult>("zitadel:index/getAction:getAction", args ?? new GetActionArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing an action belonging to an organization.
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
        ///     var @default = Zitadel.GetAction.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Default.Id,
        ///         ActionId = "123456789012345678",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["action"] = @default,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetActionResult> Invoke(GetActionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionResult>("zitadel:index/getAction:getAction", args ?? new GetActionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("actionId", required: true)]
        public string ActionId { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        public GetActionArgs()
        {
        }
        public static new GetActionArgs Empty => new GetActionArgs();
    }

    public sealed class GetActionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("actionId", required: true)]
        public Input<string> ActionId { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public GetActionInvokeArgs()
        {
        }
        public static new GetActionInvokeArgs Empty => new GetActionInvokeArgs();
    }


    [OutputType]
    public sealed class GetActionResult
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string ActionId;
        /// <summary>
        /// when true, the next action will be called even if this action fails
        /// </summary>
        public readonly bool AllowedToFail;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string? OrgId;
        public readonly string Script;
        /// <summary>
        /// the state of the action
        /// </summary>
        public readonly int State;
        /// <summary>
        /// after which time the action will be terminated if not finished
        /// </summary>
        public readonly string Timeout;

        [OutputConstructor]
        private GetActionResult(
            string actionId,

            bool allowedToFail,

            string id,

            string name,

            string? orgId,

            string script,

            int state,

            string timeout)
        {
            ActionId = actionId;
            AllowedToFail = allowedToFail;
            Id = id;
            Name = name;
            OrgId = orgId;
            Script = script;
            State = state;
            Timeout = timeout;
        }
    }
}
