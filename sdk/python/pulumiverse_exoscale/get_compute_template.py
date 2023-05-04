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
    'GetComputeTemplateResult',
    'AwaitableGetComputeTemplateResult',
    'get_compute_template',
    'get_compute_template_output',
]

@pulumi.output_type
class GetComputeTemplateResult:
    """
    A collection of values returned by getComputeTemplate.
    """
    def __init__(__self__, filter=None, id=None, name=None, username=None, zone=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        """
        A template category filter (default: `featured`); among: - `featured` - official Exoscale templates - `community` - community-contributed templates - `mine` - custom templates private to my organization
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The compute instance template ID to match (conflicts with `name`).
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The template name to match (conflicts with `id`).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        Username for logging into a compute instance based on this template
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")


class AwaitableGetComputeTemplateResult(GetComputeTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeTemplateResult(
            filter=self.filter,
            id=self.id,
            name=self.name,
            username=self.username,
            zone=self.zone)


def get_compute_template(filter: Optional[str] = None,
                         id: Optional[str] = None,
                         name: Optional[str] = None,
                         zone: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeTemplateResult:
    """
    Use this data source to access information about an existing resource.

    :param str filter: A template category filter (default: `featured`); among: - `featured` - official Exoscale templates - `community` - community-contributed templates - `mine` - custom templates private to my organization
    :param str id: The compute instance template ID to match (conflicts with `name`).
    :param str name: The template name to match (conflicts with `id`).
    :param str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['id'] = id
    __args__['name'] = name
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getComputeTemplate:getComputeTemplate', __args__, opts=opts, typ=GetComputeTemplateResult).value

    return AwaitableGetComputeTemplateResult(
        filter=__ret__.filter,
        id=__ret__.id,
        name=__ret__.name,
        username=__ret__.username,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_compute_template)
def get_compute_template_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                                id: Optional[pulumi.Input[Optional[str]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                zone: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetComputeTemplateResult]:
    """
    Use this data source to access information about an existing resource.

    :param str filter: A template category filter (default: `featured`); among: - `featured` - official Exoscale templates - `community` - community-contributed templates - `mine` - custom templates private to my organization
    :param str id: The compute instance template ID to match (conflicts with `name`).
    :param str name: The template name to match (conflicts with `id`).
    :param str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    ...
