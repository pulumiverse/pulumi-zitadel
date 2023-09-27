// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
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
// 		_default, err := zitadel.LookupOrg(ctx, &GetOrgArgs{
// 			Id: "123456789012345678",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("org", _default)
// 		return nil
// 	})
// }
// ```
func LookupOrg(ctx *pulumi.Context, args *LookupOrgArgs, opts ...pulumi.InvokeOption) (*LookupOrgResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupOrgResult
	err := ctx.Invoke("zitadel:index/getOrg:getOrg", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrg.
type LookupOrgArgs struct {
	// ID of the organization
	Id string `pulumi:"id"`
}

// A collection of values returned by getOrg.
type LookupOrgResult struct {
	// ID of the organization
	Id string `pulumi:"id"`
	// Name of the org.
	Name string `pulumi:"name"`
	// Primary domain of the org
	PrimaryDomain string `pulumi:"primaryDomain"`
	// State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
	State string `pulumi:"state"`
}

func LookupOrgOutput(ctx *pulumi.Context, args LookupOrgOutputArgs, opts ...pulumi.InvokeOption) LookupOrgResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgResult, error) {
			args := v.(LookupOrgArgs)
			r, err := LookupOrg(ctx, &args, opts...)
			var s LookupOrgResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgResultOutput)
}

// A collection of arguments for invoking getOrg.
type LookupOrgOutputArgs struct {
	// ID of the organization
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupOrgOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgArgs)(nil)).Elem()
}

// A collection of values returned by getOrg.
type LookupOrgResultOutput struct{ *pulumi.OutputState }

func (LookupOrgResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgResult)(nil)).Elem()
}

func (o LookupOrgResultOutput) ToLookupOrgResultOutput() LookupOrgResultOutput {
	return o
}

func (o LookupOrgResultOutput) ToLookupOrgResultOutputWithContext(ctx context.Context) LookupOrgResultOutput {
	return o
}

// ID of the organization
func (o LookupOrgResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the org.
func (o LookupOrgResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Name }).(pulumi.StringOutput)
}

// Primary domain of the org
func (o LookupOrgResultOutput) PrimaryDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.PrimaryDomain }).(pulumi.StringOutput)
}

// State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
func (o LookupOrgResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgResultOutput{})
}
