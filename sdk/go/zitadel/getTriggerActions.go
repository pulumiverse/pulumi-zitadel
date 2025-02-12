// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/scoretechnologies/pulumi-zitadel/sdk/go/zitadel/internal"
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
//	"github.com/scoretechnologies/pulumi-zitadel/sdk/go/zitadel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := zitadel.LookupTriggerActions(ctx, &zitadel.LookupTriggerActionsArgs{
//				OrgId:       pulumi.StringRef(data.Zitadel_org.Default.Id),
//				FlowType:    "FLOW_TYPE_EXTERNAL_AUTHENTICATION",
//				TriggerType: "TRIGGER_TYPE_POST_AUTHENTICATION",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("triggerActions", _default)
//			return nil
//		})
//	}
//
// ```
func LookupTriggerActions(ctx *pulumi.Context, args *LookupTriggerActionsArgs, opts ...pulumi.InvokeOption) (*LookupTriggerActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTriggerActionsResult
	err := ctx.Invoke("zitadel:index/getTriggerActions:getTriggerActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTriggerActions.
type LookupTriggerActionsArgs struct {
	// Type of the flow to which the action triggers belong
	FlowType string `pulumi:"flowType"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Trigger type on when the actions get triggered
	TriggerType string `pulumi:"triggerType"`
}

// A collection of values returned by getTriggerActions.
type LookupTriggerActionsResult struct {
	// IDs of the triggered actions
	ActionIds []string `pulumi:"actionIds"`
	// Type of the flow to which the action triggers belong
	FlowType string `pulumi:"flowType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Trigger type on when the actions get triggered
	TriggerType string `pulumi:"triggerType"`
}

func LookupTriggerActionsOutput(ctx *pulumi.Context, args LookupTriggerActionsOutputArgs, opts ...pulumi.InvokeOption) LookupTriggerActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTriggerActionsResult, error) {
			args := v.(LookupTriggerActionsArgs)
			r, err := LookupTriggerActions(ctx, &args, opts...)
			var s LookupTriggerActionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTriggerActionsResultOutput)
}

// A collection of arguments for invoking getTriggerActions.
type LookupTriggerActionsOutputArgs struct {
	// Type of the flow to which the action triggers belong
	FlowType pulumi.StringInput `pulumi:"flowType"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// Trigger type on when the actions get triggered
	TriggerType pulumi.StringInput `pulumi:"triggerType"`
}

func (LookupTriggerActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTriggerActionsArgs)(nil)).Elem()
}

// A collection of values returned by getTriggerActions.
type LookupTriggerActionsResultOutput struct{ *pulumi.OutputState }

func (LookupTriggerActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTriggerActionsResult)(nil)).Elem()
}

func (o LookupTriggerActionsResultOutput) ToLookupTriggerActionsResultOutput() LookupTriggerActionsResultOutput {
	return o
}

func (o LookupTriggerActionsResultOutput) ToLookupTriggerActionsResultOutputWithContext(ctx context.Context) LookupTriggerActionsResultOutput {
	return o
}

func (o LookupTriggerActionsResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTriggerActionsResult] {
	return pulumix.Output[LookupTriggerActionsResult]{
		OutputState: o.OutputState,
	}
}

// IDs of the triggered actions
func (o LookupTriggerActionsResultOutput) ActionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTriggerActionsResult) []string { return v.ActionIds }).(pulumi.StringArrayOutput)
}

// Type of the flow to which the action triggers belong
func (o LookupTriggerActionsResultOutput) FlowType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerActionsResult) string { return v.FlowType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTriggerActionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerActionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupTriggerActionsResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTriggerActionsResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Trigger type on when the actions get triggered
func (o LookupTriggerActionsResultOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerActionsResult) string { return v.TriggerType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTriggerActionsResultOutput{})
}
