// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the custom lockout policy of an organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.LockoutPolicy("default", {
 *     orgId: defaultZitadelOrg.id,
 *     maxPasswordAttempts: 5,
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<[org_id]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/lockoutPolicy:LockoutPolicy imported '123456789012345678'
 * ```
 */
export class LockoutPolicy extends pulumi.CustomResource {
    /**
     * Get an existing LockoutPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LockoutPolicyState, opts?: pulumi.CustomResourceOptions): LockoutPolicy {
        return new LockoutPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/lockoutPolicy:LockoutPolicy';

    /**
     * Returns true if the given object is an instance of LockoutPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LockoutPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LockoutPolicy.__pulumiType;
    }

    /**
     * Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
     */
    public readonly maxPasswordAttempts!: pulumi.Output<number>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;

    /**
     * Create a LockoutPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LockoutPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LockoutPolicyArgs | LockoutPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LockoutPolicyState | undefined;
            resourceInputs["maxPasswordAttempts"] = state ? state.maxPasswordAttempts : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
        } else {
            const args = argsOrState as LockoutPolicyArgs | undefined;
            if ((!args || args.maxPasswordAttempts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxPasswordAttempts'");
            }
            resourceInputs["maxPasswordAttempts"] = args ? args.maxPasswordAttempts : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LockoutPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LockoutPolicy resources.
 */
export interface LockoutPolicyState {
    /**
     * Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
     */
    maxPasswordAttempts?: pulumi.Input<number>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LockoutPolicy resource.
 */
export interface LockoutPolicyArgs {
    /**
     * Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
     */
    maxPasswordAttempts: pulumi.Input<number>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
}
