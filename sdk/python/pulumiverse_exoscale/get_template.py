# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetTemplateResult',
    'AwaitableGetTemplateResult',
    'get_template',
    'get_template_output',
]

@pulumi.output_type
class GetTemplateResult:
    """
    A collection of values returned by getTemplate.
    """
    def __init__(__self__, default_user=None, id=None, name=None, visibility=None, zone=None):
        if default_user and not isinstance(default_user, str):
            raise TypeError("Expected argument 'default_user' to be a str")
        pulumi.set(__self__, "default_user", default_user)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="defaultUser")
    def default_user(self) -> str:
        """
        Username to use to log into a compute instance based on this template
        """
        return pulumi.get(self, "default_user")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[str]:
        return pulumi.get(self, "visibility")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetTemplateResult(GetTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTemplateResult(
            default_user=self.default_user,
            id=self.id,
            name=self.name,
            visibility=self.visibility,
            zone=self.zone)


def get_template(id: Optional[str] = None,
                 name: Optional[str] = None,
                 visibility: Optional[str] = None,
                 zone: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTemplateResult:
    """
    Use this data source to access information about an existing resource.

    :param str id: The compute instance template ID to match (conflicts with `name`).
    :param str name: The template name to match (conflicts with `id`) (when multiple templates have the same name, the newest one will be returned).
    :param str visibility: A template category filter (default: `public`); among:
    :param str zone: The Exoscale [Zone][zone] name.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['visibility'] = visibility
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getTemplate:getTemplate', __args__, opts=opts, typ=GetTemplateResult).value

    return AwaitableGetTemplateResult(
        default_user=__ret__.default_user,
        id=__ret__.id,
        name=__ret__.name,
        visibility=__ret__.visibility,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_template)
def get_template_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        visibility: Optional[pulumi.Input[Optional[str]]] = None,
                        zone: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTemplateResult]:
    """
    Use this data source to access information about an existing resource.

    :param str id: The compute instance template ID to match (conflicts with `name`).
    :param str name: The template name to match (conflicts with `id`) (when multiple templates have the same name, the newest one will be returned).
    :param str visibility: A template category filter (default: `public`); among:
    :param str zone: The Exoscale [Zone][zone] name.
    """
    ...
