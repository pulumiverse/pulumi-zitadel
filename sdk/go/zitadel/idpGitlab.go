// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing a GitLab IDP on the instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.NewIdpGitlab(ctx, "default", &zitadel.IdpGitlabArgs{
// 			ClientId:          pulumi.String("15765e..."),
// 			ClientSecret:      pulumi.String("*****abcxyz"),
// 			IsAutoCreation:    pulumi.Bool(false),
// 			IsAutoUpdate:      pulumi.Bool(true),
// 			IsCreationAllowed: pulumi.Bool(true),
// 			IsLinkingAllowed:  pulumi.Bool(false),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("openid"),
// 				pulumi.String("profile"),
// 				pulumi.String("email"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// terraform # The resource can be imported using the ID format `<id[:client_secret]>`, e.g.
//
// ```sh
//  $ pulumi import zitadel:index/idpGitlab:IdpGitlab imported '123456789012345678:1234567890abcdef'
// ```
type IdpGitlab struct {
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

// NewIdpGitlab registers a new resource with the given unique name, arguments, and options.
func NewIdpGitlab(ctx *pulumi.Context,
	name string, args *IdpGitlabArgs, opts ...pulumi.ResourceOption) (*IdpGitlab, error) {
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
	var resource IdpGitlab
	err := ctx.RegisterResource("zitadel:index/idpGitlab:IdpGitlab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdpGitlab gets an existing IdpGitlab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdpGitlab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdpGitlabState, opts ...pulumi.ResourceOption) (*IdpGitlab, error) {
	var resource IdpGitlab
	err := ctx.ReadResource("zitadel:index/idpGitlab:IdpGitlab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdpGitlab resources.
type idpGitlabState struct {
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

type IdpGitlabState struct {
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

func (IdpGitlabState) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGitlabState)(nil)).Elem()
}

type idpGitlabArgs struct {
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

// The set of arguments for constructing a IdpGitlab resource.
type IdpGitlabArgs struct {
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

func (IdpGitlabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGitlabArgs)(nil)).Elem()
}

type IdpGitlabInput interface {
	pulumi.Input

	ToIdpGitlabOutput() IdpGitlabOutput
	ToIdpGitlabOutputWithContext(ctx context.Context) IdpGitlabOutput
}

func (*IdpGitlab) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGitlab)(nil)).Elem()
}

func (i *IdpGitlab) ToIdpGitlabOutput() IdpGitlabOutput {
	return i.ToIdpGitlabOutputWithContext(context.Background())
}

func (i *IdpGitlab) ToIdpGitlabOutputWithContext(ctx context.Context) IdpGitlabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGitlabOutput)
}

// IdpGitlabArrayInput is an input type that accepts IdpGitlabArray and IdpGitlabArrayOutput values.
// You can construct a concrete instance of `IdpGitlabArrayInput` via:
//
//          IdpGitlabArray{ IdpGitlabArgs{...} }
type IdpGitlabArrayInput interface {
	pulumi.Input

	ToIdpGitlabArrayOutput() IdpGitlabArrayOutput
	ToIdpGitlabArrayOutputWithContext(context.Context) IdpGitlabArrayOutput
}

type IdpGitlabArray []IdpGitlabInput

func (IdpGitlabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGitlab)(nil)).Elem()
}

func (i IdpGitlabArray) ToIdpGitlabArrayOutput() IdpGitlabArrayOutput {
	return i.ToIdpGitlabArrayOutputWithContext(context.Background())
}

func (i IdpGitlabArray) ToIdpGitlabArrayOutputWithContext(ctx context.Context) IdpGitlabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGitlabArrayOutput)
}

// IdpGitlabMapInput is an input type that accepts IdpGitlabMap and IdpGitlabMapOutput values.
// You can construct a concrete instance of `IdpGitlabMapInput` via:
//
//          IdpGitlabMap{ "key": IdpGitlabArgs{...} }
type IdpGitlabMapInput interface {
	pulumi.Input

	ToIdpGitlabMapOutput() IdpGitlabMapOutput
	ToIdpGitlabMapOutputWithContext(context.Context) IdpGitlabMapOutput
}

type IdpGitlabMap map[string]IdpGitlabInput

func (IdpGitlabMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGitlab)(nil)).Elem()
}

func (i IdpGitlabMap) ToIdpGitlabMapOutput() IdpGitlabMapOutput {
	return i.ToIdpGitlabMapOutputWithContext(context.Background())
}

func (i IdpGitlabMap) ToIdpGitlabMapOutputWithContext(ctx context.Context) IdpGitlabMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGitlabMapOutput)
}

type IdpGitlabOutput struct{ *pulumi.OutputState }

func (IdpGitlabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGitlab)(nil)).Elem()
}

func (o IdpGitlabOutput) ToIdpGitlabOutput() IdpGitlabOutput {
	return o
}

func (o IdpGitlabOutput) ToIdpGitlabOutputWithContext(ctx context.Context) IdpGitlabOutput {
	return o
}

// client id generated by the identity provider
func (o IdpGitlabOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGitlab) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o IdpGitlabOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGitlab) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o IdpGitlabOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGitlab) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o IdpGitlabOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGitlab) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o IdpGitlabOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGitlab) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o IdpGitlabOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGitlab) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o IdpGitlabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGitlab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o IdpGitlabOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpGitlab) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

type IdpGitlabArrayOutput struct{ *pulumi.OutputState }

func (IdpGitlabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGitlab)(nil)).Elem()
}

func (o IdpGitlabArrayOutput) ToIdpGitlabArrayOutput() IdpGitlabArrayOutput {
	return o
}

func (o IdpGitlabArrayOutput) ToIdpGitlabArrayOutputWithContext(ctx context.Context) IdpGitlabArrayOutput {
	return o
}

func (o IdpGitlabArrayOutput) Index(i pulumi.IntInput) IdpGitlabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdpGitlab {
		return vs[0].([]*IdpGitlab)[vs[1].(int)]
	}).(IdpGitlabOutput)
}

type IdpGitlabMapOutput struct{ *pulumi.OutputState }

func (IdpGitlabMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGitlab)(nil)).Elem()
}

func (o IdpGitlabMapOutput) ToIdpGitlabMapOutput() IdpGitlabMapOutput {
	return o
}

func (o IdpGitlabMapOutput) ToIdpGitlabMapOutputWithContext(ctx context.Context) IdpGitlabMapOutput {
	return o
}

func (o IdpGitlabMapOutput) MapIndex(k pulumi.StringInput) IdpGitlabOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdpGitlab {
		return vs[0].(map[string]*IdpGitlab)[vs[1].(string)]
	}).(IdpGitlabOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGitlabInput)(nil)).Elem(), &IdpGitlab{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGitlabArrayInput)(nil)).Elem(), IdpGitlabArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGitlabMapInput)(nil)).Elem(), IdpGitlabMap{})
	pulumi.RegisterOutputType(IdpGitlabOutput{})
	pulumi.RegisterOutputType(IdpGitlabArrayOutput{})
	pulumi.RegisterOutputType(IdpGitlabMapOutput{})
}
