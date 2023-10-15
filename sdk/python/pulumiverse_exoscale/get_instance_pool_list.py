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

__all__ = [
    'GetInstancePoolListResult',
    'AwaitableGetInstancePoolListResult',
    'get_instance_pool_list',
    'get_instance_pool_list_output',
]

@pulumi.output_type
class GetInstancePoolListResult:
    """
    A collection of values returned by getInstancePoolList.
    """
    def __init__(__self__, id=None, pools=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if pools and not isinstance(pools, list):
            raise TypeError("Expected argument 'pools' to be a list")
        pulumi.set(__self__, "pools", pools)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def pools(self) -> Sequence['outputs.GetInstancePoolListPoolResult']:
        """
        The list of exoscale*instance*pool.
        """
        return pulumi.get(self, "pools")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")


class AwaitableGetInstancePoolListResult(GetInstancePoolListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancePoolListResult(
            id=self.id,
            pools=self.pools,
            zone=self.zone)


def get_instance_pool_list(zone: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancePoolListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getInstancePoolList:getInstancePoolList', __args__, opts=opts, typ=GetInstancePoolListResult).value

    return AwaitableGetInstancePoolListResult(
        id=pulumi.get(__ret__, 'id'),
        pools=pulumi.get(__ret__, 'pools'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_instance_pool_list)
def get_instance_pool_list_output(zone: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancePoolListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
