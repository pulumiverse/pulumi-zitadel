// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing a GitLab IDP on the instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-zitadel/sdk/go/zitadel"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.LookupIdpGitlab(ctx, &GetIdpGitlabArgs{
// 			Id: "123456789012345678",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupIdpGitlab(ctx *pulumi.Context, args *LookupIdpGitlabArgs, opts ...pulumi.InvokeOption) (*LookupIdpGitlabResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupIdpGitlabResult
	err := ctx.Invoke("zitadel:index/getIdpGitlab:getIdpGitlab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdpGitlab.
type LookupIdpGitlabArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
}

// A collection of values returned by getIdpGitlab.
type LookupIdpGitlabResult struct {
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
}

func LookupIdpGitlabOutput(ctx *pulumi.Context, args LookupIdpGitlabOutputArgs, opts ...pulumi.InvokeOption) LookupIdpGitlabResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdpGitlabResult, error) {
			args := v.(LookupIdpGitlabArgs)
			r, err := LookupIdpGitlab(ctx, &args, opts...)
			var s LookupIdpGitlabResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdpGitlabResultOutput)
}

// A collection of arguments for invoking getIdpGitlab.
type LookupIdpGitlabOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIdpGitlabOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpGitlabArgs)(nil)).Elem()
}

// A collection of values returned by getIdpGitlab.
type LookupIdpGitlabResultOutput struct{ *pulumi.OutputState }

func (LookupIdpGitlabResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpGitlabResult)(nil)).Elem()
}

func (o LookupIdpGitlabResultOutput) ToLookupIdpGitlabResultOutput() LookupIdpGitlabResultOutput {
	return o
}

func (o LookupIdpGitlabResultOutput) ToLookupIdpGitlabResultOutputWithContext(ctx context.Context) LookupIdpGitlabResultOutput {
	return o
}

// client id generated by the identity provider
func (o LookupIdpGitlabResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o LookupIdpGitlabResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupIdpGitlabResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupIdpGitlabResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupIdpGitlabResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupIdpGitlabResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupIdpGitlabResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o LookupIdpGitlabResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) string { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o LookupIdpGitlabResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIdpGitlabResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdpGitlabResultOutput{})
}
