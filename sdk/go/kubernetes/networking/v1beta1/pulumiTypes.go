// *** WARNING: this file was generated by foo. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
)

// HTTPIngressPath associates a path regex with a backend. Incoming urls matching the path are forwarded to the backend.
type HTTPIngressPath struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	Backend *IngressBackend `pulumi:"backend"`
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.
	Path *string `pulumi:"path"`
}

type HTTPIngressPathInput interface {
	pulumi.Input

	ToHTTPIngressPathOutput() HTTPIngressPathOutput
	ToHTTPIngressPathOutputWithContext(context.Context) HTTPIngressPathOutput
}

// HTTPIngressPath associates a path regex with a backend. Incoming urls matching the path are forwarded to the backend.
type HTTPIngressPathArgs struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	Backend IngressBackendPtrInput `pulumi:"backend"`
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.
	Path pulumi.StringPtrInput `pulumi:"path"`
}

func (HTTPIngressPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HTTPIngressPath)(nil)).Elem()
}

func (i HTTPIngressPathArgs) ToHTTPIngressPathOutput() HTTPIngressPathOutput {
	return i.ToHTTPIngressPathOutputWithContext(context.Background())
}

func (i HTTPIngressPathArgs) ToHTTPIngressPathOutputWithContext(ctx context.Context) HTTPIngressPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HTTPIngressPathOutput)
}

type HTTPIngressPathArrayInput interface {
	pulumi.Input

	ToHTTPIngressPathArrayOutput() HTTPIngressPathArrayOutput
	ToHTTPIngressPathArrayOutputWithContext(context.Context) HTTPIngressPathArrayOutput
}

type HTTPIngressPathArray []HTTPIngressPathInput

func (HTTPIngressPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HTTPIngressPath)(nil)).Elem()
}

func (i HTTPIngressPathArray) ToHTTPIngressPathArrayOutput() HTTPIngressPathArrayOutput {
	return i.ToHTTPIngressPathArrayOutputWithContext(context.Background())
}

func (i HTTPIngressPathArray) ToHTTPIngressPathArrayOutputWithContext(ctx context.Context) HTTPIngressPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HTTPIngressPathArrayOutput)
}

// HTTPIngressPath associates a path regex with a backend. Incoming urls matching the path are forwarded to the backend.
type HTTPIngressPathOutput struct { *pulumi.OutputState }

func (HTTPIngressPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HTTPIngressPath)(nil)).Elem()
}

func (o HTTPIngressPathOutput) ToHTTPIngressPathOutput() HTTPIngressPathOutput {
	return o
}

func (o HTTPIngressPathOutput) ToHTTPIngressPathOutputWithContext(ctx context.Context) HTTPIngressPathOutput {
	return o
}

// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
func (o HTTPIngressPathOutput) Backend() IngressBackendPtrOutput {
	return o.ApplyT(func (v HTTPIngressPath) *IngressBackend { return v.Backend }).(IngressBackendPtrOutput)
}

// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.
func (o HTTPIngressPathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v HTTPIngressPath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type HTTPIngressPathArrayOutput struct { *pulumi.OutputState}

func (HTTPIngressPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HTTPIngressPath)(nil)).Elem()
}

func (o HTTPIngressPathArrayOutput) ToHTTPIngressPathArrayOutput() HTTPIngressPathArrayOutput {
	return o
}

func (o HTTPIngressPathArrayOutput) ToHTTPIngressPathArrayOutputWithContext(ctx context.Context) HTTPIngressPathArrayOutput {
	return o
}

func (o HTTPIngressPathArrayOutput) Index(i pulumi.IntInput) HTTPIngressPathOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) HTTPIngressPath {
		return vs[0].([]HTTPIngressPath)[vs[1].(int)]
	}).(HTTPIngressPathOutput)
}

// HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
type HTTPIngressRuleValue struct {
	// A collection of paths that map requests to backends.
	Paths []HTTPIngressPath `pulumi:"paths"`
}

type HTTPIngressRuleValueInput interface {
	pulumi.Input

	ToHTTPIngressRuleValueOutput() HTTPIngressRuleValueOutput
	ToHTTPIngressRuleValueOutputWithContext(context.Context) HTTPIngressRuleValueOutput
}

// HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
type HTTPIngressRuleValueArgs struct {
	// A collection of paths that map requests to backends.
	Paths HTTPIngressPathArrayInput `pulumi:"paths"`
}

func (HTTPIngressRuleValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HTTPIngressRuleValue)(nil)).Elem()
}

func (i HTTPIngressRuleValueArgs) ToHTTPIngressRuleValueOutput() HTTPIngressRuleValueOutput {
	return i.ToHTTPIngressRuleValueOutputWithContext(context.Background())
}

func (i HTTPIngressRuleValueArgs) ToHTTPIngressRuleValueOutputWithContext(ctx context.Context) HTTPIngressRuleValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HTTPIngressRuleValueOutput)
}

func (i HTTPIngressRuleValueArgs) ToHTTPIngressRuleValuePtrOutput() HTTPIngressRuleValuePtrOutput {
	return i.ToHTTPIngressRuleValuePtrOutputWithContext(context.Background())
}

func (i HTTPIngressRuleValueArgs) ToHTTPIngressRuleValuePtrOutputWithContext(ctx context.Context) HTTPIngressRuleValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HTTPIngressRuleValueOutput).ToHTTPIngressRuleValuePtrOutputWithContext(ctx)
}

type HTTPIngressRuleValuePtrInput interface {
	pulumi.Input

	ToHTTPIngressRuleValuePtrOutput() HTTPIngressRuleValuePtrOutput
	ToHTTPIngressRuleValuePtrOutputWithContext(context.Context) HTTPIngressRuleValuePtrOutput
}

type httpingressRuleValuePtrType HTTPIngressRuleValueArgs

func HTTPIngressRuleValuePtr(v *HTTPIngressRuleValueArgs) HTTPIngressRuleValuePtrInput {	return (*httpingressRuleValuePtrType)(v)
}

func (*httpingressRuleValuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HTTPIngressRuleValue)(nil)).Elem()
}

func (i *httpingressRuleValuePtrType) ToHTTPIngressRuleValuePtrOutput() HTTPIngressRuleValuePtrOutput {
	return i.ToHTTPIngressRuleValuePtrOutputWithContext(context.Background())
}

func (i *httpingressRuleValuePtrType) ToHTTPIngressRuleValuePtrOutputWithContext(ctx context.Context) HTTPIngressRuleValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HTTPIngressRuleValuePtrOutput)
}

// HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
type HTTPIngressRuleValueOutput struct { *pulumi.OutputState }

func (HTTPIngressRuleValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HTTPIngressRuleValue)(nil)).Elem()
}

func (o HTTPIngressRuleValueOutput) ToHTTPIngressRuleValueOutput() HTTPIngressRuleValueOutput {
	return o
}

func (o HTTPIngressRuleValueOutput) ToHTTPIngressRuleValueOutputWithContext(ctx context.Context) HTTPIngressRuleValueOutput {
	return o
}

func (o HTTPIngressRuleValueOutput) ToHTTPIngressRuleValuePtrOutput() HTTPIngressRuleValuePtrOutput {
	return o.ToHTTPIngressRuleValuePtrOutputWithContext(context.Background())
}

func (o HTTPIngressRuleValueOutput) ToHTTPIngressRuleValuePtrOutputWithContext(ctx context.Context) HTTPIngressRuleValuePtrOutput {
	return o.ApplyT(func(v HTTPIngressRuleValue) *HTTPIngressRuleValue {
		return &v
	}).(HTTPIngressRuleValuePtrOutput)
}
// A collection of paths that map requests to backends.
func (o HTTPIngressRuleValueOutput) Paths() HTTPIngressPathArrayOutput {
	return o.ApplyT(func (v HTTPIngressRuleValue) []HTTPIngressPath { return v.Paths }).(HTTPIngressPathArrayOutput)
}

type HTTPIngressRuleValuePtrOutput struct { *pulumi.OutputState}

func (HTTPIngressRuleValuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HTTPIngressRuleValue)(nil)).Elem()
}

func (o HTTPIngressRuleValuePtrOutput) ToHTTPIngressRuleValuePtrOutput() HTTPIngressRuleValuePtrOutput {
	return o
}

func (o HTTPIngressRuleValuePtrOutput) ToHTTPIngressRuleValuePtrOutputWithContext(ctx context.Context) HTTPIngressRuleValuePtrOutput {
	return o
}

func (o HTTPIngressRuleValuePtrOutput) Elem() HTTPIngressRuleValueOutput {
	return o.ApplyT(func (v *HTTPIngressRuleValue) HTTPIngressRuleValue { return *v }).(HTTPIngressRuleValueOutput)
}

// A collection of paths that map requests to backends.
func (o HTTPIngressRuleValuePtrOutput) Paths() HTTPIngressPathArrayOutput {
	return o.ApplyT(func (v HTTPIngressRuleValue) []HTTPIngressPath { return v.Paths }).(HTTPIngressPathArrayOutput)
}

// IngressBackend describes all endpoints for a given service and port.
type IngressBackend struct {
	// Specifies the name of the referenced service.
	ServiceName *string `pulumi:"serviceName"`
	// Specifies the port of the referenced service.
	ServicePort interface{} `pulumi:"servicePort"`
}

type IngressBackendInput interface {
	pulumi.Input

	ToIngressBackendOutput() IngressBackendOutput
	ToIngressBackendOutputWithContext(context.Context) IngressBackendOutput
}

// IngressBackend describes all endpoints for a given service and port.
type IngressBackendArgs struct {
	// Specifies the name of the referenced service.
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
	// Specifies the port of the referenced service.
	ServicePort pulumi.Input `pulumi:"servicePort"`
}

func (IngressBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressBackend)(nil)).Elem()
}

func (i IngressBackendArgs) ToIngressBackendOutput() IngressBackendOutput {
	return i.ToIngressBackendOutputWithContext(context.Background())
}

func (i IngressBackendArgs) ToIngressBackendOutputWithContext(ctx context.Context) IngressBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressBackendOutput)
}

func (i IngressBackendArgs) ToIngressBackendPtrOutput() IngressBackendPtrOutput {
	return i.ToIngressBackendPtrOutputWithContext(context.Background())
}

func (i IngressBackendArgs) ToIngressBackendPtrOutputWithContext(ctx context.Context) IngressBackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressBackendOutput).ToIngressBackendPtrOutputWithContext(ctx)
}

type IngressBackendPtrInput interface {
	pulumi.Input

	ToIngressBackendPtrOutput() IngressBackendPtrOutput
	ToIngressBackendPtrOutputWithContext(context.Context) IngressBackendPtrOutput
}

type ingressBackendPtrType IngressBackendArgs

func IngressBackendPtr(v *IngressBackendArgs) IngressBackendPtrInput {	return (*ingressBackendPtrType)(v)
}

func (*ingressBackendPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressBackend)(nil)).Elem()
}

func (i *ingressBackendPtrType) ToIngressBackendPtrOutput() IngressBackendPtrOutput {
	return i.ToIngressBackendPtrOutputWithContext(context.Background())
}

func (i *ingressBackendPtrType) ToIngressBackendPtrOutputWithContext(ctx context.Context) IngressBackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressBackendPtrOutput)
}

// IngressBackend describes all endpoints for a given service and port.
type IngressBackendOutput struct { *pulumi.OutputState }

func (IngressBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressBackend)(nil)).Elem()
}

func (o IngressBackendOutput) ToIngressBackendOutput() IngressBackendOutput {
	return o
}

func (o IngressBackendOutput) ToIngressBackendOutputWithContext(ctx context.Context) IngressBackendOutput {
	return o
}

func (o IngressBackendOutput) ToIngressBackendPtrOutput() IngressBackendPtrOutput {
	return o.ToIngressBackendPtrOutputWithContext(context.Background())
}

func (o IngressBackendOutput) ToIngressBackendPtrOutputWithContext(ctx context.Context) IngressBackendPtrOutput {
	return o.ApplyT(func(v IngressBackend) *IngressBackend {
		return &v
	}).(IngressBackendPtrOutput)
}
// Specifies the name of the referenced service.
func (o IngressBackendOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v IngressBackend) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

// Specifies the port of the referenced service.
func (o IngressBackendOutput) ServicePort() pulumi.AnyOutput {
	return o.ApplyT(func (v IngressBackend) interface{} { return v.ServicePort }).(pulumi.AnyOutput)
}

type IngressBackendPtrOutput struct { *pulumi.OutputState}

func (IngressBackendPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressBackend)(nil)).Elem()
}

func (o IngressBackendPtrOutput) ToIngressBackendPtrOutput() IngressBackendPtrOutput {
	return o
}

func (o IngressBackendPtrOutput) ToIngressBackendPtrOutputWithContext(ctx context.Context) IngressBackendPtrOutput {
	return o
}

func (o IngressBackendPtrOutput) Elem() IngressBackendOutput {
	return o.ApplyT(func (v *IngressBackend) IngressBackend { return *v }).(IngressBackendOutput)
}

// Specifies the name of the referenced service.
func (o IngressBackendPtrOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v IngressBackend) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

// Specifies the port of the referenced service.
func (o IngressBackendPtrOutput) ServicePort() pulumi.AnyOutput {
	return o.ApplyT(func (v IngressBackend) interface{} { return v.ServicePort }).(pulumi.AnyOutput)
}

// IngressRule represents the rules mapping the paths under a specified host to the related backend services. Incoming requests are first evaluated for a host match, then routed to the backend associated with the matching IngressRuleValue.
type IngressRule struct {
	// Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the "host" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the
	// 	  IP in the Spec of the parent Ingress.
	// 2. The `:` delimiter is not respected because ports are not allowed.
	// 	  Currently the port of an Ingress is implicitly :80 for http and
	// 	  :443 for https.
	// Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.
	Host *string `pulumi:"host"`
	Http *HTTPIngressRuleValue `pulumi:"http"`
}

type IngressRuleInput interface {
	pulumi.Input

	ToIngressRuleOutput() IngressRuleOutput
	ToIngressRuleOutputWithContext(context.Context) IngressRuleOutput
}

// IngressRule represents the rules mapping the paths under a specified host to the related backend services. Incoming requests are first evaluated for a host match, then routed to the backend associated with the matching IngressRuleValue.
type IngressRuleArgs struct {
	// Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the "host" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the
	// 	  IP in the Spec of the parent Ingress.
	// 2. The `:` delimiter is not respected because ports are not allowed.
	// 	  Currently the port of an Ingress is implicitly :80 for http and
	// 	  :443 for https.
	// Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.
	Host pulumi.StringPtrInput `pulumi:"host"`
	Http HTTPIngressRuleValuePtrInput `pulumi:"http"`
}

func (IngressRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressRule)(nil)).Elem()
}

func (i IngressRuleArgs) ToIngressRuleOutput() IngressRuleOutput {
	return i.ToIngressRuleOutputWithContext(context.Background())
}

func (i IngressRuleArgs) ToIngressRuleOutputWithContext(ctx context.Context) IngressRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressRuleOutput)
}

type IngressRuleArrayInput interface {
	pulumi.Input

	ToIngressRuleArrayOutput() IngressRuleArrayOutput
	ToIngressRuleArrayOutputWithContext(context.Context) IngressRuleArrayOutput
}

type IngressRuleArray []IngressRuleInput

func (IngressRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngressRule)(nil)).Elem()
}

func (i IngressRuleArray) ToIngressRuleArrayOutput() IngressRuleArrayOutput {
	return i.ToIngressRuleArrayOutputWithContext(context.Background())
}

func (i IngressRuleArray) ToIngressRuleArrayOutputWithContext(ctx context.Context) IngressRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressRuleArrayOutput)
}

// IngressRule represents the rules mapping the paths under a specified host to the related backend services. Incoming requests are first evaluated for a host match, then routed to the backend associated with the matching IngressRuleValue.
type IngressRuleOutput struct { *pulumi.OutputState }

func (IngressRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressRule)(nil)).Elem()
}

func (o IngressRuleOutput) ToIngressRuleOutput() IngressRuleOutput {
	return o
}

func (o IngressRuleOutput) ToIngressRuleOutputWithContext(ctx context.Context) IngressRuleOutput {
	return o
}

// Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the "host" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the
// 	  IP in the Spec of the parent Ingress.
// 2. The `:` delimiter is not respected because ports are not allowed.
// 	  Currently the port of an Ingress is implicitly :80 for http and
// 	  :443 for https.
// Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.
func (o IngressRuleOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func (v IngressRule) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o IngressRuleOutput) Http() HTTPIngressRuleValuePtrOutput {
	return o.ApplyT(func (v IngressRule) *HTTPIngressRuleValue { return v.Http }).(HTTPIngressRuleValuePtrOutput)
}

type IngressRuleArrayOutput struct { *pulumi.OutputState}

func (IngressRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngressRule)(nil)).Elem()
}

func (o IngressRuleArrayOutput) ToIngressRuleArrayOutput() IngressRuleArrayOutput {
	return o
}

func (o IngressRuleArrayOutput) ToIngressRuleArrayOutputWithContext(ctx context.Context) IngressRuleArrayOutput {
	return o
}

func (o IngressRuleArrayOutput) Index(i pulumi.IntInput) IngressRuleOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) IngressRule {
		return vs[0].([]IngressRule)[vs[1].(int)]
	}).(IngressRuleOutput)
}

// IngressSpec describes the Ingress the user wishes to exist.
type IngressSpec struct {
	// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
	Backend *IngressBackend `pulumi:"backend"`
	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
	Rules []IngressRule `pulumi:"rules"`
	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
	Tls []IngressTLS `pulumi:"tls"`
}

type IngressSpecInput interface {
	pulumi.Input

	ToIngressSpecOutput() IngressSpecOutput
	ToIngressSpecOutputWithContext(context.Context) IngressSpecOutput
}

// IngressSpec describes the Ingress the user wishes to exist.
type IngressSpecArgs struct {
	// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
	Backend IngressBackendPtrInput `pulumi:"backend"`
	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
	Rules IngressRuleArrayInput `pulumi:"rules"`
	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
	Tls IngressTLSArrayInput `pulumi:"tls"`
}

func (IngressSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressSpec)(nil)).Elem()
}

func (i IngressSpecArgs) ToIngressSpecOutput() IngressSpecOutput {
	return i.ToIngressSpecOutputWithContext(context.Background())
}

func (i IngressSpecArgs) ToIngressSpecOutputWithContext(ctx context.Context) IngressSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressSpecOutput)
}

func (i IngressSpecArgs) ToIngressSpecPtrOutput() IngressSpecPtrOutput {
	return i.ToIngressSpecPtrOutputWithContext(context.Background())
}

func (i IngressSpecArgs) ToIngressSpecPtrOutputWithContext(ctx context.Context) IngressSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressSpecOutput).ToIngressSpecPtrOutputWithContext(ctx)
}

type IngressSpecPtrInput interface {
	pulumi.Input

	ToIngressSpecPtrOutput() IngressSpecPtrOutput
	ToIngressSpecPtrOutputWithContext(context.Context) IngressSpecPtrOutput
}

type ingressSpecPtrType IngressSpecArgs

func IngressSpecPtr(v *IngressSpecArgs) IngressSpecPtrInput {	return (*ingressSpecPtrType)(v)
}

func (*ingressSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressSpec)(nil)).Elem()
}

func (i *ingressSpecPtrType) ToIngressSpecPtrOutput() IngressSpecPtrOutput {
	return i.ToIngressSpecPtrOutputWithContext(context.Background())
}

func (i *ingressSpecPtrType) ToIngressSpecPtrOutputWithContext(ctx context.Context) IngressSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressSpecPtrOutput)
}

// IngressSpec describes the Ingress the user wishes to exist.
type IngressSpecOutput struct { *pulumi.OutputState }

func (IngressSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressSpec)(nil)).Elem()
}

func (o IngressSpecOutput) ToIngressSpecOutput() IngressSpecOutput {
	return o
}

func (o IngressSpecOutput) ToIngressSpecOutputWithContext(ctx context.Context) IngressSpecOutput {
	return o
}

func (o IngressSpecOutput) ToIngressSpecPtrOutput() IngressSpecPtrOutput {
	return o.ToIngressSpecPtrOutputWithContext(context.Background())
}

func (o IngressSpecOutput) ToIngressSpecPtrOutputWithContext(ctx context.Context) IngressSpecPtrOutput {
	return o.ApplyT(func(v IngressSpec) *IngressSpec {
		return &v
	}).(IngressSpecPtrOutput)
}
// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
func (o IngressSpecOutput) Backend() IngressBackendPtrOutput {
	return o.ApplyT(func (v IngressSpec) *IngressBackend { return v.Backend }).(IngressBackendPtrOutput)
}

// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
func (o IngressSpecOutput) Rules() IngressRuleArrayOutput {
	return o.ApplyT(func (v IngressSpec) []IngressRule { return v.Rules }).(IngressRuleArrayOutput)
}

// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
func (o IngressSpecOutput) Tls() IngressTLSArrayOutput {
	return o.ApplyT(func (v IngressSpec) []IngressTLS { return v.Tls }).(IngressTLSArrayOutput)
}

type IngressSpecPtrOutput struct { *pulumi.OutputState}

func (IngressSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressSpec)(nil)).Elem()
}

func (o IngressSpecPtrOutput) ToIngressSpecPtrOutput() IngressSpecPtrOutput {
	return o
}

func (o IngressSpecPtrOutput) ToIngressSpecPtrOutputWithContext(ctx context.Context) IngressSpecPtrOutput {
	return o
}

func (o IngressSpecPtrOutput) Elem() IngressSpecOutput {
	return o.ApplyT(func (v *IngressSpec) IngressSpec { return *v }).(IngressSpecOutput)
}

// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
func (o IngressSpecPtrOutput) Backend() IngressBackendPtrOutput {
	return o.ApplyT(func (v IngressSpec) *IngressBackend { return v.Backend }).(IngressBackendPtrOutput)
}

// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
func (o IngressSpecPtrOutput) Rules() IngressRuleArrayOutput {
	return o.ApplyT(func (v IngressSpec) []IngressRule { return v.Rules }).(IngressRuleArrayOutput)
}

// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
func (o IngressSpecPtrOutput) Tls() IngressTLSArrayOutput {
	return o.ApplyT(func (v IngressSpec) []IngressTLS { return v.Tls }).(IngressTLSArrayOutput)
}

// IngressStatus describe the current state of the Ingress.
type IngressStatus struct {
	// LoadBalancer contains the current status of the load-balancer.
	LoadBalancer *corev1.LoadBalancerStatus `pulumi:"loadBalancer"`
}

type IngressStatusInput interface {
	pulumi.Input

	ToIngressStatusOutput() IngressStatusOutput
	ToIngressStatusOutputWithContext(context.Context) IngressStatusOutput
}

// IngressStatus describe the current state of the Ingress.
type IngressStatusArgs struct {
	// LoadBalancer contains the current status of the load-balancer.
	LoadBalancer corev1.LoadBalancerStatusPtrInput `pulumi:"loadBalancer"`
}

func (IngressStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressStatus)(nil)).Elem()
}

func (i IngressStatusArgs) ToIngressStatusOutput() IngressStatusOutput {
	return i.ToIngressStatusOutputWithContext(context.Background())
}

func (i IngressStatusArgs) ToIngressStatusOutputWithContext(ctx context.Context) IngressStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressStatusOutput)
}

func (i IngressStatusArgs) ToIngressStatusPtrOutput() IngressStatusPtrOutput {
	return i.ToIngressStatusPtrOutputWithContext(context.Background())
}

func (i IngressStatusArgs) ToIngressStatusPtrOutputWithContext(ctx context.Context) IngressStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressStatusOutput).ToIngressStatusPtrOutputWithContext(ctx)
}

type IngressStatusPtrInput interface {
	pulumi.Input

	ToIngressStatusPtrOutput() IngressStatusPtrOutput
	ToIngressStatusPtrOutputWithContext(context.Context) IngressStatusPtrOutput
}

type ingressStatusPtrType IngressStatusArgs

func IngressStatusPtr(v *IngressStatusArgs) IngressStatusPtrInput {	return (*ingressStatusPtrType)(v)
}

func (*ingressStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressStatus)(nil)).Elem()
}

func (i *ingressStatusPtrType) ToIngressStatusPtrOutput() IngressStatusPtrOutput {
	return i.ToIngressStatusPtrOutputWithContext(context.Background())
}

func (i *ingressStatusPtrType) ToIngressStatusPtrOutputWithContext(ctx context.Context) IngressStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressStatusPtrOutput)
}

// IngressStatus describe the current state of the Ingress.
type IngressStatusOutput struct { *pulumi.OutputState }

func (IngressStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressStatus)(nil)).Elem()
}

func (o IngressStatusOutput) ToIngressStatusOutput() IngressStatusOutput {
	return o
}

func (o IngressStatusOutput) ToIngressStatusOutputWithContext(ctx context.Context) IngressStatusOutput {
	return o
}

func (o IngressStatusOutput) ToIngressStatusPtrOutput() IngressStatusPtrOutput {
	return o.ToIngressStatusPtrOutputWithContext(context.Background())
}

func (o IngressStatusOutput) ToIngressStatusPtrOutputWithContext(ctx context.Context) IngressStatusPtrOutput {
	return o.ApplyT(func(v IngressStatus) *IngressStatus {
		return &v
	}).(IngressStatusPtrOutput)
}
// LoadBalancer contains the current status of the load-balancer.
func (o IngressStatusOutput) LoadBalancer() corev1.LoadBalancerStatusPtrOutput {
	return o.ApplyT(func (v IngressStatus) *corev1.LoadBalancerStatus { return v.LoadBalancer }).(corev1.LoadBalancerStatusPtrOutput)
}

type IngressStatusPtrOutput struct { *pulumi.OutputState}

func (IngressStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressStatus)(nil)).Elem()
}

func (o IngressStatusPtrOutput) ToIngressStatusPtrOutput() IngressStatusPtrOutput {
	return o
}

func (o IngressStatusPtrOutput) ToIngressStatusPtrOutputWithContext(ctx context.Context) IngressStatusPtrOutput {
	return o
}

func (o IngressStatusPtrOutput) Elem() IngressStatusOutput {
	return o.ApplyT(func (v *IngressStatus) IngressStatus { return *v }).(IngressStatusOutput)
}

// LoadBalancer contains the current status of the load-balancer.
func (o IngressStatusPtrOutput) LoadBalancer() corev1.LoadBalancerStatusPtrOutput {
	return o.ApplyT(func (v IngressStatus) *corev1.LoadBalancerStatus { return v.LoadBalancer }).(corev1.LoadBalancerStatusPtrOutput)
}

// IngressTLS describes the transport layer security associated with an Ingress.
type IngressTLS struct {
	// Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Hosts []string `pulumi:"hosts"`
	// SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
	SecretName *string `pulumi:"secretName"`
}

type IngressTLSInput interface {
	pulumi.Input

	ToIngressTLSOutput() IngressTLSOutput
	ToIngressTLSOutputWithContext(context.Context) IngressTLSOutput
}

// IngressTLS describes the transport layer security associated with an Ingress.
type IngressTLSArgs struct {
	// Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Hosts pulumi.StringArrayInput `pulumi:"hosts"`
	// SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
	SecretName pulumi.StringPtrInput `pulumi:"secretName"`
}

func (IngressTLSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressTLS)(nil)).Elem()
}

func (i IngressTLSArgs) ToIngressTLSOutput() IngressTLSOutput {
	return i.ToIngressTLSOutputWithContext(context.Background())
}

func (i IngressTLSArgs) ToIngressTLSOutputWithContext(ctx context.Context) IngressTLSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressTLSOutput)
}

type IngressTLSArrayInput interface {
	pulumi.Input

	ToIngressTLSArrayOutput() IngressTLSArrayOutput
	ToIngressTLSArrayOutputWithContext(context.Context) IngressTLSArrayOutput
}

type IngressTLSArray []IngressTLSInput

func (IngressTLSArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngressTLS)(nil)).Elem()
}

func (i IngressTLSArray) ToIngressTLSArrayOutput() IngressTLSArrayOutput {
	return i.ToIngressTLSArrayOutputWithContext(context.Background())
}

func (i IngressTLSArray) ToIngressTLSArrayOutputWithContext(ctx context.Context) IngressTLSArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressTLSArrayOutput)
}

// IngressTLS describes the transport layer security associated with an Ingress.
type IngressTLSOutput struct { *pulumi.OutputState }

func (IngressTLSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressTLS)(nil)).Elem()
}

func (o IngressTLSOutput) ToIngressTLSOutput() IngressTLSOutput {
	return o
}

func (o IngressTLSOutput) ToIngressTLSOutputWithContext(ctx context.Context) IngressTLSOutput {
	return o
}

// Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
func (o IngressTLSOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func (v IngressTLS) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

// SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
func (o IngressTLSOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v IngressTLS) *string { return v.SecretName }).(pulumi.StringPtrOutput)
}

type IngressTLSArrayOutput struct { *pulumi.OutputState}

func (IngressTLSArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngressTLS)(nil)).Elem()
}

func (o IngressTLSArrayOutput) ToIngressTLSArrayOutput() IngressTLSArrayOutput {
	return o
}

func (o IngressTLSArrayOutput) ToIngressTLSArrayOutputWithContext(ctx context.Context) IngressTLSArrayOutput {
	return o
}

func (o IngressTLSArrayOutput) Index(i pulumi.IntInput) IngressTLSOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) IngressTLS {
		return vs[0].([]IngressTLS)[vs[1].(int)]
	}).(IngressTLSOutput)
}

func init() {
	pulumi.RegisterOutputType(HTTPIngressPathOutput{})
	pulumi.RegisterOutputType(HTTPIngressPathArrayOutput{})
	pulumi.RegisterOutputType(HTTPIngressRuleValueOutput{})
	pulumi.RegisterOutputType(HTTPIngressRuleValuePtrOutput{})
	pulumi.RegisterOutputType(IngressBackendOutput{})
	pulumi.RegisterOutputType(IngressBackendPtrOutput{})
	pulumi.RegisterOutputType(IngressRuleOutput{})
	pulumi.RegisterOutputType(IngressRuleArrayOutput{})
	pulumi.RegisterOutputType(IngressSpecOutput{})
	pulumi.RegisterOutputType(IngressSpecPtrOutput{})
	pulumi.RegisterOutputType(IngressStatusOutput{})
	pulumi.RegisterOutputType(IngressStatusPtrOutput{})
	pulumi.RegisterOutputType(IngressTLSOutput{})
	pulumi.RegisterOutputType(IngressTLSArrayOutput{})
}
