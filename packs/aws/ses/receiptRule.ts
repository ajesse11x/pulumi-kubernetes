// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class ReceiptRule extends lumi.NamedResource implements ReceiptRuleArgs {
    public readonly addHeaderAction?: { headerName: string, headerValue: string, position: number }[];
    public readonly after?: string;
    public readonly bounceAction?: { message: string, position: number, sender: string, smtpReplyCode: string, statusCode?: string, topicArn?: string }[];
    public readonly enabled: boolean;
    public readonly lambdaAction?: { functionArn: string, invocationType: string, position: number, topicArn?: string }[];
    public readonly receiptRuleName?: string;
    public readonly recipients?: string[];
    public readonly ruleSetName: string;
    public readonly s3Action?: { bucketName: string, kmsKeyArn?: string, objectKeyPrefix?: string, position: number, topicArn?: string }[];
    public readonly scanEnabled: boolean;
    public readonly snsAction?: { position: number, topicArn: string }[];
    public readonly stopAction?: { position: number, scope: string, topicArn?: string }[];
    public readonly tlsPolicy: string;
    public readonly workmailAction?: { organizationArn: string, position: number, topicArn?: string }[];

    constructor(name: string, args: ReceiptRuleArgs) {
        super(name);
        this.addHeaderAction = args.addHeaderAction;
        this.after = args.after;
        this.bounceAction = args.bounceAction;
        this.enabled = args.enabled;
        this.lambdaAction = args.lambdaAction;
        this.receiptRuleName = args.receiptRuleName;
        this.recipients = args.recipients;
        if (lumirt.defaultIfComputed(args.ruleSetName, "") === undefined) {
            throw new Error("Property argument 'ruleSetName' is required, but was missing");
        }
        this.ruleSetName = args.ruleSetName;
        this.s3Action = args.s3Action;
        this.scanEnabled = args.scanEnabled;
        this.snsAction = args.snsAction;
        this.stopAction = args.stopAction;
        this.tlsPolicy = args.tlsPolicy;
        this.workmailAction = args.workmailAction;
    }
}

export interface ReceiptRuleArgs {
    readonly addHeaderAction?: { headerName: string, headerValue: string, position: number }[];
    readonly after?: string;
    readonly bounceAction?: { message: string, position: number, sender: string, smtpReplyCode: string, statusCode?: string, topicArn?: string }[];
    readonly enabled?: boolean;
    readonly lambdaAction?: { functionArn: string, invocationType: string, position: number, topicArn?: string }[];
    readonly receiptRuleName?: string;
    readonly recipients?: string[];
    readonly ruleSetName: string;
    readonly s3Action?: { bucketName: string, kmsKeyArn?: string, objectKeyPrefix?: string, position: number, topicArn?: string }[];
    readonly scanEnabled?: boolean;
    readonly snsAction?: { position: number, topicArn: string }[];
    readonly stopAction?: { position: number, scope: string, topicArn?: string }[];
    readonly tlsPolicy?: string;
    readonly workmailAction?: { organizationArn: string, position: number, topicArn?: string }[];
}

