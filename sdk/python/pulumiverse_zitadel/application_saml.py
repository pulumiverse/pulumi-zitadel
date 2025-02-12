# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationSamlArgs', 'ApplicationSaml']

@pulumi.input_type
class ApplicationSamlArgs:
    def __init__(__self__, *,
                 metadata_xml: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationSaml resource.
        :param pulumi.Input[str] metadata_xml: Metadata as XML file
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[str] name: Name of the application
        :param pulumi.Input[str] org_id: ID of the organization
        """
        ApplicationSamlArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            metadata_xml=metadata_xml,
            project_id=project_id,
            name=name,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             metadata_xml: pulumi.Input[str],
             project_id: pulumi.Input[str],
             name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'metadataXml' in kwargs:
            metadata_xml = kwargs['metadataXml']
        if 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        _setter("metadata_xml", metadata_xml)
        _setter("project_id", project_id)
        if name is not None:
            _setter("name", name)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="metadataXml")
    def metadata_xml(self) -> pulumi.Input[str]:
        """
        Metadata as XML file
        """
        return pulumi.get(self, "metadata_xml")

    @metadata_xml.setter
    def metadata_xml(self, value: pulumi.Input[str]):
        pulumi.set(self, "metadata_xml", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the application
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _ApplicationSamlState:
    def __init__(__self__, *,
                 metadata_xml: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationSaml resources.
        :param pulumi.Input[str] metadata_xml: Metadata as XML file
        :param pulumi.Input[str] name: Name of the application
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        """
        _ApplicationSamlState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            metadata_xml=metadata_xml,
            name=name,
            org_id=org_id,
            project_id=project_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             metadata_xml: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'metadataXml' in kwargs:
            metadata_xml = kwargs['metadataXml']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'projectId' in kwargs:
            project_id = kwargs['projectId']

        if metadata_xml is not None:
            _setter("metadata_xml", metadata_xml)
        if name is not None:
            _setter("name", name)
        if org_id is not None:
            _setter("org_id", org_id)
        if project_id is not None:
            _setter("project_id", project_id)

    @property
    @pulumi.getter(name="metadataXml")
    def metadata_xml(self) -> Optional[pulumi.Input[str]]:
        """
        Metadata as XML file
        """
        return pulumi.get(self, "metadata_xml")

    @metadata_xml.setter
    def metadata_xml(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_xml", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the application
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class ApplicationSaml(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metadata_xml: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing a SAML application belonging to a project, with all configuration possibilities.

        ## Example Usage

        ```python
        import pulumi
        import scoretechnologies_zitadel as zitadel

        default = zitadel.ApplicationSaml("default",
            org_id=data["zitadel_org"]["default"]["id"],
            project_id=data["zitadel_project"]["default"]["id"],
            metadata_xml=\"\"\"<?xml version="1.0"?>
        <md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata"
                             validUntil="2024-01-26T17:48:38Z"
                             cacheDuration="PT604800S"
                             entityID="http://example.com/saml/metadata">
            <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
                <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
                <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
                                             Location="http://example.com/saml/cas"
                                             index="1" />
                
            </md:SPSSODescriptor>
        </md:EntityDescriptor>\"\"\")
        ```

        ## Import

        bash The resource can be imported using the ID format `<id:project_id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/applicationSaml:ApplicationSaml imported '123456789012345678:123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] metadata_xml: Metadata as XML file
        :param pulumi.Input[str] name: Name of the application
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationSamlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a SAML application belonging to a project, with all configuration possibilities.

        ## Example Usage

        ```python
        import pulumi
        import scoretechnologies_zitadel as zitadel

        default = zitadel.ApplicationSaml("default",
            org_id=data["zitadel_org"]["default"]["id"],
            project_id=data["zitadel_project"]["default"]["id"],
            metadata_xml=\"\"\"<?xml version="1.0"?>
        <md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata"
                             validUntil="2024-01-26T17:48:38Z"
                             cacheDuration="PT604800S"
                             entityID="http://example.com/saml/metadata">
            <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
                <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
                <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
                                             Location="http://example.com/saml/cas"
                                             index="1" />
                
            </md:SPSSODescriptor>
        </md:EntityDescriptor>\"\"\")
        ```

        ## Import

        bash The resource can be imported using the ID format `<id:project_id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/applicationSaml:ApplicationSaml imported '123456789012345678:123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationSamlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationSamlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ApplicationSamlArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metadata_xml: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationSamlArgs.__new__(ApplicationSamlArgs)

            if metadata_xml is None and not opts.urn:
                raise TypeError("Missing required property 'metadata_xml'")
            __props__.__dict__["metadata_xml"] = None if metadata_xml is None else pulumi.Output.secret(metadata_xml)
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["metadataXml"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ApplicationSaml, __self__).__init__(
            'zitadel:index/applicationSaml:ApplicationSaml',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            metadata_xml: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None) -> 'ApplicationSaml':
        """
        Get an existing ApplicationSaml resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] metadata_xml: Metadata as XML file
        :param pulumi.Input[str] name: Name of the application
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] project_id: ID of the project
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationSamlState.__new__(_ApplicationSamlState)

        __props__.__dict__["metadata_xml"] = metadata_xml
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["project_id"] = project_id
        return ApplicationSaml(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="metadataXml")
    def metadata_xml(self) -> pulumi.Output[str]:
        """
        Metadata as XML file
        """
        return pulumi.get(self, "metadata_xml")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the application
        """
        return pulumi.get(self, "name")

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

