// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing a personal access token of a user
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@scoretechnologies/zitadel";
 *
 * const _default = new zitadel.PersonalAccessToken("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     userId: data.zitadel_machine_user["default"].id,
 *     expirationDate: "2519-04-01T08:45:00Z",
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<id:user_id[:org_id][:token]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/personalAccessToken:PersonalAccessToken imported '123456789012345678:123456789012345678:123456789012345678:LHt79...'
 * ```
 */
export class PersonalAccessToken extends pulumi.CustomResource {
    /**
     * Get an existing PersonalAccessToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PersonalAccessTokenState, opts?: pulumi.CustomResourceOptions): PersonalAccessToken {
        return new PersonalAccessToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/personalAccessToken:PersonalAccessToken';

    /**
     * Returns true if the given object is an instance of PersonalAccessToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PersonalAccessToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PersonalAccessToken.__pulumiType;
    }

    /**
     * Expiration date of the token in the RFC3339 format
     */
    public readonly expirationDate!: pulumi.Output<string | undefined>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * Value of the token
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * ID of the user
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a PersonalAccessToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PersonalAccessTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PersonalAccessTokenArgs | PersonalAccessTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PersonalAccessTokenState | undefined;
            resourceInputs["expirationDate"] = state ? state.expirationDate : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as PersonalAccessTokenArgs | undefined;
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["expirationDate"] = args ? args.expirationDate : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(PersonalAccessToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PersonalAccessToken resources.
 */
export interface PersonalAccessTokenState {
    /**
     * Expiration date of the token in the RFC3339 format
     */
    expirationDate?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * Value of the token
     */
    token?: pulumi.Input<string>;
    /**
     * ID of the user
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PersonalAccessToken resource.
 */
export interface PersonalAccessTokenArgs {
    /**
     * Expiration date of the token in the RFC3339 format
     */
    expirationDate?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the user
     */
    userId: pulumi.Input<string>;
}
