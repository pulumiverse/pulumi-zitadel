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
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, has_project_check=None, id=None, name=None, org_id=None, private_labeling_setting=None, project_id=None, project_role_assertion=None, project_role_check=None, state=None):
        if has_project_check and not isinstance(has_project_check, bool):
            raise TypeError("Expected argument 'has_project_check' to be a bool")
        pulumi.set(__self__, "has_project_check", has_project_check)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if private_labeling_setting and not isinstance(private_labeling_setting, str):
            raise TypeError("Expected argument 'private_labeling_setting' to be a str")
        pulumi.set(__self__, "private_labeling_setting", private_labeling_setting)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if project_role_assertion and not isinstance(project_role_assertion, bool):
            raise TypeError("Expected argument 'project_role_assertion' to be a bool")
        pulumi.set(__self__, "project_role_assertion", project_role_assertion)
        if project_role_check and not isinstance(project_role_check, bool):
            raise TypeError("Expected argument 'project_role_check' to be a bool")
        pulumi.set(__self__, "project_role_check", project_role_check)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="hasProjectCheck")
    def has_project_check(self) -> bool:
        """
        ZITADEL checks if the org of the user has permission to this project
        """
        return pulumi.get(self, "has_project_check")

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
        Name of the project
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
    @pulumi.getter(name="privateLabelingSetting")
    def private_labeling_setting(self) -> str:
        """
        Defines from where the private labeling should be triggered
        """
        return pulumi.get(self, "private_labeling_setting")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectRoleAssertion")
    def project_role_assertion(self) -> bool:
        """
        describes if roles of user should be added in token
        """
        return pulumi.get(self, "project_role_assertion")

    @property
    @pulumi.getter(name="projectRoleCheck")
    def project_role_check(self) -> bool:
        """
        ZITADEL checks if the user has at least one on this project
        """
        return pulumi.get(self, "project_role_check")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the project
        """
        return pulumi.get(self, "state")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            has_project_check=self.has_project_check,
            id=self.id,
            name=self.name,
            org_id=self.org_id,
            private_labeling_setting=self.private_labeling_setting,
            project_id=self.project_id,
            project_role_assertion=self.project_role_assertion,
            project_role_check=self.project_role_check,
            state=self.state)


def get_project(org_id: Optional[str] = None,
                project_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_project(org_id=default_zitadel_org["id"],
        project_id="123456789012345678")
    ```


    :param str org_id: ID of the organization
    :param str project_id: The ID of this resource.
    """
    __args__ = dict()
    __args__['orgId'] = org_id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        has_project_check=pulumi.get(__ret__, 'has_project_check'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        org_id=pulumi.get(__ret__, 'org_id'),
        private_labeling_setting=pulumi.get(__ret__, 'private_labeling_setting'),
        project_id=pulumi.get(__ret__, 'project_id'),
        project_role_assertion=pulumi.get(__ret__, 'project_role_assertion'),
        project_role_check=pulumi.get(__ret__, 'project_role_check'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_project)
def get_project_output(org_id: Optional[pulumi.Input[Optional[str]]] = None,
                       project_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectResult]:
    """
    Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_project(org_id=default_zitadel_org["id"],
        project_id="123456789012345678")
    ```


    :param str org_id: ID of the organization
    :param str project_id: The ID of this resource.
    """
    ...
