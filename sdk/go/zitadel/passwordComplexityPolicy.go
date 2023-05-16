// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the custom password complexity policy of an organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/vavsab/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.NewPasswordComplexityPolicy(ctx, "passwordComplexityPolicy", &zitadel.PasswordComplexityPolicyArgs{
// 			OrgId:        pulumi.Any(zitadel_org.Org.Id),
// 			MinLength:    pulumi.Int(8),
// 			HasUppercase: pulumi.Bool(true),
// 			HasLowercase: pulumi.Bool(true),
// 			HasNumber:    pulumi.Bool(true),
// 			HasSymbol:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PasswordComplexityPolicy struct {
	pulumi.CustomResourceState

	// defines if the password MUST contain a lower case letter
	HasLowercase pulumi.BoolOutput `pulumi:"hasLowercase"`
	// defines if the password MUST contain a number
	HasNumber pulumi.BoolOutput `pulumi:"hasNumber"`
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol pulumi.BoolOutput `pulumi:"hasSymbol"`
	// defines if the password MUST contain an upper case letter
	HasUppercase pulumi.BoolOutput `pulumi:"hasUppercase"`
	// Minimal length for the password
	MinLength pulumi.IntOutput `pulumi:"minLength"`
	// Id for the organization
	OrgId pulumi.StringOutput `pulumi:"orgId"`
}

// NewPasswordComplexityPolicy registers a new resource with the given unique name, arguments, and options.
func NewPasswordComplexityPolicy(ctx *pulumi.Context,
	name string, args *PasswordComplexityPolicyArgs, opts ...pulumi.ResourceOption) (*PasswordComplexityPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HasLowercase == nil {
		return nil, errors.New("invalid value for required argument 'HasLowercase'")
	}
	if args.HasNumber == nil {
		return nil, errors.New("invalid value for required argument 'HasNumber'")
	}
	if args.HasSymbol == nil {
		return nil, errors.New("invalid value for required argument 'HasSymbol'")
	}
	if args.HasUppercase == nil {
		return nil, errors.New("invalid value for required argument 'HasUppercase'")
	}
	if args.MinLength == nil {
		return nil, errors.New("invalid value for required argument 'MinLength'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	var resource PasswordComplexityPolicy
	err := ctx.RegisterResource("zitadel:index/passwordComplexityPolicy:PasswordComplexityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPasswordComplexityPolicy gets an existing PasswordComplexityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPasswordComplexityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PasswordComplexityPolicyState, opts ...pulumi.ResourceOption) (*PasswordComplexityPolicy, error) {
	var resource PasswordComplexityPolicy
	err := ctx.ReadResource("zitadel:index/passwordComplexityPolicy:PasswordComplexityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PasswordComplexityPolicy resources.
type passwordComplexityPolicyState struct {
	// defines if the password MUST contain a lower case letter
	HasLowercase *bool `pulumi:"hasLowercase"`
	// defines if the password MUST contain a number
	HasNumber *bool `pulumi:"hasNumber"`
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol *bool `pulumi:"hasSymbol"`
	// defines if the password MUST contain an upper case letter
	HasUppercase *bool `pulumi:"hasUppercase"`
	// Minimal length for the password
	MinLength *int `pulumi:"minLength"`
	// Id for the organization
	OrgId *string `pulumi:"orgId"`
}

type PasswordComplexityPolicyState struct {
	// defines if the password MUST contain a lower case letter
	HasLowercase pulumi.BoolPtrInput
	// defines if the password MUST contain a number
	HasNumber pulumi.BoolPtrInput
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol pulumi.BoolPtrInput
	// defines if the password MUST contain an upper case letter
	HasUppercase pulumi.BoolPtrInput
	// Minimal length for the password
	MinLength pulumi.IntPtrInput
	// Id for the organization
	OrgId pulumi.StringPtrInput
}

func (PasswordComplexityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordComplexityPolicyState)(nil)).Elem()
}

type passwordComplexityPolicyArgs struct {
	// defines if the password MUST contain a lower case letter
	HasLowercase bool `pulumi:"hasLowercase"`
	// defines if the password MUST contain a number
	HasNumber bool `pulumi:"hasNumber"`
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol bool `pulumi:"hasSymbol"`
	// defines if the password MUST contain an upper case letter
	HasUppercase bool `pulumi:"hasUppercase"`
	// Minimal length for the password
	MinLength int `pulumi:"minLength"`
	// Id for the organization
	OrgId string `pulumi:"orgId"`
}

// The set of arguments for constructing a PasswordComplexityPolicy resource.
type PasswordComplexityPolicyArgs struct {
	// defines if the password MUST contain a lower case letter
	HasLowercase pulumi.BoolInput
	// defines if the password MUST contain a number
	HasNumber pulumi.BoolInput
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol pulumi.BoolInput
	// defines if the password MUST contain an upper case letter
	HasUppercase pulumi.BoolInput
	// Minimal length for the password
	MinLength pulumi.IntInput
	// Id for the organization
	OrgId pulumi.StringInput
}

func (PasswordComplexityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordComplexityPolicyArgs)(nil)).Elem()
}

type PasswordComplexityPolicyInput interface {
	pulumi.Input

	ToPasswordComplexityPolicyOutput() PasswordComplexityPolicyOutput
	ToPasswordComplexityPolicyOutputWithContext(ctx context.Context) PasswordComplexityPolicyOutput
}

func (*PasswordComplexityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordComplexityPolicy)(nil)).Elem()
}

func (i *PasswordComplexityPolicy) ToPasswordComplexityPolicyOutput() PasswordComplexityPolicyOutput {
	return i.ToPasswordComplexityPolicyOutputWithContext(context.Background())
}

func (i *PasswordComplexityPolicy) ToPasswordComplexityPolicyOutputWithContext(ctx context.Context) PasswordComplexityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordComplexityPolicyOutput)
}

// PasswordComplexityPolicyArrayInput is an input type that accepts PasswordComplexityPolicyArray and PasswordComplexityPolicyArrayOutput values.
// You can construct a concrete instance of `PasswordComplexityPolicyArrayInput` via:
//
//          PasswordComplexityPolicyArray{ PasswordComplexityPolicyArgs{...} }
type PasswordComplexityPolicyArrayInput interface {
	pulumi.Input

	ToPasswordComplexityPolicyArrayOutput() PasswordComplexityPolicyArrayOutput
	ToPasswordComplexityPolicyArrayOutputWithContext(context.Context) PasswordComplexityPolicyArrayOutput
}

type PasswordComplexityPolicyArray []PasswordComplexityPolicyInput

func (PasswordComplexityPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordComplexityPolicy)(nil)).Elem()
}

func (i PasswordComplexityPolicyArray) ToPasswordComplexityPolicyArrayOutput() PasswordComplexityPolicyArrayOutput {
	return i.ToPasswordComplexityPolicyArrayOutputWithContext(context.Background())
}

func (i PasswordComplexityPolicyArray) ToPasswordComplexityPolicyArrayOutputWithContext(ctx context.Context) PasswordComplexityPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordComplexityPolicyArrayOutput)
}

// PasswordComplexityPolicyMapInput is an input type that accepts PasswordComplexityPolicyMap and PasswordComplexityPolicyMapOutput values.
// You can construct a concrete instance of `PasswordComplexityPolicyMapInput` via:
//
//          PasswordComplexityPolicyMap{ "key": PasswordComplexityPolicyArgs{...} }
type PasswordComplexityPolicyMapInput interface {
	pulumi.Input

	ToPasswordComplexityPolicyMapOutput() PasswordComplexityPolicyMapOutput
	ToPasswordComplexityPolicyMapOutputWithContext(context.Context) PasswordComplexityPolicyMapOutput
}

type PasswordComplexityPolicyMap map[string]PasswordComplexityPolicyInput

func (PasswordComplexityPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordComplexityPolicy)(nil)).Elem()
}

func (i PasswordComplexityPolicyMap) ToPasswordComplexityPolicyMapOutput() PasswordComplexityPolicyMapOutput {
	return i.ToPasswordComplexityPolicyMapOutputWithContext(context.Background())
}

func (i PasswordComplexityPolicyMap) ToPasswordComplexityPolicyMapOutputWithContext(ctx context.Context) PasswordComplexityPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordComplexityPolicyMapOutput)
}

type PasswordComplexityPolicyOutput struct{ *pulumi.OutputState }

func (PasswordComplexityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordComplexityPolicy)(nil)).Elem()
}

func (o PasswordComplexityPolicyOutput) ToPasswordComplexityPolicyOutput() PasswordComplexityPolicyOutput {
	return o
}

func (o PasswordComplexityPolicyOutput) ToPasswordComplexityPolicyOutputWithContext(ctx context.Context) PasswordComplexityPolicyOutput {
	return o
}

// defines if the password MUST contain a lower case letter
func (o PasswordComplexityPolicyOutput) HasLowercase() pulumi.BoolOutput {
	return o.ApplyT(func(v *PasswordComplexityPolicy) pulumi.BoolOutput { return v.HasLowercase }).(pulumi.BoolOutput)
}

// defines if the password MUST contain a number
func (o PasswordComplexityPolicyOutput) HasNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v *PasswordComplexityPolicy) pulumi.BoolOutput { return v.HasNumber }).(pulumi.BoolOutput)
}

// defines if the password MUST contain a symbol. E.g. "$"
func (o PasswordComplexityPolicyOutput) HasSymbol() pulumi.BoolOutput {
	return o.ApplyT(func(v *PasswordComplexityPolicy) pulumi.BoolOutput { return v.HasSymbol }).(pulumi.BoolOutput)
}

// defines if the password MUST contain an upper case letter
func (o PasswordComplexityPolicyOutput) HasUppercase() pulumi.BoolOutput {
	return o.ApplyT(func(v *PasswordComplexityPolicy) pulumi.BoolOutput { return v.HasUppercase }).(pulumi.BoolOutput)
}

// Minimal length for the password
func (o PasswordComplexityPolicyOutput) MinLength() pulumi.IntOutput {
	return o.ApplyT(func(v *PasswordComplexityPolicy) pulumi.IntOutput { return v.MinLength }).(pulumi.IntOutput)
}

// Id for the organization
func (o PasswordComplexityPolicyOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *PasswordComplexityPolicy) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

type PasswordComplexityPolicyArrayOutput struct{ *pulumi.OutputState }

func (PasswordComplexityPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordComplexityPolicy)(nil)).Elem()
}

func (o PasswordComplexityPolicyArrayOutput) ToPasswordComplexityPolicyArrayOutput() PasswordComplexityPolicyArrayOutput {
	return o
}

func (o PasswordComplexityPolicyArrayOutput) ToPasswordComplexityPolicyArrayOutputWithContext(ctx context.Context) PasswordComplexityPolicyArrayOutput {
	return o
}

func (o PasswordComplexityPolicyArrayOutput) Index(i pulumi.IntInput) PasswordComplexityPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PasswordComplexityPolicy {
		return vs[0].([]*PasswordComplexityPolicy)[vs[1].(int)]
	}).(PasswordComplexityPolicyOutput)
}

type PasswordComplexityPolicyMapOutput struct{ *pulumi.OutputState }

func (PasswordComplexityPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordComplexityPolicy)(nil)).Elem()
}

func (o PasswordComplexityPolicyMapOutput) ToPasswordComplexityPolicyMapOutput() PasswordComplexityPolicyMapOutput {
	return o
}

func (o PasswordComplexityPolicyMapOutput) ToPasswordComplexityPolicyMapOutputWithContext(ctx context.Context) PasswordComplexityPolicyMapOutput {
	return o
}

func (o PasswordComplexityPolicyMapOutput) MapIndex(k pulumi.StringInput) PasswordComplexityPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PasswordComplexityPolicy {
		return vs[0].(map[string]*PasswordComplexityPolicy)[vs[1].(string)]
	}).(PasswordComplexityPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordComplexityPolicyInput)(nil)).Elem(), &PasswordComplexityPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordComplexityPolicyArrayInput)(nil)).Elem(), PasswordComplexityPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordComplexityPolicyMapInput)(nil)).Elem(), PasswordComplexityPolicyMap{})
	pulumi.RegisterOutputType(PasswordComplexityPolicyOutput{})
	pulumi.RegisterOutputType(PasswordComplexityPolicyArrayOutput{})
	pulumi.RegisterOutputType(PasswordComplexityPolicyMapOutput{})
}
