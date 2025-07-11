// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the exoscale package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'exoscale';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * Exoscale API key
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * Exoscale API secret
     */
    public readonly secret!: pulumi.Output<string | undefined>;
    public readonly sosEndpoint!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["delay"] = pulumi.output(args ? args.delay : undefined).apply(JSON.stringify);
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["key"] = (args?.key ? pulumi.secret(args.key) : undefined) ?? utilities.getEnv("EXOSCALE_API_KEY");
            resourceInputs["secret"] = (args?.secret ? pulumi.secret(args.secret) : undefined) ?? utilities.getEnv("EXOSCALE_API_SECRET");
            resourceInputs["sosEndpoint"] = args ? args.sosEndpoint : undefined;
            resourceInputs["timeout"] = pulumi.output(args ? args.timeout : undefined).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key", "secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }

    /**
     * This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
     */
    terraformConfig(): pulumi.Output<Provider.TerraformConfigResult> {
        return pulumi.runtime.call("pulumi:providers:exoscale/terraformConfig", {
            "__self__": this,
        }, this);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * @deprecated Does nothing
     */
    delay?: pulumi.Input<number>;
    environment?: pulumi.Input<string>;
    /**
     * Exoscale API key
     */
    key?: pulumi.Input<string>;
    /**
     * Exoscale API secret
     */
    secret?: pulumi.Input<string>;
    sosEndpoint?: pulumi.Input<string>;
    /**
     * Timeout in seconds for waiting on compute resources to become available (by default: 3600)
     */
    timeout?: pulumi.Input<number>;
}

export namespace Provider {
    /**
     * The results of the Provider.terraformConfig method.
     */
    export interface TerraformConfigResult {
        readonly result: {[key: string]: any};
    }

}
