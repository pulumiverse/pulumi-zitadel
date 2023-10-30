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

// Resource representing a machine key
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
<<<<<<< HEAD
//			_, err := zitadel.NewMachineKey(ctx, "machineKey", &zitadel.MachineKeyArgs{
//				OrgId:          pulumi.Any(zitadel_org.Org.Id),
//				UserId:         pulumi.Any(zitadel_machine_user.Machine_user.Id),
=======
//			_, err := zitadel.NewMachineKey(ctx, "default", &zitadel.MachineKeyArgs{
//				OrgId:          pulumi.Any(data.Zitadel_org.Default.Id),
//				UserId:         pulumi.Any(data.Zitadel_machine_user.Default.Id),
>>>>>>> origin/master
//				KeyType:        pulumi.String("KEY_TYPE_JSON"),
//				ExpirationDate: pulumi.String("2519-04-01T08:45:00Z"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
<<<<<<< HEAD
=======
// ```
//
// ## Import
//
// terraform The resource can be imported using the ID format `<id:user_id[:org_id][:key_details]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/machineKey:MachineKey imported '123456789012345678:123456789012345678:123456789012345678:{"type":"serviceaccount","keyId":"123456789012345678","key":"-----BEGIN RSA PRIVATE KEY-----\nMIIEpQ...-----END RSA PRIVATE KEY-----\n","userId":"123456789012345678"}'
//
>>>>>>> origin/master
// ```
type MachineKey struct {
	pulumi.CustomResourceState

	// Expiration date of the machine key in the RFC3339 format
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// Value of the machine key
	KeyDetails pulumi.StringOutput `pulumi:"keyDetails"`
	// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
	KeyType pulumi.StringOutput `pulumi:"keyType"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// ID of the user
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewMachineKey registers a new resource with the given unique name, arguments, and options.
func NewMachineKey(ctx *pulumi.Context,
	name string, args *MachineKeyArgs, opts ...pulumi.ResourceOption) (*MachineKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyType == nil {
		return nil, errors.New("invalid value for required argument 'KeyType'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"keyDetails",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MachineKey
	err := ctx.RegisterResource("zitadel:index/machineKey:MachineKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineKey gets an existing MachineKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineKeyState, opts ...pulumi.ResourceOption) (*MachineKey, error) {
	var resource MachineKey
	err := ctx.ReadResource("zitadel:index/machineKey:MachineKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineKey resources.
type machineKeyState struct {
	// Expiration date of the machine key in the RFC3339 format
	ExpirationDate *string `pulumi:"expirationDate"`
	// Value of the machine key
	KeyDetails *string `pulumi:"keyDetails"`
	// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
	KeyType *string `pulumi:"keyType"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the user
	UserId *string `pulumi:"userId"`
}

type MachineKeyState struct {
	// Expiration date of the machine key in the RFC3339 format
	ExpirationDate pulumi.StringPtrInput
	// Value of the machine key
	KeyDetails pulumi.StringPtrInput
	// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
	KeyType pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the user
	UserId pulumi.StringPtrInput
}

func (MachineKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineKeyState)(nil)).Elem()
}

type machineKeyArgs struct {
	// Expiration date of the machine key in the RFC3339 format
	ExpirationDate *string `pulumi:"expirationDate"`
	// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
	KeyType string `pulumi:"keyType"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the user
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a MachineKey resource.
type MachineKeyArgs struct {
	// Expiration date of the machine key in the RFC3339 format
	ExpirationDate pulumi.StringPtrInput
	// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
	KeyType pulumi.StringInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the user
	UserId pulumi.StringInput
}

func (MachineKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineKeyArgs)(nil)).Elem()
}

type MachineKeyInput interface {
	pulumi.Input

	ToMachineKeyOutput() MachineKeyOutput
	ToMachineKeyOutputWithContext(ctx context.Context) MachineKeyOutput
}

func (*MachineKey) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineKey)(nil)).Elem()
}

func (i *MachineKey) ToMachineKeyOutput() MachineKeyOutput {
	return i.ToMachineKeyOutputWithContext(context.Background())
}

func (i *MachineKey) ToMachineKeyOutputWithContext(ctx context.Context) MachineKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineKeyOutput)
}

func (i *MachineKey) ToOutput(ctx context.Context) pulumix.Output[*MachineKey] {
	return pulumix.Output[*MachineKey]{
		OutputState: i.ToMachineKeyOutputWithContext(ctx).OutputState,
	}
}

// MachineKeyArrayInput is an input type that accepts MachineKeyArray and MachineKeyArrayOutput values.
// You can construct a concrete instance of `MachineKeyArrayInput` via:
//
//	MachineKeyArray{ MachineKeyArgs{...} }
type MachineKeyArrayInput interface {
	pulumi.Input

	ToMachineKeyArrayOutput() MachineKeyArrayOutput
	ToMachineKeyArrayOutputWithContext(context.Context) MachineKeyArrayOutput
}

type MachineKeyArray []MachineKeyInput

func (MachineKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineKey)(nil)).Elem()
}

func (i MachineKeyArray) ToMachineKeyArrayOutput() MachineKeyArrayOutput {
	return i.ToMachineKeyArrayOutputWithContext(context.Background())
}

func (i MachineKeyArray) ToMachineKeyArrayOutputWithContext(ctx context.Context) MachineKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineKeyArrayOutput)
}

func (i MachineKeyArray) ToOutput(ctx context.Context) pulumix.Output[[]*MachineKey] {
	return pulumix.Output[[]*MachineKey]{
		OutputState: i.ToMachineKeyArrayOutputWithContext(ctx).OutputState,
	}
}

// MachineKeyMapInput is an input type that accepts MachineKeyMap and MachineKeyMapOutput values.
// You can construct a concrete instance of `MachineKeyMapInput` via:
//
//	MachineKeyMap{ "key": MachineKeyArgs{...} }
type MachineKeyMapInput interface {
	pulumi.Input

	ToMachineKeyMapOutput() MachineKeyMapOutput
	ToMachineKeyMapOutputWithContext(context.Context) MachineKeyMapOutput
}

type MachineKeyMap map[string]MachineKeyInput

func (MachineKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineKey)(nil)).Elem()
}

func (i MachineKeyMap) ToMachineKeyMapOutput() MachineKeyMapOutput {
	return i.ToMachineKeyMapOutputWithContext(context.Background())
}

func (i MachineKeyMap) ToMachineKeyMapOutputWithContext(ctx context.Context) MachineKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineKeyMapOutput)
}

func (i MachineKeyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MachineKey] {
	return pulumix.Output[map[string]*MachineKey]{
		OutputState: i.ToMachineKeyMapOutputWithContext(ctx).OutputState,
	}
}

type MachineKeyOutput struct{ *pulumi.OutputState }

func (MachineKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineKey)(nil)).Elem()
}

func (o MachineKeyOutput) ToMachineKeyOutput() MachineKeyOutput {
	return o
}

func (o MachineKeyOutput) ToMachineKeyOutputWithContext(ctx context.Context) MachineKeyOutput {
	return o
}

func (o MachineKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*MachineKey] {
	return pulumix.Output[*MachineKey]{
		OutputState: o.OutputState,
	}
}

// Expiration date of the machine key in the RFC3339 format
func (o MachineKeyOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineKey) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

// Value of the machine key
func (o MachineKeyOutput) KeyDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineKey) pulumi.StringOutput { return v.KeyDetails }).(pulumi.StringOutput)
}

// Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
func (o MachineKeyOutput) KeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineKey) pulumi.StringOutput { return v.KeyType }).(pulumi.StringOutput)
}

// ID of the organization
func (o MachineKeyOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineKey) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// ID of the user
func (o MachineKeyOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineKey) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type MachineKeyArrayOutput struct{ *pulumi.OutputState }

func (MachineKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineKey)(nil)).Elem()
}

func (o MachineKeyArrayOutput) ToMachineKeyArrayOutput() MachineKeyArrayOutput {
	return o
}

func (o MachineKeyArrayOutput) ToMachineKeyArrayOutputWithContext(ctx context.Context) MachineKeyArrayOutput {
	return o
}

func (o MachineKeyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MachineKey] {
	return pulumix.Output[[]*MachineKey]{
		OutputState: o.OutputState,
	}
}

func (o MachineKeyArrayOutput) Index(i pulumi.IntInput) MachineKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MachineKey {
		return vs[0].([]*MachineKey)[vs[1].(int)]
	}).(MachineKeyOutput)
}

type MachineKeyMapOutput struct{ *pulumi.OutputState }

func (MachineKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineKey)(nil)).Elem()
}

func (o MachineKeyMapOutput) ToMachineKeyMapOutput() MachineKeyMapOutput {
	return o
}

func (o MachineKeyMapOutput) ToMachineKeyMapOutputWithContext(ctx context.Context) MachineKeyMapOutput {
	return o
}

func (o MachineKeyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MachineKey] {
	return pulumix.Output[map[string]*MachineKey]{
		OutputState: o.OutputState,
	}
}

func (o MachineKeyMapOutput) MapIndex(k pulumi.StringInput) MachineKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MachineKey {
		return vs[0].(map[string]*MachineKey)[vs[1].(string)]
	}).(MachineKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineKeyInput)(nil)).Elem(), &MachineKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineKeyArrayInput)(nil)).Elem(), MachineKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineKeyMapInput)(nil)).Elem(), MachineKeyMap{})
	pulumi.RegisterOutputType(MachineKeyOutput{})
	pulumi.RegisterOutputType(MachineKeyArrayOutput{})
	pulumi.RegisterOutputType(MachineKeyMapOutput{})
}
