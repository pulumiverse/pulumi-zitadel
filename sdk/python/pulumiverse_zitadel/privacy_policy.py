# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PrivacyPolicyArgs', 'PrivacyPolicy']

@pulumi.input_type
class PrivacyPolicyArgs:
    def __init__(__self__, *,
                 help_link: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 privacy_link: Optional[pulumi.Input[str]] = None,
                 support_email: Optional[pulumi.Input[str]] = None,
                 tos_link: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivacyPolicy resource.
        :param pulumi.Input[str] org_id: ID of the organization
        """
        if help_link is not None:
            pulumi.set(__self__, "help_link", help_link)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if privacy_link is not None:
            pulumi.set(__self__, "privacy_link", privacy_link)
        if support_email is not None:
            pulumi.set(__self__, "support_email", support_email)
        if tos_link is not None:
            pulumi.set(__self__, "tos_link", tos_link)

    @property
    @pulumi.getter(name="helpLink")
    def help_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "help_link")

    @help_link.setter
    def help_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "help_link", value)

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
    @pulumi.getter(name="privacyLink")
    def privacy_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "privacy_link")

    @privacy_link.setter
    def privacy_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy_link", value)

    @property
    @pulumi.getter(name="supportEmail")
    def support_email(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "support_email")

    @support_email.setter
    def support_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "support_email", value)

    @property
    @pulumi.getter(name="tosLink")
    def tos_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tos_link")

    @tos_link.setter
    def tos_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tos_link", value)


@pulumi.input_type
class _PrivacyPolicyState:
    def __init__(__self__, *,
                 help_link: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 privacy_link: Optional[pulumi.Input[str]] = None,
                 support_email: Optional[pulumi.Input[str]] = None,
                 tos_link: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivacyPolicy resources.
        :param pulumi.Input[str] org_id: ID of the organization
        """
        if help_link is not None:
            pulumi.set(__self__, "help_link", help_link)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if privacy_link is not None:
            pulumi.set(__self__, "privacy_link", privacy_link)
        if support_email is not None:
            pulumi.set(__self__, "support_email", support_email)
        if tos_link is not None:
            pulumi.set(__self__, "tos_link", tos_link)

    @property
    @pulumi.getter(name="helpLink")
    def help_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "help_link")

    @help_link.setter
    def help_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "help_link", value)

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
    @pulumi.getter(name="privacyLink")
    def privacy_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "privacy_link")

    @privacy_link.setter
    def privacy_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy_link", value)

    @property
    @pulumi.getter(name="supportEmail")
    def support_email(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "support_email")

    @support_email.setter
    def support_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "support_email", value)

    @property
    @pulumi.getter(name="tosLink")
    def tos_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tos_link")

    @tos_link.setter
    def tos_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tos_link", value)


class PrivacyPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 help_link: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 privacy_link: Optional[pulumi.Input[str]] = None,
                 support_email: Optional[pulumi.Input[str]] = None,
                 tos_link: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing the custom privacy policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.PrivacyPolicy("default",
            org_id=data["zitadel_org"]["default"]["id"],
            tos_link="https://example.com/tos",
            privacy_link="https://example.com/privacy",
            help_link="https://example.com/help",
            support_email="support@example.com")
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<[org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/privacyPolicy:PrivacyPolicy imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: ID of the organization
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PrivacyPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the custom privacy policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.PrivacyPolicy("default",
            org_id=data["zitadel_org"]["default"]["id"],
            tos_link="https://example.com/tos",
            privacy_link="https://example.com/privacy",
            help_link="https://example.com/help",
            support_email="support@example.com")
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<[org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/privacyPolicy:PrivacyPolicy imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param PrivacyPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivacyPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 help_link: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 privacy_link: Optional[pulumi.Input[str]] = None,
                 support_email: Optional[pulumi.Input[str]] = None,
                 tos_link: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivacyPolicyArgs.__new__(PrivacyPolicyArgs)

            __props__.__dict__["help_link"] = help_link
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["privacy_link"] = privacy_link
            __props__.__dict__["support_email"] = support_email
            __props__.__dict__["tos_link"] = tos_link
        super(PrivacyPolicy, __self__).__init__(
            'zitadel:index/privacyPolicy:PrivacyPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            help_link: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            privacy_link: Optional[pulumi.Input[str]] = None,
            support_email: Optional[pulumi.Input[str]] = None,
            tos_link: Optional[pulumi.Input[str]] = None) -> 'PrivacyPolicy':
        """
        Get an existing PrivacyPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: ID of the organization
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivacyPolicyState.__new__(_PrivacyPolicyState)

        __props__.__dict__["help_link"] = help_link
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["privacy_link"] = privacy_link
        __props__.__dict__["support_email"] = support_email
        __props__.__dict__["tos_link"] = tos_link
        return PrivacyPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="helpLink")
    def help_link(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "help_link")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="privacyLink")
    def privacy_link(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "privacy_link")

    @property
    @pulumi.getter(name="supportEmail")
    def support_email(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "support_email")

    @property
    @pulumi.getter(name="tosLink")
    def tos_link(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tos_link")

