// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

import {Deployment} from "./deployment";
import {RestApi} from "./restApi";

export class Stage extends lumi.NamedResource implements StageArgs {
    public readonly cacheClusterEnabled?: boolean;
    public readonly cacheClusterSize?: string;
    public readonly clientCertificateId?: string;
    public readonly deployment: Deployment;
    public readonly description?: string;
    public readonly documentationVersion?: string;
    public readonly restApi: RestApi;
    public readonly stageName: string;
    public readonly variables?: {[key: string]: any};

    constructor(name: string, args: StageArgs) {
        super(name);
        this.cacheClusterEnabled = args.cacheClusterEnabled;
        this.cacheClusterSize = args.cacheClusterSize;
        this.clientCertificateId = args.clientCertificateId;
        if (lumirt.defaultIfComputed(args.deployment, "") === undefined) {
            throw new Error("Property argument 'deployment' is required, but was missing");
        }
        this.deployment = args.deployment;
        this.description = args.description;
        this.documentationVersion = args.documentationVersion;
        if (lumirt.defaultIfComputed(args.restApi, "") === undefined) {
            throw new Error("Property argument 'restApi' is required, but was missing");
        }
        this.restApi = args.restApi;
        if (lumirt.defaultIfComputed(args.stageName, "") === undefined) {
            throw new Error("Property argument 'stageName' is required, but was missing");
        }
        this.stageName = args.stageName;
        this.variables = args.variables;
    }
}

export interface StageArgs {
    readonly cacheClusterEnabled?: boolean;
    readonly cacheClusterSize?: string;
    readonly clientCertificateId?: string;
    readonly deployment: Deployment;
    readonly description?: string;
    readonly documentationVersion?: string;
    readonly restApi: RestApi;
    readonly stageName: string;
    readonly variables?: {[key: string]: any};
}

