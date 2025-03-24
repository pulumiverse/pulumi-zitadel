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

// Resource representing a SAML application belonging to a project, with all configuration possibilities.
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
//			_, err := zitadel.NewApplicationSaml(ctx, "default", &zitadel.ApplicationSamlArgs{
//				OrgId:     pulumi.Any(defaultZitadelOrg.Id),
//				ProjectId: pulumi.Any(defaultZitadelProject.Id),
//				Name:      pulumi.String("applicationapi"),
//				MetadataXml: pulumi.String(`<?xml version="1.0"?>
//
// <md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata"
//
//	                 validUntil="2024-01-26T17:48:38Z"
//	                 cacheDuration="PT604800S"
//	                 entityID="http://example.com/saml/metadata">
//	<md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
//	    <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
//	    <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
//	                                 Location="http://example.com/saml/cas"
//	                                 index="1" />
//
//	</md:SPSSODescriptor>
//
// </md:EntityDescriptor>`),
//
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
// bash The resource can be imported using the ID format `<id:project_id[:org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/applicationSaml:ApplicationSaml imported '123456789012345678:123456789012345678:123456789012345678'
//
// ```
type ApplicationSaml struct {
	pulumi.CustomResourceState

	// Metadata as XML file
	MetadataXml pulumi.StringOutput `pulumi:"metadataXml"`
	// Name of the application
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewApplicationSaml registers a new resource with the given unique name, arguments, and options.
func NewApplicationSaml(ctx *pulumi.Context,
	name string, args *ApplicationSamlArgs, opts ...pulumi.ResourceOption) (*ApplicationSaml, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetadataXml == nil {
		return nil, errors.New("invalid value for required argument 'MetadataXml'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.MetadataXml != nil {
		args.MetadataXml = pulumi.ToSecret(args.MetadataXml).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"metadataXml",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationSaml
	err := ctx.RegisterResource("zitadel:index/applicationSaml:ApplicationSaml", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSaml gets an existing ApplicationSaml resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSaml(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSamlState, opts ...pulumi.ResourceOption) (*ApplicationSaml, error) {
	var resource ApplicationSaml
	err := ctx.ReadResource("zitadel:index/applicationSaml:ApplicationSaml", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSaml resources.
type applicationSamlState struct {
	// Metadata as XML file
	MetadataXml *string `pulumi:"metadataXml"`
	// Name of the application
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
}

type ApplicationSamlState struct {
	// Metadata as XML file
	MetadataXml pulumi.StringPtrInput
	// Name of the application
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
}

func (ApplicationSamlState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSamlState)(nil)).Elem()
}

type applicationSamlArgs struct {
	// Metadata as XML file
	MetadataXml string `pulumi:"metadataXml"`
	// Name of the application
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a ApplicationSaml resource.
type ApplicationSamlArgs struct {
	// Metadata as XML file
	MetadataXml pulumi.StringInput
	// Name of the application
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringInput
}

func (ApplicationSamlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSamlArgs)(nil)).Elem()
}

type ApplicationSamlInput interface {
	pulumi.Input

	ToApplicationSamlOutput() ApplicationSamlOutput
	ToApplicationSamlOutputWithContext(ctx context.Context) ApplicationSamlOutput
}

func (*ApplicationSaml) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSaml)(nil)).Elem()
}

func (i *ApplicationSaml) ToApplicationSamlOutput() ApplicationSamlOutput {
	return i.ToApplicationSamlOutputWithContext(context.Background())
}

func (i *ApplicationSaml) ToApplicationSamlOutputWithContext(ctx context.Context) ApplicationSamlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSamlOutput)
}

func (i *ApplicationSaml) ToOutput(ctx context.Context) pulumix.Output[*ApplicationSaml] {
	return pulumix.Output[*ApplicationSaml]{
		OutputState: i.ToApplicationSamlOutputWithContext(ctx).OutputState,
	}
}

// ApplicationSamlArrayInput is an input type that accepts ApplicationSamlArray and ApplicationSamlArrayOutput values.
// You can construct a concrete instance of `ApplicationSamlArrayInput` via:
//
//	ApplicationSamlArray{ ApplicationSamlArgs{...} }
type ApplicationSamlArrayInput interface {
	pulumi.Input

	ToApplicationSamlArrayOutput() ApplicationSamlArrayOutput
	ToApplicationSamlArrayOutputWithContext(context.Context) ApplicationSamlArrayOutput
}

type ApplicationSamlArray []ApplicationSamlInput

func (ApplicationSamlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSaml)(nil)).Elem()
}

func (i ApplicationSamlArray) ToApplicationSamlArrayOutput() ApplicationSamlArrayOutput {
	return i.ToApplicationSamlArrayOutputWithContext(context.Background())
}

func (i ApplicationSamlArray) ToApplicationSamlArrayOutputWithContext(ctx context.Context) ApplicationSamlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSamlArrayOutput)
}

func (i ApplicationSamlArray) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationSaml] {
	return pulumix.Output[[]*ApplicationSaml]{
		OutputState: i.ToApplicationSamlArrayOutputWithContext(ctx).OutputState,
	}
}

// ApplicationSamlMapInput is an input type that accepts ApplicationSamlMap and ApplicationSamlMapOutput values.
// You can construct a concrete instance of `ApplicationSamlMapInput` via:
//
//	ApplicationSamlMap{ "key": ApplicationSamlArgs{...} }
type ApplicationSamlMapInput interface {
	pulumi.Input

	ToApplicationSamlMapOutput() ApplicationSamlMapOutput
	ToApplicationSamlMapOutputWithContext(context.Context) ApplicationSamlMapOutput
}

type ApplicationSamlMap map[string]ApplicationSamlInput

func (ApplicationSamlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSaml)(nil)).Elem()
}

func (i ApplicationSamlMap) ToApplicationSamlMapOutput() ApplicationSamlMapOutput {
	return i.ToApplicationSamlMapOutputWithContext(context.Background())
}

func (i ApplicationSamlMap) ToApplicationSamlMapOutputWithContext(ctx context.Context) ApplicationSamlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSamlMapOutput)
}

func (i ApplicationSamlMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationSaml] {
	return pulumix.Output[map[string]*ApplicationSaml]{
		OutputState: i.ToApplicationSamlMapOutputWithContext(ctx).OutputState,
	}
}

type ApplicationSamlOutput struct{ *pulumi.OutputState }

func (ApplicationSamlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSaml)(nil)).Elem()
}

func (o ApplicationSamlOutput) ToApplicationSamlOutput() ApplicationSamlOutput {
	return o
}

func (o ApplicationSamlOutput) ToApplicationSamlOutputWithContext(ctx context.Context) ApplicationSamlOutput {
	return o
}

func (o ApplicationSamlOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationSaml] {
	return pulumix.Output[*ApplicationSaml]{
		OutputState: o.OutputState,
	}
}

// Metadata as XML file
func (o ApplicationSamlOutput) MetadataXml() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSaml) pulumi.StringOutput { return v.MetadataXml }).(pulumi.StringOutput)
}

// Name of the application
func (o ApplicationSamlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSaml) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o ApplicationSamlOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSaml) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// ID of the project
func (o ApplicationSamlOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSaml) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

type ApplicationSamlArrayOutput struct{ *pulumi.OutputState }

func (ApplicationSamlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSaml)(nil)).Elem()
}

func (o ApplicationSamlArrayOutput) ToApplicationSamlArrayOutput() ApplicationSamlArrayOutput {
	return o
}

func (o ApplicationSamlArrayOutput) ToApplicationSamlArrayOutputWithContext(ctx context.Context) ApplicationSamlArrayOutput {
	return o
}

func (o ApplicationSamlArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationSaml] {
	return pulumix.Output[[]*ApplicationSaml]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationSamlArrayOutput) Index(i pulumi.IntInput) ApplicationSamlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationSaml {
		return vs[0].([]*ApplicationSaml)[vs[1].(int)]
	}).(ApplicationSamlOutput)
}

type ApplicationSamlMapOutput struct{ *pulumi.OutputState }

func (ApplicationSamlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSaml)(nil)).Elem()
}

func (o ApplicationSamlMapOutput) ToApplicationSamlMapOutput() ApplicationSamlMapOutput {
	return o
}

func (o ApplicationSamlMapOutput) ToApplicationSamlMapOutputWithContext(ctx context.Context) ApplicationSamlMapOutput {
	return o
}

func (o ApplicationSamlMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationSaml] {
	return pulumix.Output[map[string]*ApplicationSaml]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationSamlMapOutput) MapIndex(k pulumi.StringInput) ApplicationSamlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationSaml {
		return vs[0].(map[string]*ApplicationSaml)[vs[1].(string)]
	}).(ApplicationSamlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSamlInput)(nil)).Elem(), &ApplicationSaml{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSamlArrayInput)(nil)).Elem(), ApplicationSamlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSamlMapInput)(nil)).Elem(), ApplicationSamlMap{})
	pulumi.RegisterOutputType(ApplicationSamlOutput{})
	pulumi.RegisterOutputType(ApplicationSamlArrayOutput{})
	pulumi.RegisterOutputType(ApplicationSamlMapOutput{})
}
