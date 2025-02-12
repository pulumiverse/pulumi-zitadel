// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing a app key
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@scoretechnologies/zitadel";
 *
 * const _default = new zitadel.ApplicationKey("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     projectId: data.zitadel_project["default"].id,
 *     appId: data.zitadel_application_api["default"].id,
 *     keyType: "KEY_TYPE_JSON",
 *     expirationDate: "2519-04-01T08:45:00Z",
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<id:project_id:app_id[:org_id][:key_details]>`. You can use __SEMICOLON__ to escape :, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/applicationKey:ApplicationKey imported "123456789012345678:123456789012345678:123456789012345678:123456789012345678:$(cat ~/Downloads/123456789012345678.json | sed -e 's/:/__SEMICOLON__/g')"
 * ```
 */
export class ApplicationKey extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationKeyState, opts?: pulumi.CustomResourceOptions): ApplicationKey {
        return new ApplicationKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/applicationKey:ApplicationKey';

    /**
     * Returns true if the given object is an instance of ApplicationKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationKey.__pulumiType;
    }

    /**
     * ID of the application
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * Expiration date of the app key in the RFC3339 format
     */
    public readonly expirationDate!: pulumi.Output<string>;
    /**
     * Value of the app key
     */
    public /*out*/ readonly keyDetails!: pulumi.Output<string>;
    /**
     * Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
     */
    public readonly keyType!: pulumi.Output<string>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * ID of the project
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a ApplicationKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationKeyArgs | ApplicationKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationKeyState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["expirationDate"] = state ? state.expirationDate : undefined;
            resourceInputs["keyDetails"] = state ? state.keyDetails : undefined;
            resourceInputs["keyType"] = state ? state.keyType : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as ApplicationKeyArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.expirationDate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expirationDate'");
            }
            if ((!args || args.keyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyType'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["expirationDate"] = args ? args.expirationDate : undefined;
            resourceInputs["keyType"] = args ? args.keyType : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["keyDetails"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["keyDetails"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ApplicationKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationKey resources.
 */
export interface ApplicationKeyState {
    /**
     * ID of the application
     */
    appId?: pulumi.Input<string>;
    /**
     * Expiration date of the app key in the RFC3339 format
     */
    expirationDate?: pulumi.Input<string>;
    /**
     * Value of the app key
     */
    keyDetails?: pulumi.Input<string>;
    /**
     * Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
     */
    keyType?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the project
     */
    projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationKey resource.
 */
export interface ApplicationKeyArgs {
    /**
     * ID of the application
     */
    appId: pulumi.Input<string>;
    /**
     * Expiration date of the app key in the RFC3339 format
     */
    expirationDate: pulumi.Input<string>;
    /**
     * Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
     */
    keyType: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the project
     */
    projectId: pulumi.Input<string>;
}
