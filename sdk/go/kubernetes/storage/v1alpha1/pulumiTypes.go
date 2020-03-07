// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
)

// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
type VolumeAttachmentSource struct {
	// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
	InlineVolumeSpec *corev1.PersistentVolumeSpec `pulumi:"inlineVolumeSpec"`
	// Name of the persistent volume to attach.
	PersistentVolumeName *string `pulumi:"persistentVolumeName"`
}

type VolumeAttachmentSourceInput interface {
	pulumi.Input

	ToVolumeAttachmentSourceOutput() VolumeAttachmentSourceOutput
	ToVolumeAttachmentSourceOutputWithContext(context.Context) VolumeAttachmentSourceOutput
}

// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
type VolumeAttachmentSourceArgs struct {
	// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
	InlineVolumeSpec corev1.PersistentVolumeSpecPtrInput `pulumi:"inlineVolumeSpec"`
	// Name of the persistent volume to attach.
	PersistentVolumeName pulumi.StringPtrInput `pulumi:"persistentVolumeName"`
}

func (VolumeAttachmentSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentSource)(nil)).Elem()
}

func (i VolumeAttachmentSourceArgs) ToVolumeAttachmentSourceOutput() VolumeAttachmentSourceOutput {
	return i.ToVolumeAttachmentSourceOutputWithContext(context.Background())
}

func (i VolumeAttachmentSourceArgs) ToVolumeAttachmentSourceOutputWithContext(ctx context.Context) VolumeAttachmentSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentSourceOutput)
}

// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
type VolumeAttachmentSourceOutput struct { *pulumi.OutputState }

func (VolumeAttachmentSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentSource)(nil)).Elem()
}

func (o VolumeAttachmentSourceOutput) ToVolumeAttachmentSourceOutput() VolumeAttachmentSourceOutput {
	return o
}

func (o VolumeAttachmentSourceOutput) ToVolumeAttachmentSourceOutputWithContext(ctx context.Context) VolumeAttachmentSourceOutput {
	return o
}

// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
func (o VolumeAttachmentSourceOutput) InlineVolumeSpec() corev1.PersistentVolumeSpecPtrOutput {
	return o.ApplyT(func (v VolumeAttachmentSource) *corev1.PersistentVolumeSpec { return v.InlineVolumeSpec }).(corev1.PersistentVolumeSpecPtrOutput)
}

// Name of the persistent volume to attach.
func (o VolumeAttachmentSourceOutput) PersistentVolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeAttachmentSource) *string { return v.PersistentVolumeName }).(pulumi.StringPtrOutput)
}

// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpec struct {
	// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
	Attacher *string `pulumi:"attacher"`
	// The node that the volume should be attached to.
	NodeName *string `pulumi:"nodeName"`
	// Source represents the volume that should be attached.
	Source *VolumeAttachmentSource `pulumi:"source"`
}

type VolumeAttachmentSpecInput interface {
	pulumi.Input

	ToVolumeAttachmentSpecOutput() VolumeAttachmentSpecOutput
	ToVolumeAttachmentSpecOutputWithContext(context.Context) VolumeAttachmentSpecOutput
}

// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpecArgs struct {
	// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
	Attacher pulumi.StringPtrInput `pulumi:"attacher"`
	// The node that the volume should be attached to.
	NodeName pulumi.StringPtrInput `pulumi:"nodeName"`
	// Source represents the volume that should be attached.
	Source VolumeAttachmentSourcePtrInput `pulumi:"source"`
}

func (VolumeAttachmentSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentSpec)(nil)).Elem()
}

func (i VolumeAttachmentSpecArgs) ToVolumeAttachmentSpecOutput() VolumeAttachmentSpecOutput {
	return i.ToVolumeAttachmentSpecOutputWithContext(context.Background())
}

func (i VolumeAttachmentSpecArgs) ToVolumeAttachmentSpecOutputWithContext(ctx context.Context) VolumeAttachmentSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentSpecOutput)
}

// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpecOutput struct { *pulumi.OutputState }

func (VolumeAttachmentSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentSpec)(nil)).Elem()
}

func (o VolumeAttachmentSpecOutput) ToVolumeAttachmentSpecOutput() VolumeAttachmentSpecOutput {
	return o
}

func (o VolumeAttachmentSpecOutput) ToVolumeAttachmentSpecOutputWithContext(ctx context.Context) VolumeAttachmentSpecOutput {
	return o
}

// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
func (o VolumeAttachmentSpecOutput) Attacher() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeAttachmentSpec) *string { return v.Attacher }).(pulumi.StringPtrOutput)
}

// The node that the volume should be attached to.
func (o VolumeAttachmentSpecOutput) NodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeAttachmentSpec) *string { return v.NodeName }).(pulumi.StringPtrOutput)
}

// Source represents the volume that should be attached.
func (o VolumeAttachmentSpecOutput) Source() VolumeAttachmentSourcePtrOutput {
	return o.ApplyT(func (v VolumeAttachmentSpec) *VolumeAttachmentSource { return v.Source }).(VolumeAttachmentSourcePtrOutput)
}

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatus struct {
	// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachError *VolumeError `pulumi:"attachError"`
	// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Attached *bool `pulumi:"attached"`
	// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata map[string]string `pulumi:"attachmentMetadata"`
	// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
	DetachError *VolumeError `pulumi:"detachError"`
}

type VolumeAttachmentStatusInput interface {
	pulumi.Input

	ToVolumeAttachmentStatusOutput() VolumeAttachmentStatusOutput
	ToVolumeAttachmentStatusOutputWithContext(context.Context) VolumeAttachmentStatusOutput
}

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatusArgs struct {
	// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachError VolumeErrorPtrInput `pulumi:"attachError"`
	// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Attached pulumi.BoolPtrInput `pulumi:"attached"`
	// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata pulumi.StringMapInput `pulumi:"attachmentMetadata"`
	// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
	DetachError VolumeErrorPtrInput `pulumi:"detachError"`
}

func (VolumeAttachmentStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentStatus)(nil)).Elem()
}

func (i VolumeAttachmentStatusArgs) ToVolumeAttachmentStatusOutput() VolumeAttachmentStatusOutput {
	return i.ToVolumeAttachmentStatusOutputWithContext(context.Background())
}

func (i VolumeAttachmentStatusArgs) ToVolumeAttachmentStatusOutputWithContext(ctx context.Context) VolumeAttachmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentStatusOutput)
}

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatusOutput struct { *pulumi.OutputState }

func (VolumeAttachmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentStatus)(nil)).Elem()
}

func (o VolumeAttachmentStatusOutput) ToVolumeAttachmentStatusOutput() VolumeAttachmentStatusOutput {
	return o
}

func (o VolumeAttachmentStatusOutput) ToVolumeAttachmentStatusOutputWithContext(ctx context.Context) VolumeAttachmentStatusOutput {
	return o
}

// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusOutput) AttachError() VolumeErrorPtrOutput {
	return o.ApplyT(func (v VolumeAttachmentStatus) *VolumeError { return v.AttachError }).(VolumeErrorPtrOutput)
}

// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusOutput) Attached() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v VolumeAttachmentStatus) *bool { return v.Attached }).(pulumi.BoolPtrOutput)
}

// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusOutput) AttachmentMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func (v VolumeAttachmentStatus) map[string]string { return v.AttachmentMetadata }).(pulumi.StringMapOutput)
}

// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusOutput) DetachError() VolumeErrorPtrOutput {
	return o.ApplyT(func (v VolumeAttachmentStatus) *VolumeError { return v.DetachError }).(VolumeErrorPtrOutput)
}

// VolumeError captures an error encountered during a volume operation.
type VolumeError struct {
	// String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
	Message *string `pulumi:"message"`
	// Time the error was encountered.
	Time *string `pulumi:"time"`
}

type VolumeErrorInput interface {
	pulumi.Input

	ToVolumeErrorOutput() VolumeErrorOutput
	ToVolumeErrorOutputWithContext(context.Context) VolumeErrorOutput
}

// VolumeError captures an error encountered during a volume operation.
type VolumeErrorArgs struct {
	// String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
	Message pulumi.StringPtrInput `pulumi:"message"`
	// Time the error was encountered.
	Time pulumi.StringPtrInput `pulumi:"time"`
}

func (VolumeErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeError)(nil)).Elem()
}

func (i VolumeErrorArgs) ToVolumeErrorOutput() VolumeErrorOutput {
	return i.ToVolumeErrorOutputWithContext(context.Background())
}

func (i VolumeErrorArgs) ToVolumeErrorOutputWithContext(ctx context.Context) VolumeErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeErrorOutput)
}

// VolumeError captures an error encountered during a volume operation.
type VolumeErrorOutput struct { *pulumi.OutputState }

func (VolumeErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeError)(nil)).Elem()
}

func (o VolumeErrorOutput) ToVolumeErrorOutput() VolumeErrorOutput {
	return o
}

func (o VolumeErrorOutput) ToVolumeErrorOutputWithContext(ctx context.Context) VolumeErrorOutput {
	return o
}

// String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
func (o VolumeErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// Time the error was encountered.
func (o VolumeErrorOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeError) *string { return v.Time }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeAttachmentSourceOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentSpecOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentStatusOutput{})
	pulumi.RegisterOutputType(VolumeErrorOutput{})
}
