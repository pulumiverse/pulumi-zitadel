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

// Resource representing an OIDC application belonging to a project, with all configuration possibilities.
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
//			_, err := zitadel.NewApplicationOidc(ctx, "default", &zitadel.ApplicationOidcArgs{
//				ProjectId: pulumi.Any(defaultZitadelProject.Id),
//				OrgId:     pulumi.Any(defaultZitadelOrg.Id),
//				Name:      pulumi.String("applicationoidc"),
//				RedirectUris: pulumi.StringArray{
//					pulumi.String("https://localhost.com"),
//				},
//				ResponseTypes: pulumi.StringArray{
//					pulumi.String("OIDC_RESPONSE_TYPE_CODE"),
//				},
//				GrantTypes: pulumi.StringArray{
//					pulumi.String("OIDC_GRANT_TYPE_AUTHORIZATION_CODE"),
//				},
//				PostLogoutRedirectUris: pulumi.StringArray{
//					pulumi.String("https://localhost.com"),
//				},
//				AppType:                  pulumi.String("OIDC_APP_TYPE_WEB"),
//				AuthMethodType:           pulumi.String("OIDC_AUTH_METHOD_TYPE_BASIC"),
//				Version:                  pulumi.String("OIDC_VERSION_1_0"),
//				ClockSkew:                pulumi.String("0s"),
//				DevMode:                  pulumi.Bool(true),
//				AccessTokenType:          pulumi.String("OIDC_TOKEN_TYPE_BEARER"),
//				AccessTokenRoleAssertion: pulumi.Bool(false),
//				IdTokenRoleAssertion:     pulumi.Bool(false),
//				IdTokenUserinfoAssertion: pulumi.Bool(false),
//				AdditionalOrigins:        pulumi.StringArray{},
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
// bash The resource can be imported using the ID format `<id:project_id[:org_id][:client_id][:client_secret]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/applicationOidc:ApplicationOidc imported '123456789012345678:123456789012345678:123456789012345678:123456789012345678@zitadel:JuaDFFeOak5DGE655KCYPSAclSkbMVEJXXuX1lEMBT14eLMSs0A0qhafKX5SA2Df'
//
// ```
type ApplicationOidc struct {
	pulumi.CustomResourceState

	// Access token role assertion
	AccessTokenRoleAssertion pulumi.BoolPtrOutput `pulumi:"accessTokenRoleAssertion"`
	// Access token type, supported values: OIDC*TOKEN*TYPE*BEARER, OIDC*TOKEN*TYPE*JWT
	AccessTokenType pulumi.StringPtrOutput `pulumi:"accessTokenType"`
	// Additional origins
	AdditionalOrigins pulumi.StringArrayOutput `pulumi:"additionalOrigins"`
	// App type, supported values: OIDC*APP*TYPE*WEB, OIDC*APP*TYPE*USER*AGENT, OIDC*APP*TYPE*NATIVE
	AppType pulumi.StringPtrOutput `pulumi:"appType"`
	// Auth method type, supported values: OIDC*AUTH*METHOD*TYPE*BASIC, OIDC*AUTH*METHOD*TYPE*POST, OIDC*AUTH*METHOD*TYPE*NONE, OIDC*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType pulumi.StringPtrOutput `pulumi:"authMethodType"`
	// generated ID for this config
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// generated secret for this config
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// Clockskew
	ClockSkew pulumi.StringPtrOutput `pulumi:"clockSkew"`
	// Dev mode
	DevMode pulumi.BoolPtrOutput `pulumi:"devMode"`
	// Grant types, supported values: OIDC*GRANT*TYPE*AUTHORIZATION*CODE, OIDC*GRANT*TYPE*IMPLICIT, OIDC*GRANT*TYPE*REFRESH*TOKEN, OIDC*GRANT*TYPE*DEVICE_CODE
	GrantTypes pulumi.StringArrayOutput `pulumi:"grantTypes"`
	// ID token role assertion
	IdTokenRoleAssertion pulumi.BoolPtrOutput `pulumi:"idTokenRoleAssertion"`
	// Token userinfo assertion
	IdTokenUserinfoAssertion pulumi.BoolPtrOutput `pulumi:"idTokenUserinfoAssertion"`
	// Name of the application
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Post logout redirect URIs
	PostLogoutRedirectUris pulumi.StringArrayOutput `pulumi:"postLogoutRedirectUris"`
	// ID of the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// RedirectURIs
	RedirectUris pulumi.StringArrayOutput `pulumi:"redirectUris"`
	// Response type, supported values: OIDC*RESPONSE*TYPE*CODE, OIDC*RESPONSE*TYPE*ID*TOKEN, OIDC*RESPONSE*TYPE*ID*TOKEN*TOKEN
	ResponseTypes pulumi.StringArrayOutput `pulumi:"responseTypes"`
	// Version, supported values: OIDC*VERSION*1_0
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewApplicationOidc registers a new resource with the given unique name, arguments, and options.
func NewApplicationOidc(ctx *pulumi.Context,
	name string, args *ApplicationOidcArgs, opts ...pulumi.ResourceOption) (*ApplicationOidc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GrantTypes == nil {
		return nil, errors.New("invalid value for required argument 'GrantTypes'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.RedirectUris == nil {
		return nil, errors.New("invalid value for required argument 'RedirectUris'")
	}
	if args.ResponseTypes == nil {
		return nil, errors.New("invalid value for required argument 'ResponseTypes'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientId",
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationOidc
	err := ctx.RegisterResource("zitadel:index/applicationOidc:ApplicationOidc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationOidc gets an existing ApplicationOidc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationOidc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationOidcState, opts ...pulumi.ResourceOption) (*ApplicationOidc, error) {
	var resource ApplicationOidc
	err := ctx.ReadResource("zitadel:index/applicationOidc:ApplicationOidc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationOidc resources.
type applicationOidcState struct {
	// Access token role assertion
	AccessTokenRoleAssertion *bool `pulumi:"accessTokenRoleAssertion"`
	// Access token type, supported values: OIDC*TOKEN*TYPE*BEARER, OIDC*TOKEN*TYPE*JWT
	AccessTokenType *string `pulumi:"accessTokenType"`
	// Additional origins
	AdditionalOrigins []string `pulumi:"additionalOrigins"`
	// App type, supported values: OIDC*APP*TYPE*WEB, OIDC*APP*TYPE*USER*AGENT, OIDC*APP*TYPE*NATIVE
	AppType *string `pulumi:"appType"`
	// Auth method type, supported values: OIDC*AUTH*METHOD*TYPE*BASIC, OIDC*AUTH*METHOD*TYPE*POST, OIDC*AUTH*METHOD*TYPE*NONE, OIDC*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType *string `pulumi:"authMethodType"`
	// generated ID for this config
	ClientId *string `pulumi:"clientId"`
	// generated secret for this config
	ClientSecret *string `pulumi:"clientSecret"`
	// Clockskew
	ClockSkew *string `pulumi:"clockSkew"`
	// Dev mode
	DevMode *bool `pulumi:"devMode"`
	// Grant types, supported values: OIDC*GRANT*TYPE*AUTHORIZATION*CODE, OIDC*GRANT*TYPE*IMPLICIT, OIDC*GRANT*TYPE*REFRESH*TOKEN, OIDC*GRANT*TYPE*DEVICE_CODE
	GrantTypes []string `pulumi:"grantTypes"`
	// ID token role assertion
	IdTokenRoleAssertion *bool `pulumi:"idTokenRoleAssertion"`
	// Token userinfo assertion
	IdTokenUserinfoAssertion *bool `pulumi:"idTokenUserinfoAssertion"`
	// Name of the application
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Post logout redirect URIs
	PostLogoutRedirectUris []string `pulumi:"postLogoutRedirectUris"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
	// RedirectURIs
	RedirectUris []string `pulumi:"redirectUris"`
	// Response type, supported values: OIDC*RESPONSE*TYPE*CODE, OIDC*RESPONSE*TYPE*ID*TOKEN, OIDC*RESPONSE*TYPE*ID*TOKEN*TOKEN
	ResponseTypes []string `pulumi:"responseTypes"`
	// Version, supported values: OIDC*VERSION*1_0
	Version *string `pulumi:"version"`
}

type ApplicationOidcState struct {
	// Access token role assertion
	AccessTokenRoleAssertion pulumi.BoolPtrInput
	// Access token type, supported values: OIDC*TOKEN*TYPE*BEARER, OIDC*TOKEN*TYPE*JWT
	AccessTokenType pulumi.StringPtrInput
	// Additional origins
	AdditionalOrigins pulumi.StringArrayInput
	// App type, supported values: OIDC*APP*TYPE*WEB, OIDC*APP*TYPE*USER*AGENT, OIDC*APP*TYPE*NATIVE
	AppType pulumi.StringPtrInput
	// Auth method type, supported values: OIDC*AUTH*METHOD*TYPE*BASIC, OIDC*AUTH*METHOD*TYPE*POST, OIDC*AUTH*METHOD*TYPE*NONE, OIDC*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType pulumi.StringPtrInput
	// generated ID for this config
	ClientId pulumi.StringPtrInput
	// generated secret for this config
	ClientSecret pulumi.StringPtrInput
	// Clockskew
	ClockSkew pulumi.StringPtrInput
	// Dev mode
	DevMode pulumi.BoolPtrInput
	// Grant types, supported values: OIDC*GRANT*TYPE*AUTHORIZATION*CODE, OIDC*GRANT*TYPE*IMPLICIT, OIDC*GRANT*TYPE*REFRESH*TOKEN, OIDC*GRANT*TYPE*DEVICE_CODE
	GrantTypes pulumi.StringArrayInput
	// ID token role assertion
	IdTokenRoleAssertion pulumi.BoolPtrInput
	// Token userinfo assertion
	IdTokenUserinfoAssertion pulumi.BoolPtrInput
	// Name of the application
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Post logout redirect URIs
	PostLogoutRedirectUris pulumi.StringArrayInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
	// RedirectURIs
	RedirectUris pulumi.StringArrayInput
	// Response type, supported values: OIDC*RESPONSE*TYPE*CODE, OIDC*RESPONSE*TYPE*ID*TOKEN, OIDC*RESPONSE*TYPE*ID*TOKEN*TOKEN
	ResponseTypes pulumi.StringArrayInput
	// Version, supported values: OIDC*VERSION*1_0
	Version pulumi.StringPtrInput
}

func (ApplicationOidcState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOidcState)(nil)).Elem()
}

type applicationOidcArgs struct {
	// Access token role assertion
	AccessTokenRoleAssertion *bool `pulumi:"accessTokenRoleAssertion"`
	// Access token type, supported values: OIDC*TOKEN*TYPE*BEARER, OIDC*TOKEN*TYPE*JWT
	AccessTokenType *string `pulumi:"accessTokenType"`
	// Additional origins
	AdditionalOrigins []string `pulumi:"additionalOrigins"`
	// App type, supported values: OIDC*APP*TYPE*WEB, OIDC*APP*TYPE*USER*AGENT, OIDC*APP*TYPE*NATIVE
	AppType *string `pulumi:"appType"`
	// Auth method type, supported values: OIDC*AUTH*METHOD*TYPE*BASIC, OIDC*AUTH*METHOD*TYPE*POST, OIDC*AUTH*METHOD*TYPE*NONE, OIDC*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType *string `pulumi:"authMethodType"`
	// Clockskew
	ClockSkew *string `pulumi:"clockSkew"`
	// Dev mode
	DevMode *bool `pulumi:"devMode"`
	// Grant types, supported values: OIDC*GRANT*TYPE*AUTHORIZATION*CODE, OIDC*GRANT*TYPE*IMPLICIT, OIDC*GRANT*TYPE*REFRESH*TOKEN, OIDC*GRANT*TYPE*DEVICE_CODE
	GrantTypes []string `pulumi:"grantTypes"`
	// ID token role assertion
	IdTokenRoleAssertion *bool `pulumi:"idTokenRoleAssertion"`
	// Token userinfo assertion
	IdTokenUserinfoAssertion *bool `pulumi:"idTokenUserinfoAssertion"`
	// Name of the application
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Post logout redirect URIs
	PostLogoutRedirectUris []string `pulumi:"postLogoutRedirectUris"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
	// RedirectURIs
	RedirectUris []string `pulumi:"redirectUris"`
	// Response type, supported values: OIDC*RESPONSE*TYPE*CODE, OIDC*RESPONSE*TYPE*ID*TOKEN, OIDC*RESPONSE*TYPE*ID*TOKEN*TOKEN
	ResponseTypes []string `pulumi:"responseTypes"`
	// Version, supported values: OIDC*VERSION*1_0
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a ApplicationOidc resource.
type ApplicationOidcArgs struct {
	// Access token role assertion
	AccessTokenRoleAssertion pulumi.BoolPtrInput
	// Access token type, supported values: OIDC*TOKEN*TYPE*BEARER, OIDC*TOKEN*TYPE*JWT
	AccessTokenType pulumi.StringPtrInput
	// Additional origins
	AdditionalOrigins pulumi.StringArrayInput
	// App type, supported values: OIDC*APP*TYPE*WEB, OIDC*APP*TYPE*USER*AGENT, OIDC*APP*TYPE*NATIVE
	AppType pulumi.StringPtrInput
	// Auth method type, supported values: OIDC*AUTH*METHOD*TYPE*BASIC, OIDC*AUTH*METHOD*TYPE*POST, OIDC*AUTH*METHOD*TYPE*NONE, OIDC*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType pulumi.StringPtrInput
	// Clockskew
	ClockSkew pulumi.StringPtrInput
	// Dev mode
	DevMode pulumi.BoolPtrInput
	// Grant types, supported values: OIDC*GRANT*TYPE*AUTHORIZATION*CODE, OIDC*GRANT*TYPE*IMPLICIT, OIDC*GRANT*TYPE*REFRESH*TOKEN, OIDC*GRANT*TYPE*DEVICE_CODE
	GrantTypes pulumi.StringArrayInput
	// ID token role assertion
	IdTokenRoleAssertion pulumi.BoolPtrInput
	// Token userinfo assertion
	IdTokenUserinfoAssertion pulumi.BoolPtrInput
	// Name of the application
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Post logout redirect URIs
	PostLogoutRedirectUris pulumi.StringArrayInput
	// ID of the project
	ProjectId pulumi.StringInput
	// RedirectURIs
	RedirectUris pulumi.StringArrayInput
	// Response type, supported values: OIDC*RESPONSE*TYPE*CODE, OIDC*RESPONSE*TYPE*ID*TOKEN, OIDC*RESPONSE*TYPE*ID*TOKEN*TOKEN
	ResponseTypes pulumi.StringArrayInput
	// Version, supported values: OIDC*VERSION*1_0
	Version pulumi.StringPtrInput
}

func (ApplicationOidcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOidcArgs)(nil)).Elem()
}

type ApplicationOidcInput interface {
	pulumi.Input

	ToApplicationOidcOutput() ApplicationOidcOutput
	ToApplicationOidcOutputWithContext(ctx context.Context) ApplicationOidcOutput
}

func (*ApplicationOidc) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOidc)(nil)).Elem()
}

func (i *ApplicationOidc) ToApplicationOidcOutput() ApplicationOidcOutput {
	return i.ToApplicationOidcOutputWithContext(context.Background())
}

func (i *ApplicationOidc) ToApplicationOidcOutputWithContext(ctx context.Context) ApplicationOidcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOidcOutput)
}

func (i *ApplicationOidc) ToOutput(ctx context.Context) pulumix.Output[*ApplicationOidc] {
	return pulumix.Output[*ApplicationOidc]{
		OutputState: i.ToApplicationOidcOutputWithContext(ctx).OutputState,
	}
}

// ApplicationOidcArrayInput is an input type that accepts ApplicationOidcArray and ApplicationOidcArrayOutput values.
// You can construct a concrete instance of `ApplicationOidcArrayInput` via:
//
//	ApplicationOidcArray{ ApplicationOidcArgs{...} }
type ApplicationOidcArrayInput interface {
	pulumi.Input

	ToApplicationOidcArrayOutput() ApplicationOidcArrayOutput
	ToApplicationOidcArrayOutputWithContext(context.Context) ApplicationOidcArrayOutput
}

type ApplicationOidcArray []ApplicationOidcInput

func (ApplicationOidcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationOidc)(nil)).Elem()
}

func (i ApplicationOidcArray) ToApplicationOidcArrayOutput() ApplicationOidcArrayOutput {
	return i.ToApplicationOidcArrayOutputWithContext(context.Background())
}

func (i ApplicationOidcArray) ToApplicationOidcArrayOutputWithContext(ctx context.Context) ApplicationOidcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOidcArrayOutput)
}

func (i ApplicationOidcArray) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationOidc] {
	return pulumix.Output[[]*ApplicationOidc]{
		OutputState: i.ToApplicationOidcArrayOutputWithContext(ctx).OutputState,
	}
}

// ApplicationOidcMapInput is an input type that accepts ApplicationOidcMap and ApplicationOidcMapOutput values.
// You can construct a concrete instance of `ApplicationOidcMapInput` via:
//
//	ApplicationOidcMap{ "key": ApplicationOidcArgs{...} }
type ApplicationOidcMapInput interface {
	pulumi.Input

	ToApplicationOidcMapOutput() ApplicationOidcMapOutput
	ToApplicationOidcMapOutputWithContext(context.Context) ApplicationOidcMapOutput
}

type ApplicationOidcMap map[string]ApplicationOidcInput

func (ApplicationOidcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationOidc)(nil)).Elem()
}

func (i ApplicationOidcMap) ToApplicationOidcMapOutput() ApplicationOidcMapOutput {
	return i.ToApplicationOidcMapOutputWithContext(context.Background())
}

func (i ApplicationOidcMap) ToApplicationOidcMapOutputWithContext(ctx context.Context) ApplicationOidcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOidcMapOutput)
}

func (i ApplicationOidcMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationOidc] {
	return pulumix.Output[map[string]*ApplicationOidc]{
		OutputState: i.ToApplicationOidcMapOutputWithContext(ctx).OutputState,
	}
}

type ApplicationOidcOutput struct{ *pulumi.OutputState }

func (ApplicationOidcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOidc)(nil)).Elem()
}

func (o ApplicationOidcOutput) ToApplicationOidcOutput() ApplicationOidcOutput {
	return o
}

func (o ApplicationOidcOutput) ToApplicationOidcOutputWithContext(ctx context.Context) ApplicationOidcOutput {
	return o
}

func (o ApplicationOidcOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationOidc] {
	return pulumix.Output[*ApplicationOidc]{
		OutputState: o.OutputState,
	}
}

// Access token role assertion
func (o ApplicationOidcOutput) AccessTokenRoleAssertion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.BoolPtrOutput { return v.AccessTokenRoleAssertion }).(pulumi.BoolPtrOutput)
}

// Access token type, supported values: OIDC*TOKEN*TYPE*BEARER, OIDC*TOKEN*TYPE*JWT
func (o ApplicationOidcOutput) AccessTokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringPtrOutput { return v.AccessTokenType }).(pulumi.StringPtrOutput)
}

// Additional origins
func (o ApplicationOidcOutput) AdditionalOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringArrayOutput { return v.AdditionalOrigins }).(pulumi.StringArrayOutput)
}

// App type, supported values: OIDC*APP*TYPE*WEB, OIDC*APP*TYPE*USER*AGENT, OIDC*APP*TYPE*NATIVE
func (o ApplicationOidcOutput) AppType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringPtrOutput { return v.AppType }).(pulumi.StringPtrOutput)
}

// Auth method type, supported values: OIDC*AUTH*METHOD*TYPE*BASIC, OIDC*AUTH*METHOD*TYPE*POST, OIDC*AUTH*METHOD*TYPE*NONE, OIDC*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
func (o ApplicationOidcOutput) AuthMethodType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringPtrOutput { return v.AuthMethodType }).(pulumi.StringPtrOutput)
}

// generated ID for this config
func (o ApplicationOidcOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// generated secret for this config
func (o ApplicationOidcOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// Clockskew
func (o ApplicationOidcOutput) ClockSkew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringPtrOutput { return v.ClockSkew }).(pulumi.StringPtrOutput)
}

// Dev mode
func (o ApplicationOidcOutput) DevMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.BoolPtrOutput { return v.DevMode }).(pulumi.BoolPtrOutput)
}

// Grant types, supported values: OIDC*GRANT*TYPE*AUTHORIZATION*CODE, OIDC*GRANT*TYPE*IMPLICIT, OIDC*GRANT*TYPE*REFRESH*TOKEN, OIDC*GRANT*TYPE*DEVICE_CODE
func (o ApplicationOidcOutput) GrantTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringArrayOutput { return v.GrantTypes }).(pulumi.StringArrayOutput)
}

// ID token role assertion
func (o ApplicationOidcOutput) IdTokenRoleAssertion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.BoolPtrOutput { return v.IdTokenRoleAssertion }).(pulumi.BoolPtrOutput)
}

// Token userinfo assertion
func (o ApplicationOidcOutput) IdTokenUserinfoAssertion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.BoolPtrOutput { return v.IdTokenUserinfoAssertion }).(pulumi.BoolPtrOutput)
}

// Name of the application
func (o ApplicationOidcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o ApplicationOidcOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Post logout redirect URIs
func (o ApplicationOidcOutput) PostLogoutRedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringArrayOutput { return v.PostLogoutRedirectUris }).(pulumi.StringArrayOutput)
}

// ID of the project
func (o ApplicationOidcOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// RedirectURIs
func (o ApplicationOidcOutput) RedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringArrayOutput { return v.RedirectUris }).(pulumi.StringArrayOutput)
}

// Response type, supported values: OIDC*RESPONSE*TYPE*CODE, OIDC*RESPONSE*TYPE*ID*TOKEN, OIDC*RESPONSE*TYPE*ID*TOKEN*TOKEN
func (o ApplicationOidcOutput) ResponseTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringArrayOutput { return v.ResponseTypes }).(pulumi.StringArrayOutput)
}

// Version, supported values: OIDC*VERSION*1_0
func (o ApplicationOidcOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationOidc) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

type ApplicationOidcArrayOutput struct{ *pulumi.OutputState }

func (ApplicationOidcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationOidc)(nil)).Elem()
}

func (o ApplicationOidcArrayOutput) ToApplicationOidcArrayOutput() ApplicationOidcArrayOutput {
	return o
}

func (o ApplicationOidcArrayOutput) ToApplicationOidcArrayOutputWithContext(ctx context.Context) ApplicationOidcArrayOutput {
	return o
}

func (o ApplicationOidcArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationOidc] {
	return pulumix.Output[[]*ApplicationOidc]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationOidcArrayOutput) Index(i pulumi.IntInput) ApplicationOidcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationOidc {
		return vs[0].([]*ApplicationOidc)[vs[1].(int)]
	}).(ApplicationOidcOutput)
}

type ApplicationOidcMapOutput struct{ *pulumi.OutputState }

func (ApplicationOidcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationOidc)(nil)).Elem()
}

func (o ApplicationOidcMapOutput) ToApplicationOidcMapOutput() ApplicationOidcMapOutput {
	return o
}

func (o ApplicationOidcMapOutput) ToApplicationOidcMapOutputWithContext(ctx context.Context) ApplicationOidcMapOutput {
	return o
}

func (o ApplicationOidcMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationOidc] {
	return pulumix.Output[map[string]*ApplicationOidc]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationOidcMapOutput) MapIndex(k pulumi.StringInput) ApplicationOidcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationOidc {
		return vs[0].(map[string]*ApplicationOidc)[vs[1].(string)]
	}).(ApplicationOidcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationOidcInput)(nil)).Elem(), &ApplicationOidc{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationOidcArrayInput)(nil)).Elem(), ApplicationOidcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationOidcMapInput)(nil)).Elem(), ApplicationOidcMap{})
	pulumi.RegisterOutputType(ApplicationOidcOutput{})
	pulumi.RegisterOutputType(ApplicationOidcArrayOutput{})
	pulumi.RegisterOutputType(ApplicationOidcMapOutput{})
}
