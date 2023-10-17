// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the membership of a user on an project, defined with the given role.
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
//			_, err := zitadel.NewProjectMember(ctx, "projectMember", &zitadel.ProjectMemberArgs{
//				OrgId:     pulumi.Any(zitadel_org.Org.Id),
//				ProjectId: pulumi.Any(zitadel_project.Project.Id),
//				UserId:    pulumi.Any(zitadel_human_user.Human_user.Id),
//				Roles: pulumi.StringArray{
//					pulumi.String("PROJECT_OWNER"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ProjectMember struct {
	pulumi.CustomResourceState

	// ID of the organization which owns the resource
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// List of roles granted
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// ID of the user
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewProjectMember registers a new resource with the given unique name, arguments, and options.
func NewProjectMember(ctx *pulumi.Context,
	name string, args *ProjectMemberArgs, opts ...pulumi.ResourceOption) (*ProjectMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ProjectMember
	err := ctx.RegisterResource("zitadel:index/projectMember:ProjectMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMember gets an existing ProjectMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMemberState, opts ...pulumi.ResourceOption) (*ProjectMember, error) {
	var resource ProjectMember
	err := ctx.ReadResource("zitadel:index/projectMember:ProjectMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMember resources.
type projectMemberState struct {
	// ID of the organization which owns the resource
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
	// List of roles granted
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId *string `pulumi:"userId"`
}

type ProjectMemberState struct {
	// ID of the organization which owns the resource
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
	// List of roles granted
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringPtrInput
}

func (ProjectMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMemberState)(nil)).Elem()
}

type projectMemberArgs struct {
	// ID of the organization which owns the resource
	OrgId string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
	// List of roles granted
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a ProjectMember resource.
type ProjectMemberArgs struct {
	// ID of the organization which owns the resource
	OrgId pulumi.StringInput
	// ID of the project
	ProjectId pulumi.StringInput
	// List of roles granted
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringInput
}

func (ProjectMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMemberArgs)(nil)).Elem()
}

type ProjectMemberInput interface {
	pulumi.Input

	ToProjectMemberOutput() ProjectMemberOutput
	ToProjectMemberOutputWithContext(ctx context.Context) ProjectMemberOutput
}

func (*ProjectMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMember)(nil)).Elem()
}

func (i *ProjectMember) ToProjectMemberOutput() ProjectMemberOutput {
	return i.ToProjectMemberOutputWithContext(context.Background())
}

func (i *ProjectMember) ToProjectMemberOutputWithContext(ctx context.Context) ProjectMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberOutput)
}

// ProjectMemberArrayInput is an input type that accepts ProjectMemberArray and ProjectMemberArrayOutput values.
// You can construct a concrete instance of `ProjectMemberArrayInput` via:
//
//	ProjectMemberArray{ ProjectMemberArgs{...} }
type ProjectMemberArrayInput interface {
	pulumi.Input

	ToProjectMemberArrayOutput() ProjectMemberArrayOutput
	ToProjectMemberArrayOutputWithContext(context.Context) ProjectMemberArrayOutput
}

type ProjectMemberArray []ProjectMemberInput

func (ProjectMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMember)(nil)).Elem()
}

func (i ProjectMemberArray) ToProjectMemberArrayOutput() ProjectMemberArrayOutput {
	return i.ToProjectMemberArrayOutputWithContext(context.Background())
}

func (i ProjectMemberArray) ToProjectMemberArrayOutputWithContext(ctx context.Context) ProjectMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberArrayOutput)
}

// ProjectMemberMapInput is an input type that accepts ProjectMemberMap and ProjectMemberMapOutput values.
// You can construct a concrete instance of `ProjectMemberMapInput` via:
//
//	ProjectMemberMap{ "key": ProjectMemberArgs{...} }
type ProjectMemberMapInput interface {
	pulumi.Input

	ToProjectMemberMapOutput() ProjectMemberMapOutput
	ToProjectMemberMapOutputWithContext(context.Context) ProjectMemberMapOutput
}

type ProjectMemberMap map[string]ProjectMemberInput

func (ProjectMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMember)(nil)).Elem()
}

func (i ProjectMemberMap) ToProjectMemberMapOutput() ProjectMemberMapOutput {
	return i.ToProjectMemberMapOutputWithContext(context.Background())
}

func (i ProjectMemberMap) ToProjectMemberMapOutputWithContext(ctx context.Context) ProjectMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberMapOutput)
}

type ProjectMemberOutput struct{ *pulumi.OutputState }

func (ProjectMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMember)(nil)).Elem()
}

func (o ProjectMemberOutput) ToProjectMemberOutput() ProjectMemberOutput {
	return o
}

func (o ProjectMemberOutput) ToProjectMemberOutputWithContext(ctx context.Context) ProjectMemberOutput {
	return o
}

// ID of the organization which owns the resource
func (o ProjectMemberOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMember) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// ID of the project
func (o ProjectMemberOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMember) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// List of roles granted
func (o ProjectMemberOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectMember) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// ID of the user
func (o ProjectMemberOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMember) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type ProjectMemberArrayOutput struct{ *pulumi.OutputState }

func (ProjectMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMember)(nil)).Elem()
}

func (o ProjectMemberArrayOutput) ToProjectMemberArrayOutput() ProjectMemberArrayOutput {
	return o
}

func (o ProjectMemberArrayOutput) ToProjectMemberArrayOutputWithContext(ctx context.Context) ProjectMemberArrayOutput {
	return o
}

func (o ProjectMemberArrayOutput) Index(i pulumi.IntInput) ProjectMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectMember {
		return vs[0].([]*ProjectMember)[vs[1].(int)]
	}).(ProjectMemberOutput)
}

type ProjectMemberMapOutput struct{ *pulumi.OutputState }

func (ProjectMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMember)(nil)).Elem()
}

func (o ProjectMemberMapOutput) ToProjectMemberMapOutput() ProjectMemberMapOutput {
	return o
}

func (o ProjectMemberMapOutput) ToProjectMemberMapOutputWithContext(ctx context.Context) ProjectMemberMapOutput {
	return o
}

func (o ProjectMemberMapOutput) MapIndex(k pulumi.StringInput) ProjectMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectMember {
		return vs[0].(map[string]*ProjectMember)[vs[1].(string)]
	}).(ProjectMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberInput)(nil)).Elem(), &ProjectMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberArrayInput)(nil)).Elem(), ProjectMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberMapInput)(nil)).Elem(), ProjectMemberMap{})
	pulumi.RegisterOutputType(ProjectMemberOutput{})
	pulumi.RegisterOutputType(ProjectMemberArrayOutput{})
	pulumi.RegisterOutputType(ProjectMemberMapOutput{})
}
