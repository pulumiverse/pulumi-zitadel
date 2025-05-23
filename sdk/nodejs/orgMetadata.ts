// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.OrgMetadata("default", {
 *     orgId: defaultZitadelOrg.id,
 *     key: "a_key",
 *     value: "a_value",
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<key[:org_id]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/orgMetadata:OrgMetadata imported 'a_key:123456789012345678'
 * ```
 */
export class OrgMetadata extends pulumi.CustomResource {
    /**
     * Get an existing OrgMetadata resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgMetadataState, opts?: pulumi.CustomResourceOptions): OrgMetadata {
        return new OrgMetadata(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/orgMetadata:OrgMetadata';

    /**
     * Returns true if the given object is an instance of OrgMetadata.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgMetadata {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgMetadata.__pulumiType;
    }

    /**
     * The key of a metadata entry
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * The string representation of a metadata entry value. For binary data, use the base64encode function.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a OrgMetadata resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgMetadataArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgMetadataArgs | OrgMetadataState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgMetadataState | undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as OrgMetadataArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgMetadata.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgMetadata resources.
 */
export interface OrgMetadataState {
    /**
     * The key of a metadata entry
     */
    key?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * The string representation of a metadata entry value. For binary data, use the base64encode function.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgMetadata resource.
 */
export interface OrgMetadataArgs {
    /**
     * The key of a metadata entry
     */
    key: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * The string representation of a metadata entry value. For binary data, use the base64encode function.
     */
    value: pulumi.Input<string>;
}
