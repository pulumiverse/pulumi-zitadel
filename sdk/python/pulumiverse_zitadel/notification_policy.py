# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NotificationPolicyArgs', 'NotificationPolicy']

@pulumi.input_type
class NotificationPolicyArgs:
    def __init__(__self__, *,
                 password_change: pulumi.Input[bool],
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NotificationPolicy resource.
        :param pulumi.Input[bool] password_change: Send notification if a user changes his password
        :param pulumi.Input[str] org_id: ID of the organization
        """
        NotificationPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            password_change=password_change,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             password_change: pulumi.Input[bool],
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'passwordChange' in kwargs:
            password_change = kwargs['passwordChange']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        _setter("password_change", password_change)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="passwordChange")
    def password_change(self) -> pulumi.Input[bool]:
        """
        Send notification if a user changes his password
        """
        return pulumi.get(self, "password_change")

    @password_change.setter
    def password_change(self, value: pulumi.Input[bool]):
        pulumi.set(self, "password_change", value)

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
class _NotificationPolicyState:
    def __init__(__self__, *,
                 org_id: Optional[pulumi.Input[str]] = None,
                 password_change: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering NotificationPolicy resources.
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[bool] password_change: Send notification if a user changes his password
        """
        _NotificationPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            org_id=org_id,
            password_change=password_change,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             org_id: Optional[pulumi.Input[str]] = None,
             password_change: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'passwordChange' in kwargs:
            password_change = kwargs['passwordChange']

        if org_id is not None:
            _setter("org_id", org_id)
        if password_change is not None:
            _setter("password_change", password_change)

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
    @pulumi.getter(name="passwordChange")
    def password_change(self) -> Optional[pulumi.Input[bool]]:
        """
        Send notification if a user changes his password
        """
        return pulumi.get(self, "password_change")

    @password_change.setter
    def password_change(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "password_change", value)


class NotificationPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 password_change: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Resource representing the custom notification policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.NotificationPolicy("default",
            org_id=default_zitadel_org["id"],
            password_change=False)
        ```

        ## Import

        bash The resource can be imported using the ID format `<[org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/notificationPolicy:NotificationPolicy imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[bool] password_change: Send notification if a user changes his password
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the custom notification policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.NotificationPolicy("default",
            org_id=default_zitadel_org["id"],
            password_change=False)
        ```

        ## Import

        bash The resource can be imported using the ID format `<[org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/notificationPolicy:NotificationPolicy imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param NotificationPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            NotificationPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 password_change: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NotificationPolicyArgs.__new__(NotificationPolicyArgs)

            __props__.__dict__["org_id"] = org_id
            if password_change is None and not opts.urn:
                raise TypeError("Missing required property 'password_change'")
            __props__.__dict__["password_change"] = password_change
        super(NotificationPolicy, __self__).__init__(
            'zitadel:index/notificationPolicy:NotificationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            password_change: Optional[pulumi.Input[bool]] = None) -> 'NotificationPolicy':
        """
        Get an existing NotificationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[bool] password_change: Send notification if a user changes his password
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NotificationPolicyState.__new__(_NotificationPolicyState)

        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["password_change"] = password_change
        return NotificationPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="passwordChange")
    def password_change(self) -> pulumi.Output[bool]:
        """
        Send notification if a user changes his password
        """
        return pulumi.get(self, "password_change")

