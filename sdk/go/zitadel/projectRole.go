// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the project roles, which can be given as authorizations to users.
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
// 		_, err := zitadel.NewProjectRole(ctx, "projectRole", &zitadel.ProjectRoleArgs{
// 			OrgId:       pulumi.Any(zitadel_org.Org.Id),
// 			ProjectId:   pulumi.Any(zitadel_project.Project.Id),
// 			RoleKey:     pulumi.String("key"),
// 			DisplayName: pulumi.String("display_name2"),
// 			Group:       pulumi.String("role_group"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProjectRole struct {
	pulumi.CustomResourceState

	// Name used for project role
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Group used for project role
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// ID of the organization
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Key used for project role
	RoleKey pulumi.StringOutput `pulumi:"roleKey"`
}

// NewProjectRole registers a new resource with the given unique name, arguments, and options.
func NewProjectRole(ctx *pulumi.Context,
	name string, args *ProjectRoleArgs, opts ...pulumi.ResourceOption) (*ProjectRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.RoleKey == nil {
		return nil, errors.New("invalid value for required argument 'RoleKey'")
	}
	var resource ProjectRole
	err := ctx.RegisterResource("zitadel:index/projectRole:ProjectRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectRole gets an existing ProjectRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectRoleState, opts ...pulumi.ResourceOption) (*ProjectRole, error) {
	var resource ProjectRole
	err := ctx.ReadResource("zitadel:index/projectRole:ProjectRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectRole resources.
type projectRoleState struct {
	// Name used for project role
	DisplayName *string `pulumi:"displayName"`
	// Group used for project role
	Group *string `pulumi:"group"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
	// Key used for project role
	RoleKey *string `pulumi:"roleKey"`
}

type ProjectRoleState struct {
	// Name used for project role
	DisplayName pulumi.StringPtrInput
	// Group used for project role
	Group pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
	// Key used for project role
	RoleKey pulumi.StringPtrInput
}

func (ProjectRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectRoleState)(nil)).Elem()
}

type projectRoleArgs struct {
	// Name used for project role
	DisplayName string `pulumi:"displayName"`
	// Group used for project role
	Group *string `pulumi:"group"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
	// Key used for project role
	RoleKey string `pulumi:"roleKey"`
}

// The set of arguments for constructing a ProjectRole resource.
type ProjectRoleArgs struct {
	// Name used for project role
	DisplayName pulumi.StringInput
	// Group used for project role
	Group pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringInput
	// ID of the project
	ProjectId pulumi.StringInput
	// Key used for project role
	RoleKey pulumi.StringInput
}

func (ProjectRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectRoleArgs)(nil)).Elem()
}

type ProjectRoleInput interface {
	pulumi.Input

	ToProjectRoleOutput() ProjectRoleOutput
	ToProjectRoleOutputWithContext(ctx context.Context) ProjectRoleOutput
}

func (*ProjectRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectRole)(nil)).Elem()
}

func (i *ProjectRole) ToProjectRoleOutput() ProjectRoleOutput {
	return i.ToProjectRoleOutputWithContext(context.Background())
}

func (i *ProjectRole) ToProjectRoleOutputWithContext(ctx context.Context) ProjectRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRoleOutput)
}

// ProjectRoleArrayInput is an input type that accepts ProjectRoleArray and ProjectRoleArrayOutput values.
// You can construct a concrete instance of `ProjectRoleArrayInput` via:
//
//          ProjectRoleArray{ ProjectRoleArgs{...} }
type ProjectRoleArrayInput interface {
	pulumi.Input

	ToProjectRoleArrayOutput() ProjectRoleArrayOutput
	ToProjectRoleArrayOutputWithContext(context.Context) ProjectRoleArrayOutput
}

type ProjectRoleArray []ProjectRoleInput

func (ProjectRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectRole)(nil)).Elem()
}

func (i ProjectRoleArray) ToProjectRoleArrayOutput() ProjectRoleArrayOutput {
	return i.ToProjectRoleArrayOutputWithContext(context.Background())
}

func (i ProjectRoleArray) ToProjectRoleArrayOutputWithContext(ctx context.Context) ProjectRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRoleArrayOutput)
}

// ProjectRoleMapInput is an input type that accepts ProjectRoleMap and ProjectRoleMapOutput values.
// You can construct a concrete instance of `ProjectRoleMapInput` via:
//
//          ProjectRoleMap{ "key": ProjectRoleArgs{...} }
type ProjectRoleMapInput interface {
	pulumi.Input

	ToProjectRoleMapOutput() ProjectRoleMapOutput
	ToProjectRoleMapOutputWithContext(context.Context) ProjectRoleMapOutput
}

type ProjectRoleMap map[string]ProjectRoleInput

func (ProjectRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectRole)(nil)).Elem()
}

func (i ProjectRoleMap) ToProjectRoleMapOutput() ProjectRoleMapOutput {
	return i.ToProjectRoleMapOutputWithContext(context.Background())
}

func (i ProjectRoleMap) ToProjectRoleMapOutputWithContext(ctx context.Context) ProjectRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRoleMapOutput)
}

type ProjectRoleOutput struct{ *pulumi.OutputState }

func (ProjectRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectRole)(nil)).Elem()
}

func (o ProjectRoleOutput) ToProjectRoleOutput() ProjectRoleOutput {
	return o
}

func (o ProjectRoleOutput) ToProjectRoleOutputWithContext(ctx context.Context) ProjectRoleOutput {
	return o
}

// Name used for project role
func (o ProjectRoleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectRole) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Group used for project role
func (o ProjectRoleOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectRole) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

// ID of the organization
func (o ProjectRoleOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectRole) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// ID of the project
func (o ProjectRoleOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectRole) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Key used for project role
func (o ProjectRoleOutput) RoleKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectRole) pulumi.StringOutput { return v.RoleKey }).(pulumi.StringOutput)
}

type ProjectRoleArrayOutput struct{ *pulumi.OutputState }

func (ProjectRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectRole)(nil)).Elem()
}

func (o ProjectRoleArrayOutput) ToProjectRoleArrayOutput() ProjectRoleArrayOutput {
	return o
}

func (o ProjectRoleArrayOutput) ToProjectRoleArrayOutputWithContext(ctx context.Context) ProjectRoleArrayOutput {
	return o
}

func (o ProjectRoleArrayOutput) Index(i pulumi.IntInput) ProjectRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectRole {
		return vs[0].([]*ProjectRole)[vs[1].(int)]
	}).(ProjectRoleOutput)
}

type ProjectRoleMapOutput struct{ *pulumi.OutputState }

func (ProjectRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectRole)(nil)).Elem()
}

func (o ProjectRoleMapOutput) ToProjectRoleMapOutput() ProjectRoleMapOutput {
	return o
}

func (o ProjectRoleMapOutput) ToProjectRoleMapOutputWithContext(ctx context.Context) ProjectRoleMapOutput {
	return o
}

func (o ProjectRoleMapOutput) MapIndex(k pulumi.StringInput) ProjectRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectRole {
		return vs[0].(map[string]*ProjectRole)[vs[1].(string)]
	}).(ProjectRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRoleInput)(nil)).Elem(), &ProjectRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRoleArrayInput)(nil)).Elem(), ProjectRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRoleMapInput)(nil)).Elem(), ProjectRoleMap{})
	pulumi.RegisterOutputType(ProjectRoleOutput{})
	pulumi.RegisterOutputType(ProjectRoleArrayOutput{})
	pulumi.RegisterOutputType(ProjectRoleMapOutput{})
}
