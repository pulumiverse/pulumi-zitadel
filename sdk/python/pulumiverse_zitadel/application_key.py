# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationKeyArgs', 'ApplicationKey']

@pulumi.input_type
class ApplicationKeyArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 expiration_date: pulumi.Input[str],
                 key_type: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationKey resource.
        :param pulumi.Input[str] app_id: ID of the application
        :param pulumi.Input[str] expiration_date: Expiration date of the app key in the RFC3339 format
        :param pulumi.Input[str] key_type: Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[str] org_id: ID of the organization
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "expiration_date", expiration_date)
        pulumi.set(__self__, "key_type", key_type)
        pulumi.set(__self__, "project_id", project_id)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        ID of the application
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Input[str]:
        """
        Expiration date of the app key in the RFC3339 format
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: pulumi.Input[str]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Input[str]:
        """
        Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_type", value)

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


@pulumi.input_type
class _ApplicationKeyState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_details: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationKey resources.
        :param pulumi.Input[str] app_id: ID of the application
        :param pulumi.Input[str] expiration_date: Expiration date of the app key in the RFC3339 format
        :param pulumi.Input[str] key_details: Value of the app key
        :param pulumi.Input[str] key_type: Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if key_details is not None:
            pulumi.set(__self__, "key_details", key_details)
        if key_type is not None:
            pulumi.set(__self__, "key_type", key_type)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the application
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date of the app key in the RFC3339 format
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="keyDetails")
    def key_details(self) -> Optional[pulumi.Input[str]]:
        """
        Value of the app key
        """
        return pulumi.get(self, "key_details")

    @key_details.setter
    def key_details(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_details", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


class ApplicationKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing a app key

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.ApplicationKey("default",
            org_id=data["zitadel_org"]["default"]["id"],
            project_id=data["zitadel_project"]["default"]["id"],
            app_id=data["zitadel_application_api"]["default"]["id"],
            key_type="KEY_TYPE_JSON",
            expiration_date="2519-04-01T08:45:00Z")
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<id:project_id:app_id[:org_id][:key_details]>`. # You can use __SEMICOLON__ to escape :, e.g.

        ```sh
         $ pulumi import zitadel:index/applicationKey:ApplicationKey imported "123456789012345678:123456789012345678:123456789012345678:123456789012345678:$(cat ~/Downloads/123456789012345678.json | sed -e 's/:/__SEMICOLON__/g')"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: ID of the application
        :param pulumi.Input[str] expiration_date: Expiration date of the app key in the RFC3339 format
        :param pulumi.Input[str] key_type: Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a app key

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.ApplicationKey("default",
            org_id=data["zitadel_org"]["default"]["id"],
            project_id=data["zitadel_project"]["default"]["id"],
            app_id=data["zitadel_application_api"]["default"]["id"],
            key_type="KEY_TYPE_JSON",
            expiration_date="2519-04-01T08:45:00Z")
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<id:project_id:app_id[:org_id][:key_details]>`. # You can use __SEMICOLON__ to escape :, e.g.

        ```sh
         $ pulumi import zitadel:index/applicationKey:ApplicationKey imported "123456789012345678:123456789012345678:123456789012345678:123456789012345678:$(cat ~/Downloads/123456789012345678.json | sed -e 's/:/__SEMICOLON__/g')"
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationKeyArgs.__new__(ApplicationKeyArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            if expiration_date is None and not opts.urn:
                raise TypeError("Missing required property 'expiration_date'")
            __props__.__dict__["expiration_date"] = expiration_date
            if key_type is None and not opts.urn:
                raise TypeError("Missing required property 'key_type'")
            __props__.__dict__["key_type"] = key_type
            __props__.__dict__["org_id"] = org_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["key_details"] = None
        super(ApplicationKey, __self__).__init__(
            'zitadel:index/applicationKey:ApplicationKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            key_details: Optional[pulumi.Input[str]] = None,
            key_type: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None) -> 'ApplicationKey':
        """
        Get an existing ApplicationKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: ID of the application
        :param pulumi.Input[str] expiration_date: Expiration date of the app key in the RFC3339 format
        :param pulumi.Input[str] key_details: Value of the app key
        :param pulumi.Input[str] key_type: Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationKeyState.__new__(_ApplicationKeyState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["expiration_date"] = expiration_date
        __props__.__dict__["key_details"] = key_details
        __props__.__dict__["key_type"] = key_type
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["project_id"] = project_id
        return ApplicationKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        ID of the application
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Output[str]:
        """
        Expiration date of the app key in the RFC3339 format
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="keyDetails")
    def key_details(self) -> pulumi.Output[str]:
        """
        Value of the app key
        """
        return pulumi.get(self, "key_details")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[str]:
        """
        Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
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
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")

