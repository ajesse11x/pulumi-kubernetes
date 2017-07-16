// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Domain extends lumi.NamedResource implements DomainArgs {
    public /*out*/ readonly arn: string;
    public readonly domainName: string;

    constructor(name: string, args: DomainArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.domainName, "") === undefined) {
            throw new Error("Property argument 'domainName' is required, but was missing");
        }
        this.domainName = args.domainName;
    }
}

export interface DomainArgs {
    readonly domainName: string;
}

