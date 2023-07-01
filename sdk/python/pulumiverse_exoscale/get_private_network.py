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
    'GetPrivateNetworkResult',
    'AwaitableGetPrivateNetworkResult',
    'get_private_network',
    'get_private_network_output',
]

@pulumi.output_type
class GetPrivateNetworkResult:
    """
    A collection of values returned by getPrivateNetwork.
    """
    def __init__(__self__, description=None, end_ip=None, id=None, name=None, netmask=None, start_ip=None, zone=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if end_ip and not isinstance(end_ip, str):
            raise TypeError("Expected argument 'end_ip' to be a str")
        pulumi.set(__self__, "end_ip", end_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if netmask and not isinstance(netmask, str):
            raise TypeError("Expected argument 'netmask' to be a str")
        pulumi.set(__self__, "netmask", netmask)
        if start_ip and not isinstance(start_ip, str):
            raise TypeError("Expected argument 'start_ip' to be a str")
        pulumi.set(__self__, "start_ip", start_ip)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The private network description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> str:
        """
        The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The private network ID to match (conflicts with `name`).
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The network name to match (conflicts with `id`).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netmask(self) -> str:
        """
        The network mask defining the IPv4 network allowed for static leases.
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> str:
        """
        The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        return pulumi.get(self, "start_ip")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")


class AwaitableGetPrivateNetworkResult(GetPrivateNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateNetworkResult(
            description=self.description,
            end_ip=self.end_ip,
            id=self.id,
            name=self.name,
            netmask=self.netmask,
            start_ip=self.start_ip,
            zone=self.zone)


def get_private_network(description: Optional[str] = None,
                        id: Optional[str] = None,
                        name: Optional[str] = None,
                        zone: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateNetworkResult:
    """
    Use this data source to access information about an existing resource.

    :param str description: The private network description.
    :param str id: The private network ID to match (conflicts with `name`).
    :param str name: The network name to match (conflicts with `id`).
    :param str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['id'] = id
    __args__['name'] = name
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getPrivateNetwork:getPrivateNetwork', __args__, opts=opts, typ=GetPrivateNetworkResult).value

    return AwaitableGetPrivateNetworkResult(
        description=pulumi.get(__ret__, 'description'),
        end_ip=pulumi.get(__ret__, 'end_ip'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        netmask=pulumi.get(__ret__, 'netmask'),
        start_ip=pulumi.get(__ret__, 'start_ip'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_private_network)
def get_private_network_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                               id: Optional[pulumi.Input[Optional[str]]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               zone: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivateNetworkResult]:
    """
    Use this data source to access information about an existing resource.

    :param str description: The private network description.
    :param str id: The private network ID to match (conflicts with `name`).
    :param str name: The network name to match (conflicts with `id`).
    :param str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    ...
