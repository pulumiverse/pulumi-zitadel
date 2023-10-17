// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing a Google IDP on the instance.
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
//			_, err := zitadel.NewIdpGoogle(ctx, "google", &zitadel.IdpGoogleArgs{
//				ClientId:          pulumi.String("182902..."),
//				ClientSecret:      pulumi.String("GOCSPX-*****"),
//				IsAutoCreation:    pulumi.Bool(false),
//				IsAutoUpdate:      pulumi.Bool(true),
//				IsCreationAllowed: pulumi.Bool(true),
//				IsLinkingAllowed:  pulumi.Bool(false),
//				Scopes: pulumi.StringArray{
//					pulumi.String("openid"),
//					pulumi.String("profile"),
//					pulumi.String("email"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IdpGoogle struct {
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
}

// NewIdpGoogle registers a new resource with the given unique name, arguments, and options.
func NewIdpGoogle(ctx *pulumi.Context,
	name string, args *IdpGoogleArgs, opts ...pulumi.ResourceOption) (*IdpGoogle, error) {
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
	opts = pkgResourceDefaultOpts(opts)
	var resource IdpGoogle
	err := ctx.RegisterResource("zitadel:index/idpGoogle:IdpGoogle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdpGoogle gets an existing IdpGoogle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdpGoogle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdpGoogleState, opts ...pulumi.ResourceOption) (*IdpGoogle, error) {
	var resource IdpGoogle
	err := ctx.ReadResource("zitadel:index/idpGoogle:IdpGoogle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdpGoogle resources.
type idpGoogleState struct {
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

type IdpGoogleState struct {
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (IdpGoogleState) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGoogleState)(nil)).Elem()
}

type idpGoogleArgs struct {
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a IdpGoogle resource.
type IdpGoogleArgs struct {
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (IdpGoogleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGoogleArgs)(nil)).Elem()
}

type IdpGoogleInput interface {
	pulumi.Input

	ToIdpGoogleOutput() IdpGoogleOutput
	ToIdpGoogleOutputWithContext(ctx context.Context) IdpGoogleOutput
}

func (*IdpGoogle) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGoogle)(nil)).Elem()
}

func (i *IdpGoogle) ToIdpGoogleOutput() IdpGoogleOutput {
	return i.ToIdpGoogleOutputWithContext(context.Background())
}

func (i *IdpGoogle) ToIdpGoogleOutputWithContext(ctx context.Context) IdpGoogleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGoogleOutput)
}

// IdpGoogleArrayInput is an input type that accepts IdpGoogleArray and IdpGoogleArrayOutput values.
// You can construct a concrete instance of `IdpGoogleArrayInput` via:
//
//	IdpGoogleArray{ IdpGoogleArgs{...} }
type IdpGoogleArrayInput interface {
	pulumi.Input

	ToIdpGoogleArrayOutput() IdpGoogleArrayOutput
	ToIdpGoogleArrayOutputWithContext(context.Context) IdpGoogleArrayOutput
}

type IdpGoogleArray []IdpGoogleInput

func (IdpGoogleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGoogle)(nil)).Elem()
}

func (i IdpGoogleArray) ToIdpGoogleArrayOutput() IdpGoogleArrayOutput {
	return i.ToIdpGoogleArrayOutputWithContext(context.Background())
}

func (i IdpGoogleArray) ToIdpGoogleArrayOutputWithContext(ctx context.Context) IdpGoogleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGoogleArrayOutput)
}

// IdpGoogleMapInput is an input type that accepts IdpGoogleMap and IdpGoogleMapOutput values.
// You can construct a concrete instance of `IdpGoogleMapInput` via:
//
//	IdpGoogleMap{ "key": IdpGoogleArgs{...} }
type IdpGoogleMapInput interface {
	pulumi.Input

	ToIdpGoogleMapOutput() IdpGoogleMapOutput
	ToIdpGoogleMapOutputWithContext(context.Context) IdpGoogleMapOutput
}

type IdpGoogleMap map[string]IdpGoogleInput

func (IdpGoogleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGoogle)(nil)).Elem()
}

func (i IdpGoogleMap) ToIdpGoogleMapOutput() IdpGoogleMapOutput {
	return i.ToIdpGoogleMapOutputWithContext(context.Background())
}

func (i IdpGoogleMap) ToIdpGoogleMapOutputWithContext(ctx context.Context) IdpGoogleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGoogleMapOutput)
}

type IdpGoogleOutput struct{ *pulumi.OutputState }

func (IdpGoogleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGoogle)(nil)).Elem()
}

func (o IdpGoogleOutput) ToIdpGoogleOutput() IdpGoogleOutput {
	return o
}

func (o IdpGoogleOutput) ToIdpGoogleOutputWithContext(ctx context.Context) IdpGoogleOutput {
	return o
}

// client id generated by the identity provider
func (o IdpGoogleOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGoogle) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o IdpGoogleOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGoogle) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o IdpGoogleOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGoogle) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o IdpGoogleOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGoogle) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o IdpGoogleOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGoogle) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o IdpGoogleOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGoogle) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o IdpGoogleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGoogle) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o IdpGoogleOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpGoogle) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

type IdpGoogleArrayOutput struct{ *pulumi.OutputState }

func (IdpGoogleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGoogle)(nil)).Elem()
}

func (o IdpGoogleArrayOutput) ToIdpGoogleArrayOutput() IdpGoogleArrayOutput {
	return o
}

func (o IdpGoogleArrayOutput) ToIdpGoogleArrayOutputWithContext(ctx context.Context) IdpGoogleArrayOutput {
	return o
}

func (o IdpGoogleArrayOutput) Index(i pulumi.IntInput) IdpGoogleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdpGoogle {
		return vs[0].([]*IdpGoogle)[vs[1].(int)]
	}).(IdpGoogleOutput)
}

type IdpGoogleMapOutput struct{ *pulumi.OutputState }

func (IdpGoogleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGoogle)(nil)).Elem()
}

func (o IdpGoogleMapOutput) ToIdpGoogleMapOutput() IdpGoogleMapOutput {
	return o
}

func (o IdpGoogleMapOutput) ToIdpGoogleMapOutputWithContext(ctx context.Context) IdpGoogleMapOutput {
	return o
}

func (o IdpGoogleMapOutput) MapIndex(k pulumi.StringInput) IdpGoogleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdpGoogle {
		return vs[0].(map[string]*IdpGoogle)[vs[1].(string)]
	}).(IdpGoogleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGoogleInput)(nil)).Elem(), &IdpGoogle{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGoogleArrayInput)(nil)).Elem(), IdpGoogleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGoogleMapInput)(nil)).Elem(), IdpGoogleMap{})
	pulumi.RegisterOutputType(IdpGoogleOutput{})
	pulumi.RegisterOutputType(IdpGoogleArrayOutput{})
	pulumi.RegisterOutputType(IdpGoogleMapOutput{})
}
