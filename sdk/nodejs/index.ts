// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./action";
export * from "./applicationApi";
export * from "./applicationKey";
export * from "./applicationOidc";
export * from "./defaultDomainPolicy";
export * from "./defaultLabelPolicy";
export * from "./defaultLockoutPolicy";
export * from "./defaultLoginPolicy";
export * from "./defaultNotificationPolicy";
export * from "./defaultPasswordComplexityPolicy";
export * from "./defaultPrivacyPolicy";
export * from "./domain";
export * from "./domainPolicy";
export * from "./getAction";
export * from "./getApplicationApi";
export * from "./getApplicationOidc";
export * from "./getHumanUser";
export * from "./getIdpAzureAd";
export * from "./getIdpGithub";
export * from "./getIdpGithubEs";
export * from "./getIdpGitlab";
export * from "./getIdpGitlabSelfHosted";
export * from "./getIdpGoogle";
export * from "./getIdpLdap";
export * from "./getMachineUser";
export * from "./getOrg";
export * from "./getOrgIdpAzureAd";
export * from "./getOrgIdpGithub";
export * from "./getOrgIdpGithubEs";
export * from "./getOrgIdpGitlab";
export * from "./getOrgIdpGitlabSelfHosted";
export * from "./getOrgIdpGoogle";
export * from "./getOrgIdpLdap";
export * from "./getOrgJwtIdp";
export * from "./getOrgOidcIdp";
export * from "./getProject";
export * from "./getProjectRole";
export * from "./getTriggerActions";
export * from "./humanUser";
export * from "./idpAzureAd";
export * from "./idpGithub";
export * from "./idpGithubEs";
export * from "./idpGitlab";
export * from "./idpGitlabSelfHosted";
export * from "./idpGoogle";
export * from "./idpLdap";
export * from "./instanceMember";
export * from "./labelPolicy";
export * from "./lockoutPolicy";
export * from "./loginPolicy";
export * from "./machineKey";
export * from "./machineUser";
export * from "./notificationPolicy";
export * from "./org";
export * from "./orgIdpAzureAd";
export * from "./orgIdpGithub";
export * from "./orgIdpGithubEs";
export * from "./orgIdpGitlab";
export * from "./orgIdpGitlabSelfHosted";
export * from "./orgIdpGoogle";
export * from "./orgIdpJwt";
export * from "./orgIdpLdap";
export * from "./orgIdpOidc";
export * from "./orgMember";
export * from "./passwordComplexityPolicy";
export * from "./personalAccessToken";
export * from "./privacyPolicy";
export * from "./project";
export * from "./projectGrant";
export * from "./projectGrantMember";
export * from "./projectMember";
export * from "./projectRole";
export * from "./provider";
export * from "./smsProviderTwilio";
export * from "./smtpConfig";
export * from "./triggerActions";
export * from "./userGrant";

// Export sub-modules:
import * as config from "./config";

export {
    config,
};

// Import resources to register:
import { Action } from "./action";
import { ApplicationApi } from "./applicationApi";
import { ApplicationKey } from "./applicationKey";
import { ApplicationOidc } from "./applicationOidc";
import { DefaultDomainPolicy } from "./defaultDomainPolicy";
import { DefaultLabelPolicy } from "./defaultLabelPolicy";
import { DefaultLockoutPolicy } from "./defaultLockoutPolicy";
import { DefaultLoginPolicy } from "./defaultLoginPolicy";
import { DefaultNotificationPolicy } from "./defaultNotificationPolicy";
import { DefaultPasswordComplexityPolicy } from "./defaultPasswordComplexityPolicy";
import { DefaultPrivacyPolicy } from "./defaultPrivacyPolicy";
import { Domain } from "./domain";
import { DomainPolicy } from "./domainPolicy";
import { HumanUser } from "./humanUser";
import { IdpAzureAd } from "./idpAzureAd";
import { IdpGithub } from "./idpGithub";
import { IdpGithubEs } from "./idpGithubEs";
import { IdpGitlab } from "./idpGitlab";
import { IdpGitlabSelfHosted } from "./idpGitlabSelfHosted";
import { IdpGoogle } from "./idpGoogle";
import { IdpLdap } from "./idpLdap";
import { InstanceMember } from "./instanceMember";
import { LabelPolicy } from "./labelPolicy";
import { LockoutPolicy } from "./lockoutPolicy";
import { LoginPolicy } from "./loginPolicy";
import { MachineKey } from "./machineKey";
import { MachineUser } from "./machineUser";
import { NotificationPolicy } from "./notificationPolicy";
import { Org } from "./org";
import { OrgIdpAzureAd } from "./orgIdpAzureAd";
import { OrgIdpGithub } from "./orgIdpGithub";
import { OrgIdpGithubEs } from "./orgIdpGithubEs";
import { OrgIdpGitlab } from "./orgIdpGitlab";
import { OrgIdpGitlabSelfHosted } from "./orgIdpGitlabSelfHosted";
import { OrgIdpGoogle } from "./orgIdpGoogle";
import { OrgIdpJwt } from "./orgIdpJwt";
import { OrgIdpLdap } from "./orgIdpLdap";
import { OrgIdpOidc } from "./orgIdpOidc";
import { OrgMember } from "./orgMember";
import { PasswordComplexityPolicy } from "./passwordComplexityPolicy";
import { PersonalAccessToken } from "./personalAccessToken";
import { PrivacyPolicy } from "./privacyPolicy";
import { Project } from "./project";
import { ProjectGrant } from "./projectGrant";
import { ProjectGrantMember } from "./projectGrantMember";
import { ProjectMember } from "./projectMember";
import { ProjectRole } from "./projectRole";
import { SmsProviderTwilio } from "./smsProviderTwilio";
import { SmtpConfig } from "./smtpConfig";
import { TriggerActions } from "./triggerActions";
import { UserGrant } from "./userGrant";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zitadel:index/action:Action":
                return new Action(name, <any>undefined, { urn })
            case "zitadel:index/applicationApi:ApplicationApi":
                return new ApplicationApi(name, <any>undefined, { urn })
            case "zitadel:index/applicationKey:ApplicationKey":
                return new ApplicationKey(name, <any>undefined, { urn })
            case "zitadel:index/applicationOidc:ApplicationOidc":
                return new ApplicationOidc(name, <any>undefined, { urn })
            case "zitadel:index/defaultDomainPolicy:DefaultDomainPolicy":
                return new DefaultDomainPolicy(name, <any>undefined, { urn })
            case "zitadel:index/defaultLabelPolicy:DefaultLabelPolicy":
                return new DefaultLabelPolicy(name, <any>undefined, { urn })
            case "zitadel:index/defaultLockoutPolicy:DefaultLockoutPolicy":
                return new DefaultLockoutPolicy(name, <any>undefined, { urn })
            case "zitadel:index/defaultLoginPolicy:DefaultLoginPolicy":
                return new DefaultLoginPolicy(name, <any>undefined, { urn })
            case "zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy":
                return new DefaultNotificationPolicy(name, <any>undefined, { urn })
            case "zitadel:index/defaultPasswordComplexityPolicy:DefaultPasswordComplexityPolicy":
                return new DefaultPasswordComplexityPolicy(name, <any>undefined, { urn })
            case "zitadel:index/defaultPrivacyPolicy:DefaultPrivacyPolicy":
                return new DefaultPrivacyPolicy(name, <any>undefined, { urn })
            case "zitadel:index/domain:Domain":
                return new Domain(name, <any>undefined, { urn })
            case "zitadel:index/domainPolicy:DomainPolicy":
                return new DomainPolicy(name, <any>undefined, { urn })
            case "zitadel:index/humanUser:HumanUser":
                return new HumanUser(name, <any>undefined, { urn })
            case "zitadel:index/idpAzureAd:IdpAzureAd":
                return new IdpAzureAd(name, <any>undefined, { urn })
            case "zitadel:index/idpGithub:IdpGithub":
                return new IdpGithub(name, <any>undefined, { urn })
            case "zitadel:index/idpGithubEs:IdpGithubEs":
                return new IdpGithubEs(name, <any>undefined, { urn })
            case "zitadel:index/idpGitlab:IdpGitlab":
                return new IdpGitlab(name, <any>undefined, { urn })
            case "zitadel:index/idpGitlabSelfHosted:IdpGitlabSelfHosted":
                return new IdpGitlabSelfHosted(name, <any>undefined, { urn })
            case "zitadel:index/idpGoogle:IdpGoogle":
                return new IdpGoogle(name, <any>undefined, { urn })
            case "zitadel:index/idpLdap:IdpLdap":
                return new IdpLdap(name, <any>undefined, { urn })
            case "zitadel:index/instanceMember:InstanceMember":
                return new InstanceMember(name, <any>undefined, { urn })
            case "zitadel:index/labelPolicy:LabelPolicy":
                return new LabelPolicy(name, <any>undefined, { urn })
            case "zitadel:index/lockoutPolicy:LockoutPolicy":
                return new LockoutPolicy(name, <any>undefined, { urn })
            case "zitadel:index/loginPolicy:LoginPolicy":
                return new LoginPolicy(name, <any>undefined, { urn })
            case "zitadel:index/machineKey:MachineKey":
                return new MachineKey(name, <any>undefined, { urn })
            case "zitadel:index/machineUser:MachineUser":
                return new MachineUser(name, <any>undefined, { urn })
            case "zitadel:index/notificationPolicy:NotificationPolicy":
                return new NotificationPolicy(name, <any>undefined, { urn })
            case "zitadel:index/org:Org":
                return new Org(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpAzureAd:OrgIdpAzureAd":
                return new OrgIdpAzureAd(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpGithub:OrgIdpGithub":
                return new OrgIdpGithub(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpGithubEs:OrgIdpGithubEs":
                return new OrgIdpGithubEs(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpGitlab:OrgIdpGitlab":
                return new OrgIdpGitlab(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpGitlabSelfHosted:OrgIdpGitlabSelfHosted":
                return new OrgIdpGitlabSelfHosted(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpGoogle:OrgIdpGoogle":
                return new OrgIdpGoogle(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpJwt:OrgIdpJwt":
                return new OrgIdpJwt(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpLdap:OrgIdpLdap":
                return new OrgIdpLdap(name, <any>undefined, { urn })
            case "zitadel:index/orgIdpOidc:OrgIdpOidc":
                return new OrgIdpOidc(name, <any>undefined, { urn })
            case "zitadel:index/orgMember:OrgMember":
                return new OrgMember(name, <any>undefined, { urn })
            case "zitadel:index/passwordComplexityPolicy:PasswordComplexityPolicy":
                return new PasswordComplexityPolicy(name, <any>undefined, { urn })
            case "zitadel:index/personalAccessToken:PersonalAccessToken":
                return new PersonalAccessToken(name, <any>undefined, { urn })
            case "zitadel:index/privacyPolicy:PrivacyPolicy":
                return new PrivacyPolicy(name, <any>undefined, { urn })
            case "zitadel:index/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "zitadel:index/projectGrant:ProjectGrant":
                return new ProjectGrant(name, <any>undefined, { urn })
            case "zitadel:index/projectGrantMember:ProjectGrantMember":
                return new ProjectGrantMember(name, <any>undefined, { urn })
            case "zitadel:index/projectMember:ProjectMember":
                return new ProjectMember(name, <any>undefined, { urn })
            case "zitadel:index/projectRole:ProjectRole":
                return new ProjectRole(name, <any>undefined, { urn })
            case "zitadel:index/smsProviderTwilio:SmsProviderTwilio":
                return new SmsProviderTwilio(name, <any>undefined, { urn })
            case "zitadel:index/smtpConfig:SmtpConfig":
                return new SmtpConfig(name, <any>undefined, { urn })
            case "zitadel:index/triggerActions:TriggerActions":
                return new TriggerActions(name, <any>undefined, { urn })
            case "zitadel:index/userGrant:UserGrant":
                return new UserGrant(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zitadel", "index/action", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/applicationApi", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/applicationKey", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/applicationOidc", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/defaultDomainPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/defaultLabelPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/defaultLockoutPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/defaultLoginPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/defaultNotificationPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/defaultPasswordComplexityPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/defaultPrivacyPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/domain", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/domainPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/humanUser", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/idpAzureAd", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/idpGithub", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/idpGithubEs", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/idpGitlab", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/idpGitlabSelfHosted", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/idpGoogle", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/idpLdap", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/instanceMember", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/labelPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/lockoutPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/loginPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/machineKey", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/machineUser", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/notificationPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/org", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpAzureAd", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpGithub", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpGithubEs", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpGitlab", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpGitlabSelfHosted", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpGoogle", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpJwt", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpLdap", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgIdpOidc", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/orgMember", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/passwordComplexityPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/personalAccessToken", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/privacyPolicy", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/project", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/projectGrant", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/projectGrantMember", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/projectMember", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/projectRole", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/smsProviderTwilio", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/smtpConfig", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/triggerActions", _module)
pulumi.runtime.registerResourceModule("zitadel", "index/userGrant", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("zitadel", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:zitadel") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
