// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/scoretechnologies/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Datasource representing a GitHub IDP on the instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/scoretechnologies/pulumi-zitadel/sdk/go/zitadel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zitadel.LookupIdpGithub(ctx, &zitadel.LookupIdpGithubArgs{
//				Id: "123456789012345678",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIdpGithub(ctx *pulumi.Context, args *LookupIdpGithubArgs, opts ...pulumi.InvokeOption) (*LookupIdpGithubResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdpGithubResult
	err := ctx.Invoke("zitadel:index/getIdpGithub:getIdpGithub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdpGithub.
type LookupIdpGithubArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
}

// A collection of values returned by getIdpGithub.
type LookupIdpGithubResult struct {
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

func LookupIdpGithubOutput(ctx *pulumi.Context, args LookupIdpGithubOutputArgs, opts ...pulumi.InvokeOption) LookupIdpGithubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdpGithubResult, error) {
			args := v.(LookupIdpGithubArgs)
			r, err := LookupIdpGithub(ctx, &args, opts...)
			var s LookupIdpGithubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdpGithubResultOutput)
}

// A collection of arguments for invoking getIdpGithub.
type LookupIdpGithubOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIdpGithubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpGithubArgs)(nil)).Elem()
}

// A collection of values returned by getIdpGithub.
type LookupIdpGithubResultOutput struct{ *pulumi.OutputState }

func (LookupIdpGithubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpGithubResult)(nil)).Elem()
}

func (o LookupIdpGithubResultOutput) ToLookupIdpGithubResultOutput() LookupIdpGithubResultOutput {
	return o
}

func (o LookupIdpGithubResultOutput) ToLookupIdpGithubResultOutputWithContext(ctx context.Context) LookupIdpGithubResultOutput {
	return o
}

func (o LookupIdpGithubResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIdpGithubResult] {
	return pulumix.Output[LookupIdpGithubResult]{
		OutputState: o.OutputState,
	}
}

// client id generated by the identity provider
func (o LookupIdpGithubResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o LookupIdpGithubResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupIdpGithubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupIdpGithubResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupIdpGithubResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupIdpGithubResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupIdpGithubResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o LookupIdpGithubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) string { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o LookupIdpGithubResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIdpGithubResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdpGithubResultOutput{})
}
