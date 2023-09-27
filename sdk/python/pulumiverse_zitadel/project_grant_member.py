# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectGrantMemberArgs', 'ProjectGrantMember']

@pulumi.input_type
class ProjectGrantMemberArgs:
    def __init__(__self__, *,
                 grant_id: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 roles: pulumi.Input[Sequence[pulumi.Input[str]]],
                 user_id: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectGrantMember resource.
        :param pulumi.Input[str] grant_id: ID of the grant
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles granted
        :param pulumi.Input[str] user_id: ID of the user
        :param pulumi.Input[str] org_id: ID of the organization
        """
        pulumi.set(__self__, "grant_id", grant_id)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "roles", roles)
        pulumi.set(__self__, "user_id", user_id)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter(name="grantId")
    def grant_id(self) -> pulumi.Input[str]:
        """
        ID of the grant
        """
        return pulumi.get(self, "grant_id")

    @grant_id.setter
    def grant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "grant_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of roles granted
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        ID of the user
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)


@pulumi.input_type
class _ProjectGrantMemberState:
    def __init__(__self__, *,
                 grant_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectGrantMember resources.
        :param pulumi.Input[str] grant_id: ID of the grant
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles granted
        :param pulumi.Input[str] user_id: ID of the user
        """
        if grant_id is not None:
            pulumi.set(__self__, "grant_id", grant_id)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="grantId")
    def grant_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the grant
        """
        return pulumi.get(self, "grant_id")

    @grant_id.setter
    def grant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_id", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of roles granted
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the user
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class ProjectGrantMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grant_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing the membership of a user on an granted project, defined with the given role.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.ProjectGrantMember("default",
            org_id=data["zitadel_org"]["default"]["id"],
            project_id=data["zitadel_project"]["default"]["id"],
            user_id=data["zitadel_human_user"]["default"]["id"],
            grant_id="123456789012345678",
            roles=["PROJECT_GRANT_OWNER"])
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<project_id:grant_id:user_id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/projectGrantMember:ProjectGrantMember imported '123456789012345678:123456789012345678:123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grant_id: ID of the grant
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles granted
        :param pulumi.Input[str] user_id: ID of the user
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectGrantMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the membership of a user on an granted project, defined with the given role.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.ProjectGrantMember("default",
            org_id=data["zitadel_org"]["default"]["id"],
            project_id=data["zitadel_project"]["default"]["id"],
            user_id=data["zitadel_human_user"]["default"]["id"],
            grant_id="123456789012345678",
            roles=["PROJECT_GRANT_OWNER"])
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<project_id:grant_id:user_id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/projectGrantMember:ProjectGrantMember imported '123456789012345678:123456789012345678:123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param ProjectGrantMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectGrantMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grant_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectGrantMemberArgs.__new__(ProjectGrantMemberArgs)

            if grant_id is None and not opts.urn:
                raise TypeError("Missing required property 'grant_id'")
            __props__.__dict__["grant_id"] = grant_id
            __props__.__dict__["org_id"] = org_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if roles is None and not opts.urn:
                raise TypeError("Missing required property 'roles'")
            __props__.__dict__["roles"] = roles
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(ProjectGrantMember, __self__).__init__(
            'zitadel:index/projectGrantMember:ProjectGrantMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            grant_id: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'ProjectGrantMember':
        """
        Get an existing ProjectGrantMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grant_id: ID of the grant
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles granted
        :param pulumi.Input[str] user_id: ID of the user
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectGrantMemberState.__new__(_ProjectGrantMemberState)

        __props__.__dict__["grant_id"] = grant_id
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["roles"] = roles
        __props__.__dict__["user_id"] = user_id
        return ProjectGrantMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="grantId")
    def grant_id(self) -> pulumi.Output[str]:
        """
        ID of the grant
        """
        return pulumi.get(self, "grant_id")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence[str]]:
        """
        List of roles granted
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        ID of the user
        """
        return pulumi.get(self, "user_id")

