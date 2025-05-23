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

// Resource representing a GitHub IDP on the instance.
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
//			_, err := zitadel.NewIdpGithub(ctx, "default", &zitadel.IdpGithubArgs{
//				Name:         pulumi.String("GitHub"),
//				ClientId:     pulumi.String("86a165..."),
//				ClientSecret: pulumi.String("*****afdbac18"),
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
// bash The resource can be imported using the ID format `<id[:client_secret]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/idpGithub:IdpGithub imported '123456789012345678:1234567890123456781234567890123456787890'
//
// ```
type IdpGithub struct {
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

// NewIdpGithub registers a new resource with the given unique name, arguments, and options.
func NewIdpGithub(ctx *pulumi.Context,
	name string, args *IdpGithubArgs, opts ...pulumi.ResourceOption) (*IdpGithub, error) {
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
	var resource IdpGithub
	err := ctx.RegisterResource("zitadel:index/idpGithub:IdpGithub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdpGithub gets an existing IdpGithub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdpGithub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdpGithubState, opts ...pulumi.ResourceOption) (*IdpGithub, error) {
	var resource IdpGithub
	err := ctx.ReadResource("zitadel:index/idpGithub:IdpGithub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdpGithub resources.
type idpGithubState struct {
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

type IdpGithubState struct {
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

func (IdpGithubState) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGithubState)(nil)).Elem()
}

type idpGithubArgs struct {
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

// The set of arguments for constructing a IdpGithub resource.
type IdpGithubArgs struct {
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

func (IdpGithubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGithubArgs)(nil)).Elem()
}

type IdpGithubInput interface {
	pulumi.Input

	ToIdpGithubOutput() IdpGithubOutput
	ToIdpGithubOutputWithContext(ctx context.Context) IdpGithubOutput
}

func (*IdpGithub) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGithub)(nil)).Elem()
}

func (i *IdpGithub) ToIdpGithubOutput() IdpGithubOutput {
	return i.ToIdpGithubOutputWithContext(context.Background())
}

func (i *IdpGithub) ToIdpGithubOutputWithContext(ctx context.Context) IdpGithubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGithubOutput)
}

func (i *IdpGithub) ToOutput(ctx context.Context) pulumix.Output[*IdpGithub] {
	return pulumix.Output[*IdpGithub]{
		OutputState: i.ToIdpGithubOutputWithContext(ctx).OutputState,
	}
}

// IdpGithubArrayInput is an input type that accepts IdpGithubArray and IdpGithubArrayOutput values.
// You can construct a concrete instance of `IdpGithubArrayInput` via:
//
//	IdpGithubArray{ IdpGithubArgs{...} }
type IdpGithubArrayInput interface {
	pulumi.Input

	ToIdpGithubArrayOutput() IdpGithubArrayOutput
	ToIdpGithubArrayOutputWithContext(context.Context) IdpGithubArrayOutput
}

type IdpGithubArray []IdpGithubInput

func (IdpGithubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGithub)(nil)).Elem()
}

func (i IdpGithubArray) ToIdpGithubArrayOutput() IdpGithubArrayOutput {
	return i.ToIdpGithubArrayOutputWithContext(context.Background())
}

func (i IdpGithubArray) ToIdpGithubArrayOutputWithContext(ctx context.Context) IdpGithubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGithubArrayOutput)
}

func (i IdpGithubArray) ToOutput(ctx context.Context) pulumix.Output[[]*IdpGithub] {
	return pulumix.Output[[]*IdpGithub]{
		OutputState: i.ToIdpGithubArrayOutputWithContext(ctx).OutputState,
	}
}

// IdpGithubMapInput is an input type that accepts IdpGithubMap and IdpGithubMapOutput values.
// You can construct a concrete instance of `IdpGithubMapInput` via:
//
//	IdpGithubMap{ "key": IdpGithubArgs{...} }
type IdpGithubMapInput interface {
	pulumi.Input

	ToIdpGithubMapOutput() IdpGithubMapOutput
	ToIdpGithubMapOutputWithContext(context.Context) IdpGithubMapOutput
}

type IdpGithubMap map[string]IdpGithubInput

func (IdpGithubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGithub)(nil)).Elem()
}

func (i IdpGithubMap) ToIdpGithubMapOutput() IdpGithubMapOutput {
	return i.ToIdpGithubMapOutputWithContext(context.Background())
}

func (i IdpGithubMap) ToIdpGithubMapOutputWithContext(ctx context.Context) IdpGithubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGithubMapOutput)
}

func (i IdpGithubMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IdpGithub] {
	return pulumix.Output[map[string]*IdpGithub]{
		OutputState: i.ToIdpGithubMapOutputWithContext(ctx).OutputState,
	}
}

type IdpGithubOutput struct{ *pulumi.OutputState }

func (IdpGithubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGithub)(nil)).Elem()
}

func (o IdpGithubOutput) ToIdpGithubOutput() IdpGithubOutput {
	return o
}

func (o IdpGithubOutput) ToIdpGithubOutputWithContext(ctx context.Context) IdpGithubOutput {
	return o
}

func (o IdpGithubOutput) ToOutput(ctx context.Context) pulumix.Output[*IdpGithub] {
	return pulumix.Output[*IdpGithub]{
		OutputState: o.OutputState,
	}
}

// client id generated by the identity provider
func (o IdpGithubOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithub) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o IdpGithubOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithub) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o IdpGithubOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGithub) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o IdpGithubOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGithub) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o IdpGithubOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGithub) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o IdpGithubOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGithub) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o IdpGithubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o IdpGithubOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpGithub) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

type IdpGithubArrayOutput struct{ *pulumi.OutputState }

func (IdpGithubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGithub)(nil)).Elem()
}

func (o IdpGithubArrayOutput) ToIdpGithubArrayOutput() IdpGithubArrayOutput {
	return o
}

func (o IdpGithubArrayOutput) ToIdpGithubArrayOutputWithContext(ctx context.Context) IdpGithubArrayOutput {
	return o
}

func (o IdpGithubArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IdpGithub] {
	return pulumix.Output[[]*IdpGithub]{
		OutputState: o.OutputState,
	}
}

func (o IdpGithubArrayOutput) Index(i pulumi.IntInput) IdpGithubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdpGithub {
		return vs[0].([]*IdpGithub)[vs[1].(int)]
	}).(IdpGithubOutput)
}

type IdpGithubMapOutput struct{ *pulumi.OutputState }

func (IdpGithubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGithub)(nil)).Elem()
}

func (o IdpGithubMapOutput) ToIdpGithubMapOutput() IdpGithubMapOutput {
	return o
}

func (o IdpGithubMapOutput) ToIdpGithubMapOutputWithContext(ctx context.Context) IdpGithubMapOutput {
	return o
}

func (o IdpGithubMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IdpGithub] {
	return pulumix.Output[map[string]*IdpGithub]{
		OutputState: o.OutputState,
	}
}

func (o IdpGithubMapOutput) MapIndex(k pulumi.StringInput) IdpGithubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdpGithub {
		return vs[0].(map[string]*IdpGithub)[vs[1].(string)]
	}).(IdpGithubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGithubInput)(nil)).Elem(), &IdpGithub{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGithubArrayInput)(nil)).Elem(), IdpGithubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGithubMapInput)(nil)).Elem(), IdpGithubMap{})
	pulumi.RegisterOutputType(IdpGithubOutput{})
	pulumi.RegisterOutputType(IdpGithubArrayOutput{})
	pulumi.RegisterOutputType(IdpGithubMapOutput{})
}
