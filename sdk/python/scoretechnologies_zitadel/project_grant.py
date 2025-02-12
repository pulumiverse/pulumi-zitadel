# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectGrantArgs', 'ProjectGrant']

@pulumi.input_type
class ProjectGrantArgs:
    def __init__(__self__, *,
                 granted_org_id: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None,
                 role_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ProjectGrant resource.
        :param pulumi.Input[str] granted_org_id: ID of the organization granted the project
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_keys: List of roles granted
        """
        ProjectGrantArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            granted_org_id=granted_org_id,
            project_id=project_id,
            org_id=org_id,
            role_keys=role_keys,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             granted_org_id: pulumi.Input[str],
             project_id: pulumi.Input[str],
             org_id: Optional[pulumi.Input[str]] = None,
             role_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'grantedOrgId' in kwargs:
            granted_org_id = kwargs['grantedOrgId']
        if 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'roleKeys' in kwargs:
            role_keys = kwargs['roleKeys']

        _setter("granted_org_id", granted_org_id)
        _setter("project_id", project_id)
        if org_id is not None:
            _setter("org_id", org_id)
        if role_keys is not None:
            _setter("role_keys", role_keys)

    @property
    @pulumi.getter(name="grantedOrgId")
    def granted_org_id(self) -> pulumi.Input[str]:
        """
        ID of the organization granted the project
        """
        return pulumi.get(self, "granted_org_id")

    @granted_org_id.setter
    def granted_org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "granted_org_id", value)

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
    @pulumi.getter(name="roleKeys")
    def role_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of roles granted
        """
        return pulumi.get(self, "role_keys")

    @role_keys.setter
    def role_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_keys", value)


@pulumi.input_type
class _ProjectGrantState:
    def __init__(__self__, *,
                 granted_org_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 role_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ProjectGrant resources.
        :param pulumi.Input[str] granted_org_id: ID of the organization granted the project
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_keys: List of roles granted
        """
        _ProjectGrantState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            granted_org_id=granted_org_id,
            org_id=org_id,
            project_id=project_id,
            role_keys=role_keys,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             granted_org_id: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             role_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'grantedOrgId' in kwargs:
            granted_org_id = kwargs['grantedOrgId']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if 'roleKeys' in kwargs:
            role_keys = kwargs['roleKeys']

        if granted_org_id is not None:
            _setter("granted_org_id", granted_org_id)
        if org_id is not None:
            _setter("org_id", org_id)
        if project_id is not None:
            _setter("project_id", project_id)
        if role_keys is not None:
            _setter("role_keys", role_keys)

    @property
    @pulumi.getter(name="grantedOrgId")
    def granted_org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization granted the project
        """
        return pulumi.get(self, "granted_org_id")

    @granted_org_id.setter
    def granted_org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "granted_org_id", value)

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
    @pulumi.getter(name="roleKeys")
    def role_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of roles granted
        """
        return pulumi.get(self, "role_keys")

    @role_keys.setter
    def role_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_keys", value)


class ProjectGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 granted_org_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 role_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource representing the grant of a project to a different organization, also containing the available roles which can be given to the members of the projectgrant.

        ## Example Usage

        ```python
        import pulumi
        import scoretechnologies_zitadel as zitadel

        default = zitadel.ProjectGrant("default",
            org_id=data["zitadel_org"]["default"]["id"],
            project_id=data["zitadel_project"]["default"]["id"],
            granted_org_id=data["zitadel_org"]["granted_org"]["id"],
            role_keys=["super-user"])
        ```

        ## Import

        bash The resource can be imported using the ID format `<id:project_id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/projectGrant:ProjectGrant imported '123456789012345678:123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] granted_org_id: ID of the organization granted the project
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_keys: List of roles granted
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the grant of a project to a different organization, also containing the available roles which can be given to the members of the projectgrant.

        ## Example Usage

        ```python
        import pulumi
        import scoretechnologies_zitadel as zitadel

        default = zitadel.ProjectGrant("default",
            org_id=data["zitadel_org"]["default"]["id"],
            project_id=data["zitadel_project"]["default"]["id"],
            granted_org_id=data["zitadel_org"]["granted_org"]["id"],
            role_keys=["super-user"])
        ```

        ## Import

        bash The resource can be imported using the ID format `<id:project_id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/projectGrant:ProjectGrant imported '123456789012345678:123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param ProjectGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProjectGrantArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 granted_org_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 role_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectGrantArgs.__new__(ProjectGrantArgs)

            if granted_org_id is None and not opts.urn:
                raise TypeError("Missing required property 'granted_org_id'")
            __props__.__dict__["granted_org_id"] = granted_org_id
            __props__.__dict__["org_id"] = org_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["role_keys"] = role_keys
        super(ProjectGrant, __self__).__init__(
            'zitadel:index/projectGrant:ProjectGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            granted_org_id: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            role_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ProjectGrant':
        """
        Get an existing ProjectGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] granted_org_id: ID of the organization granted the project
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_keys: List of roles granted
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectGrantState.__new__(_ProjectGrantState)

        __props__.__dict__["granted_org_id"] = granted_org_id
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["role_keys"] = role_keys
        return ProjectGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="grantedOrgId")
    def granted_org_id(self) -> pulumi.Output[str]:
        """
        ID of the organization granted the project
        """
        return pulumi.get(self, "granted_org_id")

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
    @pulumi.getter(name="roleKeys")
    def role_keys(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of roles granted
        """
        return pulumi.get(self, "role_keys")

