// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing a GitHub IdP on the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.OrgIdpGithub("default", {
 *     orgId: defaultZitadelOrg.id,
 *     name: "GitHub",
 *     clientId: "86a165...",
 *     clientSecret: "*****afdbac18",
 *     scopes: [
 *         "openid",
 *         "profile",
 *         "email",
 *     ],
 *     isLinkingAllowed: false,
 *     isCreationAllowed: true,
 *     isAutoCreation: false,
 *     isAutoUpdate: true,
 * });
 * ```
 *
 * ## Import
 *
 * terraform The resource can be imported using the ID format `<id[:org_id][:client_secret]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/orgIdpGithub:OrgIdpGithub imported '123456789012345678:123456789012345678:1234567890123456781234567890123456787890'
 * ```
 */
export class OrgIdpGithub extends pulumi.CustomResource {
    /**
     * Get an existing OrgIdpGithub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgIdpGithubState, opts?: pulumi.CustomResourceOptions): OrgIdpGithub {
        return new OrgIdpGithub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/orgIdpGithub:OrgIdpGithub';

    /**
     * Returns true if the given object is an instance of OrgIdpGithub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgIdpGithub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgIdpGithub.__pulumiType;
    }

    /**
     * client id generated by the identity provider
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * client secret generated by the identity provider
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * enable if a new account in ZITADEL should be created automatically on login with an external account
     */
    public readonly isAutoCreation!: pulumi.Output<boolean>;
    /**
     * enable if a the ZITADEL account fields should be updated automatically on each login
     */
    public readonly isAutoUpdate!: pulumi.Output<boolean>;
    /**
     * enable if users should be able to create a new account in ZITADEL when using an external account
     */
    public readonly isCreationAllowed!: pulumi.Output<boolean>;
    /**
     * enable if users should be able to link an existing ZITADEL user with an external account
     */
    public readonly isLinkingAllowed!: pulumi.Output<boolean>;
    /**
     * Name of the IDP
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    public readonly scopes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a OrgIdpGithub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgIdpGithubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgIdpGithubArgs | OrgIdpGithubState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgIdpGithubState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["isAutoCreation"] = state ? state.isAutoCreation : undefined;
            resourceInputs["isAutoUpdate"] = state ? state.isAutoUpdate : undefined;
            resourceInputs["isCreationAllowed"] = state ? state.isCreationAllowed : undefined;
            resourceInputs["isLinkingAllowed"] = state ? state.isLinkingAllowed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
        } else {
            const args = argsOrState as OrgIdpGithubArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if ((!args || args.isAutoCreation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isAutoCreation'");
            }
            if ((!args || args.isAutoUpdate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isAutoUpdate'");
            }
            if ((!args || args.isCreationAllowed === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isCreationAllowed'");
            }
            if ((!args || args.isLinkingAllowed === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isLinkingAllowed'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["isAutoCreation"] = args ? args.isAutoCreation : undefined;
            resourceInputs["isAutoUpdate"] = args ? args.isAutoUpdate : undefined;
            resourceInputs["isCreationAllowed"] = args ? args.isCreationAllowed : undefined;
            resourceInputs["isLinkingAllowed"] = args ? args.isLinkingAllowed : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(OrgIdpGithub.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgIdpGithub resources.
 */
export interface OrgIdpGithubState {
    /**
     * client id generated by the identity provider
     */
    clientId?: pulumi.Input<string>;
    /**
     * client secret generated by the identity provider
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * enable if a new account in ZITADEL should be created automatically on login with an external account
     */
    isAutoCreation?: pulumi.Input<boolean>;
    /**
     * enable if a the ZITADEL account fields should be updated automatically on each login
     */
    isAutoUpdate?: pulumi.Input<boolean>;
    /**
     * enable if users should be able to create a new account in ZITADEL when using an external account
     */
    isCreationAllowed?: pulumi.Input<boolean>;
    /**
     * enable if users should be able to link an existing ZITADEL user with an external account
     */
    isLinkingAllowed?: pulumi.Input<boolean>;
    /**
     * Name of the IDP
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a OrgIdpGithub resource.
 */
export interface OrgIdpGithubArgs {
    /**
     * client id generated by the identity provider
     */
    clientId: pulumi.Input<string>;
    /**
     * client secret generated by the identity provider
     */
    clientSecret: pulumi.Input<string>;
    /**
     * enable if a new account in ZITADEL should be created automatically on login with an external account
     */
    isAutoCreation: pulumi.Input<boolean>;
    /**
     * enable if a the ZITADEL account fields should be updated automatically on each login
     */
    isAutoUpdate: pulumi.Input<boolean>;
    /**
     * enable if users should be able to create a new account in ZITADEL when using an external account
     */
    isCreationAllowed: pulumi.Input<boolean>;
    /**
     * enable if users should be able to link an existing ZITADEL user with an external account
     */
    isLinkingAllowed: pulumi.Input<boolean>;
    /**
     * Name of the IDP
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
}
