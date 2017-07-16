// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class SrvRecord extends lumi.NamedResource implements SrvRecordArgs {
    public readonly srvRecordName?: string;
    public readonly record: { port: number, priority: number, target: string, weight: number }[];
    public readonly resourceGroupName: string;
    public readonly tags: {[key: string]: any};
    public readonly ttl: number;
    public readonly zoneName: string;

    constructor(name: string, args: SrvRecordArgs) {
        super(name);
        this.srvRecordName = args.srvRecordName;
        if (lumirt.defaultIfComputed(args.record, "") === undefined) {
            throw new Error("Property argument 'record' is required, but was missing");
        }
        this.record = args.record;
        if (lumirt.defaultIfComputed(args.resourceGroupName, "") === undefined) {
            throw new Error("Property argument 'resourceGroupName' is required, but was missing");
        }
        this.resourceGroupName = args.resourceGroupName;
        this.tags = args.tags;
        if (lumirt.defaultIfComputed(args.ttl, "") === undefined) {
            throw new Error("Property argument 'ttl' is required, but was missing");
        }
        this.ttl = args.ttl;
        if (lumirt.defaultIfComputed(args.zoneName, "") === undefined) {
            throw new Error("Property argument 'zoneName' is required, but was missing");
        }
        this.zoneName = args.zoneName;
    }
}

export interface SrvRecordArgs {
    readonly srvRecordName?: string;
    readonly record: { port: number, priority: number, target: string, weight: number }[];
    readonly resourceGroupName: string;
    readonly tags?: {[key: string]: any};
    readonly ttl: number;
    readonly zoneName: string;
}

