// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the custom notification policy of an organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.NewNotificationPolicy(ctx, "notificationPolicy", &zitadel.NotificationPolicyArgs{
// 			OrgId:          pulumi.Any(zitadel_org.Org.Id),
// 			PasswordChange: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NotificationPolicy struct {
	pulumi.CustomResourceState

	// Id for the organization
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Send notification if a user changes his password
	PasswordChange pulumi.BoolOutput `pulumi:"passwordChange"`
}

// NewNotificationPolicy registers a new resource with the given unique name, arguments, and options.
func NewNotificationPolicy(ctx *pulumi.Context,
	name string, args *NotificationPolicyArgs, opts ...pulumi.ResourceOption) (*NotificationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.PasswordChange == nil {
		return nil, errors.New("invalid value for required argument 'PasswordChange'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NotificationPolicy
	err := ctx.RegisterResource("zitadel:index/notificationPolicy:NotificationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationPolicy gets an existing NotificationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationPolicyState, opts ...pulumi.ResourceOption) (*NotificationPolicy, error) {
	var resource NotificationPolicy
	err := ctx.ReadResource("zitadel:index/notificationPolicy:NotificationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationPolicy resources.
type notificationPolicyState struct {
	// Id for the organization
	OrgId *string `pulumi:"orgId"`
	// Send notification if a user changes his password
	PasswordChange *bool `pulumi:"passwordChange"`
}

type NotificationPolicyState struct {
	// Id for the organization
	OrgId pulumi.StringPtrInput
	// Send notification if a user changes his password
	PasswordChange pulumi.BoolPtrInput
}

func (NotificationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationPolicyState)(nil)).Elem()
}

type notificationPolicyArgs struct {
	// Id for the organization
	OrgId string `pulumi:"orgId"`
	// Send notification if a user changes his password
	PasswordChange bool `pulumi:"passwordChange"`
}

// The set of arguments for constructing a NotificationPolicy resource.
type NotificationPolicyArgs struct {
	// Id for the organization
	OrgId pulumi.StringInput
	// Send notification if a user changes his password
	PasswordChange pulumi.BoolInput
}

func (NotificationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationPolicyArgs)(nil)).Elem()
}

type NotificationPolicyInput interface {
	pulumi.Input

	ToNotificationPolicyOutput() NotificationPolicyOutput
	ToNotificationPolicyOutputWithContext(ctx context.Context) NotificationPolicyOutput
}

func (*NotificationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationPolicy)(nil)).Elem()
}

func (i *NotificationPolicy) ToNotificationPolicyOutput() NotificationPolicyOutput {
	return i.ToNotificationPolicyOutputWithContext(context.Background())
}

func (i *NotificationPolicy) ToNotificationPolicyOutputWithContext(ctx context.Context) NotificationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPolicyOutput)
}

// NotificationPolicyArrayInput is an input type that accepts NotificationPolicyArray and NotificationPolicyArrayOutput values.
// You can construct a concrete instance of `NotificationPolicyArrayInput` via:
//
//          NotificationPolicyArray{ NotificationPolicyArgs{...} }
type NotificationPolicyArrayInput interface {
	pulumi.Input

	ToNotificationPolicyArrayOutput() NotificationPolicyArrayOutput
	ToNotificationPolicyArrayOutputWithContext(context.Context) NotificationPolicyArrayOutput
}

type NotificationPolicyArray []NotificationPolicyInput

func (NotificationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationPolicy)(nil)).Elem()
}

func (i NotificationPolicyArray) ToNotificationPolicyArrayOutput() NotificationPolicyArrayOutput {
	return i.ToNotificationPolicyArrayOutputWithContext(context.Background())
}

func (i NotificationPolicyArray) ToNotificationPolicyArrayOutputWithContext(ctx context.Context) NotificationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPolicyArrayOutput)
}

// NotificationPolicyMapInput is an input type that accepts NotificationPolicyMap and NotificationPolicyMapOutput values.
// You can construct a concrete instance of `NotificationPolicyMapInput` via:
//
//          NotificationPolicyMap{ "key": NotificationPolicyArgs{...} }
type NotificationPolicyMapInput interface {
	pulumi.Input

	ToNotificationPolicyMapOutput() NotificationPolicyMapOutput
	ToNotificationPolicyMapOutputWithContext(context.Context) NotificationPolicyMapOutput
}

type NotificationPolicyMap map[string]NotificationPolicyInput

func (NotificationPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationPolicy)(nil)).Elem()
}

func (i NotificationPolicyMap) ToNotificationPolicyMapOutput() NotificationPolicyMapOutput {
	return i.ToNotificationPolicyMapOutputWithContext(context.Background())
}

func (i NotificationPolicyMap) ToNotificationPolicyMapOutputWithContext(ctx context.Context) NotificationPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPolicyMapOutput)
}

type NotificationPolicyOutput struct{ *pulumi.OutputState }

func (NotificationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationPolicy)(nil)).Elem()
}

func (o NotificationPolicyOutput) ToNotificationPolicyOutput() NotificationPolicyOutput {
	return o
}

func (o NotificationPolicyOutput) ToNotificationPolicyOutputWithContext(ctx context.Context) NotificationPolicyOutput {
	return o
}

// Id for the organization
func (o NotificationPolicyOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationPolicy) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Send notification if a user changes his password
func (o NotificationPolicyOutput) PasswordChange() pulumi.BoolOutput {
	return o.ApplyT(func(v *NotificationPolicy) pulumi.BoolOutput { return v.PasswordChange }).(pulumi.BoolOutput)
}

type NotificationPolicyArrayOutput struct{ *pulumi.OutputState }

func (NotificationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationPolicy)(nil)).Elem()
}

func (o NotificationPolicyArrayOutput) ToNotificationPolicyArrayOutput() NotificationPolicyArrayOutput {
	return o
}

func (o NotificationPolicyArrayOutput) ToNotificationPolicyArrayOutputWithContext(ctx context.Context) NotificationPolicyArrayOutput {
	return o
}

func (o NotificationPolicyArrayOutput) Index(i pulumi.IntInput) NotificationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NotificationPolicy {
		return vs[0].([]*NotificationPolicy)[vs[1].(int)]
	}).(NotificationPolicyOutput)
}

type NotificationPolicyMapOutput struct{ *pulumi.OutputState }

func (NotificationPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationPolicy)(nil)).Elem()
}

func (o NotificationPolicyMapOutput) ToNotificationPolicyMapOutput() NotificationPolicyMapOutput {
	return o
}

func (o NotificationPolicyMapOutput) ToNotificationPolicyMapOutputWithContext(ctx context.Context) NotificationPolicyMapOutput {
	return o
}

func (o NotificationPolicyMapOutput) MapIndex(k pulumi.StringInput) NotificationPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NotificationPolicy {
		return vs[0].(map[string]*NotificationPolicy)[vs[1].(string)]
	}).(NotificationPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationPolicyInput)(nil)).Elem(), &NotificationPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationPolicyArrayInput)(nil)).Elem(), NotificationPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationPolicyMapInput)(nil)).Elem(), NotificationPolicyMap{})
	pulumi.RegisterOutputType(NotificationPolicyOutput{})
	pulumi.RegisterOutputType(NotificationPolicyArrayOutput{})
	pulumi.RegisterOutputType(NotificationPolicyMapOutput{})
}
