// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getProject({
 *     orgId: data.zitadel_org["default"].id,
 *     projectId: "123456789012345678",
 * });
 * ```
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zitadel:index/getProject:getProject", {
        "orgId": args.orgId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * ID of the organization
     */
    orgId?: string;
    /**
     * The ID of this resource.
     */
    projectId: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * ZITADEL checks if the org of the user has permission to this project
     */
    readonly hasProjectCheck: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the project
     */
    readonly name: string;
    /**
     * ID of the organization
     */
    readonly orgId?: string;
    /**
     * Defines from where the private labeling should be triggered
     */
    readonly privateLabelingSetting: string;
    /**
     * The ID of this resource.
     */
    readonly projectId: string;
    /**
     * describes if roles of user should be added in token
     */
    readonly projectRoleAssertion: boolean;
    /**
     * ZITADEL checks if the user has at least one on this project
     */
    readonly projectRoleCheck: boolean;
    /**
     * State of the project
     */
    readonly state: string;
}
/**
 * Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getProject({
 *     orgId: data.zitadel_org["default"].id,
 *     projectId: "123456789012345678",
 * });
 * ```
 */
export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply((a: any) => getProject(a, opts))
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectOutputArgs {
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * The ID of this resource.
     */
    projectId: pulumi.Input<string>;
}
