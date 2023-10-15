# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['IAMAPIKeyArgs', 'IAMAPIKey']

@pulumi.input_type
class IAMAPIKeyArgs:
    def __init__(__self__, *,
                 role_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input['IAMAPIKeyTimeoutsArgs']] = None):
        """
        The set of arguments for constructing a IAMAPIKey resource.
        :param pulumi.Input[str] role_id: ❗ IAM API role ID.
        :param pulumi.Input[str] name: ❗ IAM API Key name.
        """
        IAMAPIKeyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            role_id=role_id,
            name=name,
            timeouts=timeouts,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             role_id: pulumi.Input[str],
             name: Optional[pulumi.Input[str]] = None,
             timeouts: Optional[pulumi.Input['IAMAPIKeyTimeoutsArgs']] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("role_id", role_id)
        if name is not None:
            _setter("name", name)
        if timeouts is not None:
            _setter("timeouts", timeouts)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[str]:
        """
        ❗ IAM API role ID.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ IAM API Key name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['IAMAPIKeyTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['IAMAPIKeyTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


@pulumi.input_type
class _IAMAPIKeyState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input['IAMAPIKeyTimeoutsArgs']] = None):
        """
        Input properties used for looking up and filtering IAMAPIKey resources.
        :param pulumi.Input[str] key: The IAM API Key to match.
        :param pulumi.Input[str] name: ❗ IAM API Key name.
        :param pulumi.Input[str] role_id: ❗ IAM API role ID.
        :param pulumi.Input[str] secret: Secret for the IAM API Key.
        """
        _IAMAPIKeyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            key=key,
            name=name,
            role_id=role_id,
            secret=secret,
            timeouts=timeouts,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             key: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             role_id: Optional[pulumi.Input[str]] = None,
             secret: Optional[pulumi.Input[str]] = None,
             timeouts: Optional[pulumi.Input['IAMAPIKeyTimeoutsArgs']] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if key is not None:
            _setter("key", key)
        if name is not None:
            _setter("name", name)
        if role_id is not None:
            _setter("role_id", role_id)
        if secret is not None:
            _setter("secret", secret)
        if timeouts is not None:
            _setter("timeouts", timeouts)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The IAM API Key to match.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ IAM API Key name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ IAM API role ID.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        Secret for the IAM API Key.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['IAMAPIKeyTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['IAMAPIKeyTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


class IAMAPIKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input[pulumi.InputType['IAMAPIKeyTimeoutsArgs']]] = None,
                 __props__=None):
        """
        Create a IAMAPIKey resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: ❗ IAM API Key name.
        :param pulumi.Input[str] role_id: ❗ IAM API role ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IAMAPIKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IAMAPIKey resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IAMAPIKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IAMAPIKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IAMAPIKeyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input[pulumi.InputType['IAMAPIKeyTimeoutsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IAMAPIKeyArgs.__new__(IAMAPIKeyArgs)

            __props__.__dict__["name"] = name
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
            if timeouts is not None and not isinstance(timeouts, IAMAPIKeyTimeoutsArgs):
                timeouts = timeouts or {}
                def _setter(key, value):
                    timeouts[key] = value
                IAMAPIKeyTimeoutsArgs._configure(_setter, **timeouts)
            __props__.__dict__["timeouts"] = timeouts
            __props__.__dict__["key"] = None
            __props__.__dict__["secret"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(IAMAPIKey, __self__).__init__(
            'exoscale:index/iAMAPIKey:IAMAPIKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            role_id: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            timeouts: Optional[pulumi.Input[pulumi.InputType['IAMAPIKeyTimeoutsArgs']]] = None) -> 'IAMAPIKey':
        """
        Get an existing IAMAPIKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: The IAM API Key to match.
        :param pulumi.Input[str] name: ❗ IAM API Key name.
        :param pulumi.Input[str] role_id: ❗ IAM API role ID.
        :param pulumi.Input[str] secret: Secret for the IAM API Key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMAPIKeyState.__new__(_IAMAPIKeyState)

        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["role_id"] = role_id
        __props__.__dict__["secret"] = secret
        __props__.__dict__["timeouts"] = timeouts
        return IAMAPIKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The IAM API Key to match.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        ❗ IAM API Key name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        ❗ IAM API role ID.
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        Secret for the IAM API Key.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def timeouts(self) -> pulumi.Output[Optional['outputs.IAMAPIKeyTimeouts']]:
        return pulumi.get(self, "timeouts")

