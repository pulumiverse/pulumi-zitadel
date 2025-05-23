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

// Resource representing an Azure AD IDP on the instance.
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
//			_, err := zitadel.NewIdpAzureAd(ctx, "default", &zitadel.IdpAzureAdArgs{
//				Name:         pulumi.String("Azure AD"),
//				ClientId:     pulumi.String("9065bfc8-a08a..."),
//				ClientSecret: pulumi.String("H2n***"),
//				Scopes: pulumi.StringArray{
//					pulumi.String("openid"),
//					pulumi.String("profile"),
//					pulumi.String("email"),
//					pulumi.String("User.Read"),
//				},
//				TenantType:        pulumi.String("AZURE_AD_TENANT_TYPE_ORGANISATIONS"),
//				EmailVerified:     pulumi.Bool(true),
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
//	$ pulumi import zitadel:index/idpAzureAd:IdpAzureAd imported '123456789012345678:12345678-1234-1234-1234-123456789012'
//
// ```
type IdpAzureAd struct {
	pulumi.CustomResourceState

	// client id generated by the identity provider
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// automatically mark emails as verified
	EmailVerified pulumi.BoolOutput `pulumi:"emailVerified"`
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
	// if tenant*id is not set, the tenant*type is used
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// the azure ad tenant type
	TenantType pulumi.StringPtrOutput `pulumi:"tenantType"`
}

// NewIdpAzureAd registers a new resource with the given unique name, arguments, and options.
func NewIdpAzureAd(ctx *pulumi.Context,
	name string, args *IdpAzureAdArgs, opts ...pulumi.ResourceOption) (*IdpAzureAd, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.EmailVerified == nil {
		return nil, errors.New("invalid value for required argument 'EmailVerified'")
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
	var resource IdpAzureAd
	err := ctx.RegisterResource("zitadel:index/idpAzureAd:IdpAzureAd", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdpAzureAd gets an existing IdpAzureAd resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdpAzureAd(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdpAzureAdState, opts ...pulumi.ResourceOption) (*IdpAzureAd, error) {
	var resource IdpAzureAd
	err := ctx.ReadResource("zitadel:index/idpAzureAd:IdpAzureAd", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdpAzureAd resources.
type idpAzureAdState struct {
	// client id generated by the identity provider
	ClientId *string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret *string `pulumi:"clientSecret"`
	// automatically mark emails as verified
	EmailVerified *bool `pulumi:"emailVerified"`
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
	// if tenant*id is not set, the tenant*type is used
	TenantId *string `pulumi:"tenantId"`
	// the azure ad tenant type
	TenantType *string `pulumi:"tenantType"`
}

type IdpAzureAdState struct {
	// client id generated by the identity provider
	ClientId pulumi.StringPtrInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringPtrInput
	// automatically mark emails as verified
	EmailVerified pulumi.BoolPtrInput
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
	// if tenant*id is not set, the tenant*type is used
	TenantId pulumi.StringPtrInput
	// the azure ad tenant type
	TenantType pulumi.StringPtrInput
}

func (IdpAzureAdState) ElementType() reflect.Type {
	return reflect.TypeOf((*idpAzureAdState)(nil)).Elem()
}

type idpAzureAdArgs struct {
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// automatically mark emails as verified
	EmailVerified bool `pulumi:"emailVerified"`
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
	// if tenant*id is not set, the tenant*type is used
	TenantId *string `pulumi:"tenantId"`
	// the azure ad tenant type
	TenantType *string `pulumi:"tenantType"`
}

// The set of arguments for constructing a IdpAzureAd resource.
type IdpAzureAdArgs struct {
	// client id generated by the identity provider
	ClientId pulumi.StringInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringInput
	// automatically mark emails as verified
	EmailVerified pulumi.BoolInput
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
	// if tenant*id is not set, the tenant*type is used
	TenantId pulumi.StringPtrInput
	// the azure ad tenant type
	TenantType pulumi.StringPtrInput
}

func (IdpAzureAdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*idpAzureAdArgs)(nil)).Elem()
}

type IdpAzureAdInput interface {
	pulumi.Input

	ToIdpAzureAdOutput() IdpAzureAdOutput
	ToIdpAzureAdOutputWithContext(ctx context.Context) IdpAzureAdOutput
}

func (*IdpAzureAd) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpAzureAd)(nil)).Elem()
}

func (i *IdpAzureAd) ToIdpAzureAdOutput() IdpAzureAdOutput {
	return i.ToIdpAzureAdOutputWithContext(context.Background())
}

func (i *IdpAzureAd) ToIdpAzureAdOutputWithContext(ctx context.Context) IdpAzureAdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpAzureAdOutput)
}

func (i *IdpAzureAd) ToOutput(ctx context.Context) pulumix.Output[*IdpAzureAd] {
	return pulumix.Output[*IdpAzureAd]{
		OutputState: i.ToIdpAzureAdOutputWithContext(ctx).OutputState,
	}
}

// IdpAzureAdArrayInput is an input type that accepts IdpAzureAdArray and IdpAzureAdArrayOutput values.
// You can construct a concrete instance of `IdpAzureAdArrayInput` via:
//
//	IdpAzureAdArray{ IdpAzureAdArgs{...} }
type IdpAzureAdArrayInput interface {
	pulumi.Input

	ToIdpAzureAdArrayOutput() IdpAzureAdArrayOutput
	ToIdpAzureAdArrayOutputWithContext(context.Context) IdpAzureAdArrayOutput
}

type IdpAzureAdArray []IdpAzureAdInput

func (IdpAzureAdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpAzureAd)(nil)).Elem()
}

func (i IdpAzureAdArray) ToIdpAzureAdArrayOutput() IdpAzureAdArrayOutput {
	return i.ToIdpAzureAdArrayOutputWithContext(context.Background())
}

func (i IdpAzureAdArray) ToIdpAzureAdArrayOutputWithContext(ctx context.Context) IdpAzureAdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpAzureAdArrayOutput)
}

func (i IdpAzureAdArray) ToOutput(ctx context.Context) pulumix.Output[[]*IdpAzureAd] {
	return pulumix.Output[[]*IdpAzureAd]{
		OutputState: i.ToIdpAzureAdArrayOutputWithContext(ctx).OutputState,
	}
}

// IdpAzureAdMapInput is an input type that accepts IdpAzureAdMap and IdpAzureAdMapOutput values.
// You can construct a concrete instance of `IdpAzureAdMapInput` via:
//
//	IdpAzureAdMap{ "key": IdpAzureAdArgs{...} }
type IdpAzureAdMapInput interface {
	pulumi.Input

	ToIdpAzureAdMapOutput() IdpAzureAdMapOutput
	ToIdpAzureAdMapOutputWithContext(context.Context) IdpAzureAdMapOutput
}

type IdpAzureAdMap map[string]IdpAzureAdInput

func (IdpAzureAdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpAzureAd)(nil)).Elem()
}

func (i IdpAzureAdMap) ToIdpAzureAdMapOutput() IdpAzureAdMapOutput {
	return i.ToIdpAzureAdMapOutputWithContext(context.Background())
}

func (i IdpAzureAdMap) ToIdpAzureAdMapOutputWithContext(ctx context.Context) IdpAzureAdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpAzureAdMapOutput)
}

func (i IdpAzureAdMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IdpAzureAd] {
	return pulumix.Output[map[string]*IdpAzureAd]{
		OutputState: i.ToIdpAzureAdMapOutputWithContext(ctx).OutputState,
	}
}

type IdpAzureAdOutput struct{ *pulumi.OutputState }

func (IdpAzureAdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpAzureAd)(nil)).Elem()
}

func (o IdpAzureAdOutput) ToIdpAzureAdOutput() IdpAzureAdOutput {
	return o
}

func (o IdpAzureAdOutput) ToIdpAzureAdOutputWithContext(ctx context.Context) IdpAzureAdOutput {
	return o
}

func (o IdpAzureAdOutput) ToOutput(ctx context.Context) pulumix.Output[*IdpAzureAd] {
	return pulumix.Output[*IdpAzureAd]{
		OutputState: o.OutputState,
	}
}

// client id generated by the identity provider
func (o IdpAzureAdOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o IdpAzureAdOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// automatically mark emails as verified
func (o IdpAzureAdOutput) EmailVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.BoolOutput { return v.EmailVerified }).(pulumi.BoolOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o IdpAzureAdOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o IdpAzureAdOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o IdpAzureAdOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o IdpAzureAdOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o IdpAzureAdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o IdpAzureAdOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// if tenant*id is not set, the tenant*type is used
func (o IdpAzureAdOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// the azure ad tenant type
func (o IdpAzureAdOutput) TenantType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpAzureAd) pulumi.StringPtrOutput { return v.TenantType }).(pulumi.StringPtrOutput)
}

type IdpAzureAdArrayOutput struct{ *pulumi.OutputState }

func (IdpAzureAdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpAzureAd)(nil)).Elem()
}

func (o IdpAzureAdArrayOutput) ToIdpAzureAdArrayOutput() IdpAzureAdArrayOutput {
	return o
}

func (o IdpAzureAdArrayOutput) ToIdpAzureAdArrayOutputWithContext(ctx context.Context) IdpAzureAdArrayOutput {
	return o
}

func (o IdpAzureAdArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IdpAzureAd] {
	return pulumix.Output[[]*IdpAzureAd]{
		OutputState: o.OutputState,
	}
}

func (o IdpAzureAdArrayOutput) Index(i pulumi.IntInput) IdpAzureAdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdpAzureAd {
		return vs[0].([]*IdpAzureAd)[vs[1].(int)]
	}).(IdpAzureAdOutput)
}

type IdpAzureAdMapOutput struct{ *pulumi.OutputState }

func (IdpAzureAdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpAzureAd)(nil)).Elem()
}

func (o IdpAzureAdMapOutput) ToIdpAzureAdMapOutput() IdpAzureAdMapOutput {
	return o
}

func (o IdpAzureAdMapOutput) ToIdpAzureAdMapOutputWithContext(ctx context.Context) IdpAzureAdMapOutput {
	return o
}

func (o IdpAzureAdMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IdpAzureAd] {
	return pulumix.Output[map[string]*IdpAzureAd]{
		OutputState: o.OutputState,
	}
}

func (o IdpAzureAdMapOutput) MapIndex(k pulumi.StringInput) IdpAzureAdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdpAzureAd {
		return vs[0].(map[string]*IdpAzureAd)[vs[1].(string)]
	}).(IdpAzureAdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdpAzureAdInput)(nil)).Elem(), &IdpAzureAd{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpAzureAdArrayInput)(nil)).Elem(), IdpAzureAdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpAzureAdMapInput)(nil)).Elem(), IdpAzureAdMap{})
	pulumi.RegisterOutputType(IdpAzureAdOutput{})
	pulumi.RegisterOutputType(IdpAzureAdArrayOutput{})
	pulumi.RegisterOutputType(IdpAzureAdMapOutput{})
}
