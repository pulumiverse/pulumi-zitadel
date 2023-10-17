// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
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
//			_, err := zitadel.NewOrg(ctx, "default", nil)
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
// terraform # The resource can be imported using the ID format `<id>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/org:Org imported '123456789012345678'
//
// ```
type Org struct {
	pulumi.CustomResourceState

	// Name of the org
	Name pulumi.StringOutput `pulumi:"name"`
	// Primary domain of the org
	PrimaryDomain pulumi.StringOutput `pulumi:"primaryDomain"`
	// State of the org
	State pulumi.StringOutput `pulumi:"state"`
}

// NewOrg registers a new resource with the given unique name, arguments, and options.
func NewOrg(ctx *pulumi.Context,
	name string, args *OrgArgs, opts ...pulumi.ResourceOption) (*Org, error) {
	if args == nil {
		args = &OrgArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Org
	err := ctx.RegisterResource("zitadel:index/org:Org", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrg gets an existing Org resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrg(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgState, opts ...pulumi.ResourceOption) (*Org, error) {
	var resource Org
	err := ctx.ReadResource("zitadel:index/org:Org", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Org resources.
type orgState struct {
	// Name of the org
	Name *string `pulumi:"name"`
	// Primary domain of the org
	PrimaryDomain *string `pulumi:"primaryDomain"`
	// State of the org
	State *string `pulumi:"state"`
}

type OrgState struct {
	// Name of the org
	Name pulumi.StringPtrInput
	// Primary domain of the org
	PrimaryDomain pulumi.StringPtrInput
	// State of the org
	State pulumi.StringPtrInput
}

func (OrgState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgState)(nil)).Elem()
}

type orgArgs struct {
	// Name of the org
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Org resource.
type OrgArgs struct {
	// Name of the org
	Name pulumi.StringPtrInput
}

func (OrgArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgArgs)(nil)).Elem()
}

type OrgInput interface {
	pulumi.Input

	ToOrgOutput() OrgOutput
	ToOrgOutputWithContext(ctx context.Context) OrgOutput
}

func (*Org) ElementType() reflect.Type {
	return reflect.TypeOf((**Org)(nil)).Elem()
}

func (i *Org) ToOrgOutput() OrgOutput {
	return i.ToOrgOutputWithContext(context.Background())
}

func (i *Org) ToOrgOutputWithContext(ctx context.Context) OrgOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgOutput)
}

// OrgArrayInput is an input type that accepts OrgArray and OrgArrayOutput values.
// You can construct a concrete instance of `OrgArrayInput` via:
//
//	OrgArray{ OrgArgs{...} }
type OrgArrayInput interface {
	pulumi.Input

	ToOrgArrayOutput() OrgArrayOutput
	ToOrgArrayOutputWithContext(context.Context) OrgArrayOutput
}

type OrgArray []OrgInput

func (OrgArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Org)(nil)).Elem()
}

func (i OrgArray) ToOrgArrayOutput() OrgArrayOutput {
	return i.ToOrgArrayOutputWithContext(context.Background())
}

func (i OrgArray) ToOrgArrayOutputWithContext(ctx context.Context) OrgArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgArrayOutput)
}

// OrgMapInput is an input type that accepts OrgMap and OrgMapOutput values.
// You can construct a concrete instance of `OrgMapInput` via:
//
//	OrgMap{ "key": OrgArgs{...} }
type OrgMapInput interface {
	pulumi.Input

	ToOrgMapOutput() OrgMapOutput
	ToOrgMapOutputWithContext(context.Context) OrgMapOutput
}

type OrgMap map[string]OrgInput

func (OrgMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Org)(nil)).Elem()
}

func (i OrgMap) ToOrgMapOutput() OrgMapOutput {
	return i.ToOrgMapOutputWithContext(context.Background())
}

func (i OrgMap) ToOrgMapOutputWithContext(ctx context.Context) OrgMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgMapOutput)
}

type OrgOutput struct{ *pulumi.OutputState }

func (OrgOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Org)(nil)).Elem()
}

func (o OrgOutput) ToOrgOutput() OrgOutput {
	return o
}

func (o OrgOutput) ToOrgOutputWithContext(ctx context.Context) OrgOutput {
	return o
}

// Name of the org
func (o OrgOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Org) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Primary domain of the org
func (o OrgOutput) PrimaryDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Org) pulumi.StringOutput { return v.PrimaryDomain }).(pulumi.StringOutput)
}

// State of the org
func (o OrgOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Org) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type OrgArrayOutput struct{ *pulumi.OutputState }

func (OrgArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Org)(nil)).Elem()
}

func (o OrgArrayOutput) ToOrgArrayOutput() OrgArrayOutput {
	return o
}

func (o OrgArrayOutput) ToOrgArrayOutputWithContext(ctx context.Context) OrgArrayOutput {
	return o
}

func (o OrgArrayOutput) Index(i pulumi.IntInput) OrgOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Org {
		return vs[0].([]*Org)[vs[1].(int)]
	}).(OrgOutput)
}

type OrgMapOutput struct{ *pulumi.OutputState }

func (OrgMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Org)(nil)).Elem()
}

func (o OrgMapOutput) ToOrgMapOutput() OrgMapOutput {
	return o
}

func (o OrgMapOutput) ToOrgMapOutputWithContext(ctx context.Context) OrgMapOutput {
	return o
}

func (o OrgMapOutput) MapIndex(k pulumi.StringInput) OrgOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Org {
		return vs[0].(map[string]*Org)[vs[1].(string)]
	}).(OrgOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgInput)(nil)).Elem(), &Org{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgArrayInput)(nil)).Elem(), OrgArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgMapInput)(nil)).Elem(), OrgMap{})
	pulumi.RegisterOutputType(OrgOutput{})
	pulumi.RegisterOutputType(OrgArrayOutput{})
	pulumi.RegisterOutputType(OrgMapOutput{})
}
