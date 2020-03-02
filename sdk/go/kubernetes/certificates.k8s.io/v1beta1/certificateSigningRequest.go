// *** WARNING: this file was generated by foo. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/certificates/v1beta1"
	"github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Describes a certificate signing request
type CertificateSigningRequest struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec certificatesv1beta1.CertificateSigningRequestSpecPtrOutput `pulumi:"spec"`
	// Derived information about the request.
	Status certificatesv1beta1.CertificateSigningRequestStatusPtrOutput `pulumi:"status"`
}

// NewCertificateSigningRequest registers a new resource with the given unique name, arguments, and options.
func NewCertificateSigningRequest(ctx *pulumi.Context,
	name string, args *CertificateSigningRequestArgs, opts ...pulumi.ResourceOption) (*CertificateSigningRequest, error) {
	if args == nil {
		args = &CertificateSigningRequestArgs{}
	}
	var resource CertificateSigningRequest
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateSigningRequest gets an existing CertificateSigningRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateSigningRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateSigningRequestState, opts ...pulumi.ResourceOption) (*CertificateSigningRequest, error) {
	var resource CertificateSigningRequest
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateSigningRequest resources.
type certificateSigningRequestState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec *certificatesv1beta1.CertificateSigningRequestSpec `pulumi:"spec"`
	// Derived information about the request.
	Status *certificatesv1beta1.CertificateSigningRequestStatus `pulumi:"status"`
}

type CertificateSigningRequestState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// The certificate request itself and any additional information.
	Spec certificatesv1beta1.CertificateSigningRequestSpecPtrInput
	// Derived information about the request.
	Status certificatesv1beta1.CertificateSigningRequestStatusPtrInput
}

func (CertificateSigningRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestState)(nil)).Elem()
}

type certificateSigningRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec *certificatesv1beta1.CertificateSigningRequestSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CertificateSigningRequest resource.
type CertificateSigningRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// The certificate request itself and any additional information.
	Spec certificatesv1beta1.CertificateSigningRequestSpecPtrInput
}

func (CertificateSigningRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestArgs)(nil)).Elem()
}

