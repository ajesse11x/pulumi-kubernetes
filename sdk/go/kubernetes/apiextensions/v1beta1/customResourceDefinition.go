// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format <.spec.name>.<.spec.group>. Deprecated in v1.16, planned for removal in v1.19. Use apiextensions.k8s.io/v1 CustomResourceDefinition instead.
type CustomResourceDefinition struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecPtrOutput `pulumi:"spec"`
	// status indicates the actual state of the CustomResourceDefinition
	Status CustomResourceDefinitionStatusPtrOutput `pulumi:"status"`
}

// NewCustomResourceDefinition registers a new resource with the given unique name, arguments, and options.
func NewCustomResourceDefinition(ctx *pulumi.Context,
	name string, args *CustomResourceDefinitionArgs, opts ...pulumi.ResourceOption) (*CustomResourceDefinition, error) {
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil {
		args = &CustomResourceDefinitionArgs{}
	}
	if args.ApiVersion == nil {
		args.ApiVersion = pulumi.StringPtr("apiextensions.k8s.io/v1beta1")
	}
	if args.Kind == nil {
		args.Kind = pulumi.StringPtr("CustomResourceDefinition")
	}
	var resource CustomResourceDefinition
	err := ctx.RegisterResource("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomResourceDefinition gets an existing CustomResourceDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomResourceDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomResourceDefinitionState, opts ...pulumi.ResourceOption) (*CustomResourceDefinition, error) {
	var resource CustomResourceDefinition
	err := ctx.ReadResource("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomResourceDefinition resources.
type customResourceDefinitionState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec *CustomResourceDefinitionSpec `pulumi:"spec"`
	// status indicates the actual state of the CustomResourceDefinition
	Status *CustomResourceDefinitionStatus `pulumi:"status"`
}

type CustomResourceDefinitionState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecPtrInput
	// status indicates the actual state of the CustomResourceDefinition
	Status CustomResourceDefinitionStatusPtrInput
}

func (CustomResourceDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionState)(nil)).Elem()
}

type customResourceDefinitionArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CustomResourceDefinition resource.
type CustomResourceDefinitionArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecInput
}

func (CustomResourceDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionArgs)(nil)).Elem()
}
