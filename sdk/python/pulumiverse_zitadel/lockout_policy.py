# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LockoutPolicyArgs', 'LockoutPolicy']

@pulumi.input_type
class LockoutPolicyArgs:
    def __init__(__self__, *,
                 max_password_attempts: pulumi.Input[int],
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LockoutPolicy resource.
        :param pulumi.Input[int] max_password_attempts: Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        :param pulumi.Input[str] org_id: ID of the organization
        """
        LockoutPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            max_password_attempts=max_password_attempts,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             max_password_attempts: pulumi.Input[int],
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'maxPasswordAttempts' in kwargs:
            max_password_attempts = kwargs['maxPasswordAttempts']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        _setter("max_password_attempts", max_password_attempts)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="maxPasswordAttempts")
    def max_password_attempts(self) -> pulumi.Input[int]:
        """
        Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        """
        return pulumi.get(self, "max_password_attempts")

    @max_password_attempts.setter
    def max_password_attempts(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_password_attempts", value)

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
class _LockoutPolicyState:
    def __init__(__self__, *,
                 max_password_attempts: Optional[pulumi.Input[int]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LockoutPolicy resources.
        :param pulumi.Input[int] max_password_attempts: Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        :param pulumi.Input[str] org_id: ID of the organization
        """
        _LockoutPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            max_password_attempts=max_password_attempts,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             max_password_attempts: Optional[pulumi.Input[int]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'maxPasswordAttempts' in kwargs:
            max_password_attempts = kwargs['maxPasswordAttempts']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        if max_password_attempts is not None:
            _setter("max_password_attempts", max_password_attempts)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="maxPasswordAttempts")
    def max_password_attempts(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        """
        return pulumi.get(self, "max_password_attempts")

    @max_password_attempts.setter
    def max_password_attempts(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_password_attempts", value)

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


class LockoutPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_password_attempts: Optional[pulumi.Input[int]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing the custom lockout policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import scoretechnologies_zitadel as zitadel

        default = zitadel.LockoutPolicy("default",
            org_id=data["zitadel_org"]["default"]["id"],
            max_password_attempts=5)
        ```

        ## Import

        bash The resource can be imported using the ID format `<[org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/lockoutPolicy:LockoutPolicy imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] max_password_attempts: Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        :param pulumi.Input[str] org_id: ID of the organization
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LockoutPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the custom lockout policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import scoretechnologies_zitadel as zitadel

        default = zitadel.LockoutPolicy("default",
            org_id=data["zitadel_org"]["default"]["id"],
            max_password_attempts=5)
        ```

        ## Import

        bash The resource can be imported using the ID format `<[org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/lockoutPolicy:LockoutPolicy imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param LockoutPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LockoutPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            LockoutPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_password_attempts: Optional[pulumi.Input[int]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LockoutPolicyArgs.__new__(LockoutPolicyArgs)

            if max_password_attempts is None and not opts.urn:
                raise TypeError("Missing required property 'max_password_attempts'")
            __props__.__dict__["max_password_attempts"] = max_password_attempts
            __props__.__dict__["org_id"] = org_id
        super(LockoutPolicy, __self__).__init__(
            'zitadel:index/lockoutPolicy:LockoutPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            max_password_attempts: Optional[pulumi.Input[int]] = None,
            org_id: Optional[pulumi.Input[str]] = None) -> 'LockoutPolicy':
        """
        Get an existing LockoutPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] max_password_attempts: Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        :param pulumi.Input[str] org_id: ID of the organization
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LockoutPolicyState.__new__(_LockoutPolicyState)

        __props__.__dict__["max_password_attempts"] = max_password_attempts
        __props__.__dict__["org_id"] = org_id
        return LockoutPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="maxPasswordAttempts")
    def max_password_attempts(self) -> pulumi.Output[int]:
        """
        Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        """
        return pulumi.get(self, "max_password_attempts")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

