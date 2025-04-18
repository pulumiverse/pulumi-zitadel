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

// Datasource representing a SAML IdP of the organization.
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
//			_, err := zitadel.LookupOrgIdpSaml(ctx, &zitadel.LookupOrgIdpSamlArgs{
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
func LookupOrgIdpSaml(ctx *pulumi.Context, args *LookupOrgIdpSamlArgs, opts ...pulumi.InvokeOption) (*LookupOrgIdpSamlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrgIdpSamlResult
	err := ctx.Invoke("zitadel:index/getOrgIdpSaml:getOrgIdpSaml", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgIdpSaml.
type LookupOrgIdpSamlArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getOrgIdpSaml.
type LookupOrgIdpSamlResult struct {
	// The binding
	Binding string `pulumi:"binding"`
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
	// The metadata XML as plain string
	MetadataXml string `pulumi:"metadataXml"`
	// Name of the IDP
	Name string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Whether the SAML IDP requires signed requests
	WithSignedRequest string `pulumi:"withSignedRequest"`
}

func LookupOrgIdpSamlOutput(ctx *pulumi.Context, args LookupOrgIdpSamlOutputArgs, opts ...pulumi.InvokeOption) LookupOrgIdpSamlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgIdpSamlResult, error) {
			args := v.(LookupOrgIdpSamlArgs)
			r, err := LookupOrgIdpSaml(ctx, &args, opts...)
			var s LookupOrgIdpSamlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgIdpSamlResultOutput)
}

// A collection of arguments for invoking getOrgIdpSaml.
type LookupOrgIdpSamlOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (LookupOrgIdpSamlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpSamlArgs)(nil)).Elem()
}

// A collection of values returned by getOrgIdpSaml.
type LookupOrgIdpSamlResultOutput struct{ *pulumi.OutputState }

func (LookupOrgIdpSamlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpSamlResult)(nil)).Elem()
}

func (o LookupOrgIdpSamlResultOutput) ToLookupOrgIdpSamlResultOutput() LookupOrgIdpSamlResultOutput {
	return o
}

func (o LookupOrgIdpSamlResultOutput) ToLookupOrgIdpSamlResultOutputWithContext(ctx context.Context) LookupOrgIdpSamlResultOutput {
	return o
}

func (o LookupOrgIdpSamlResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOrgIdpSamlResult] {
	return pulumix.Output[LookupOrgIdpSamlResult]{
		OutputState: o.OutputState,
	}
}

// The binding
func (o LookupOrgIdpSamlResultOutput) Binding() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) string { return v.Binding }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupOrgIdpSamlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupOrgIdpSamlResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupOrgIdpSamlResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupOrgIdpSamlResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupOrgIdpSamlResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// The metadata XML as plain string
func (o LookupOrgIdpSamlResultOutput) MetadataXml() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) string { return v.MetadataXml }).(pulumi.StringOutput)
}

// Name of the IDP
func (o LookupOrgIdpSamlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupOrgIdpSamlResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Whether the SAML IDP requires signed requests
func (o LookupOrgIdpSamlResultOutput) WithSignedRequest() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpSamlResult) string { return v.WithSignedRequest }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgIdpSamlResultOutput{})
}
