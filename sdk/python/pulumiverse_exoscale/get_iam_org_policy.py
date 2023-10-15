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

__all__ = [
    'GetIAMOrgPolicyResult',
    'AwaitableGetIAMOrgPolicyResult',
    'get_iam_org_policy',
    'get_iam_org_policy_output',
]

@pulumi.output_type
class GetIAMOrgPolicyResult:
    """
    A collection of values returned by getIAMOrgPolicy.
    """
    def __init__(__self__, default_service_strategy=None, id=None, services=None, timeouts=None):
        if default_service_strategy and not isinstance(default_service_strategy, str):
            raise TypeError("Expected argument 'default_service_strategy' to be a str")
        pulumi.set(__self__, "default_service_strategy", default_service_strategy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if services and not isinstance(services, dict):
            raise TypeError("Expected argument 'services' to be a dict")
        pulumi.set(__self__, "services", services)
        if timeouts and not isinstance(timeouts, dict):
            raise TypeError("Expected argument 'timeouts' to be a dict")
        pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter(name="defaultServiceStrategy")
    def default_service_strategy(self) -> str:
        """
        Default service strategy (`allow` or `deny`).
        """
        return pulumi.get(self, "default_service_strategy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def services(self) -> Mapping[str, 'outputs.GetIAMOrgPolicyServicesResult']:
        """
        IAM policy services.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def timeouts(self) -> Optional['outputs.GetIAMOrgPolicyTimeoutsResult']:
        return pulumi.get(self, "timeouts")


class AwaitableGetIAMOrgPolicyResult(GetIAMOrgPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIAMOrgPolicyResult(
            default_service_strategy=self.default_service_strategy,
            id=self.id,
            services=self.services,
            timeouts=self.timeouts)


def get_iam_org_policy(timeouts: Optional[pulumi.InputType['GetIAMOrgPolicyTimeoutsArgs']] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIAMOrgPolicyResult:
    """
    Fetch Exoscale [IAM](https://community.exoscale.com/documentation/iam/) Organization Policy.

    Corresponding resource: exoscale_iam_org_policy.
    """
    __args__ = dict()
    __args__['timeouts'] = timeouts
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getIAMOrgPolicy:getIAMOrgPolicy', __args__, opts=opts, typ=GetIAMOrgPolicyResult).value

    return AwaitableGetIAMOrgPolicyResult(
        default_service_strategy=pulumi.get(__ret__, 'default_service_strategy'),
        id=pulumi.get(__ret__, 'id'),
        services=pulumi.get(__ret__, 'services'),
        timeouts=pulumi.get(__ret__, 'timeouts'))


@_utilities.lift_output_func(get_iam_org_policy)
def get_iam_org_policy_output(timeouts: Optional[pulumi.Input[Optional[pulumi.InputType['GetIAMOrgPolicyTimeoutsArgs']]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIAMOrgPolicyResult]:
    """
    Fetch Exoscale [IAM](https://community.exoscale.com/documentation/iam/) Organization Policy.

    Corresponding resource: exoscale_iam_org_policy.
    """
    ...
