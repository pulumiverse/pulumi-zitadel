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
    'GetMachineUsersResult',
    'AwaitableGetMachineUsersResult',
    'get_machine_users',
    'get_machine_users_output',
]

@pulumi.output_type
class GetMachineUsersResult:
    """
    A collection of values returned by getMachineUsers.
    """
    def __init__(__self__, id=None, org_id=None, user_ids=None, user_name=None, user_name_method=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if user_ids and not isinstance(user_ids, list):
            raise TypeError("Expected argument 'user_ids' to be a list")
        pulumi.set(__self__, "user_ids", user_ids)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)
        if user_name_method and not isinstance(user_name_method, str):
            raise TypeError("Expected argument 'user_name_method' to be a str")
        pulumi.set(__self__, "user_name_method", user_name_method)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[str]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Sequence[str]:
        """
        A set of all IDs.
        """
        return pulumi.get(self, "user_ids")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        Username
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="userNameMethod")
    def user_name_method(self) -> Optional[str]:
        """
        Method for querying machine users by username, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        """
        return pulumi.get(self, "user_name_method")


class AwaitableGetMachineUsersResult(GetMachineUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMachineUsersResult(
            id=self.id,
            org_id=self.org_id,
            user_ids=self.user_ids,
            user_name=self.user_name,
            user_name_method=self.user_name_method)


def get_machine_users(org_id: Optional[str] = None,
                      user_name: Optional[str] = None,
                      user_name_method: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMachineUsersResult:
    """
    Datasource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.


    :param str org_id: ID of the organization
    :param str user_name: Username
    :param str user_name_method: Method for querying machine users by username, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    """
    __args__ = dict()
    __args__['orgId'] = org_id
    __args__['userName'] = user_name
    __args__['userNameMethod'] = user_name_method
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getMachineUsers:getMachineUsers', __args__, opts=opts, typ=GetMachineUsersResult).value

    return AwaitableGetMachineUsersResult(
        id=pulumi.get(__ret__, 'id'),
        org_id=pulumi.get(__ret__, 'org_id'),
        user_ids=pulumi.get(__ret__, 'user_ids'),
        user_name=pulumi.get(__ret__, 'user_name'),
        user_name_method=pulumi.get(__ret__, 'user_name_method'))


@_utilities.lift_output_func(get_machine_users)
def get_machine_users_output(org_id: Optional[pulumi.Input[Optional[str]]] = None,
                             user_name: Optional[pulumi.Input[str]] = None,
                             user_name_method: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMachineUsersResult]:
    """
    Datasource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.


    :param str org_id: ID of the organization
    :param str user_name: Username
    :param str user_name_method: Method for querying machine users by username, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    """
    ...
