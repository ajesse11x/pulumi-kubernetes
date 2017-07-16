// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Schedule extends lumi.NamedResource implements ScheduleArgs {
    public /*out*/ readonly arn: string;
    public readonly autoscalingGroupName: string;
    public readonly desiredCapacity: number;
    public readonly endTime: string;
    public readonly maxSize: number;
    public readonly minSize: number;
    public readonly recurrence: string;
    public readonly scheduledActionName: string;
    public readonly startTime: string;

    constructor(name: string, args: ScheduleArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.autoscalingGroupName, "") === undefined) {
            throw new Error("Property argument 'autoscalingGroupName' is required, but was missing");
        }
        this.autoscalingGroupName = args.autoscalingGroupName;
        this.desiredCapacity = args.desiredCapacity;
        this.endTime = args.endTime;
        this.maxSize = args.maxSize;
        this.minSize = args.minSize;
        this.recurrence = args.recurrence;
        if (lumirt.defaultIfComputed(args.scheduledActionName, "") === undefined) {
            throw new Error("Property argument 'scheduledActionName' is required, but was missing");
        }
        this.scheduledActionName = args.scheduledActionName;
        this.startTime = args.startTime;
    }
}

export interface ScheduleArgs {
    readonly autoscalingGroupName: string;
    readonly desiredCapacity?: number;
    readonly endTime?: string;
    readonly maxSize?: number;
    readonly minSize?: number;
    readonly recurrence?: string;
    readonly scheduledActionName: string;
    readonly startTime?: string;
}

