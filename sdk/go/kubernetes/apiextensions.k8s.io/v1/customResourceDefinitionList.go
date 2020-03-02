// *** WARNING: this file was generated by foo. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// CustomResourceDefinitionList is a list of CustomResourceDefinition objects.
type CustomResourceDefinitionList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items list individual CustomResourceDefinition objects
	Items apiextensionsv1.CustomResourceDefinitionArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewCustomResourceDefinitionList registers a new resource with the given unique name, arguments, and options.
func NewCustomResourceDefinitionList(ctx *pulumi.Context,
	name string, args *CustomResourceDefinitionListArgs, opts ...pulumi.ResourceOption) (*CustomResourceDefinitionList, error) {
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args == nil {
		args = &CustomResourceDefinitionListArgs{}
	}
	var resource CustomResourceDefinitionList
	err := ctx.RegisterResource("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinitionList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomResourceDefinitionList gets an existing CustomResourceDefinitionList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomResourceDefinitionList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomResourceDefinitionListState, opts ...pulumi.ResourceOption) (*CustomResourceDefinitionList, error) {
	var resource CustomResourceDefinitionList
	err := ctx.ReadResource("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinitionList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomResourceDefinitionList resources.
type customResourceDefinitionListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items list individual CustomResourceDefinition objects
	Items []apiextensionsv1.CustomResourceDefinition `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type CustomResourceDefinitionListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items list individual CustomResourceDefinition objects
	Items apiextensionsv1.CustomResourceDefinitionArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (CustomResourceDefinitionListState) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionListState)(nil)).Elem()
}

type customResourceDefinitionListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items list individual CustomResourceDefinition objects
	Items []apiextensionsv1.CustomResourceDefinition `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CustomResourceDefinitionList resource.
type CustomResourceDefinitionListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items list individual CustomResourceDefinition objects
	Items apiextensionsv1.CustomResourceDefinitionArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (CustomResourceDefinitionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionListArgs)(nil)).Elem()
}

