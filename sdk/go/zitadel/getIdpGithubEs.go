// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing a GitHub Enterprise IDP on the instance.
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
//			_, err := zitadel.LookupIdpGithubEs(ctx, &GetIdpGithubEsArgs{
//				Id: "177073614158299139",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIdpGithubEs(ctx *pulumi.Context, args *LookupIdpGithubEsArgs, opts ...pulumi.InvokeOption) (*LookupIdpGithubEsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupIdpGithubEsResult
	err := ctx.Invoke("zitadel:index/getIdpGithubEs:getIdpGithubEs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdpGithubEs.
type LookupIdpGithubEsArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
}

// A collection of values returned by getIdpGithubEs.
type LookupIdpGithubEsResult struct {
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
	// the providers token endpoint
	TokenEndpoint string `pulumi:"tokenEndpoint"`
	// the providers user endpoint
	UserEndpoint string `pulumi:"userEndpoint"`
}

func LookupIdpGithubEsOutput(ctx *pulumi.Context, args LookupIdpGithubEsOutputArgs, opts ...pulumi.InvokeOption) LookupIdpGithubEsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdpGithubEsResult, error) {
			args := v.(LookupIdpGithubEsArgs)
			r, err := LookupIdpGithubEs(ctx, &args, opts...)
			var s LookupIdpGithubEsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdpGithubEsResultOutput)
}

// A collection of arguments for invoking getIdpGithubEs.
type LookupIdpGithubEsOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIdpGithubEsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpGithubEsArgs)(nil)).Elem()
}

// A collection of values returned by getIdpGithubEs.
type LookupIdpGithubEsResultOutput struct{ *pulumi.OutputState }

func (LookupIdpGithubEsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpGithubEsResult)(nil)).Elem()
}

func (o LookupIdpGithubEsResultOutput) ToLookupIdpGithubEsResultOutput() LookupIdpGithubEsResultOutput {
	return o
}

func (o LookupIdpGithubEsResultOutput) ToLookupIdpGithubEsResultOutputWithContext(ctx context.Context) LookupIdpGithubEsResultOutput {
	return o
}

// the providers authorization endpoint
func (o LookupIdpGithubEsResultOutput) AuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) string { return v.AuthorizationEndpoint }).(pulumi.StringOutput)
}

// client id generated by the identity provider
func (o LookupIdpGithubEsResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o LookupIdpGithubEsResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupIdpGithubEsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupIdpGithubEsResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupIdpGithubEsResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupIdpGithubEsResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupIdpGithubEsResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o LookupIdpGithubEsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) string { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o LookupIdpGithubEsResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// the providers token endpoint
func (o LookupIdpGithubEsResultOutput) TokenEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) string { return v.TokenEndpoint }).(pulumi.StringOutput)
}

// the providers user endpoint
func (o LookupIdpGithubEsResultOutput) UserEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubEsResult) string { return v.UserEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdpGithubEsResultOutput{})
}
