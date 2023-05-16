// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing a personal access token of a user
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
// 		_, err := zitadel.NewPersonalAccessToken(ctx, "pat", &zitadel.PersonalAccessTokenArgs{
// 			OrgId:          pulumi.Any(zitadel_org.Org.Id),
// 			UserId:         pulumi.Any(zitadel_machine_user.Machine_user.Id),
// 			ExpirationDate: pulumi.String("2519-04-01T08:45:00Z"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PersonalAccessToken struct {
	pulumi.CustomResourceState

	// Expiration date of the token in the RFC3339 format
	ExpirationDate pulumi.StringPtrOutput `pulumi:"expirationDate"`
	// ID of the organization
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Value of the token
	Token pulumi.StringOutput `pulumi:"token"`
	// ID of the user
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewPersonalAccessToken registers a new resource with the given unique name, arguments, and options.
func NewPersonalAccessToken(ctx *pulumi.Context,
	name string, args *PersonalAccessTokenArgs, opts ...pulumi.ResourceOption) (*PersonalAccessToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource PersonalAccessToken
	err := ctx.RegisterResource("zitadel:index/personalAccessToken:PersonalAccessToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersonalAccessToken gets an existing PersonalAccessToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersonalAccessToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersonalAccessTokenState, opts ...pulumi.ResourceOption) (*PersonalAccessToken, error) {
	var resource PersonalAccessToken
	err := ctx.ReadResource("zitadel:index/personalAccessToken:PersonalAccessToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersonalAccessToken resources.
type personalAccessTokenState struct {
	// Expiration date of the token in the RFC3339 format
	ExpirationDate *string `pulumi:"expirationDate"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Value of the token
	Token *string `pulumi:"token"`
	// ID of the user
	UserId *string `pulumi:"userId"`
}

type PersonalAccessTokenState struct {
	// Expiration date of the token in the RFC3339 format
	ExpirationDate pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Value of the token
	Token pulumi.StringPtrInput
	// ID of the user
	UserId pulumi.StringPtrInput
}

func (PersonalAccessTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*personalAccessTokenState)(nil)).Elem()
}

type personalAccessTokenArgs struct {
	// Expiration date of the token in the RFC3339 format
	ExpirationDate *string `pulumi:"expirationDate"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// ID of the user
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a PersonalAccessToken resource.
type PersonalAccessTokenArgs struct {
	// Expiration date of the token in the RFC3339 format
	ExpirationDate pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringInput
	// ID of the user
	UserId pulumi.StringInput
}

func (PersonalAccessTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*personalAccessTokenArgs)(nil)).Elem()
}

type PersonalAccessTokenInput interface {
	pulumi.Input

	ToPersonalAccessTokenOutput() PersonalAccessTokenOutput
	ToPersonalAccessTokenOutputWithContext(ctx context.Context) PersonalAccessTokenOutput
}

func (*PersonalAccessToken) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalAccessToken)(nil)).Elem()
}

func (i *PersonalAccessToken) ToPersonalAccessTokenOutput() PersonalAccessTokenOutput {
	return i.ToPersonalAccessTokenOutputWithContext(context.Background())
}

func (i *PersonalAccessToken) ToPersonalAccessTokenOutputWithContext(ctx context.Context) PersonalAccessTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalAccessTokenOutput)
}

// PersonalAccessTokenArrayInput is an input type that accepts PersonalAccessTokenArray and PersonalAccessTokenArrayOutput values.
// You can construct a concrete instance of `PersonalAccessTokenArrayInput` via:
//
//          PersonalAccessTokenArray{ PersonalAccessTokenArgs{...} }
type PersonalAccessTokenArrayInput interface {
	pulumi.Input

	ToPersonalAccessTokenArrayOutput() PersonalAccessTokenArrayOutput
	ToPersonalAccessTokenArrayOutputWithContext(context.Context) PersonalAccessTokenArrayOutput
}

type PersonalAccessTokenArray []PersonalAccessTokenInput

func (PersonalAccessTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersonalAccessToken)(nil)).Elem()
}

func (i PersonalAccessTokenArray) ToPersonalAccessTokenArrayOutput() PersonalAccessTokenArrayOutput {
	return i.ToPersonalAccessTokenArrayOutputWithContext(context.Background())
}

func (i PersonalAccessTokenArray) ToPersonalAccessTokenArrayOutputWithContext(ctx context.Context) PersonalAccessTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalAccessTokenArrayOutput)
}

// PersonalAccessTokenMapInput is an input type that accepts PersonalAccessTokenMap and PersonalAccessTokenMapOutput values.
// You can construct a concrete instance of `PersonalAccessTokenMapInput` via:
//
//          PersonalAccessTokenMap{ "key": PersonalAccessTokenArgs{...} }
type PersonalAccessTokenMapInput interface {
	pulumi.Input

	ToPersonalAccessTokenMapOutput() PersonalAccessTokenMapOutput
	ToPersonalAccessTokenMapOutputWithContext(context.Context) PersonalAccessTokenMapOutput
}

type PersonalAccessTokenMap map[string]PersonalAccessTokenInput

func (PersonalAccessTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersonalAccessToken)(nil)).Elem()
}

func (i PersonalAccessTokenMap) ToPersonalAccessTokenMapOutput() PersonalAccessTokenMapOutput {
	return i.ToPersonalAccessTokenMapOutputWithContext(context.Background())
}

func (i PersonalAccessTokenMap) ToPersonalAccessTokenMapOutputWithContext(ctx context.Context) PersonalAccessTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalAccessTokenMapOutput)
}

type PersonalAccessTokenOutput struct{ *pulumi.OutputState }

func (PersonalAccessTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalAccessToken)(nil)).Elem()
}

func (o PersonalAccessTokenOutput) ToPersonalAccessTokenOutput() PersonalAccessTokenOutput {
	return o
}

func (o PersonalAccessTokenOutput) ToPersonalAccessTokenOutputWithContext(ctx context.Context) PersonalAccessTokenOutput {
	return o
}

// Expiration date of the token in the RFC3339 format
func (o PersonalAccessTokenOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringPtrOutput { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

// ID of the organization
func (o PersonalAccessTokenOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Value of the token
func (o PersonalAccessTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// ID of the user
func (o PersonalAccessTokenOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type PersonalAccessTokenArrayOutput struct{ *pulumi.OutputState }

func (PersonalAccessTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersonalAccessToken)(nil)).Elem()
}

func (o PersonalAccessTokenArrayOutput) ToPersonalAccessTokenArrayOutput() PersonalAccessTokenArrayOutput {
	return o
}

func (o PersonalAccessTokenArrayOutput) ToPersonalAccessTokenArrayOutputWithContext(ctx context.Context) PersonalAccessTokenArrayOutput {
	return o
}

func (o PersonalAccessTokenArrayOutput) Index(i pulumi.IntInput) PersonalAccessTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersonalAccessToken {
		return vs[0].([]*PersonalAccessToken)[vs[1].(int)]
	}).(PersonalAccessTokenOutput)
}

type PersonalAccessTokenMapOutput struct{ *pulumi.OutputState }

func (PersonalAccessTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersonalAccessToken)(nil)).Elem()
}

func (o PersonalAccessTokenMapOutput) ToPersonalAccessTokenMapOutput() PersonalAccessTokenMapOutput {
	return o
}

func (o PersonalAccessTokenMapOutput) ToPersonalAccessTokenMapOutputWithContext(ctx context.Context) PersonalAccessTokenMapOutput {
	return o
}

func (o PersonalAccessTokenMapOutput) MapIndex(k pulumi.StringInput) PersonalAccessTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersonalAccessToken {
		return vs[0].(map[string]*PersonalAccessToken)[vs[1].(string)]
	}).(PersonalAccessTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersonalAccessTokenInput)(nil)).Elem(), &PersonalAccessToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersonalAccessTokenArrayInput)(nil)).Elem(), PersonalAccessTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersonalAccessTokenMapInput)(nil)).Elem(), PersonalAccessTokenMap{})
	pulumi.RegisterOutputType(PersonalAccessTokenOutput{})
	pulumi.RegisterOutputType(PersonalAccessTokenArrayOutput{})
	pulumi.RegisterOutputType(PersonalAccessTokenMapOutput{})
}
