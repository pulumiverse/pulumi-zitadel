// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.Org("default", {name: "terraform-test"});
 * ```
 *
 * ## Import
 *
 * terraform The resource can be imported using the ID format `<id>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/org:Org imported '123456789012345678'
 * ```
 */
export class Org extends pulumi.CustomResource {
    /**
     * Get an existing Org resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgState, opts?: pulumi.CustomResourceOptions): Org {
        return new Org(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/org:Org';

    /**
     * Returns true if the given object is an instance of Org.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Org {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Org.__pulumiType;
    }

    /**
     * True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
     */
    public readonly isDefault!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the org
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Primary domain of the org
     */
    public /*out*/ readonly primaryDomain!: pulumi.Output<string>;
    /**
     * State of the org
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Org resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OrgArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgArgs | OrgState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgState | undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["primaryDomain"] = state ? state.primaryDomain : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as OrgArgs | undefined;
            resourceInputs["isDefault"] = args ? args.isDefault : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["primaryDomain"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Org.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Org resources.
 */
export interface OrgState {
    /**
     * True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Name of the org
     */
    name?: pulumi.Input<string>;
    /**
     * Primary domain of the org
     */
    primaryDomain?: pulumi.Input<string>;
    /**
     * State of the org
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Org resource.
 */
export interface OrgArgs {
    /**
     * True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Name of the org
     */
    name?: pulumi.Input<string>;
}
