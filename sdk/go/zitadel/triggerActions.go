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

// Resource representing triggers, when actions get started
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
//			_, err := zitadel.NewTriggerActions(ctx, "default", &zitadel.TriggerActionsArgs{
//				OrgId:       pulumi.Any(defaultZitadelOrg.Id),
//				FlowType:    pulumi.String("FLOW_TYPE_CUSTOMISE_TOKEN"),
//				TriggerType: pulumi.String("TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION"),
//				ActionIds: pulumi.StringArray{
//					defaultZitadelAction.Id,
//				},
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
// bash The resource can be imported using the ID format `<flow_type:trigger_type[:org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/triggerActions:TriggerActions imported 'FLOW_TYPE_EXTERNAL_AUTHENTICATION:TRIGGER_TYPE_POST_CREATION:123456789012345678'
//
// ```
type TriggerActions struct {
	pulumi.CustomResourceState

	// IDs of the triggered actions
	ActionIds pulumi.StringArrayOutput `pulumi:"actionIds"`
	// Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
	FlowType pulumi.StringOutput `pulumi:"flowType"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
	TriggerType pulumi.StringOutput `pulumi:"triggerType"`
}

// NewTriggerActions registers a new resource with the given unique name, arguments, and options.
func NewTriggerActions(ctx *pulumi.Context,
	name string, args *TriggerActionsArgs, opts ...pulumi.ResourceOption) (*TriggerActions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionIds == nil {
		return nil, errors.New("invalid value for required argument 'ActionIds'")
	}
	if args.FlowType == nil {
		return nil, errors.New("invalid value for required argument 'FlowType'")
	}
	if args.TriggerType == nil {
		return nil, errors.New("invalid value for required argument 'TriggerType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TriggerActions
	err := ctx.RegisterResource("zitadel:index/triggerActions:TriggerActions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTriggerActions gets an existing TriggerActions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTriggerActions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerActionsState, opts ...pulumi.ResourceOption) (*TriggerActions, error) {
	var resource TriggerActions
	err := ctx.ReadResource("zitadel:index/triggerActions:TriggerActions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TriggerActions resources.
type triggerActionsState struct {
	// IDs of the triggered actions
	ActionIds []string `pulumi:"actionIds"`
	// Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
	FlowType *string `pulumi:"flowType"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
	TriggerType *string `pulumi:"triggerType"`
}

type TriggerActionsState struct {
	// IDs of the triggered actions
	ActionIds pulumi.StringArrayInput
	// Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
	FlowType pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
	TriggerType pulumi.StringPtrInput
}

func (TriggerActionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerActionsState)(nil)).Elem()
}

type triggerActionsArgs struct {
	// IDs of the triggered actions
	ActionIds []string `pulumi:"actionIds"`
	// Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
	FlowType string `pulumi:"flowType"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
	TriggerType string `pulumi:"triggerType"`
}

// The set of arguments for constructing a TriggerActions resource.
type TriggerActionsArgs struct {
	// IDs of the triggered actions
	ActionIds pulumi.StringArrayInput
	// Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
	FlowType pulumi.StringInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
	TriggerType pulumi.StringInput
}

func (TriggerActionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerActionsArgs)(nil)).Elem()
}

type TriggerActionsInput interface {
	pulumi.Input

	ToTriggerActionsOutput() TriggerActionsOutput
	ToTriggerActionsOutputWithContext(ctx context.Context) TriggerActionsOutput
}

func (*TriggerActions) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerActions)(nil)).Elem()
}

func (i *TriggerActions) ToTriggerActionsOutput() TriggerActionsOutput {
	return i.ToTriggerActionsOutputWithContext(context.Background())
}

func (i *TriggerActions) ToTriggerActionsOutputWithContext(ctx context.Context) TriggerActionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerActionsOutput)
}

func (i *TriggerActions) ToOutput(ctx context.Context) pulumix.Output[*TriggerActions] {
	return pulumix.Output[*TriggerActions]{
		OutputState: i.ToTriggerActionsOutputWithContext(ctx).OutputState,
	}
}

// TriggerActionsArrayInput is an input type that accepts TriggerActionsArray and TriggerActionsArrayOutput values.
// You can construct a concrete instance of `TriggerActionsArrayInput` via:
//
//	TriggerActionsArray{ TriggerActionsArgs{...} }
type TriggerActionsArrayInput interface {
	pulumi.Input

	ToTriggerActionsArrayOutput() TriggerActionsArrayOutput
	ToTriggerActionsArrayOutputWithContext(context.Context) TriggerActionsArrayOutput
}

type TriggerActionsArray []TriggerActionsInput

func (TriggerActionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TriggerActions)(nil)).Elem()
}

func (i TriggerActionsArray) ToTriggerActionsArrayOutput() TriggerActionsArrayOutput {
	return i.ToTriggerActionsArrayOutputWithContext(context.Background())
}

func (i TriggerActionsArray) ToTriggerActionsArrayOutputWithContext(ctx context.Context) TriggerActionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerActionsArrayOutput)
}

func (i TriggerActionsArray) ToOutput(ctx context.Context) pulumix.Output[[]*TriggerActions] {
	return pulumix.Output[[]*TriggerActions]{
		OutputState: i.ToTriggerActionsArrayOutputWithContext(ctx).OutputState,
	}
}

// TriggerActionsMapInput is an input type that accepts TriggerActionsMap and TriggerActionsMapOutput values.
// You can construct a concrete instance of `TriggerActionsMapInput` via:
//
//	TriggerActionsMap{ "key": TriggerActionsArgs{...} }
type TriggerActionsMapInput interface {
	pulumi.Input

	ToTriggerActionsMapOutput() TriggerActionsMapOutput
	ToTriggerActionsMapOutputWithContext(context.Context) TriggerActionsMapOutput
}

type TriggerActionsMap map[string]TriggerActionsInput

func (TriggerActionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TriggerActions)(nil)).Elem()
}

func (i TriggerActionsMap) ToTriggerActionsMapOutput() TriggerActionsMapOutput {
	return i.ToTriggerActionsMapOutputWithContext(context.Background())
}

func (i TriggerActionsMap) ToTriggerActionsMapOutputWithContext(ctx context.Context) TriggerActionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerActionsMapOutput)
}

func (i TriggerActionsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*TriggerActions] {
	return pulumix.Output[map[string]*TriggerActions]{
		OutputState: i.ToTriggerActionsMapOutputWithContext(ctx).OutputState,
	}
}

type TriggerActionsOutput struct{ *pulumi.OutputState }

func (TriggerActionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerActions)(nil)).Elem()
}

func (o TriggerActionsOutput) ToTriggerActionsOutput() TriggerActionsOutput {
	return o
}

func (o TriggerActionsOutput) ToTriggerActionsOutputWithContext(ctx context.Context) TriggerActionsOutput {
	return o
}

func (o TriggerActionsOutput) ToOutput(ctx context.Context) pulumix.Output[*TriggerActions] {
	return pulumix.Output[*TriggerActions]{
		OutputState: o.OutputState,
	}
}

// IDs of the triggered actions
func (o TriggerActionsOutput) ActionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TriggerActions) pulumi.StringArrayOutput { return v.ActionIds }).(pulumi.StringArrayOutput)
}

// Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
func (o TriggerActionsOutput) FlowType() pulumi.StringOutput {
	return o.ApplyT(func(v *TriggerActions) pulumi.StringOutput { return v.FlowType }).(pulumi.StringOutput)
}

// ID of the organization
func (o TriggerActionsOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerActions) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
func (o TriggerActionsOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v *TriggerActions) pulumi.StringOutput { return v.TriggerType }).(pulumi.StringOutput)
}

type TriggerActionsArrayOutput struct{ *pulumi.OutputState }

func (TriggerActionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TriggerActions)(nil)).Elem()
}

func (o TriggerActionsArrayOutput) ToTriggerActionsArrayOutput() TriggerActionsArrayOutput {
	return o
}

func (o TriggerActionsArrayOutput) ToTriggerActionsArrayOutputWithContext(ctx context.Context) TriggerActionsArrayOutput {
	return o
}

func (o TriggerActionsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*TriggerActions] {
	return pulumix.Output[[]*TriggerActions]{
		OutputState: o.OutputState,
	}
}

func (o TriggerActionsArrayOutput) Index(i pulumi.IntInput) TriggerActionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TriggerActions {
		return vs[0].([]*TriggerActions)[vs[1].(int)]
	}).(TriggerActionsOutput)
}

type TriggerActionsMapOutput struct{ *pulumi.OutputState }

func (TriggerActionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TriggerActions)(nil)).Elem()
}

func (o TriggerActionsMapOutput) ToTriggerActionsMapOutput() TriggerActionsMapOutput {
	return o
}

func (o TriggerActionsMapOutput) ToTriggerActionsMapOutputWithContext(ctx context.Context) TriggerActionsMapOutput {
	return o
}

func (o TriggerActionsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*TriggerActions] {
	return pulumix.Output[map[string]*TriggerActions]{
		OutputState: o.OutputState,
	}
}

func (o TriggerActionsMapOutput) MapIndex(k pulumi.StringInput) TriggerActionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TriggerActions {
		return vs[0].(map[string]*TriggerActions)[vs[1].(string)]
	}).(TriggerActionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerActionsInput)(nil)).Elem(), &TriggerActions{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerActionsArrayInput)(nil)).Elem(), TriggerActionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerActionsMapInput)(nil)).Elem(), TriggerActionsMap{})
	pulumi.RegisterOutputType(TriggerActionsOutput{})
	pulumi.RegisterOutputType(TriggerActionsArrayOutput{})
	pulumi.RegisterOutputType(TriggerActionsMapOutput{})
}
