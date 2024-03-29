# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('exoscale')


class _ExportableConfig(types.ModuleType):
    @property
    def delay(self) -> Optional[int]:
        return __config__.get_int('delay')

    @property
    def environment(self) -> Optional[str]:
        return __config__.get('environment')

    @property
    def key(self) -> Optional[str]:
        """
        Exoscale API key
        """
        return __config__.get('key') or _utilities.get_env('EXOSCALE_API_KEY')

    @property
    def secret(self) -> Optional[str]:
        """
        Exoscale API secret
        """
        return __config__.get('secret') or _utilities.get_env('EXOSCALE_API_SECRET')

    @property
    def timeout(self) -> Optional[int]:
        """
        Timeout in seconds for waiting on compute resources to become available (by default: 300)
        """
        return __config__.get_int('timeout')

