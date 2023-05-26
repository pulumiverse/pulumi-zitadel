// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing an action belonging to an organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-zitadel/sdk/go/zitadel"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		actionAction, err := zitadel.LookupAction(ctx, &GetActionArgs{
// 			OrgId:    data.Zitadel_org.Org.Id,
// 			ActionId: "177073621691269123",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("action", actionAction)
// 		return nil
// 	})
// }
// ```
func LookupAction(ctx *pulumi.Context, args *LookupActionArgs, opts ...pulumi.InvokeOption) (*LookupActionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupActionResult
	err := ctx.Invoke("zitadel:index/getAction:getAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAction.
type LookupActionArgs struct {
	// The ID of this resource.
	ActionId string `pulumi:"actionId"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getAction.
type LookupActionResult struct {
	// The ID of this resource.
	ActionId string `pulumi:"actionId"`
	// when true, the next action will be called even if this action fails
	AllowedToFail bool `pulumi:"allowedToFail"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// ID of the organization
	OrgId  string `pulumi:"orgId"`
	Script string `pulumi:"script"`
	// the state of the action
	State int `pulumi:"state"`
	// after which time the action will be terminated if not finished
	Timeout string `pulumi:"timeout"`
}

func LookupActionOutput(ctx *pulumi.Context, args LookupActionOutputArgs, opts ...pulumi.InvokeOption) LookupActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActionResult, error) {
			args := v.(LookupActionArgs)
			r, err := LookupAction(ctx, &args, opts...)
			var s LookupActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActionResultOutput)
}

// A collection of arguments for invoking getAction.
type LookupActionOutputArgs struct {
	// The ID of this resource.
	ActionId pulumi.StringInput `pulumi:"actionId"`
	// ID of the organization
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (LookupActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionArgs)(nil)).Elem()
}

// A collection of values returned by getAction.
type LookupActionResultOutput struct{ *pulumi.OutputState }

func (LookupActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionResult)(nil)).Elem()
}

func (o LookupActionResultOutput) ToLookupActionResultOutput() LookupActionResultOutput {
	return o
}

func (o LookupActionResultOutput) ToLookupActionResultOutputWithContext(ctx context.Context) LookupActionResultOutput {
	return o
}

// The ID of this resource.
func (o LookupActionResultOutput) ActionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.ActionId }).(pulumi.StringOutput)
}

// when true, the next action will be called even if this action fails
func (o LookupActionResultOutput) AllowedToFail() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupActionResult) bool { return v.AllowedToFail }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupActionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupActionResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func (o LookupActionResultOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Script }).(pulumi.StringOutput)
}

// the state of the action
func (o LookupActionResultOutput) State() pulumi.IntOutput {
	return o.ApplyT(func(v LookupActionResult) int { return v.State }).(pulumi.IntOutput)
}

// after which time the action will be terminated if not finished
func (o LookupActionResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Timeout }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActionResultOutput{})
}
