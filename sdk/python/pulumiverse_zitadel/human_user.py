# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['HumanUserArgs', 'HumanUser']

@pulumi.input_type
class HumanUserArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 first_name: pulumi.Input[str],
                 last_name: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 initial_password: Optional[pulumi.Input[str]] = None,
                 is_email_verified: Optional[pulumi.Input[bool]] = None,
                 is_phone_verified: Optional[pulumi.Input[bool]] = None,
                 nick_name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 preferred_language: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HumanUser resource.
        :param pulumi.Input[str] email: Email of the user
        :param pulumi.Input[str] first_name: First name of the user
        :param pulumi.Input[str] last_name: Last name of the user
        :param pulumi.Input[str] user_name: Username
        :param pulumi.Input[str] display_name: Display name of the user
        :param pulumi.Input[str] gender: Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        :param pulumi.Input[str] initial_password: Initially set password for the user, not changeable after creation
        :param pulumi.Input[bool] is_email_verified: Is the email verified of the user, can only be true if password of the user is set
        :param pulumi.Input[bool] is_phone_verified: Is the phone verified of the user
        :param pulumi.Input[str] nick_name: Nick name of the user
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] phone: Phone of the user
        :param pulumi.Input[str] preferred_language: Preferred language of the user
        """
        HumanUserArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            email=email,
            first_name=first_name,
            last_name=last_name,
            user_name=user_name,
            display_name=display_name,
            gender=gender,
            initial_password=initial_password,
            is_email_verified=is_email_verified,
            is_phone_verified=is_phone_verified,
            nick_name=nick_name,
            org_id=org_id,
            phone=phone,
            preferred_language=preferred_language,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             email: pulumi.Input[str],
             first_name: pulumi.Input[str],
             last_name: pulumi.Input[str],
             user_name: pulumi.Input[str],
             display_name: Optional[pulumi.Input[str]] = None,
             gender: Optional[pulumi.Input[str]] = None,
             initial_password: Optional[pulumi.Input[str]] = None,
             is_email_verified: Optional[pulumi.Input[bool]] = None,
             is_phone_verified: Optional[pulumi.Input[bool]] = None,
             nick_name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             phone: Optional[pulumi.Input[str]] = None,
             preferred_language: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'firstName' in kwargs:
            first_name = kwargs['firstName']
        if 'lastName' in kwargs:
            last_name = kwargs['lastName']
        if 'userName' in kwargs:
            user_name = kwargs['userName']
        if 'displayName' in kwargs:
            display_name = kwargs['displayName']
        if 'initialPassword' in kwargs:
            initial_password = kwargs['initialPassword']
        if 'isEmailVerified' in kwargs:
            is_email_verified = kwargs['isEmailVerified']
        if 'isPhoneVerified' in kwargs:
            is_phone_verified = kwargs['isPhoneVerified']
        if 'nickName' in kwargs:
            nick_name = kwargs['nickName']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'preferredLanguage' in kwargs:
            preferred_language = kwargs['preferredLanguage']

        _setter("email", email)
        _setter("first_name", first_name)
        _setter("last_name", last_name)
        _setter("user_name", user_name)
        if display_name is not None:
            _setter("display_name", display_name)
        if gender is not None:
            _setter("gender", gender)
        if initial_password is not None:
            _setter("initial_password", initial_password)
        if is_email_verified is not None:
            _setter("is_email_verified", is_email_verified)
        if is_phone_verified is not None:
            _setter("is_phone_verified", is_phone_verified)
        if nick_name is not None:
            _setter("nick_name", nick_name)
        if org_id is not None:
            _setter("org_id", org_id)
        if phone is not None:
            _setter("phone", phone)
        if preferred_language is not None:
            _setter("preferred_language", preferred_language)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        Email of the user
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> pulumi.Input[str]:
        """
        First name of the user
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> pulumi.Input[str]:
        """
        Last name of the user
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        Username
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the user
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def gender(self) -> Optional[pulumi.Input[str]]:
        """
        Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        """
        return pulumi.get(self, "gender")

    @gender.setter
    def gender(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gender", value)

    @property
    @pulumi.getter(name="initialPassword")
    def initial_password(self) -> Optional[pulumi.Input[str]]:
        """
        Initially set password for the user, not changeable after creation
        """
        return pulumi.get(self, "initial_password")

    @initial_password.setter
    def initial_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initial_password", value)

    @property
    @pulumi.getter(name="isEmailVerified")
    def is_email_verified(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the email verified of the user, can only be true if password of the user is set
        """
        return pulumi.get(self, "is_email_verified")

    @is_email_verified.setter
    def is_email_verified(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_email_verified", value)

    @property
    @pulumi.getter(name="isPhoneVerified")
    def is_phone_verified(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the phone verified of the user
        """
        return pulumi.get(self, "is_phone_verified")

    @is_phone_verified.setter
    def is_phone_verified(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_phone_verified", value)

    @property
    @pulumi.getter(name="nickName")
    def nick_name(self) -> Optional[pulumi.Input[str]]:
        """
        Nick name of the user
        """
        return pulumi.get(self, "nick_name")

    @nick_name.setter
    def nick_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nick_name", value)

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
    def phone(self) -> Optional[pulumi.Input[str]]:
        """
        Phone of the user
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter(name="preferredLanguage")
    def preferred_language(self) -> Optional[pulumi.Input[str]]:
        """
        Preferred language of the user
        """
        return pulumi.get(self, "preferred_language")

    @preferred_language.setter
    def preferred_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_language", value)


@pulumi.input_type
class _HumanUserState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 initial_password: Optional[pulumi.Input[str]] = None,
                 is_email_verified: Optional[pulumi.Input[bool]] = None,
                 is_phone_verified: Optional[pulumi.Input[bool]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 login_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 nick_name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 preferred_language: Optional[pulumi.Input[str]] = None,
                 preferred_login_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HumanUser resources.
        :param pulumi.Input[str] display_name: Display name of the user
        :param pulumi.Input[str] email: Email of the user
        :param pulumi.Input[str] first_name: First name of the user
        :param pulumi.Input[str] gender: Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        :param pulumi.Input[str] initial_password: Initially set password for the user, not changeable after creation
        :param pulumi.Input[bool] is_email_verified: Is the email verified of the user, can only be true if password of the user is set
        :param pulumi.Input[bool] is_phone_verified: Is the phone verified of the user
        :param pulumi.Input[str] last_name: Last name of the user
        :param pulumi.Input[Sequence[pulumi.Input[str]]] login_names: Loginnames
        :param pulumi.Input[str] nick_name: Nick name of the user
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] phone: Phone of the user
        :param pulumi.Input[str] preferred_language: Preferred language of the user
        :param pulumi.Input[str] preferred_login_name: Preferred login name
        :param pulumi.Input[str] state: State of the user
        :param pulumi.Input[str] user_name: Username
        """
        _HumanUserState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            display_name=display_name,
            email=email,
            first_name=first_name,
            gender=gender,
            initial_password=initial_password,
            is_email_verified=is_email_verified,
            is_phone_verified=is_phone_verified,
            last_name=last_name,
            login_names=login_names,
            nick_name=nick_name,
            org_id=org_id,
            phone=phone,
            preferred_language=preferred_language,
            preferred_login_name=preferred_login_name,
            state=state,
            user_name=user_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             display_name: Optional[pulumi.Input[str]] = None,
             email: Optional[pulumi.Input[str]] = None,
             first_name: Optional[pulumi.Input[str]] = None,
             gender: Optional[pulumi.Input[str]] = None,
             initial_password: Optional[pulumi.Input[str]] = None,
             is_email_verified: Optional[pulumi.Input[bool]] = None,
             is_phone_verified: Optional[pulumi.Input[bool]] = None,
             last_name: Optional[pulumi.Input[str]] = None,
             login_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             nick_name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             phone: Optional[pulumi.Input[str]] = None,
             preferred_language: Optional[pulumi.Input[str]] = None,
             preferred_login_name: Optional[pulumi.Input[str]] = None,
             state: Optional[pulumi.Input[str]] = None,
             user_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'displayName' in kwargs:
            display_name = kwargs['displayName']
        if 'firstName' in kwargs:
            first_name = kwargs['firstName']
        if 'initialPassword' in kwargs:
            initial_password = kwargs['initialPassword']
        if 'isEmailVerified' in kwargs:
            is_email_verified = kwargs['isEmailVerified']
        if 'isPhoneVerified' in kwargs:
            is_phone_verified = kwargs['isPhoneVerified']
        if 'lastName' in kwargs:
            last_name = kwargs['lastName']
        if 'loginNames' in kwargs:
            login_names = kwargs['loginNames']
        if 'nickName' in kwargs:
            nick_name = kwargs['nickName']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'preferredLanguage' in kwargs:
            preferred_language = kwargs['preferredLanguage']
        if 'preferredLoginName' in kwargs:
            preferred_login_name = kwargs['preferredLoginName']
        if 'userName' in kwargs:
            user_name = kwargs['userName']

        if display_name is not None:
            _setter("display_name", display_name)
        if email is not None:
            _setter("email", email)
        if first_name is not None:
            _setter("first_name", first_name)
        if gender is not None:
            _setter("gender", gender)
        if initial_password is not None:
            _setter("initial_password", initial_password)
        if is_email_verified is not None:
            _setter("is_email_verified", is_email_verified)
        if is_phone_verified is not None:
            _setter("is_phone_verified", is_phone_verified)
        if last_name is not None:
            _setter("last_name", last_name)
        if login_names is not None:
            _setter("login_names", login_names)
        if nick_name is not None:
            _setter("nick_name", nick_name)
        if org_id is not None:
            _setter("org_id", org_id)
        if phone is not None:
            _setter("phone", phone)
        if preferred_language is not None:
            _setter("preferred_language", preferred_language)
        if preferred_login_name is not None:
            _setter("preferred_login_name", preferred_login_name)
        if state is not None:
            _setter("state", state)
        if user_name is not None:
            _setter("user_name", user_name)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the user
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the user
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        First name of the user
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter
    def gender(self) -> Optional[pulumi.Input[str]]:
        """
        Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        """
        return pulumi.get(self, "gender")

    @gender.setter
    def gender(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gender", value)

    @property
    @pulumi.getter(name="initialPassword")
    def initial_password(self) -> Optional[pulumi.Input[str]]:
        """
        Initially set password for the user, not changeable after creation
        """
        return pulumi.get(self, "initial_password")

    @initial_password.setter
    def initial_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initial_password", value)

    @property
    @pulumi.getter(name="isEmailVerified")
    def is_email_verified(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the email verified of the user, can only be true if password of the user is set
        """
        return pulumi.get(self, "is_email_verified")

    @is_email_verified.setter
    def is_email_verified(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_email_verified", value)

    @property
    @pulumi.getter(name="isPhoneVerified")
    def is_phone_verified(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the phone verified of the user
        """
        return pulumi.get(self, "is_phone_verified")

    @is_phone_verified.setter
    def is_phone_verified(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_phone_verified", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        Last name of the user
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter(name="loginNames")
    def login_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Loginnames
        """
        return pulumi.get(self, "login_names")

    @login_names.setter
    def login_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "login_names", value)

    @property
    @pulumi.getter(name="nickName")
    def nick_name(self) -> Optional[pulumi.Input[str]]:
        """
        Nick name of the user
        """
        return pulumi.get(self, "nick_name")

    @nick_name.setter
    def nick_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nick_name", value)

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
    def phone(self) -> Optional[pulumi.Input[str]]:
        """
        Phone of the user
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter(name="preferredLanguage")
    def preferred_language(self) -> Optional[pulumi.Input[str]]:
        """
        Preferred language of the user
        """
        return pulumi.get(self, "preferred_language")

    @preferred_language.setter
    def preferred_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_language", value)

    @property
    @pulumi.getter(name="preferredLoginName")
    def preferred_login_name(self) -> Optional[pulumi.Input[str]]:
        """
        Preferred login name
        """
        return pulumi.get(self, "preferred_login_name")

    @preferred_login_name.setter
    def preferred_login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_login_name", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the user
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Username
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class HumanUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 initial_password: Optional[pulumi.Input[str]] = None,
                 is_email_verified: Optional[pulumi.Input[bool]] = None,
                 is_phone_verified: Optional[pulumi.Input[bool]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 nick_name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 preferred_language: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Caution: Email can only be set verified if a password is set for the user, either with initial_password or during runtime**

        Resource representing a human user situated under an organization, which then can be authorized through memberships or direct grants on other resources.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.HumanUser("default",
            org_id=default_zitadel_org["id"],
            user_name="humanfull@localhost.com",
            first_name="firstname",
            last_name="lastname",
            nick_name="nickname",
            display_name="displayname",
            preferred_language="de",
            gender="GENDER_MALE",
            phone="+41799999999",
            is_phone_verified=True,
            email="test@zitadel.com",
            is_email_verified=True,
            initial_password="Password1!")
        ```

        ## Import

        bash The resource can be imported using the ID format `id[:org_id][:initial_password]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/humanUser:HumanUser imported '123456789012345678:123456789012345678:Password1!'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Display name of the user
        :param pulumi.Input[str] email: Email of the user
        :param pulumi.Input[str] first_name: First name of the user
        :param pulumi.Input[str] gender: Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        :param pulumi.Input[str] initial_password: Initially set password for the user, not changeable after creation
        :param pulumi.Input[bool] is_email_verified: Is the email verified of the user, can only be true if password of the user is set
        :param pulumi.Input[bool] is_phone_verified: Is the phone verified of the user
        :param pulumi.Input[str] last_name: Last name of the user
        :param pulumi.Input[str] nick_name: Nick name of the user
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] phone: Phone of the user
        :param pulumi.Input[str] preferred_language: Preferred language of the user
        :param pulumi.Input[str] user_name: Username
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HumanUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Caution: Email can only be set verified if a password is set for the user, either with initial_password or during runtime**

        Resource representing a human user situated under an organization, which then can be authorized through memberships or direct grants on other resources.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.HumanUser("default",
            org_id=default_zitadel_org["id"],
            user_name="humanfull@localhost.com",
            first_name="firstname",
            last_name="lastname",
            nick_name="nickname",
            display_name="displayname",
            preferred_language="de",
            gender="GENDER_MALE",
            phone="+41799999999",
            is_phone_verified=True,
            email="test@zitadel.com",
            is_email_verified=True,
            initial_password="Password1!")
        ```

        ## Import

        bash The resource can be imported using the ID format `id[:org_id][:initial_password]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/humanUser:HumanUser imported '123456789012345678:123456789012345678:Password1!'
        ```

        :param str resource_name: The name of the resource.
        :param HumanUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HumanUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            HumanUserArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 initial_password: Optional[pulumi.Input[str]] = None,
                 is_email_verified: Optional[pulumi.Input[bool]] = None,
                 is_phone_verified: Optional[pulumi.Input[bool]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 nick_name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 preferred_language: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HumanUserArgs.__new__(HumanUserArgs)

            __props__.__dict__["display_name"] = display_name
            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            if first_name is None and not opts.urn:
                raise TypeError("Missing required property 'first_name'")
            __props__.__dict__["first_name"] = first_name
            __props__.__dict__["gender"] = gender
            __props__.__dict__["initial_password"] = None if initial_password is None else pulumi.Output.secret(initial_password)
            __props__.__dict__["is_email_verified"] = is_email_verified
            __props__.__dict__["is_phone_verified"] = is_phone_verified
            if last_name is None and not opts.urn:
                raise TypeError("Missing required property 'last_name'")
            __props__.__dict__["last_name"] = last_name
            __props__.__dict__["nick_name"] = nick_name
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["phone"] = phone
            __props__.__dict__["preferred_language"] = preferred_language
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["login_names"] = None
            __props__.__dict__["preferred_login_name"] = None
            __props__.__dict__["state"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["initialPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(HumanUser, __self__).__init__(
            'zitadel:index/humanUser:HumanUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            first_name: Optional[pulumi.Input[str]] = None,
            gender: Optional[pulumi.Input[str]] = None,
            initial_password: Optional[pulumi.Input[str]] = None,
            is_email_verified: Optional[pulumi.Input[bool]] = None,
            is_phone_verified: Optional[pulumi.Input[bool]] = None,
            last_name: Optional[pulumi.Input[str]] = None,
            login_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            nick_name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            phone: Optional[pulumi.Input[str]] = None,
            preferred_language: Optional[pulumi.Input[str]] = None,
            preferred_login_name: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'HumanUser':
        """
        Get an existing HumanUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Display name of the user
        :param pulumi.Input[str] email: Email of the user
        :param pulumi.Input[str] first_name: First name of the user
        :param pulumi.Input[str] gender: Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        :param pulumi.Input[str] initial_password: Initially set password for the user, not changeable after creation
        :param pulumi.Input[bool] is_email_verified: Is the email verified of the user, can only be true if password of the user is set
        :param pulumi.Input[bool] is_phone_verified: Is the phone verified of the user
        :param pulumi.Input[str] last_name: Last name of the user
        :param pulumi.Input[Sequence[pulumi.Input[str]]] login_names: Loginnames
        :param pulumi.Input[str] nick_name: Nick name of the user
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] phone: Phone of the user
        :param pulumi.Input[str] preferred_language: Preferred language of the user
        :param pulumi.Input[str] preferred_login_name: Preferred login name
        :param pulumi.Input[str] state: State of the user
        :param pulumi.Input[str] user_name: Username
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HumanUserState.__new__(_HumanUserState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["email"] = email
        __props__.__dict__["first_name"] = first_name
        __props__.__dict__["gender"] = gender
        __props__.__dict__["initial_password"] = initial_password
        __props__.__dict__["is_email_verified"] = is_email_verified
        __props__.__dict__["is_phone_verified"] = is_phone_verified
        __props__.__dict__["last_name"] = last_name
        __props__.__dict__["login_names"] = login_names
        __props__.__dict__["nick_name"] = nick_name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["phone"] = phone
        __props__.__dict__["preferred_language"] = preferred_language
        __props__.__dict__["preferred_login_name"] = preferred_login_name
        __props__.__dict__["state"] = state
        __props__.__dict__["user_name"] = user_name
        return HumanUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name of the user
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        Email of the user
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> pulumi.Output[str]:
        """
        First name of the user
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter
    def gender(self) -> pulumi.Output[Optional[str]]:
        """
        Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        """
        return pulumi.get(self, "gender")

    @property
    @pulumi.getter(name="initialPassword")
    def initial_password(self) -> pulumi.Output[Optional[str]]:
        """
        Initially set password for the user, not changeable after creation
        """
        return pulumi.get(self, "initial_password")

    @property
    @pulumi.getter(name="isEmailVerified")
    def is_email_verified(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the email verified of the user, can only be true if password of the user is set
        """
        return pulumi.get(self, "is_email_verified")

    @property
    @pulumi.getter(name="isPhoneVerified")
    def is_phone_verified(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the phone verified of the user
        """
        return pulumi.get(self, "is_phone_verified")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> pulumi.Output[str]:
        """
        Last name of the user
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter(name="loginNames")
    def login_names(self) -> pulumi.Output[Sequence[str]]:
        """
        Loginnames
        """
        return pulumi.get(self, "login_names")

    @property
    @pulumi.getter(name="nickName")
    def nick_name(self) -> pulumi.Output[Optional[str]]:
        """
        Nick name of the user
        """
        return pulumi.get(self, "nick_name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Output[Optional[str]]:
        """
        Phone of the user
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter(name="preferredLanguage")
    def preferred_language(self) -> pulumi.Output[Optional[str]]:
        """
        Preferred language of the user
        """
        return pulumi.get(self, "preferred_language")

    @property
    @pulumi.getter(name="preferredLoginName")
    def preferred_login_name(self) -> pulumi.Output[str]:
        """
        Preferred login name
        """
        return pulumi.get(self, "preferred_login_name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the user
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Username
        """
        return pulumi.get(self, "user_name")

