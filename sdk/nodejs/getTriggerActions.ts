// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing triggers, when actions get started
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getTriggerActions({
 *     orgId: defaultZitadelOrg.id,
 *     flowType: "FLOW_TYPE_EXTERNAL_AUTHENTICATION",
 *     triggerType: "TRIGGER_TYPE_POST_AUTHENTICATION",
 * });
 * export const triggerActions = _default;
 * ```
 */
export function getTriggerActions(args: GetTriggerActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTriggerActionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zitadel:index/getTriggerActions:getTriggerActions", {
        "flowType": args.flowType,
        "orgId": args.orgId,
        "triggerType": args.triggerType,
    }, opts);
}

/**
 * A collection of arguments for invoking getTriggerActions.
 */
export interface GetTriggerActionsArgs {
    /**
     * Type of the flow to which the action triggers belong
     */
    flowType: string;
    /**
     * ID of the organization
     */
    orgId?: string;
    /**
     * Trigger type on when the actions get triggered
     */
    triggerType: string;
}

/**
 * A collection of values returned by getTriggerActions.
 */
export interface GetTriggerActionsResult {
    /**
     * IDs of the triggered actions
     */
    readonly actionIds: string[];
    /**
     * Type of the flow to which the action triggers belong
     */
    readonly flowType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * ID of the organization
     */
    readonly orgId?: string;
    /**
     * Trigger type on when the actions get triggered
     */
    readonly triggerType: string;
}
/**
 * Resource representing triggers, when actions get started
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getTriggerActions({
 *     orgId: defaultZitadelOrg.id,
 *     flowType: "FLOW_TYPE_EXTERNAL_AUTHENTICATION",
 *     triggerType: "TRIGGER_TYPE_POST_AUTHENTICATION",
 * });
 * export const triggerActions = _default;
 * ```
 */
export function getTriggerActionsOutput(args: GetTriggerActionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTriggerActionsResult> {
    return pulumi.output(args).apply((a: any) => getTriggerActions(a, opts))
}

/**
 * A collection of arguments for invoking getTriggerActions.
 */
export interface GetTriggerActionsOutputArgs {
    /**
     * Type of the flow to which the action triggers belong
     */
    flowType: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * Trigger type on when the actions get triggered
     */
    triggerType: pulumi.Input<string>;
}
