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

// Datasource representing a GitHub Enterprise IdP of the organization.
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
//			_, err := zitadel.LookupOrgIdpGithubEs(ctx, &zitadel.LookupOrgIdpGithubEsArgs{
//				OrgId: pulumi.StringRef(defaultZitadelOrg.Id),
//				Id:    "123456789012345678",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOrgIdpGithubEs(ctx *pulumi.Context, args *LookupOrgIdpGithubEsArgs, opts ...pulumi.InvokeOption) (*LookupOrgIdpGithubEsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrgIdpGithubEsResult
	err := ctx.Invoke("zitadel:index/getOrgIdpGithubEs:getOrgIdpGithubEs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgIdpGithubEs.
type LookupOrgIdpGithubEsArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getOrgIdpGithubEs.
type LookupOrgIdpGithubEsResult struct {
	// the providers authorization endpoint
	AuthorizationEndpoint string `pulumi:"authorizationEndpoint"`
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// enabled if a new account in ZITADEL are created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enabled if a the ZITADEL account fields are updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enabled if users are able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enabled if users are able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
	// the providers token endpoint
	TokenEndpoint string `pulumi:"tokenEndpoint"`
	// the providers user endpoint
	UserEndpoint string `pulumi:"userEndpoint"`
}

func LookupOrgIdpGithubEsOutput(ctx *pulumi.Context, args LookupOrgIdpGithubEsOutputArgs, opts ...pulumi.InvokeOption) LookupOrgIdpGithubEsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgIdpGithubEsResult, error) {
			args := v.(LookupOrgIdpGithubEsArgs)
			r, err := LookupOrgIdpGithubEs(ctx, &args, opts...)
			var s LookupOrgIdpGithubEsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgIdpGithubEsResultOutput)
}

// A collection of arguments for invoking getOrgIdpGithubEs.
type LookupOrgIdpGithubEsOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (LookupOrgIdpGithubEsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpGithubEsArgs)(nil)).Elem()
}

// A collection of values returned by getOrgIdpGithubEs.
type LookupOrgIdpGithubEsResultOutput struct{ *pulumi.OutputState }

func (LookupOrgIdpGithubEsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpGithubEsResult)(nil)).Elem()
}

func (o LookupOrgIdpGithubEsResultOutput) ToLookupOrgIdpGithubEsResultOutput() LookupOrgIdpGithubEsResultOutput {
	return o
}

func (o LookupOrgIdpGithubEsResultOutput) ToLookupOrgIdpGithubEsResultOutputWithContext(ctx context.Context) LookupOrgIdpGithubEsResultOutput {
	return o
}

func (o LookupOrgIdpGithubEsResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOrgIdpGithubEsResult] {
	return pulumix.Output[LookupOrgIdpGithubEsResult]{
		OutputState: o.OutputState,
	}
}

// the providers authorization endpoint
func (o LookupOrgIdpGithubEsResultOutput) AuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) string { return v.AuthorizationEndpoint }).(pulumi.StringOutput)
}

// client id generated by the identity provider
func (o LookupOrgIdpGithubEsResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o LookupOrgIdpGithubEsResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupOrgIdpGithubEsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupOrgIdpGithubEsResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupOrgIdpGithubEsResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupOrgIdpGithubEsResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupOrgIdpGithubEsResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o LookupOrgIdpGithubEsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupOrgIdpGithubEsResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o LookupOrgIdpGithubEsResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// the providers token endpoint
func (o LookupOrgIdpGithubEsResultOutput) TokenEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) string { return v.TokenEndpoint }).(pulumi.StringOutput)
}

// the providers user endpoint
func (o LookupOrgIdpGithubEsResultOutput) UserEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGithubEsResult) string { return v.UserEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgIdpGithubEsResultOutput{})
}
