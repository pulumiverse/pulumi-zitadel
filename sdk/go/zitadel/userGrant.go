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

// Resource representing the authorization given to a user directly, including the given roles.
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
//			_, err := zitadel.NewUserGrant(ctx, "userGrant", &zitadel.UserGrantArgs{
//				ProjectId: pulumi.Any(zitadel_project.Project.Id),
//				OrgId:     pulumi.Any(zitadel_org.Org.Id),
//				RoleKeys: pulumi.StringArray{
//					pulumi.String("key"),
//				},
//				UserId: pulumi.Any(zitadel_human_user.Granted_human_user.Id),
=======
//			_, err := zitadel.NewUserGrant(ctx, "default", &zitadel.UserGrantArgs{
//				ProjectId: pulumi.Any(data.Zitadel_project.Default.Id),
//				OrgId:     pulumi.Any(data.Zitadel_org.Default.Id),
//				RoleKeys: pulumi.StringArray{
//					pulumi.String("super-user"),
//				},
//				UserId: pulumi.Any(data.Zitadel_human_user.Default.Id),
>>>>>>> origin/master
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
// terraform The resource can be imported using the ID format `<flow_type:trigger_type[:org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/userGrant:UserGrant imported '123456789012345678:123456789012345678:123456789012345678'
//
>>>>>>> origin/master
// ```
type UserGrant struct {
	pulumi.CustomResourceState

	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// ID of the granted project
	ProjectGrantId pulumi.StringPtrOutput `pulumi:"projectGrantId"`
	// ID of the project
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// List of roles granted
	RoleKeys pulumi.StringArrayOutput `pulumi:"roleKeys"`
	// ID of the user
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserGrant registers a new resource with the given unique name, arguments, and options.
func NewUserGrant(ctx *pulumi.Context,
	name string, args *UserGrantArgs, opts ...pulumi.ResourceOption) (*UserGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserGrant
	err := ctx.RegisterResource("zitadel:index/userGrant:UserGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGrant gets an existing UserGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGrantState, opts ...pulumi.ResourceOption) (*UserGrant, error) {
	var resource UserGrant
	err := ctx.ReadResource("zitadel:index/userGrant:UserGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGrant resources.
type userGrantState struct {
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the granted project
	ProjectGrantId *string `pulumi:"projectGrantId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
	// List of roles granted
	RoleKeys []string `pulumi:"roleKeys"`
	// ID of the user
	UserId *string `pulumi:"userId"`
}

type UserGrantState struct {
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the granted project
	ProjectGrantId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
	// List of roles granted
	RoleKeys pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringPtrInput
}

func (UserGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGrantState)(nil)).Elem()
}

type userGrantArgs struct {
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the granted project
	ProjectGrantId *string `pulumi:"projectGrantId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
	// List of roles granted
	RoleKeys []string `pulumi:"roleKeys"`
	// ID of the user
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserGrant resource.
type UserGrantArgs struct {
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the granted project
	ProjectGrantId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
	// List of roles granted
	RoleKeys pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringInput
}

func (UserGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGrantArgs)(nil)).Elem()
}

type UserGrantInput interface {
	pulumi.Input

	ToUserGrantOutput() UserGrantOutput
	ToUserGrantOutputWithContext(ctx context.Context) UserGrantOutput
}

func (*UserGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGrant)(nil)).Elem()
}

func (i *UserGrant) ToUserGrantOutput() UserGrantOutput {
	return i.ToUserGrantOutputWithContext(context.Background())
}

func (i *UserGrant) ToUserGrantOutputWithContext(ctx context.Context) UserGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGrantOutput)
}

func (i *UserGrant) ToOutput(ctx context.Context) pulumix.Output[*UserGrant] {
	return pulumix.Output[*UserGrant]{
		OutputState: i.ToUserGrantOutputWithContext(ctx).OutputState,
	}
}

// UserGrantArrayInput is an input type that accepts UserGrantArray and UserGrantArrayOutput values.
// You can construct a concrete instance of `UserGrantArrayInput` via:
//
//	UserGrantArray{ UserGrantArgs{...} }
type UserGrantArrayInput interface {
	pulumi.Input

	ToUserGrantArrayOutput() UserGrantArrayOutput
	ToUserGrantArrayOutputWithContext(context.Context) UserGrantArrayOutput
}

type UserGrantArray []UserGrantInput

func (UserGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGrant)(nil)).Elem()
}

func (i UserGrantArray) ToUserGrantArrayOutput() UserGrantArrayOutput {
	return i.ToUserGrantArrayOutputWithContext(context.Background())
}

func (i UserGrantArray) ToUserGrantArrayOutputWithContext(ctx context.Context) UserGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGrantArrayOutput)
}

func (i UserGrantArray) ToOutput(ctx context.Context) pulumix.Output[[]*UserGrant] {
	return pulumix.Output[[]*UserGrant]{
		OutputState: i.ToUserGrantArrayOutputWithContext(ctx).OutputState,
	}
}

// UserGrantMapInput is an input type that accepts UserGrantMap and UserGrantMapOutput values.
// You can construct a concrete instance of `UserGrantMapInput` via:
//
//	UserGrantMap{ "key": UserGrantArgs{...} }
type UserGrantMapInput interface {
	pulumi.Input

	ToUserGrantMapOutput() UserGrantMapOutput
	ToUserGrantMapOutputWithContext(context.Context) UserGrantMapOutput
}

type UserGrantMap map[string]UserGrantInput

func (UserGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGrant)(nil)).Elem()
}

func (i UserGrantMap) ToUserGrantMapOutput() UserGrantMapOutput {
	return i.ToUserGrantMapOutputWithContext(context.Background())
}

func (i UserGrantMap) ToUserGrantMapOutputWithContext(ctx context.Context) UserGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGrantMapOutput)
}

func (i UserGrantMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserGrant] {
	return pulumix.Output[map[string]*UserGrant]{
		OutputState: i.ToUserGrantMapOutputWithContext(ctx).OutputState,
	}
}

type UserGrantOutput struct{ *pulumi.OutputState }

func (UserGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGrant)(nil)).Elem()
}

func (o UserGrantOutput) ToUserGrantOutput() UserGrantOutput {
	return o
}

func (o UserGrantOutput) ToUserGrantOutputWithContext(ctx context.Context) UserGrantOutput {
	return o
}

func (o UserGrantOutput) ToOutput(ctx context.Context) pulumix.Output[*UserGrant] {
	return pulumix.Output[*UserGrant]{
		OutputState: o.OutputState,
	}
}

// ID of the organization
func (o UserGrantOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserGrant) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// ID of the granted project
func (o UserGrantOutput) ProjectGrantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserGrant) pulumi.StringPtrOutput { return v.ProjectGrantId }).(pulumi.StringPtrOutput)
}

// ID of the project
func (o UserGrantOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserGrant) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// List of roles granted
func (o UserGrantOutput) RoleKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserGrant) pulumi.StringArrayOutput { return v.RoleKeys }).(pulumi.StringArrayOutput)
}

// ID of the user
func (o UserGrantOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGrant) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type UserGrantArrayOutput struct{ *pulumi.OutputState }

func (UserGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGrant)(nil)).Elem()
}

func (o UserGrantArrayOutput) ToUserGrantArrayOutput() UserGrantArrayOutput {
	return o
}

func (o UserGrantArrayOutput) ToUserGrantArrayOutputWithContext(ctx context.Context) UserGrantArrayOutput {
	return o
}

func (o UserGrantArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*UserGrant] {
	return pulumix.Output[[]*UserGrant]{
		OutputState: o.OutputState,
	}
}

func (o UserGrantArrayOutput) Index(i pulumi.IntInput) UserGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserGrant {
		return vs[0].([]*UserGrant)[vs[1].(int)]
	}).(UserGrantOutput)
}

type UserGrantMapOutput struct{ *pulumi.OutputState }

func (UserGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGrant)(nil)).Elem()
}

func (o UserGrantMapOutput) ToUserGrantMapOutput() UserGrantMapOutput {
	return o
}

func (o UserGrantMapOutput) ToUserGrantMapOutputWithContext(ctx context.Context) UserGrantMapOutput {
	return o
}

func (o UserGrantMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserGrant] {
	return pulumix.Output[map[string]*UserGrant]{
		OutputState: o.OutputState,
	}
}

func (o UserGrantMapOutput) MapIndex(k pulumi.StringInput) UserGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserGrant {
		return vs[0].(map[string]*UserGrant)[vs[1].(string)]
	}).(UserGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserGrantInput)(nil)).Elem(), &UserGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGrantArrayInput)(nil)).Elem(), UserGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGrantMapInput)(nil)).Elem(), UserGrantMap{})
	pulumi.RegisterOutputType(UserGrantOutput{})
	pulumi.RegisterOutputType(UserGrantArrayOutput{})
	pulumi.RegisterOutputType(UserGrantMapOutput{})
}
