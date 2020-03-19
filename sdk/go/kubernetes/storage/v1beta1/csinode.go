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

// CSINode holds information about all CSI drivers installed on a node. CSI drivers do not need to create the CSINode object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node object.
type CSINode struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec CSINodeSpecPtrOutput `pulumi:"spec"`
}

// NewCSINode registers a new resource with the given unique name, arguments, and options.
func NewCSINode(ctx *pulumi.Context,
	name string, args *CSINodeArgs, opts ...pulumi.ResourceOption) (*CSINode, error) {
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil {
		args = &CSINodeArgs{}
	}
	if args.ApiVersion == nil {
		args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1beta1")
	}
	if args.Kind == nil {
		args.Kind = pulumi.StringPtr("CSINode")
	}
	var resource CSINode
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1beta1:CSINode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSINode gets an existing CSINode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSINode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSINodeState, opts ...pulumi.ResourceOption) (*CSINode, error) {
	var resource CSINode
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1beta1:CSINode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSINode resources.
type csinodeState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata.name must be the Kubernetes node name.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec *CSINodeSpec `pulumi:"spec"`
}

type CSINodeState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaPtrInput
	// spec is the specification of CSINode
	Spec CSINodeSpecPtrInput
}

func (CSINodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodeState)(nil)).Elem()
}

type csinodeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata.name must be the Kubernetes node name.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec CSINodeSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CSINode resource.
type CSINodeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaPtrInput
	// spec is the specification of CSINode
	Spec CSINodeSpecInput
}

func (CSINodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodeArgs)(nil)).Elem()
}
