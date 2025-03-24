// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the custom privacy policy of an organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.PrivacyPolicy("default", {
 *     orgId: defaultZitadelOrg.id,
 *     tosLink: "https://example.com/tos",
 *     privacyLink: "https://example.com/privacy",
 *     helpLink: "https://example.com/help",
 *     supportEmail: "support@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<[org_id]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/privacyPolicy:PrivacyPolicy imported '123456789012345678'
 * ```
 */
export class PrivacyPolicy extends pulumi.CustomResource {
    /**
     * Get an existing PrivacyPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivacyPolicyState, opts?: pulumi.CustomResourceOptions): PrivacyPolicy {
        return new PrivacyPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/privacyPolicy:PrivacyPolicy';

    /**
     * Returns true if the given object is an instance of PrivacyPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivacyPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivacyPolicy.__pulumiType;
    }

    public readonly helpLink!: pulumi.Output<string | undefined>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    public readonly privacyLink!: pulumi.Output<string | undefined>;
    public readonly supportEmail!: pulumi.Output<string | undefined>;
    public readonly tosLink!: pulumi.Output<string | undefined>;

    /**
     * Create a PrivacyPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PrivacyPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivacyPolicyArgs | PrivacyPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivacyPolicyState | undefined;
            resourceInputs["helpLink"] = state ? state.helpLink : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["privacyLink"] = state ? state.privacyLink : undefined;
            resourceInputs["supportEmail"] = state ? state.supportEmail : undefined;
            resourceInputs["tosLink"] = state ? state.tosLink : undefined;
        } else {
            const args = argsOrState as PrivacyPolicyArgs | undefined;
            resourceInputs["helpLink"] = args ? args.helpLink : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["privacyLink"] = args ? args.privacyLink : undefined;
            resourceInputs["supportEmail"] = args ? args.supportEmail : undefined;
            resourceInputs["tosLink"] = args ? args.tosLink : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivacyPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivacyPolicy resources.
 */
export interface PrivacyPolicyState {
    helpLink?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    privacyLink?: pulumi.Input<string>;
    supportEmail?: pulumi.Input<string>;
    tosLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivacyPolicy resource.
 */
export interface PrivacyPolicyArgs {
    helpLink?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    privacyLink?: pulumi.Input<string>;
    supportEmail?: pulumi.Input<string>;
    tosLink?: pulumi.Input<string>;
}
