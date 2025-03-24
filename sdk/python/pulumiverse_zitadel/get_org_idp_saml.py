# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOrgIdpSamlResult',
    'AwaitableGetOrgIdpSamlResult',
    'get_org_idp_saml',
    'get_org_idp_saml_output',
]

@pulumi.output_type
class GetOrgIdpSamlResult:
    """
    A collection of values returned by getOrgIdpSaml.
    """
    def __init__(__self__, binding=None, id=None, is_auto_creation=None, is_auto_update=None, is_creation_allowed=None, is_linking_allowed=None, metadata_xml=None, name=None, org_id=None, with_signed_request=None):
        if binding and not isinstance(binding, str):
            raise TypeError("Expected argument 'binding' to be a str")
        pulumi.set(__self__, "binding", binding)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_auto_creation and not isinstance(is_auto_creation, bool):
            raise TypeError("Expected argument 'is_auto_creation' to be a bool")
        pulumi.set(__self__, "is_auto_creation", is_auto_creation)
        if is_auto_update and not isinstance(is_auto_update, bool):
            raise TypeError("Expected argument 'is_auto_update' to be a bool")
        pulumi.set(__self__, "is_auto_update", is_auto_update)
        if is_creation_allowed and not isinstance(is_creation_allowed, bool):
            raise TypeError("Expected argument 'is_creation_allowed' to be a bool")
        pulumi.set(__self__, "is_creation_allowed", is_creation_allowed)
        if is_linking_allowed and not isinstance(is_linking_allowed, bool):
            raise TypeError("Expected argument 'is_linking_allowed' to be a bool")
        pulumi.set(__self__, "is_linking_allowed", is_linking_allowed)
        if metadata_xml and not isinstance(metadata_xml, str):
            raise TypeError("Expected argument 'metadata_xml' to be a str")
        pulumi.set(__self__, "metadata_xml", metadata_xml)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if with_signed_request and not isinstance(with_signed_request, str):
            raise TypeError("Expected argument 'with_signed_request' to be a str")
        pulumi.set(__self__, "with_signed_request", with_signed_request)

    @property
    @pulumi.getter
    def binding(self) -> str:
        """
        The binding
        """
        return pulumi.get(self, "binding")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> bool:
        """
        enabled if a new account in ZITADEL are created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> bool:
        """
        enabled if a the ZITADEL account fields are updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> bool:
        """
        enabled if users are able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> bool:
        """
        enabled if users are able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @property
    @pulumi.getter(name="metadataXml")
    def metadata_xml(self) -> str:
        """
        The metadata XML as plain string
        """
        return pulumi.get(self, "metadata_xml")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[str]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="withSignedRequest")
    def with_signed_request(self) -> str:
        """
        Whether the SAML IDP requires signed requests
        """
        return pulumi.get(self, "with_signed_request")


class AwaitableGetOrgIdpSamlResult(GetOrgIdpSamlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgIdpSamlResult(
            binding=self.binding,
            id=self.id,
            is_auto_creation=self.is_auto_creation,
            is_auto_update=self.is_auto_update,
            is_creation_allowed=self.is_creation_allowed,
            is_linking_allowed=self.is_linking_allowed,
            metadata_xml=self.metadata_xml,
            name=self.name,
            org_id=self.org_id,
            with_signed_request=self.with_signed_request)


def get_org_idp_saml(id: Optional[str] = None,
                     org_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgIdpSamlResult:
    """
    Datasource representing a SAML IdP of the organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_org_idp_saml(org_id=default_zitadel_org["id"],
        id="123456789012345678")
    ```


    :param str id: The ID of this resource.
    :param str org_id: ID of the organization
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getOrgIdpSaml:getOrgIdpSaml', __args__, opts=opts, typ=GetOrgIdpSamlResult).value

    return AwaitableGetOrgIdpSamlResult(
        binding=pulumi.get(__ret__, 'binding'),
        id=pulumi.get(__ret__, 'id'),
        is_auto_creation=pulumi.get(__ret__, 'is_auto_creation'),
        is_auto_update=pulumi.get(__ret__, 'is_auto_update'),
        is_creation_allowed=pulumi.get(__ret__, 'is_creation_allowed'),
        is_linking_allowed=pulumi.get(__ret__, 'is_linking_allowed'),
        metadata_xml=pulumi.get(__ret__, 'metadata_xml'),
        name=pulumi.get(__ret__, 'name'),
        org_id=pulumi.get(__ret__, 'org_id'),
        with_signed_request=pulumi.get(__ret__, 'with_signed_request'))


@_utilities.lift_output_func(get_org_idp_saml)
def get_org_idp_saml_output(id: Optional[pulumi.Input[str]] = None,
                            org_id: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrgIdpSamlResult]:
    """
    Datasource representing a SAML IdP of the organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_org_idp_saml(org_id=default_zitadel_org["id"],
        id="123456789012345678")
    ```


    :param str id: The ID of this resource.
    :param str org_id: ID of the organization
    """
    ...
