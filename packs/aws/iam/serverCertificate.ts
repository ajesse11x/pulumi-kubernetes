// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class ServerCertificate extends lumi.NamedResource implements ServerCertificateArgs {
    public readonly arn: string;
    public readonly certificateBody: string;
    public readonly certificateChain?: string;
    public readonly serverCertificateName: string;
    public readonly namePrefix?: string;
    public readonly path?: string;
    public readonly privateKey: string;

    constructor(name: string, args: ServerCertificateArgs) {
        super(name);
        this.arn = args.arn;
        if (lumirt.defaultIfComputed(args.certificateBody, "") === undefined) {
            throw new Error("Property argument 'certificateBody' is required, but was missing");
        }
        this.certificateBody = args.certificateBody;
        this.certificateChain = args.certificateChain;
        this.serverCertificateName = args.serverCertificateName;
        this.namePrefix = args.namePrefix;
        this.path = args.path;
        if (lumirt.defaultIfComputed(args.privateKey, "") === undefined) {
            throw new Error("Property argument 'privateKey' is required, but was missing");
        }
        this.privateKey = args.privateKey;
    }
}

export interface ServerCertificateArgs {
    readonly arn?: string;
    readonly certificateBody: string;
    readonly certificateChain?: string;
    readonly serverCertificateName?: string;
    readonly namePrefix?: string;
    readonly path?: string;
    readonly privateKey: string;
}

