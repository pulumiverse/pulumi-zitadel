// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the default login policy.
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
// 		_, err := zitadel.NewDefaultLoginPolicy(ctx, "loginPolicy", &zitadel.DefaultLoginPolicyArgs{
// 			AllowDomainDiscovery:       pulumi.Bool(true),
// 			AllowExternalIdp:           pulumi.Bool(true),
// 			AllowRegister:              pulumi.Bool(true),
// 			DefaultRedirectUri:         pulumi.String("localhost:8080"),
// 			DisableLoginWithEmail:      pulumi.Bool(true),
// 			DisableLoginWithPhone:      pulumi.Bool(true),
// 			ExternalLoginCheckLifetime: pulumi.String("240h0m0s"),
// 			ForceMfa:                   pulumi.Bool(false),
// 			HidePasswordReset:          pulumi.Bool(false),
// 			IgnoreUnknownUsernames:     pulumi.Bool(true),
// 			MfaInitSkipLifetime:        pulumi.String("720h0m0s"),
// 			MultiFactorCheckLifetime:   pulumi.String("24h0m0s"),
// 			MultiFactors: pulumi.StringArray{
// 				pulumi.String("MULTI_FACTOR_TYPE_U2F_WITH_VERIFICATION"),
// 			},
// 			PasswordCheckLifetime:     pulumi.String("240h0m0s"),
// 			PasswordlessType:          pulumi.String("PASSWORDLESS_TYPE_ALLOWED"),
// 			SecondFactorCheckLifetime: pulumi.String("24h0m0s"),
// 			SecondFactors: pulumi.StringArray{
// 				pulumi.String("SECOND_FACTOR_TYPE_OTP"),
// 				pulumi.String("SECOND_FACTOR_TYPE_U2F"),
// 			},
// 			UserLogin: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DefaultLoginPolicy struct {
	pulumi.CustomResourceState

	// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	AllowDomainDiscovery pulumi.BoolPtrOutput `pulumi:"allowDomainDiscovery"`
	// defines if a user is allowed to add a defined identity provider. E.g. Google auth
	AllowExternalIdp pulumi.BoolOutput `pulumi:"allowExternalIdp"`
	// defines if a person is allowed to register a user on this organisation
	AllowRegister pulumi.BoolOutput `pulumi:"allowRegister"`
	// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	DefaultRedirectUri pulumi.StringOutput `pulumi:"defaultRedirectUri"`
	// defines if user can additionally (to the loginname) be identified by their verified email address
	DisableLoginWithEmail pulumi.BoolPtrOutput `pulumi:"disableLoginWithEmail"`
	// defines if user can additionally (to the loginname) be identified by their verified phone number
	DisableLoginWithPhone      pulumi.BoolPtrOutput `pulumi:"disableLoginWithPhone"`
	ExternalLoginCheckLifetime pulumi.StringOutput  `pulumi:"externalLoginCheckLifetime"`
	// defines if a user MUST use a multi factor to log in
	ForceMfa pulumi.BoolOutput `pulumi:"forceMfa"`
	// defines if password reset link should be shown in the login screen
	HidePasswordReset pulumi.BoolOutput `pulumi:"hidePasswordReset"`
	// allowed idps to login or register
	Idps pulumi.StringArrayOutput `pulumi:"idps"`
	// defines if unknown username on login screen directly return an error or always display the password screen
	IgnoreUnknownUsernames   pulumi.BoolOutput   `pulumi:"ignoreUnknownUsernames"`
	MfaInitSkipLifetime      pulumi.StringOutput `pulumi:"mfaInitSkipLifetime"`
	MultiFactorCheckLifetime pulumi.StringOutput `pulumi:"multiFactorCheckLifetime"`
	// allowed multi factors
	MultiFactors          pulumi.StringArrayOutput `pulumi:"multiFactors"`
	PasswordCheckLifetime pulumi.StringOutput      `pulumi:"passwordCheckLifetime"`
	// defines if passwordless is allowed for users
	PasswordlessType          pulumi.StringOutput `pulumi:"passwordlessType"`
	SecondFactorCheckLifetime pulumi.StringOutput `pulumi:"secondFactorCheckLifetime"`
	// allowed second factors
	SecondFactors pulumi.StringArrayOutput `pulumi:"secondFactors"`
	// defines if a user is allowed to login with his username and password
	UserLogin pulumi.BoolOutput `pulumi:"userLogin"`
}

// NewDefaultLoginPolicy registers a new resource with the given unique name, arguments, and options.
func NewDefaultLoginPolicy(ctx *pulumi.Context,
	name string, args *DefaultLoginPolicyArgs, opts ...pulumi.ResourceOption) (*DefaultLoginPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowExternalIdp == nil {
		return nil, errors.New("invalid value for required argument 'AllowExternalIdp'")
	}
	if args.AllowRegister == nil {
		return nil, errors.New("invalid value for required argument 'AllowRegister'")
	}
	if args.DefaultRedirectUri == nil {
		return nil, errors.New("invalid value for required argument 'DefaultRedirectUri'")
	}
	if args.ExternalLoginCheckLifetime == nil {
		return nil, errors.New("invalid value for required argument 'ExternalLoginCheckLifetime'")
	}
	if args.ForceMfa == nil {
		return nil, errors.New("invalid value for required argument 'ForceMfa'")
	}
	if args.HidePasswordReset == nil {
		return nil, errors.New("invalid value for required argument 'HidePasswordReset'")
	}
	if args.IgnoreUnknownUsernames == nil {
		return nil, errors.New("invalid value for required argument 'IgnoreUnknownUsernames'")
	}
	if args.MfaInitSkipLifetime == nil {
		return nil, errors.New("invalid value for required argument 'MfaInitSkipLifetime'")
	}
	if args.MultiFactorCheckLifetime == nil {
		return nil, errors.New("invalid value for required argument 'MultiFactorCheckLifetime'")
	}
	if args.PasswordCheckLifetime == nil {
		return nil, errors.New("invalid value for required argument 'PasswordCheckLifetime'")
	}
	if args.PasswordlessType == nil {
		return nil, errors.New("invalid value for required argument 'PasswordlessType'")
	}
	if args.SecondFactorCheckLifetime == nil {
		return nil, errors.New("invalid value for required argument 'SecondFactorCheckLifetime'")
	}
	if args.UserLogin == nil {
		return nil, errors.New("invalid value for required argument 'UserLogin'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DefaultLoginPolicy
	err := ctx.RegisterResource("zitadel:index/defaultLoginPolicy:DefaultLoginPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultLoginPolicy gets an existing DefaultLoginPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultLoginPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultLoginPolicyState, opts ...pulumi.ResourceOption) (*DefaultLoginPolicy, error) {
	var resource DefaultLoginPolicy
	err := ctx.ReadResource("zitadel:index/defaultLoginPolicy:DefaultLoginPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultLoginPolicy resources.
type defaultLoginPolicyState struct {
	// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	AllowDomainDiscovery *bool `pulumi:"allowDomainDiscovery"`
	// defines if a user is allowed to add a defined identity provider. E.g. Google auth
	AllowExternalIdp *bool `pulumi:"allowExternalIdp"`
	// defines if a person is allowed to register a user on this organisation
	AllowRegister *bool `pulumi:"allowRegister"`
	// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	DefaultRedirectUri *string `pulumi:"defaultRedirectUri"`
	// defines if user can additionally (to the loginname) be identified by their verified email address
	DisableLoginWithEmail *bool `pulumi:"disableLoginWithEmail"`
	// defines if user can additionally (to the loginname) be identified by their verified phone number
	DisableLoginWithPhone      *bool   `pulumi:"disableLoginWithPhone"`
	ExternalLoginCheckLifetime *string `pulumi:"externalLoginCheckLifetime"`
	// defines if a user MUST use a multi factor to log in
	ForceMfa *bool `pulumi:"forceMfa"`
	// defines if password reset link should be shown in the login screen
	HidePasswordReset *bool `pulumi:"hidePasswordReset"`
	// allowed idps to login or register
	Idps []string `pulumi:"idps"`
	// defines if unknown username on login screen directly return an error or always display the password screen
	IgnoreUnknownUsernames   *bool   `pulumi:"ignoreUnknownUsernames"`
	MfaInitSkipLifetime      *string `pulumi:"mfaInitSkipLifetime"`
	MultiFactorCheckLifetime *string `pulumi:"multiFactorCheckLifetime"`
	// allowed multi factors
	MultiFactors          []string `pulumi:"multiFactors"`
	PasswordCheckLifetime *string  `pulumi:"passwordCheckLifetime"`
	// defines if passwordless is allowed for users
	PasswordlessType          *string `pulumi:"passwordlessType"`
	SecondFactorCheckLifetime *string `pulumi:"secondFactorCheckLifetime"`
	// allowed second factors
	SecondFactors []string `pulumi:"secondFactors"`
	// defines if a user is allowed to login with his username and password
	UserLogin *bool `pulumi:"userLogin"`
}

type DefaultLoginPolicyState struct {
	// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	AllowDomainDiscovery pulumi.BoolPtrInput
	// defines if a user is allowed to add a defined identity provider. E.g. Google auth
	AllowExternalIdp pulumi.BoolPtrInput
	// defines if a person is allowed to register a user on this organisation
	AllowRegister pulumi.BoolPtrInput
	// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	DefaultRedirectUri pulumi.StringPtrInput
	// defines if user can additionally (to the loginname) be identified by their verified email address
	DisableLoginWithEmail pulumi.BoolPtrInput
	// defines if user can additionally (to the loginname) be identified by their verified phone number
	DisableLoginWithPhone      pulumi.BoolPtrInput
	ExternalLoginCheckLifetime pulumi.StringPtrInput
	// defines if a user MUST use a multi factor to log in
	ForceMfa pulumi.BoolPtrInput
	// defines if password reset link should be shown in the login screen
	HidePasswordReset pulumi.BoolPtrInput
	// allowed idps to login or register
	Idps pulumi.StringArrayInput
	// defines if unknown username on login screen directly return an error or always display the password screen
	IgnoreUnknownUsernames   pulumi.BoolPtrInput
	MfaInitSkipLifetime      pulumi.StringPtrInput
	MultiFactorCheckLifetime pulumi.StringPtrInput
	// allowed multi factors
	MultiFactors          pulumi.StringArrayInput
	PasswordCheckLifetime pulumi.StringPtrInput
	// defines if passwordless is allowed for users
	PasswordlessType          pulumi.StringPtrInput
	SecondFactorCheckLifetime pulumi.StringPtrInput
	// allowed second factors
	SecondFactors pulumi.StringArrayInput
	// defines if a user is allowed to login with his username and password
	UserLogin pulumi.BoolPtrInput
}

func (DefaultLoginPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultLoginPolicyState)(nil)).Elem()
}

type defaultLoginPolicyArgs struct {
	// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	AllowDomainDiscovery *bool `pulumi:"allowDomainDiscovery"`
	// defines if a user is allowed to add a defined identity provider. E.g. Google auth
	AllowExternalIdp bool `pulumi:"allowExternalIdp"`
	// defines if a person is allowed to register a user on this organisation
	AllowRegister bool `pulumi:"allowRegister"`
	// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	DefaultRedirectUri string `pulumi:"defaultRedirectUri"`
	// defines if user can additionally (to the loginname) be identified by their verified email address
	DisableLoginWithEmail *bool `pulumi:"disableLoginWithEmail"`
	// defines if user can additionally (to the loginname) be identified by their verified phone number
	DisableLoginWithPhone      *bool  `pulumi:"disableLoginWithPhone"`
	ExternalLoginCheckLifetime string `pulumi:"externalLoginCheckLifetime"`
	// defines if a user MUST use a multi factor to log in
	ForceMfa bool `pulumi:"forceMfa"`
	// defines if password reset link should be shown in the login screen
	HidePasswordReset bool `pulumi:"hidePasswordReset"`
	// allowed idps to login or register
	Idps []string `pulumi:"idps"`
	// defines if unknown username on login screen directly return an error or always display the password screen
	IgnoreUnknownUsernames   bool   `pulumi:"ignoreUnknownUsernames"`
	MfaInitSkipLifetime      string `pulumi:"mfaInitSkipLifetime"`
	MultiFactorCheckLifetime string `pulumi:"multiFactorCheckLifetime"`
	// allowed multi factors
	MultiFactors          []string `pulumi:"multiFactors"`
	PasswordCheckLifetime string   `pulumi:"passwordCheckLifetime"`
	// defines if passwordless is allowed for users
	PasswordlessType          string `pulumi:"passwordlessType"`
	SecondFactorCheckLifetime string `pulumi:"secondFactorCheckLifetime"`
	// allowed second factors
	SecondFactors []string `pulumi:"secondFactors"`
	// defines if a user is allowed to login with his username and password
	UserLogin bool `pulumi:"userLogin"`
}

// The set of arguments for constructing a DefaultLoginPolicy resource.
type DefaultLoginPolicyArgs struct {
	// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	AllowDomainDiscovery pulumi.BoolPtrInput
	// defines if a user is allowed to add a defined identity provider. E.g. Google auth
	AllowExternalIdp pulumi.BoolInput
	// defines if a person is allowed to register a user on this organisation
	AllowRegister pulumi.BoolInput
	// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	DefaultRedirectUri pulumi.StringInput
	// defines if user can additionally (to the loginname) be identified by their verified email address
	DisableLoginWithEmail pulumi.BoolPtrInput
	// defines if user can additionally (to the loginname) be identified by their verified phone number
	DisableLoginWithPhone      pulumi.BoolPtrInput
	ExternalLoginCheckLifetime pulumi.StringInput
	// defines if a user MUST use a multi factor to log in
	ForceMfa pulumi.BoolInput
	// defines if password reset link should be shown in the login screen
	HidePasswordReset pulumi.BoolInput
	// allowed idps to login or register
	Idps pulumi.StringArrayInput
	// defines if unknown username on login screen directly return an error or always display the password screen
	IgnoreUnknownUsernames   pulumi.BoolInput
	MfaInitSkipLifetime      pulumi.StringInput
	MultiFactorCheckLifetime pulumi.StringInput
	// allowed multi factors
	MultiFactors          pulumi.StringArrayInput
	PasswordCheckLifetime pulumi.StringInput
	// defines if passwordless is allowed for users
	PasswordlessType          pulumi.StringInput
	SecondFactorCheckLifetime pulumi.StringInput
	// allowed second factors
	SecondFactors pulumi.StringArrayInput
	// defines if a user is allowed to login with his username and password
	UserLogin pulumi.BoolInput
}

func (DefaultLoginPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultLoginPolicyArgs)(nil)).Elem()
}

type DefaultLoginPolicyInput interface {
	pulumi.Input

	ToDefaultLoginPolicyOutput() DefaultLoginPolicyOutput
	ToDefaultLoginPolicyOutputWithContext(ctx context.Context) DefaultLoginPolicyOutput
}

func (*DefaultLoginPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultLoginPolicy)(nil)).Elem()
}

func (i *DefaultLoginPolicy) ToDefaultLoginPolicyOutput() DefaultLoginPolicyOutput {
	return i.ToDefaultLoginPolicyOutputWithContext(context.Background())
}

func (i *DefaultLoginPolicy) ToDefaultLoginPolicyOutputWithContext(ctx context.Context) DefaultLoginPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLoginPolicyOutput)
}

// DefaultLoginPolicyArrayInput is an input type that accepts DefaultLoginPolicyArray and DefaultLoginPolicyArrayOutput values.
// You can construct a concrete instance of `DefaultLoginPolicyArrayInput` via:
//
//          DefaultLoginPolicyArray{ DefaultLoginPolicyArgs{...} }
type DefaultLoginPolicyArrayInput interface {
	pulumi.Input

	ToDefaultLoginPolicyArrayOutput() DefaultLoginPolicyArrayOutput
	ToDefaultLoginPolicyArrayOutputWithContext(context.Context) DefaultLoginPolicyArrayOutput
}

type DefaultLoginPolicyArray []DefaultLoginPolicyInput

func (DefaultLoginPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultLoginPolicy)(nil)).Elem()
}

func (i DefaultLoginPolicyArray) ToDefaultLoginPolicyArrayOutput() DefaultLoginPolicyArrayOutput {
	return i.ToDefaultLoginPolicyArrayOutputWithContext(context.Background())
}

func (i DefaultLoginPolicyArray) ToDefaultLoginPolicyArrayOutputWithContext(ctx context.Context) DefaultLoginPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLoginPolicyArrayOutput)
}

// DefaultLoginPolicyMapInput is an input type that accepts DefaultLoginPolicyMap and DefaultLoginPolicyMapOutput values.
// You can construct a concrete instance of `DefaultLoginPolicyMapInput` via:
//
//          DefaultLoginPolicyMap{ "key": DefaultLoginPolicyArgs{...} }
type DefaultLoginPolicyMapInput interface {
	pulumi.Input

	ToDefaultLoginPolicyMapOutput() DefaultLoginPolicyMapOutput
	ToDefaultLoginPolicyMapOutputWithContext(context.Context) DefaultLoginPolicyMapOutput
}

type DefaultLoginPolicyMap map[string]DefaultLoginPolicyInput

func (DefaultLoginPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultLoginPolicy)(nil)).Elem()
}

func (i DefaultLoginPolicyMap) ToDefaultLoginPolicyMapOutput() DefaultLoginPolicyMapOutput {
	return i.ToDefaultLoginPolicyMapOutputWithContext(context.Background())
}

func (i DefaultLoginPolicyMap) ToDefaultLoginPolicyMapOutputWithContext(ctx context.Context) DefaultLoginPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLoginPolicyMapOutput)
}

type DefaultLoginPolicyOutput struct{ *pulumi.OutputState }

func (DefaultLoginPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultLoginPolicy)(nil)).Elem()
}

func (o DefaultLoginPolicyOutput) ToDefaultLoginPolicyOutput() DefaultLoginPolicyOutput {
	return o
}

func (o DefaultLoginPolicyOutput) ToDefaultLoginPolicyOutputWithContext(ctx context.Context) DefaultLoginPolicyOutput {
	return o
}

// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
func (o DefaultLoginPolicyOutput) AllowDomainDiscovery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolPtrOutput { return v.AllowDomainDiscovery }).(pulumi.BoolPtrOutput)
}

// defines if a user is allowed to add a defined identity provider. E.g. Google auth
func (o DefaultLoginPolicyOutput) AllowExternalIdp() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolOutput { return v.AllowExternalIdp }).(pulumi.BoolOutput)
}

// defines if a person is allowed to register a user on this organisation
func (o DefaultLoginPolicyOutput) AllowRegister() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolOutput { return v.AllowRegister }).(pulumi.BoolOutput)
}

// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
func (o DefaultLoginPolicyOutput) DefaultRedirectUri() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringOutput { return v.DefaultRedirectUri }).(pulumi.StringOutput)
}

// defines if user can additionally (to the loginname) be identified by their verified email address
func (o DefaultLoginPolicyOutput) DisableLoginWithEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolPtrOutput { return v.DisableLoginWithEmail }).(pulumi.BoolPtrOutput)
}

// defines if user can additionally (to the loginname) be identified by their verified phone number
func (o DefaultLoginPolicyOutput) DisableLoginWithPhone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolPtrOutput { return v.DisableLoginWithPhone }).(pulumi.BoolPtrOutput)
}

func (o DefaultLoginPolicyOutput) ExternalLoginCheckLifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringOutput { return v.ExternalLoginCheckLifetime }).(pulumi.StringOutput)
}

// defines if a user MUST use a multi factor to log in
func (o DefaultLoginPolicyOutput) ForceMfa() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolOutput { return v.ForceMfa }).(pulumi.BoolOutput)
}

// defines if password reset link should be shown in the login screen
func (o DefaultLoginPolicyOutput) HidePasswordReset() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolOutput { return v.HidePasswordReset }).(pulumi.BoolOutput)
}

// allowed idps to login or register
func (o DefaultLoginPolicyOutput) Idps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringArrayOutput { return v.Idps }).(pulumi.StringArrayOutput)
}

// defines if unknown username on login screen directly return an error or always display the password screen
func (o DefaultLoginPolicyOutput) IgnoreUnknownUsernames() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolOutput { return v.IgnoreUnknownUsernames }).(pulumi.BoolOutput)
}

func (o DefaultLoginPolicyOutput) MfaInitSkipLifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringOutput { return v.MfaInitSkipLifetime }).(pulumi.StringOutput)
}

func (o DefaultLoginPolicyOutput) MultiFactorCheckLifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringOutput { return v.MultiFactorCheckLifetime }).(pulumi.StringOutput)
}

// allowed multi factors
func (o DefaultLoginPolicyOutput) MultiFactors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringArrayOutput { return v.MultiFactors }).(pulumi.StringArrayOutput)
}

func (o DefaultLoginPolicyOutput) PasswordCheckLifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringOutput { return v.PasswordCheckLifetime }).(pulumi.StringOutput)
}

// defines if passwordless is allowed for users
func (o DefaultLoginPolicyOutput) PasswordlessType() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringOutput { return v.PasswordlessType }).(pulumi.StringOutput)
}

func (o DefaultLoginPolicyOutput) SecondFactorCheckLifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringOutput { return v.SecondFactorCheckLifetime }).(pulumi.StringOutput)
}

// allowed second factors
func (o DefaultLoginPolicyOutput) SecondFactors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.StringArrayOutput { return v.SecondFactors }).(pulumi.StringArrayOutput)
}

// defines if a user is allowed to login with his username and password
func (o DefaultLoginPolicyOutput) UserLogin() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultLoginPolicy) pulumi.BoolOutput { return v.UserLogin }).(pulumi.BoolOutput)
}

type DefaultLoginPolicyArrayOutput struct{ *pulumi.OutputState }

func (DefaultLoginPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultLoginPolicy)(nil)).Elem()
}

func (o DefaultLoginPolicyArrayOutput) ToDefaultLoginPolicyArrayOutput() DefaultLoginPolicyArrayOutput {
	return o
}

func (o DefaultLoginPolicyArrayOutput) ToDefaultLoginPolicyArrayOutputWithContext(ctx context.Context) DefaultLoginPolicyArrayOutput {
	return o
}

func (o DefaultLoginPolicyArrayOutput) Index(i pulumi.IntInput) DefaultLoginPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultLoginPolicy {
		return vs[0].([]*DefaultLoginPolicy)[vs[1].(int)]
	}).(DefaultLoginPolicyOutput)
}

type DefaultLoginPolicyMapOutput struct{ *pulumi.OutputState }

func (DefaultLoginPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultLoginPolicy)(nil)).Elem()
}

func (o DefaultLoginPolicyMapOutput) ToDefaultLoginPolicyMapOutput() DefaultLoginPolicyMapOutput {
	return o
}

func (o DefaultLoginPolicyMapOutput) ToDefaultLoginPolicyMapOutputWithContext(ctx context.Context) DefaultLoginPolicyMapOutput {
	return o
}

func (o DefaultLoginPolicyMapOutput) MapIndex(k pulumi.StringInput) DefaultLoginPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultLoginPolicy {
		return vs[0].(map[string]*DefaultLoginPolicy)[vs[1].(string)]
	}).(DefaultLoginPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLoginPolicyInput)(nil)).Elem(), &DefaultLoginPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLoginPolicyArrayInput)(nil)).Elem(), DefaultLoginPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLoginPolicyMapInput)(nil)).Elem(), DefaultLoginPolicyMap{})
	pulumi.RegisterOutputType(DefaultLoginPolicyOutput{})
	pulumi.RegisterOutputType(DefaultLoginPolicyArrayOutput{})
	pulumi.RegisterOutputType(DefaultLoginPolicyMapOutput{})
}