# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetNlbServiceListResult',
    'AwaitableGetNlbServiceListResult',
    'get_nlb_service_list',
    'get_nlb_service_list_output',
]

@pulumi.output_type
class GetNlbServiceListResult:
    """
    A collection of values returned by getNlbServiceList.
    """
    def __init__(__self__, id=None, nlb_id=None, nlb_name=None, services=None, timeouts=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if nlb_id and not isinstance(nlb_id, str):
            raise TypeError("Expected argument 'nlb_id' to be a str")
        pulumi.set(__self__, "nlb_id", nlb_id)
        if nlb_name and not isinstance(nlb_name, str):
            raise TypeError("Expected argument 'nlb_name' to be a str")
        pulumi.set(__self__, "nlb_name", nlb_name)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
        if timeouts and not isinstance(timeouts, dict):
            raise TypeError("Expected argument 'timeouts' to be a dict")
        pulumi.set(__self__, "timeouts", timeouts)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nlbId")
    def nlb_id(self) -> Optional[builtins.str]:
        """
        The NLB ID to match (conflicts with `name`).
        """
        return pulumi.get(self, "nlb_id")

    @property
    @pulumi.getter(name="nlbName")
    def nlb_name(self) -> Optional[builtins.str]:
        """
        The NLB name to match (conflicts with `id`).
        """
        return pulumi.get(self, "nlb_name")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetNlbServiceListServiceResult']:
        """
        The list of exoscale*nlb*service.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def timeouts(self) -> Optional['outputs.GetNlbServiceListTimeoutsResult']:
        return pulumi.get(self, "timeouts")

    @property
    @pulumi.getter
    def zone(self) -> builtins.str:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")


class AwaitableGetNlbServiceListResult(GetNlbServiceListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNlbServiceListResult(
            id=self.id,
            nlb_id=self.nlb_id,
            nlb_name=self.nlb_name,
            services=self.services,
            timeouts=self.timeouts,
            zone=self.zone)


def get_nlb_service_list(nlb_id: Optional[builtins.str] = None,
                         nlb_name: Optional[builtins.str] = None,
                         timeouts: Optional[Union['GetNlbServiceListTimeoutsArgs', 'GetNlbServiceListTimeoutsArgsDict']] = None,
                         zone: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNlbServiceListResult:
    """
    Fetch Exoscale [Network Load Balancers (NLB)](https://community.exoscale.com/product/networking/nlb/) Services.

    Corresponding resource: exoscale_nlb.


    :param builtins.str nlb_id: The NLB ID to match (conflicts with `name`).
    :param builtins.str nlb_name: The NLB name to match (conflicts with `id`).
    :param builtins.str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['nlbId'] = nlb_id
    __args__['nlbName'] = nlb_name
    __args__['timeouts'] = timeouts
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getNlbServiceList:getNlbServiceList', __args__, opts=opts, typ=GetNlbServiceListResult).value

    return AwaitableGetNlbServiceListResult(
        id=pulumi.get(__ret__, 'id'),
        nlb_id=pulumi.get(__ret__, 'nlb_id'),
        nlb_name=pulumi.get(__ret__, 'nlb_name'),
        services=pulumi.get(__ret__, 'services'),
        timeouts=pulumi.get(__ret__, 'timeouts'),
        zone=pulumi.get(__ret__, 'zone'))
def get_nlb_service_list_output(nlb_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                nlb_name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                timeouts: Optional[pulumi.Input[Optional[Union['GetNlbServiceListTimeoutsArgs', 'GetNlbServiceListTimeoutsArgsDict']]]] = None,
                                zone: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNlbServiceListResult]:
    """
    Fetch Exoscale [Network Load Balancers (NLB)](https://community.exoscale.com/product/networking/nlb/) Services.

    Corresponding resource: exoscale_nlb.


    :param builtins.str nlb_id: The NLB ID to match (conflicts with `name`).
    :param builtins.str nlb_name: The NLB name to match (conflicts with `id`).
    :param builtins.str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['nlbId'] = nlb_id
    __args__['nlbName'] = nlb_name
    __args__['timeouts'] = timeouts
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('exoscale:index/getNlbServiceList:getNlbServiceList', __args__, opts=opts, typ=GetNlbServiceListResult)
    return __ret__.apply(lambda __response__: GetNlbServiceListResult(
        id=pulumi.get(__response__, 'id'),
        nlb_id=pulumi.get(__response__, 'nlb_id'),
        nlb_name=pulumi.get(__response__, 'nlb_name'),
        services=pulumi.get(__response__, 'services'),
        timeouts=pulumi.get(__response__, 'timeouts'),
        zone=pulumi.get(__response__, 'zone')))
