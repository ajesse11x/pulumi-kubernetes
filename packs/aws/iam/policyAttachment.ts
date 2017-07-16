// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

import {ARN} from "../index";
import {Group} from "./group";
import {Role} from "./role";
import {User} from "./user";

export class PolicyAttachment extends lumi.NamedResource implements PolicyAttachmentArgs {
    public readonly groups?: Group[];
    public readonly policyAttachmentName?: string;
    public readonly policyArn: ARN;
    public readonly roles?: Role[];
    public readonly users?: User[];

    constructor(name: string, args: PolicyAttachmentArgs) {
        super(name);
        this.groups = args.groups;
        this.policyAttachmentName = args.policyAttachmentName;
        if (lumirt.defaultIfComputed(args.policyArn, "") === undefined) {
            throw new Error("Property argument 'policyArn' is required, but was missing");
        }
        this.policyArn = args.policyArn;
        this.roles = args.roles;
        this.users = args.users;
    }
}

export interface PolicyAttachmentArgs {
    readonly groups?: Group[];
    readonly policyAttachmentName?: string;
    readonly policyArn: ARN;
    readonly roles?: Role[];
    readonly users?: User[];
}

