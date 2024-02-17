// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Datasource representing multiple OIDC applications belonging to a project.
func GetApplicationOidcs(ctx *pulumi.Context, args *GetApplicationOidcsArgs, opts ...pulumi.InvokeOption) (*GetApplicationOidcsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApplicationOidcsResult
	err := ctx.Invoke("zitadel:index/getApplicationOidcs:getApplicationOidcs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationOidcs.
type GetApplicationOidcsArgs struct {
	// Name of the application
	Name string `pulumi:"name"`
	// Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod *string `pulumi:"nameMethod"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getApplicationOidcs.
type GetApplicationOidcsResult struct {
	// A set of all IDs.
	AppIds []string `pulumi:"appIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the application
	Name string `pulumi:"name"`
	// Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod *string `pulumi:"nameMethod"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
}

func GetApplicationOidcsOutput(ctx *pulumi.Context, args GetApplicationOidcsOutputArgs, opts ...pulumi.InvokeOption) GetApplicationOidcsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationOidcsResult, error) {
			args := v.(GetApplicationOidcsArgs)
			r, err := GetApplicationOidcs(ctx, &args, opts...)
			var s GetApplicationOidcsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationOidcsResultOutput)
}

// A collection of arguments for invoking getApplicationOidcs.
type GetApplicationOidcsOutputArgs struct {
	// Name of the application
	Name pulumi.StringInput `pulumi:"name"`
	// Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod pulumi.StringPtrInput `pulumi:"nameMethod"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetApplicationOidcsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationOidcsArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationOidcs.
type GetApplicationOidcsResultOutput struct{ *pulumi.OutputState }

func (GetApplicationOidcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationOidcsResult)(nil)).Elem()
}

func (o GetApplicationOidcsResultOutput) ToGetApplicationOidcsResultOutput() GetApplicationOidcsResultOutput {
	return o
}

func (o GetApplicationOidcsResultOutput) ToGetApplicationOidcsResultOutputWithContext(ctx context.Context) GetApplicationOidcsResultOutput {
	return o
}

func (o GetApplicationOidcsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetApplicationOidcsResult] {
	return pulumix.Output[GetApplicationOidcsResult]{
		OutputState: o.OutputState,
	}
}

// A set of all IDs.
func (o GetApplicationOidcsResultOutput) AppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApplicationOidcsResult) []string { return v.AppIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApplicationOidcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationOidcsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the application
func (o GetApplicationOidcsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationOidcsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
func (o GetApplicationOidcsResultOutput) NameMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationOidcsResult) *string { return v.NameMethod }).(pulumi.StringPtrOutput)
}

// ID of the organization
func (o GetApplicationOidcsResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationOidcsResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// ID of the project
func (o GetApplicationOidcsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationOidcsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationOidcsResultOutput{})
}