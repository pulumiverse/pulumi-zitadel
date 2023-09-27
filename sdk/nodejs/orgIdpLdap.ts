// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing an LDAP IdP on the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.OrgIdpLdap("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     servers: [
 *         "ldaps://my.primary.server:389",
 *         "ldaps://my.secondary.server:389",
 *     ],
 *     startTls: false,
 *     baseDn: "dc=example,dc=com",
 *     bindDn: "cn=admin,dc=example,dc=com",
 *     bindPassword: "Password1!",
 *     userBase: "dn",
 *     userObjectClasses: ["inetOrgPerson"],
 *     userFilters: [
 *         "uid",
 *         "email",
 *     ],
 *     timeout: "10s",
 *     idAttribute: "uid",
 *     firstNameAttribute: "firstname",
 *     lastNameAttribute: "lastname",
 *     isLinkingAllowed: false,
 *     isCreationAllowed: true,
 *     isAutoCreation: false,
 *     isAutoUpdate: true,
 * });
 * ```
 *
 * ## Import
 *
 * terraform # The resource can be imported using the ID format `<id[:org_id][:bind_password]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/orgIdpLdap:OrgIdpLdap imported '123456789012345678:123456789012345678:b1nd_p4ssw0rd'
 * ```
 */
export class OrgIdpLdap extends pulumi.CustomResource {
    /**
     * Get an existing OrgIdpLdap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgIdpLdapState, opts?: pulumi.CustomResourceOptions): OrgIdpLdap {
        return new OrgIdpLdap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/orgIdpLdap:OrgIdpLdap';

    /**
     * Returns true if the given object is an instance of OrgIdpLdap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgIdpLdap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgIdpLdap.__pulumiType;
    }

    /**
     * User attribute for the avatar url
     */
    public readonly avatarUrlAttribute!: pulumi.Output<string | undefined>;
    /**
     * Base DN for LDAP connections
     */
    public readonly baseDn!: pulumi.Output<string>;
    /**
     * Bind DN for LDAP connections
     */
    public readonly bindDn!: pulumi.Output<string>;
    /**
     * Bind password for LDAP connections
     */
    public readonly bindPassword!: pulumi.Output<string>;
    /**
     * User attribute for the display name
     */
    public readonly displayNameAttribute!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the email
     */
    public readonly emailAttribute!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the email verified state
     */
    public readonly emailVerifiedAttribute!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the first name
     */
    public readonly firstNameAttribute!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the id
     */
    public readonly idAttribute!: pulumi.Output<string | undefined>;
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
     * User attribute for the last name
     */
    public readonly lastNameAttribute!: pulumi.Output<string | undefined>;
    /**
     * Name of the IDP
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * User attribute for the nick name
     */
    public readonly nickNameAttribute!: pulumi.Output<string | undefined>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the phone
     */
    public readonly phoneAttribute!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the phone verified state
     */
    public readonly phoneVerifiedAttribute!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the preferred language
     */
    public readonly preferredLanguageAttribute!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the preferred username
     */
    public readonly preferredUsernameAttribute!: pulumi.Output<string | undefined>;
    /**
     * User attribute for the profile
     */
    public readonly profileAttribute!: pulumi.Output<string | undefined>;
    /**
     * Servers to try in order for establishing LDAP connections
     */
    public readonly servers!: pulumi.Output<string[]>;
    /**
     * Wether to use StartTLS for LDAP connections
     */
    public readonly startTls!: pulumi.Output<boolean>;
    /**
     * Timeout for LDAP connections
     */
    public readonly timeout!: pulumi.Output<string>;
    /**
     * User base for LDAP connections
     */
    public readonly userBase!: pulumi.Output<string>;
    /**
     * User filters for LDAP connections
     */
    public readonly userFilters!: pulumi.Output<string[]>;
    /**
     * User object classes for LDAP connections
     */
    public readonly userObjectClasses!: pulumi.Output<string[]>;

    /**
     * Create a OrgIdpLdap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgIdpLdapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgIdpLdapArgs | OrgIdpLdapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgIdpLdapState | undefined;
            resourceInputs["avatarUrlAttribute"] = state ? state.avatarUrlAttribute : undefined;
            resourceInputs["baseDn"] = state ? state.baseDn : undefined;
            resourceInputs["bindDn"] = state ? state.bindDn : undefined;
            resourceInputs["bindPassword"] = state ? state.bindPassword : undefined;
            resourceInputs["displayNameAttribute"] = state ? state.displayNameAttribute : undefined;
            resourceInputs["emailAttribute"] = state ? state.emailAttribute : undefined;
            resourceInputs["emailVerifiedAttribute"] = state ? state.emailVerifiedAttribute : undefined;
            resourceInputs["firstNameAttribute"] = state ? state.firstNameAttribute : undefined;
            resourceInputs["idAttribute"] = state ? state.idAttribute : undefined;
            resourceInputs["isAutoCreation"] = state ? state.isAutoCreation : undefined;
            resourceInputs["isAutoUpdate"] = state ? state.isAutoUpdate : undefined;
            resourceInputs["isCreationAllowed"] = state ? state.isCreationAllowed : undefined;
            resourceInputs["isLinkingAllowed"] = state ? state.isLinkingAllowed : undefined;
            resourceInputs["lastNameAttribute"] = state ? state.lastNameAttribute : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nickNameAttribute"] = state ? state.nickNameAttribute : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["phoneAttribute"] = state ? state.phoneAttribute : undefined;
            resourceInputs["phoneVerifiedAttribute"] = state ? state.phoneVerifiedAttribute : undefined;
            resourceInputs["preferredLanguageAttribute"] = state ? state.preferredLanguageAttribute : undefined;
            resourceInputs["preferredUsernameAttribute"] = state ? state.preferredUsernameAttribute : undefined;
            resourceInputs["profileAttribute"] = state ? state.profileAttribute : undefined;
            resourceInputs["servers"] = state ? state.servers : undefined;
            resourceInputs["startTls"] = state ? state.startTls : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["userBase"] = state ? state.userBase : undefined;
            resourceInputs["userFilters"] = state ? state.userFilters : undefined;
            resourceInputs["userObjectClasses"] = state ? state.userObjectClasses : undefined;
        } else {
            const args = argsOrState as OrgIdpLdapArgs | undefined;
            if ((!args || args.baseDn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseDn'");
            }
            if ((!args || args.bindDn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindDn'");
            }
            if ((!args || args.bindPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindPassword'");
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
            if ((!args || args.servers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servers'");
            }
            if ((!args || args.startTls === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTls'");
            }
            if ((!args || args.timeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeout'");
            }
            if ((!args || args.userBase === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userBase'");
            }
            if ((!args || args.userFilters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userFilters'");
            }
            if ((!args || args.userObjectClasses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userObjectClasses'");
            }
            resourceInputs["avatarUrlAttribute"] = args ? args.avatarUrlAttribute : undefined;
            resourceInputs["baseDn"] = args ? args.baseDn : undefined;
            resourceInputs["bindDn"] = args ? args.bindDn : undefined;
            resourceInputs["bindPassword"] = args ? args.bindPassword : undefined;
            resourceInputs["displayNameAttribute"] = args ? args.displayNameAttribute : undefined;
            resourceInputs["emailAttribute"] = args ? args.emailAttribute : undefined;
            resourceInputs["emailVerifiedAttribute"] = args ? args.emailVerifiedAttribute : undefined;
            resourceInputs["firstNameAttribute"] = args ? args.firstNameAttribute : undefined;
            resourceInputs["idAttribute"] = args ? args.idAttribute : undefined;
            resourceInputs["isAutoCreation"] = args ? args.isAutoCreation : undefined;
            resourceInputs["isAutoUpdate"] = args ? args.isAutoUpdate : undefined;
            resourceInputs["isCreationAllowed"] = args ? args.isCreationAllowed : undefined;
            resourceInputs["isLinkingAllowed"] = args ? args.isLinkingAllowed : undefined;
            resourceInputs["lastNameAttribute"] = args ? args.lastNameAttribute : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nickNameAttribute"] = args ? args.nickNameAttribute : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["phoneAttribute"] = args ? args.phoneAttribute : undefined;
            resourceInputs["phoneVerifiedAttribute"] = args ? args.phoneVerifiedAttribute : undefined;
            resourceInputs["preferredLanguageAttribute"] = args ? args.preferredLanguageAttribute : undefined;
            resourceInputs["preferredUsernameAttribute"] = args ? args.preferredUsernameAttribute : undefined;
            resourceInputs["profileAttribute"] = args ? args.profileAttribute : undefined;
            resourceInputs["servers"] = args ? args.servers : undefined;
            resourceInputs["startTls"] = args ? args.startTls : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["userBase"] = args ? args.userBase : undefined;
            resourceInputs["userFilters"] = args ? args.userFilters : undefined;
            resourceInputs["userObjectClasses"] = args ? args.userObjectClasses : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgIdpLdap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgIdpLdap resources.
 */
export interface OrgIdpLdapState {
    /**
     * User attribute for the avatar url
     */
    avatarUrlAttribute?: pulumi.Input<string>;
    /**
     * Base DN for LDAP connections
     */
    baseDn?: pulumi.Input<string>;
    /**
     * Bind DN for LDAP connections
     */
    bindDn?: pulumi.Input<string>;
    /**
     * Bind password for LDAP connections
     */
    bindPassword?: pulumi.Input<string>;
    /**
     * User attribute for the display name
     */
    displayNameAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the email
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the email verified state
     */
    emailVerifiedAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the first name
     */
    firstNameAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the id
     */
    idAttribute?: pulumi.Input<string>;
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
     * User attribute for the last name
     */
    lastNameAttribute?: pulumi.Input<string>;
    /**
     * Name of the IDP
     */
    name?: pulumi.Input<string>;
    /**
     * User attribute for the nick name
     */
    nickNameAttribute?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * User attribute for the phone
     */
    phoneAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the phone verified state
     */
    phoneVerifiedAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the preferred language
     */
    preferredLanguageAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the preferred username
     */
    preferredUsernameAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the profile
     */
    profileAttribute?: pulumi.Input<string>;
    /**
     * Servers to try in order for establishing LDAP connections
     */
    servers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Wether to use StartTLS for LDAP connections
     */
    startTls?: pulumi.Input<boolean>;
    /**
     * Timeout for LDAP connections
     */
    timeout?: pulumi.Input<string>;
    /**
     * User base for LDAP connections
     */
    userBase?: pulumi.Input<string>;
    /**
     * User filters for LDAP connections
     */
    userFilters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User object classes for LDAP connections
     */
    userObjectClasses?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a OrgIdpLdap resource.
 */
export interface OrgIdpLdapArgs {
    /**
     * User attribute for the avatar url
     */
    avatarUrlAttribute?: pulumi.Input<string>;
    /**
     * Base DN for LDAP connections
     */
    baseDn: pulumi.Input<string>;
    /**
     * Bind DN for LDAP connections
     */
    bindDn: pulumi.Input<string>;
    /**
     * Bind password for LDAP connections
     */
    bindPassword: pulumi.Input<string>;
    /**
     * User attribute for the display name
     */
    displayNameAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the email
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the email verified state
     */
    emailVerifiedAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the first name
     */
    firstNameAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the id
     */
    idAttribute?: pulumi.Input<string>;
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
     * User attribute for the last name
     */
    lastNameAttribute?: pulumi.Input<string>;
    /**
     * Name of the IDP
     */
    name?: pulumi.Input<string>;
    /**
     * User attribute for the nick name
     */
    nickNameAttribute?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * User attribute for the phone
     */
    phoneAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the phone verified state
     */
    phoneVerifiedAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the preferred language
     */
    preferredLanguageAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the preferred username
     */
    preferredUsernameAttribute?: pulumi.Input<string>;
    /**
     * User attribute for the profile
     */
    profileAttribute?: pulumi.Input<string>;
    /**
     * Servers to try in order for establishing LDAP connections
     */
    servers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Wether to use StartTLS for LDAP connections
     */
    startTls: pulumi.Input<boolean>;
    /**
     * Timeout for LDAP connections
     */
    timeout: pulumi.Input<string>;
    /**
     * User base for LDAP connections
     */
    userBase: pulumi.Input<string>;
    /**
     * User filters for LDAP connections
     */
    userFilters: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User object classes for LDAP connections
     */
    userObjectClasses: pulumi.Input<pulumi.Input<string>[]>;
}
