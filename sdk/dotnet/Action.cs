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
    /// Resource representing an action belonging to an organization.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zitadel = Vavsab.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var action = new Zitadel.Action("action", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         Script = "testscript",
    ///         Timeout = "10s",
    ///         AllowedToFail = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/action:Action")]
    public partial class Action : global::Pulumi.CustomResource
    {
        /// <summary>
        /// when true, the next action will be called even if this action fails
        /// </summary>
        [Output("allowedToFail")]
        public Output<bool> AllowedToFail { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        [Output("script")]
        public Output<string> Script { get; private set; } = null!;

        /// <summary>
        /// the state of the action
        /// </summary>
        [Output("state")]
        public Output<int> State { get; private set; } = null!;

        /// <summary>
        /// after which time the action will be terminated if not finished
        /// </summary>
        [Output("timeout")]
        public Output<string> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a Action resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Action(string name, ActionArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/action:Action", name, args ?? new ActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Action(string name, Input<string> id, ActionState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/action:Action", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Action resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Action Get(string name, Input<string> id, ActionState? state = null, CustomResourceOptions? options = null)
        {
            return new Action(name, id, state, options);
        }
    }

    public sealed class ActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// when true, the next action will be called even if this action fails
        /// </summary>
        [Input("allowedToFail", required: true)]
        public Input<bool> AllowedToFail { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        [Input("script", required: true)]
        public Input<string> Script { get; set; } = null!;

        /// <summary>
        /// after which time the action will be terminated if not finished
        /// </summary>
        [Input("timeout", required: true)]
        public Input<string> Timeout { get; set; } = null!;

        public ActionArgs()
        {
        }
        public static new ActionArgs Empty => new ActionArgs();
    }

    public sealed class ActionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// when true, the next action will be called even if this action fails
        /// </summary>
        [Input("allowedToFail")]
        public Input<bool>? AllowedToFail { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// the state of the action
        /// </summary>
        [Input("state")]
        public Input<int>? State { get; set; }

        /// <summary>
        /// after which time the action will be terminated if not finished
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public ActionState()
        {
        }
        public static new ActionState Empty => new ActionState();
    }
}
