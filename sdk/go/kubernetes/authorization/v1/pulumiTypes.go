// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
type NonResourceAttributes struct {
	// Path is the URL path of the request
	Path *string `pulumi:"path"`
	// Verb is the standard HTTP verb
	Verb *string `pulumi:"verb"`
}

type NonResourceAttributesInput interface {
	pulumi.Input

	ToNonResourceAttributesOutput() NonResourceAttributesOutput
	ToNonResourceAttributesOutputWithContext(context.Context) NonResourceAttributesOutput
}

// NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
type NonResourceAttributesArgs struct {
	// Path is the URL path of the request
	Path pulumi.StringPtrInput `pulumi:"path"`
	// Verb is the standard HTTP verb
	Verb pulumi.StringPtrInput `pulumi:"verb"`
}

func (NonResourceAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NonResourceAttributes)(nil)).Elem()
}

func (i NonResourceAttributesArgs) ToNonResourceAttributesOutput() NonResourceAttributesOutput {
	return i.ToNonResourceAttributesOutputWithContext(context.Background())
}

func (i NonResourceAttributesArgs) ToNonResourceAttributesOutputWithContext(ctx context.Context) NonResourceAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonResourceAttributesOutput)
}

// NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
type NonResourceAttributesOutput struct { *pulumi.OutputState }

func (NonResourceAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonResourceAttributes)(nil)).Elem()
}

func (o NonResourceAttributesOutput) ToNonResourceAttributesOutput() NonResourceAttributesOutput {
	return o
}

func (o NonResourceAttributesOutput) ToNonResourceAttributesOutputWithContext(ctx context.Context) NonResourceAttributesOutput {
	return o
}

// Path is the URL path of the request
func (o NonResourceAttributesOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v NonResourceAttributes) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// Verb is the standard HTTP verb
func (o NonResourceAttributesOutput) Verb() pulumi.StringPtrOutput {
	return o.ApplyT(func (v NonResourceAttributes) *string { return v.Verb }).(pulumi.StringPtrOutput)
}

// NonResourceRule holds information that describes a rule for the non-resource
type NonResourceRule struct {
	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path.  "*" means all.
	NonResourceURLs []string `pulumi:"nonResourceURLs"`
	// Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  "*" means all.
	Verbs []string `pulumi:"verbs"`
}

type NonResourceRuleInput interface {
	pulumi.Input

	ToNonResourceRuleOutput() NonResourceRuleOutput
	ToNonResourceRuleOutputWithContext(context.Context) NonResourceRuleOutput
}

// NonResourceRule holds information that describes a rule for the non-resource
type NonResourceRuleArgs struct {
	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path.  "*" means all.
	NonResourceURLs pulumi.StringArrayInput `pulumi:"nonResourceURLs"`
	// Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  "*" means all.
	Verbs pulumi.StringArrayInput `pulumi:"verbs"`
}

func (NonResourceRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NonResourceRule)(nil)).Elem()
}

func (i NonResourceRuleArgs) ToNonResourceRuleOutput() NonResourceRuleOutput {
	return i.ToNonResourceRuleOutputWithContext(context.Background())
}

func (i NonResourceRuleArgs) ToNonResourceRuleOutputWithContext(ctx context.Context) NonResourceRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonResourceRuleOutput)
}

// NonResourceRule holds information that describes a rule for the non-resource
type NonResourceRuleOutput struct { *pulumi.OutputState }

func (NonResourceRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonResourceRule)(nil)).Elem()
}

func (o NonResourceRuleOutput) ToNonResourceRuleOutput() NonResourceRuleOutput {
	return o
}

func (o NonResourceRuleOutput) ToNonResourceRuleOutputWithContext(ctx context.Context) NonResourceRuleOutput {
	return o
}

// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path.  "*" means all.
func (o NonResourceRuleOutput) NonResourceURLs() pulumi.StringArrayOutput {
	return o.ApplyT(func (v NonResourceRule) []string { return v.NonResourceURLs }).(pulumi.StringArrayOutput)
}

// Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  "*" means all.
func (o NonResourceRuleOutput) Verbs() pulumi.StringArrayOutput {
	return o.ApplyT(func (v NonResourceRule) []string { return v.Verbs }).(pulumi.StringArrayOutput)
}

// ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
type ResourceAttributes struct {
	// Group is the API Group of the Resource.  "*" means all.
	Group *string `pulumi:"group"`
	// Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	Name *string `pulumi:"name"`
	// Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Namespace *string `pulumi:"namespace"`
	// Resource is one of the existing resource types.  "*" means all.
	Resource *string `pulumi:"resource"`
	// Subresource is one of the existing resource types.  "" means none.
	Subresource *string `pulumi:"subresource"`
	// Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
	Verb *string `pulumi:"verb"`
	// Version is the API Version of the Resource.  "*" means all.
	Version *string `pulumi:"version"`
}

type ResourceAttributesInput interface {
	pulumi.Input

	ToResourceAttributesOutput() ResourceAttributesOutput
	ToResourceAttributesOutputWithContext(context.Context) ResourceAttributesOutput
}

// ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
type ResourceAttributesArgs struct {
	// Group is the API Group of the Resource.  "*" means all.
	Group pulumi.StringPtrInput `pulumi:"group"`
	// Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Resource is one of the existing resource types.  "*" means all.
	Resource pulumi.StringPtrInput `pulumi:"resource"`
	// Subresource is one of the existing resource types.  "" means none.
	Subresource pulumi.StringPtrInput `pulumi:"subresource"`
	// Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
	Verb pulumi.StringPtrInput `pulumi:"verb"`
	// Version is the API Version of the Resource.  "*" means all.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ResourceAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAttributes)(nil)).Elem()
}

func (i ResourceAttributesArgs) ToResourceAttributesOutput() ResourceAttributesOutput {
	return i.ToResourceAttributesOutputWithContext(context.Background())
}

func (i ResourceAttributesArgs) ToResourceAttributesOutputWithContext(ctx context.Context) ResourceAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAttributesOutput)
}

// ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
type ResourceAttributesOutput struct { *pulumi.OutputState }

func (ResourceAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAttributes)(nil)).Elem()
}

func (o ResourceAttributesOutput) ToResourceAttributesOutput() ResourceAttributesOutput {
	return o
}

func (o ResourceAttributesOutput) ToResourceAttributesOutputWithContext(ctx context.Context) ResourceAttributesOutput {
	return o
}

// Group is the API Group of the Resource.  "*" means all.
func (o ResourceAttributesOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ResourceAttributes) *string { return v.Group }).(pulumi.StringPtrOutput)
}

// Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
func (o ResourceAttributesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ResourceAttributes) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
func (o ResourceAttributesOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ResourceAttributes) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Resource is one of the existing resource types.  "*" means all.
func (o ResourceAttributesOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ResourceAttributes) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

// Subresource is one of the existing resource types.  "" means none.
func (o ResourceAttributesOutput) Subresource() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ResourceAttributes) *string { return v.Subresource }).(pulumi.StringPtrOutput)
}

// Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
func (o ResourceAttributesOutput) Verb() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ResourceAttributes) *string { return v.Verb }).(pulumi.StringPtrOutput)
}

// Version is the API Version of the Resource.  "*" means all.
func (o ResourceAttributesOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ResourceAttributes) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
type ResourceRule struct {
	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.  "*" means all.
	ApiGroups []string `pulumi:"apiGroups"`
	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.  "*" means all.
	ResourceNames []string `pulumi:"resourceNames"`
	// Resources is a list of resources this rule applies to.  "*" means all in the specified apiGroups.
	//  "*/foo" represents the subresource 'foo' for all resources in the specified apiGroups.
	Resources []string `pulumi:"resources"`
	// Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy.  "*" means all.
	Verbs []string `pulumi:"verbs"`
}

type ResourceRuleInput interface {
	pulumi.Input

	ToResourceRuleOutput() ResourceRuleOutput
	ToResourceRuleOutputWithContext(context.Context) ResourceRuleOutput
}

// ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
type ResourceRuleArgs struct {
	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.  "*" means all.
	ApiGroups pulumi.StringArrayInput `pulumi:"apiGroups"`
	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.  "*" means all.
	ResourceNames pulumi.StringArrayInput `pulumi:"resourceNames"`
	// Resources is a list of resources this rule applies to.  "*" means all in the specified apiGroups.
	//  "*/foo" represents the subresource 'foo' for all resources in the specified apiGroups.
	Resources pulumi.StringArrayInput `pulumi:"resources"`
	// Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy.  "*" means all.
	Verbs pulumi.StringArrayInput `pulumi:"verbs"`
}

func (ResourceRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRule)(nil)).Elem()
}

func (i ResourceRuleArgs) ToResourceRuleOutput() ResourceRuleOutput {
	return i.ToResourceRuleOutputWithContext(context.Background())
}

func (i ResourceRuleArgs) ToResourceRuleOutputWithContext(ctx context.Context) ResourceRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRuleOutput)
}

// ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
type ResourceRuleOutput struct { *pulumi.OutputState }

func (ResourceRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRule)(nil)).Elem()
}

func (o ResourceRuleOutput) ToResourceRuleOutput() ResourceRuleOutput {
	return o
}

func (o ResourceRuleOutput) ToResourceRuleOutputWithContext(ctx context.Context) ResourceRuleOutput {
	return o
}

// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.  "*" means all.
func (o ResourceRuleOutput) ApiGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ResourceRule) []string { return v.ApiGroups }).(pulumi.StringArrayOutput)
}

// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.  "*" means all.
func (o ResourceRuleOutput) ResourceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ResourceRule) []string { return v.ResourceNames }).(pulumi.StringArrayOutput)
}

// Resources is a list of resources this rule applies to.  "*" means all in the specified apiGroups.
//  "*/foo" represents the subresource 'foo' for all resources in the specified apiGroups.
func (o ResourceRuleOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ResourceRule) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

// Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy.  "*" means all.
func (o ResourceRuleOutput) Verbs() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ResourceRule) []string { return v.Verbs }).(pulumi.StringArrayOutput)
}

// SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type SelfSubjectAccessReviewSpec struct {
	// NonResourceAttributes describes information for a non-resource access request
	NonResourceAttributes *NonResourceAttributes `pulumi:"nonResourceAttributes"`
	// ResourceAuthorizationAttributes describes information for a resource access request
	ResourceAttributes *ResourceAttributes `pulumi:"resourceAttributes"`
}

type SelfSubjectAccessReviewSpecInput interface {
	pulumi.Input

	ToSelfSubjectAccessReviewSpecOutput() SelfSubjectAccessReviewSpecOutput
	ToSelfSubjectAccessReviewSpecOutputWithContext(context.Context) SelfSubjectAccessReviewSpecOutput
}

// SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type SelfSubjectAccessReviewSpecArgs struct {
	// NonResourceAttributes describes information for a non-resource access request
	NonResourceAttributes NonResourceAttributesPtrInput `pulumi:"nonResourceAttributes"`
	// ResourceAuthorizationAttributes describes information for a resource access request
	ResourceAttributes ResourceAttributesPtrInput `pulumi:"resourceAttributes"`
}

func (SelfSubjectAccessReviewSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSubjectAccessReviewSpec)(nil)).Elem()
}

func (i SelfSubjectAccessReviewSpecArgs) ToSelfSubjectAccessReviewSpecOutput() SelfSubjectAccessReviewSpecOutput {
	return i.ToSelfSubjectAccessReviewSpecOutputWithContext(context.Background())
}

func (i SelfSubjectAccessReviewSpecArgs) ToSelfSubjectAccessReviewSpecOutputWithContext(ctx context.Context) SelfSubjectAccessReviewSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewSpecOutput)
}

// SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type SelfSubjectAccessReviewSpecOutput struct { *pulumi.OutputState }

func (SelfSubjectAccessReviewSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSubjectAccessReviewSpec)(nil)).Elem()
}

func (o SelfSubjectAccessReviewSpecOutput) ToSelfSubjectAccessReviewSpecOutput() SelfSubjectAccessReviewSpecOutput {
	return o
}

func (o SelfSubjectAccessReviewSpecOutput) ToSelfSubjectAccessReviewSpecOutputWithContext(ctx context.Context) SelfSubjectAccessReviewSpecOutput {
	return o
}

// NonResourceAttributes describes information for a non-resource access request
func (o SelfSubjectAccessReviewSpecOutput) NonResourceAttributes() NonResourceAttributesPtrOutput {
	return o.ApplyT(func (v SelfSubjectAccessReviewSpec) *NonResourceAttributes { return v.NonResourceAttributes }).(NonResourceAttributesPtrOutput)
}

// ResourceAuthorizationAttributes describes information for a resource access request
func (o SelfSubjectAccessReviewSpecOutput) ResourceAttributes() ResourceAttributesPtrOutput {
	return o.ApplyT(func (v SelfSubjectAccessReviewSpec) *ResourceAttributes { return v.ResourceAttributes }).(ResourceAttributesPtrOutput)
}

type SelfSubjectRulesReviewSpec struct {
	// Namespace to evaluate rules for. Required.
	Namespace *string `pulumi:"namespace"`
}

type SelfSubjectRulesReviewSpecInput interface {
	pulumi.Input

	ToSelfSubjectRulesReviewSpecOutput() SelfSubjectRulesReviewSpecOutput
	ToSelfSubjectRulesReviewSpecOutputWithContext(context.Context) SelfSubjectRulesReviewSpecOutput
}

type SelfSubjectRulesReviewSpecArgs struct {
	// Namespace to evaluate rules for. Required.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (SelfSubjectRulesReviewSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSubjectRulesReviewSpec)(nil)).Elem()
}

func (i SelfSubjectRulesReviewSpecArgs) ToSelfSubjectRulesReviewSpecOutput() SelfSubjectRulesReviewSpecOutput {
	return i.ToSelfSubjectRulesReviewSpecOutputWithContext(context.Background())
}

func (i SelfSubjectRulesReviewSpecArgs) ToSelfSubjectRulesReviewSpecOutputWithContext(ctx context.Context) SelfSubjectRulesReviewSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewSpecOutput)
}

type SelfSubjectRulesReviewSpecOutput struct { *pulumi.OutputState }

func (SelfSubjectRulesReviewSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSubjectRulesReviewSpec)(nil)).Elem()
}

func (o SelfSubjectRulesReviewSpecOutput) ToSelfSubjectRulesReviewSpecOutput() SelfSubjectRulesReviewSpecOutput {
	return o
}

func (o SelfSubjectRulesReviewSpecOutput) ToSelfSubjectRulesReviewSpecOutputWithContext(ctx context.Context) SelfSubjectRulesReviewSpecOutput {
	return o
}

// Namespace to evaluate rules for. Required.
func (o SelfSubjectRulesReviewSpecOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SelfSubjectRulesReviewSpec) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type SubjectAccessReviewSpec struct {
	// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
	Extra map[string][]string `pulumi:"extra"`
	// Groups is the groups you're testing for.
	Groups []string `pulumi:"groups"`
	// NonResourceAttributes describes information for a non-resource access request
	NonResourceAttributes *NonResourceAttributes `pulumi:"nonResourceAttributes"`
	// ResourceAuthorizationAttributes describes information for a resource access request
	ResourceAttributes *ResourceAttributes `pulumi:"resourceAttributes"`
	// UID information about the requesting user.
	Uid *string `pulumi:"uid"`
	// User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
	User *string `pulumi:"user"`
}

type SubjectAccessReviewSpecInput interface {
	pulumi.Input

	ToSubjectAccessReviewSpecOutput() SubjectAccessReviewSpecOutput
	ToSubjectAccessReviewSpecOutputWithContext(context.Context) SubjectAccessReviewSpecOutput
}

// SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type SubjectAccessReviewSpecArgs struct {
	// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
	Extra pulumi.StringArrayMapInput `pulumi:"extra"`
	// Groups is the groups you're testing for.
	Groups pulumi.StringArrayInput `pulumi:"groups"`
	// NonResourceAttributes describes information for a non-resource access request
	NonResourceAttributes NonResourceAttributesPtrInput `pulumi:"nonResourceAttributes"`
	// ResourceAuthorizationAttributes describes information for a resource access request
	ResourceAttributes ResourceAttributesPtrInput `pulumi:"resourceAttributes"`
	// UID information about the requesting user.
	Uid pulumi.StringPtrInput `pulumi:"uid"`
	// User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
	User pulumi.StringPtrInput `pulumi:"user"`
}

func (SubjectAccessReviewSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubjectAccessReviewSpec)(nil)).Elem()
}

func (i SubjectAccessReviewSpecArgs) ToSubjectAccessReviewSpecOutput() SubjectAccessReviewSpecOutput {
	return i.ToSubjectAccessReviewSpecOutputWithContext(context.Background())
}

func (i SubjectAccessReviewSpecArgs) ToSubjectAccessReviewSpecOutputWithContext(ctx context.Context) SubjectAccessReviewSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubjectAccessReviewSpecOutput)
}

// SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type SubjectAccessReviewSpecOutput struct { *pulumi.OutputState }

func (SubjectAccessReviewSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubjectAccessReviewSpec)(nil)).Elem()
}

func (o SubjectAccessReviewSpecOutput) ToSubjectAccessReviewSpecOutput() SubjectAccessReviewSpecOutput {
	return o
}

func (o SubjectAccessReviewSpecOutput) ToSubjectAccessReviewSpecOutputWithContext(ctx context.Context) SubjectAccessReviewSpecOutput {
	return o
}

// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
func (o SubjectAccessReviewSpecOutput) Extra() pulumi.StringArrayMapOutput {
	return o.ApplyT(func (v SubjectAccessReviewSpec) map[string][]string { return v.Extra }).(pulumi.StringArrayMapOutput)
}

// Groups is the groups you're testing for.
func (o SubjectAccessReviewSpecOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func (v SubjectAccessReviewSpec) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// NonResourceAttributes describes information for a non-resource access request
func (o SubjectAccessReviewSpecOutput) NonResourceAttributes() NonResourceAttributesPtrOutput {
	return o.ApplyT(func (v SubjectAccessReviewSpec) *NonResourceAttributes { return v.NonResourceAttributes }).(NonResourceAttributesPtrOutput)
}

// ResourceAuthorizationAttributes describes information for a resource access request
func (o SubjectAccessReviewSpecOutput) ResourceAttributes() ResourceAttributesPtrOutput {
	return o.ApplyT(func (v SubjectAccessReviewSpec) *ResourceAttributes { return v.ResourceAttributes }).(ResourceAttributesPtrOutput)
}

// UID information about the requesting user.
func (o SubjectAccessReviewSpecOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SubjectAccessReviewSpec) *string { return v.Uid }).(pulumi.StringPtrOutput)
}

// User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
func (o SubjectAccessReviewSpecOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SubjectAccessReviewSpec) *string { return v.User }).(pulumi.StringPtrOutput)
}

// SubjectAccessReviewStatus
type SubjectAccessReviewStatus struct {
	// Allowed is required. True if the action would be allowed, false otherwise.
	Allowed *bool `pulumi:"allowed"`
	// Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
	Denied *bool `pulumi:"denied"`
	// EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
	EvaluationError *string `pulumi:"evaluationError"`
	// Reason is optional.  It indicates why a request was allowed or denied.
	Reason *string `pulumi:"reason"`
}

type SubjectAccessReviewStatusInput interface {
	pulumi.Input

	ToSubjectAccessReviewStatusOutput() SubjectAccessReviewStatusOutput
	ToSubjectAccessReviewStatusOutputWithContext(context.Context) SubjectAccessReviewStatusOutput
}

// SubjectAccessReviewStatus
type SubjectAccessReviewStatusArgs struct {
	// Allowed is required. True if the action would be allowed, false otherwise.
	Allowed pulumi.BoolPtrInput `pulumi:"allowed"`
	// Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
	Denied pulumi.BoolPtrInput `pulumi:"denied"`
	// EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
	EvaluationError pulumi.StringPtrInput `pulumi:"evaluationError"`
	// Reason is optional.  It indicates why a request was allowed or denied.
	Reason pulumi.StringPtrInput `pulumi:"reason"`
}

func (SubjectAccessReviewStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubjectAccessReviewStatus)(nil)).Elem()
}

func (i SubjectAccessReviewStatusArgs) ToSubjectAccessReviewStatusOutput() SubjectAccessReviewStatusOutput {
	return i.ToSubjectAccessReviewStatusOutputWithContext(context.Background())
}

func (i SubjectAccessReviewStatusArgs) ToSubjectAccessReviewStatusOutputWithContext(ctx context.Context) SubjectAccessReviewStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubjectAccessReviewStatusOutput)
}

// SubjectAccessReviewStatus
type SubjectAccessReviewStatusOutput struct { *pulumi.OutputState }

func (SubjectAccessReviewStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubjectAccessReviewStatus)(nil)).Elem()
}

func (o SubjectAccessReviewStatusOutput) ToSubjectAccessReviewStatusOutput() SubjectAccessReviewStatusOutput {
	return o
}

func (o SubjectAccessReviewStatusOutput) ToSubjectAccessReviewStatusOutputWithContext(ctx context.Context) SubjectAccessReviewStatusOutput {
	return o
}

// Allowed is required. True if the action would be allowed, false otherwise.
func (o SubjectAccessReviewStatusOutput) Allowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v SubjectAccessReviewStatus) *bool { return v.Allowed }).(pulumi.BoolPtrOutput)
}

// Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
func (o SubjectAccessReviewStatusOutput) Denied() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v SubjectAccessReviewStatus) *bool { return v.Denied }).(pulumi.BoolPtrOutput)
}

// EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
func (o SubjectAccessReviewStatusOutput) EvaluationError() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SubjectAccessReviewStatus) *string { return v.EvaluationError }).(pulumi.StringPtrOutput)
}

// Reason is optional.  It indicates why a request was allowed or denied.
func (o SubjectAccessReviewStatusOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SubjectAccessReviewStatus) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

// SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it's safe to assume the subject has that permission, even if that list is incomplete.
type SubjectRulesReviewStatus struct {
	// EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.
	EvaluationError *string `pulumi:"evaluationError"`
	// Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn't support rules evaluation.
	Incomplete *bool `pulumi:"incomplete"`
	// NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	NonResourceRules []NonResourceRule `pulumi:"nonResourceRules"`
	// ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	ResourceRules []ResourceRule `pulumi:"resourceRules"`
}

type SubjectRulesReviewStatusInput interface {
	pulumi.Input

	ToSubjectRulesReviewStatusOutput() SubjectRulesReviewStatusOutput
	ToSubjectRulesReviewStatusOutputWithContext(context.Context) SubjectRulesReviewStatusOutput
}

// SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it's safe to assume the subject has that permission, even if that list is incomplete.
type SubjectRulesReviewStatusArgs struct {
	// EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.
	EvaluationError pulumi.StringPtrInput `pulumi:"evaluationError"`
	// Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn't support rules evaluation.
	Incomplete pulumi.BoolPtrInput `pulumi:"incomplete"`
	// NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	NonResourceRules NonResourceRuleArrayInput `pulumi:"nonResourceRules"`
	// ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	ResourceRules ResourceRuleArrayInput `pulumi:"resourceRules"`
}

func (SubjectRulesReviewStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubjectRulesReviewStatus)(nil)).Elem()
}

func (i SubjectRulesReviewStatusArgs) ToSubjectRulesReviewStatusOutput() SubjectRulesReviewStatusOutput {
	return i.ToSubjectRulesReviewStatusOutputWithContext(context.Background())
}

func (i SubjectRulesReviewStatusArgs) ToSubjectRulesReviewStatusOutputWithContext(ctx context.Context) SubjectRulesReviewStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubjectRulesReviewStatusOutput)
}

// SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it's safe to assume the subject has that permission, even if that list is incomplete.
type SubjectRulesReviewStatusOutput struct { *pulumi.OutputState }

func (SubjectRulesReviewStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubjectRulesReviewStatus)(nil)).Elem()
}

func (o SubjectRulesReviewStatusOutput) ToSubjectRulesReviewStatusOutput() SubjectRulesReviewStatusOutput {
	return o
}

func (o SubjectRulesReviewStatusOutput) ToSubjectRulesReviewStatusOutputWithContext(ctx context.Context) SubjectRulesReviewStatusOutput {
	return o
}

// EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.
func (o SubjectRulesReviewStatusOutput) EvaluationError() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SubjectRulesReviewStatus) *string { return v.EvaluationError }).(pulumi.StringPtrOutput)
}

// Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn't support rules evaluation.
func (o SubjectRulesReviewStatusOutput) Incomplete() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v SubjectRulesReviewStatus) *bool { return v.Incomplete }).(pulumi.BoolPtrOutput)
}

// NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
func (o SubjectRulesReviewStatusOutput) NonResourceRules() NonResourceRuleArrayOutput {
	return o.ApplyT(func (v SubjectRulesReviewStatus) []NonResourceRule { return v.NonResourceRules }).(NonResourceRuleArrayOutput)
}

// ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
func (o SubjectRulesReviewStatusOutput) ResourceRules() ResourceRuleArrayOutput {
	return o.ApplyT(func (v SubjectRulesReviewStatus) []ResourceRule { return v.ResourceRules }).(ResourceRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(NonResourceAttributesOutput{})
	pulumi.RegisterOutputType(NonResourceRuleOutput{})
	pulumi.RegisterOutputType(ResourceAttributesOutput{})
	pulumi.RegisterOutputType(ResourceRuleOutput{})
	pulumi.RegisterOutputType(SelfSubjectAccessReviewSpecOutput{})
	pulumi.RegisterOutputType(SelfSubjectRulesReviewSpecOutput{})
	pulumi.RegisterOutputType(SubjectAccessReviewSpecOutput{})
	pulumi.RegisterOutputType(SubjectAccessReviewStatusOutput{})
	pulumi.RegisterOutputType(SubjectRulesReviewStatusOutput{})
}
