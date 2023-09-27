// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the custom notification policy of an organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.NotificationPolicy("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     passwordChange: false,
 * });
 * ```
 *
 * ## Import
 *
 * terraform # The resource can be imported using the ID format `<[org_id]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/notificationPolicy:NotificationPolicy imported '123456789012345678'
 * ```
 */
export class NotificationPolicy extends pulumi.CustomResource {
    /**
     * Get an existing NotificationPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationPolicyState, opts?: pulumi.CustomResourceOptions): NotificationPolicy {
        return new NotificationPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/notificationPolicy:NotificationPolicy';

    /**
     * Returns true if the given object is an instance of NotificationPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationPolicy.__pulumiType;
    }

    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * Send notification if a user changes his password
     */
    public readonly passwordChange!: pulumi.Output<boolean>;

    /**
     * Create a NotificationPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationPolicyArgs | NotificationPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationPolicyState | undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["passwordChange"] = state ? state.passwordChange : undefined;
        } else {
            const args = argsOrState as NotificationPolicyArgs | undefined;
            if ((!args || args.passwordChange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'passwordChange'");
            }
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["passwordChange"] = args ? args.passwordChange : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NotificationPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationPolicy resources.
 */
export interface NotificationPolicyState {
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * Send notification if a user changes his password
     */
    passwordChange?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a NotificationPolicy resource.
 */
export interface NotificationPolicyArgs {
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * Send notification if a user changes his password
     */
    passwordChange: pulumi.Input<boolean>;
}
