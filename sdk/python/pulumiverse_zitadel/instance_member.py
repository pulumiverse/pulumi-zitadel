# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InstanceMemberArgs', 'InstanceMember']

@pulumi.input_type
class InstanceMemberArgs:
    def __init__(__self__, *,
                 roles: pulumi.Input[Sequence[pulumi.Input[str]]],
                 user_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a InstanceMember resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
        :param pulumi.Input[str] user_id: ID of the user
        """
        InstanceMemberArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            roles=roles,
            user_id=user_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             roles: pulumi.Input[Sequence[pulumi.Input[str]]],
             user_id: pulumi.Input[str],
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'userId' in kwargs:
            user_id = kwargs['userId']

        _setter("roles", roles)
        _setter("user_id", user_id)

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
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


@pulumi.input_type
class _InstanceMemberState:
    def __init__(__self__, *,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceMember resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
        :param pulumi.Input[str] user_id: ID of the user
        """
        _InstanceMemberState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            roles=roles,
            user_id=user_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             user_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'userId' in kwargs:
            user_id = kwargs['userId']

        if roles is not None:
            _setter("roles", roles)
        if user_id is not None:
            _setter("user_id", user_id)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
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


class InstanceMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing the membership of a user on an instance, defined with the given role.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.InstanceMember("default",
            user_id=default_zitadel_human_user["id"],
            roles=["IAM_OWNER"])
        ```

        ## Import

        bash The resource can be imported using the ID format `<user_id>`, e.g.

        ```sh
         $ pulumi import zitadel:index/instanceMember:InstanceMember imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
        :param pulumi.Input[str] user_id: ID of the user
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the membership of a user on an instance, defined with the given role.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.InstanceMember("default",
            user_id=default_zitadel_human_user["id"],
            roles=["IAM_OWNER"])
        ```

        ## Import

        bash The resource can be imported using the ID format `<user_id>`, e.g.

        ```sh
         $ pulumi import zitadel:index/instanceMember:InstanceMember imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param InstanceMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            InstanceMemberArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceMemberArgs.__new__(InstanceMemberArgs)

            if roles is None and not opts.urn:
                raise TypeError("Missing required property 'roles'")
            __props__.__dict__["roles"] = roles
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(InstanceMember, __self__).__init__(
            'zitadel:index/instanceMember:InstanceMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'InstanceMember':
        """
        Get an existing InstanceMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
        :param pulumi.Input[str] user_id: ID of the user
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceMemberState.__new__(_InstanceMemberState)

        __props__.__dict__["roles"] = roles
        __props__.__dict__["user_id"] = user_id
        return InstanceMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence[str]]:
        """
        List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        ID of the user
        """
        return pulumi.get(self, "user_id")

