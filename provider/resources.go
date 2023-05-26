// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zitadel

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumiverse/pulumi-zitadel/provider/pkg/version"
	"github.com/zitadel/terraform-provider-zitadel/zitadel"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "zitadel"
	// modules:
	mainMod = "index" // the zitadel module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(zitadel.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "zitadel",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Description:       "A Pulumi package for creating and managing zitadel cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "zitadel", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/pulumi/pulumi-zitadel",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "zitadel",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"zitadel_org":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Org")},
			"zitadel_human_user":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HumanUser")},
			"zitadel_machine_user":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MachineUser")},
			"zitadel_project":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Project")},
			"zitadel_project_role":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectRole")},
			"zitadel_domain":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Domain")},
			"zitadel_action":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Action")},
			"zitadel_application_oidc":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ApplicationOidc")},
			"zitadel_application_api":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ApplicationApi")},
			"zitadel_application_key":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ApplicationKey")},
			"zitadel_project_grant":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectGrant")},
			"zitadel_user_grant":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "UserGrant")},
			"zitadel_org_member":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgMember")},
			"zitadel_instance_member":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceMember")},
			"zitadel_project_member":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectMember")},
			"zitadel_project_grant_member":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectGrantMember")},
			"zitadel_domain_policy":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DomainPolicy")},
			"zitadel_label_policy":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LabelPolicy")},
			"zitadel_lockout_policy":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LockoutPolicy")},
			"zitadel_login_policy":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LoginPolicy")},
			"zitadel_password_complexity_policy":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PasswordComplexityPolicy")},
			"zitadel_privacy_policy":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PrivacyPolicy")},
			"zitadel_trigger_actions":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TriggerActions")},
			"zitadel_personal_access_token":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PersonalAccessToken")},
			"zitadel_machine_key":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MachineKey")},
			"zitadel_default_label_policy":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DefaultLabelPolicy")},
			"zitadel_default_login_policy":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DefaultLoginPolicy")},
			"zitadel_default_lockout_policy":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DefaultLockoutPolicy")},
			"zitadel_default_domain_policy":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DefaultDomainPolicy")},
			"zitadel_default_privacy_policy":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DefaultPrivacyPolicy")},
			"zitadel_default_password_complexity_policy": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DefaultPasswordComplexityPolicy")},
			"zitadel_sms_provider_twilio":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SmsProviderTwilio")},
			"zitadel_smtp_config":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SmtpConfig")},
			"zitadel_default_notification_policy":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DefaultNotificationPolicy")},
			"zitadel_notification_policy":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NotificationPolicy")},
			"zitadel_idp_github":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdpGithub")},
			"zitadel_idp_github_es":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdpGithubEs")},
			"zitadel_idp_gitlab":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdpGitlab")},
			"zitadel_idp_gitlab_self_hosted":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdpGitlabSelfHosted")},
			"zitadel_idp_google":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdpGoogle")},
			"zitadel_idp_azure_ad":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdpAzureAd")},
			"zitadel_idp_ldap":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdpLdap")},
			"zitadel_org_idp_jwt":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpJwt")},
			"zitadel_org_idp_oidc":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpOidc")},
			"zitadel_org_idp_github":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpGithub")},
			"zitadel_org_idp_github_es":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpGithubEs")},
			"zitadel_org_idp_gitlab":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpGitlab")},
			"zitadel_org_idp_gitlab_self_hosted":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpGitlabSelfHosted")},
			"zitadel_org_idp_google":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpGoogle")},
			"zitadel_org_idp_azure_ad":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpAzureAd")},
			"zitadel_org_idp_ldap":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgIdpLdap")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"zitadel_org":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrg")},
			"zitadel_human_user":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getHumanUser")},
			"zitadel_machine_user":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMachineUser")},
			"zitadel_project":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProject")},
			"zitadel_project_role":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProjectRole")},
			"zitadel_action":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAction")},
			"zitadel_application_oidc":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getApplicationOidc")},
			"zitadel_application_api":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getApplicationApi")},
			"zitadel_trigger_actions":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTriggerActions")},
			"zitadel_idp_github":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIdpGithub")},
			"zitadel_idp_github_es":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIdpGithubEs")},
			"zitadel_idp_gitlab":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIdpGitlab")},
			"zitadel_idp_gitlab_self_hosted":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIdpGitlabSelfHosted")},
			"zitadel_idp_google":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIdpGoogle")},
			"zitadel_idp_azure_ad":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIdpAzureAd")},
			"zitadel_idp_ldap":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIdpLdap")},
			"zitadel_org_jwt_idp":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgJwtIdp")},
			"zitadel_org_oidc_idp":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgOidcIdp")},
			"zitadel_org_idp_github":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgIdpGithub")},
			"zitadel_org_idp_github_es":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgIdpGithubEs")},
			"zitadel_org_idp_gitlab":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgIdpGitlab")},
			"zitadel_org_idp_gitlab_self_hosted": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgIdpGitlabSelfHosted")},
			"zitadel_org_idp_google":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgIdpGoogle")},
			"zitadel_org_idp_azure_ad":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgIdpAzureAd")},
			"zitadel_org_idp_ldap":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgIdpLdap")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/zitadel",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},

			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_zitadel",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
