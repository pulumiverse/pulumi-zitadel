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

// Resource representing an LDAP IdP on the organization.
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
//			_, err := zitadel.NewOrgIdpLdap(ctx, "default", &zitadel.OrgIdpLdapArgs{
//				OrgId: pulumi.Any(defaultZitadelOrg.Id),
//				Name:  pulumi.String("LDAP"),
//				Servers: pulumi.StringArray{
//					pulumi.String("ldaps://my.primary.server:389"),
//					pulumi.String("ldaps://my.secondary.server:389"),
//				},
//				StartTls:     pulumi.Bool(false),
//				BaseDn:       pulumi.String("dc=example,dc=com"),
//				BindDn:       pulumi.String("cn=admin,dc=example,dc=com"),
//				BindPassword: pulumi.String("Password1!"),
//				UserBase:     pulumi.String("dn"),
//				UserObjectClasses: pulumi.StringArray{
//					pulumi.String("inetOrgPerson"),
//				},
//				UserFilters: pulumi.StringArray{
//					pulumi.String("uid"),
//					pulumi.String("email"),
//				},
//				Timeout:            pulumi.String("10s"),
//				IdAttribute:        pulumi.String("uid"),
//				FirstNameAttribute: pulumi.String("firstname"),
//				LastNameAttribute:  pulumi.String("lastname"),
//				IsLinkingAllowed:   pulumi.Bool(false),
//				IsCreationAllowed:  pulumi.Bool(true),
//				IsAutoCreation:     pulumi.Bool(false),
//				IsAutoUpdate:       pulumi.Bool(true),
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
// bash The resource can be imported using the ID format `<id[:org_id][:bind_password]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/orgIdpLdap:OrgIdpLdap imported '123456789012345678:123456789012345678:b1nd_p4ssw0rd'
//
// ```
type OrgIdpLdap struct {
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
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
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

// NewOrgIdpLdap registers a new resource with the given unique name, arguments, and options.
func NewOrgIdpLdap(ctx *pulumi.Context,
	name string, args *OrgIdpLdapArgs, opts ...pulumi.ResourceOption) (*OrgIdpLdap, error) {
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
	if args.BindPassword != nil {
		args.BindPassword = pulumi.ToSecret(args.BindPassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"bindPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgIdpLdap
	err := ctx.RegisterResource("zitadel:index/orgIdpLdap:OrgIdpLdap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgIdpLdap gets an existing OrgIdpLdap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgIdpLdap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgIdpLdapState, opts ...pulumi.ResourceOption) (*OrgIdpLdap, error) {
	var resource OrgIdpLdap
	err := ctx.ReadResource("zitadel:index/orgIdpLdap:OrgIdpLdap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgIdpLdap resources.
type orgIdpLdapState struct {
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
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
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

type OrgIdpLdapState struct {
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
	// ID of the organization
	OrgId pulumi.StringPtrInput
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

func (OrgIdpLdapState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpLdapState)(nil)).Elem()
}

type orgIdpLdapArgs struct {
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
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
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

// The set of arguments for constructing a OrgIdpLdap resource.
type OrgIdpLdapArgs struct {
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
	// ID of the organization
	OrgId pulumi.StringPtrInput
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

func (OrgIdpLdapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpLdapArgs)(nil)).Elem()
}

type OrgIdpLdapInput interface {
	pulumi.Input

	ToOrgIdpLdapOutput() OrgIdpLdapOutput
	ToOrgIdpLdapOutputWithContext(ctx context.Context) OrgIdpLdapOutput
}

func (*OrgIdpLdap) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpLdap)(nil)).Elem()
}

func (i *OrgIdpLdap) ToOrgIdpLdapOutput() OrgIdpLdapOutput {
	return i.ToOrgIdpLdapOutputWithContext(context.Background())
}

func (i *OrgIdpLdap) ToOrgIdpLdapOutputWithContext(ctx context.Context) OrgIdpLdapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpLdapOutput)
}

func (i *OrgIdpLdap) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpLdap] {
	return pulumix.Output[*OrgIdpLdap]{
		OutputState: i.ToOrgIdpLdapOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpLdapArrayInput is an input type that accepts OrgIdpLdapArray and OrgIdpLdapArrayOutput values.
// You can construct a concrete instance of `OrgIdpLdapArrayInput` via:
//
//	OrgIdpLdapArray{ OrgIdpLdapArgs{...} }
type OrgIdpLdapArrayInput interface {
	pulumi.Input

	ToOrgIdpLdapArrayOutput() OrgIdpLdapArrayOutput
	ToOrgIdpLdapArrayOutputWithContext(context.Context) OrgIdpLdapArrayOutput
}

type OrgIdpLdapArray []OrgIdpLdapInput

func (OrgIdpLdapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpLdap)(nil)).Elem()
}

func (i OrgIdpLdapArray) ToOrgIdpLdapArrayOutput() OrgIdpLdapArrayOutput {
	return i.ToOrgIdpLdapArrayOutputWithContext(context.Background())
}

func (i OrgIdpLdapArray) ToOrgIdpLdapArrayOutputWithContext(ctx context.Context) OrgIdpLdapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpLdapArrayOutput)
}

func (i OrgIdpLdapArray) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpLdap] {
	return pulumix.Output[[]*OrgIdpLdap]{
		OutputState: i.ToOrgIdpLdapArrayOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpLdapMapInput is an input type that accepts OrgIdpLdapMap and OrgIdpLdapMapOutput values.
// You can construct a concrete instance of `OrgIdpLdapMapInput` via:
//
//	OrgIdpLdapMap{ "key": OrgIdpLdapArgs{...} }
type OrgIdpLdapMapInput interface {
	pulumi.Input

	ToOrgIdpLdapMapOutput() OrgIdpLdapMapOutput
	ToOrgIdpLdapMapOutputWithContext(context.Context) OrgIdpLdapMapOutput
}

type OrgIdpLdapMap map[string]OrgIdpLdapInput

func (OrgIdpLdapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpLdap)(nil)).Elem()
}

func (i OrgIdpLdapMap) ToOrgIdpLdapMapOutput() OrgIdpLdapMapOutput {
	return i.ToOrgIdpLdapMapOutputWithContext(context.Background())
}

func (i OrgIdpLdapMap) ToOrgIdpLdapMapOutputWithContext(ctx context.Context) OrgIdpLdapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpLdapMapOutput)
}

func (i OrgIdpLdapMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpLdap] {
	return pulumix.Output[map[string]*OrgIdpLdap]{
		OutputState: i.ToOrgIdpLdapMapOutputWithContext(ctx).OutputState,
	}
}

type OrgIdpLdapOutput struct{ *pulumi.OutputState }

func (OrgIdpLdapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpLdap)(nil)).Elem()
}

func (o OrgIdpLdapOutput) ToOrgIdpLdapOutput() OrgIdpLdapOutput {
	return o
}

func (o OrgIdpLdapOutput) ToOrgIdpLdapOutputWithContext(ctx context.Context) OrgIdpLdapOutput {
	return o
}

func (o OrgIdpLdapOutput) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpLdap] {
	return pulumix.Output[*OrgIdpLdap]{
		OutputState: o.OutputState,
	}
}

// User attribute for the avatar url
func (o OrgIdpLdapOutput) AvatarUrlAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.AvatarUrlAttribute }).(pulumi.StringPtrOutput)
}

// Base DN for LDAP connections
func (o OrgIdpLdapOutput) BaseDn() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringOutput { return v.BaseDn }).(pulumi.StringOutput)
}

// Bind DN for LDAP connections
func (o OrgIdpLdapOutput) BindDn() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringOutput { return v.BindDn }).(pulumi.StringOutput)
}

// Bind password for LDAP connections
func (o OrgIdpLdapOutput) BindPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringOutput { return v.BindPassword }).(pulumi.StringOutput)
}

// User attribute for the display name
func (o OrgIdpLdapOutput) DisplayNameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.DisplayNameAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the email
func (o OrgIdpLdapOutput) EmailAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.EmailAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the email verified state
func (o OrgIdpLdapOutput) EmailVerifiedAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.EmailVerifiedAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the first name
func (o OrgIdpLdapOutput) FirstNameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.FirstNameAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the id
func (o OrgIdpLdapOutput) IdAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.IdAttribute }).(pulumi.StringPtrOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o OrgIdpLdapOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o OrgIdpLdapOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o OrgIdpLdapOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o OrgIdpLdapOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// User attribute for the last name
func (o OrgIdpLdapOutput) LastNameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.LastNameAttribute }).(pulumi.StringPtrOutput)
}

// Name of the IDP
func (o OrgIdpLdapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// User attribute for the nick name
func (o OrgIdpLdapOutput) NickNameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.NickNameAttribute }).(pulumi.StringPtrOutput)
}

// ID of the organization
func (o OrgIdpLdapOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// User attribute for the phone
func (o OrgIdpLdapOutput) PhoneAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.PhoneAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the phone verified state
func (o OrgIdpLdapOutput) PhoneVerifiedAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.PhoneVerifiedAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the preferred language
func (o OrgIdpLdapOutput) PreferredLanguageAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.PreferredLanguageAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the preferred username
func (o OrgIdpLdapOutput) PreferredUsernameAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.PreferredUsernameAttribute }).(pulumi.StringPtrOutput)
}

// User attribute for the profile
func (o OrgIdpLdapOutput) ProfileAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringPtrOutput { return v.ProfileAttribute }).(pulumi.StringPtrOutput)
}

// Servers to try in order for establishing LDAP connections
func (o OrgIdpLdapOutput) Servers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringArrayOutput { return v.Servers }).(pulumi.StringArrayOutput)
}

// Wether to use StartTLS for LDAP connections
func (o OrgIdpLdapOutput) StartTls() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.BoolOutput { return v.StartTls }).(pulumi.BoolOutput)
}

// Timeout for LDAP connections
func (o OrgIdpLdapOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringOutput { return v.Timeout }).(pulumi.StringOutput)
}

// User base for LDAP connections
func (o OrgIdpLdapOutput) UserBase() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringOutput { return v.UserBase }).(pulumi.StringOutput)
}

// User filters for LDAP connections
func (o OrgIdpLdapOutput) UserFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringArrayOutput { return v.UserFilters }).(pulumi.StringArrayOutput)
}

// User object classes for LDAP connections
func (o OrgIdpLdapOutput) UserObjectClasses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgIdpLdap) pulumi.StringArrayOutput { return v.UserObjectClasses }).(pulumi.StringArrayOutput)
}

type OrgIdpLdapArrayOutput struct{ *pulumi.OutputState }

func (OrgIdpLdapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpLdap)(nil)).Elem()
}

func (o OrgIdpLdapArrayOutput) ToOrgIdpLdapArrayOutput() OrgIdpLdapArrayOutput {
	return o
}

func (o OrgIdpLdapArrayOutput) ToOrgIdpLdapArrayOutputWithContext(ctx context.Context) OrgIdpLdapArrayOutput {
	return o
}

func (o OrgIdpLdapArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpLdap] {
	return pulumix.Output[[]*OrgIdpLdap]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpLdapArrayOutput) Index(i pulumi.IntInput) OrgIdpLdapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgIdpLdap {
		return vs[0].([]*OrgIdpLdap)[vs[1].(int)]
	}).(OrgIdpLdapOutput)
}

type OrgIdpLdapMapOutput struct{ *pulumi.OutputState }

func (OrgIdpLdapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpLdap)(nil)).Elem()
}

func (o OrgIdpLdapMapOutput) ToOrgIdpLdapMapOutput() OrgIdpLdapMapOutput {
	return o
}

func (o OrgIdpLdapMapOutput) ToOrgIdpLdapMapOutputWithContext(ctx context.Context) OrgIdpLdapMapOutput {
	return o
}

func (o OrgIdpLdapMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpLdap] {
	return pulumix.Output[map[string]*OrgIdpLdap]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpLdapMapOutput) MapIndex(k pulumi.StringInput) OrgIdpLdapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgIdpLdap {
		return vs[0].(map[string]*OrgIdpLdap)[vs[1].(string)]
	}).(OrgIdpLdapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpLdapInput)(nil)).Elem(), &OrgIdpLdap{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpLdapArrayInput)(nil)).Elem(), OrgIdpLdapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpLdapMapInput)(nil)).Elem(), OrgIdpLdapMap{})
	pulumi.RegisterOutputType(OrgIdpLdapOutput{})
	pulumi.RegisterOutputType(OrgIdpLdapArrayOutput{})
	pulumi.RegisterOutputType(OrgIdpLdapMapOutput{})
}
