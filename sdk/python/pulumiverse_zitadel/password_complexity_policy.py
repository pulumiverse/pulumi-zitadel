# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PasswordComplexityPolicyArgs', 'PasswordComplexityPolicy']

@pulumi.input_type
class PasswordComplexityPolicyArgs:
    def __init__(__self__, *,
                 has_lowercase: pulumi.Input[bool],
                 has_number: pulumi.Input[bool],
                 has_symbol: pulumi.Input[bool],
                 has_uppercase: pulumi.Input[bool],
                 min_length: pulumi.Input[int],
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PasswordComplexityPolicy resource.
        :param pulumi.Input[bool] has_lowercase: defines if the password MUST contain a lower case letter
        :param pulumi.Input[bool] has_number: defines if the password MUST contain a number
        :param pulumi.Input[bool] has_symbol: defines if the password MUST contain a symbol. E.g. "$"
        :param pulumi.Input[bool] has_uppercase: defines if the password MUST contain an upper case letter
        :param pulumi.Input[int] min_length: Minimal length for the password
        :param pulumi.Input[str] org_id: ID of the organization
        """
        PasswordComplexityPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            has_lowercase=has_lowercase,
            has_number=has_number,
            has_symbol=has_symbol,
            has_uppercase=has_uppercase,
            min_length=min_length,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             has_lowercase: pulumi.Input[bool],
             has_number: pulumi.Input[bool],
             has_symbol: pulumi.Input[bool],
             has_uppercase: pulumi.Input[bool],
             min_length: pulumi.Input[int],
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'hasLowercase' in kwargs:
            has_lowercase = kwargs['hasLowercase']
        if 'hasNumber' in kwargs:
            has_number = kwargs['hasNumber']
        if 'hasSymbol' in kwargs:
            has_symbol = kwargs['hasSymbol']
        if 'hasUppercase' in kwargs:
            has_uppercase = kwargs['hasUppercase']
        if 'minLength' in kwargs:
            min_length = kwargs['minLength']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        _setter("has_lowercase", has_lowercase)
        _setter("has_number", has_number)
        _setter("has_symbol", has_symbol)
        _setter("has_uppercase", has_uppercase)
        _setter("min_length", min_length)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="hasLowercase")
    def has_lowercase(self) -> pulumi.Input[bool]:
        """
        defines if the password MUST contain a lower case letter
        """
        return pulumi.get(self, "has_lowercase")

    @has_lowercase.setter
    def has_lowercase(self, value: pulumi.Input[bool]):
        pulumi.set(self, "has_lowercase", value)

    @property
    @pulumi.getter(name="hasNumber")
    def has_number(self) -> pulumi.Input[bool]:
        """
        defines if the password MUST contain a number
        """
        return pulumi.get(self, "has_number")

    @has_number.setter
    def has_number(self, value: pulumi.Input[bool]):
        pulumi.set(self, "has_number", value)

    @property
    @pulumi.getter(name="hasSymbol")
    def has_symbol(self) -> pulumi.Input[bool]:
        """
        defines if the password MUST contain a symbol. E.g. "$"
        """
        return pulumi.get(self, "has_symbol")

    @has_symbol.setter
    def has_symbol(self, value: pulumi.Input[bool]):
        pulumi.set(self, "has_symbol", value)

    @property
    @pulumi.getter(name="hasUppercase")
    def has_uppercase(self) -> pulumi.Input[bool]:
        """
        defines if the password MUST contain an upper case letter
        """
        return pulumi.get(self, "has_uppercase")

    @has_uppercase.setter
    def has_uppercase(self, value: pulumi.Input[bool]):
        pulumi.set(self, "has_uppercase", value)

    @property
    @pulumi.getter(name="minLength")
    def min_length(self) -> pulumi.Input[int]:
        """
        Minimal length for the password
        """
        return pulumi.get(self, "min_length")

    @min_length.setter
    def min_length(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_length", value)

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
class _PasswordComplexityPolicyState:
    def __init__(__self__, *,
                 has_lowercase: Optional[pulumi.Input[bool]] = None,
                 has_number: Optional[pulumi.Input[bool]] = None,
                 has_symbol: Optional[pulumi.Input[bool]] = None,
                 has_uppercase: Optional[pulumi.Input[bool]] = None,
                 min_length: Optional[pulumi.Input[int]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PasswordComplexityPolicy resources.
        :param pulumi.Input[bool] has_lowercase: defines if the password MUST contain a lower case letter
        :param pulumi.Input[bool] has_number: defines if the password MUST contain a number
        :param pulumi.Input[bool] has_symbol: defines if the password MUST contain a symbol. E.g. "$"
        :param pulumi.Input[bool] has_uppercase: defines if the password MUST contain an upper case letter
        :param pulumi.Input[int] min_length: Minimal length for the password
        :param pulumi.Input[str] org_id: ID of the organization
        """
        _PasswordComplexityPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            has_lowercase=has_lowercase,
            has_number=has_number,
            has_symbol=has_symbol,
            has_uppercase=has_uppercase,
            min_length=min_length,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             has_lowercase: Optional[pulumi.Input[bool]] = None,
             has_number: Optional[pulumi.Input[bool]] = None,
             has_symbol: Optional[pulumi.Input[bool]] = None,
             has_uppercase: Optional[pulumi.Input[bool]] = None,
             min_length: Optional[pulumi.Input[int]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'hasLowercase' in kwargs:
            has_lowercase = kwargs['hasLowercase']
        if 'hasNumber' in kwargs:
            has_number = kwargs['hasNumber']
        if 'hasSymbol' in kwargs:
            has_symbol = kwargs['hasSymbol']
        if 'hasUppercase' in kwargs:
            has_uppercase = kwargs['hasUppercase']
        if 'minLength' in kwargs:
            min_length = kwargs['minLength']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        if has_lowercase is not None:
            _setter("has_lowercase", has_lowercase)
        if has_number is not None:
            _setter("has_number", has_number)
        if has_symbol is not None:
            _setter("has_symbol", has_symbol)
        if has_uppercase is not None:
            _setter("has_uppercase", has_uppercase)
        if min_length is not None:
            _setter("min_length", min_length)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="hasLowercase")
    def has_lowercase(self) -> Optional[pulumi.Input[bool]]:
        """
        defines if the password MUST contain a lower case letter
        """
        return pulumi.get(self, "has_lowercase")

    @has_lowercase.setter
    def has_lowercase(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "has_lowercase", value)

    @property
    @pulumi.getter(name="hasNumber")
    def has_number(self) -> Optional[pulumi.Input[bool]]:
        """
        defines if the password MUST contain a number
        """
        return pulumi.get(self, "has_number")

    @has_number.setter
    def has_number(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "has_number", value)

    @property
    @pulumi.getter(name="hasSymbol")
    def has_symbol(self) -> Optional[pulumi.Input[bool]]:
        """
        defines if the password MUST contain a symbol. E.g. "$"
        """
        return pulumi.get(self, "has_symbol")

    @has_symbol.setter
    def has_symbol(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "has_symbol", value)

    @property
    @pulumi.getter(name="hasUppercase")
    def has_uppercase(self) -> Optional[pulumi.Input[bool]]:
        """
        defines if the password MUST contain an upper case letter
        """
        return pulumi.get(self, "has_uppercase")

    @has_uppercase.setter
    def has_uppercase(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "has_uppercase", value)

    @property
    @pulumi.getter(name="minLength")
    def min_length(self) -> Optional[pulumi.Input[int]]:
        """
        Minimal length for the password
        """
        return pulumi.get(self, "min_length")

    @min_length.setter
    def min_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_length", value)

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


class PasswordComplexityPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 has_lowercase: Optional[pulumi.Input[bool]] = None,
                 has_number: Optional[pulumi.Input[bool]] = None,
                 has_symbol: Optional[pulumi.Input[bool]] = None,
                 has_uppercase: Optional[pulumi.Input[bool]] = None,
                 min_length: Optional[pulumi.Input[int]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing the custom password complexity policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.PasswordComplexityPolicy("default",
            org_id=default_zitadel_org["id"],
            min_length=8,
            has_uppercase=True,
            has_lowercase=True,
            has_number=True,
            has_symbol=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] has_lowercase: defines if the password MUST contain a lower case letter
        :param pulumi.Input[bool] has_number: defines if the password MUST contain a number
        :param pulumi.Input[bool] has_symbol: defines if the password MUST contain a symbol. E.g. "$"
        :param pulumi.Input[bool] has_uppercase: defines if the password MUST contain an upper case letter
        :param pulumi.Input[int] min_length: Minimal length for the password
        :param pulumi.Input[str] org_id: ID of the organization
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PasswordComplexityPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the custom password complexity policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.PasswordComplexityPolicy("default",
            org_id=default_zitadel_org["id"],
            min_length=8,
            has_uppercase=True,
            has_lowercase=True,
            has_number=True,
            has_symbol=True)
        ```

        :param str resource_name: The name of the resource.
        :param PasswordComplexityPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PasswordComplexityPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            PasswordComplexityPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 has_lowercase: Optional[pulumi.Input[bool]] = None,
                 has_number: Optional[pulumi.Input[bool]] = None,
                 has_symbol: Optional[pulumi.Input[bool]] = None,
                 has_uppercase: Optional[pulumi.Input[bool]] = None,
                 min_length: Optional[pulumi.Input[int]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PasswordComplexityPolicyArgs.__new__(PasswordComplexityPolicyArgs)

            if has_lowercase is None and not opts.urn:
                raise TypeError("Missing required property 'has_lowercase'")
            __props__.__dict__["has_lowercase"] = has_lowercase
            if has_number is None and not opts.urn:
                raise TypeError("Missing required property 'has_number'")
            __props__.__dict__["has_number"] = has_number
            if has_symbol is None and not opts.urn:
                raise TypeError("Missing required property 'has_symbol'")
            __props__.__dict__["has_symbol"] = has_symbol
            if has_uppercase is None and not opts.urn:
                raise TypeError("Missing required property 'has_uppercase'")
            __props__.__dict__["has_uppercase"] = has_uppercase
            if min_length is None and not opts.urn:
                raise TypeError("Missing required property 'min_length'")
            __props__.__dict__["min_length"] = min_length
            __props__.__dict__["org_id"] = org_id
        super(PasswordComplexityPolicy, __self__).__init__(
            'zitadel:index/passwordComplexityPolicy:PasswordComplexityPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            has_lowercase: Optional[pulumi.Input[bool]] = None,
            has_number: Optional[pulumi.Input[bool]] = None,
            has_symbol: Optional[pulumi.Input[bool]] = None,
            has_uppercase: Optional[pulumi.Input[bool]] = None,
            min_length: Optional[pulumi.Input[int]] = None,
            org_id: Optional[pulumi.Input[str]] = None) -> 'PasswordComplexityPolicy':
        """
        Get an existing PasswordComplexityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] has_lowercase: defines if the password MUST contain a lower case letter
        :param pulumi.Input[bool] has_number: defines if the password MUST contain a number
        :param pulumi.Input[bool] has_symbol: defines if the password MUST contain a symbol. E.g. "$"
        :param pulumi.Input[bool] has_uppercase: defines if the password MUST contain an upper case letter
        :param pulumi.Input[int] min_length: Minimal length for the password
        :param pulumi.Input[str] org_id: ID of the organization
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PasswordComplexityPolicyState.__new__(_PasswordComplexityPolicyState)

        __props__.__dict__["has_lowercase"] = has_lowercase
        __props__.__dict__["has_number"] = has_number
        __props__.__dict__["has_symbol"] = has_symbol
        __props__.__dict__["has_uppercase"] = has_uppercase
        __props__.__dict__["min_length"] = min_length
        __props__.__dict__["org_id"] = org_id
        return PasswordComplexityPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="hasLowercase")
    def has_lowercase(self) -> pulumi.Output[bool]:
        """
        defines if the password MUST contain a lower case letter
        """
        return pulumi.get(self, "has_lowercase")

    @property
    @pulumi.getter(name="hasNumber")
    def has_number(self) -> pulumi.Output[bool]:
        """
        defines if the password MUST contain a number
        """
        return pulumi.get(self, "has_number")

    @property
    @pulumi.getter(name="hasSymbol")
    def has_symbol(self) -> pulumi.Output[bool]:
        """
        defines if the password MUST contain a symbol. E.g. "$"
        """
        return pulumi.get(self, "has_symbol")

    @property
    @pulumi.getter(name="hasUppercase")
    def has_uppercase(self) -> pulumi.Output[bool]:
        """
        defines if the password MUST contain an upper case letter
        """
        return pulumi.get(self, "has_uppercase")

    @property
    @pulumi.getter(name="minLength")
    def min_length(self) -> pulumi.Output[int]:
        """
        Minimal length for the password
        """
        return pulumi.get(self, "min_length")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

