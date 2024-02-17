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
    'GetProjectsResult',
    'AwaitableGetProjectsResult',
    'get_projects',
    'get_projects_output',
]

@pulumi.output_type
class GetProjectsResult:
    """
    A collection of values returned by getProjects.
    """
    def __init__(__self__, id=None, name=None, name_method=None, org_id=None, project_ids=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_method and not isinstance(name_method, str):
            raise TypeError("Expected argument 'name_method' to be a str")
        pulumi.set(__self__, "name_method", name_method)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if project_ids and not isinstance(project_ids, list):
            raise TypeError("Expected argument 'project_ids' to be a list")
        pulumi.set(__self__, "project_ids", project_ids)

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
    @pulumi.getter(name="nameMethod")
    def name_method(self) -> Optional[str]:
        """
        Method for querying projects by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        """
        return pulumi.get(self, "name_method")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[str]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> Sequence[str]:
        """
        A set of all project IDs.
        """
        return pulumi.get(self, "project_ids")


class AwaitableGetProjectsResult(GetProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectsResult(
            id=self.id,
            name=self.name,
            name_method=self.name_method,
            org_id=self.org_id,
            project_ids=self.project_ids)


def get_projects(name: Optional[str] = None,
                 name_method: Optional[str] = None,
                 org_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectsResult:
    """
    Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.


    :param str name: Name of the project
    :param str name_method: Method for querying projects by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    :param str org_id: ID of the organization
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['nameMethod'] = name_method
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getProjects:getProjects', __args__, opts=opts, typ=GetProjectsResult).value

    return AwaitableGetProjectsResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        name_method=pulumi.get(__ret__, 'name_method'),
        org_id=pulumi.get(__ret__, 'org_id'),
        project_ids=pulumi.get(__ret__, 'project_ids'))


@_utilities.lift_output_func(get_projects)
def get_projects_output(name: Optional[pulumi.Input[str]] = None,
                        name_method: Optional[pulumi.Input[Optional[str]]] = None,
                        org_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectsResult]:
    """
    Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.


    :param str name: Name of the project
    :param str name_method: Method for querying projects by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    :param str org_id: ID of the organization
    """
    ...