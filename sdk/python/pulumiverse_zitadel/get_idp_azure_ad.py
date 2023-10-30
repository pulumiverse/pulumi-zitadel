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
    'GetIdpAzureAdResult',
    'AwaitableGetIdpAzureAdResult',
    'get_idp_azure_ad',
    'get_idp_azure_ad_output',
]

@pulumi.output_type
class GetIdpAzureAdResult:
    """
    A collection of values returned by getIdpAzureAd.
    """
    def __init__(__self__, client_id=None, client_secret=None, email_verified=None, id=None, is_auto_creation=None, is_auto_update=None, is_creation_allowed=None, is_linking_allowed=None, name=None, scopes=None, tenant_id=None, tenant_type=None):
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if email_verified and not isinstance(email_verified, bool):
            raise TypeError("Expected argument 'email_verified' to be a bool")
        pulumi.set(__self__, "email_verified", email_verified)
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
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)
        if tenant_type and not isinstance(tenant_type, str):
            raise TypeError("Expected argument 'tenant_type' to be a str")
        pulumi.set(__self__, "tenant_type", tenant_type)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> str:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> bool:
        """
        automatically mark emails as verified
        """
        return pulumi.get(self, "email_verified")

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
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence[str]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        the azure ad tenant id
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="tenantType")
    def tenant_type(self) -> str:
        """
        the azure ad tenant type
        """
        return pulumi.get(self, "tenant_type")


class AwaitableGetIdpAzureAdResult(GetIdpAzureAdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdpAzureAdResult(
            client_id=self.client_id,
            client_secret=self.client_secret,
            email_verified=self.email_verified,
            id=self.id,
            is_auto_creation=self.is_auto_creation,
            is_auto_update=self.is_auto_update,
            is_creation_allowed=self.is_creation_allowed,
            is_linking_allowed=self.is_linking_allowed,
            name=self.name,
            scopes=self.scopes,
            tenant_id=self.tenant_id,
            tenant_type=self.tenant_type)


def get_idp_azure_ad(id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdpAzureAdResult:
    """
    Datasource representing an Azure AD IDP on the instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_idp_azure_ad(id="123456789012345678")
    ```


    :param str id: The ID of this resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getIdpAzureAd:getIdpAzureAd', __args__, opts=opts, typ=GetIdpAzureAdResult).value

    return AwaitableGetIdpAzureAdResult(
        client_id=pulumi.get(__ret__, 'client_id'),
        client_secret=pulumi.get(__ret__, 'client_secret'),
        email_verified=pulumi.get(__ret__, 'email_verified'),
        id=pulumi.get(__ret__, 'id'),
        is_auto_creation=pulumi.get(__ret__, 'is_auto_creation'),
        is_auto_update=pulumi.get(__ret__, 'is_auto_update'),
        is_creation_allowed=pulumi.get(__ret__, 'is_creation_allowed'),
        is_linking_allowed=pulumi.get(__ret__, 'is_linking_allowed'),
        name=pulumi.get(__ret__, 'name'),
        scopes=pulumi.get(__ret__, 'scopes'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'),
        tenant_type=pulumi.get(__ret__, 'tenant_type'))


@_utilities.lift_output_func(get_idp_azure_ad)
def get_idp_azure_ad_output(id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIdpAzureAdResult]:
    """
    Datasource representing an Azure AD IDP on the instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_idp_azure_ad(id="123456789012345678")
    ```


    :param str id: The ID of this resource.
    """
    ...
