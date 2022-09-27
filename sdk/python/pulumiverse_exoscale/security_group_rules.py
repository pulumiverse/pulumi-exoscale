# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SecurityGroupRulesArgs', 'SecurityGroupRules']

@pulumi.input_type
class SecurityGroupRulesArgs:
    def __init__(__self__, *,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesEgressArgs']]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesIngressArgs']]]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecurityGroupRules resource.
        :param pulumi.Input[str] security_group: The security group (name) the rules apply to (conflicts with `security_group_id`).
        :param pulumi.Input[str] security_group_id: The security group (ID) the rules apply to (conficts with `security_group)`.
        """
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesEgressArgs']]]]:
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesIngressArgs']]]]:
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[str]]:
        """
        The security group (name) the rules apply to (conflicts with `security_group_id`).
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The security group (ID) the rules apply to (conficts with `security_group)`.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)


@pulumi.input_type
class _SecurityGroupRulesState:
    def __init__(__self__, *,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesEgressArgs']]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesIngressArgs']]]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityGroupRules resources.
        :param pulumi.Input[str] security_group: The security group (name) the rules apply to (conflicts with `security_group_id`).
        :param pulumi.Input[str] security_group_id: The security group (ID) the rules apply to (conficts with `security_group)`.
        """
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesEgressArgs']]]]:
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesIngressArgs']]]]:
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupRulesIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[str]]:
        """
        The security group (name) the rules apply to (conflicts with `security_group_id`).
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The security group (ID) the rules apply to (conficts with `security_group)`.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)


class SecurityGroupRules(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupRulesEgressArgs']]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupRulesIngressArgs']]]]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use the SecurityGroupRule instead (or refer to the ad-hoc migration guide).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] security_group: The security group (name) the rules apply to (conflicts with `security_group_id`).
        :param pulumi.Input[str] security_group_id: The security group (ID) the rules apply to (conficts with `security_group)`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecurityGroupRulesArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use the SecurityGroupRule instead (or refer to the ad-hoc migration guide).

        :param str resource_name: The name of the resource.
        :param SecurityGroupRulesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityGroupRulesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupRulesEgressArgs']]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupRulesIngressArgs']]]]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityGroupRulesArgs.__new__(SecurityGroupRulesArgs)

            __props__.__dict__["egresses"] = egresses
            __props__.__dict__["ingresses"] = ingresses
            __props__.__dict__["security_group"] = security_group
            __props__.__dict__["security_group_id"] = security_group_id
        super(SecurityGroupRules, __self__).__init__(
            'exoscale:index/securityGroupRules:SecurityGroupRules',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupRulesEgressArgs']]]]] = None,
            ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupRulesIngressArgs']]]]] = None,
            security_group: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None) -> 'SecurityGroupRules':
        """
        Get an existing SecurityGroupRules resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] security_group: The security group (name) the rules apply to (conflicts with `security_group_id`).
        :param pulumi.Input[str] security_group_id: The security group (ID) the rules apply to (conficts with `security_group)`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityGroupRulesState.__new__(_SecurityGroupRulesState)

        __props__.__dict__["egresses"] = egresses
        __props__.__dict__["ingresses"] = ingresses
        __props__.__dict__["security_group"] = security_group
        __props__.__dict__["security_group_id"] = security_group_id
        return SecurityGroupRules(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def egresses(self) -> pulumi.Output[Optional[Sequence['outputs.SecurityGroupRulesEgress']]]:
        return pulumi.get(self, "egresses")

    @property
    @pulumi.getter
    def ingresses(self) -> pulumi.Output[Optional[Sequence['outputs.SecurityGroupRulesIngress']]]:
        return pulumi.get(self, "ingresses")

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> pulumi.Output[str]:
        """
        The security group (name) the rules apply to (conflicts with `security_group_id`).
        """
        return pulumi.get(self, "security_group")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The security group (ID) the rules apply to (conficts with `security_group)`.
        """
        return pulumi.get(self, "security_group_id")

