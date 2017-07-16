// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class PatchBaseline extends lumi.NamedResource implements PatchBaselineArgs {
    public readonly approvalRule?: { approveAfterDays: number, patchFilter: { key: string, values: string[] }[] }[];
    public readonly approvedPatches?: string[];
    public readonly description?: string;
    public readonly globalFilter?: { key: string, values: string[] }[];
    public readonly patchBaselineName?: string;
    public readonly rejectedPatches?: string[];

    constructor(name: string, args: PatchBaselineArgs) {
        super(name);
        this.approvalRule = args.approvalRule;
        this.approvedPatches = args.approvedPatches;
        this.description = args.description;
        this.globalFilter = args.globalFilter;
        this.patchBaselineName = args.patchBaselineName;
        this.rejectedPatches = args.rejectedPatches;
    }
}

export interface PatchBaselineArgs {
    readonly approvalRule?: { approveAfterDays: number, patchFilter: { key: string, values: string[] }[] }[];
    readonly approvedPatches?: string[];
    readonly description?: string;
    readonly globalFilter?: { key: string, values: string[] }[];
    readonly patchBaselineName?: string;
    readonly rejectedPatches?: string[];
}

