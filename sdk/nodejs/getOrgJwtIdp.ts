// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing a generic JWT IdP on the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const orgJwtIdpOrgJwtIdp = zitadel.getOrgJwtIdp({
 *     orgId: data.zitadel_org.org.id,
 *     idpId: "177073612581240835",
 * });
 * export const orgJwtIdp = orgJwtIdpOrgJwtIdp;
 * ```
 */
export function getOrgJwtIdp(args: GetOrgJwtIdpArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgJwtIdpResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("zitadel:index/getOrgJwtIdp:getOrgJwtIdp", {
        "idpId": args.idpId,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgJwtIdp.
 */
export interface GetOrgJwtIdpArgs {
    /**
     * The ID of this resource.
     */
    idpId: string;
    /**
     * ID of the organization
     */
    orgId: string;
}

/**
 * A collection of values returned by getOrgJwtIdp.
 */
export interface GetOrgJwtIdpResult {
    /**
     * auto register for users from this idp
     */
    readonly autoRegister: boolean;
    /**
     * the name of the header where the JWT is sent in, default is authorization
     */
    readonly headerName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID of this resource.
     */
    readonly idpId: string;
    /**
     * the issuer of the jwt (for validation)
     */
    readonly issuer: string;
    /**
     * the endpoint where the jwt can be extracted
     */
    readonly jwtEndpoint: string;
    /**
     * the endpoint to the key (JWK) which are used to sign the JWT with
     */
    readonly keysEndpoint: string;
    /**
     * Name of the IDP
     */
    readonly name: string;
    /**
     * ID of the organization
     */
    readonly orgId: string;
    /**
     * Some identity providers specify the styling of the button to their login
     */
    readonly stylingType: string;
}

export function getOrgJwtIdpOutput(args: GetOrgJwtIdpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrgJwtIdpResult> {
    return pulumi.output(args).apply(a => getOrgJwtIdp(a, opts))
}

/**
 * A collection of arguments for invoking getOrgJwtIdp.
 */
export interface GetOrgJwtIdpOutputArgs {
    /**
     * The ID of this resource.
     */
    idpId: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId: pulumi.Input<string>;
}
