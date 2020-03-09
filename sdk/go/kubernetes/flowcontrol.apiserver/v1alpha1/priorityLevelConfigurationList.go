// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1alpha1

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// PriorityLevelConfigurationList is a list of PriorityLevelConfiguration objects.
type PriorityLevelConfigurationList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// `items` is a list of request-priorities.
	Items flowcontrolv1alpha1.PriorityLevelConfigurationArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewPriorityLevelConfigurationList registers a new resource with the given unique name, arguments, and options.
func NewPriorityLevelConfigurationList(ctx *pulumi.Context,
	name string, args *PriorityLevelConfigurationListArgs, opts ...pulumi.ResourceOption) (*PriorityLevelConfigurationList, error) {
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args == nil {
		args = &PriorityLevelConfigurationListArgs{}
	}
	var resource PriorityLevelConfigurationList
	err := ctx.RegisterResource("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:PriorityLevelConfigurationList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPriorityLevelConfigurationList gets an existing PriorityLevelConfigurationList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPriorityLevelConfigurationList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PriorityLevelConfigurationListState, opts ...pulumi.ResourceOption) (*PriorityLevelConfigurationList, error) {
	var resource PriorityLevelConfigurationList
	err := ctx.ReadResource("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:PriorityLevelConfigurationList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PriorityLevelConfigurationList resources.
type priorityLevelConfigurationListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// `items` is a list of request-priorities.
	Items []flowcontrolv1alpha1.PriorityLevelConfiguration `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type PriorityLevelConfigurationListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// `items` is a list of request-priorities.
	Items flowcontrolv1alpha1.PriorityLevelConfigurationArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (PriorityLevelConfigurationListState) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityLevelConfigurationListState)(nil)).Elem()
}

type priorityLevelConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// `items` is a list of request-priorities.
	Items []flowcontrolv1alpha1.PriorityLevelConfiguration `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PriorityLevelConfigurationList resource.
type PriorityLevelConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// `items` is a list of request-priorities.
	Items flowcontrolv1alpha1.PriorityLevelConfigurationArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (PriorityLevelConfigurationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityLevelConfigurationListArgs)(nil)).Elem()
}
