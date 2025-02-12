// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the project roles, which can be given as authorizations to users.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@scoretechnologies/zitadel";
 *
 * const _default = new zitadel.ProjectRole("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     projectId: data.zitadel_project["default"].id,
 *     roleKey: "super-user",
 *     displayName: "display_name2",
 *     group: "role_group",
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<project_id:role_key[:org_id]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/projectRole:ProjectRole imported '123456789012345678:my-role-key:123456789012345678'
 * ```
 */
export class ProjectRole extends pulumi.CustomResource {
    /**
     * Get an existing ProjectRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectRoleState, opts?: pulumi.CustomResourceOptions): ProjectRole {
        return new ProjectRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/projectRole:ProjectRole';

    /**
     * Returns true if the given object is an instance of ProjectRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectRole.__pulumiType;
    }

    /**
     * Name used for project role
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Group used for project role
     */
    public readonly group!: pulumi.Output<string | undefined>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * ID of the project
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Key used for project role
     */
    public readonly roleKey!: pulumi.Output<string>;

    /**
     * Create a ProjectRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectRoleArgs | ProjectRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectRoleState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["roleKey"] = state ? state.roleKey : undefined;
        } else {
            const args = argsOrState as ProjectRoleArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.roleKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleKey'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["roleKey"] = args ? args.roleKey : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectRole resources.
 */
export interface ProjectRoleState {
    /**
     * Name used for project role
     */
    displayName?: pulumi.Input<string>;
    /**
     * Group used for project role
     */
    group?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the project
     */
    projectId?: pulumi.Input<string>;
    /**
     * Key used for project role
     */
    roleKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectRole resource.
 */
export interface ProjectRoleArgs {
    /**
     * Name used for project role
     */
    displayName: pulumi.Input<string>;
    /**
     * Group used for project role
     */
    group?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the project
     */
    projectId: pulumi.Input<string>;
    /**
     * Key used for project role
     */
    roleKey: pulumi.Input<string>;
}
