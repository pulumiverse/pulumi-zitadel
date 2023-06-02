// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing an LDAP IDP on the instance.
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
// 		_, err := zitadel.NewIdpLdap(ctx, "ldap", &zitadel.IdpLdapArgs{
// 			BaseDn:             pulumi.String("dc=example,dc=com"),
// 			BindDn:             pulumi.String("cn=admin,dc=example,dc=com"),
// 			BindPassword:       pulumi.String("Password1!"),
// 			FirstNameAttribute: pulumi.String("firstname"),
// 			IdAttribute:        pulumi.String("uid"),
// 			IsAutoCreation:     pulumi.Bool(false),
// 			IsAutoUpdate:       pulumi.Bool(true),
// 			IsCreationAllowed:  pulumi.Bool(true),
// 			IsLinkingAllowed:   pulumi.Bool(false),
// 			LastNameAttribute:  pulumi.String("lastname"),
// 			Servers: pulumi.StringArray{
// 				pulumi.String("ldaps://my.primary.server:389"),
// 				pulumi.String("ldaps://my.secondary.server:389"),
// 			},
// 			StartTls: pulumi.Bool(false),
// 			Timeout:  pulumi.String("10s"),
// 			UserBase: pulumi.String("dn"),
// 			UserFilters: pulumi.StringArray{
// 				pulumi.String("uid"),
// 				pulumi.String("email"),
// 			},
// 			UserObjectClasses: pulumi.StringArray{
// 				pulumi.String("inetOrgPerson"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IdpLdap struct {
	pulumi.CustomResourceState

	// User attribute for the avatar url
	AvatarUrlAttribute pulumi.StringPtrOutput `pulumi:"avatarUrlAttribute"`
	// Base DN for LDAP connections
	BaseDn pulumi.StringOutput `pulumi:"baseDn"`
	// Bind DN for LDAP connections
	BindDn pulumi.StringOutput `pulumi:"bindDn"`
	// Bind password for LDAP connections
	BindPassword pulumi.StringOutput `pulumi:"bindPassword"`
	// User attribute for the display name
	DisplayNameAttribute pulumi.StringPtrOutput `pulumi:"displayNameAttribute"`
	// User attribute for the email
	EmailAttribute pulumi.StringPtrOutput `pulumi:"emailAttribute"`
	// User attribute for the email verified state
	EmailVerifiedAttribute pulumi.StringPtrOutput `pulumi:"emailVerifiedAttribute"`
	// User attribute for the first name
	FirstNameAttribute pulumi.StringPtrOutput `pulumi:"firstNameAttribute"`
	// User attribute for the id
	IdAttribute pulumi.StringPtrOutput `pulumi:"idAttribute"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolOutput `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolOutput `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolOutput `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolOutput `pulumi:"isLinkingAllowed"`
	// User attribute for the last name
	LastNameAttribute pulumi.StringPtrOutput `pulumi:"lastNameAttribute"`
	// Name of the IDP
	Name pulumi.StringOutput `pulumi:"name"`
	// User attribute for the nick name
	NickNameAttribute pulumi.StringPtrOutput `pulumi:"nickNameAttribute"`
	// User attribute for the phone
	PhoneAttribute pulumi.StringPtrOutput `pulumi:"phoneAttribute"`
	// User attribute for the phone verified state
	PhoneVerifiedAttribute pulumi.StringPtrOutput `pulumi:"phoneVerifiedAttribute"`
	// User attribute for the preferred language
	PreferredLanguageAttribute pulumi.StringPtrOutput `pulumi:"preferredLanguageAttribute"`
	// User attribute for the preferred username
	PreferredUsernameAttribute pulumi.StringPtrOutput `pulumi:"preferredUsernameAttribute"`
	// User attribute for the profile
	ProfileAttribute pulumi.StringPtrOutput `pulumi:"profileAttribute"`
	// Servers to try in order for establishing LDAP connections
	Servers pulumi.StringArrayOutput `pulumi:"servers"`
	// Wether to use StartTLS for LDAP connections
	StartTls pulumi.BoolOutput `pulumi:"startTls"`
	// Timeout for LDAP connections
	Timeout pulumi.StringOutput `pulumi:"timeout"`
	// User base for LDAP connections
	UserBase pulumi.StringOutput `pulumi:"userBase"`
	// User filters for LDAP connections
	UserFilters pulumi.StringArrayOutput `pulumi:"userFilters"`
	// User object classes for LDAP connections
	UserObjectClasses pulumi.StringArrayOutput `pulumi:"userObjectClasses"`
}

// NewIdpLdap registers a new resource with the given unique name, arguments, and options.
func NewIdpLdap(ctx *pulumi.Context,
	name string, args *IdpLdapArgs, opts ...pulumi.ResourceOption) (*IdpLdap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseDn == nil {
		return nil, errors.New("invalid value for required argument 'BaseDn'")
	}
	if args.BindDn == nil {
		return nil, errors.New("invalid value for required argument 'BindDn'")
	}
	if args.BindPassword == nil {
		return nil, errors.New("invalid value for required argument 'BindPassword'")
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
	if args.Servers == nil {
		return nil, errors.New("invalid value for required argument 'Servers'")
	}
	if args.StartTls == nil {
		return nil, errors.New("invalid value for required argument 'StartTls'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	if args.UserBase == nil {
		return nil, errors.New("invalid value for required argument 'UserBase'")
	}
	if args.UserFilters == nil {
		return nil, errors.New("invalid value for required argument 'UserFilters'")
	}
	if args.UserObjectClasses == nil {
		return nil, errors.New("invalid value for required argument 'UserObjectClasses'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IdpLdap
	err := ctx.RegisterResource("zitadel:index/idpLdap:IdpLdap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdpLdap gets an existing IdpLdap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdpLdap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdpLdapState, opts ...pulumi.ResourceOption) (*IdpLdap, error) {
	var resource IdpLdap
	err := ctx.ReadResource("zitadel:index/idpLdap:IdpLdap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdpLdap resources.
type idpLdapState struct {
	// User attribute for the avatar url
	AvatarUrlAttribute *string `pulumi:"avatarUrlAttribute"`
	// Base DN for LDAP connections
	BaseDn *string `pulumi:"baseDn"`
	// Bind DN for LDAP connections
	BindDn *string `pulumi:"bindDn"`
	// Bind password for LDAP connections
	BindPassword *string `pulumi:"bindPassword"`
	// User attribute for the display name
	DisplayNameAttribute *string `pulumi:"displayNameAttribute"`
	// User attribute for the email
	EmailAttribute *string `pulumi:"emailAttribute"`
	// User attribute for the email verified state
	EmailVerifiedAttribute *string `pulumi:"emailVerifiedAttribute"`
	// User attribute for the first name
	FirstNameAttribute *string `pulumi:"firstNameAttribute"`
	// User attribute for the id
	IdAttribute *string `pulumi:"idAttribute"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation *bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate *bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed *bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed *bool `pulumi:"isLinkingAllowed"`
	// User attribute for the last name
	LastNameAttribute *string `pulumi:"lastNameAttribute"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// User attribute for the nick name
	NickNameAttribute *string `pulumi:"nickNameAttribute"`
	// User attribute for the phone
	PhoneAttribute *string `pulumi:"phoneAttribute"`
	// User attribute for the phone verified state
	PhoneVerifiedAttribute *string `pulumi:"phoneVerifiedAttribute"`
	// User attribute for the preferred language
	PreferredLanguageAttribute *string `pulumi:"preferredLanguageAttribute"`
	// User attribute for the preferred username
	PreferredUsernameAttribute *string `pulumi:"preferredUsernameAttribute"`
	// User attribute for the profile
	ProfileAttribute *string `pulumi:"profileAttribute"`
	// Servers to try in order for establishing LDAP connections
	Servers []string `pulumi:"servers"`
	// Wether to use StartTLS for LDAP connections
	StartTls *bool `pulumi:"startTls"`
	// Timeout for LDAP connections
	Timeout *string `pulumi:"timeout"`
	// User base for LDAP connections
	UserBase *string `pulumi:"userBase"`
	// User filters for LDAP connections
	UserFilters []string `pulumi:"userFilters"`
	// User object classes for LDAP connections
	UserObjectClasses []string `pulumi:"userObjectClasses"`
}

type IdpLdapState struct {
	// User attribute for the avatar url
	AvatarUrlAttribute pulumi.StringPtrInput
	// Base DN for LDAP connections
	BaseDn pulumi.StringPtrInput
	// Bind DN for LDAP connections
	BindDn pulumi.StringPtrInput
	// Bind password for LDAP connections
	BindPassword pulumi.StringPtrInput
	// User attribute for the display name
	DisplayNameAttribute pulumi.StringPtrInput
	// User attribute for the email
	EmailAttribute pulumi.StringPtrInput
	// User attribute for the email verified state
	EmailVerifiedAttribute pulumi.StringPtrInput
	// User attribute for the first name
	FirstNameAttribute pulumi.StringPtrInput
	// User attribute for the id
	IdAttribute pulumi.StringPtrInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolPtrInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolPtrInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolPtrInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolPtrInput
	// User attribute for the last name
	LastNameAttribute pulumi.StringPtrInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// User attribute for the nick name
	NickNameAttribute pulumi.StringPtrInput
	// User attribute for the phone
	PhoneAttribute pulumi.StringPtrInput
	// User attribute for the phone verified state
	PhoneVerifiedAttribute pulumi.StringPtrInput
	// User attribute for the preferred language
	PreferredLanguageAttribute pulumi.StringPtrInput
	// User attribute for the preferred username
	PreferredUsernameAttribute pulumi.StringPtrInput
	// User attribute for the profile
	ProfileAttribute pulumi.StringPtrInput
	// Servers to try in order for establishing LDAP connections
	Servers pulumi.StringArrayInput
	// Wether to use StartTLS for LDAP connections
	StartTls pulumi.BoolPtrInput
	// Timeout for LDAP connections
	Timeout pulumi.StringPtrInput
	// User base for LDAP connections
	UserBase pulumi.StringPtrInput
	// User filters for LDAP connections
	UserFilters pulumi.StringArrayInput
	// User object classes for LDAP connections
	UserObjectClasses pulumi.StringArrayInput
}

func (IdpLdapState) ElementType() reflect.Type {
	return reflect.TypeOf((*idpLdapState)(nil)).Elem()
}

type idpLdapArgs struct {
	// User attribute for the avatar url
	AvatarUrlAttribute *string `pulumi:"avatarUrlAttribute"`
	// Base DN for LDAP connections
	BaseDn string `pulumi:"baseDn"`
	// Bind DN for LDAP connections
	BindDn string `pulumi:"bindDn"`
	// Bind password for LDAP connections
	BindPassword string `pulumi:"bindPassword"`
	// User attribute for the display name
	DisplayNameAttribute *string `pulumi:"displayNameAttribute"`
	// User attribute for the email
	EmailAttribute *string `pulumi:"emailAttribute"`
	// User attribute for the email verified state
	EmailVerifiedAttribute *string `pulumi:"emailVerifiedAttribute"`
	// User attribute for the first name
	FirstNameAttribute *string `pulumi:"firstNameAttribute"`
	// User attribute for the id
	IdAttribute *string `pulumi:"idAttribute"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// User attribute for the last name
	LastNameAttribute *string `pulumi:"lastNameAttribute"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// User attribute for the nick name
	NickNameAttribute *string `pulumi:"nickNameAttribute"`
	// User attribute for the phone
	PhoneAttribute *string `pulumi:"phoneAttribute"`
	// User attribute for the phone verified state
	PhoneVerifiedAttribute *string `pulumi:"phoneVerifiedAttribute"`
	// User attribute for the preferred language
	PreferredLanguageAttribute *string `pulumi:"preferredLanguageAttribute"`
	// User attribute for the preferred username
	PreferredUsernameAttribute *string `pulumi:"preferredUsernameAttribute"`
	// User attribute for the profile
	ProfileAttribute *string `pulumi:"profileAttribute"`
	// Servers to try in order for establishing LDAP connections
	Servers []string `pulumi:"servers"`
	// Wether to use StartTLS for LDAP connections
	StartTls bool `pulumi:"startTls"`
	// Timeout for LDAP connections
	Timeout string `pulumi:"timeout"`
	// User base for LDAP connections
	UserBase string `pulumi:"userBase"`
	// User filters for LDAP connections
	UserFilters []string `pulumi:"userFilters"`
	// User object classes for LDAP connections
	UserObjectClasses []string `pulumi:"userObjectClasses"`
}

// The set of arguments for constructing a IdpLdap resource.
type IdpLdapArgs struct {
	// User attribute for the avatar url
	AvatarUrlAttribute pulumi.StringPtrInput
	// Base DN for LDAP connections
	BaseDn pulumi.StringInput
	// Bind DN for LDAP connections
	BindDn pulumi.StringInput
	// Bind password for LDAP connections
	BindPassword pulumi.StringInput
	// User attribute for the display name
	DisplayNameAttribute pulumi.StringPtrInput
	// User attribute for the email
	EmailAttribute pulumi.StringPtrInput
	// User attribute for the email verified state
	EmailVerifiedAttribute pulumi.StringPtrInput
	// User attribute for the first name
	FirstNameAttribute pulumi.StringPtrInput
	// User attribute for the id
	IdAttribute pulumi.StringPtrInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolInput
	// User attribute for the last name
	LastNameAttribute pulumi.StringPtrInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// User attribute for the nick name
	NickNameAttribute pulumi.StringPtrInput
	// User attribute for the phone
	PhoneAttribute pulumi.StringPtrInput
	// User attribute for the phone verified state
	PhoneVerifiedAttribute pulumi.StringPtrInput
	// User attribute for the preferred language
	PreferredLanguageAttribute pulumi.StringPtrInput
	// User attribute for the preferred username
	PreferredUsernameAttribute pulumi.StringPtrInput
	// User attribute for the profile
	ProfileAttribute pulumi.StringPtrInput
	// Servers to try in order for establishing LDAP connections
	Servers pulumi.StringArrayInput
	// Wether to use StartTLS for LDAP connections
	StartTls pulumi.BoolInput
	// Timeout for LDAP connections
	Timeout pulumi.StringInput
	// User base for LDAP connections
	UserBase pulumi.StringInput
	// User filters for LDAP connections
	UserFilters pulumi.StringArrayInput
	// User object classes for LDAP connections
	UserObjectClasses pulumi.StringArrayInput
}

func (IdpLdapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*idpLdapArgs)(nil)).Elem()
}

type IdpLdapInput interface {
	pulumi.Input

	ToIdpLdapOutput() IdpLdapOutput
	ToIdpLdapOutputWithContext(ctx context.Context) IdpLdapOutput
}

func (*IdpLdap) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpLdap)(nil)).Elem()
}

func (i *IdpLdap) ToIdpLdapOutput() IdpLdapOutput {
	return i.ToIdpLdapOutputWithContext(context.Background())
}

func (i *IdpLdap) ToIdpLdapOutputWithContext(ctx context.Context) IdpLdapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpLdapOutput)
}

// IdpLdapArrayInput is an input type that accepts IdpLdapArray and IdpLdapArrayOutput values.
// You can construct a concrete instance of `IdpLdapArrayInput` via:
//
//          IdpLdapArray{ IdpLdapArgs{...} }
type IdpLdapArrayInput interface {
	pulumi.Input

	ToIdpLdapArrayOutput() IdpLdapArrayOutput
	ToIdpLdapArrayOutputWithContext(context.Context) IdpLdapArrayOutput
}

type IdpLdapArray []IdpLdapInput

func (IdpLdapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpLdap)(nil)).Elem()
}

func (i IdpLdapArray) ToIdpLdapArrayOutput() IdpLdapArrayOutput {
	return i.ToIdpLdapArrayOutputWithContext(context.Background())
}

func (i IdpLdapArray) ToIdpLdapArrayOutputWithContext(ctx context.Context) IdpLdapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpLdapArrayOutput)
}

// IdpLdapMapInput is an input type that accepts IdpLdapMap and IdpLdapMapOutput values.
// You can construct a concrete instance of `IdpLdapMapInput` via:
//
//          IdpLdapMap{ "key": IdpLdapArgs{...} }
type IdpLdapMapInput interface {
	pulumi.Input

	ToIdpLdapMapOutput() IdpLdapMapOutput
	ToIdpLdapMapOutputWithContext(context.Context) IdpLdapMapOutput
}

type IdpLdapMap map[string]IdpLdapInput

func (IdpLdapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpLdap)(nil)).Elem()
}

func (i IdpLdapMap) ToIdpLdapMapOutput() IdpLdapMapOutput {
	return i.ToIdpLdapMapOutputWithContext(context.Background())
}

func (i IdpLdapMap) ToIdpLdapMapOutputWithContext(ctx context.Context) IdpLdapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpLdapMapOutput)
}

type IdpLdapOutput struct{ *pulumi.OutputState }

func (IdpLdapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpLdap)(nil)).Elem()
}

func (o IdpLdapOutput) ToIdpLdapOutput() IdpLdapOutput {
	return o
}

func (o IdpLdapOutput) ToIdpLdapOutputWithContext(ctx context.Context) IdpLdapOutput {
	return o
}

// User attribute for the avatar url
func (o IdpLdapOutput) AvatarUrlAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.AvatarUrlAttribute }).(pulumi.StringPtrOutput)
}

// Base DN for LDAP connections
func (o IdpLdapOutput) BaseDn() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringOutput { return v.BaseDn }).(pulumi.StringOutput)
}

// Bind DN for LDAP connections
func (o IdpLdapOutput) BindDn() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringOutput { return v.BindDn }).(pulumi.StringOutput)
}

// Bind password for LDAP connections
func (o IdpLdapOutput) BindPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringOutput { return v.BindPassword }).(pulumi.StringOutput)
}

// User attribute for the display name
func (o IdpLdapOutput) DisplayNameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.DisplayNameAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the email
func (o IdpLdapOutput) EmailAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.EmailAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the email verified state
func (o IdpLdapOutput) EmailVerifiedAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.EmailVerifiedAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the first name
func (o IdpLdapOutput) FirstNameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.FirstNameAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the id
func (o IdpLdapOutput) IdAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.IdAttribute }).(pulumi.StringPtrOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o IdpLdapOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o IdpLdapOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o IdpLdapOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o IdpLdapOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// User attribute for the last name
func (o IdpLdapOutput) LastNameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.LastNameAttribute }).(pulumi.StringPtrOutput)
}

// Name of the IDP
func (o IdpLdapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// User attribute for the nick name
func (o IdpLdapOutput) NickNameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.NickNameAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the phone
func (o IdpLdapOutput) PhoneAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.PhoneAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the phone verified state
func (o IdpLdapOutput) PhoneVerifiedAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.PhoneVerifiedAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the preferred language
func (o IdpLdapOutput) PreferredLanguageAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.PreferredLanguageAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the preferred username
func (o IdpLdapOutput) PreferredUsernameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.PreferredUsernameAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the profile
func (o IdpLdapOutput) ProfileAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringPtrOutput { return v.ProfileAttribute }).(pulumi.StringPtrOutput)
}

// Servers to try in order for establishing LDAP connections
func (o IdpLdapOutput) Servers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringArrayOutput { return v.Servers }).(pulumi.StringArrayOutput)
}

// Wether to use StartTLS for LDAP connections
func (o IdpLdapOutput) StartTls() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.BoolOutput { return v.StartTls }).(pulumi.BoolOutput)
}

// Timeout for LDAP connections
func (o IdpLdapOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringOutput { return v.Timeout }).(pulumi.StringOutput)
}

// User base for LDAP connections
func (o IdpLdapOutput) UserBase() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringOutput { return v.UserBase }).(pulumi.StringOutput)
}

// User filters for LDAP connections
func (o IdpLdapOutput) UserFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringArrayOutput { return v.UserFilters }).(pulumi.StringArrayOutput)
}

// User object classes for LDAP connections
func (o IdpLdapOutput) UserObjectClasses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpLdap) pulumi.StringArrayOutput { return v.UserObjectClasses }).(pulumi.StringArrayOutput)
}

type IdpLdapArrayOutput struct{ *pulumi.OutputState }

func (IdpLdapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpLdap)(nil)).Elem()
}

func (o IdpLdapArrayOutput) ToIdpLdapArrayOutput() IdpLdapArrayOutput {
	return o
}

func (o IdpLdapArrayOutput) ToIdpLdapArrayOutputWithContext(ctx context.Context) IdpLdapArrayOutput {
	return o
}

func (o IdpLdapArrayOutput) Index(i pulumi.IntInput) IdpLdapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdpLdap {
		return vs[0].([]*IdpLdap)[vs[1].(int)]
	}).(IdpLdapOutput)
}

type IdpLdapMapOutput struct{ *pulumi.OutputState }

func (IdpLdapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpLdap)(nil)).Elem()
}

func (o IdpLdapMapOutput) ToIdpLdapMapOutput() IdpLdapMapOutput {
	return o
}

func (o IdpLdapMapOutput) ToIdpLdapMapOutputWithContext(ctx context.Context) IdpLdapMapOutput {
	return o
}

func (o IdpLdapMapOutput) MapIndex(k pulumi.StringInput) IdpLdapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdpLdap {
		return vs[0].(map[string]*IdpLdap)[vs[1].(string)]
	}).(IdpLdapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdpLdapInput)(nil)).Elem(), &IdpLdap{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpLdapArrayInput)(nil)).Elem(), IdpLdapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpLdapMapInput)(nil)).Elem(), IdpLdapMap{})
	pulumi.RegisterOutputType(IdpLdapOutput{})
	pulumi.RegisterOutputType(IdpLdapArrayOutput{})
	pulumi.RegisterOutputType(IdpLdapMapOutput{})
}