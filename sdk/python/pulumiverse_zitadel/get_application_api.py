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
    'GetApplicationApiResult',
    'AwaitableGetApplicationApiResult',
    'get_application_api',
    'get_application_api_output',
]

@pulumi.output_type
class GetApplicationApiResult:
    """
    A collection of values returned by getApplicationApi.
    """
    def __init__(__self__, app_id=None, auth_method_type=None, id=None, name=None, org_id=None, project_id=None):
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        pulumi.set(__self__, "app_id", app_id)
        if auth_method_type and not isinstance(auth_method_type, str):
            raise TypeError("Expected argument 'auth_method_type' to be a str")
        pulumi.set(__self__, "auth_method_type", auth_method_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="authMethodType")
    def auth_method_type(self) -> str:
        """
        Auth method type
        """
        return pulumi.get(self, "auth_method_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the application
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
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")


class AwaitableGetApplicationApiResult(GetApplicationApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationApiResult(
            app_id=self.app_id,
            auth_method_type=self.auth_method_type,
            id=self.id,
            name=self.name,
            org_id=self.org_id,
            project_id=self.project_id)


def get_application_api(app_id: Optional[str] = None,
                        org_id: Optional[str] = None,
                        project_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationApiResult:
    """
    Datasource representing an API application belonging to a project, with all configuration possibilities.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_application_api(org_id=data["zitadel_org"]["default"]["id"],
        project_id=data["zitadel_project"]["default"]["id"],
        app_id="123456789012345678")
    ```


    :param str app_id: The ID of this resource.
    :param str org_id: ID of the organization
    :param str project_id: ID of the project
    """
    __args__ = dict()
    __args__['appId'] = app_id
    __args__['orgId'] = org_id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getApplicationApi:getApplicationApi', __args__, opts=opts, typ=GetApplicationApiResult).value

    return AwaitableGetApplicationApiResult(
        app_id=pulumi.get(__ret__, 'app_id'),
        auth_method_type=pulumi.get(__ret__, 'auth_method_type'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        org_id=pulumi.get(__ret__, 'org_id'),
        project_id=pulumi.get(__ret__, 'project_id'))


@_utilities.lift_output_func(get_application_api)
def get_application_api_output(app_id: Optional[pulumi.Input[str]] = None,
                               org_id: Optional[pulumi.Input[Optional[str]]] = None,
                               project_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationApiResult]:
    """
    Datasource representing an API application belonging to a project, with all configuration possibilities.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_application_api(org_id=data["zitadel_org"]["default"]["id"],
        project_id=data["zitadel_project"]["default"]["id"],
        app_id="123456789012345678")
    ```


    :param str app_id: The ID of this resource.
    :param str org_id: ID of the organization
    :param str project_id: ID of the project
    """
    ...
