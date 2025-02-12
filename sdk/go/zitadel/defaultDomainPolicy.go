// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/scoretechnologies/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Resource representing the default domain policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/scoretechnologies/pulumi-zitadel/sdk/go/zitadel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zitadel.NewDefaultDomainPolicy(ctx, "default", &zitadel.DefaultDomainPolicyArgs{
//				SmtpSenderAddressMatchesInstanceDomain: pulumi.Bool(true),
//				UserLoginMustBeDomain:                  pulumi.Bool(false),
//				ValidateOrgDomains:                     pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// bash The resource can be imported using the ID format `<>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/defaultDomainPolicy:DefaultDomainPolicy imported ''
//
// ```
type DefaultDomainPolicy struct {
	pulumi.CustomResourceState

	SmtpSenderAddressMatchesInstanceDomain pulumi.BoolOutput `pulumi:"smtpSenderAddressMatchesInstanceDomain"`
	// User login must be domain
	UserLoginMustBeDomain pulumi.BoolOutput `pulumi:"userLoginMustBeDomain"`
	// Validate organization domains
	ValidateOrgDomains pulumi.BoolOutput `pulumi:"validateOrgDomains"`
}

// NewDefaultDomainPolicy registers a new resource with the given unique name, arguments, and options.
func NewDefaultDomainPolicy(ctx *pulumi.Context,
	name string, args *DefaultDomainPolicyArgs, opts ...pulumi.ResourceOption) (*DefaultDomainPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SmtpSenderAddressMatchesInstanceDomain == nil {
		return nil, errors.New("invalid value for required argument 'SmtpSenderAddressMatchesInstanceDomain'")
	}
	if args.UserLoginMustBeDomain == nil {
		return nil, errors.New("invalid value for required argument 'UserLoginMustBeDomain'")
	}
	if args.ValidateOrgDomains == nil {
		return nil, errors.New("invalid value for required argument 'ValidateOrgDomains'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultDomainPolicy
	err := ctx.RegisterResource("zitadel:index/defaultDomainPolicy:DefaultDomainPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultDomainPolicy gets an existing DefaultDomainPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultDomainPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultDomainPolicyState, opts ...pulumi.ResourceOption) (*DefaultDomainPolicy, error) {
	var resource DefaultDomainPolicy
	err := ctx.ReadResource("zitadel:index/defaultDomainPolicy:DefaultDomainPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultDomainPolicy resources.
type defaultDomainPolicyState struct {
	SmtpSenderAddressMatchesInstanceDomain *bool `pulumi:"smtpSenderAddressMatchesInstanceDomain"`
	// User login must be domain
	UserLoginMustBeDomain *bool `pulumi:"userLoginMustBeDomain"`
	// Validate organization domains
	ValidateOrgDomains *bool `pulumi:"validateOrgDomains"`
}

type DefaultDomainPolicyState struct {
	SmtpSenderAddressMatchesInstanceDomain pulumi.BoolPtrInput
	// User login must be domain
	UserLoginMustBeDomain pulumi.BoolPtrInput
	// Validate organization domains
	ValidateOrgDomains pulumi.BoolPtrInput
}

func (DefaultDomainPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultDomainPolicyState)(nil)).Elem()
}

type defaultDomainPolicyArgs struct {
	SmtpSenderAddressMatchesInstanceDomain bool `pulumi:"smtpSenderAddressMatchesInstanceDomain"`
	// User login must be domain
	UserLoginMustBeDomain bool `pulumi:"userLoginMustBeDomain"`
	// Validate organization domains
	ValidateOrgDomains bool `pulumi:"validateOrgDomains"`
}

// The set of arguments for constructing a DefaultDomainPolicy resource.
type DefaultDomainPolicyArgs struct {
	SmtpSenderAddressMatchesInstanceDomain pulumi.BoolInput
	// User login must be domain
	UserLoginMustBeDomain pulumi.BoolInput
	// Validate organization domains
	ValidateOrgDomains pulumi.BoolInput
}

func (DefaultDomainPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultDomainPolicyArgs)(nil)).Elem()
}

type DefaultDomainPolicyInput interface {
	pulumi.Input

	ToDefaultDomainPolicyOutput() DefaultDomainPolicyOutput
	ToDefaultDomainPolicyOutputWithContext(ctx context.Context) DefaultDomainPolicyOutput
}

func (*DefaultDomainPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultDomainPolicy)(nil)).Elem()
}

func (i *DefaultDomainPolicy) ToDefaultDomainPolicyOutput() DefaultDomainPolicyOutput {
	return i.ToDefaultDomainPolicyOutputWithContext(context.Background())
}

func (i *DefaultDomainPolicy) ToDefaultDomainPolicyOutputWithContext(ctx context.Context) DefaultDomainPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultDomainPolicyOutput)
}

func (i *DefaultDomainPolicy) ToOutput(ctx context.Context) pulumix.Output[*DefaultDomainPolicy] {
	return pulumix.Output[*DefaultDomainPolicy]{
		OutputState: i.ToDefaultDomainPolicyOutputWithContext(ctx).OutputState,
	}
}

// DefaultDomainPolicyArrayInput is an input type that accepts DefaultDomainPolicyArray and DefaultDomainPolicyArrayOutput values.
// You can construct a concrete instance of `DefaultDomainPolicyArrayInput` via:
//
//	DefaultDomainPolicyArray{ DefaultDomainPolicyArgs{...} }
type DefaultDomainPolicyArrayInput interface {
	pulumi.Input

	ToDefaultDomainPolicyArrayOutput() DefaultDomainPolicyArrayOutput
	ToDefaultDomainPolicyArrayOutputWithContext(context.Context) DefaultDomainPolicyArrayOutput
}

type DefaultDomainPolicyArray []DefaultDomainPolicyInput

func (DefaultDomainPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultDomainPolicy)(nil)).Elem()
}

func (i DefaultDomainPolicyArray) ToDefaultDomainPolicyArrayOutput() DefaultDomainPolicyArrayOutput {
	return i.ToDefaultDomainPolicyArrayOutputWithContext(context.Background())
}

func (i DefaultDomainPolicyArray) ToDefaultDomainPolicyArrayOutputWithContext(ctx context.Context) DefaultDomainPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultDomainPolicyArrayOutput)
}

func (i DefaultDomainPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultDomainPolicy] {
	return pulumix.Output[[]*DefaultDomainPolicy]{
		OutputState: i.ToDefaultDomainPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// DefaultDomainPolicyMapInput is an input type that accepts DefaultDomainPolicyMap and DefaultDomainPolicyMapOutput values.
// You can construct a concrete instance of `DefaultDomainPolicyMapInput` via:
//
//	DefaultDomainPolicyMap{ "key": DefaultDomainPolicyArgs{...} }
type DefaultDomainPolicyMapInput interface {
	pulumi.Input

	ToDefaultDomainPolicyMapOutput() DefaultDomainPolicyMapOutput
	ToDefaultDomainPolicyMapOutputWithContext(context.Context) DefaultDomainPolicyMapOutput
}

type DefaultDomainPolicyMap map[string]DefaultDomainPolicyInput

func (DefaultDomainPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultDomainPolicy)(nil)).Elem()
}

func (i DefaultDomainPolicyMap) ToDefaultDomainPolicyMapOutput() DefaultDomainPolicyMapOutput {
	return i.ToDefaultDomainPolicyMapOutputWithContext(context.Background())
}

func (i DefaultDomainPolicyMap) ToDefaultDomainPolicyMapOutputWithContext(ctx context.Context) DefaultDomainPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultDomainPolicyMapOutput)
}

func (i DefaultDomainPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultDomainPolicy] {
	return pulumix.Output[map[string]*DefaultDomainPolicy]{
		OutputState: i.ToDefaultDomainPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type DefaultDomainPolicyOutput struct{ *pulumi.OutputState }

func (DefaultDomainPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultDomainPolicy)(nil)).Elem()
}

func (o DefaultDomainPolicyOutput) ToDefaultDomainPolicyOutput() DefaultDomainPolicyOutput {
	return o
}

func (o DefaultDomainPolicyOutput) ToDefaultDomainPolicyOutputWithContext(ctx context.Context) DefaultDomainPolicyOutput {
	return o
}

func (o DefaultDomainPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*DefaultDomainPolicy] {
	return pulumix.Output[*DefaultDomainPolicy]{
		OutputState: o.OutputState,
	}
}

func (o DefaultDomainPolicyOutput) SmtpSenderAddressMatchesInstanceDomain() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultDomainPolicy) pulumi.BoolOutput { return v.SmtpSenderAddressMatchesInstanceDomain }).(pulumi.BoolOutput)
}

// User login must be domain
func (o DefaultDomainPolicyOutput) UserLoginMustBeDomain() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultDomainPolicy) pulumi.BoolOutput { return v.UserLoginMustBeDomain }).(pulumi.BoolOutput)
}

// Validate organization domains
func (o DefaultDomainPolicyOutput) ValidateOrgDomains() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultDomainPolicy) pulumi.BoolOutput { return v.ValidateOrgDomains }).(pulumi.BoolOutput)
}

type DefaultDomainPolicyArrayOutput struct{ *pulumi.OutputState }

func (DefaultDomainPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultDomainPolicy)(nil)).Elem()
}

func (o DefaultDomainPolicyArrayOutput) ToDefaultDomainPolicyArrayOutput() DefaultDomainPolicyArrayOutput {
	return o
}

func (o DefaultDomainPolicyArrayOutput) ToDefaultDomainPolicyArrayOutputWithContext(ctx context.Context) DefaultDomainPolicyArrayOutput {
	return o
}

func (o DefaultDomainPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultDomainPolicy] {
	return pulumix.Output[[]*DefaultDomainPolicy]{
		OutputState: o.OutputState,
	}
}

func (o DefaultDomainPolicyArrayOutput) Index(i pulumi.IntInput) DefaultDomainPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultDomainPolicy {
		return vs[0].([]*DefaultDomainPolicy)[vs[1].(int)]
	}).(DefaultDomainPolicyOutput)
}

type DefaultDomainPolicyMapOutput struct{ *pulumi.OutputState }

func (DefaultDomainPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultDomainPolicy)(nil)).Elem()
}

func (o DefaultDomainPolicyMapOutput) ToDefaultDomainPolicyMapOutput() DefaultDomainPolicyMapOutput {
	return o
}

func (o DefaultDomainPolicyMapOutput) ToDefaultDomainPolicyMapOutputWithContext(ctx context.Context) DefaultDomainPolicyMapOutput {
	return o
}

func (o DefaultDomainPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultDomainPolicy] {
	return pulumix.Output[map[string]*DefaultDomainPolicy]{
		OutputState: o.OutputState,
	}
}

func (o DefaultDomainPolicyMapOutput) MapIndex(k pulumi.StringInput) DefaultDomainPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultDomainPolicy {
		return vs[0].(map[string]*DefaultDomainPolicy)[vs[1].(string)]
	}).(DefaultDomainPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultDomainPolicyInput)(nil)).Elem(), &DefaultDomainPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultDomainPolicyArrayInput)(nil)).Elem(), DefaultDomainPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultDomainPolicyMapInput)(nil)).Elem(), DefaultDomainPolicyMap{})
	pulumi.RegisterOutputType(DefaultDomainPolicyOutput{})
	pulumi.RegisterOutputType(DefaultDomainPolicyArrayOutput{})
	pulumi.RegisterOutputType(DefaultDomainPolicyMapOutput{})
}
