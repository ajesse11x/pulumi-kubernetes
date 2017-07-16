// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class SecurityGroup extends lumi.NamedResource implements SecurityGroupArgs {
    public readonly description?: string;
    public readonly ingress: { cidr?: string, securityGroupName: string, securityGroupOwnerId: string }[];
    public readonly securityGroupName?: string;

    constructor(name: string, args: SecurityGroupArgs) {
        super(name);
        this.description = args.description;
        if (lumirt.defaultIfComputed(args.ingress, "") === undefined) {
            throw new Error("Property argument 'ingress' is required, but was missing");
        }
        this.ingress = args.ingress;
        this.securityGroupName = args.securityGroupName;
    }
}

export interface SecurityGroupArgs {
    readonly description?: string;
    readonly ingress: { cidr?: string, securityGroupName: string, securityGroupOwnerId: string }[];
    readonly securityGroupName?: string;
}

