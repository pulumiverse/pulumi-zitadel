# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SmsProviderTwilioArgs', 'SmsProviderTwilio']

@pulumi.input_type
class SmsProviderTwilioArgs:
    def __init__(__self__, *,
                 sender_number: pulumi.Input[str],
                 sid: pulumi.Input[str],
                 token: pulumi.Input[str]):
        """
        The set of arguments for constructing a SmsProviderTwilio resource.
        :param pulumi.Input[str] sender_number: Sender number which is used to send the SMS.
        :param pulumi.Input[str] sid: SID used to communicate with Twilio.
        :param pulumi.Input[str] token: Token used to communicate with Twilio.
        """
        SmsProviderTwilioArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            sender_number=sender_number,
            sid=sid,
            token=token,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             sender_number: pulumi.Input[str],
             sid: pulumi.Input[str],
             token: pulumi.Input[str],
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'senderNumber' in kwargs:
            sender_number = kwargs['senderNumber']

        _setter("sender_number", sender_number)
        _setter("sid", sid)
        _setter("token", token)

    @property
    @pulumi.getter(name="senderNumber")
    def sender_number(self) -> pulumi.Input[str]:
        """
        Sender number which is used to send the SMS.
        """
        return pulumi.get(self, "sender_number")

    @sender_number.setter
    def sender_number(self, value: pulumi.Input[str]):
        pulumi.set(self, "sender_number", value)

    @property
    @pulumi.getter
    def sid(self) -> pulumi.Input[str]:
        """
        SID used to communicate with Twilio.
        """
        return pulumi.get(self, "sid")

    @sid.setter
    def sid(self, value: pulumi.Input[str]):
        pulumi.set(self, "sid", value)

    @property
    @pulumi.getter
    def token(self) -> pulumi.Input[str]:
        """
        Token used to communicate with Twilio.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: pulumi.Input[str]):
        pulumi.set(self, "token", value)


@pulumi.input_type
class _SmsProviderTwilioState:
    def __init__(__self__, *,
                 sender_number: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SmsProviderTwilio resources.
        :param pulumi.Input[str] sender_number: Sender number which is used to send the SMS.
        :param pulumi.Input[str] sid: SID used to communicate with Twilio.
        :param pulumi.Input[str] token: Token used to communicate with Twilio.
        """
        _SmsProviderTwilioState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            sender_number=sender_number,
            sid=sid,
            token=token,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             sender_number: Optional[pulumi.Input[str]] = None,
             sid: Optional[pulumi.Input[str]] = None,
             token: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'senderNumber' in kwargs:
            sender_number = kwargs['senderNumber']

        if sender_number is not None:
            _setter("sender_number", sender_number)
        if sid is not None:
            _setter("sid", sid)
        if token is not None:
            _setter("token", token)

    @property
    @pulumi.getter(name="senderNumber")
    def sender_number(self) -> Optional[pulumi.Input[str]]:
        """
        Sender number which is used to send the SMS.
        """
        return pulumi.get(self, "sender_number")

    @sender_number.setter
    def sender_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sender_number", value)

    @property
    @pulumi.getter
    def sid(self) -> Optional[pulumi.Input[str]]:
        """
        SID used to communicate with Twilio.
        """
        return pulumi.get(self, "sid")

    @sid.setter
    def sid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sid", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Token used to communicate with Twilio.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class SmsProviderTwilio(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 sender_number: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing the SMS provider Twilio configuration of an instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.SmsProviderTwilio("default",
            sid="sid",
            sender_number="019920892",
            token="twilio_token")
        ```

        ## Import

        bash The resource can be imported using the ID format `<id[:token]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/smsProviderTwilio:SmsProviderTwilio imported '123456789012345678:12345678901234567890123456abcdef'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] sender_number: Sender number which is used to send the SMS.
        :param pulumi.Input[str] sid: SID used to communicate with Twilio.
        :param pulumi.Input[str] token: Token used to communicate with Twilio.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SmsProviderTwilioArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the SMS provider Twilio configuration of an instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.SmsProviderTwilio("default",
            sid="sid",
            sender_number="019920892",
            token="twilio_token")
        ```

        ## Import

        bash The resource can be imported using the ID format `<id[:token]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/smsProviderTwilio:SmsProviderTwilio imported '123456789012345678:12345678901234567890123456abcdef'
        ```

        :param str resource_name: The name of the resource.
        :param SmsProviderTwilioArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SmsProviderTwilioArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            SmsProviderTwilioArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 sender_number: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SmsProviderTwilioArgs.__new__(SmsProviderTwilioArgs)

            if sender_number is None and not opts.urn:
                raise TypeError("Missing required property 'sender_number'")
            __props__.__dict__["sender_number"] = sender_number
            if sid is None and not opts.urn:
                raise TypeError("Missing required property 'sid'")
            __props__.__dict__["sid"] = sid
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SmsProviderTwilio, __self__).__init__(
            'zitadel:index/smsProviderTwilio:SmsProviderTwilio',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            sender_number: Optional[pulumi.Input[str]] = None,
            sid: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None) -> 'SmsProviderTwilio':
        """
        Get an existing SmsProviderTwilio resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] sender_number: Sender number which is used to send the SMS.
        :param pulumi.Input[str] sid: SID used to communicate with Twilio.
        :param pulumi.Input[str] token: Token used to communicate with Twilio.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SmsProviderTwilioState.__new__(_SmsProviderTwilioState)

        __props__.__dict__["sender_number"] = sender_number
        __props__.__dict__["sid"] = sid
        __props__.__dict__["token"] = token
        return SmsProviderTwilio(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="senderNumber")
    def sender_number(self) -> pulumi.Output[str]:
        """
        Sender number which is used to send the SMS.
        """
        return pulumi.get(self, "sender_number")

    @property
    @pulumi.getter
    def sid(self) -> pulumi.Output[str]:
        """
        SID used to communicate with Twilio.
        """
        return pulumi.get(self, "sid")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        Token used to communicate with Twilio.
        """
        return pulumi.get(self, "token")

