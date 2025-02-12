// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the default notification policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@scoretechnologies/zitadel";
 *
 * const _default = new zitadel.DefaultNotificationPolicy("default", {passwordChange: false});
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy imported ''
 * ```
 */
export class DefaultNotificationPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DefaultNotificationPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultNotificationPolicyState, opts?: pulumi.CustomResourceOptions): DefaultNotificationPolicy {
        return new DefaultNotificationPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy';

    /**
     * Returns true if the given object is an instance of DefaultNotificationPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultNotificationPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultNotificationPolicy.__pulumiType;
    }

    /**
     * Send notification if a user changes his password
     */
    public readonly passwordChange!: pulumi.Output<boolean>;

    /**
     * Create a DefaultNotificationPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultNotificationPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultNotificationPolicyArgs | DefaultNotificationPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultNotificationPolicyState | undefined;
            resourceInputs["passwordChange"] = state ? state.passwordChange : undefined;
        } else {
            const args = argsOrState as DefaultNotificationPolicyArgs | undefined;
            if ((!args || args.passwordChange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'passwordChange'");
            }
            resourceInputs["passwordChange"] = args ? args.passwordChange : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefaultNotificationPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultNotificationPolicy resources.
 */
export interface DefaultNotificationPolicyState {
    /**
     * Send notification if a user changes his password
     */
    passwordChange?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DefaultNotificationPolicy resource.
 */
export interface DefaultNotificationPolicyArgs {
    /**
     * Send notification if a user changes his password
     */
    passwordChange: pulumi.Input<boolean>;
}
