// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing a Google IdP of the organization.
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
// 		_, err := zitadel.LookupOrgIdpGoogle(ctx, &GetOrgIdpGoogleArgs{
// 			OrgId: pulumi.StringRef(data.Zitadel_org.Default.Id),
// 			Id:    "123456789012345678",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupOrgIdpGoogle(ctx *pulumi.Context, args *LookupOrgIdpGoogleArgs, opts ...pulumi.InvokeOption) (*LookupOrgIdpGoogleResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupOrgIdpGoogleResult
	err := ctx.Invoke("zitadel:index/getOrgIdpGoogle:getOrgIdpGoogle", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgIdpGoogle.
type LookupOrgIdpGoogleArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getOrgIdpGoogle.
type LookupOrgIdpGoogleResult struct {
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
}

func LookupOrgIdpGoogleOutput(ctx *pulumi.Context, args LookupOrgIdpGoogleOutputArgs, opts ...pulumi.InvokeOption) LookupOrgIdpGoogleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgIdpGoogleResult, error) {
			args := v.(LookupOrgIdpGoogleArgs)
			r, err := LookupOrgIdpGoogle(ctx, &args, opts...)
			var s LookupOrgIdpGoogleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgIdpGoogleResultOutput)
}

// A collection of arguments for invoking getOrgIdpGoogle.
type LookupOrgIdpGoogleOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (LookupOrgIdpGoogleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpGoogleArgs)(nil)).Elem()
}

// A collection of values returned by getOrgIdpGoogle.
type LookupOrgIdpGoogleResultOutput struct{ *pulumi.OutputState }

func (LookupOrgIdpGoogleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpGoogleResult)(nil)).Elem()
}

func (o LookupOrgIdpGoogleResultOutput) ToLookupOrgIdpGoogleResultOutput() LookupOrgIdpGoogleResultOutput {
	return o
}

func (o LookupOrgIdpGoogleResultOutput) ToLookupOrgIdpGoogleResultOutputWithContext(ctx context.Context) LookupOrgIdpGoogleResultOutput {
	return o
}

// client id generated by the identity provider
func (o LookupOrgIdpGoogleResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o LookupOrgIdpGoogleResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupOrgIdpGoogleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupOrgIdpGoogleResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupOrgIdpGoogleResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupOrgIdpGoogleResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupOrgIdpGoogleResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o LookupOrgIdpGoogleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupOrgIdpGoogleResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o LookupOrgIdpGoogleResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgIdpGoogleResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgIdpGoogleResultOutput{})
}
