// *** WARNING: this file was generated by foo. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// StatefulSet represents a set of pods with consistent identities. Identities are defined as:
//  - Network: A single stable DNS and hostname.
//  - Storage: As many VolumeClaims as requested.
// The StatefulSet guarantees that a given network identity will always map to the same storage identity.
type StatefulSet struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecPtrOutput `pulumi:"spec"`
	// Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
	Status StatefulSetStatusPtrOutput `pulumi:"status"`
}

// NewStatefulSet registers a new resource with the given unique name, arguments, and options.
func NewStatefulSet(ctx *pulumi.Context,
	name string, args *StatefulSetArgs, opts ...pulumi.ResourceOption) (*StatefulSet, error) {
	if args == nil {
		args = &StatefulSetArgs{}
	}
	var resource StatefulSet
	err := ctx.RegisterResource("kubernetes:apps/v1beta1:StatefulSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatefulSet gets an existing StatefulSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatefulSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatefulSetState, opts ...pulumi.ResourceOption) (*StatefulSet, error) {
	var resource StatefulSet
	err := ctx.ReadResource("kubernetes:apps/v1beta1:StatefulSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StatefulSet resources.
type statefulSetState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired identities of pods in this set.
	Spec *StatefulSetSpec `pulumi:"spec"`
	// Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
	Status *StatefulSetStatus `pulumi:"status"`
}

type StatefulSetState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecPtrInput
	// Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
	Status StatefulSetStatusPtrInput
}

func (StatefulSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetState)(nil)).Elem()
}

type statefulSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired identities of pods in this set.
	Spec *StatefulSetSpec `pulumi:"spec"`
}

// The set of arguments for constructing a StatefulSet resource.
type StatefulSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecPtrInput
}

func (StatefulSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetArgs)(nil)).Elem()
}

