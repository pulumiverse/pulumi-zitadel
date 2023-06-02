// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the project, which can then be granted to different organizations or users directly, containing different applications.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const project = new zitadel.Project("project", {
 *     orgId: zitadel_org.org.id,
 *     projectRoleAssertion: true,
 *     projectRoleCheck: true,
 *     hasProjectCheck: true,
 *     privateLabelingSetting: "PRIVATE_LABELING_SETTING_ENFORCE_PROJECT_RESOURCE_OWNER_POLICY",
 * });
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * ZITADEL checks if the org of the user has permission to this project
     */
    public readonly hasProjectCheck!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the project
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Organization in which the project is located
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * Defines from where the private labeling should be triggered, supported values: PRIVATE*LABELING*SETTING*UNSPECIFIED, PRIVATE*LABELING*SETTING*ENFORCE*PROJECT*RESOURCE*OWNER*POLICY, PRIVATE*LABELING*SETTING*ALLOW*LOGIN*USER*RESOURCE*OWNER*POLICY
     */
    public readonly privateLabelingSetting!: pulumi.Output<string | undefined>;
    /**
     * describes if roles of user should be added in token
     */
    public readonly projectRoleAssertion!: pulumi.Output<boolean | undefined>;
    /**
     * ZITADEL checks if the user has at least one on this project
     */
    public readonly projectRoleCheck!: pulumi.Output<boolean | undefined>;
    /**
     * State of the project
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["hasProjectCheck"] = state ? state.hasProjectCheck : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["privateLabelingSetting"] = state ? state.privateLabelingSetting : undefined;
            resourceInputs["projectRoleAssertion"] = state ? state.projectRoleAssertion : undefined;
            resourceInputs["projectRoleCheck"] = state ? state.projectRoleCheck : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            resourceInputs["hasProjectCheck"] = args ? args.hasProjectCheck : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["privateLabelingSetting"] = args ? args.privateLabelingSetting : undefined;
            resourceInputs["projectRoleAssertion"] = args ? args.projectRoleAssertion : undefined;
            resourceInputs["projectRoleCheck"] = args ? args.projectRoleCheck : undefined;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * ZITADEL checks if the org of the user has permission to this project
     */
    hasProjectCheck?: pulumi.Input<boolean>;
    /**
     * Name of the project
     */
    name?: pulumi.Input<string>;
    /**
     * Organization in which the project is located
     */
    orgId?: pulumi.Input<string>;
    /**
     * Defines from where the private labeling should be triggered, supported values: PRIVATE*LABELING*SETTING*UNSPECIFIED, PRIVATE*LABELING*SETTING*ENFORCE*PROJECT*RESOURCE*OWNER*POLICY, PRIVATE*LABELING*SETTING*ALLOW*LOGIN*USER*RESOURCE*OWNER*POLICY
     */
    privateLabelingSetting?: pulumi.Input<string>;
    /**
     * describes if roles of user should be added in token
     */
    projectRoleAssertion?: pulumi.Input<boolean>;
    /**
     * ZITADEL checks if the user has at least one on this project
     */
    projectRoleCheck?: pulumi.Input<boolean>;
    /**
     * State of the project
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * ZITADEL checks if the org of the user has permission to this project
     */
    hasProjectCheck?: pulumi.Input<boolean>;
    /**
     * Name of the project
     */
    name?: pulumi.Input<string>;
    /**
     * Organization in which the project is located
     */
    orgId: pulumi.Input<string>;
    /**
     * Defines from where the private labeling should be triggered, supported values: PRIVATE*LABELING*SETTING*UNSPECIFIED, PRIVATE*LABELING*SETTING*ENFORCE*PROJECT*RESOURCE*OWNER*POLICY, PRIVATE*LABELING*SETTING*ALLOW*LOGIN*USER*RESOURCE*OWNER*POLICY
     */
    privateLabelingSetting?: pulumi.Input<string>;
    /**
     * describes if roles of user should be added in token
     */
    projectRoleAssertion?: pulumi.Input<boolean>;
    /**
     * ZITADEL checks if the user has at least one on this project
     */
    projectRoleCheck?: pulumi.Input<boolean>;
}