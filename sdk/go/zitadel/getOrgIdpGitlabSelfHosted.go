// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing a GitLab Self Hosted IdP of the organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-zitadel/sdk/go/zitadel"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/vavsab/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.LookupOrgIdpGitlabSelfHosted(ctx, &GetOrgIdpGitlabSelfHostedArgs{
// 			Id: "177073614158299139",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupOrgIdpGitlabSelfHosted(ctx *pulumi.Context, args *LookupOrgIdpGitlabSelfHostedArgs, opts ...pulumi.InvokeOption) (*LookupOrgIdpGitlabSelfHostedResult, error) {
	var rv LookupOrgIdpGitlabSelfHostedResult
	err := ctx.Invoke("zitadel:index/getOrgIdpGitlabSelfHosted:getOrgIdpGitlabSelfHosted", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgIdpGitlabSelfHosted.
type LookupOrgIdpGitlabSelfHostedArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getOrgIdpGitlabSelfHosted.
type LookupOrgIdpGitlabSelfHostedResult struct {
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
	// the providers issuer
	Issuer string `pulumi:"issuer"`
	// Name of the IDP
	Name string `pulumi:"name"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

func LookupOrgIdpGitlabSelfHostedOutput(ctx *pulumi.Context, args LookupOrgIdpGitlabSelfHostedOutputArgs, opts ...pulumi.InvokeOption) LookupOrgIdpGitlabSelfHostedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgIdpGitlabSelfHostedResult, error) {
			args := v.(LookupOrgIdpGitlabSelfHostedArgs)
			r, err := LookupOrgIdpGitlabSelfHosted(ctx, &args, opts...)
			var s LookupOrgIdpGitlabSelfHostedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgIdpGitlabSelfHostedResultOutput)
}

// A collection of arguments for invoking getOrgIdpGitlabSelfHosted.
type LookupOrgIdpGitlabSelfHostedOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the organization
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (LookupOrgIdpGitlabSelfHostedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpGitlabSelfHostedArgs)(nil)).Elem()
}

// A collection of values returned by getOrgIdpGitlabSelfHosted.
type LookupOrgIdpGitlabSelfHostedResultOutput struct{ *pulumi.OutputState }

func (LookupOrgIdpGitlabSelfHostedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpGitlabSelfHostedResult)(nil)).Elem()
}

func (o LookupOrgIdpGitlabSelfHostedResultOutput) ToLookupOrgIdpGitlabSelfHostedResultOutput() LookupOrgIdpGitlabSelfHostedResultOutput {
	return o
}

func (o LookupOrgIdpGitlabSelfHostedResultOutput) ToLookupOrgIdpGitlabSelfHostedResultOutputWithContext(ctx context.Context) LookupOrgIdpGitlabSelfHostedResultOutput {
	return o
}

// client id generated by the identity provider
func (o LookupOrgIdpGitlabSelfHostedResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o LookupOrgIdpGitlabSelfHostedResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupOrgIdpGitlabSelfHostedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupOrgIdpGitlabSelfHostedResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupOrgIdpGitlabSelfHostedResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupOrgIdpGitlabSelfHostedResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupOrgIdpGitlabSelfHostedResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// the providers issuer
func (o LookupOrgIdpGitlabSelfHostedResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// Name of the IDP
func (o LookupOrgIdpGitlabSelfHostedResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupOrgIdpGitlabSelfHostedResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o LookupOrgIdpGitlabSelfHostedResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgIdpGitlabSelfHostedResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgIdpGitlabSelfHostedResultOutput{})
}
