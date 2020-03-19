// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass. PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
type PriorityClass struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description *string `pulumi:"description"`
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault *bool `pulumi:"globalDefault"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
	PreemptionPolicy *string `pulumi:"preemptionPolicy"`
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value *int `pulumi:"value"`
}

type PriorityClassInput interface {
	pulumi.Input

	ToPriorityClassOutput() PriorityClassOutput
	ToPriorityClassOutputWithContext(context.Context) PriorityClassOutput
}

// DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass. PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
type PriorityClassArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault pulumi.BoolPtrInput `pulumi:"globalDefault"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput `pulumi:"metadata"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
	PreemptionPolicy pulumi.StringPtrInput `pulumi:"preemptionPolicy"`
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value pulumi.IntPtrInput `pulumi:"value"`
}

func (PriorityClassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PriorityClass)(nil)).Elem()
}

func (i PriorityClassArgs) ToPriorityClassOutput() PriorityClassOutput {
	return i.ToPriorityClassOutputWithContext(context.Background())
}

func (i PriorityClassArgs) ToPriorityClassOutputWithContext(ctx context.Context) PriorityClassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassOutput)
}

type PriorityClassArrayInput interface {
	pulumi.Input

	ToPriorityClassArrayOutput() PriorityClassArrayOutput
	ToPriorityClassArrayOutputWithContext(context.Context) PriorityClassArrayOutput
}

type PriorityClassArray []PriorityClassInput

func (PriorityClassArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PriorityClass)(nil)).Elem()
}

func (i PriorityClassArray) ToPriorityClassArrayOutput() PriorityClassArrayOutput {
	return i.ToPriorityClassArrayOutputWithContext(context.Background())
}

func (i PriorityClassArray) ToPriorityClassArrayOutputWithContext(ctx context.Context) PriorityClassArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassArrayOutput)
}

// DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass. PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
type PriorityClassOutput struct { *pulumi.OutputState }

func (PriorityClassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PriorityClass)(nil)).Elem()
}

func (o PriorityClassOutput) ToPriorityClassOutput() PriorityClassOutput {
	return o
}

func (o PriorityClassOutput) ToPriorityClassOutputWithContext(ctx context.Context) PriorityClassOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PriorityClassOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PriorityClass) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
func (o PriorityClassOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PriorityClass) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
func (o PriorityClassOutput) GlobalDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v PriorityClass) *bool { return v.GlobalDefault }).(pulumi.BoolPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PriorityClassOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PriorityClass) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PriorityClassOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func (v PriorityClass) *metav1.ObjectMeta { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
func (o PriorityClassOutput) PreemptionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PriorityClass) *string { return v.PreemptionPolicy }).(pulumi.StringPtrOutput)
}

// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
func (o PriorityClassOutput) Value() pulumi.IntPtrOutput {
	return o.ApplyT(func (v PriorityClass) *int { return v.Value }).(pulumi.IntPtrOutput)
}

type PriorityClassArrayOutput struct { *pulumi.OutputState }

func (PriorityClassArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PriorityClass)(nil)).Elem()
}

func (o PriorityClassArrayOutput) ToPriorityClassArrayOutput() PriorityClassArrayOutput {
	return o
}

func (o PriorityClassArrayOutput) ToPriorityClassArrayOutputWithContext(ctx context.Context) PriorityClassArrayOutput {
	return o
}

func (o PriorityClassArrayOutput) Index(i pulumi.IntInput) PriorityClassOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) PriorityClass {
		return vs[0].([]PriorityClass)[vs[1].(int)]
	}).(PriorityClassOutput)
}

// PriorityClassList is a collection of priority classes.
type PriorityClassList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of PriorityClasses
	Items []PriorityClass `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type PriorityClassListInput interface {
	pulumi.Input

	ToPriorityClassListOutput() PriorityClassListOutput
	ToPriorityClassListOutputWithContext(context.Context) PriorityClassListOutput
}

// PriorityClassList is a collection of priority classes.
type PriorityClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// items is the list of PriorityClasses
	Items PriorityClassArrayInput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput `pulumi:"metadata"`
}

func (PriorityClassListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PriorityClassList)(nil)).Elem()
}

func (i PriorityClassListArgs) ToPriorityClassListOutput() PriorityClassListOutput {
	return i.ToPriorityClassListOutputWithContext(context.Background())
}

func (i PriorityClassListArgs) ToPriorityClassListOutputWithContext(ctx context.Context) PriorityClassListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassListOutput)
}

// PriorityClassList is a collection of priority classes.
type PriorityClassListOutput struct { *pulumi.OutputState }

func (PriorityClassListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PriorityClassList)(nil)).Elem()
}

func (o PriorityClassListOutput) ToPriorityClassListOutput() PriorityClassListOutput {
	return o
}

func (o PriorityClassListOutput) ToPriorityClassListOutputWithContext(ctx context.Context) PriorityClassListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PriorityClassListOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PriorityClassList) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// items is the list of PriorityClasses
func (o PriorityClassListOutput) Items() PriorityClassArrayOutput {
	return o.ApplyT(func (v PriorityClassList) []PriorityClass { return v.Items }).(PriorityClassArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PriorityClassListOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PriorityClassList) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PriorityClassListOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func (v PriorityClassList) *metav1.ListMeta { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PriorityClassOutput{})
	pulumi.RegisterOutputType(PriorityClassArrayOutput{})
	pulumi.RegisterOutputType(PriorityClassListOutput{})
}
