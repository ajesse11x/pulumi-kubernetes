// *** WARNING: this file was generated by foo. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
type AggregationRule struct {
	// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
	ClusterRoleSelectors []metav1.LabelSelector `pulumi:"clusterRoleSelectors"`
}

type AggregationRuleInput interface {
	pulumi.Input

	ToAggregationRuleOutput() AggregationRuleOutput
	ToAggregationRuleOutputWithContext(context.Context) AggregationRuleOutput
}

// AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
type AggregationRuleArgs struct {
	// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
	ClusterRoleSelectors metav1.LabelSelectorArrayInput `pulumi:"clusterRoleSelectors"`
}

func (AggregationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AggregationRule)(nil)).Elem()
}

func (i AggregationRuleArgs) ToAggregationRuleOutput() AggregationRuleOutput {
	return i.ToAggregationRuleOutputWithContext(context.Background())
}

func (i AggregationRuleArgs) ToAggregationRuleOutputWithContext(ctx context.Context) AggregationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregationRuleOutput)
}

func (i AggregationRuleArgs) ToAggregationRulePtrOutput() AggregationRulePtrOutput {
	return i.ToAggregationRulePtrOutputWithContext(context.Background())
}

func (i AggregationRuleArgs) ToAggregationRulePtrOutputWithContext(ctx context.Context) AggregationRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregationRuleOutput).ToAggregationRulePtrOutputWithContext(ctx)
}

type AggregationRulePtrInput interface {
	pulumi.Input

	ToAggregationRulePtrOutput() AggregationRulePtrOutput
	ToAggregationRulePtrOutputWithContext(context.Context) AggregationRulePtrOutput
}

type aggregationRulePtrType AggregationRuleArgs

func AggregationRulePtr(v *AggregationRuleArgs) AggregationRulePtrInput {	return (*aggregationRulePtrType)(v)
}

func (*aggregationRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregationRule)(nil)).Elem()
}

func (i *aggregationRulePtrType) ToAggregationRulePtrOutput() AggregationRulePtrOutput {
	return i.ToAggregationRulePtrOutputWithContext(context.Background())
}

func (i *aggregationRulePtrType) ToAggregationRulePtrOutputWithContext(ctx context.Context) AggregationRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregationRulePtrOutput)
}

// AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
type AggregationRuleOutput struct { *pulumi.OutputState }

func (AggregationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AggregationRule)(nil)).Elem()
}

func (o AggregationRuleOutput) ToAggregationRuleOutput() AggregationRuleOutput {
	return o
}

func (o AggregationRuleOutput) ToAggregationRuleOutputWithContext(ctx context.Context) AggregationRuleOutput {
	return o
}

func (o AggregationRuleOutput) ToAggregationRulePtrOutput() AggregationRulePtrOutput {
	return o.ToAggregationRulePtrOutputWithContext(context.Background())
}

func (o AggregationRuleOutput) ToAggregationRulePtrOutputWithContext(ctx context.Context) AggregationRulePtrOutput {
	return o.ApplyT(func(v AggregationRule) *AggregationRule {
		return &v
	}).(AggregationRulePtrOutput)
}
// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
func (o AggregationRuleOutput) ClusterRoleSelectors() metav1.LabelSelectorArrayOutput {
	return o.ApplyT(func (v AggregationRule) []metav1.LabelSelector { return v.ClusterRoleSelectors }).(metav1.LabelSelectorArrayOutput)
}

type AggregationRulePtrOutput struct { *pulumi.OutputState}

func (AggregationRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregationRule)(nil)).Elem()
}

func (o AggregationRulePtrOutput) ToAggregationRulePtrOutput() AggregationRulePtrOutput {
	return o
}

func (o AggregationRulePtrOutput) ToAggregationRulePtrOutputWithContext(ctx context.Context) AggregationRulePtrOutput {
	return o
}

func (o AggregationRulePtrOutput) Elem() AggregationRuleOutput {
	return o.ApplyT(func (v *AggregationRule) AggregationRule { return *v }).(AggregationRuleOutput)
}

// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
func (o AggregationRulePtrOutput) ClusterRoleSelectors() metav1.LabelSelectorArrayOutput {
	return o.ApplyT(func (v AggregationRule) []metav1.LabelSelector { return v.ClusterRoleSelectors }).(metav1.LabelSelectorArrayOutput)
}

// PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.
type PolicyRule struct {
	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
	ApiGroups []string `pulumi:"apiGroups"`
	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
	NonResourceURLs []string `pulumi:"nonResourceURLs"`
	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
	ResourceNames []string `pulumi:"resourceNames"`
	// Resources is a list of resources this rule applies to.  '*' represents all resources in the specified apiGroups. '*/foo' represents the subresource 'foo' for all resources in the specified apiGroups.
	Resources []string `pulumi:"resources"`
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
	Verbs []string `pulumi:"verbs"`
}

type PolicyRuleInput interface {
	pulumi.Input

	ToPolicyRuleOutput() PolicyRuleOutput
	ToPolicyRuleOutputWithContext(context.Context) PolicyRuleOutput
}

// PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.
type PolicyRuleArgs struct {
	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
	ApiGroups pulumi.StringArrayInput `pulumi:"apiGroups"`
	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
	NonResourceURLs pulumi.StringArrayInput `pulumi:"nonResourceURLs"`
	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
	ResourceNames pulumi.StringArrayInput `pulumi:"resourceNames"`
	// Resources is a list of resources this rule applies to.  '*' represents all resources in the specified apiGroups. '*/foo' represents the subresource 'foo' for all resources in the specified apiGroups.
	Resources pulumi.StringArrayInput `pulumi:"resources"`
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
	Verbs pulumi.StringArrayInput `pulumi:"verbs"`
}

func (PolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyRule)(nil)).Elem()
}

func (i PolicyRuleArgs) ToPolicyRuleOutput() PolicyRuleOutput {
	return i.ToPolicyRuleOutputWithContext(context.Background())
}

func (i PolicyRuleArgs) ToPolicyRuleOutputWithContext(ctx context.Context) PolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyRuleOutput)
}

type PolicyRuleArrayInput interface {
	pulumi.Input

	ToPolicyRuleArrayOutput() PolicyRuleArrayOutput
	ToPolicyRuleArrayOutputWithContext(context.Context) PolicyRuleArrayOutput
}

type PolicyRuleArray []PolicyRuleInput

func (PolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyRule)(nil)).Elem()
}

func (i PolicyRuleArray) ToPolicyRuleArrayOutput() PolicyRuleArrayOutput {
	return i.ToPolicyRuleArrayOutputWithContext(context.Background())
}

func (i PolicyRuleArray) ToPolicyRuleArrayOutputWithContext(ctx context.Context) PolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyRuleArrayOutput)
}

// PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.
type PolicyRuleOutput struct { *pulumi.OutputState }

func (PolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyRule)(nil)).Elem()
}

func (o PolicyRuleOutput) ToPolicyRuleOutput() PolicyRuleOutput {
	return o
}

func (o PolicyRuleOutput) ToPolicyRuleOutputWithContext(ctx context.Context) PolicyRuleOutput {
	return o
}

// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
func (o PolicyRuleOutput) ApiGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func (v PolicyRule) []string { return v.ApiGroups }).(pulumi.StringArrayOutput)
}

// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
func (o PolicyRuleOutput) NonResourceURLs() pulumi.StringArrayOutput {
	return o.ApplyT(func (v PolicyRule) []string { return v.NonResourceURLs }).(pulumi.StringArrayOutput)
}

// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
func (o PolicyRuleOutput) ResourceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func (v PolicyRule) []string { return v.ResourceNames }).(pulumi.StringArrayOutput)
}

// Resources is a list of resources this rule applies to.  '*' represents all resources in the specified apiGroups. '*/foo' represents the subresource 'foo' for all resources in the specified apiGroups.
func (o PolicyRuleOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func (v PolicyRule) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
func (o PolicyRuleOutput) Verbs() pulumi.StringArrayOutput {
	return o.ApplyT(func (v PolicyRule) []string { return v.Verbs }).(pulumi.StringArrayOutput)
}

type PolicyRuleArrayOutput struct { *pulumi.OutputState}

func (PolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyRule)(nil)).Elem()
}

func (o PolicyRuleArrayOutput) ToPolicyRuleArrayOutput() PolicyRuleArrayOutput {
	return o
}

func (o PolicyRuleArrayOutput) ToPolicyRuleArrayOutputWithContext(ctx context.Context) PolicyRuleArrayOutput {
	return o
}

func (o PolicyRuleArrayOutput) Index(i pulumi.IntInput) PolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) PolicyRule {
		return vs[0].([]PolicyRule)[vs[1].(int)]
	}).(PolicyRuleOutput)
}

// RoleRef contains information that points to the role being used
type RoleRef struct {
	// APIGroup is the group for the resource being referenced
	ApiGroup *string `pulumi:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind *string `pulumi:"kind"`
	// Name is the name of resource being referenced
	Name *string `pulumi:"name"`
}

type RoleRefInput interface {
	pulumi.Input

	ToRoleRefOutput() RoleRefOutput
	ToRoleRefOutputWithContext(context.Context) RoleRefOutput
}

// RoleRef contains information that points to the role being used
type RoleRefArgs struct {
	// APIGroup is the group for the resource being referenced
	ApiGroup pulumi.StringPtrInput `pulumi:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Name is the name of resource being referenced
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (RoleRefArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleRef)(nil)).Elem()
}

func (i RoleRefArgs) ToRoleRefOutput() RoleRefOutput {
	return i.ToRoleRefOutputWithContext(context.Background())
}

func (i RoleRefArgs) ToRoleRefOutputWithContext(ctx context.Context) RoleRefOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleRefOutput)
}

func (i RoleRefArgs) ToRoleRefPtrOutput() RoleRefPtrOutput {
	return i.ToRoleRefPtrOutputWithContext(context.Background())
}

func (i RoleRefArgs) ToRoleRefPtrOutputWithContext(ctx context.Context) RoleRefPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleRefOutput).ToRoleRefPtrOutputWithContext(ctx)
}

type RoleRefPtrInput interface {
	pulumi.Input

	ToRoleRefPtrOutput() RoleRefPtrOutput
	ToRoleRefPtrOutputWithContext(context.Context) RoleRefPtrOutput
}

type roleRefPtrType RoleRefArgs

func RoleRefPtr(v *RoleRefArgs) RoleRefPtrInput {	return (*roleRefPtrType)(v)
}

func (*roleRefPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleRef)(nil)).Elem()
}

func (i *roleRefPtrType) ToRoleRefPtrOutput() RoleRefPtrOutput {
	return i.ToRoleRefPtrOutputWithContext(context.Background())
}

func (i *roleRefPtrType) ToRoleRefPtrOutputWithContext(ctx context.Context) RoleRefPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleRefPtrOutput)
}

// RoleRef contains information that points to the role being used
type RoleRefOutput struct { *pulumi.OutputState }

func (RoleRefOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleRef)(nil)).Elem()
}

func (o RoleRefOutput) ToRoleRefOutput() RoleRefOutput {
	return o
}

func (o RoleRefOutput) ToRoleRefOutputWithContext(ctx context.Context) RoleRefOutput {
	return o
}

func (o RoleRefOutput) ToRoleRefPtrOutput() RoleRefPtrOutput {
	return o.ToRoleRefPtrOutputWithContext(context.Background())
}

func (o RoleRefOutput) ToRoleRefPtrOutputWithContext(ctx context.Context) RoleRefPtrOutput {
	return o.ApplyT(func(v RoleRef) *RoleRef {
		return &v
	}).(RoleRefPtrOutput)
}
// APIGroup is the group for the resource being referenced
func (o RoleRefOutput) ApiGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RoleRef) *string { return v.ApiGroup }).(pulumi.StringPtrOutput)
}

// Kind is the type of resource being referenced
func (o RoleRefOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RoleRef) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Name is the name of resource being referenced
func (o RoleRefOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RoleRef) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RoleRefPtrOutput struct { *pulumi.OutputState}

func (RoleRefPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleRef)(nil)).Elem()
}

func (o RoleRefPtrOutput) ToRoleRefPtrOutput() RoleRefPtrOutput {
	return o
}

func (o RoleRefPtrOutput) ToRoleRefPtrOutputWithContext(ctx context.Context) RoleRefPtrOutput {
	return o
}

func (o RoleRefPtrOutput) Elem() RoleRefOutput {
	return o.ApplyT(func (v *RoleRef) RoleRef { return *v }).(RoleRefOutput)
}

// APIGroup is the group for the resource being referenced
func (o RoleRefPtrOutput) ApiGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RoleRef) *string { return v.ApiGroup }).(pulumi.StringPtrOutput)
}

// Kind is the type of resource being referenced
func (o RoleRefPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RoleRef) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Name is the name of resource being referenced
func (o RoleRefPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RoleRef) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.
type Subject struct {
	// APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
	ApiGroup *string `pulumi:"apiGroup"`
	// Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Kind *string `pulumi:"kind"`
	// Name of the object being referenced.
	Name *string `pulumi:"name"`
	// Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
	Namespace *string `pulumi:"namespace"`
}

type SubjectInput interface {
	pulumi.Input

	ToSubjectOutput() SubjectOutput
	ToSubjectOutputWithContext(context.Context) SubjectOutput
}

// Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.
type SubjectArgs struct {
	// APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
	ApiGroup pulumi.StringPtrInput `pulumi:"apiGroup"`
	// Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Name of the object being referenced.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (SubjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Subject)(nil)).Elem()
}

func (i SubjectArgs) ToSubjectOutput() SubjectOutput {
	return i.ToSubjectOutputWithContext(context.Background())
}

func (i SubjectArgs) ToSubjectOutputWithContext(ctx context.Context) SubjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubjectOutput)
}

type SubjectArrayInput interface {
	pulumi.Input

	ToSubjectArrayOutput() SubjectArrayOutput
	ToSubjectArrayOutputWithContext(context.Context) SubjectArrayOutput
}

type SubjectArray []SubjectInput

func (SubjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subject)(nil)).Elem()
}

func (i SubjectArray) ToSubjectArrayOutput() SubjectArrayOutput {
	return i.ToSubjectArrayOutputWithContext(context.Background())
}

func (i SubjectArray) ToSubjectArrayOutputWithContext(ctx context.Context) SubjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubjectArrayOutput)
}

// Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.
type SubjectOutput struct { *pulumi.OutputState }

func (SubjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subject)(nil)).Elem()
}

func (o SubjectOutput) ToSubjectOutput() SubjectOutput {
	return o
}

func (o SubjectOutput) ToSubjectOutputWithContext(ctx context.Context) SubjectOutput {
	return o
}

// APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
func (o SubjectOutput) ApiGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func (v Subject) *string { return v.ApiGroup }).(pulumi.StringPtrOutput)
}

// Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
func (o SubjectOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func (v Subject) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Name of the object being referenced.
func (o SubjectOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v Subject) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
func (o SubjectOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func (v Subject) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

type SubjectArrayOutput struct { *pulumi.OutputState}

func (SubjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subject)(nil)).Elem()
}

func (o SubjectArrayOutput) ToSubjectArrayOutput() SubjectArrayOutput {
	return o
}

func (o SubjectArrayOutput) ToSubjectArrayOutputWithContext(ctx context.Context) SubjectArrayOutput {
	return o
}

func (o SubjectArrayOutput) Index(i pulumi.IntInput) SubjectOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) Subject {
		return vs[0].([]Subject)[vs[1].(int)]
	}).(SubjectOutput)
}

func init() {
	pulumi.RegisterOutputType(AggregationRuleOutput{})
	pulumi.RegisterOutputType(AggregationRulePtrOutput{})
	pulumi.RegisterOutputType(PolicyRuleOutput{})
	pulumi.RegisterOutputType(PolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(RoleRefOutput{})
	pulumi.RegisterOutputType(RoleRefPtrOutput{})
	pulumi.RegisterOutputType(SubjectOutput{})
	pulumi.RegisterOutputType(SubjectArrayOutput{})
}
