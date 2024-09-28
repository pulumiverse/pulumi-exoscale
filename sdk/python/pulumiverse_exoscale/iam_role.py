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

__all__ = ['IamRoleArgs', 'IamRole']

@pulumi.input_type
class IamRoleArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 editable: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy: Optional[pulumi.Input['IamRolePolicyArgs']] = None,
                 timeouts: Optional[pulumi.Input['IamRoleTimeoutsArgs']] = None):
        """
        The set of arguments for constructing a IamRole resource.
        :param pulumi.Input[str] description: A free-form text describing the IAM Role
        :param pulumi.Input[bool] editable: Defines if IAM Role Policy is editable or not.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: IAM Role labels.
        :param pulumi.Input[str] name: ❗Name of IAM Role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: IAM Role permissions.
        :param pulumi.Input['IamRolePolicyArgs'] policy: IAM Policy.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if editable is not None:
            pulumi.set(__self__, "editable", editable)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the IAM Role
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def editable(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if IAM Role Policy is editable or not.
        """
        return pulumi.get(self, "editable")

    @editable.setter
    def editable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "editable", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        IAM Role labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ❗Name of IAM Role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IAM Role permissions.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input['IamRolePolicyArgs']]:
        """
        IAM Policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input['IamRolePolicyArgs']]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['IamRoleTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['IamRoleTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


@pulumi.input_type
class _IamRoleState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 editable: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy: Optional[pulumi.Input['IamRolePolicyArgs']] = None,
                 timeouts: Optional[pulumi.Input['IamRoleTimeoutsArgs']] = None):
        """
        Input properties used for looking up and filtering IamRole resources.
        :param pulumi.Input[str] description: A free-form text describing the IAM Role
        :param pulumi.Input[bool] editable: Defines if IAM Role Policy is editable or not.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: IAM Role labels.
        :param pulumi.Input[str] name: ❗Name of IAM Role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: IAM Role permissions.
        :param pulumi.Input['IamRolePolicyArgs'] policy: IAM Policy.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if editable is not None:
            pulumi.set(__self__, "editable", editable)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the IAM Role
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def editable(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if IAM Role Policy is editable or not.
        """
        return pulumi.get(self, "editable")

    @editable.setter
    def editable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "editable", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        IAM Role labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ❗Name of IAM Role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IAM Role permissions.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input['IamRolePolicyArgs']]:
        """
        IAM Policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input['IamRolePolicyArgs']]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['IamRoleTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['IamRoleTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


class IamRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 editable: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy: Optional[pulumi.Input[Union['IamRolePolicyArgs', 'IamRolePolicyArgsDict']]] = None,
                 timeouts: Optional[pulumi.Input[Union['IamRoleTimeoutsArgs', 'IamRoleTimeoutsArgsDict']]] = None,
                 __props__=None):
        """
        Manage Exoscale [IAM](https://community.exoscale.com/documentation/iam/) Role.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A free-form text describing the IAM Role
        :param pulumi.Input[bool] editable: Defines if IAM Role Policy is editable or not.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: IAM Role labels.
        :param pulumi.Input[str] name: ❗Name of IAM Role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: IAM Role permissions.
        :param pulumi.Input[Union['IamRolePolicyArgs', 'IamRolePolicyArgsDict']] policy: IAM Policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IamRoleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage Exoscale [IAM](https://community.exoscale.com/documentation/iam/) Role.

        :param str resource_name: The name of the resource.
        :param IamRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 editable: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy: Optional[pulumi.Input[Union['IamRolePolicyArgs', 'IamRolePolicyArgsDict']]] = None,
                 timeouts: Optional[pulumi.Input[Union['IamRoleTimeoutsArgs', 'IamRoleTimeoutsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamRoleArgs.__new__(IamRoleArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["editable"] = editable
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["policy"] = policy
            __props__.__dict__["timeouts"] = timeouts
        super(IamRole, __self__).__init__(
            'exoscale:index/iamRole:IamRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            editable: Optional[pulumi.Input[bool]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            policy: Optional[pulumi.Input[Union['IamRolePolicyArgs', 'IamRolePolicyArgsDict']]] = None,
            timeouts: Optional[pulumi.Input[Union['IamRoleTimeoutsArgs', 'IamRoleTimeoutsArgsDict']]] = None) -> 'IamRole':
        """
        Get an existing IamRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A free-form text describing the IAM Role
        :param pulumi.Input[bool] editable: Defines if IAM Role Policy is editable or not.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: IAM Role labels.
        :param pulumi.Input[str] name: ❗Name of IAM Role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: IAM Role permissions.
        :param pulumi.Input[Union['IamRolePolicyArgs', 'IamRolePolicyArgsDict']] policy: IAM Policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamRoleState.__new__(_IamRoleState)

        __props__.__dict__["description"] = description
        __props__.__dict__["editable"] = editable
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["policy"] = policy
        __props__.__dict__["timeouts"] = timeouts
        return IamRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A free-form text describing the IAM Role
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def editable(self) -> pulumi.Output[bool]:
        """
        Defines if IAM Role Policy is editable or not.
        """
        return pulumi.get(self, "editable")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        IAM Role labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        ❗Name of IAM Role.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence[str]]:
        """
        IAM Role permissions.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output['outputs.IamRolePolicy']:
        """
        IAM Policy.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter
    def timeouts(self) -> pulumi.Output[Optional['outputs.IamRoleTimeouts']]:
        return pulumi.get(self, "timeouts")

