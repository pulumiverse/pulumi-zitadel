// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getMachineUser({
 *     orgId: data.zitadel_org["default"].id,
 *     userId: "123456789012345678",
 * });
 * ```
 */
export function getMachineUser(args: GetMachineUserArgs, opts?: pulumi.InvokeOptions): Promise<GetMachineUserResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zitadel:index/getMachineUser:getMachineUser", {
        "orgId": args.orgId,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMachineUser.
 */
export interface GetMachineUserArgs {
    /**
     * ID of the organization
     */
    orgId?: string;
    /**
     * The ID of this resource.
     */
    userId: string;
}

/**
 * A collection of values returned by getMachineUser.
 */
export interface GetMachineUserResult {
    /**
     * Access token type
     */
    readonly accessTokenType: string;
    /**
     * Description of the user
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Loginnames
     */
    readonly loginNames: string[];
    /**
     * Name of the machine user
     */
    readonly name: string;
    /**
     * ID of the organization
     */
    readonly orgId?: string;
    /**
     * Preferred login name
     */
    readonly preferredLoginName: string;
    /**
     * State of the user
     */
    readonly state: string;
    /**
     * The ID of this resource.
     */
    readonly userId: string;
    /**
     * Username
     */
    readonly userName: string;
}
/**
 * Datasource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getMachineUser({
 *     orgId: data.zitadel_org["default"].id,
 *     userId: "123456789012345678",
 * });
 * ```
 */
export function getMachineUserOutput(args: GetMachineUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMachineUserResult> {
    return pulumi.output(args).apply((a: any) => getMachineUser(a, opts))
}

/**
 * A collection of arguments for invoking getMachineUser.
 */
export interface GetMachineUserOutputArgs {
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * The ID of this resource.
     */
    userId: pulumi.Input<string>;
}
