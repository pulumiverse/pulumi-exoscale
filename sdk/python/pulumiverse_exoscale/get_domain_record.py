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
    'GetDomainRecordResult',
    'AwaitableGetDomainRecordResult',
    'get_domain_record',
    'get_domain_record_output',
]

@pulumi.output_type
class GetDomainRecordResult:
    """
    A collection of values returned by getDomainRecord.
    """
    def __init__(__self__, domain=None, filter=None, id=None, records=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if filter and not isinstance(filter, dict):
            raise TypeError("Expected argument 'filter' to be a dict")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if records and not isinstance(records, list):
            raise TypeError("Expected argument 'records' to be a list")
        pulumi.set(__self__, "records", records)

    @property
    @pulumi.getter
    def domain(self) -> builtins.str:
        """
        The Domain name to match.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def filter(self) -> 'outputs.GetDomainRecordFilterResult':
        """
        Filter to apply when looking up domain records.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def records(self) -> Sequence['outputs.GetDomainRecordRecordResult']:
        """
        The list of matching records. Structure is documented below.
        """
        return pulumi.get(self, "records")


class AwaitableGetDomainRecordResult(GetDomainRecordResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainRecordResult(
            domain=self.domain,
            filter=self.filter,
            id=self.id,
            records=self.records)


def get_domain_record(domain: Optional[builtins.str] = None,
                      filter: Optional[Union['GetDomainRecordFilterArgs', 'GetDomainRecordFilterArgsDict']] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainRecordResult:
    """
    Fetch Exoscale [DNS](https://community.exoscale.com/product/networking/dns/) Domain Records data.

    Corresponding resource: exoscale_domain_record.


    :param builtins.str domain: The Domain name to match.
    :param Union['GetDomainRecordFilterArgs', 'GetDomainRecordFilterArgsDict'] filter: Filter to apply when looking up domain records.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['filter'] = filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getDomainRecord:getDomainRecord', __args__, opts=opts, typ=GetDomainRecordResult).value

    return AwaitableGetDomainRecordResult(
        domain=pulumi.get(__ret__, 'domain'),
        filter=pulumi.get(__ret__, 'filter'),
        id=pulumi.get(__ret__, 'id'),
        records=pulumi.get(__ret__, 'records'))
def get_domain_record_output(domain: Optional[pulumi.Input[builtins.str]] = None,
                             filter: Optional[pulumi.Input[Union['GetDomainRecordFilterArgs', 'GetDomainRecordFilterArgsDict']]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDomainRecordResult]:
    """
    Fetch Exoscale [DNS](https://community.exoscale.com/product/networking/dns/) Domain Records data.

    Corresponding resource: exoscale_domain_record.


    :param builtins.str domain: The Domain name to match.
    :param Union['GetDomainRecordFilterArgs', 'GetDomainRecordFilterArgsDict'] filter: Filter to apply when looking up domain records.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['filter'] = filter
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('exoscale:index/getDomainRecord:getDomainRecord', __args__, opts=opts, typ=GetDomainRecordResult)
    return __ret__.apply(lambda __response__: GetDomainRecordResult(
        domain=pulumi.get(__response__, 'domain'),
        filter=pulumi.get(__response__, 'filter'),
        id=pulumi.get(__response__, 'id'),
        records=pulumi.get(__response__, 'records')))
