// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing the project roles, which can be given as authorizations to users.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-zitadel/sdk/go/zitadel"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			projectRoleProjectRole, err := zitadel.LookupProjectRole(ctx, &GetProjectRoleArgs{
//				OrgId:     data.Zitadel_org.Org.Id,
//				ProjectId: data.Zitadel_project.Project.Id,
//				RoleKey:   "key",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("projectRole", projectRoleProjectRole)
//			return nil
//		})
//	}
//
// ```
func LookupProjectRole(ctx *pulumi.Context, args *LookupProjectRoleArgs, opts ...pulumi.InvokeOption) (*LookupProjectRoleResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupProjectRoleResult
	err := ctx.Invoke("zitadel:index/getProjectRole:getProjectRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectRole.
type LookupProjectRoleArgs struct {
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
	// Key used for project role
	RoleKey string `pulumi:"roleKey"`
}

// A collection of values returned by getProjectRole.
type LookupProjectRoleResult struct {
	// Name used for project role
	DisplayName string `pulumi:"displayName"`
	// Group used for project role
	Group string `pulumi:"group"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
	// Key used for project role
	RoleKey string `pulumi:"roleKey"`
}

func LookupProjectRoleOutput(ctx *pulumi.Context, args LookupProjectRoleOutputArgs, opts ...pulumi.InvokeOption) LookupProjectRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectRoleResult, error) {
			args := v.(LookupProjectRoleArgs)
			r, err := LookupProjectRole(ctx, &args, opts...)
			var s LookupProjectRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectRoleResultOutput)
}

// A collection of arguments for invoking getProjectRole.
type LookupProjectRoleOutputArgs struct {
	// ID of the organization
	OrgId pulumi.StringInput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// Key used for project role
	RoleKey pulumi.StringInput `pulumi:"roleKey"`
}

func (LookupProjectRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectRoleArgs)(nil)).Elem()
}

// A collection of values returned by getProjectRole.
type LookupProjectRoleResultOutput struct{ *pulumi.OutputState }

func (LookupProjectRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectRoleResult)(nil)).Elem()
}

func (o LookupProjectRoleResultOutput) ToLookupProjectRoleResultOutput() LookupProjectRoleResultOutput {
	return o
}

func (o LookupProjectRoleResultOutput) ToLookupProjectRoleResultOutputWithContext(ctx context.Context) LookupProjectRoleResultOutput {
	return o
}

// Name used for project role
func (o LookupProjectRoleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectRoleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Group used for project role
func (o LookupProjectRoleResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectRoleResult) string { return v.Group }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupProjectRoleResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectRoleResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// ID of the project
func (o LookupProjectRoleResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectRoleResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Key used for project role
func (o LookupProjectRoleResultOutput) RoleKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectRoleResult) string { return v.RoleKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectRoleResultOutput{})
}
