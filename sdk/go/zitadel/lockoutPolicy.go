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

// Resource representing the custom lockout policy of an organization.
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
<<<<<<< HEAD
//			_, err := zitadel.NewLockoutPolicy(ctx, "lockoutPolicy", &zitadel.LockoutPolicyArgs{
//				OrgId:               pulumi.Any(zitadel_org.Org.Id),
=======
//			_, err := zitadel.NewLockoutPolicy(ctx, "default", &zitadel.LockoutPolicyArgs{
//				OrgId:               pulumi.Any(data.Zitadel_org.Default.Id),
>>>>>>> origin/master
//				MaxPasswordAttempts: pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
<<<<<<< HEAD
=======
// ```
//
// ## Import
//
// terraform The resource can be imported using the ID format `<[org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/lockoutPolicy:LockoutPolicy imported '123456789012345678'
//
>>>>>>> origin/master
// ```
type LockoutPolicy struct {
	pulumi.CustomResourceState

	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
	MaxPasswordAttempts pulumi.IntOutput `pulumi:"maxPasswordAttempts"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
}

// NewLockoutPolicy registers a new resource with the given unique name, arguments, and options.
func NewLockoutPolicy(ctx *pulumi.Context,
	name string, args *LockoutPolicyArgs, opts ...pulumi.ResourceOption) (*LockoutPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxPasswordAttempts == nil {
		return nil, errors.New("invalid value for required argument 'MaxPasswordAttempts'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LockoutPolicy
	err := ctx.RegisterResource("zitadel:index/lockoutPolicy:LockoutPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLockoutPolicy gets an existing LockoutPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLockoutPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LockoutPolicyState, opts ...pulumi.ResourceOption) (*LockoutPolicy, error) {
	var resource LockoutPolicy
	err := ctx.ReadResource("zitadel:index/lockoutPolicy:LockoutPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LockoutPolicy resources.
type lockoutPolicyState struct {
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
	MaxPasswordAttempts *int `pulumi:"maxPasswordAttempts"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
}

type LockoutPolicyState struct {
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
	MaxPasswordAttempts pulumi.IntPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
}

func (LockoutPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lockoutPolicyState)(nil)).Elem()
}

type lockoutPolicyArgs struct {
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
	MaxPasswordAttempts int `pulumi:"maxPasswordAttempts"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
}

// The set of arguments for constructing a LockoutPolicy resource.
type LockoutPolicyArgs struct {
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
	MaxPasswordAttempts pulumi.IntInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
}

func (LockoutPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lockoutPolicyArgs)(nil)).Elem()
}

type LockoutPolicyInput interface {
	pulumi.Input

	ToLockoutPolicyOutput() LockoutPolicyOutput
	ToLockoutPolicyOutputWithContext(ctx context.Context) LockoutPolicyOutput
}

func (*LockoutPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**LockoutPolicy)(nil)).Elem()
}

func (i *LockoutPolicy) ToLockoutPolicyOutput() LockoutPolicyOutput {
	return i.ToLockoutPolicyOutputWithContext(context.Background())
}

func (i *LockoutPolicy) ToLockoutPolicyOutputWithContext(ctx context.Context) LockoutPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LockoutPolicyOutput)
}

func (i *LockoutPolicy) ToOutput(ctx context.Context) pulumix.Output[*LockoutPolicy] {
	return pulumix.Output[*LockoutPolicy]{
		OutputState: i.ToLockoutPolicyOutputWithContext(ctx).OutputState,
	}
}

// LockoutPolicyArrayInput is an input type that accepts LockoutPolicyArray and LockoutPolicyArrayOutput values.
// You can construct a concrete instance of `LockoutPolicyArrayInput` via:
//
//	LockoutPolicyArray{ LockoutPolicyArgs{...} }
type LockoutPolicyArrayInput interface {
	pulumi.Input

	ToLockoutPolicyArrayOutput() LockoutPolicyArrayOutput
	ToLockoutPolicyArrayOutputWithContext(context.Context) LockoutPolicyArrayOutput
}

type LockoutPolicyArray []LockoutPolicyInput

func (LockoutPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LockoutPolicy)(nil)).Elem()
}

func (i LockoutPolicyArray) ToLockoutPolicyArrayOutput() LockoutPolicyArrayOutput {
	return i.ToLockoutPolicyArrayOutputWithContext(context.Background())
}

func (i LockoutPolicyArray) ToLockoutPolicyArrayOutputWithContext(ctx context.Context) LockoutPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LockoutPolicyArrayOutput)
}

func (i LockoutPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*LockoutPolicy] {
	return pulumix.Output[[]*LockoutPolicy]{
		OutputState: i.ToLockoutPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// LockoutPolicyMapInput is an input type that accepts LockoutPolicyMap and LockoutPolicyMapOutput values.
// You can construct a concrete instance of `LockoutPolicyMapInput` via:
//
//	LockoutPolicyMap{ "key": LockoutPolicyArgs{...} }
type LockoutPolicyMapInput interface {
	pulumi.Input

	ToLockoutPolicyMapOutput() LockoutPolicyMapOutput
	ToLockoutPolicyMapOutputWithContext(context.Context) LockoutPolicyMapOutput
}

type LockoutPolicyMap map[string]LockoutPolicyInput

func (LockoutPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LockoutPolicy)(nil)).Elem()
}

func (i LockoutPolicyMap) ToLockoutPolicyMapOutput() LockoutPolicyMapOutput {
	return i.ToLockoutPolicyMapOutputWithContext(context.Background())
}

func (i LockoutPolicyMap) ToLockoutPolicyMapOutputWithContext(ctx context.Context) LockoutPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LockoutPolicyMapOutput)
}

func (i LockoutPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LockoutPolicy] {
	return pulumix.Output[map[string]*LockoutPolicy]{
		OutputState: i.ToLockoutPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type LockoutPolicyOutput struct{ *pulumi.OutputState }

func (LockoutPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LockoutPolicy)(nil)).Elem()
}

func (o LockoutPolicyOutput) ToLockoutPolicyOutput() LockoutPolicyOutput {
	return o
}

func (o LockoutPolicyOutput) ToLockoutPolicyOutputWithContext(ctx context.Context) LockoutPolicyOutput {
	return o
}

func (o LockoutPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*LockoutPolicy] {
	return pulumix.Output[*LockoutPolicy]{
		OutputState: o.OutputState,
	}
}

// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
func (o LockoutPolicyOutput) MaxPasswordAttempts() pulumi.IntOutput {
	return o.ApplyT(func(v *LockoutPolicy) pulumi.IntOutput { return v.MaxPasswordAttempts }).(pulumi.IntOutput)
}

// ID of the organization
func (o LockoutPolicyOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LockoutPolicy) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

type LockoutPolicyArrayOutput struct{ *pulumi.OutputState }

func (LockoutPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LockoutPolicy)(nil)).Elem()
}

func (o LockoutPolicyArrayOutput) ToLockoutPolicyArrayOutput() LockoutPolicyArrayOutput {
	return o
}

func (o LockoutPolicyArrayOutput) ToLockoutPolicyArrayOutputWithContext(ctx context.Context) LockoutPolicyArrayOutput {
	return o
}

func (o LockoutPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LockoutPolicy] {
	return pulumix.Output[[]*LockoutPolicy]{
		OutputState: o.OutputState,
	}
}

func (o LockoutPolicyArrayOutput) Index(i pulumi.IntInput) LockoutPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LockoutPolicy {
		return vs[0].([]*LockoutPolicy)[vs[1].(int)]
	}).(LockoutPolicyOutput)
}

type LockoutPolicyMapOutput struct{ *pulumi.OutputState }

func (LockoutPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LockoutPolicy)(nil)).Elem()
}

func (o LockoutPolicyMapOutput) ToLockoutPolicyMapOutput() LockoutPolicyMapOutput {
	return o
}

func (o LockoutPolicyMapOutput) ToLockoutPolicyMapOutputWithContext(ctx context.Context) LockoutPolicyMapOutput {
	return o
}

func (o LockoutPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LockoutPolicy] {
	return pulumix.Output[map[string]*LockoutPolicy]{
		OutputState: o.OutputState,
	}
}

func (o LockoutPolicyMapOutput) MapIndex(k pulumi.StringInput) LockoutPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LockoutPolicy {
		return vs[0].(map[string]*LockoutPolicy)[vs[1].(string)]
	}).(LockoutPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LockoutPolicyInput)(nil)).Elem(), &LockoutPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LockoutPolicyArrayInput)(nil)).Elem(), LockoutPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LockoutPolicyMapInput)(nil)).Elem(), LockoutPolicyMap{})
	pulumi.RegisterOutputType(LockoutPolicyOutput{})
	pulumi.RegisterOutputType(LockoutPolicyArrayOutput{})
	pulumi.RegisterOutputType(LockoutPolicyMapOutput{})
}
