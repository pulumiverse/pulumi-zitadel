// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Resource representing the default lockout policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zitadel.NewDefaultLockoutPolicy(ctx, "default", &zitadel.DefaultLockoutPolicyArgs{
//				MaxPasswordAttempts: pulumi.Int(5),
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
//	$ pulumi import zitadel:index/defaultLockoutPolicy:DefaultLockoutPolicy imported ''
//
// ```
type DefaultLockoutPolicy struct {
	pulumi.CustomResourceState

	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	MaxPasswordAttempts pulumi.IntOutput `pulumi:"maxPasswordAttempts"`
}

// NewDefaultLockoutPolicy registers a new resource with the given unique name, arguments, and options.
func NewDefaultLockoutPolicy(ctx *pulumi.Context,
	name string, args *DefaultLockoutPolicyArgs, opts ...pulumi.ResourceOption) (*DefaultLockoutPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxPasswordAttempts == nil {
		return nil, errors.New("invalid value for required argument 'MaxPasswordAttempts'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultLockoutPolicy
	err := ctx.RegisterResource("zitadel:index/defaultLockoutPolicy:DefaultLockoutPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultLockoutPolicy gets an existing DefaultLockoutPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultLockoutPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultLockoutPolicyState, opts ...pulumi.ResourceOption) (*DefaultLockoutPolicy, error) {
	var resource DefaultLockoutPolicy
	err := ctx.ReadResource("zitadel:index/defaultLockoutPolicy:DefaultLockoutPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultLockoutPolicy resources.
type defaultLockoutPolicyState struct {
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	MaxPasswordAttempts *int `pulumi:"maxPasswordAttempts"`
}

type DefaultLockoutPolicyState struct {
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	MaxPasswordAttempts pulumi.IntPtrInput
}

func (DefaultLockoutPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultLockoutPolicyState)(nil)).Elem()
}

type defaultLockoutPolicyArgs struct {
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	MaxPasswordAttempts int `pulumi:"maxPasswordAttempts"`
}

// The set of arguments for constructing a DefaultLockoutPolicy resource.
type DefaultLockoutPolicyArgs struct {
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	MaxPasswordAttempts pulumi.IntInput
}

func (DefaultLockoutPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultLockoutPolicyArgs)(nil)).Elem()
}

type DefaultLockoutPolicyInput interface {
	pulumi.Input

	ToDefaultLockoutPolicyOutput() DefaultLockoutPolicyOutput
	ToDefaultLockoutPolicyOutputWithContext(ctx context.Context) DefaultLockoutPolicyOutput
}

func (*DefaultLockoutPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultLockoutPolicy)(nil)).Elem()
}

func (i *DefaultLockoutPolicy) ToDefaultLockoutPolicyOutput() DefaultLockoutPolicyOutput {
	return i.ToDefaultLockoutPolicyOutputWithContext(context.Background())
}

func (i *DefaultLockoutPolicy) ToDefaultLockoutPolicyOutputWithContext(ctx context.Context) DefaultLockoutPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLockoutPolicyOutput)
}

func (i *DefaultLockoutPolicy) ToOutput(ctx context.Context) pulumix.Output[*DefaultLockoutPolicy] {
	return pulumix.Output[*DefaultLockoutPolicy]{
		OutputState: i.ToDefaultLockoutPolicyOutputWithContext(ctx).OutputState,
	}
}

// DefaultLockoutPolicyArrayInput is an input type that accepts DefaultLockoutPolicyArray and DefaultLockoutPolicyArrayOutput values.
// You can construct a concrete instance of `DefaultLockoutPolicyArrayInput` via:
//
//	DefaultLockoutPolicyArray{ DefaultLockoutPolicyArgs{...} }
type DefaultLockoutPolicyArrayInput interface {
	pulumi.Input

	ToDefaultLockoutPolicyArrayOutput() DefaultLockoutPolicyArrayOutput
	ToDefaultLockoutPolicyArrayOutputWithContext(context.Context) DefaultLockoutPolicyArrayOutput
}

type DefaultLockoutPolicyArray []DefaultLockoutPolicyInput

func (DefaultLockoutPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultLockoutPolicy)(nil)).Elem()
}

func (i DefaultLockoutPolicyArray) ToDefaultLockoutPolicyArrayOutput() DefaultLockoutPolicyArrayOutput {
	return i.ToDefaultLockoutPolicyArrayOutputWithContext(context.Background())
}

func (i DefaultLockoutPolicyArray) ToDefaultLockoutPolicyArrayOutputWithContext(ctx context.Context) DefaultLockoutPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLockoutPolicyArrayOutput)
}

func (i DefaultLockoutPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultLockoutPolicy] {
	return pulumix.Output[[]*DefaultLockoutPolicy]{
		OutputState: i.ToDefaultLockoutPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// DefaultLockoutPolicyMapInput is an input type that accepts DefaultLockoutPolicyMap and DefaultLockoutPolicyMapOutput values.
// You can construct a concrete instance of `DefaultLockoutPolicyMapInput` via:
//
//	DefaultLockoutPolicyMap{ "key": DefaultLockoutPolicyArgs{...} }
type DefaultLockoutPolicyMapInput interface {
	pulumi.Input

	ToDefaultLockoutPolicyMapOutput() DefaultLockoutPolicyMapOutput
	ToDefaultLockoutPolicyMapOutputWithContext(context.Context) DefaultLockoutPolicyMapOutput
}

type DefaultLockoutPolicyMap map[string]DefaultLockoutPolicyInput

func (DefaultLockoutPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultLockoutPolicy)(nil)).Elem()
}

func (i DefaultLockoutPolicyMap) ToDefaultLockoutPolicyMapOutput() DefaultLockoutPolicyMapOutput {
	return i.ToDefaultLockoutPolicyMapOutputWithContext(context.Background())
}

func (i DefaultLockoutPolicyMap) ToDefaultLockoutPolicyMapOutputWithContext(ctx context.Context) DefaultLockoutPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLockoutPolicyMapOutput)
}

func (i DefaultLockoutPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultLockoutPolicy] {
	return pulumix.Output[map[string]*DefaultLockoutPolicy]{
		OutputState: i.ToDefaultLockoutPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type DefaultLockoutPolicyOutput struct{ *pulumi.OutputState }

func (DefaultLockoutPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultLockoutPolicy)(nil)).Elem()
}

func (o DefaultLockoutPolicyOutput) ToDefaultLockoutPolicyOutput() DefaultLockoutPolicyOutput {
	return o
}

func (o DefaultLockoutPolicyOutput) ToDefaultLockoutPolicyOutputWithContext(ctx context.Context) DefaultLockoutPolicyOutput {
	return o
}

func (o DefaultLockoutPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*DefaultLockoutPolicy] {
	return pulumix.Output[*DefaultLockoutPolicy]{
		OutputState: o.OutputState,
	}
}

// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
func (o DefaultLockoutPolicyOutput) MaxPasswordAttempts() pulumi.IntOutput {
	return o.ApplyT(func(v *DefaultLockoutPolicy) pulumi.IntOutput { return v.MaxPasswordAttempts }).(pulumi.IntOutput)
}

type DefaultLockoutPolicyArrayOutput struct{ *pulumi.OutputState }

func (DefaultLockoutPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultLockoutPolicy)(nil)).Elem()
}

func (o DefaultLockoutPolicyArrayOutput) ToDefaultLockoutPolicyArrayOutput() DefaultLockoutPolicyArrayOutput {
	return o
}

func (o DefaultLockoutPolicyArrayOutput) ToDefaultLockoutPolicyArrayOutputWithContext(ctx context.Context) DefaultLockoutPolicyArrayOutput {
	return o
}

func (o DefaultLockoutPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultLockoutPolicy] {
	return pulumix.Output[[]*DefaultLockoutPolicy]{
		OutputState: o.OutputState,
	}
}

func (o DefaultLockoutPolicyArrayOutput) Index(i pulumi.IntInput) DefaultLockoutPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultLockoutPolicy {
		return vs[0].([]*DefaultLockoutPolicy)[vs[1].(int)]
	}).(DefaultLockoutPolicyOutput)
}

type DefaultLockoutPolicyMapOutput struct{ *pulumi.OutputState }

func (DefaultLockoutPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultLockoutPolicy)(nil)).Elem()
}

func (o DefaultLockoutPolicyMapOutput) ToDefaultLockoutPolicyMapOutput() DefaultLockoutPolicyMapOutput {
	return o
}

func (o DefaultLockoutPolicyMapOutput) ToDefaultLockoutPolicyMapOutputWithContext(ctx context.Context) DefaultLockoutPolicyMapOutput {
	return o
}

func (o DefaultLockoutPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultLockoutPolicy] {
	return pulumix.Output[map[string]*DefaultLockoutPolicy]{
		OutputState: o.OutputState,
	}
}

func (o DefaultLockoutPolicyMapOutput) MapIndex(k pulumi.StringInput) DefaultLockoutPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultLockoutPolicy {
		return vs[0].(map[string]*DefaultLockoutPolicy)[vs[1].(string)]
	}).(DefaultLockoutPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLockoutPolicyInput)(nil)).Elem(), &DefaultLockoutPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLockoutPolicyArrayInput)(nil)).Elem(), DefaultLockoutPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLockoutPolicyMapInput)(nil)).Elem(), DefaultLockoutPolicyMap{})
	pulumi.RegisterOutputType(DefaultLockoutPolicyOutput{})
	pulumi.RegisterOutputType(DefaultLockoutPolicyArrayOutput{})
	pulumi.RegisterOutputType(DefaultLockoutPolicyMapOutput{})
}
