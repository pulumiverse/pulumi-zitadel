# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MachineKeyArgs', 'MachineKey']

@pulumi.input_type
class MachineKeyArgs:
    def __init__(__self__, *,
                 key_type: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MachineKey resource.
        :param pulumi.Input[str] key_type: Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        :param pulumi.Input[str] user_id: ID of the user
        :param pulumi.Input[str] expiration_date: Expiration date of the machine key in the RFC3339 format
        :param pulumi.Input[str] org_id: ID of the organization
        """
        MachineKeyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            key_type=key_type,
            user_id=user_id,
            expiration_date=expiration_date,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             key_type: pulumi.Input[str],
             user_id: pulumi.Input[str],
             expiration_date: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'keyType' in kwargs:
            key_type = kwargs['keyType']
        if 'userId' in kwargs:
            user_id = kwargs['userId']
        if 'expirationDate' in kwargs:
            expiration_date = kwargs['expirationDate']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        _setter("key_type", key_type)
        _setter("user_id", user_id)
        if expiration_date is not None:
            _setter("expiration_date", expiration_date)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Input[str]:
        """
        Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_type", value)

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
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date of the machine key in the RFC3339 format
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

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
class _MachineKeyState:
    def __init__(__self__, *,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_details: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MachineKey resources.
        :param pulumi.Input[str] expiration_date: Expiration date of the machine key in the RFC3339 format
        :param pulumi.Input[str] key_details: Value of the machine key
        :param pulumi.Input[str] key_type: Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] user_id: ID of the user
        """
        _MachineKeyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            expiration_date=expiration_date,
            key_details=key_details,
            key_type=key_type,
            org_id=org_id,
            user_id=user_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             expiration_date: Optional[pulumi.Input[str]] = None,
             key_details: Optional[pulumi.Input[str]] = None,
             key_type: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             user_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'expirationDate' in kwargs:
            expiration_date = kwargs['expirationDate']
        if 'keyDetails' in kwargs:
            key_details = kwargs['keyDetails']
        if 'keyType' in kwargs:
            key_type = kwargs['keyType']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'userId' in kwargs:
            user_id = kwargs['userId']

        if expiration_date is not None:
            _setter("expiration_date", expiration_date)
        if key_details is not None:
            _setter("key_details", key_details)
        if key_type is not None:
            _setter("key_type", key_type)
        if org_id is not None:
            _setter("org_id", org_id)
        if user_id is not None:
            _setter("user_id", user_id)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date of the machine key in the RFC3339 format
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="keyDetails")
    def key_details(self) -> Optional[pulumi.Input[str]]:
        """
        Value of the machine key
        """
        return pulumi.get(self, "key_details")

    @key_details.setter
    def key_details(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_details", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_type", value)

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
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the user
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class MachineKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing a machine key

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.MachineKey("default",
            org_id=default_zitadel_org["id"],
            user_id=default_zitadel_machine_user["id"],
            key_type="KEY_TYPE_JSON",
            expiration_date="2519-04-01T08:45:00Z")
        ```

        ## Import

        bash The resource can be imported using the ID format `<id:user_id[:org_id][:key_details]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/machineKey:MachineKey imported '123456789012345678:123456789012345678:123456789012345678:{"type":"serviceaccount","keyId":"123456789012345678","key":"-----BEGIN RSA PRIVATE KEY-----\\nMIIEpQ...-----END RSA PRIVATE KEY-----\\n","userId":"123456789012345678"}'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_date: Expiration date of the machine key in the RFC3339 format
        :param pulumi.Input[str] key_type: Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] user_id: ID of the user
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MachineKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a machine key

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.MachineKey("default",
            org_id=default_zitadel_org["id"],
            user_id=default_zitadel_machine_user["id"],
            key_type="KEY_TYPE_JSON",
            expiration_date="2519-04-01T08:45:00Z")
        ```

        ## Import

        bash The resource can be imported using the ID format `<id:user_id[:org_id][:key_details]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/machineKey:MachineKey imported '123456789012345678:123456789012345678:123456789012345678:{"type":"serviceaccount","keyId":"123456789012345678","key":"-----BEGIN RSA PRIVATE KEY-----\\nMIIEpQ...-----END RSA PRIVATE KEY-----\\n","userId":"123456789012345678"}'
        ```

        :param str resource_name: The name of the resource.
        :param MachineKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MachineKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            MachineKeyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MachineKeyArgs.__new__(MachineKeyArgs)

            __props__.__dict__["expiration_date"] = expiration_date
            if key_type is None and not opts.urn:
                raise TypeError("Missing required property 'key_type'")
            __props__.__dict__["key_type"] = key_type
            __props__.__dict__["org_id"] = org_id
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["key_details"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["keyDetails"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MachineKey, __self__).__init__(
            'zitadel:index/machineKey:MachineKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            key_details: Optional[pulumi.Input[str]] = None,
            key_type: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'MachineKey':
        """
        Get an existing MachineKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_date: Expiration date of the machine key in the RFC3339 format
        :param pulumi.Input[str] key_details: Value of the machine key
        :param pulumi.Input[str] key_type: Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] user_id: ID of the user
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MachineKeyState.__new__(_MachineKeyState)

        __props__.__dict__["expiration_date"] = expiration_date
        __props__.__dict__["key_details"] = key_details
        __props__.__dict__["key_type"] = key_type
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["user_id"] = user_id
        return MachineKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Output[str]:
        """
        Expiration date of the machine key in the RFC3339 format
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="keyDetails")
    def key_details(self) -> pulumi.Output[str]:
        """
        Value of the machine key
        """
        return pulumi.get(self, "key_details")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[str]:
        """
        Type of the machine key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        ID of the user
        """
        return pulumi.get(self, "user_id")

