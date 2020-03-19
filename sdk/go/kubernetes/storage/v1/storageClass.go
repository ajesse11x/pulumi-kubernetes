// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1

import (
	"reflect"

	"github.com/pkg/errors"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.
//
// StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
type StorageClass struct {
	pulumi.CustomResourceState

	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion pulumi.BoolPtrOutput `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies corev1.TopologySelectorTermArrayOutput `pulumi:"allowedTopologies"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions pulumi.StringArrayOutput `pulumi:"mountOptions"`
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Provisioner indicates the type of the provisioner.
	Provisioner pulumi.StringPtrOutput `pulumi:"provisioner"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy pulumi.StringPtrOutput `pulumi:"reclaimPolicy"`
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode pulumi.StringPtrOutput `pulumi:"volumeBindingMode"`
}

// NewStorageClass registers a new resource with the given unique name, arguments, and options.
func NewStorageClass(ctx *pulumi.Context,
	name string, args *StorageClassArgs, opts ...pulumi.ResourceOption) (*StorageClass, error) {
	if args == nil || args.Provisioner == nil {
		return nil, errors.New("missing required argument 'Provisioner'")
	}
	if args == nil {
		args = &StorageClassArgs{}
	}
	if args.ApiVersion == nil {
		args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	}
	if args.Kind == nil {
		args.Kind = pulumi.StringPtr("StorageClass")
	}
	var resource StorageClass
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:StorageClass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageClass gets an existing StorageClass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageClass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageClassState, opts ...pulumi.ResourceOption) (*StorageClass, error) {
	var resource StorageClass
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:StorageClass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageClass resources.
type storageClassState struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion *bool `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies []corev1.TopologySelectorTerm `pulumi:"allowedTopologies"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions []string `pulumi:"mountOptions"`
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters map[string]string `pulumi:"parameters"`
	// Provisioner indicates the type of the provisioner.
	Provisioner *string `pulumi:"provisioner"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy *string `pulumi:"reclaimPolicy"`
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode *string `pulumi:"volumeBindingMode"`
}

type StorageClassState struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion pulumi.BoolPtrInput
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies corev1.TopologySelectorTermArrayInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions pulumi.StringArrayInput
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters pulumi.StringMapInput
	// Provisioner indicates the type of the provisioner.
	Provisioner pulumi.StringPtrInput
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy pulumi.StringPtrInput
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode pulumi.StringPtrInput
}

func (StorageClassState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageClassState)(nil)).Elem()
}

type storageClassArgs struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion *bool `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies []corev1.TopologySelectorTerm `pulumi:"allowedTopologies"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions []string `pulumi:"mountOptions"`
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters map[string]string `pulumi:"parameters"`
	// Provisioner indicates the type of the provisioner.
	Provisioner string `pulumi:"provisioner"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy *string `pulumi:"reclaimPolicy"`
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode *string `pulumi:"volumeBindingMode"`
}

// The set of arguments for constructing a StorageClass resource.
type StorageClassArgs struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion pulumi.BoolPtrInput
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies corev1.TopologySelectorTermArrayInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions pulumi.StringArrayInput
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters pulumi.StringMapInput
	// Provisioner indicates the type of the provisioner.
	Provisioner pulumi.StringInput
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy pulumi.StringPtrInput
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode pulumi.StringPtrInput
}

func (StorageClassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageClassArgs)(nil)).Elem()
}
