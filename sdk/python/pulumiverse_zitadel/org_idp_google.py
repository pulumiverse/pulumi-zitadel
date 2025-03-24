# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['OrgIdpGoogleArgs', 'OrgIdpGoogle']

@pulumi.input_type
class OrgIdpGoogleArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 client_secret: pulumi.Input[str],
                 is_auto_creation: pulumi.Input[bool],
                 is_auto_update: pulumi.Input[bool],
                 is_creation_allowed: pulumi.Input[bool],
                 is_linking_allowed: pulumi.Input[bool],
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a OrgIdpGoogle resource.
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        OrgIdpGoogleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            client_id=client_id,
            client_secret=client_secret,
            is_auto_creation=is_auto_creation,
            is_auto_update=is_auto_update,
            is_creation_allowed=is_creation_allowed,
            is_linking_allowed=is_linking_allowed,
            name=name,
            org_id=org_id,
            scopes=scopes,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             client_id: pulumi.Input[str],
             client_secret: pulumi.Input[str],
             is_auto_creation: pulumi.Input[bool],
             is_auto_update: pulumi.Input[bool],
             is_creation_allowed: pulumi.Input[bool],
             is_linking_allowed: pulumi.Input[bool],
             name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'clientId' in kwargs:
            client_id = kwargs['clientId']
        if 'clientSecret' in kwargs:
            client_secret = kwargs['clientSecret']
        if 'isAutoCreation' in kwargs:
            is_auto_creation = kwargs['isAutoCreation']
        if 'isAutoUpdate' in kwargs:
            is_auto_update = kwargs['isAutoUpdate']
        if 'isCreationAllowed' in kwargs:
            is_creation_allowed = kwargs['isCreationAllowed']
        if 'isLinkingAllowed' in kwargs:
            is_linking_allowed = kwargs['isLinkingAllowed']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        _setter("client_id", client_id)
        _setter("client_secret", client_secret)
        _setter("is_auto_creation", is_auto_creation)
        _setter("is_auto_update", is_auto_update)
        _setter("is_creation_allowed", is_creation_allowed)
        _setter("is_linking_allowed", is_linking_allowed)
        if name is not None:
            _setter("name", name)
        if org_id is not None:
            _setter("org_id", org_id)
        if scopes is not None:
            _setter("scopes", scopes)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Input[str]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> pulumi.Input[bool]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @is_auto_creation.setter
    def is_auto_creation(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_auto_creation", value)

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> pulumi.Input[bool]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @is_auto_update.setter
    def is_auto_update(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_auto_update", value)

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> pulumi.Input[bool]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @is_creation_allowed.setter
    def is_creation_allowed(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_creation_allowed", value)

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> pulumi.Input[bool]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @is_linking_allowed.setter
    def is_linking_allowed(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_linking_allowed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IDP
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
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)


@pulumi.input_type
class _OrgIdpGoogleState:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering OrgIdpGoogle resources.
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        _OrgIdpGoogleState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            client_id=client_id,
            client_secret=client_secret,
            is_auto_creation=is_auto_creation,
            is_auto_update=is_auto_update,
            is_creation_allowed=is_creation_allowed,
            is_linking_allowed=is_linking_allowed,
            name=name,
            org_id=org_id,
            scopes=scopes,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             client_id: Optional[pulumi.Input[str]] = None,
             client_secret: Optional[pulumi.Input[str]] = None,
             is_auto_creation: Optional[pulumi.Input[bool]] = None,
             is_auto_update: Optional[pulumi.Input[bool]] = None,
             is_creation_allowed: Optional[pulumi.Input[bool]] = None,
             is_linking_allowed: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'clientId' in kwargs:
            client_id = kwargs['clientId']
        if 'clientSecret' in kwargs:
            client_secret = kwargs['clientSecret']
        if 'isAutoCreation' in kwargs:
            is_auto_creation = kwargs['isAutoCreation']
        if 'isAutoUpdate' in kwargs:
            is_auto_update = kwargs['isAutoUpdate']
        if 'isCreationAllowed' in kwargs:
            is_creation_allowed = kwargs['isCreationAllowed']
        if 'isLinkingAllowed' in kwargs:
            is_linking_allowed = kwargs['isLinkingAllowed']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        if client_id is not None:
            _setter("client_id", client_id)
        if client_secret is not None:
            _setter("client_secret", client_secret)
        if is_auto_creation is not None:
            _setter("is_auto_creation", is_auto_creation)
        if is_auto_update is not None:
            _setter("is_auto_update", is_auto_update)
        if is_creation_allowed is not None:
            _setter("is_creation_allowed", is_creation_allowed)
        if is_linking_allowed is not None:
            _setter("is_linking_allowed", is_linking_allowed)
        if name is not None:
            _setter("name", name)
        if org_id is not None:
            _setter("org_id", org_id)
        if scopes is not None:
            _setter("scopes", scopes)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @is_auto_creation.setter
    def is_auto_creation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_auto_creation", value)

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @is_auto_update.setter
    def is_auto_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_auto_update", value)

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @is_creation_allowed.setter
    def is_creation_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_creation_allowed", value)

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @is_linking_allowed.setter
    def is_linking_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_linking_allowed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IDP
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
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)


class OrgIdpGoogle(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource representing a Google IdP on the organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.OrgIdpGoogle("default",
            org_id=default_zitadel_org["id"],
            name="Google",
            client_id="182902...",
            client_secret="GOCSPX-*****",
            scopes=[
                "openid",
                "profile",
                "email",
            ],
            is_linking_allowed=False,
            is_creation_allowed=True,
            is_auto_creation=False,
            is_auto_update=True)
        ```

        ## Import

        bash The resource can be imported using the ID format `<id[:org_id][:client_secret]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/orgIdpGoogle:OrgIdpGoogle imported '123456789012345678:123456789012345678:G1234567890123'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrgIdpGoogleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a Google IdP on the organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.OrgIdpGoogle("default",
            org_id=default_zitadel_org["id"],
            name="Google",
            client_id="182902...",
            client_secret="GOCSPX-*****",
            scopes=[
                "openid",
                "profile",
                "email",
            ],
            is_linking_allowed=False,
            is_creation_allowed=True,
            is_auto_creation=False,
            is_auto_update=True)
        ```

        ## Import

        bash The resource can be imported using the ID format `<id[:org_id][:client_secret]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/orgIdpGoogle:OrgIdpGoogle imported '123456789012345678:123456789012345678:G1234567890123'
        ```

        :param str resource_name: The name of the resource.
        :param OrgIdpGoogleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrgIdpGoogleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            OrgIdpGoogleArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrgIdpGoogleArgs.__new__(OrgIdpGoogleArgs)

            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            if client_secret is None and not opts.urn:
                raise TypeError("Missing required property 'client_secret'")
            __props__.__dict__["client_secret"] = None if client_secret is None else pulumi.Output.secret(client_secret)
            if is_auto_creation is None and not opts.urn:
                raise TypeError("Missing required property 'is_auto_creation'")
            __props__.__dict__["is_auto_creation"] = is_auto_creation
            if is_auto_update is None and not opts.urn:
                raise TypeError("Missing required property 'is_auto_update'")
            __props__.__dict__["is_auto_update"] = is_auto_update
            if is_creation_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'is_creation_allowed'")
            __props__.__dict__["is_creation_allowed"] = is_creation_allowed
            if is_linking_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'is_linking_allowed'")
            __props__.__dict__["is_linking_allowed"] = is_linking_allowed
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["scopes"] = scopes
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["clientSecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(OrgIdpGoogle, __self__).__init__(
            'zitadel:index/orgIdpGoogle:OrgIdpGoogle',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            is_auto_creation: Optional[pulumi.Input[bool]] = None,
            is_auto_update: Optional[pulumi.Input[bool]] = None,
            is_creation_allowed: Optional[pulumi.Input[bool]] = None,
            is_linking_allowed: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'OrgIdpGoogle':
        """
        Get an existing OrgIdpGoogle resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrgIdpGoogleState.__new__(_OrgIdpGoogleState)

        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["is_auto_creation"] = is_auto_creation
        __props__.__dict__["is_auto_update"] = is_auto_update
        __props__.__dict__["is_creation_allowed"] = is_creation_allowed
        __props__.__dict__["is_linking_allowed"] = is_linking_allowed
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["scopes"] = scopes
        return OrgIdpGoogle(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[str]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> pulumi.Output[bool]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> pulumi.Output[bool]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> pulumi.Output[bool]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> pulumi.Output[bool]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the IDP
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
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

