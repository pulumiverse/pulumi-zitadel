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

// Resource representing the grant of a project to a different organization, also containing the available roles which can be given to the members of the projectgrant.
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
//			_, err := zitadel.NewProjectGrant(ctx, "projectGrant", &zitadel.ProjectGrantArgs{
//				OrgId:        pulumi.Any(zitadel_org.Org.Id),
//				ProjectId:    pulumi.Any(zitadel_project.Project.Id),
//				GrantedOrgId: pulumi.Any(zitadel_org.Grantedorg.Id),
//				RoleKeys: pulumi.StringArray{
//					pulumi.Any(zitadel_project_role.Project_role.Role_key),
=======
//			_, err := zitadel.NewProjectGrant(ctx, "default", &zitadel.ProjectGrantArgs{
//				OrgId:        pulumi.Any(data.Zitadel_org.Default.Id),
//				ProjectId:    pulumi.Any(data.Zitadel_project.Default.Id),
//				GrantedOrgId: pulumi.Any(data.Zitadel_org.Granted_org.Id),
//				RoleKeys: pulumi.StringArray{
//					pulumi.String("super-user"),
>>>>>>> origin/master
//				},
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
// terraform The resource can be imported using the ID format `<id:project_id[:org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/projectGrant:ProjectGrant imported '123456789012345678:123456789012345678:123456789012345678'
//
>>>>>>> origin/master
// ```
type ProjectGrant struct {
	pulumi.CustomResourceState

	// ID of the organization granted the project
	GrantedOrgId pulumi.StringOutput `pulumi:"grantedOrgId"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// List of roles granted
	RoleKeys pulumi.StringArrayOutput `pulumi:"roleKeys"`
}

// NewProjectGrant registers a new resource with the given unique name, arguments, and options.
func NewProjectGrant(ctx *pulumi.Context,
	name string, args *ProjectGrantArgs, opts ...pulumi.ResourceOption) (*ProjectGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GrantedOrgId == nil {
		return nil, errors.New("invalid value for required argument 'GrantedOrgId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectGrant
	err := ctx.RegisterResource("zitadel:index/projectGrant:ProjectGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectGrant gets an existing ProjectGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectGrantState, opts ...pulumi.ResourceOption) (*ProjectGrant, error) {
	var resource ProjectGrant
	err := ctx.ReadResource("zitadel:index/projectGrant:ProjectGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectGrant resources.
type projectGrantState struct {
	// ID of the organization granted the project
	GrantedOrgId *string `pulumi:"grantedOrgId"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
	// List of roles granted
	RoleKeys []string `pulumi:"roleKeys"`
}

type ProjectGrantState struct {
	// ID of the organization granted the project
	GrantedOrgId pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
	// List of roles granted
	RoleKeys pulumi.StringArrayInput
}

func (ProjectGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectGrantState)(nil)).Elem()
}

type projectGrantArgs struct {
	// ID of the organization granted the project
	GrantedOrgId string `pulumi:"grantedOrgId"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
	// List of roles granted
	RoleKeys []string `pulumi:"roleKeys"`
}

// The set of arguments for constructing a ProjectGrant resource.
type ProjectGrantArgs struct {
	// ID of the organization granted the project
	GrantedOrgId pulumi.StringInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringInput
	// List of roles granted
	RoleKeys pulumi.StringArrayInput
}

func (ProjectGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectGrantArgs)(nil)).Elem()
}

type ProjectGrantInput interface {
	pulumi.Input

	ToProjectGrantOutput() ProjectGrantOutput
	ToProjectGrantOutputWithContext(ctx context.Context) ProjectGrantOutput
}

func (*ProjectGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectGrant)(nil)).Elem()
}

func (i *ProjectGrant) ToProjectGrantOutput() ProjectGrantOutput {
	return i.ToProjectGrantOutputWithContext(context.Background())
}

func (i *ProjectGrant) ToProjectGrantOutputWithContext(ctx context.Context) ProjectGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectGrantOutput)
}

func (i *ProjectGrant) ToOutput(ctx context.Context) pulumix.Output[*ProjectGrant] {
	return pulumix.Output[*ProjectGrant]{
		OutputState: i.ToProjectGrantOutputWithContext(ctx).OutputState,
	}
}

// ProjectGrantArrayInput is an input type that accepts ProjectGrantArray and ProjectGrantArrayOutput values.
// You can construct a concrete instance of `ProjectGrantArrayInput` via:
//
//	ProjectGrantArray{ ProjectGrantArgs{...} }
type ProjectGrantArrayInput interface {
	pulumi.Input

	ToProjectGrantArrayOutput() ProjectGrantArrayOutput
	ToProjectGrantArrayOutputWithContext(context.Context) ProjectGrantArrayOutput
}

type ProjectGrantArray []ProjectGrantInput

func (ProjectGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectGrant)(nil)).Elem()
}

func (i ProjectGrantArray) ToProjectGrantArrayOutput() ProjectGrantArrayOutput {
	return i.ToProjectGrantArrayOutputWithContext(context.Background())
}

func (i ProjectGrantArray) ToProjectGrantArrayOutputWithContext(ctx context.Context) ProjectGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectGrantArrayOutput)
}

func (i ProjectGrantArray) ToOutput(ctx context.Context) pulumix.Output[[]*ProjectGrant] {
	return pulumix.Output[[]*ProjectGrant]{
		OutputState: i.ToProjectGrantArrayOutputWithContext(ctx).OutputState,
	}
}

// ProjectGrantMapInput is an input type that accepts ProjectGrantMap and ProjectGrantMapOutput values.
// You can construct a concrete instance of `ProjectGrantMapInput` via:
//
//	ProjectGrantMap{ "key": ProjectGrantArgs{...} }
type ProjectGrantMapInput interface {
	pulumi.Input

	ToProjectGrantMapOutput() ProjectGrantMapOutput
	ToProjectGrantMapOutputWithContext(context.Context) ProjectGrantMapOutput
}

type ProjectGrantMap map[string]ProjectGrantInput

func (ProjectGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectGrant)(nil)).Elem()
}

func (i ProjectGrantMap) ToProjectGrantMapOutput() ProjectGrantMapOutput {
	return i.ToProjectGrantMapOutputWithContext(context.Background())
}

func (i ProjectGrantMap) ToProjectGrantMapOutputWithContext(ctx context.Context) ProjectGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectGrantMapOutput)
}

func (i ProjectGrantMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ProjectGrant] {
	return pulumix.Output[map[string]*ProjectGrant]{
		OutputState: i.ToProjectGrantMapOutputWithContext(ctx).OutputState,
	}
}

type ProjectGrantOutput struct{ *pulumi.OutputState }

func (ProjectGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectGrant)(nil)).Elem()
}

func (o ProjectGrantOutput) ToProjectGrantOutput() ProjectGrantOutput {
	return o
}

func (o ProjectGrantOutput) ToProjectGrantOutputWithContext(ctx context.Context) ProjectGrantOutput {
	return o
}

func (o ProjectGrantOutput) ToOutput(ctx context.Context) pulumix.Output[*ProjectGrant] {
	return pulumix.Output[*ProjectGrant]{
		OutputState: o.OutputState,
	}
}

// ID of the organization granted the project
func (o ProjectGrantOutput) GrantedOrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectGrant) pulumi.StringOutput { return v.GrantedOrgId }).(pulumi.StringOutput)
}

// ID of the organization
func (o ProjectGrantOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectGrant) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// ID of the project
func (o ProjectGrantOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectGrant) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// List of roles granted
func (o ProjectGrantOutput) RoleKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectGrant) pulumi.StringArrayOutput { return v.RoleKeys }).(pulumi.StringArrayOutput)
}

type ProjectGrantArrayOutput struct{ *pulumi.OutputState }

func (ProjectGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectGrant)(nil)).Elem()
}

func (o ProjectGrantArrayOutput) ToProjectGrantArrayOutput() ProjectGrantArrayOutput {
	return o
}

func (o ProjectGrantArrayOutput) ToProjectGrantArrayOutputWithContext(ctx context.Context) ProjectGrantArrayOutput {
	return o
}

func (o ProjectGrantArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ProjectGrant] {
	return pulumix.Output[[]*ProjectGrant]{
		OutputState: o.OutputState,
	}
}

func (o ProjectGrantArrayOutput) Index(i pulumi.IntInput) ProjectGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectGrant {
		return vs[0].([]*ProjectGrant)[vs[1].(int)]
	}).(ProjectGrantOutput)
}

type ProjectGrantMapOutput struct{ *pulumi.OutputState }

func (ProjectGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectGrant)(nil)).Elem()
}

func (o ProjectGrantMapOutput) ToProjectGrantMapOutput() ProjectGrantMapOutput {
	return o
}

func (o ProjectGrantMapOutput) ToProjectGrantMapOutputWithContext(ctx context.Context) ProjectGrantMapOutput {
	return o
}

func (o ProjectGrantMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ProjectGrant] {
	return pulumix.Output[map[string]*ProjectGrant]{
		OutputState: o.OutputState,
	}
}

func (o ProjectGrantMapOutput) MapIndex(k pulumi.StringInput) ProjectGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectGrant {
		return vs[0].(map[string]*ProjectGrant)[vs[1].(string)]
	}).(ProjectGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectGrantInput)(nil)).Elem(), &ProjectGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectGrantArrayInput)(nil)).Elem(), ProjectGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectGrantMapInput)(nil)).Elem(), ProjectGrantMap{})
	pulumi.RegisterOutputType(ProjectGrantOutput{})
	pulumi.RegisterOutputType(ProjectGrantArrayOutput{})
	pulumi.RegisterOutputType(ProjectGrantMapOutput{})
}
