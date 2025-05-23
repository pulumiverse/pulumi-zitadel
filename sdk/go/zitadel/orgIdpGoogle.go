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

// Resource representing a Google IdP on the organization.
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
//			_, err := zitadel.NewOrgIdpGoogle(ctx, "default", &zitadel.OrgIdpGoogleArgs{
//				OrgId:        pulumi.Any(defaultZitadelOrg.Id),
//				Name:         pulumi.String("Google"),
//				ClientId:     pulumi.String("182902..."),
//				ClientSecret: pulumi.String("GOCSPX-*****"),
//				Scopes: pulumi.StringArray{
//					pulumi.String("openid"),
//					pulumi.String("profile"),
//					pulumi.String("email"),
//				},
//				IsLinkingAllowed:  pulumi.Bool(false),
//				IsCreationAllowed: pulumi.Bool(true),
//				IsAutoCreation:    pulumi.Bool(false),
//				IsAutoUpdate:      pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// bash The resource can be imported using the ID format `<id[:org_id][:client_secret]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/orgIdpGoogle:OrgIdpGoogle imported '123456789012345678:123456789012345678:G1234567890123'
//
// ```
type OrgIdpGoogle struct {
	pulumi.CustomResourceState

	// client id generated by the identity provider
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolOutput `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolOutput `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolOutput `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolOutput `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
}

// NewOrgIdpGoogle registers a new resource with the given unique name, arguments, and options.
func NewOrgIdpGoogle(ctx *pulumi.Context,
	name string, args *OrgIdpGoogleArgs, opts ...pulumi.ResourceOption) (*OrgIdpGoogle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
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
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgIdpGoogle
	err := ctx.RegisterResource("zitadel:index/orgIdpGoogle:OrgIdpGoogle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgIdpGoogle gets an existing OrgIdpGoogle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgIdpGoogle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgIdpGoogleState, opts ...pulumi.ResourceOption) (*OrgIdpGoogle, error) {
	var resource OrgIdpGoogle
	err := ctx.ReadResource("zitadel:index/orgIdpGoogle:OrgIdpGoogle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgIdpGoogle resources.
type orgIdpGoogleState struct {
	// client id generated by the identity provider
	ClientId *string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret *string `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation *bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate *bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed *bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed *bool `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

type OrgIdpGoogleState struct {
	// client id generated by the identity provider
	ClientId pulumi.StringPtrInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringPtrInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolPtrInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolPtrInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolPtrInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolPtrInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (OrgIdpGoogleState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpGoogleState)(nil)).Elem()
}

type orgIdpGoogleArgs struct {
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a OrgIdpGoogle resource.
type OrgIdpGoogleArgs struct {
	// client id generated by the identity provider
	ClientId pulumi.StringInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (OrgIdpGoogleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpGoogleArgs)(nil)).Elem()
}

type OrgIdpGoogleInput interface {
	pulumi.Input

	ToOrgIdpGoogleOutput() OrgIdpGoogleOutput
	ToOrgIdpGoogleOutputWithContext(ctx context.Context) OrgIdpGoogleOutput
}

func (*OrgIdpGoogle) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpGoogle)(nil)).Elem()
}

func (i *OrgIdpGoogle) ToOrgIdpGoogleOutput() OrgIdpGoogleOutput {
	return i.ToOrgIdpGoogleOutputWithContext(context.Background())
}

func (i *OrgIdpGoogle) ToOrgIdpGoogleOutputWithContext(ctx context.Context) OrgIdpGoogleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGoogleOutput)
}

func (i *OrgIdpGoogle) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpGoogle] {
	return pulumix.Output[*OrgIdpGoogle]{
		OutputState: i.ToOrgIdpGoogleOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpGoogleArrayInput is an input type that accepts OrgIdpGoogleArray and OrgIdpGoogleArrayOutput values.
// You can construct a concrete instance of `OrgIdpGoogleArrayInput` via:
//
//	OrgIdpGoogleArray{ OrgIdpGoogleArgs{...} }
type OrgIdpGoogleArrayInput interface {
	pulumi.Input

	ToOrgIdpGoogleArrayOutput() OrgIdpGoogleArrayOutput
	ToOrgIdpGoogleArrayOutputWithContext(context.Context) OrgIdpGoogleArrayOutput
}

type OrgIdpGoogleArray []OrgIdpGoogleInput

func (OrgIdpGoogleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpGoogle)(nil)).Elem()
}

func (i OrgIdpGoogleArray) ToOrgIdpGoogleArrayOutput() OrgIdpGoogleArrayOutput {
	return i.ToOrgIdpGoogleArrayOutputWithContext(context.Background())
}

func (i OrgIdpGoogleArray) ToOrgIdpGoogleArrayOutputWithContext(ctx context.Context) OrgIdpGoogleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGoogleArrayOutput)
}

func (i OrgIdpGoogleArray) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpGoogle] {
	return pulumix.Output[[]*OrgIdpGoogle]{
		OutputState: i.ToOrgIdpGoogleArrayOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpGoogleMapInput is an input type that accepts OrgIdpGoogleMap and OrgIdpGoogleMapOutput values.
// You can construct a concrete instance of `OrgIdpGoogleMapInput` via:
//
//	OrgIdpGoogleMap{ "key": OrgIdpGoogleArgs{...} }
type OrgIdpGoogleMapInput interface {
	pulumi.Input

	ToOrgIdpGoogleMapOutput() OrgIdpGoogleMapOutput
	ToOrgIdpGoogleMapOutputWithContext(context.Context) OrgIdpGoogleMapOutput
}

type OrgIdpGoogleMap map[string]OrgIdpGoogleInput

func (OrgIdpGoogleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpGoogle)(nil)).Elem()
}

func (i OrgIdpGoogleMap) ToOrgIdpGoogleMapOutput() OrgIdpGoogleMapOutput {
	return i.ToOrgIdpGoogleMapOutputWithContext(context.Background())
}

func (i OrgIdpGoogleMap) ToOrgIdpGoogleMapOutputWithContext(ctx context.Context) OrgIdpGoogleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGoogleMapOutput)
}

func (i OrgIdpGoogleMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpGoogle] {
	return pulumix.Output[map[string]*OrgIdpGoogle]{
		OutputState: i.ToOrgIdpGoogleMapOutputWithContext(ctx).OutputState,
	}
}

type OrgIdpGoogleOutput struct{ *pulumi.OutputState }

func (OrgIdpGoogleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpGoogle)(nil)).Elem()
}

func (o OrgIdpGoogleOutput) ToOrgIdpGoogleOutput() OrgIdpGoogleOutput {
	return o
}

func (o OrgIdpGoogleOutput) ToOrgIdpGoogleOutputWithContext(ctx context.Context) OrgIdpGoogleOutput {
	return o
}

func (o OrgIdpGoogleOutput) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpGoogle] {
	return pulumix.Output[*OrgIdpGoogle]{
		OutputState: o.OutputState,
	}
}

// client id generated by the identity provider
func (o OrgIdpGoogleOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o OrgIdpGoogleOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o OrgIdpGoogleOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o OrgIdpGoogleOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o OrgIdpGoogleOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o OrgIdpGoogleOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o OrgIdpGoogleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o OrgIdpGoogleOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o OrgIdpGoogleOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgIdpGoogle) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

type OrgIdpGoogleArrayOutput struct{ *pulumi.OutputState }

func (OrgIdpGoogleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpGoogle)(nil)).Elem()
}

func (o OrgIdpGoogleArrayOutput) ToOrgIdpGoogleArrayOutput() OrgIdpGoogleArrayOutput {
	return o
}

func (o OrgIdpGoogleArrayOutput) ToOrgIdpGoogleArrayOutputWithContext(ctx context.Context) OrgIdpGoogleArrayOutput {
	return o
}

func (o OrgIdpGoogleArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpGoogle] {
	return pulumix.Output[[]*OrgIdpGoogle]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpGoogleArrayOutput) Index(i pulumi.IntInput) OrgIdpGoogleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgIdpGoogle {
		return vs[0].([]*OrgIdpGoogle)[vs[1].(int)]
	}).(OrgIdpGoogleOutput)
}

type OrgIdpGoogleMapOutput struct{ *pulumi.OutputState }

func (OrgIdpGoogleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpGoogle)(nil)).Elem()
}

func (o OrgIdpGoogleMapOutput) ToOrgIdpGoogleMapOutput() OrgIdpGoogleMapOutput {
	return o
}

func (o OrgIdpGoogleMapOutput) ToOrgIdpGoogleMapOutputWithContext(ctx context.Context) OrgIdpGoogleMapOutput {
	return o
}

func (o OrgIdpGoogleMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpGoogle] {
	return pulumix.Output[map[string]*OrgIdpGoogle]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpGoogleMapOutput) MapIndex(k pulumi.StringInput) OrgIdpGoogleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgIdpGoogle {
		return vs[0].(map[string]*OrgIdpGoogle)[vs[1].(string)]
	}).(OrgIdpGoogleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGoogleInput)(nil)).Elem(), &OrgIdpGoogle{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGoogleArrayInput)(nil)).Elem(), OrgIdpGoogleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGoogleMapInput)(nil)).Elem(), OrgIdpGoogleMap{})
	pulumi.RegisterOutputType(OrgIdpGoogleOutput{})
	pulumi.RegisterOutputType(OrgIdpGoogleArrayOutput{})
	pulumi.RegisterOutputType(OrgIdpGoogleMapOutput{})
}
