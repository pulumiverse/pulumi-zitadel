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

// Datasource representing a generic OAuth2 IDP of the organization.
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
//			_, err := zitadel.LookupOrgIdpOauth(ctx, &zitadel.LookupOrgIdpOauthArgs{
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
func LookupOrgIdpOauth(ctx *pulumi.Context, args *LookupOrgIdpOauthArgs, opts ...pulumi.InvokeOption) (*LookupOrgIdpOauthResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrgIdpOauthResult
	err := ctx.Invoke("zitadel:index/getOrgIdpOauth:getOrgIdpOauth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgIdpOauth.
type LookupOrgIdpOauthArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getOrgIdpOauth.
type LookupOrgIdpOauthResult struct {
	// The authorization endpoint
	AuthorizationEndpoint string `pulumi:"authorizationEndpoint"`
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// The id attribute
	IdAttribute string `pulumi:"idAttribute"`
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
	// The token endpoint
	TokenEndpoint string `pulumi:"tokenEndpoint"`
	// The user endpoint
	UserEndpoint string `pulumi:"userEndpoint"`
}

func LookupOrgIdpOauthOutput(ctx *pulumi.Context, args LookupOrgIdpOauthOutputArgs, opts ...pulumi.InvokeOption) LookupOrgIdpOauthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgIdpOauthResult, error) {
			args := v.(LookupOrgIdpOauthArgs)
			r, err := LookupOrgIdpOauth(ctx, &args, opts...)
			var s LookupOrgIdpOauthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgIdpOauthResultOutput)
}

// A collection of arguments for invoking getOrgIdpOauth.
type LookupOrgIdpOauthOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (LookupOrgIdpOauthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpOauthArgs)(nil)).Elem()
}

// A collection of values returned by getOrgIdpOauth.
type LookupOrgIdpOauthResultOutput struct{ *pulumi.OutputState }

func (LookupOrgIdpOauthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpOauthResult)(nil)).Elem()
}

func (o LookupOrgIdpOauthResultOutput) ToLookupOrgIdpOauthResultOutput() LookupOrgIdpOauthResultOutput {
	return o
}

func (o LookupOrgIdpOauthResultOutput) ToLookupOrgIdpOauthResultOutputWithContext(ctx context.Context) LookupOrgIdpOauthResultOutput {
	return o
}

func (o LookupOrgIdpOauthResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOrgIdpOauthResult] {
	return pulumix.Output[LookupOrgIdpOauthResult]{
		OutputState: o.OutputState,
	}
}

// The authorization endpoint
func (o LookupOrgIdpOauthResultOutput) AuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) string { return v.AuthorizationEndpoint }).(pulumi.StringOutput)
}

// client id generated by the identity provider
func (o LookupOrgIdpOauthResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o LookupOrgIdpOauthResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupOrgIdpOauthResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id attribute
func (o LookupOrgIdpOauthResultOutput) IdAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) string { return v.IdAttribute }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupOrgIdpOauthResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupOrgIdpOauthResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupOrgIdpOauthResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupOrgIdpOauthResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o LookupOrgIdpOauthResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupOrgIdpOauthResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o LookupOrgIdpOauthResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// The token endpoint
func (o LookupOrgIdpOauthResultOutput) TokenEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) string { return v.TokenEndpoint }).(pulumi.StringOutput)
}

// The user endpoint
func (o LookupOrgIdpOauthResultOutput) UserEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpOauthResult) string { return v.UserEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgIdpOauthResultOutput{})
}
