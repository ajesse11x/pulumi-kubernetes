// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// AuditSinkSpec holds the spec for the audit sink
type AuditSinkSpec struct {
	// Policy defines the policy for selecting which events should be sent to the webhook required
	Policy *Policy `pulumi:"policy"`
	// Webhook to send events required
	Webhook *Webhook `pulumi:"webhook"`
}

type AuditSinkSpecInput interface {
	pulumi.Input

	ToAuditSinkSpecOutput() AuditSinkSpecOutput
	ToAuditSinkSpecOutputWithContext(context.Context) AuditSinkSpecOutput
}

// AuditSinkSpec holds the spec for the audit sink
type AuditSinkSpecArgs struct {
	// Policy defines the policy for selecting which events should be sent to the webhook required
	Policy PolicyPtrInput `pulumi:"policy"`
	// Webhook to send events required
	Webhook WebhookPtrInput `pulumi:"webhook"`
}

func (AuditSinkSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditSinkSpec)(nil)).Elem()
}

func (i AuditSinkSpecArgs) ToAuditSinkSpecOutput() AuditSinkSpecOutput {
	return i.ToAuditSinkSpecOutputWithContext(context.Background())
}

func (i AuditSinkSpecArgs) ToAuditSinkSpecOutputWithContext(ctx context.Context) AuditSinkSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditSinkSpecOutput)
}

func (i AuditSinkSpecArgs) ToAuditSinkSpecPtrOutput() AuditSinkSpecPtrOutput {
	return i.ToAuditSinkSpecPtrOutputWithContext(context.Background())
}

func (i AuditSinkSpecArgs) ToAuditSinkSpecPtrOutputWithContext(ctx context.Context) AuditSinkSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditSinkSpecOutput).ToAuditSinkSpecPtrOutputWithContext(ctx)
}

type AuditSinkSpecPtrInput interface {
	pulumi.Input

	ToAuditSinkSpecPtrOutput() AuditSinkSpecPtrOutput
	ToAuditSinkSpecPtrOutputWithContext(context.Context) AuditSinkSpecPtrOutput
}

type auditSinkSpecPtrType AuditSinkSpecArgs

func AuditSinkSpecPtr(v *AuditSinkSpecArgs) AuditSinkSpecPtrInput {	return (*auditSinkSpecPtrType)(v)
}

func (*auditSinkSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditSinkSpec)(nil)).Elem()
}

func (i *auditSinkSpecPtrType) ToAuditSinkSpecPtrOutput() AuditSinkSpecPtrOutput {
	return i.ToAuditSinkSpecPtrOutputWithContext(context.Background())
}

func (i *auditSinkSpecPtrType) ToAuditSinkSpecPtrOutputWithContext(ctx context.Context) AuditSinkSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditSinkSpecPtrOutput)
}

// AuditSinkSpec holds the spec for the audit sink
type AuditSinkSpecOutput struct { *pulumi.OutputState }

func (AuditSinkSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditSinkSpec)(nil)).Elem()
}

func (o AuditSinkSpecOutput) ToAuditSinkSpecOutput() AuditSinkSpecOutput {
	return o
}

func (o AuditSinkSpecOutput) ToAuditSinkSpecOutputWithContext(ctx context.Context) AuditSinkSpecOutput {
	return o
}

func (o AuditSinkSpecOutput) ToAuditSinkSpecPtrOutput() AuditSinkSpecPtrOutput {
	return o.ToAuditSinkSpecPtrOutputWithContext(context.Background())
}

func (o AuditSinkSpecOutput) ToAuditSinkSpecPtrOutputWithContext(ctx context.Context) AuditSinkSpecPtrOutput {
	return o.ApplyT(func(v AuditSinkSpec) *AuditSinkSpec {
		return &v
	}).(AuditSinkSpecPtrOutput)
}
// Policy defines the policy for selecting which events should be sent to the webhook required
func (o AuditSinkSpecOutput) Policy() PolicyPtrOutput {
	return o.ApplyT(func (v AuditSinkSpec) *Policy { return v.Policy }).(PolicyPtrOutput)
}

// Webhook to send events required
func (o AuditSinkSpecOutput) Webhook() WebhookPtrOutput {
	return o.ApplyT(func (v AuditSinkSpec) *Webhook { return v.Webhook }).(WebhookPtrOutput)
}

type AuditSinkSpecPtrOutput struct { *pulumi.OutputState}

func (AuditSinkSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditSinkSpec)(nil)).Elem()
}

func (o AuditSinkSpecPtrOutput) ToAuditSinkSpecPtrOutput() AuditSinkSpecPtrOutput {
	return o
}

func (o AuditSinkSpecPtrOutput) ToAuditSinkSpecPtrOutputWithContext(ctx context.Context) AuditSinkSpecPtrOutput {
	return o
}

func (o AuditSinkSpecPtrOutput) Elem() AuditSinkSpecOutput {
	return o.ApplyT(func (v *AuditSinkSpec) AuditSinkSpec { return *v }).(AuditSinkSpecOutput)
}

// Policy defines the policy for selecting which events should be sent to the webhook required
func (o AuditSinkSpecPtrOutput) Policy() PolicyPtrOutput {
	return o.ApplyT(func (v AuditSinkSpec) *Policy { return v.Policy }).(PolicyPtrOutput)
}

// Webhook to send events required
func (o AuditSinkSpecPtrOutput) Webhook() WebhookPtrOutput {
	return o.ApplyT(func (v AuditSinkSpec) *Webhook { return v.Webhook }).(WebhookPtrOutput)
}

// Policy defines the configuration of how audit events are logged
type Policy struct {
	// The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
	Level *string `pulumi:"level"`
	// Stages is a list of stages for which events are created.
	Stages []string `pulumi:"stages"`
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(context.Context) PolicyOutput
}

// Policy defines the configuration of how audit events are logged
type PolicyArgs struct {
	// The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
	Level pulumi.StringPtrInput `pulumi:"level"`
	// Stages is a list of stages for which events are created.
	Stages pulumi.StringArrayInput `pulumi:"stages"`
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil)).Elem()
}

func (i PolicyArgs) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i PolicyArgs) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

func (i PolicyArgs) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i PolicyArgs) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput).ToPolicyPtrOutputWithContext(ctx)
}

type PolicyPtrInput interface {
	pulumi.Input

	ToPolicyPtrOutput() PolicyPtrOutput
	ToPolicyPtrOutputWithContext(context.Context) PolicyPtrOutput
}

type policyPtrType PolicyArgs

func PolicyPtr(v *PolicyArgs) PolicyPtrInput {	return (*policyPtrType)(v)
}

func (*policyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *policyPtrType) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *policyPtrType) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

// Policy defines the configuration of how audit events are logged
type PolicyOutput struct { *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o.ToPolicyPtrOutputWithContext(context.Background())
}

func (o PolicyOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o.ApplyT(func(v Policy) *Policy {
		return &v
	}).(PolicyPtrOutput)
}
// The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
func (o PolicyOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func (v Policy) *string { return v.Level }).(pulumi.StringPtrOutput)
}

// Stages is a list of stages for which events are created.
func (o PolicyOutput) Stages() pulumi.StringArrayOutput {
	return o.ApplyT(func (v Policy) []string { return v.Stages }).(pulumi.StringArrayOutput)
}

type PolicyPtrOutput struct { *pulumi.OutputState}

func (PolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyPtrOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) Elem() PolicyOutput {
	return o.ApplyT(func (v *Policy) Policy { return *v }).(PolicyOutput)
}

// The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
func (o PolicyPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func (v Policy) *string { return v.Level }).(pulumi.StringPtrOutput)
}

// Stages is a list of stages for which events are created.
func (o PolicyPtrOutput) Stages() pulumi.StringArrayOutput {
	return o.ApplyT(func (v Policy) []string { return v.Stages }).(pulumi.StringArrayOutput)
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
	// `name` is the name of the service. Required
	Name *string `pulumi:"name"`
	// `namespace` is the namespace of the service. Required
	Namespace *string `pulumi:"namespace"`
	// `path` is an optional URL path which will be sent in any request to this service.
	Path *string `pulumi:"path"`
	// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	Port *int `pulumi:"port"`
}

type ServiceReferenceInput interface {
	pulumi.Input

	ToServiceReferenceOutput() ServiceReferenceOutput
	ToServiceReferenceOutputWithContext(context.Context) ServiceReferenceOutput
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReferenceArgs struct {
	// `name` is the name of the service. Required
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `namespace` is the namespace of the service. Required
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// `path` is an optional URL path which will be sent in any request to this service.
	Path pulumi.StringPtrInput `pulumi:"path"`
	// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	Port pulumi.IntPtrInput `pulumi:"port"`
}

func (ServiceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceReference)(nil)).Elem()
}

func (i ServiceReferenceArgs) ToServiceReferenceOutput() ServiceReferenceOutput {
	return i.ToServiceReferenceOutputWithContext(context.Background())
}

func (i ServiceReferenceArgs) ToServiceReferenceOutputWithContext(ctx context.Context) ServiceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceReferenceOutput)
}

func (i ServiceReferenceArgs) ToServiceReferencePtrOutput() ServiceReferencePtrOutput {
	return i.ToServiceReferencePtrOutputWithContext(context.Background())
}

func (i ServiceReferenceArgs) ToServiceReferencePtrOutputWithContext(ctx context.Context) ServiceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceReferenceOutput).ToServiceReferencePtrOutputWithContext(ctx)
}

type ServiceReferencePtrInput interface {
	pulumi.Input

	ToServiceReferencePtrOutput() ServiceReferencePtrOutput
	ToServiceReferencePtrOutputWithContext(context.Context) ServiceReferencePtrOutput
}

type serviceReferencePtrType ServiceReferenceArgs

func ServiceReferencePtr(v *ServiceReferenceArgs) ServiceReferencePtrInput {	return (*serviceReferencePtrType)(v)
}

func (*serviceReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceReference)(nil)).Elem()
}

func (i *serviceReferencePtrType) ToServiceReferencePtrOutput() ServiceReferencePtrOutput {
	return i.ToServiceReferencePtrOutputWithContext(context.Background())
}

func (i *serviceReferencePtrType) ToServiceReferencePtrOutputWithContext(ctx context.Context) ServiceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceReferencePtrOutput)
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReferenceOutput struct { *pulumi.OutputState }

func (ServiceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceReference)(nil)).Elem()
}

func (o ServiceReferenceOutput) ToServiceReferenceOutput() ServiceReferenceOutput {
	return o
}

func (o ServiceReferenceOutput) ToServiceReferenceOutputWithContext(ctx context.Context) ServiceReferenceOutput {
	return o
}

func (o ServiceReferenceOutput) ToServiceReferencePtrOutput() ServiceReferencePtrOutput {
	return o.ToServiceReferencePtrOutputWithContext(context.Background())
}

func (o ServiceReferenceOutput) ToServiceReferencePtrOutputWithContext(ctx context.Context) ServiceReferencePtrOutput {
	return o.ApplyT(func(v ServiceReference) *ServiceReference {
		return &v
	}).(ServiceReferencePtrOutput)
}
// `name` is the name of the service. Required
func (o ServiceReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// `namespace` is the namespace of the service. Required
func (o ServiceReferenceOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceReference) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// `path` is an optional URL path which will be sent in any request to this service.
func (o ServiceReferenceOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceReference) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
func (o ServiceReferenceOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ServiceReference) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type ServiceReferencePtrOutput struct { *pulumi.OutputState}

func (ServiceReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceReference)(nil)).Elem()
}

func (o ServiceReferencePtrOutput) ToServiceReferencePtrOutput() ServiceReferencePtrOutput {
	return o
}

func (o ServiceReferencePtrOutput) ToServiceReferencePtrOutputWithContext(ctx context.Context) ServiceReferencePtrOutput {
	return o
}

func (o ServiceReferencePtrOutput) Elem() ServiceReferenceOutput {
	return o.ApplyT(func (v *ServiceReference) ServiceReference { return *v }).(ServiceReferenceOutput)
}

// `name` is the name of the service. Required
func (o ServiceReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// `namespace` is the namespace of the service. Required
func (o ServiceReferencePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceReference) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// `path` is an optional URL path which will be sent in any request to this service.
func (o ServiceReferencePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceReference) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
func (o ServiceReferencePtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ServiceReference) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// Webhook holds the configuration of the webhook
type Webhook struct {
	// ClientConfig holds the connection parameters for the webhook required
	ClientConfig *WebhookClientConfig `pulumi:"clientConfig"`
	// Throttle holds the options for throttling the webhook
	Throttle *WebhookThrottleConfig `pulumi:"throttle"`
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(context.Context) WebhookOutput
}

// Webhook holds the configuration of the webhook
type WebhookArgs struct {
	// ClientConfig holds the connection parameters for the webhook required
	ClientConfig WebhookClientConfigPtrInput `pulumi:"clientConfig"`
	// Throttle holds the options for throttling the webhook
	Throttle WebhookThrottleConfigPtrInput `pulumi:"throttle"`
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil)).Elem()
}

func (i WebhookArgs) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i WebhookArgs) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

func (i WebhookArgs) ToWebhookPtrOutput() WebhookPtrOutput {
	return i.ToWebhookPtrOutputWithContext(context.Background())
}

func (i WebhookArgs) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput).ToWebhookPtrOutputWithContext(ctx)
}

type WebhookPtrInput interface {
	pulumi.Input

	ToWebhookPtrOutput() WebhookPtrOutput
	ToWebhookPtrOutputWithContext(context.Context) WebhookPtrOutput
}

type webhookPtrType WebhookArgs

func WebhookPtr(v *WebhookArgs) WebhookPtrInput {	return (*webhookPtrType)(v)
}

func (*webhookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (i *webhookPtrType) ToWebhookPtrOutput() WebhookPtrOutput {
	return i.ToWebhookPtrOutputWithContext(context.Background())
}

func (i *webhookPtrType) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPtrOutput)
}

// Webhook holds the configuration of the webhook
type WebhookOutput struct { *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil)).Elem()
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookPtrOutput() WebhookPtrOutput {
	return o.ToWebhookPtrOutputWithContext(context.Background())
}

func (o WebhookOutput) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return o.ApplyT(func(v Webhook) *Webhook {
		return &v
	}).(WebhookPtrOutput)
}
// ClientConfig holds the connection parameters for the webhook required
func (o WebhookOutput) ClientConfig() WebhookClientConfigPtrOutput {
	return o.ApplyT(func (v Webhook) *WebhookClientConfig { return v.ClientConfig }).(WebhookClientConfigPtrOutput)
}

// Throttle holds the options for throttling the webhook
func (o WebhookOutput) Throttle() WebhookThrottleConfigPtrOutput {
	return o.ApplyT(func (v Webhook) *WebhookThrottleConfig { return v.Throttle }).(WebhookThrottleConfigPtrOutput)
}

type WebhookPtrOutput struct { *pulumi.OutputState}

func (WebhookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (o WebhookPtrOutput) ToWebhookPtrOutput() WebhookPtrOutput {
	return o
}

func (o WebhookPtrOutput) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return o
}

func (o WebhookPtrOutput) Elem() WebhookOutput {
	return o.ApplyT(func (v *Webhook) Webhook { return *v }).(WebhookOutput)
}

// ClientConfig holds the connection parameters for the webhook required
func (o WebhookPtrOutput) ClientConfig() WebhookClientConfigPtrOutput {
	return o.ApplyT(func (v Webhook) *WebhookClientConfig { return v.ClientConfig }).(WebhookClientConfigPtrOutput)
}

// Throttle holds the options for throttling the webhook
func (o WebhookPtrOutput) Throttle() WebhookThrottleConfigPtrOutput {
	return o.ApplyT(func (v Webhook) *WebhookThrottleConfig { return v.Throttle }).(WebhookThrottleConfigPtrOutput)
}

// WebhookClientConfig contains the information to make a connection with the webhook
type WebhookClientConfig struct {
	// `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
	CaBundle *string `pulumi:"caBundle"`
	// `service` is a reference to the service for this webhook. Either `service` or `url` must be specified.
	//
	// If the webhook is running within the cluster, then you should use `service`.
	Service *ServiceReference `pulumi:"service"`
	// `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.
	//
	// The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.
	//
	// Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.
	//
	// The scheme must be "https"; the URL must begin with "https://".
	//
	// A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.
	//
	// Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
	Url *string `pulumi:"url"`
}

type WebhookClientConfigInput interface {
	pulumi.Input

	ToWebhookClientConfigOutput() WebhookClientConfigOutput
	ToWebhookClientConfigOutputWithContext(context.Context) WebhookClientConfigOutput
}

// WebhookClientConfig contains the information to make a connection with the webhook
type WebhookClientConfigArgs struct {
	// `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
	CaBundle pulumi.StringPtrInput `pulumi:"caBundle"`
	// `service` is a reference to the service for this webhook. Either `service` or `url` must be specified.
	//
	// If the webhook is running within the cluster, then you should use `service`.
	Service ServiceReferencePtrInput `pulumi:"service"`
	// `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.
	//
	// The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.
	//
	// Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.
	//
	// The scheme must be "https"; the URL must begin with "https://".
	//
	// A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.
	//
	// Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
	Url pulumi.StringPtrInput `pulumi:"url"`
}

func (WebhookClientConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookClientConfig)(nil)).Elem()
}

func (i WebhookClientConfigArgs) ToWebhookClientConfigOutput() WebhookClientConfigOutput {
	return i.ToWebhookClientConfigOutputWithContext(context.Background())
}

func (i WebhookClientConfigArgs) ToWebhookClientConfigOutputWithContext(ctx context.Context) WebhookClientConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookClientConfigOutput)
}

func (i WebhookClientConfigArgs) ToWebhookClientConfigPtrOutput() WebhookClientConfigPtrOutput {
	return i.ToWebhookClientConfigPtrOutputWithContext(context.Background())
}

func (i WebhookClientConfigArgs) ToWebhookClientConfigPtrOutputWithContext(ctx context.Context) WebhookClientConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookClientConfigOutput).ToWebhookClientConfigPtrOutputWithContext(ctx)
}

type WebhookClientConfigPtrInput interface {
	pulumi.Input

	ToWebhookClientConfigPtrOutput() WebhookClientConfigPtrOutput
	ToWebhookClientConfigPtrOutputWithContext(context.Context) WebhookClientConfigPtrOutput
}

type webhookClientConfigPtrType WebhookClientConfigArgs

func WebhookClientConfigPtr(v *WebhookClientConfigArgs) WebhookClientConfigPtrInput {	return (*webhookClientConfigPtrType)(v)
}

func (*webhookClientConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookClientConfig)(nil)).Elem()
}

func (i *webhookClientConfigPtrType) ToWebhookClientConfigPtrOutput() WebhookClientConfigPtrOutput {
	return i.ToWebhookClientConfigPtrOutputWithContext(context.Background())
}

func (i *webhookClientConfigPtrType) ToWebhookClientConfigPtrOutputWithContext(ctx context.Context) WebhookClientConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookClientConfigPtrOutput)
}

// WebhookClientConfig contains the information to make a connection with the webhook
type WebhookClientConfigOutput struct { *pulumi.OutputState }

func (WebhookClientConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookClientConfig)(nil)).Elem()
}

func (o WebhookClientConfigOutput) ToWebhookClientConfigOutput() WebhookClientConfigOutput {
	return o
}

func (o WebhookClientConfigOutput) ToWebhookClientConfigOutputWithContext(ctx context.Context) WebhookClientConfigOutput {
	return o
}

func (o WebhookClientConfigOutput) ToWebhookClientConfigPtrOutput() WebhookClientConfigPtrOutput {
	return o.ToWebhookClientConfigPtrOutputWithContext(context.Background())
}

func (o WebhookClientConfigOutput) ToWebhookClientConfigPtrOutputWithContext(ctx context.Context) WebhookClientConfigPtrOutput {
	return o.ApplyT(func(v WebhookClientConfig) *WebhookClientConfig {
		return &v
	}).(WebhookClientConfigPtrOutput)
}
// `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
func (o WebhookClientConfigOutput) CaBundle() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebhookClientConfig) *string { return v.CaBundle }).(pulumi.StringPtrOutput)
}

// `service` is a reference to the service for this webhook. Either `service` or `url` must be specified.
//
// If the webhook is running within the cluster, then you should use `service`.
func (o WebhookClientConfigOutput) Service() ServiceReferencePtrOutput {
	return o.ApplyT(func (v WebhookClientConfig) *ServiceReference { return v.Service }).(ServiceReferencePtrOutput)
}

// `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.
//
// The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.
//
// Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.
//
// The scheme must be "https"; the URL must begin with "https://".
//
// A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.
//
// Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
func (o WebhookClientConfigOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebhookClientConfig) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type WebhookClientConfigPtrOutput struct { *pulumi.OutputState}

func (WebhookClientConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookClientConfig)(nil)).Elem()
}

func (o WebhookClientConfigPtrOutput) ToWebhookClientConfigPtrOutput() WebhookClientConfigPtrOutput {
	return o
}

func (o WebhookClientConfigPtrOutput) ToWebhookClientConfigPtrOutputWithContext(ctx context.Context) WebhookClientConfigPtrOutput {
	return o
}

func (o WebhookClientConfigPtrOutput) Elem() WebhookClientConfigOutput {
	return o.ApplyT(func (v *WebhookClientConfig) WebhookClientConfig { return *v }).(WebhookClientConfigOutput)
}

// `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
func (o WebhookClientConfigPtrOutput) CaBundle() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebhookClientConfig) *string { return v.CaBundle }).(pulumi.StringPtrOutput)
}

// `service` is a reference to the service for this webhook. Either `service` or `url` must be specified.
//
// If the webhook is running within the cluster, then you should use `service`.
func (o WebhookClientConfigPtrOutput) Service() ServiceReferencePtrOutput {
	return o.ApplyT(func (v WebhookClientConfig) *ServiceReference { return v.Service }).(ServiceReferencePtrOutput)
}

// `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.
//
// The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.
//
// Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.
//
// The scheme must be "https"; the URL must begin with "https://".
//
// A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.
//
// Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
func (o WebhookClientConfigPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebhookClientConfig) *string { return v.Url }).(pulumi.StringPtrOutput)
}

// WebhookThrottleConfig holds the configuration for throttling events
type WebhookThrottleConfig struct {
	// ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
	Burst *int `pulumi:"burst"`
	// ThrottleQPS maximum number of batches per second default 10 QPS
	Qps *int `pulumi:"qps"`
}

type WebhookThrottleConfigInput interface {
	pulumi.Input

	ToWebhookThrottleConfigOutput() WebhookThrottleConfigOutput
	ToWebhookThrottleConfigOutputWithContext(context.Context) WebhookThrottleConfigOutput
}

// WebhookThrottleConfig holds the configuration for throttling events
type WebhookThrottleConfigArgs struct {
	// ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
	Burst pulumi.IntPtrInput `pulumi:"burst"`
	// ThrottleQPS maximum number of batches per second default 10 QPS
	Qps pulumi.IntPtrInput `pulumi:"qps"`
}

func (WebhookThrottleConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookThrottleConfig)(nil)).Elem()
}

func (i WebhookThrottleConfigArgs) ToWebhookThrottleConfigOutput() WebhookThrottleConfigOutput {
	return i.ToWebhookThrottleConfigOutputWithContext(context.Background())
}

func (i WebhookThrottleConfigArgs) ToWebhookThrottleConfigOutputWithContext(ctx context.Context) WebhookThrottleConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookThrottleConfigOutput)
}

func (i WebhookThrottleConfigArgs) ToWebhookThrottleConfigPtrOutput() WebhookThrottleConfigPtrOutput {
	return i.ToWebhookThrottleConfigPtrOutputWithContext(context.Background())
}

func (i WebhookThrottleConfigArgs) ToWebhookThrottleConfigPtrOutputWithContext(ctx context.Context) WebhookThrottleConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookThrottleConfigOutput).ToWebhookThrottleConfigPtrOutputWithContext(ctx)
}

type WebhookThrottleConfigPtrInput interface {
	pulumi.Input

	ToWebhookThrottleConfigPtrOutput() WebhookThrottleConfigPtrOutput
	ToWebhookThrottleConfigPtrOutputWithContext(context.Context) WebhookThrottleConfigPtrOutput
}

type webhookThrottleConfigPtrType WebhookThrottleConfigArgs

func WebhookThrottleConfigPtr(v *WebhookThrottleConfigArgs) WebhookThrottleConfigPtrInput {	return (*webhookThrottleConfigPtrType)(v)
}

func (*webhookThrottleConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookThrottleConfig)(nil)).Elem()
}

func (i *webhookThrottleConfigPtrType) ToWebhookThrottleConfigPtrOutput() WebhookThrottleConfigPtrOutput {
	return i.ToWebhookThrottleConfigPtrOutputWithContext(context.Background())
}

func (i *webhookThrottleConfigPtrType) ToWebhookThrottleConfigPtrOutputWithContext(ctx context.Context) WebhookThrottleConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookThrottleConfigPtrOutput)
}

// WebhookThrottleConfig holds the configuration for throttling events
type WebhookThrottleConfigOutput struct { *pulumi.OutputState }

func (WebhookThrottleConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookThrottleConfig)(nil)).Elem()
}

func (o WebhookThrottleConfigOutput) ToWebhookThrottleConfigOutput() WebhookThrottleConfigOutput {
	return o
}

func (o WebhookThrottleConfigOutput) ToWebhookThrottleConfigOutputWithContext(ctx context.Context) WebhookThrottleConfigOutput {
	return o
}

func (o WebhookThrottleConfigOutput) ToWebhookThrottleConfigPtrOutput() WebhookThrottleConfigPtrOutput {
	return o.ToWebhookThrottleConfigPtrOutputWithContext(context.Background())
}

func (o WebhookThrottleConfigOutput) ToWebhookThrottleConfigPtrOutputWithContext(ctx context.Context) WebhookThrottleConfigPtrOutput {
	return o.ApplyT(func(v WebhookThrottleConfig) *WebhookThrottleConfig {
		return &v
	}).(WebhookThrottleConfigPtrOutput)
}
// ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
func (o WebhookThrottleConfigOutput) Burst() pulumi.IntPtrOutput {
	return o.ApplyT(func (v WebhookThrottleConfig) *int { return v.Burst }).(pulumi.IntPtrOutput)
}

// ThrottleQPS maximum number of batches per second default 10 QPS
func (o WebhookThrottleConfigOutput) Qps() pulumi.IntPtrOutput {
	return o.ApplyT(func (v WebhookThrottleConfig) *int { return v.Qps }).(pulumi.IntPtrOutput)
}

type WebhookThrottleConfigPtrOutput struct { *pulumi.OutputState}

func (WebhookThrottleConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookThrottleConfig)(nil)).Elem()
}

func (o WebhookThrottleConfigPtrOutput) ToWebhookThrottleConfigPtrOutput() WebhookThrottleConfigPtrOutput {
	return o
}

func (o WebhookThrottleConfigPtrOutput) ToWebhookThrottleConfigPtrOutputWithContext(ctx context.Context) WebhookThrottleConfigPtrOutput {
	return o
}

func (o WebhookThrottleConfigPtrOutput) Elem() WebhookThrottleConfigOutput {
	return o.ApplyT(func (v *WebhookThrottleConfig) WebhookThrottleConfig { return *v }).(WebhookThrottleConfigOutput)
}

// ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
func (o WebhookThrottleConfigPtrOutput) Burst() pulumi.IntPtrOutput {
	return o.ApplyT(func (v WebhookThrottleConfig) *int { return v.Burst }).(pulumi.IntPtrOutput)
}

// ThrottleQPS maximum number of batches per second default 10 QPS
func (o WebhookThrottleConfigPtrOutput) Qps() pulumi.IntPtrOutput {
	return o.ApplyT(func (v WebhookThrottleConfig) *int { return v.Qps }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuditSinkSpecOutput{})
	pulumi.RegisterOutputType(AuditSinkSpecPtrOutput{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyPtrOutput{})
	pulumi.RegisterOutputType(ServiceReferenceOutput{})
	pulumi.RegisterOutputType(ServiceReferencePtrOutput{})
	pulumi.RegisterOutputType(WebhookOutput{})
	pulumi.RegisterOutputType(WebhookPtrOutput{})
	pulumi.RegisterOutputType(WebhookClientConfigOutput{})
	pulumi.RegisterOutputType(WebhookClientConfigPtrOutput{})
	pulumi.RegisterOutputType(WebhookThrottleConfigOutput{})
	pulumi.RegisterOutputType(WebhookThrottleConfigPtrOutput{})
}
