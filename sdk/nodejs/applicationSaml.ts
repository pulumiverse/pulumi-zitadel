// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing a SAML application belonging to a project, with all configuration possibilities.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@scoretechnologies/zitadel";
 *
 * const _default = new zitadel.ApplicationSaml("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     projectId: data.zitadel_project["default"].id,
 *     metadataXml: `<?xml version="1.0"?>
 * <md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata"
 *                      validUntil="2024-01-26T17:48:38Z"
 *                      cacheDuration="PT604800S"
 *                      entityID="http://example.com/saml/metadata">
 *     <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
 *         <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
 *         <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
 *                                      Location="http://example.com/saml/cas"
 *                                      index="1" />
 *         
 *     </md:SPSSODescriptor>
 * </md:EntityDescriptor>`,
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<id:project_id[:org_id]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/applicationSaml:ApplicationSaml imported '123456789012345678:123456789012345678:123456789012345678'
 * ```
 */
export class ApplicationSaml extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationSaml resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationSamlState, opts?: pulumi.CustomResourceOptions): ApplicationSaml {
        return new ApplicationSaml(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/applicationSaml:ApplicationSaml';

    /**
     * Returns true if the given object is an instance of ApplicationSaml.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationSaml {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationSaml.__pulumiType;
    }

    /**
     * Metadata as XML file
     */
    public readonly metadataXml!: pulumi.Output<string>;
    /**
     * Name of the application
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * ID of the project
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a ApplicationSaml resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationSamlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationSamlArgs | ApplicationSamlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationSamlState | undefined;
            resourceInputs["metadataXml"] = state ? state.metadataXml : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as ApplicationSamlArgs | undefined;
            if ((!args || args.metadataXml === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadataXml'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["metadataXml"] = args?.metadataXml ? pulumi.secret(args.metadataXml) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["metadataXml"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ApplicationSaml.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationSaml resources.
 */
export interface ApplicationSamlState {
    /**
     * Metadata as XML file
     */
    metadataXml?: pulumi.Input<string>;
    /**
     * Name of the application
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the project
     */
    projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationSaml resource.
 */
export interface ApplicationSamlArgs {
    /**
     * Metadata as XML file
     */
    metadataXml: pulumi.Input<string>;
    /**
     * Name of the application
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the project
     */
    projectId: pulumi.Input<string>;
}
