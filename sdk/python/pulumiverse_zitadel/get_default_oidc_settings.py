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
    'GetDefaultOidcSettingsResult',
    'AwaitableGetDefaultOidcSettingsResult',
    'get_default_oidc_settings',
    'get_default_oidc_settings_output',
]

@pulumi.output_type
class GetDefaultOidcSettingsResult:
    """
    A collection of values returned by getDefaultOidcSettings.
    """
    def __init__(__self__, access_token_lifetime=None, id=None, id_token_lifetime=None, refresh_token_expiration=None, refresh_token_idle_expiration=None):
        if access_token_lifetime and not isinstance(access_token_lifetime, str):
            raise TypeError("Expected argument 'access_token_lifetime' to be a str")
        pulumi.set(__self__, "access_token_lifetime", access_token_lifetime)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if id_token_lifetime and not isinstance(id_token_lifetime, str):
            raise TypeError("Expected argument 'id_token_lifetime' to be a str")
        pulumi.set(__self__, "id_token_lifetime", id_token_lifetime)
        if refresh_token_expiration and not isinstance(refresh_token_expiration, str):
            raise TypeError("Expected argument 'refresh_token_expiration' to be a str")
        pulumi.set(__self__, "refresh_token_expiration", refresh_token_expiration)
        if refresh_token_idle_expiration and not isinstance(refresh_token_idle_expiration, str):
            raise TypeError("Expected argument 'refresh_token_idle_expiration' to be a str")
        pulumi.set(__self__, "refresh_token_idle_expiration", refresh_token_idle_expiration)

    @property
    @pulumi.getter(name="accessTokenLifetime")
    def access_token_lifetime(self) -> str:
        """
        lifetime duration of access tokens
        """
        return pulumi.get(self, "access_token_lifetime")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idTokenLifetime")
    def id_token_lifetime(self) -> str:
        """
        lifetime duration of id tokens
        """
        return pulumi.get(self, "id_token_lifetime")

    @property
    @pulumi.getter(name="refreshTokenExpiration")
    def refresh_token_expiration(self) -> str:
        """
        expiration duration of refresh tokens
        """
        return pulumi.get(self, "refresh_token_expiration")

    @property
    @pulumi.getter(name="refreshTokenIdleExpiration")
    def refresh_token_idle_expiration(self) -> str:
        """
        expiration duration of idle refresh tokens
        """
        return pulumi.get(self, "refresh_token_idle_expiration")


class AwaitableGetDefaultOidcSettingsResult(GetDefaultOidcSettingsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultOidcSettingsResult(
            access_token_lifetime=self.access_token_lifetime,
            id=self.id,
            id_token_lifetime=self.id_token_lifetime,
            refresh_token_expiration=self.refresh_token_expiration,
            refresh_token_idle_expiration=self.refresh_token_idle_expiration)


def get_default_oidc_settings(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultOidcSettingsResult:
    """
    Datasource representing the default oidc settings.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_default_oidc_settings()
    pulumi.export("oidcSettings", default)
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getDefaultOidcSettings:getDefaultOidcSettings', __args__, opts=opts, typ=GetDefaultOidcSettingsResult).value

    return AwaitableGetDefaultOidcSettingsResult(
        access_token_lifetime=pulumi.get(__ret__, 'access_token_lifetime'),
        id=pulumi.get(__ret__, 'id'),
        id_token_lifetime=pulumi.get(__ret__, 'id_token_lifetime'),
        refresh_token_expiration=pulumi.get(__ret__, 'refresh_token_expiration'),
        refresh_token_idle_expiration=pulumi.get(__ret__, 'refresh_token_idle_expiration'))


@_utilities.lift_output_func(get_default_oidc_settings)
def get_default_oidc_settings_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultOidcSettingsResult]:
    """
    Datasource representing the default oidc settings.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_default_oidc_settings()
    pulumi.export("oidcSettings", default)
    ```
    """
    ...
