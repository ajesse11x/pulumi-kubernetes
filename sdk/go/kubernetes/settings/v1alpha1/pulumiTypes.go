// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// PodPresetSpec is a description of a pod preset.
type PodPresetSpec struct {
	// Env defines the collection of EnvVar to inject into containers.
	Env []corev1.EnvVar `pulumi:"env"`
	// EnvFrom defines the collection of EnvFromSource to inject into containers.
	EnvFrom []corev1.EnvFromSource `pulumi:"envFrom"`
	// Selector is a label query over a set of resources, in this case pods. Required.
	Selector *metav1.LabelSelector `pulumi:"selector"`
	// VolumeMounts defines the collection of VolumeMount to inject into containers.
	VolumeMounts []corev1.VolumeMount `pulumi:"volumeMounts"`
	// Volumes defines the collection of Volume to inject into the pod.
	Volumes []corev1.Volume `pulumi:"volumes"`
}

type PodPresetSpecInput interface {
	pulumi.Input

	ToPodPresetSpecOutput() PodPresetSpecOutput
	ToPodPresetSpecOutputWithContext(context.Context) PodPresetSpecOutput
}

// PodPresetSpec is a description of a pod preset.
type PodPresetSpecArgs struct {
	// Env defines the collection of EnvVar to inject into containers.
	Env corev1.EnvVarArrayInput `pulumi:"env"`
	// EnvFrom defines the collection of EnvFromSource to inject into containers.
	EnvFrom corev1.EnvFromSourceArrayInput `pulumi:"envFrom"`
	// Selector is a label query over a set of resources, in this case pods. Required.
	Selector metav1.LabelSelectorPtrInput `pulumi:"selector"`
	// VolumeMounts defines the collection of VolumeMount to inject into containers.
	VolumeMounts corev1.VolumeMountArrayInput `pulumi:"volumeMounts"`
	// Volumes defines the collection of Volume to inject into the pod.
	Volumes corev1.VolumeArrayInput `pulumi:"volumes"`
}

func (PodPresetSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PodPresetSpec)(nil)).Elem()
}

func (i PodPresetSpecArgs) ToPodPresetSpecOutput() PodPresetSpecOutput {
	return i.ToPodPresetSpecOutputWithContext(context.Background())
}

func (i PodPresetSpecArgs) ToPodPresetSpecOutputWithContext(ctx context.Context) PodPresetSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetSpecOutput)
}

// PodPresetSpec is a description of a pod preset.
type PodPresetSpecOutput struct { *pulumi.OutputState }

func (PodPresetSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PodPresetSpec)(nil)).Elem()
}

func (o PodPresetSpecOutput) ToPodPresetSpecOutput() PodPresetSpecOutput {
	return o
}

func (o PodPresetSpecOutput) ToPodPresetSpecOutputWithContext(ctx context.Context) PodPresetSpecOutput {
	return o
}

// Env defines the collection of EnvVar to inject into containers.
func (o PodPresetSpecOutput) Env() corev1.EnvVarArrayOutput {
	return o.ApplyT(func (v PodPresetSpec) []corev1.EnvVar { return v.Env }).(corev1.EnvVarArrayOutput)
}

// EnvFrom defines the collection of EnvFromSource to inject into containers.
func (o PodPresetSpecOutput) EnvFrom() corev1.EnvFromSourceArrayOutput {
	return o.ApplyT(func (v PodPresetSpec) []corev1.EnvFromSource { return v.EnvFrom }).(corev1.EnvFromSourceArrayOutput)
}

// Selector is a label query over a set of resources, in this case pods. Required.
func (o PodPresetSpecOutput) Selector() metav1.LabelSelectorPtrOutput {
	return o.ApplyT(func (v PodPresetSpec) *metav1.LabelSelector { return v.Selector }).(metav1.LabelSelectorPtrOutput)
}

// VolumeMounts defines the collection of VolumeMount to inject into containers.
func (o PodPresetSpecOutput) VolumeMounts() corev1.VolumeMountArrayOutput {
	return o.ApplyT(func (v PodPresetSpec) []corev1.VolumeMount { return v.VolumeMounts }).(corev1.VolumeMountArrayOutput)
}

// Volumes defines the collection of Volume to inject into the pod.
func (o PodPresetSpecOutput) Volumes() corev1.VolumeArrayOutput {
	return o.ApplyT(func (v PodPresetSpec) []corev1.Volume { return v.Volumes }).(corev1.VolumeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(PodPresetSpecOutput{})
}
