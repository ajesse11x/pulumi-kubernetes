// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1

import (
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.
type ClusterRoleBinding struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefPtrOutput `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayOutput `pulumi:"subjects"`
}

// NewClusterRoleBinding registers a new resource with the given unique name, arguments, and options.
func NewClusterRoleBinding(ctx *pulumi.Context,
	name string, args *ClusterRoleBindingArgs, opts ...pulumi.ResourceOption) (*ClusterRoleBinding, error) {
	if args == nil || args.RoleRef == nil {
		return nil, errors.New("missing required argument 'RoleRef'")
	}
	if args == nil {
		args = &ClusterRoleBindingArgs{}
	}
	var resource ClusterRoleBinding
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRoleBinding gets an existing ClusterRoleBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRoleBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRoleBindingState, opts ...pulumi.ResourceOption) (*ClusterRoleBinding, error) {
	var resource ClusterRoleBinding
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRoleBinding resources.
type clusterRoleBindingState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef *RoleRef `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects []Subject `pulumi:"subjects"`
}

type ClusterRoleBindingState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefPtrInput
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayInput
}

func (ClusterRoleBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleBindingState)(nil)).Elem()
}

type clusterRoleBindingArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRef `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects []Subject `pulumi:"subjects"`
}

// The set of arguments for constructing a ClusterRoleBinding resource.
type ClusterRoleBindingArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefInput
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayInput
}

func (ClusterRoleBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleBindingArgs)(nil)).Elem()
}
