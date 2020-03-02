// *** WARNING: this file was generated by foo. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// JobCondition describes current state of a job.
type JobCondition struct {
	// Last time the condition was checked.
	LastProbeTime *string `pulumi:"lastProbeTime"`
	// Last time the condition transit from one status to another.
	LastTransitionTime *string `pulumi:"lastTransitionTime"`
	// Human readable message indicating details about last transition.
	Message *string `pulumi:"message"`
	// (brief) reason for the condition's last transition.
	Reason *string `pulumi:"reason"`
	// Status of the condition, one of True, False, Unknown.
	Status *string `pulumi:"status"`
	// Type of job condition, Complete or Failed.
	Type *string `pulumi:"type"`
}

type JobConditionInput interface {
	pulumi.Input

	ToJobConditionOutput() JobConditionOutput
	ToJobConditionOutputWithContext(context.Context) JobConditionOutput
}

// JobCondition describes current state of a job.
type JobConditionArgs struct {
	// Last time the condition was checked.
	LastProbeTime pulumi.StringPtrInput `pulumi:"lastProbeTime"`
	// Last time the condition transit from one status to another.
	LastTransitionTime pulumi.StringPtrInput `pulumi:"lastTransitionTime"`
	// Human readable message indicating details about last transition.
	Message pulumi.StringPtrInput `pulumi:"message"`
	// (brief) reason for the condition's last transition.
	Reason pulumi.StringPtrInput `pulumi:"reason"`
	// Status of the condition, one of True, False, Unknown.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Type of job condition, Complete or Failed.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (JobConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCondition)(nil)).Elem()
}

func (i JobConditionArgs) ToJobConditionOutput() JobConditionOutput {
	return i.ToJobConditionOutputWithContext(context.Background())
}

func (i JobConditionArgs) ToJobConditionOutputWithContext(ctx context.Context) JobConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobConditionOutput)
}

type JobConditionArrayInput interface {
	pulumi.Input

	ToJobConditionArrayOutput() JobConditionArrayOutput
	ToJobConditionArrayOutputWithContext(context.Context) JobConditionArrayOutput
}

type JobConditionArray []JobConditionInput

func (JobConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobCondition)(nil)).Elem()
}

func (i JobConditionArray) ToJobConditionArrayOutput() JobConditionArrayOutput {
	return i.ToJobConditionArrayOutputWithContext(context.Background())
}

func (i JobConditionArray) ToJobConditionArrayOutputWithContext(ctx context.Context) JobConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobConditionArrayOutput)
}

// JobCondition describes current state of a job.
type JobConditionOutput struct { *pulumi.OutputState }

func (JobConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCondition)(nil)).Elem()
}

func (o JobConditionOutput) ToJobConditionOutput() JobConditionOutput {
	return o
}

func (o JobConditionOutput) ToJobConditionOutputWithContext(ctx context.Context) JobConditionOutput {
	return o
}

// Last time the condition was checked.
func (o JobConditionOutput) LastProbeTime() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobCondition) *string { return v.LastProbeTime }).(pulumi.StringPtrOutput)
}

// Last time the condition transit from one status to another.
func (o JobConditionOutput) LastTransitionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobCondition) *string { return v.LastTransitionTime }).(pulumi.StringPtrOutput)
}

// Human readable message indicating details about last transition.
func (o JobConditionOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobCondition) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// (brief) reason for the condition's last transition.
func (o JobConditionOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobCondition) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

// Status of the condition, one of True, False, Unknown.
func (o JobConditionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobCondition) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Type of job condition, Complete or Failed.
func (o JobConditionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobCondition) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobConditionArrayOutput struct { *pulumi.OutputState}

func (JobConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobCondition)(nil)).Elem()
}

func (o JobConditionArrayOutput) ToJobConditionArrayOutput() JobConditionArrayOutput {
	return o
}

func (o JobConditionArrayOutput) ToJobConditionArrayOutputWithContext(ctx context.Context) JobConditionArrayOutput {
	return o
}

func (o JobConditionArrayOutput) Index(i pulumi.IntInput) JobConditionOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) JobCondition {
		return vs[0].([]JobCondition)[vs[1].(int)]
	}).(JobConditionOutput)
}

// JobSpec describes how the job execution will look like.
type JobSpec struct {
	// Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
	ActiveDeadlineSeconds *int `pulumi:"activeDeadlineSeconds"`
	// Specifies the number of retries before marking this job failed. Defaults to 6
	BackoffLimit *int `pulumi:"backoffLimit"`
	// Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Completions *int `pulumi:"completions"`
	// manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
	ManualSelector *bool `pulumi:"manualSelector"`
	// Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Parallelism *int `pulumi:"parallelism"`
	// A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *metav1.LabelSelector `pulumi:"selector"`
	// Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Template *corev1.PodTemplateSpec `pulumi:"template"`
	// ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.
	TtlSecondsAfterFinished *int `pulumi:"ttlSecondsAfterFinished"`
}

type JobSpecInput interface {
	pulumi.Input

	ToJobSpecOutput() JobSpecOutput
	ToJobSpecOutputWithContext(context.Context) JobSpecOutput
}

// JobSpec describes how the job execution will look like.
type JobSpecArgs struct {
	// Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
	ActiveDeadlineSeconds pulumi.IntPtrInput `pulumi:"activeDeadlineSeconds"`
	// Specifies the number of retries before marking this job failed. Defaults to 6
	BackoffLimit pulumi.IntPtrInput `pulumi:"backoffLimit"`
	// Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Completions pulumi.IntPtrInput `pulumi:"completions"`
	// manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
	ManualSelector pulumi.BoolPtrInput `pulumi:"manualSelector"`
	// Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Parallelism pulumi.IntPtrInput `pulumi:"parallelism"`
	// A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector metav1.LabelSelectorPtrInput `pulumi:"selector"`
	// Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Template corev1.PodTemplateSpecPtrInput `pulumi:"template"`
	// ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.
	TtlSecondsAfterFinished pulumi.IntPtrInput `pulumi:"ttlSecondsAfterFinished"`
}

func (JobSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSpec)(nil)).Elem()
}

func (i JobSpecArgs) ToJobSpecOutput() JobSpecOutput {
	return i.ToJobSpecOutputWithContext(context.Background())
}

func (i JobSpecArgs) ToJobSpecOutputWithContext(ctx context.Context) JobSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSpecOutput)
}

func (i JobSpecArgs) ToJobSpecPtrOutput() JobSpecPtrOutput {
	return i.ToJobSpecPtrOutputWithContext(context.Background())
}

func (i JobSpecArgs) ToJobSpecPtrOutputWithContext(ctx context.Context) JobSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSpecOutput).ToJobSpecPtrOutputWithContext(ctx)
}

type JobSpecPtrInput interface {
	pulumi.Input

	ToJobSpecPtrOutput() JobSpecPtrOutput
	ToJobSpecPtrOutputWithContext(context.Context) JobSpecPtrOutput
}

type jobSpecPtrType JobSpecArgs

func JobSpecPtr(v *JobSpecArgs) JobSpecPtrInput {	return (*jobSpecPtrType)(v)
}

func (*jobSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSpec)(nil)).Elem()
}

func (i *jobSpecPtrType) ToJobSpecPtrOutput() JobSpecPtrOutput {
	return i.ToJobSpecPtrOutputWithContext(context.Background())
}

func (i *jobSpecPtrType) ToJobSpecPtrOutputWithContext(ctx context.Context) JobSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSpecPtrOutput)
}

// JobSpec describes how the job execution will look like.
type JobSpecOutput struct { *pulumi.OutputState }

func (JobSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSpec)(nil)).Elem()
}

func (o JobSpecOutput) ToJobSpecOutput() JobSpecOutput {
	return o
}

func (o JobSpecOutput) ToJobSpecOutputWithContext(ctx context.Context) JobSpecOutput {
	return o
}

func (o JobSpecOutput) ToJobSpecPtrOutput() JobSpecPtrOutput {
	return o.ToJobSpecPtrOutputWithContext(context.Background())
}

func (o JobSpecOutput) ToJobSpecPtrOutputWithContext(ctx context.Context) JobSpecPtrOutput {
	return o.ApplyT(func(v JobSpec) *JobSpec {
		return &v
	}).(JobSpecPtrOutput)
}
// Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
func (o JobSpecOutput) ActiveDeadlineSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.ActiveDeadlineSeconds }).(pulumi.IntPtrOutput)
}

// Specifies the number of retries before marking this job failed. Defaults to 6
func (o JobSpecOutput) BackoffLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.BackoffLimit }).(pulumi.IntPtrOutput)
}

// Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
func (o JobSpecOutput) Completions() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.Completions }).(pulumi.IntPtrOutput)
}

// manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
func (o JobSpecOutput) ManualSelector() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v JobSpec) *bool { return v.ManualSelector }).(pulumi.BoolPtrOutput)
}

// Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
func (o JobSpecOutput) Parallelism() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.Parallelism }).(pulumi.IntPtrOutput)
}

// A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
func (o JobSpecOutput) Selector() metav1.LabelSelectorPtrOutput {
	return o.ApplyT(func (v JobSpec) *metav1.LabelSelector { return v.Selector }).(metav1.LabelSelectorPtrOutput)
}

// Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
func (o JobSpecOutput) Template() corev1.PodTemplateSpecPtrOutput {
	return o.ApplyT(func (v JobSpec) *corev1.PodTemplateSpec { return v.Template }).(corev1.PodTemplateSpecPtrOutput)
}

// ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.
func (o JobSpecOutput) TtlSecondsAfterFinished() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.TtlSecondsAfterFinished }).(pulumi.IntPtrOutput)
}

type JobSpecPtrOutput struct { *pulumi.OutputState}

func (JobSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSpec)(nil)).Elem()
}

func (o JobSpecPtrOutput) ToJobSpecPtrOutput() JobSpecPtrOutput {
	return o
}

func (o JobSpecPtrOutput) ToJobSpecPtrOutputWithContext(ctx context.Context) JobSpecPtrOutput {
	return o
}

func (o JobSpecPtrOutput) Elem() JobSpecOutput {
	return o.ApplyT(func (v *JobSpec) JobSpec { return *v }).(JobSpecOutput)
}

// Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
func (o JobSpecPtrOutput) ActiveDeadlineSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.ActiveDeadlineSeconds }).(pulumi.IntPtrOutput)
}

// Specifies the number of retries before marking this job failed. Defaults to 6
func (o JobSpecPtrOutput) BackoffLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.BackoffLimit }).(pulumi.IntPtrOutput)
}

// Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
func (o JobSpecPtrOutput) Completions() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.Completions }).(pulumi.IntPtrOutput)
}

// manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
func (o JobSpecPtrOutput) ManualSelector() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v JobSpec) *bool { return v.ManualSelector }).(pulumi.BoolPtrOutput)
}

// Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
func (o JobSpecPtrOutput) Parallelism() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.Parallelism }).(pulumi.IntPtrOutput)
}

// A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
func (o JobSpecPtrOutput) Selector() metav1.LabelSelectorPtrOutput {
	return o.ApplyT(func (v JobSpec) *metav1.LabelSelector { return v.Selector }).(metav1.LabelSelectorPtrOutput)
}

// Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
func (o JobSpecPtrOutput) Template() corev1.PodTemplateSpecPtrOutput {
	return o.ApplyT(func (v JobSpec) *corev1.PodTemplateSpec { return v.Template }).(corev1.PodTemplateSpecPtrOutput)
}

// ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.
func (o JobSpecPtrOutput) TtlSecondsAfterFinished() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobSpec) *int { return v.TtlSecondsAfterFinished }).(pulumi.IntPtrOutput)
}

// JobStatus represents the current state of a Job.
type JobStatus struct {
	// The number of actively running pods.
	Active *int `pulumi:"active"`
	// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
	CompletionTime *string `pulumi:"completionTime"`
	// The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Conditions []JobCondition `pulumi:"conditions"`
	// The number of pods which reached phase Failed.
	Failed *int `pulumi:"failed"`
	// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
	StartTime *string `pulumi:"startTime"`
	// The number of pods which reached phase Succeeded.
	Succeeded *int `pulumi:"succeeded"`
}

type JobStatusInput interface {
	pulumi.Input

	ToJobStatusOutput() JobStatusOutput
	ToJobStatusOutputWithContext(context.Context) JobStatusOutput
}

// JobStatus represents the current state of a Job.
type JobStatusArgs struct {
	// The number of actively running pods.
	Active pulumi.IntPtrInput `pulumi:"active"`
	// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
	CompletionTime pulumi.StringPtrInput `pulumi:"completionTime"`
	// The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Conditions JobConditionArrayInput `pulumi:"conditions"`
	// The number of pods which reached phase Failed.
	Failed pulumi.IntPtrInput `pulumi:"failed"`
	// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
	// The number of pods which reached phase Succeeded.
	Succeeded pulumi.IntPtrInput `pulumi:"succeeded"`
}

func (JobStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStatus)(nil)).Elem()
}

func (i JobStatusArgs) ToJobStatusOutput() JobStatusOutput {
	return i.ToJobStatusOutputWithContext(context.Background())
}

func (i JobStatusArgs) ToJobStatusOutputWithContext(ctx context.Context) JobStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusOutput)
}

func (i JobStatusArgs) ToJobStatusPtrOutput() JobStatusPtrOutput {
	return i.ToJobStatusPtrOutputWithContext(context.Background())
}

func (i JobStatusArgs) ToJobStatusPtrOutputWithContext(ctx context.Context) JobStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusOutput).ToJobStatusPtrOutputWithContext(ctx)
}

type JobStatusPtrInput interface {
	pulumi.Input

	ToJobStatusPtrOutput() JobStatusPtrOutput
	ToJobStatusPtrOutputWithContext(context.Context) JobStatusPtrOutput
}

type jobStatusPtrType JobStatusArgs

func JobStatusPtr(v *JobStatusArgs) JobStatusPtrInput {	return (*jobStatusPtrType)(v)
}

func (*jobStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStatus)(nil)).Elem()
}

func (i *jobStatusPtrType) ToJobStatusPtrOutput() JobStatusPtrOutput {
	return i.ToJobStatusPtrOutputWithContext(context.Background())
}

func (i *jobStatusPtrType) ToJobStatusPtrOutputWithContext(ctx context.Context) JobStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusPtrOutput)
}

// JobStatus represents the current state of a Job.
type JobStatusOutput struct { *pulumi.OutputState }

func (JobStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStatus)(nil)).Elem()
}

func (o JobStatusOutput) ToJobStatusOutput() JobStatusOutput {
	return o
}

func (o JobStatusOutput) ToJobStatusOutputWithContext(ctx context.Context) JobStatusOutput {
	return o
}

func (o JobStatusOutput) ToJobStatusPtrOutput() JobStatusPtrOutput {
	return o.ToJobStatusPtrOutputWithContext(context.Background())
}

func (o JobStatusOutput) ToJobStatusPtrOutputWithContext(ctx context.Context) JobStatusPtrOutput {
	return o.ApplyT(func(v JobStatus) *JobStatus {
		return &v
	}).(JobStatusPtrOutput)
}
// The number of actively running pods.
func (o JobStatusOutput) Active() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobStatus) *int { return v.Active }).(pulumi.IntPtrOutput)
}

// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
func (o JobStatusOutput) CompletionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobStatus) *string { return v.CompletionTime }).(pulumi.StringPtrOutput)
}

// The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
func (o JobStatusOutput) Conditions() JobConditionArrayOutput {
	return o.ApplyT(func (v JobStatus) []JobCondition { return v.Conditions }).(JobConditionArrayOutput)
}

// The number of pods which reached phase Failed.
func (o JobStatusOutput) Failed() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobStatus) *int { return v.Failed }).(pulumi.IntPtrOutput)
}

// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
func (o JobStatusOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobStatus) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// The number of pods which reached phase Succeeded.
func (o JobStatusOutput) Succeeded() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobStatus) *int { return v.Succeeded }).(pulumi.IntPtrOutput)
}

type JobStatusPtrOutput struct { *pulumi.OutputState}

func (JobStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStatus)(nil)).Elem()
}

func (o JobStatusPtrOutput) ToJobStatusPtrOutput() JobStatusPtrOutput {
	return o
}

func (o JobStatusPtrOutput) ToJobStatusPtrOutputWithContext(ctx context.Context) JobStatusPtrOutput {
	return o
}

func (o JobStatusPtrOutput) Elem() JobStatusOutput {
	return o.ApplyT(func (v *JobStatus) JobStatus { return *v }).(JobStatusOutput)
}

// The number of actively running pods.
func (o JobStatusPtrOutput) Active() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobStatus) *int { return v.Active }).(pulumi.IntPtrOutput)
}

// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
func (o JobStatusPtrOutput) CompletionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobStatus) *string { return v.CompletionTime }).(pulumi.StringPtrOutput)
}

// The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
func (o JobStatusPtrOutput) Conditions() JobConditionArrayOutput {
	return o.ApplyT(func (v JobStatus) []JobCondition { return v.Conditions }).(JobConditionArrayOutput)
}

// The number of pods which reached phase Failed.
func (o JobStatusPtrOutput) Failed() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobStatus) *int { return v.Failed }).(pulumi.IntPtrOutput)
}

// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
func (o JobStatusPtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobStatus) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// The number of pods which reached phase Succeeded.
func (o JobStatusPtrOutput) Succeeded() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobStatus) *int { return v.Succeeded }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(JobConditionOutput{})
	pulumi.RegisterOutputType(JobConditionArrayOutput{})
	pulumi.RegisterOutputType(JobSpecOutput{})
	pulumi.RegisterOutputType(JobSpecPtrOutput{})
	pulumi.RegisterOutputType(JobStatusOutput{})
	pulumi.RegisterOutputType(JobStatusPtrOutput{})
}
