// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing an Azure AD IdP of the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getOrgIdpAzureAd({
 *     orgId: data.zitadel_org["default"].id,
 *     id: "123456789012345678",
 * });
 * ```
 */
export function getOrgIdpAzureAd(args: GetOrgIdpAzureAdArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgIdpAzureAdResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("zitadel:index/getOrgIdpAzureAd:getOrgIdpAzureAd", {
        "id": args.id,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgIdpAzureAd.
 */
export interface GetOrgIdpAzureAdArgs {
    /**
     * The ID of this resource.
     */
    id: string;
    /**
     * ID of the organization
     */
    orgId?: string;
}

/**
 * A collection of values returned by getOrgIdpAzureAd.
 */
export interface GetOrgIdpAzureAdResult {
    /**
     * client id generated by the identity provider
     */
    readonly clientId: string;
    /**
     * client secret generated by the identity provider
     */
    readonly clientSecret: string;
    /**
     * automatically mark emails as verified
     */
    readonly emailVerified: boolean;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * enabled if a new account in ZITADEL are created automatically on login with an external account
     */
    readonly isAutoCreation: boolean;
    /**
     * enabled if a the ZITADEL account fields are updated automatically on each login
     */
    readonly isAutoUpdate: boolean;
    /**
     * enabled if users are able to create a new account in ZITADEL when using an external account
     */
    readonly isCreationAllowed: boolean;
    /**
     * enabled if users are able to link an existing ZITADEL user with an external account
     */
    readonly isLinkingAllowed: boolean;
    /**
     * Name of the IDP
     */
    readonly name: string;
    /**
     * ID of the organization
     */
    readonly orgId?: string;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    readonly scopes: string[];
    /**
     * the azure ad tenant id
     */
    readonly tenantId: string;
    /**
     * the azure ad tenant type
     */
    readonly tenantType: string;
}

export function getOrgIdpAzureAdOutput(args: GetOrgIdpAzureAdOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrgIdpAzureAdResult> {
    return pulumi.output(args).apply(a => getOrgIdpAzureAd(a, opts))
}

/**
 * A collection of arguments for invoking getOrgIdpAzureAd.
 */
export interface GetOrgIdpAzureAdOutputArgs {
    /**
     * The ID of this resource.
     */
    id: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
}
