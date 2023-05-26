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
    /// Resource representing a machine key
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zitadel = Pulumiverse.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var machineKey = new Zitadel.MachineKey("machineKey", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         UserId = zitadel_machine_user.Machine_user.Id,
    ///         KeyType = "KEY_TYPE_JSON",
    ///         ExpirationDate = "2519-04-01T08:45:00Z",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/machineKey:MachineKey")]
    public partial class MachineKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Expiration date of the machine key in the RFC3339 format
        /// </summary>
        [Output("expirationDate")]
        public Output<string> ExpirationDate { get; private set; } = null!;

        /// <summary>
        /// Value of the machine key
        /// </summary>
        [Output("keyDetails")]
        public Output<string> KeyDetails { get; private set; } = null!;

        /// <summary>
        /// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        /// </summary>
        [Output("keyType")]
        public Output<string> KeyType { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of the user
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a MachineKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MachineKey(string name, MachineKeyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/machineKey:MachineKey", name, args ?? new MachineKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MachineKey(string name, Input<string> id, MachineKeyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/machineKey:MachineKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MachineKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MachineKey Get(string name, Input<string> id, MachineKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new MachineKey(name, id, state, options);
        }
    }

    public sealed class MachineKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expiration date of the machine key in the RFC3339 format
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        /// </summary>
        [Input("keyType", required: true)]
        public Input<string> KeyType { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public MachineKeyArgs()
        {
        }
        public static new MachineKeyArgs Empty => new MachineKeyArgs();
    }

    public sealed class MachineKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expiration date of the machine key in the RFC3339 format
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// Value of the machine key
        /// </summary>
        [Input("keyDetails")]
        public Input<string>? KeyDetails { get; set; }

        /// <summary>
        /// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public MachineKeyState()
        {
        }
        public static new MachineKeyState Empty => new MachineKeyState();
    }
}
