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

// Resource representing a generic JWT IdP of the organization.
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
//			_, err := zitadel.NewOrgIdpJwt(ctx, "default", &zitadel.OrgIdpJwtArgs{
//				OrgId:        pulumi.Any(data.Zitadel_org.Default.Id),
//				StylingType:  pulumi.String("STYLING_TYPE_UNSPECIFIED"),
//				JwtEndpoint:  pulumi.String("https://jwtendpoint.com/jwt"),
//				Issuer:       pulumi.String("https://google.com"),
//				KeysEndpoint: pulumi.String("https://jwtendpoint.com/keys"),
//				HeaderName:   pulumi.String("x-auth-token"),
//				AutoRegister: pulumi.Bool(false),
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
// terraform The resource can be imported using the ID format `<id[:org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/orgIdpJwt:OrgIdpJwt imported '123456789012345678:123456789012345678'
//
// ```
type OrgIdpJwt struct {
	pulumi.CustomResourceState

	// auto register for users from this idp
	AutoRegister pulumi.BoolOutput `pulumi:"autoRegister"`
	// the name of the header where the JWT is sent in, default is authorization
	HeaderName pulumi.StringOutput `pulumi:"headerName"`
	// the issuer of the jwt (for validation)
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// the endpoint where the jwt can be extracted
	JwtEndpoint pulumi.StringOutput `pulumi:"jwtEndpoint"`
	// the endpoint to the key (JWK) which are used to sign the JWT with
	KeysEndpoint pulumi.StringOutput `pulumi:"keysEndpoint"`
	// Name of the IDP
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
	StylingType pulumi.StringOutput `pulumi:"stylingType"`
}

// NewOrgIdpJwt registers a new resource with the given unique name, arguments, and options.
func NewOrgIdpJwt(ctx *pulumi.Context,
	name string, args *OrgIdpJwtArgs, opts ...pulumi.ResourceOption) (*OrgIdpJwt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoRegister == nil {
		return nil, errors.New("invalid value for required argument 'AutoRegister'")
	}
	if args.HeaderName == nil {
		return nil, errors.New("invalid value for required argument 'HeaderName'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	if args.JwtEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'JwtEndpoint'")
	}
	if args.KeysEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'KeysEndpoint'")
	}
	if args.StylingType == nil {
		return nil, errors.New("invalid value for required argument 'StylingType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgIdpJwt
	err := ctx.RegisterResource("zitadel:index/orgIdpJwt:OrgIdpJwt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgIdpJwt gets an existing OrgIdpJwt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgIdpJwt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgIdpJwtState, opts ...pulumi.ResourceOption) (*OrgIdpJwt, error) {
	var resource OrgIdpJwt
	err := ctx.ReadResource("zitadel:index/orgIdpJwt:OrgIdpJwt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgIdpJwt resources.
type orgIdpJwtState struct {
	// auto register for users from this idp
	AutoRegister *bool `pulumi:"autoRegister"`
	// the name of the header where the JWT is sent in, default is authorization
	HeaderName *string `pulumi:"headerName"`
	// the issuer of the jwt (for validation)
	Issuer *string `pulumi:"issuer"`
	// the endpoint where the jwt can be extracted
	JwtEndpoint *string `pulumi:"jwtEndpoint"`
	// the endpoint to the key (JWK) which are used to sign the JWT with
	KeysEndpoint *string `pulumi:"keysEndpoint"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
	StylingType *string `pulumi:"stylingType"`
}

type OrgIdpJwtState struct {
	// auto register for users from this idp
	AutoRegister pulumi.BoolPtrInput
	// the name of the header where the JWT is sent in, default is authorization
	HeaderName pulumi.StringPtrInput
	// the issuer of the jwt (for validation)
	Issuer pulumi.StringPtrInput
	// the endpoint where the jwt can be extracted
	JwtEndpoint pulumi.StringPtrInput
	// the endpoint to the key (JWK) which are used to sign the JWT with
	KeysEndpoint pulumi.StringPtrInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
	StylingType pulumi.StringPtrInput
}

func (OrgIdpJwtState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpJwtState)(nil)).Elem()
}

type orgIdpJwtArgs struct {
	// auto register for users from this idp
	AutoRegister bool `pulumi:"autoRegister"`
	// the name of the header where the JWT is sent in, default is authorization
	HeaderName string `pulumi:"headerName"`
	// the issuer of the jwt (for validation)
	Issuer string `pulumi:"issuer"`
	// the endpoint where the jwt can be extracted
	JwtEndpoint string `pulumi:"jwtEndpoint"`
	// the endpoint to the key (JWK) which are used to sign the JWT with
	KeysEndpoint string `pulumi:"keysEndpoint"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
	StylingType string `pulumi:"stylingType"`
}

// The set of arguments for constructing a OrgIdpJwt resource.
type OrgIdpJwtArgs struct {
	// auto register for users from this idp
	AutoRegister pulumi.BoolInput
	// the name of the header where the JWT is sent in, default is authorization
	HeaderName pulumi.StringInput
	// the issuer of the jwt (for validation)
	Issuer pulumi.StringInput
	// the endpoint where the jwt can be extracted
	JwtEndpoint pulumi.StringInput
	// the endpoint to the key (JWK) which are used to sign the JWT with
	KeysEndpoint pulumi.StringInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
	StylingType pulumi.StringInput
}

func (OrgIdpJwtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpJwtArgs)(nil)).Elem()
}

type OrgIdpJwtInput interface {
	pulumi.Input

	ToOrgIdpJwtOutput() OrgIdpJwtOutput
	ToOrgIdpJwtOutputWithContext(ctx context.Context) OrgIdpJwtOutput
}

func (*OrgIdpJwt) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpJwt)(nil)).Elem()
}

func (i *OrgIdpJwt) ToOrgIdpJwtOutput() OrgIdpJwtOutput {
	return i.ToOrgIdpJwtOutputWithContext(context.Background())
}

func (i *OrgIdpJwt) ToOrgIdpJwtOutputWithContext(ctx context.Context) OrgIdpJwtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpJwtOutput)
}

func (i *OrgIdpJwt) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpJwt] {
	return pulumix.Output[*OrgIdpJwt]{
		OutputState: i.ToOrgIdpJwtOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpJwtArrayInput is an input type that accepts OrgIdpJwtArray and OrgIdpJwtArrayOutput values.
// You can construct a concrete instance of `OrgIdpJwtArrayInput` via:
//
//	OrgIdpJwtArray{ OrgIdpJwtArgs{...} }
type OrgIdpJwtArrayInput interface {
	pulumi.Input

	ToOrgIdpJwtArrayOutput() OrgIdpJwtArrayOutput
	ToOrgIdpJwtArrayOutputWithContext(context.Context) OrgIdpJwtArrayOutput
}

type OrgIdpJwtArray []OrgIdpJwtInput

func (OrgIdpJwtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpJwt)(nil)).Elem()
}

func (i OrgIdpJwtArray) ToOrgIdpJwtArrayOutput() OrgIdpJwtArrayOutput {
	return i.ToOrgIdpJwtArrayOutputWithContext(context.Background())
}

func (i OrgIdpJwtArray) ToOrgIdpJwtArrayOutputWithContext(ctx context.Context) OrgIdpJwtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpJwtArrayOutput)
}

func (i OrgIdpJwtArray) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpJwt] {
	return pulumix.Output[[]*OrgIdpJwt]{
		OutputState: i.ToOrgIdpJwtArrayOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpJwtMapInput is an input type that accepts OrgIdpJwtMap and OrgIdpJwtMapOutput values.
// You can construct a concrete instance of `OrgIdpJwtMapInput` via:
//
//	OrgIdpJwtMap{ "key": OrgIdpJwtArgs{...} }
type OrgIdpJwtMapInput interface {
	pulumi.Input

	ToOrgIdpJwtMapOutput() OrgIdpJwtMapOutput
	ToOrgIdpJwtMapOutputWithContext(context.Context) OrgIdpJwtMapOutput
}

type OrgIdpJwtMap map[string]OrgIdpJwtInput

func (OrgIdpJwtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpJwt)(nil)).Elem()
}

func (i OrgIdpJwtMap) ToOrgIdpJwtMapOutput() OrgIdpJwtMapOutput {
	return i.ToOrgIdpJwtMapOutputWithContext(context.Background())
}

func (i OrgIdpJwtMap) ToOrgIdpJwtMapOutputWithContext(ctx context.Context) OrgIdpJwtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpJwtMapOutput)
}

func (i OrgIdpJwtMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpJwt] {
	return pulumix.Output[map[string]*OrgIdpJwt]{
		OutputState: i.ToOrgIdpJwtMapOutputWithContext(ctx).OutputState,
	}
}

type OrgIdpJwtOutput struct{ *pulumi.OutputState }

func (OrgIdpJwtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpJwt)(nil)).Elem()
}

func (o OrgIdpJwtOutput) ToOrgIdpJwtOutput() OrgIdpJwtOutput {
	return o
}

func (o OrgIdpJwtOutput) ToOrgIdpJwtOutputWithContext(ctx context.Context) OrgIdpJwtOutput {
	return o
}

func (o OrgIdpJwtOutput) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpJwt] {
	return pulumix.Output[*OrgIdpJwt]{
		OutputState: o.OutputState,
	}
}

// auto register for users from this idp
func (o OrgIdpJwtOutput) AutoRegister() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpJwt) pulumi.BoolOutput { return v.AutoRegister }).(pulumi.BoolOutput)
}

// the name of the header where the JWT is sent in, default is authorization
func (o OrgIdpJwtOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpJwt) pulumi.StringOutput { return v.HeaderName }).(pulumi.StringOutput)
}

// the issuer of the jwt (for validation)
func (o OrgIdpJwtOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpJwt) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// the endpoint where the jwt can be extracted
func (o OrgIdpJwtOutput) JwtEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpJwt) pulumi.StringOutput { return v.JwtEndpoint }).(pulumi.StringOutput)
}

// the endpoint to the key (JWK) which are used to sign the JWT with
func (o OrgIdpJwtOutput) KeysEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpJwt) pulumi.StringOutput { return v.KeysEndpoint }).(pulumi.StringOutput)
}

// Name of the IDP
func (o OrgIdpJwtOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpJwt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o OrgIdpJwtOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpJwt) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
func (o OrgIdpJwtOutput) StylingType() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpJwt) pulumi.StringOutput { return v.StylingType }).(pulumi.StringOutput)
}

type OrgIdpJwtArrayOutput struct{ *pulumi.OutputState }

func (OrgIdpJwtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpJwt)(nil)).Elem()
}

func (o OrgIdpJwtArrayOutput) ToOrgIdpJwtArrayOutput() OrgIdpJwtArrayOutput {
	return o
}

func (o OrgIdpJwtArrayOutput) ToOrgIdpJwtArrayOutputWithContext(ctx context.Context) OrgIdpJwtArrayOutput {
	return o
}

func (o OrgIdpJwtArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpJwt] {
	return pulumix.Output[[]*OrgIdpJwt]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpJwtArrayOutput) Index(i pulumi.IntInput) OrgIdpJwtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgIdpJwt {
		return vs[0].([]*OrgIdpJwt)[vs[1].(int)]
	}).(OrgIdpJwtOutput)
}

type OrgIdpJwtMapOutput struct{ *pulumi.OutputState }

func (OrgIdpJwtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpJwt)(nil)).Elem()
}

func (o OrgIdpJwtMapOutput) ToOrgIdpJwtMapOutput() OrgIdpJwtMapOutput {
	return o
}

func (o OrgIdpJwtMapOutput) ToOrgIdpJwtMapOutputWithContext(ctx context.Context) OrgIdpJwtMapOutput {
	return o
}

func (o OrgIdpJwtMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpJwt] {
	return pulumix.Output[map[string]*OrgIdpJwt]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpJwtMapOutput) MapIndex(k pulumi.StringInput) OrgIdpJwtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgIdpJwt {
		return vs[0].(map[string]*OrgIdpJwt)[vs[1].(string)]
	}).(OrgIdpJwtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpJwtInput)(nil)).Elem(), &OrgIdpJwt{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpJwtArrayInput)(nil)).Elem(), OrgIdpJwtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpJwtMapInput)(nil)).Elem(), OrgIdpJwtMap{})
	pulumi.RegisterOutputType(OrgIdpJwtOutput{})
	pulumi.RegisterOutputType(OrgIdpJwtArrayOutput{})
	pulumi.RegisterOutputType(OrgIdpJwtMapOutput{})
}
