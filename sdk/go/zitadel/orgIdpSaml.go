// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Resource representing a SAML IdP on the organization.
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
//			_, err := zitadel.NewOrgIdpSaml(ctx, "default", &zitadel.OrgIdpSamlArgs{
//				OrgId:             pulumi.Any(defaultZitadelOrg.Id),
//				Name:              pulumi.String("LDAP"),
//				Binding:           pulumi.String("SAML_BINDING_POST"),
//				WithSignedRequest: pulumi.Bool(true),
//				IsLinkingAllowed:  pulumi.Bool(false),
//				IsCreationAllowed: pulumi.Bool(true),
//				IsAutoCreation:    pulumi.Bool(false),
//				IsAutoUpdate:      pulumi.Bool(true),
//				MetadataXml: pulumi.String(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
//
// <md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://saml.example.com/entityid" validUntil="2034-05-15T14:21:58.979Z">
//
//	<md:IDPSSODescriptor WantAuthnRequestsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
//	  <md:KeyDescriptor use="signing">
//	    <ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
//	      <ds:X509Data>
//	        <ds:X509Certificate>MIIC4jCCAcoCCQC33wnybT5QZDANBgkqhkiG9w0BAQsFADAyMQswCQYDVQQGEwJV
//
// SzEPMA0GA1UECgwGQm94eUhRMRIwEAYDVQQDDAlNb2NrIFNBTUwwIBcNMjIwMjI4
// MjE0NjM4WhgPMzAyMTA3MDEyMTQ2MzhaMDIxCzAJBgNVBAYTAlVLMQ8wDQYDVQQK
// DAZCb3h5SFExEjAQBgNVBAMMCU1vY2sgU0FNTDCCASIwDQYJKoZIhvcNAQEBBQAD
// ggEPADCCAQoCggEBALGfYettMsct1T6tVUwTudNJH5Pnb9GGnkXi9Zw/e6x45DD0
// RuRONbFlJ2T4RjAE/uG+AjXxXQ8o2SZfb9+GgmCHuTJFNgHoZ1nFVXCmb/Hg8Hpd
// 4vOAGXndixaReOiq3EH5XvpMjMkJ3+8+9VYMzMZOjkgQtAqO36eAFFfNKX7dTj3V
// pwLkvz6/KFCq8OAwY+AUi4eZm5J57D31GzjHwfjH9WTeX0MyndmnNB1qV75qQR3b
// 2/W5sGHRv+9AarggJkF+ptUkXoLtVA51wcfYm6hILptpde5FQC8RWY1YrswBWAEZ
// NfyrR4JeSweElNHg4NVOs4TwGjOPwWGqzTfgTlECAwEAATANBgkqhkiG9w0BAQsF
// AAOCAQEAAYRlYflSXAWoZpFfwNiCQVE5d9zZ0DPzNdWhAybXcTyMf0z5mDf6FWBW
// 5Gyoi9u3EMEDnzLcJNkwJAAc39Apa4I2/tml+Jy29dk8bTyX6m93ngmCgdLh5Za4
// khuU3AM3L63g7VexCuO7kwkjh/+LqdcIXsVGO6XDfu2QOs1Xpe9zIzLpwm/RNYeX
// UjbSj5ce/jekpAw7qyVVL4xOyh8AtUW1ek3wIw1MJvEgEPt0d16oshWJpoS1OT8L
// r/22SvYEo3EmSGdTVGgk3x3s+A0qWAqTcyjr7Q4s/GKYRFfomGwz0TZ4Iw1ZN99M
// m0eo2USlSRTVl7QHRTuiuSThHpLKQQ==</ds:X509Certificate>
//
//	      </ds:X509Data>
//	    </ds:KeyInfo>
//	  </md:KeyDescriptor>
//	  <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat>
//	  <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://mocksaml.com/api/saml/sso"/>
//	  <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://mocksaml.com/api/saml/sso"/>
//	</md:IDPSSODescriptor>
//
// </md:EntityDescriptor>
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Loading the XML Metadata
//
// If you don't want to pass the XML metadata inline, you have plenty of options. For example:
// - localFile Data Source
// - http Data Source
// - terracurlRequest Data Source
// - ...
//
// ## Import
//
// bash The resource can be imported using the ID format `<id[:org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/orgIdpSaml:OrgIdpSaml imported '123456789012345678:123456789012345678'
//
// ```
type OrgIdpSaml struct {
	pulumi.CustomResourceState

	// The binding, supported values: SAML*BINDING*UNSPECIFIED, SAML*BINDING*POST, SAML*BINDING*REDIRECT, SAML*BINDING*ARTIFACT
	Binding pulumi.StringPtrOutput `pulumi:"binding"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolOutput `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolOutput `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolOutput `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolOutput `pulumi:"isLinkingAllowed"`
	// The metadata XML as plain string
	MetadataXml pulumi.StringOutput `pulumi:"metadataXml"`
	// Name of the IDP
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Whether the SAML IDP requires signed requests
	WithSignedRequest pulumi.BoolPtrOutput `pulumi:"withSignedRequest"`
}

// NewOrgIdpSaml registers a new resource with the given unique name, arguments, and options.
func NewOrgIdpSaml(ctx *pulumi.Context,
	name string, args *OrgIdpSamlArgs, opts ...pulumi.ResourceOption) (*OrgIdpSaml, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IsAutoCreation == nil {
		return nil, errors.New("invalid value for required argument 'IsAutoCreation'")
	}
	if args.IsAutoUpdate == nil {
		return nil, errors.New("invalid value for required argument 'IsAutoUpdate'")
	}
	if args.IsCreationAllowed == nil {
		return nil, errors.New("invalid value for required argument 'IsCreationAllowed'")
	}
	if args.IsLinkingAllowed == nil {
		return nil, errors.New("invalid value for required argument 'IsLinkingAllowed'")
	}
	if args.MetadataXml == nil {
		return nil, errors.New("invalid value for required argument 'MetadataXml'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgIdpSaml
	err := ctx.RegisterResource("zitadel:index/orgIdpSaml:OrgIdpSaml", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgIdpSaml gets an existing OrgIdpSaml resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgIdpSaml(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgIdpSamlState, opts ...pulumi.ResourceOption) (*OrgIdpSaml, error) {
	var resource OrgIdpSaml
	err := ctx.ReadResource("zitadel:index/orgIdpSaml:OrgIdpSaml", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgIdpSaml resources.
type orgIdpSamlState struct {
	// The binding, supported values: SAML*BINDING*UNSPECIFIED, SAML*BINDING*POST, SAML*BINDING*REDIRECT, SAML*BINDING*ARTIFACT
	Binding *string `pulumi:"binding"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation *bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate *bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed *bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed *bool `pulumi:"isLinkingAllowed"`
	// The metadata XML as plain string
	MetadataXml *string `pulumi:"metadataXml"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Whether the SAML IDP requires signed requests
	WithSignedRequest *bool `pulumi:"withSignedRequest"`
}

type OrgIdpSamlState struct {
	// The binding, supported values: SAML*BINDING*UNSPECIFIED, SAML*BINDING*POST, SAML*BINDING*REDIRECT, SAML*BINDING*ARTIFACT
	Binding pulumi.StringPtrInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolPtrInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolPtrInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolPtrInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolPtrInput
	// The metadata XML as plain string
	MetadataXml pulumi.StringPtrInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Whether the SAML IDP requires signed requests
	WithSignedRequest pulumi.BoolPtrInput
}

func (OrgIdpSamlState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpSamlState)(nil)).Elem()
}

type orgIdpSamlArgs struct {
	// The binding, supported values: SAML*BINDING*UNSPECIFIED, SAML*BINDING*POST, SAML*BINDING*REDIRECT, SAML*BINDING*ARTIFACT
	Binding *string `pulumi:"binding"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// The metadata XML as plain string
	MetadataXml string `pulumi:"metadataXml"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Whether the SAML IDP requires signed requests
	WithSignedRequest *bool `pulumi:"withSignedRequest"`
}

// The set of arguments for constructing a OrgIdpSaml resource.
type OrgIdpSamlArgs struct {
	// The binding, supported values: SAML*BINDING*UNSPECIFIED, SAML*BINDING*POST, SAML*BINDING*REDIRECT, SAML*BINDING*ARTIFACT
	Binding pulumi.StringPtrInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolInput
	// The metadata XML as plain string
	MetadataXml pulumi.StringInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Whether the SAML IDP requires signed requests
	WithSignedRequest pulumi.BoolPtrInput
}

func (OrgIdpSamlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpSamlArgs)(nil)).Elem()
}

type OrgIdpSamlInput interface {
	pulumi.Input

	ToOrgIdpSamlOutput() OrgIdpSamlOutput
	ToOrgIdpSamlOutputWithContext(ctx context.Context) OrgIdpSamlOutput
}

func (*OrgIdpSaml) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpSaml)(nil)).Elem()
}

func (i *OrgIdpSaml) ToOrgIdpSamlOutput() OrgIdpSamlOutput {
	return i.ToOrgIdpSamlOutputWithContext(context.Background())
}

func (i *OrgIdpSaml) ToOrgIdpSamlOutputWithContext(ctx context.Context) OrgIdpSamlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpSamlOutput)
}

func (i *OrgIdpSaml) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpSaml] {
	return pulumix.Output[*OrgIdpSaml]{
		OutputState: i.ToOrgIdpSamlOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpSamlArrayInput is an input type that accepts OrgIdpSamlArray and OrgIdpSamlArrayOutput values.
// You can construct a concrete instance of `OrgIdpSamlArrayInput` via:
//
//	OrgIdpSamlArray{ OrgIdpSamlArgs{...} }
type OrgIdpSamlArrayInput interface {
	pulumi.Input

	ToOrgIdpSamlArrayOutput() OrgIdpSamlArrayOutput
	ToOrgIdpSamlArrayOutputWithContext(context.Context) OrgIdpSamlArrayOutput
}

type OrgIdpSamlArray []OrgIdpSamlInput

func (OrgIdpSamlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpSaml)(nil)).Elem()
}

func (i OrgIdpSamlArray) ToOrgIdpSamlArrayOutput() OrgIdpSamlArrayOutput {
	return i.ToOrgIdpSamlArrayOutputWithContext(context.Background())
}

func (i OrgIdpSamlArray) ToOrgIdpSamlArrayOutputWithContext(ctx context.Context) OrgIdpSamlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpSamlArrayOutput)
}

func (i OrgIdpSamlArray) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpSaml] {
	return pulumix.Output[[]*OrgIdpSaml]{
		OutputState: i.ToOrgIdpSamlArrayOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpSamlMapInput is an input type that accepts OrgIdpSamlMap and OrgIdpSamlMapOutput values.
// You can construct a concrete instance of `OrgIdpSamlMapInput` via:
//
//	OrgIdpSamlMap{ "key": OrgIdpSamlArgs{...} }
type OrgIdpSamlMapInput interface {
	pulumi.Input

	ToOrgIdpSamlMapOutput() OrgIdpSamlMapOutput
	ToOrgIdpSamlMapOutputWithContext(context.Context) OrgIdpSamlMapOutput
}

type OrgIdpSamlMap map[string]OrgIdpSamlInput

func (OrgIdpSamlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpSaml)(nil)).Elem()
}

func (i OrgIdpSamlMap) ToOrgIdpSamlMapOutput() OrgIdpSamlMapOutput {
	return i.ToOrgIdpSamlMapOutputWithContext(context.Background())
}

func (i OrgIdpSamlMap) ToOrgIdpSamlMapOutputWithContext(ctx context.Context) OrgIdpSamlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpSamlMapOutput)
}

func (i OrgIdpSamlMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpSaml] {
	return pulumix.Output[map[string]*OrgIdpSaml]{
		OutputState: i.ToOrgIdpSamlMapOutputWithContext(ctx).OutputState,
	}
}

type OrgIdpSamlOutput struct{ *pulumi.OutputState }

func (OrgIdpSamlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpSaml)(nil)).Elem()
}

func (o OrgIdpSamlOutput) ToOrgIdpSamlOutput() OrgIdpSamlOutput {
	return o
}

func (o OrgIdpSamlOutput) ToOrgIdpSamlOutputWithContext(ctx context.Context) OrgIdpSamlOutput {
	return o
}

func (o OrgIdpSamlOutput) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpSaml] {
	return pulumix.Output[*OrgIdpSaml]{
		OutputState: o.OutputState,
	}
}

// The binding, supported values: SAML*BINDING*UNSPECIFIED, SAML*BINDING*POST, SAML*BINDING*REDIRECT, SAML*BINDING*ARTIFACT
func (o OrgIdpSamlOutput) Binding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.StringPtrOutput { return v.Binding }).(pulumi.StringPtrOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o OrgIdpSamlOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o OrgIdpSamlOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o OrgIdpSamlOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o OrgIdpSamlOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// The metadata XML as plain string
func (o OrgIdpSamlOutput) MetadataXml() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.StringOutput { return v.MetadataXml }).(pulumi.StringOutput)
}

// Name of the IDP
func (o OrgIdpSamlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o OrgIdpSamlOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Whether the SAML IDP requires signed requests
func (o OrgIdpSamlOutput) WithSignedRequest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrgIdpSaml) pulumi.BoolPtrOutput { return v.WithSignedRequest }).(pulumi.BoolPtrOutput)
}

type OrgIdpSamlArrayOutput struct{ *pulumi.OutputState }

func (OrgIdpSamlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpSaml)(nil)).Elem()
}

func (o OrgIdpSamlArrayOutput) ToOrgIdpSamlArrayOutput() OrgIdpSamlArrayOutput {
	return o
}

func (o OrgIdpSamlArrayOutput) ToOrgIdpSamlArrayOutputWithContext(ctx context.Context) OrgIdpSamlArrayOutput {
	return o
}

func (o OrgIdpSamlArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpSaml] {
	return pulumix.Output[[]*OrgIdpSaml]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpSamlArrayOutput) Index(i pulumi.IntInput) OrgIdpSamlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgIdpSaml {
		return vs[0].([]*OrgIdpSaml)[vs[1].(int)]
	}).(OrgIdpSamlOutput)
}

type OrgIdpSamlMapOutput struct{ *pulumi.OutputState }

func (OrgIdpSamlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpSaml)(nil)).Elem()
}

func (o OrgIdpSamlMapOutput) ToOrgIdpSamlMapOutput() OrgIdpSamlMapOutput {
	return o
}

func (o OrgIdpSamlMapOutput) ToOrgIdpSamlMapOutputWithContext(ctx context.Context) OrgIdpSamlMapOutput {
	return o
}

func (o OrgIdpSamlMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpSaml] {
	return pulumix.Output[map[string]*OrgIdpSaml]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpSamlMapOutput) MapIndex(k pulumi.StringInput) OrgIdpSamlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgIdpSaml {
		return vs[0].(map[string]*OrgIdpSaml)[vs[1].(string)]
	}).(OrgIdpSamlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpSamlInput)(nil)).Elem(), &OrgIdpSaml{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpSamlArrayInput)(nil)).Elem(), OrgIdpSamlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpSamlMapInput)(nil)).Elem(), OrgIdpSamlMap{})
	pulumi.RegisterOutputType(OrgIdpSamlOutput{})
	pulumi.RegisterOutputType(OrgIdpSamlArrayOutput{})
	pulumi.RegisterOutputType(OrgIdpSamlMapOutput{})
}
