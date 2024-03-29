// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing multiple API applications belonging to a project.
 */
export function getApplicationApis(args: GetApplicationApisArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationApisResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zitadel:index/getApplicationApis:getApplicationApis", {
        "name": args.name,
        "nameMethod": args.nameMethod,
        "orgId": args.orgId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationApis.
 */
export interface GetApplicationApisArgs {
    /**
     * Name of the application
     */
    name: string;
    /**
     * Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
     */
    nameMethod?: string;
    /**
     * ID of the organization
     */
    orgId?: string;
    /**
     * ID of the project
     */
    projectId: string;
}

/**
 * A collection of values returned by getApplicationApis.
 */
export interface GetApplicationApisResult {
    /**
     * A set of all IDs.
     */
    readonly appIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the application
     */
    readonly name: string;
    /**
     * Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
     */
    readonly nameMethod?: string;
    /**
     * ID of the organization
     */
    readonly orgId?: string;
    /**
     * ID of the project
     */
    readonly projectId: string;
}
/**
 * Datasource representing multiple API applications belonging to a project.
 */
export function getApplicationApisOutput(args: GetApplicationApisOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationApisResult> {
    return pulumi.output(args).apply((a: any) => getApplicationApis(a, opts))
}

/**
 * A collection of arguments for invoking getApplicationApis.
 */
export interface GetApplicationApisOutputArgs {
    /**
     * Name of the application
     */
    name: pulumi.Input<string>;
    /**
     * Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
     */
    nameMethod?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the project
     */
    projectId: pulumi.Input<string>;
}
