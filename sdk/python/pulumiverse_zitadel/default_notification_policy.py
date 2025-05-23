# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DefaultNotificationPolicyArgs', 'DefaultNotificationPolicy']

@pulumi.input_type
class DefaultNotificationPolicyArgs:
    def __init__(__self__, *,
                 password_change: pulumi.Input[bool]):
        """
        The set of arguments for constructing a DefaultNotificationPolicy resource.
        :param pulumi.Input[bool] password_change: Send notification if a user changes his password
        """
        DefaultNotificationPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            password_change=password_change,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             password_change: pulumi.Input[bool],
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'passwordChange' in kwargs:
            password_change = kwargs['passwordChange']

        _setter("password_change", password_change)

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


@pulumi.input_type
class _DefaultNotificationPolicyState:
    def __init__(__self__, *,
                 password_change: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering DefaultNotificationPolicy resources.
        :param pulumi.Input[bool] password_change: Send notification if a user changes his password
        """
        _DefaultNotificationPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            password_change=password_change,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             password_change: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'passwordChange' in kwargs:
            password_change = kwargs['passwordChange']

        if password_change is not None:
            _setter("password_change", password_change)

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


class DefaultNotificationPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 password_change: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Resource representing the default notification policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.DefaultNotificationPolicy("default", password_change=False)
        ```

        ## Import

        bash The resource can be imported using the ID format `<>`, e.g.

        ```sh
         $ pulumi import zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy imported ''
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] password_change: Send notification if a user changes his password
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DefaultNotificationPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the default notification policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.DefaultNotificationPolicy("default", password_change=False)
        ```

        ## Import

        bash The resource can be imported using the ID format `<>`, e.g.

        ```sh
         $ pulumi import zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy imported ''
        ```

        :param str resource_name: The name of the resource.
        :param DefaultNotificationPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultNotificationPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DefaultNotificationPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 password_change: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DefaultNotificationPolicyArgs.__new__(DefaultNotificationPolicyArgs)

            if password_change is None and not opts.urn:
                raise TypeError("Missing required property 'password_change'")
            __props__.__dict__["password_change"] = password_change
        super(DefaultNotificationPolicy, __self__).__init__(
            'zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            password_change: Optional[pulumi.Input[bool]] = None) -> 'DefaultNotificationPolicy':
        """
        Get an existing DefaultNotificationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] password_change: Send notification if a user changes his password
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultNotificationPolicyState.__new__(_DefaultNotificationPolicyState)

        __props__.__dict__["password_change"] = password_change
        return DefaultNotificationPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="passwordChange")
    def password_change(self) -> pulumi.Output[bool]:
        """
        Send notification if a user changes his password
        """
        return pulumi.get(self, "password_change")

