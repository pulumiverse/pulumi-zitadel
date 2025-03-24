# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ActionArgs', 'Action']

@pulumi.input_type
class ActionArgs:
    def __init__(__self__, *,
                 allowed_to_fail: pulumi.Input[bool],
                 script: pulumi.Input[str],
                 timeout: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Action resource.
        :param pulumi.Input[bool] allowed_to_fail: when true, the next action will be called even if this action fails
        :param pulumi.Input[str] timeout: after which time the action will be terminated if not finished
        :param pulumi.Input[str] org_id: ID of the organization
        """
        ActionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            allowed_to_fail=allowed_to_fail,
            script=script,
            timeout=timeout,
            name=name,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             allowed_to_fail: pulumi.Input[bool],
             script: pulumi.Input[str],
             timeout: pulumi.Input[str],
             name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'allowedToFail' in kwargs:
            allowed_to_fail = kwargs['allowedToFail']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        _setter("allowed_to_fail", allowed_to_fail)
        _setter("script", script)
        _setter("timeout", timeout)
        if name is not None:
            _setter("name", name)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="allowedToFail")
    def allowed_to_fail(self) -> pulumi.Input[bool]:
        """
        when true, the next action will be called even if this action fails
        """
        return pulumi.get(self, "allowed_to_fail")

    @allowed_to_fail.setter
    def allowed_to_fail(self, value: pulumi.Input[bool]):
        pulumi.set(self, "allowed_to_fail", value)

    @property
    @pulumi.getter
    def script(self) -> pulumi.Input[str]:
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: pulumi.Input[str]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Input[str]:
        """
        after which time the action will be terminated if not finished
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: pulumi.Input[str]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
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
class _ActionState:
    def __init__(__self__, *,
                 allowed_to_fail: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Action resources.
        :param pulumi.Input[bool] allowed_to_fail: when true, the next action will be called even if this action fails
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[int] state: the state of the action
        :param pulumi.Input[str] timeout: after which time the action will be terminated if not finished
        """
        _ActionState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            allowed_to_fail=allowed_to_fail,
            name=name,
            org_id=org_id,
            script=script,
            state=state,
            timeout=timeout,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             allowed_to_fail: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             script: Optional[pulumi.Input[str]] = None,
             state: Optional[pulumi.Input[int]] = None,
             timeout: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'allowedToFail' in kwargs:
            allowed_to_fail = kwargs['allowedToFail']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        if allowed_to_fail is not None:
            _setter("allowed_to_fail", allowed_to_fail)
        if name is not None:
            _setter("name", name)
        if org_id is not None:
            _setter("org_id", org_id)
        if script is not None:
            _setter("script", script)
        if state is not None:
            _setter("state", state)
        if timeout is not None:
            _setter("timeout", timeout)

    @property
    @pulumi.getter(name="allowedToFail")
    def allowed_to_fail(self) -> Optional[pulumi.Input[bool]]:
        """
        when true, the next action will be called even if this action fails
        """
        return pulumi.get(self, "allowed_to_fail")

    @allowed_to_fail.setter
    def allowed_to_fail(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allowed_to_fail", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
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
    @pulumi.getter
    def script(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[int]]:
        """
        the state of the action
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        """
        after which time the action will be terminated if not finished
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)


class Action(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_to_fail: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing an action belonging to an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.Action("default",
            org_id=default_zitadel_org["id"],
            name="actionname",
            script="testscript",
            timeout="10s",
            allowed_to_fail=True)
        ```

        ## Import

        bash The resource can be imported using the ID format `<id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/action:Action imported '123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allowed_to_fail: when true, the next action will be called even if this action fails
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] timeout: after which time the action will be terminated if not finished
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing an action belonging to an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.Action("default",
            org_id=default_zitadel_org["id"],
            name="actionname",
            script="testscript",
            timeout="10s",
            allowed_to_fail=True)
        ```

        ## Import

        bash The resource can be imported using the ID format `<id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/action:Action imported '123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param ActionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ActionArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_to_fail: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActionArgs.__new__(ActionArgs)

            if allowed_to_fail is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_to_fail'")
            __props__.__dict__["allowed_to_fail"] = allowed_to_fail
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
            if script is None and not opts.urn:
                raise TypeError("Missing required property 'script'")
            __props__.__dict__["script"] = script
            if timeout is None and not opts.urn:
                raise TypeError("Missing required property 'timeout'")
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["state"] = None
        super(Action, __self__).__init__(
            'zitadel:index/action:Action',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_to_fail: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            script: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[int]] = None,
            timeout: Optional[pulumi.Input[str]] = None) -> 'Action':
        """
        Get an existing Action resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allowed_to_fail: when true, the next action will be called even if this action fails
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[int] state: the state of the action
        :param pulumi.Input[str] timeout: after which time the action will be terminated if not finished
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActionState.__new__(_ActionState)

        __props__.__dict__["allowed_to_fail"] = allowed_to_fail
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["script"] = script
        __props__.__dict__["state"] = state
        __props__.__dict__["timeout"] = timeout
        return Action(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedToFail")
    def allowed_to_fail(self) -> pulumi.Output[bool]:
        """
        when true, the next action will be called even if this action fails
        """
        return pulumi.get(self, "allowed_to_fail")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def script(self) -> pulumi.Output[str]:
        return pulumi.get(self, "script")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[int]:
        """
        the state of the action
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[str]:
        """
        after which time the action will be terminated if not finished
        """
        return pulumi.get(self, "timeout")

