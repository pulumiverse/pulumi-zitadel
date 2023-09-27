# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SmtpConfigArgs', 'SmtpConfig']

@pulumi.input_type
class SmtpConfigArgs:
    def __init__(__self__, *,
                 host: pulumi.Input[str],
                 sender_address: pulumi.Input[str],
                 sender_name: pulumi.Input[str],
                 password: Optional[pulumi.Input[str]] = None,
                 tls: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SmtpConfig resource.
        :param pulumi.Input[str] host: Host and port address to your SMTP server.
        :param pulumi.Input[str] sender_address: Address used to send emails.
        :param pulumi.Input[str] sender_name: Sender name used to send emails.
        :param pulumi.Input[str] password: Password used to communicate with your SMTP server.
        :param pulumi.Input[bool] tls: TLS used to communicate with your SMTP server.
        :param pulumi.Input[str] user: User used to communicate with your SMTP server.
        """
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "sender_address", sender_address)
        pulumi.set(__self__, "sender_name", sender_name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if tls is not None:
            pulumi.set(__self__, "tls", tls)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        Host and port address to your SMTP server.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="senderAddress")
    def sender_address(self) -> pulumi.Input[str]:
        """
        Address used to send emails.
        """
        return pulumi.get(self, "sender_address")

    @sender_address.setter
    def sender_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "sender_address", value)

    @property
    @pulumi.getter(name="senderName")
    def sender_name(self) -> pulumi.Input[str]:
        """
        Sender name used to send emails.
        """
        return pulumi.get(self, "sender_name")

    @sender_name.setter
    def sender_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "sender_name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password used to communicate with your SMTP server.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def tls(self) -> Optional[pulumi.Input[bool]]:
        """
        TLS used to communicate with your SMTP server.
        """
        return pulumi.get(self, "tls")

    @tls.setter
    def tls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        User used to communicate with your SMTP server.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


@pulumi.input_type
class _SmtpConfigState:
    def __init__(__self__, *,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 sender_address: Optional[pulumi.Input[str]] = None,
                 sender_name: Optional[pulumi.Input[str]] = None,
                 tls: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SmtpConfig resources.
        :param pulumi.Input[str] host: Host and port address to your SMTP server.
        :param pulumi.Input[str] password: Password used to communicate with your SMTP server.
        :param pulumi.Input[str] sender_address: Address used to send emails.
        :param pulumi.Input[str] sender_name: Sender name used to send emails.
        :param pulumi.Input[bool] tls: TLS used to communicate with your SMTP server.
        :param pulumi.Input[str] user: User used to communicate with your SMTP server.
        """
        if host is not None:
            pulumi.set(__self__, "host", host)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if sender_address is not None:
            pulumi.set(__self__, "sender_address", sender_address)
        if sender_name is not None:
            pulumi.set(__self__, "sender_name", sender_name)
        if tls is not None:
            pulumi.set(__self__, "tls", tls)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Host and port address to your SMTP server.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password used to communicate with your SMTP server.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="senderAddress")
    def sender_address(self) -> Optional[pulumi.Input[str]]:
        """
        Address used to send emails.
        """
        return pulumi.get(self, "sender_address")

    @sender_address.setter
    def sender_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sender_address", value)

    @property
    @pulumi.getter(name="senderName")
    def sender_name(self) -> Optional[pulumi.Input[str]]:
        """
        Sender name used to send emails.
        """
        return pulumi.get(self, "sender_name")

    @sender_name.setter
    def sender_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sender_name", value)

    @property
    @pulumi.getter
    def tls(self) -> Optional[pulumi.Input[bool]]:
        """
        TLS used to communicate with your SMTP server.
        """
        return pulumi.get(self, "tls")

    @tls.setter
    def tls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        User used to communicate with your SMTP server.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class SmtpConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 sender_address: Optional[pulumi.Input[str]] = None,
                 sender_name: Optional[pulumi.Input[str]] = None,
                 tls: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing the SMTP configuration of an instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.SmtpConfig("default",
            host="localhost:25",
            password="secret_password",
            sender_address="sender@example.com",
            sender_name="no-reply",
            tls=True,
            user="user")
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<[password]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/smtpConfig:SmtpConfig imported 'p4ssw0rd'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: Host and port address to your SMTP server.
        :param pulumi.Input[str] password: Password used to communicate with your SMTP server.
        :param pulumi.Input[str] sender_address: Address used to send emails.
        :param pulumi.Input[str] sender_name: Sender name used to send emails.
        :param pulumi.Input[bool] tls: TLS used to communicate with your SMTP server.
        :param pulumi.Input[str] user: User used to communicate with your SMTP server.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SmtpConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the SMTP configuration of an instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.SmtpConfig("default",
            host="localhost:25",
            password="secret_password",
            sender_address="sender@example.com",
            sender_name="no-reply",
            tls=True,
            user="user")
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<[password]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/smtpConfig:SmtpConfig imported 'p4ssw0rd'
        ```

        :param str resource_name: The name of the resource.
        :param SmtpConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SmtpConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 sender_address: Optional[pulumi.Input[str]] = None,
                 sender_name: Optional[pulumi.Input[str]] = None,
                 tls: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SmtpConfigArgs.__new__(SmtpConfigArgs)

            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            __props__.__dict__["password"] = password
            if sender_address is None and not opts.urn:
                raise TypeError("Missing required property 'sender_address'")
            __props__.__dict__["sender_address"] = sender_address
            if sender_name is None and not opts.urn:
                raise TypeError("Missing required property 'sender_name'")
            __props__.__dict__["sender_name"] = sender_name
            __props__.__dict__["tls"] = tls
            __props__.__dict__["user"] = user
        super(SmtpConfig, __self__).__init__(
            'zitadel:index/smtpConfig:SmtpConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            host: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            sender_address: Optional[pulumi.Input[str]] = None,
            sender_name: Optional[pulumi.Input[str]] = None,
            tls: Optional[pulumi.Input[bool]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'SmtpConfig':
        """
        Get an existing SmtpConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: Host and port address to your SMTP server.
        :param pulumi.Input[str] password: Password used to communicate with your SMTP server.
        :param pulumi.Input[str] sender_address: Address used to send emails.
        :param pulumi.Input[str] sender_name: Sender name used to send emails.
        :param pulumi.Input[bool] tls: TLS used to communicate with your SMTP server.
        :param pulumi.Input[str] user: User used to communicate with your SMTP server.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SmtpConfigState.__new__(_SmtpConfigState)

        __props__.__dict__["host"] = host
        __props__.__dict__["password"] = password
        __props__.__dict__["sender_address"] = sender_address
        __props__.__dict__["sender_name"] = sender_name
        __props__.__dict__["tls"] = tls
        __props__.__dict__["user"] = user
        return SmtpConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Host and port address to your SMTP server.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password used to communicate with your SMTP server.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="senderAddress")
    def sender_address(self) -> pulumi.Output[str]:
        """
        Address used to send emails.
        """
        return pulumi.get(self, "sender_address")

    @property
    @pulumi.getter(name="senderName")
    def sender_name(self) -> pulumi.Output[str]:
        """
        Sender name used to send emails.
        """
        return pulumi.get(self, "sender_name")

    @property
    @pulumi.getter
    def tls(self) -> pulumi.Output[Optional[bool]]:
        """
        TLS used to communicate with your SMTP server.
        """
        return pulumi.get(self, "tls")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[Optional[str]]:
        """
        User used to communicate with your SMTP server.
        """
        return pulumi.get(self, "user")

