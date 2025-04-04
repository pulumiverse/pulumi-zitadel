// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the SMTP configuration of an instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.SmtpConfig("default", {
 *     senderAddress: "sender@example.com",
 *     senderName: "no-reply",
 *     tls: true,
 *     host: "localhost:25",
 *     user: "user",
 *     password: "secret_password",
 *     replyToAddress: "replyto@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<[password]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/smtpConfig:SmtpConfig imported 'p4ssw0rd'
 * ```
 */
export class SmtpConfig extends pulumi.CustomResource {
    /**
     * Get an existing SmtpConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SmtpConfigState, opts?: pulumi.CustomResourceOptions): SmtpConfig {
        return new SmtpConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/smtpConfig:SmtpConfig';

    /**
     * Returns true if the given object is an instance of SmtpConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SmtpConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SmtpConfig.__pulumiType;
    }

    /**
     * Host and port address to your SMTP server.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * Password used to communicate with your SMTP server.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Address to reply to.
     */
    public readonly replyToAddress!: pulumi.Output<string | undefined>;
    /**
     * Address used to send emails.
     */
    public readonly senderAddress!: pulumi.Output<string>;
    /**
     * Sender name used to send emails.
     */
    public readonly senderName!: pulumi.Output<string>;
    /**
     * TLS used to communicate with your SMTP server.
     */
    public readonly tls!: pulumi.Output<boolean | undefined>;
    /**
     * User used to communicate with your SMTP server.
     */
    public readonly user!: pulumi.Output<string | undefined>;

    /**
     * Create a SmtpConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SmtpConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SmtpConfigArgs | SmtpConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SmtpConfigState | undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["replyToAddress"] = state ? state.replyToAddress : undefined;
            resourceInputs["senderAddress"] = state ? state.senderAddress : undefined;
            resourceInputs["senderName"] = state ? state.senderName : undefined;
            resourceInputs["tls"] = state ? state.tls : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as SmtpConfigArgs | undefined;
            if ((!args || args.host === undefined) && !opts.urn) {
                throw new Error("Missing required property 'host'");
            }
            if ((!args || args.senderAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'senderAddress'");
            }
            if ((!args || args.senderName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'senderName'");
            }
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["replyToAddress"] = args ? args.replyToAddress : undefined;
            resourceInputs["senderAddress"] = args ? args.senderAddress : undefined;
            resourceInputs["senderName"] = args ? args.senderName : undefined;
            resourceInputs["tls"] = args ? args.tls : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SmtpConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SmtpConfig resources.
 */
export interface SmtpConfigState {
    /**
     * Host and port address to your SMTP server.
     */
    host?: pulumi.Input<string>;
    /**
     * Password used to communicate with your SMTP server.
     */
    password?: pulumi.Input<string>;
    /**
     * Address to reply to.
     */
    replyToAddress?: pulumi.Input<string>;
    /**
     * Address used to send emails.
     */
    senderAddress?: pulumi.Input<string>;
    /**
     * Sender name used to send emails.
     */
    senderName?: pulumi.Input<string>;
    /**
     * TLS used to communicate with your SMTP server.
     */
    tls?: pulumi.Input<boolean>;
    /**
     * User used to communicate with your SMTP server.
     */
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SmtpConfig resource.
 */
export interface SmtpConfigArgs {
    /**
     * Host and port address to your SMTP server.
     */
    host: pulumi.Input<string>;
    /**
     * Password used to communicate with your SMTP server.
     */
    password?: pulumi.Input<string>;
    /**
     * Address to reply to.
     */
    replyToAddress?: pulumi.Input<string>;
    /**
     * Address used to send emails.
     */
    senderAddress: pulumi.Input<string>;
    /**
     * Sender name used to send emails.
     */
    senderName: pulumi.Input<string>;
    /**
     * TLS used to communicate with your SMTP server.
     */
    tls?: pulumi.Input<boolean>;
    /**
     * User used to communicate with your SMTP server.
     */
    user?: pulumi.Input<string>;
}
