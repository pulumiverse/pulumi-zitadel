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

// Datasource representing a generic OIDC IdP on the organization.
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
//			_default, err := zitadel.GetOrgOidcIdp(ctx, &zitadel.GetOrgOidcIdpArgs{
//				OrgId: pulumi.StringRef(defaultZitadelOrg.Id),
//				Id:    "123456789012345678",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("orgOidcIdp", _default)
//			return nil
//		})
//	}
//
// ```
func GetOrgOidcIdp(ctx *pulumi.Context, args *GetOrgOidcIdpArgs, opts ...pulumi.InvokeOption) (*GetOrgOidcIdpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrgOidcIdpResult
	err := ctx.Invoke("zitadel:index/getOrgOidcIdp:getOrgOidcIdp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgOidcIdp.
type GetOrgOidcIdpArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getOrgOidcIdp.
type GetOrgOidcIdpResult struct {
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
	// if true, provider information get mapped from the id token, not from the userinfo endpoint.
	IsIdTokenMapping bool `pulumi:"isIdTokenMapping"`
	// enabled if users are able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// the issuer of the idp
	Issuer string `pulumi:"issuer"`
	// Name of the IDP
	Name string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

func GetOrgOidcIdpOutput(ctx *pulumi.Context, args GetOrgOidcIdpOutputArgs, opts ...pulumi.InvokeOption) GetOrgOidcIdpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrgOidcIdpResult, error) {
			args := v.(GetOrgOidcIdpArgs)
			r, err := GetOrgOidcIdp(ctx, &args, opts...)
			var s GetOrgOidcIdpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrgOidcIdpResultOutput)
}

// A collection of arguments for invoking getOrgOidcIdp.
type GetOrgOidcIdpOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (GetOrgOidcIdpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrgOidcIdpArgs)(nil)).Elem()
}

// A collection of values returned by getOrgOidcIdp.
type GetOrgOidcIdpResultOutput struct{ *pulumi.OutputState }

func (GetOrgOidcIdpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrgOidcIdpResult)(nil)).Elem()
}

func (o GetOrgOidcIdpResultOutput) ToGetOrgOidcIdpResultOutput() GetOrgOidcIdpResultOutput {
	return o
}

func (o GetOrgOidcIdpResultOutput) ToGetOrgOidcIdpResultOutputWithContext(ctx context.Context) GetOrgOidcIdpResultOutput {
	return o
}

func (o GetOrgOidcIdpResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetOrgOidcIdpResult] {
	return pulumix.Output[GetOrgOidcIdpResult]{
		OutputState: o.OutputState,
	}
}

// client id generated by the identity provider
func (o GetOrgOidcIdpResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o GetOrgOidcIdpResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o GetOrgOidcIdpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o GetOrgOidcIdpResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o GetOrgOidcIdpResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o GetOrgOidcIdpResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// if true, provider information get mapped from the id token, not from the userinfo endpoint.
func (o GetOrgOidcIdpResultOutput) IsIdTokenMapping() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) bool { return v.IsIdTokenMapping }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o GetOrgOidcIdpResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// the issuer of the idp
func (o GetOrgOidcIdpResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// Name of the IDP
func (o GetOrgOidcIdpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o GetOrgOidcIdpResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o GetOrgOidcIdpResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrgOidcIdpResultOutput{})
}
