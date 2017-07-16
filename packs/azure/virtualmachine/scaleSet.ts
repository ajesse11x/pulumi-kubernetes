// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class ScaleSet extends lumi.NamedResource implements ScaleSetArgs {
    public readonly extension?: { autoUpgradeMinorVersion?: boolean, name: string, protectedSettings?: string, publisher: string, settings?: string, type: string, typeHandlerVersion: string }[];
    public readonly location: string;
    public readonly scaleSetName?: string;
    public readonly networkProfile: { ipConfiguration: { loadBalancerBackendAddressPoolIds?: string[], loadBalancerInboundNatRulesIds: string[], name: string, subnetId: string }[], name: string, primary: boolean }[];
    public readonly osProfile: { adminPassword: string, adminUsername: string, computerNamePrefix: string, customData?: string }[];
    public readonly osProfileLinuxConfig: { disablePasswordAuthentication?: boolean, sshKeys?: { keyData?: string, path: string }[] }[];
    public readonly osProfileSecrets?: { sourceVaultId: string, vaultCertificates?: { certificateStore?: string, certificateUrl: string }[] }[];
    public readonly osProfileWindowsConfig?: { additionalUnattendConfig?: { component: string, content: string, pass: string, settingName: string }[], enableAutomaticUpgrades?: boolean, provisionVmAgent?: boolean, winrm?: { certificateUrl?: string, protocol: string }[] }[];
    public readonly overprovision?: boolean;
    public readonly resourceGroupName: string;
    public readonly singlePlacementGroup?: boolean;
    public readonly sku: { capacity: number, name: string, tier: string }[];
    public readonly storageProfileDataDisk?: { caching: string, createOption: string, diskSizeGb: number, lun: number, managedDiskType: string }[];
    public readonly storageProfileImageReference: { offer: string, publisher: string, sku: string, version: string }[];
    public readonly storageProfileOsDisk: { caching: string, createOption: string, image?: string, managedDiskType: string, name?: string, osType?: string, vhdContainers?: string[] }[];
    public readonly tags: {[key: string]: any};
    public readonly upgradePolicyMode: string;

    constructor(name: string, args: ScaleSetArgs) {
        super(name);
        this.extension = args.extension;
        if (lumirt.defaultIfComputed(args.location, "") === undefined) {
            throw new Error("Property argument 'location' is required, but was missing");
        }
        this.location = args.location;
        this.scaleSetName = args.scaleSetName;
        if (lumirt.defaultIfComputed(args.networkProfile, "") === undefined) {
            throw new Error("Property argument 'networkProfile' is required, but was missing");
        }
        this.networkProfile = args.networkProfile;
        if (lumirt.defaultIfComputed(args.osProfile, "") === undefined) {
            throw new Error("Property argument 'osProfile' is required, but was missing");
        }
        this.osProfile = args.osProfile;
        this.osProfileLinuxConfig = args.osProfileLinuxConfig;
        this.osProfileSecrets = args.osProfileSecrets;
        this.osProfileWindowsConfig = args.osProfileWindowsConfig;
        this.overprovision = args.overprovision;
        if (lumirt.defaultIfComputed(args.resourceGroupName, "") === undefined) {
            throw new Error("Property argument 'resourceGroupName' is required, but was missing");
        }
        this.resourceGroupName = args.resourceGroupName;
        this.singlePlacementGroup = args.singlePlacementGroup;
        if (lumirt.defaultIfComputed(args.sku, "") === undefined) {
            throw new Error("Property argument 'sku' is required, but was missing");
        }
        this.sku = args.sku;
        this.storageProfileDataDisk = args.storageProfileDataDisk;
        this.storageProfileImageReference = args.storageProfileImageReference;
        if (lumirt.defaultIfComputed(args.storageProfileOsDisk, "") === undefined) {
            throw new Error("Property argument 'storageProfileOsDisk' is required, but was missing");
        }
        this.storageProfileOsDisk = args.storageProfileOsDisk;
        this.tags = args.tags;
        if (lumirt.defaultIfComputed(args.upgradePolicyMode, "") === undefined) {
            throw new Error("Property argument 'upgradePolicyMode' is required, but was missing");
        }
        this.upgradePolicyMode = args.upgradePolicyMode;
    }
}

export interface ScaleSetArgs {
    readonly extension?: { autoUpgradeMinorVersion?: boolean, name: string, protectedSettings?: string, publisher: string, settings?: string, type: string, typeHandlerVersion: string }[];
    readonly location: string;
    readonly scaleSetName?: string;
    readonly networkProfile: { ipConfiguration: { loadBalancerBackendAddressPoolIds?: string[], loadBalancerInboundNatRulesIds: string[], name: string, subnetId: string }[], name: string, primary: boolean }[];
    readonly osProfile: { adminPassword: string, adminUsername: string, computerNamePrefix: string, customData?: string }[];
    readonly osProfileLinuxConfig?: { disablePasswordAuthentication?: boolean, sshKeys?: { keyData?: string, path: string }[] }[];
    readonly osProfileSecrets?: { sourceVaultId: string, vaultCertificates?: { certificateStore?: string, certificateUrl: string }[] }[];
    readonly osProfileWindowsConfig?: { additionalUnattendConfig?: { component: string, content: string, pass: string, settingName: string }[], enableAutomaticUpgrades?: boolean, provisionVmAgent?: boolean, winrm?: { certificateUrl?: string, protocol: string }[] }[];
    readonly overprovision?: boolean;
    readonly resourceGroupName: string;
    readonly singlePlacementGroup?: boolean;
    readonly sku: { capacity: number, name: string, tier: string }[];
    readonly storageProfileDataDisk?: { caching: string, createOption: string, diskSizeGb: number, lun: number, managedDiskType: string }[];
    readonly storageProfileImageReference?: { offer: string, publisher: string, sku: string, version: string }[];
    readonly storageProfileOsDisk: { caching: string, createOption: string, image?: string, managedDiskType: string, name?: string, osType?: string, vhdContainers?: string[] }[];
    readonly tags?: {[key: string]: any};
    readonly upgradePolicyMode: string;
}

