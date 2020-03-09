// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
type Event struct {
	pulumi.CustomResourceState

	// What action was taken/failed regarding to the regarding object.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedCount pulumi.IntPtrOutput `pulumi:"deprecatedCount"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedFirstTimestamp pulumi.StringPtrOutput `pulumi:"deprecatedFirstTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedLastTimestamp pulumi.StringPtrOutput `pulumi:"deprecatedLastTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedSource corev1.EventSourcePtrOutput `pulumi:"deprecatedSource"`
	// Required. Time when this Event was first observed.
	EventTime pulumi.StringPtrOutput `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// Why the action was taken.
	Reason pulumi.StringPtrOutput `pulumi:"reason"`
	// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding corev1.ObjectReferencePtrOutput `pulumi:"regarding"`
	// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related corev1.ObjectReferencePtrOutput `pulumi:"related"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingController pulumi.StringPtrOutput `pulumi:"reportingController"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringPtrOutput `pulumi:"reportingInstance"`
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPtrOutput `pulumi:"series"`
	// Type of this event (Normal, Warning), new types could be added in the future.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewEvent registers a new resource with the given unique name, arguments, and options.
func NewEvent(ctx *pulumi.Context,
	name string, args *EventArgs, opts ...pulumi.ResourceOption) (*Event, error) {
	if args == nil || args.EventTime == nil {
		return nil, errors.New("missing required argument 'EventTime'")
	}
	if args == nil {
		args = &EventArgs{}
	}
	var resource Event
	err := ctx.RegisterResource("kubernetes:events.k8s.io/v1beta1:Event", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEvent gets an existing Event resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventState, opts ...pulumi.ResourceOption) (*Event, error) {
	var resource Event
	err := ctx.ReadResource("kubernetes:events.k8s.io/v1beta1:Event", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Event resources.
type eventState struct {
	// What action was taken/failed regarding to the regarding object.
	Action *string `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedCount *int `pulumi:"deprecatedCount"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedFirstTimestamp *string `pulumi:"deprecatedFirstTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedLastTimestamp *string `pulumi:"deprecatedLastTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedSource *corev1.EventSource `pulumi:"deprecatedSource"`
	// Required. Time when this Event was first observed.
	EventTime *string `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note *string `pulumi:"note"`
	// Why the action was taken.
	Reason *string `pulumi:"reason"`
	// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding *corev1.ObjectReference `pulumi:"regarding"`
	// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related *corev1.ObjectReference `pulumi:"related"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingController *string `pulumi:"reportingController"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance *string `pulumi:"reportingInstance"`
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeries `pulumi:"series"`
	// Type of this event (Normal, Warning), new types could be added in the future.
	Type *string `pulumi:"type"`
}

type EventState struct {
	// What action was taken/failed regarding to the regarding object.
	Action pulumi.StringPtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedCount pulumi.IntPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedFirstTimestamp pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedLastTimestamp pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedSource corev1.EventSourcePtrInput
	// Required. Time when this Event was first observed.
	EventTime pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note pulumi.StringPtrInput
	// Why the action was taken.
	Reason pulumi.StringPtrInput
	// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding corev1.ObjectReferencePtrInput
	// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related corev1.ObjectReferencePtrInput
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingController pulumi.StringPtrInput
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringPtrInput
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPtrInput
	// Type of this event (Normal, Warning), new types could be added in the future.
	Type pulumi.StringPtrInput
}

func (EventState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventState)(nil)).Elem()
}

type eventArgs struct {
	// What action was taken/failed regarding to the regarding object.
	Action *string `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedCount *int `pulumi:"deprecatedCount"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedFirstTimestamp *string `pulumi:"deprecatedFirstTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedLastTimestamp *string `pulumi:"deprecatedLastTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedSource *corev1.EventSource `pulumi:"deprecatedSource"`
	// Required. Time when this Event was first observed.
	EventTime string `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note *string `pulumi:"note"`
	// Why the action was taken.
	Reason *string `pulumi:"reason"`
	// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding *corev1.ObjectReference `pulumi:"regarding"`
	// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related *corev1.ObjectReference `pulumi:"related"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingController *string `pulumi:"reportingController"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance *string `pulumi:"reportingInstance"`
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeries `pulumi:"series"`
	// Type of this event (Normal, Warning), new types could be added in the future.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Event resource.
type EventArgs struct {
	// What action was taken/failed regarding to the regarding object.
	Action pulumi.StringPtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedCount pulumi.IntPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedFirstTimestamp pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedLastTimestamp pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedSource corev1.EventSourcePtrInput
	// Required. Time when this Event was first observed.
	EventTime pulumi.StringInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note pulumi.StringPtrInput
	// Why the action was taken.
	Reason pulumi.StringPtrInput
	// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding corev1.ObjectReferencePtrInput
	// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related corev1.ObjectReferencePtrInput
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingController pulumi.StringPtrInput
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringPtrInput
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPtrInput
	// Type of this event (Normal, Warning), new types could be added in the future.
	Type pulumi.StringPtrInput
}

func (EventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventArgs)(nil)).Elem()
}

