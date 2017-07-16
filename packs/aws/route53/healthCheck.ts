// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class HealthCheck extends lumi.NamedResource implements HealthCheckArgs {
    public readonly childHealthThreshold?: number;
    public readonly childHealthchecks?: string[];
    public readonly cloudwatchAlarmName?: string;
    public readonly cloudwatchAlarmRegion?: string;
    public readonly enableSni: boolean;
    public readonly failureThreshold?: number;
    public readonly fqdn?: string;
    public readonly insufficientDataHealthStatus?: string;
    public readonly invertHealthcheck?: boolean;
    public readonly ipAddress?: string;
    public readonly measureLatency?: boolean;
    public readonly port?: number;
    public readonly referenceName?: string;
    public readonly requestInterval?: number;
    public readonly resourcePath?: string;
    public readonly searchString?: string;
    public readonly tags?: {[key: string]: any};
    public readonly type: string;

    constructor(name: string, args: HealthCheckArgs) {
        super(name);
        this.childHealthThreshold = args.childHealthThreshold;
        this.childHealthchecks = args.childHealthchecks;
        this.cloudwatchAlarmName = args.cloudwatchAlarmName;
        this.cloudwatchAlarmRegion = args.cloudwatchAlarmRegion;
        this.enableSni = args.enableSni;
        this.failureThreshold = args.failureThreshold;
        this.fqdn = args.fqdn;
        this.insufficientDataHealthStatus = args.insufficientDataHealthStatus;
        this.invertHealthcheck = args.invertHealthcheck;
        this.ipAddress = args.ipAddress;
        this.measureLatency = args.measureLatency;
        this.port = args.port;
        this.referenceName = args.referenceName;
        this.requestInterval = args.requestInterval;
        this.resourcePath = args.resourcePath;
        this.searchString = args.searchString;
        this.tags = args.tags;
        if (lumirt.defaultIfComputed(args.type, "") === undefined) {
            throw new Error("Property argument 'type' is required, but was missing");
        }
        this.type = args.type;
    }
}

export interface HealthCheckArgs {
    readonly childHealthThreshold?: number;
    readonly childHealthchecks?: string[];
    readonly cloudwatchAlarmName?: string;
    readonly cloudwatchAlarmRegion?: string;
    readonly enableSni?: boolean;
    readonly failureThreshold?: number;
    readonly fqdn?: string;
    readonly insufficientDataHealthStatus?: string;
    readonly invertHealthcheck?: boolean;
    readonly ipAddress?: string;
    readonly measureLatency?: boolean;
    readonly port?: number;
    readonly referenceName?: string;
    readonly requestInterval?: number;
    readonly resourcePath?: string;
    readonly searchString?: string;
    readonly tags?: {[key: string]: any};
    readonly type: string;
}

