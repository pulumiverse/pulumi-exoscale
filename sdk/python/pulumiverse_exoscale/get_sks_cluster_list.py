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

__all__ = [
    'GetSksClusterListResult',
    'AwaitableGetSksClusterListResult',
    'get_sks_cluster_list',
    'get_sks_cluster_list_output',
]

@pulumi.output_type
class GetSksClusterListResult:
    """
    A collection of values returned by getSksClusterList.
    """
    def __init__(__self__, aggregation_ca=None, auto_upgrade=None, clusters=None, cni=None, control_plane_ca=None, created_at=None, description=None, enable_kube_proxy=None, endpoint=None, exoscale_ccm=None, exoscale_csi=None, id=None, kubelet_ca=None, labels=None, metrics_server=None, name=None, service_level=None, state=None, version=None, zone=None):
        if aggregation_ca and not isinstance(aggregation_ca, str):
            raise TypeError("Expected argument 'aggregation_ca' to be a str")
        pulumi.set(__self__, "aggregation_ca", aggregation_ca)
        if auto_upgrade and not isinstance(auto_upgrade, bool):
            raise TypeError("Expected argument 'auto_upgrade' to be a bool")
        pulumi.set(__self__, "auto_upgrade", auto_upgrade)
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if cni and not isinstance(cni, str):
            raise TypeError("Expected argument 'cni' to be a str")
        pulumi.set(__self__, "cni", cni)
        if control_plane_ca and not isinstance(control_plane_ca, str):
            raise TypeError("Expected argument 'control_plane_ca' to be a str")
        pulumi.set(__self__, "control_plane_ca", control_plane_ca)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_kube_proxy and not isinstance(enable_kube_proxy, bool):
            raise TypeError("Expected argument 'enable_kube_proxy' to be a bool")
        pulumi.set(__self__, "enable_kube_proxy", enable_kube_proxy)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if exoscale_ccm and not isinstance(exoscale_ccm, bool):
            raise TypeError("Expected argument 'exoscale_ccm' to be a bool")
        pulumi.set(__self__, "exoscale_ccm", exoscale_ccm)
        if exoscale_csi and not isinstance(exoscale_csi, bool):
            raise TypeError("Expected argument 'exoscale_csi' to be a bool")
        pulumi.set(__self__, "exoscale_csi", exoscale_csi)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kubelet_ca and not isinstance(kubelet_ca, str):
            raise TypeError("Expected argument 'kubelet_ca' to be a str")
        pulumi.set(__self__, "kubelet_ca", kubelet_ca)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if metrics_server and not isinstance(metrics_server, bool):
            raise TypeError("Expected argument 'metrics_server' to be a bool")
        pulumi.set(__self__, "metrics_server", metrics_server)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_level and not isinstance(service_level, str):
            raise TypeError("Expected argument 'service_level' to be a str")
        pulumi.set(__self__, "service_level", service_level)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="aggregationCa")
    def aggregation_ca(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "aggregation_ca")

    @property
    @pulumi.getter(name="autoUpgrade")
    def auto_upgrade(self) -> Optional[builtins.bool]:
        """
        Match against this bool
        """
        return pulumi.get(self, "auto_upgrade")

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetSksClusterListClusterResult']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter
    def cni(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "cni")

    @property
    @pulumi.getter(name="controlPlaneCa")
    def control_plane_ca(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "control_plane_ca")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableKubeProxy")
    def enable_kube_proxy(self) -> Optional[builtins.bool]:
        """
        Match against this bool
        """
        return pulumi.get(self, "enable_kube_proxy")

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="exoscaleCcm")
    def exoscale_ccm(self) -> Optional[builtins.bool]:
        """
        Match against this bool
        """
        return pulumi.get(self, "exoscale_ccm")

    @property
    @pulumi.getter(name="exoscaleCsi")
    def exoscale_csi(self) -> Optional[builtins.bool]:
        """
        Match against this bool
        """
        return pulumi.get(self, "exoscale_csi")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kubeletCa")
    def kubelet_ca(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "kubelet_ca")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="metricsServer")
    def metrics_server(self) -> Optional[builtins.bool]:
        """
        Match against this bool
        """
        return pulumi.get(self, "metrics_server")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "service_level")

    @property
    @pulumi.getter
    def state(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def version(self) -> Optional[builtins.str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zone(self) -> builtins.str:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")


class AwaitableGetSksClusterListResult(GetSksClusterListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSksClusterListResult(
            aggregation_ca=self.aggregation_ca,
            auto_upgrade=self.auto_upgrade,
            clusters=self.clusters,
            cni=self.cni,
            control_plane_ca=self.control_plane_ca,
            created_at=self.created_at,
            description=self.description,
            enable_kube_proxy=self.enable_kube_proxy,
            endpoint=self.endpoint,
            exoscale_ccm=self.exoscale_ccm,
            exoscale_csi=self.exoscale_csi,
            id=self.id,
            kubelet_ca=self.kubelet_ca,
            labels=self.labels,
            metrics_server=self.metrics_server,
            name=self.name,
            service_level=self.service_level,
            state=self.state,
            version=self.version,
            zone=self.zone)


def get_sks_cluster_list(aggregation_ca: Optional[builtins.str] = None,
                         auto_upgrade: Optional[builtins.bool] = None,
                         cni: Optional[builtins.str] = None,
                         control_plane_ca: Optional[builtins.str] = None,
                         created_at: Optional[builtins.str] = None,
                         description: Optional[builtins.str] = None,
                         enable_kube_proxy: Optional[builtins.bool] = None,
                         endpoint: Optional[builtins.str] = None,
                         exoscale_ccm: Optional[builtins.bool] = None,
                         exoscale_csi: Optional[builtins.bool] = None,
                         id: Optional[builtins.str] = None,
                         kubelet_ca: Optional[builtins.str] = None,
                         labels: Optional[Mapping[str, builtins.str]] = None,
                         metrics_server: Optional[builtins.bool] = None,
                         name: Optional[builtins.str] = None,
                         service_level: Optional[builtins.str] = None,
                         state: Optional[builtins.str] = None,
                         version: Optional[builtins.str] = None,
                         zone: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSksClusterListResult:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str aggregation_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.bool auto_upgrade: Match against this bool
    :param builtins.str cni: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str control_plane_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str created_at: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str description: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.bool enable_kube_proxy: Match against this bool
    :param builtins.str endpoint: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.bool exoscale_ccm: Match against this bool
    :param builtins.bool exoscale_csi: Match against this bool
    :param builtins.str id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str kubelet_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param Mapping[str, builtins.str] labels: Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
    :param builtins.bool metrics_server: Match against this bool
    :param builtins.str name: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str service_level: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str state: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str version: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['aggregationCa'] = aggregation_ca
    __args__['autoUpgrade'] = auto_upgrade
    __args__['cni'] = cni
    __args__['controlPlaneCa'] = control_plane_ca
    __args__['createdAt'] = created_at
    __args__['description'] = description
    __args__['enableKubeProxy'] = enable_kube_proxy
    __args__['endpoint'] = endpoint
    __args__['exoscaleCcm'] = exoscale_ccm
    __args__['exoscaleCsi'] = exoscale_csi
    __args__['id'] = id
    __args__['kubeletCa'] = kubelet_ca
    __args__['labels'] = labels
    __args__['metricsServer'] = metrics_server
    __args__['name'] = name
    __args__['serviceLevel'] = service_level
    __args__['state'] = state
    __args__['version'] = version
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getSksClusterList:getSksClusterList', __args__, opts=opts, typ=GetSksClusterListResult).value

    return AwaitableGetSksClusterListResult(
        aggregation_ca=pulumi.get(__ret__, 'aggregation_ca'),
        auto_upgrade=pulumi.get(__ret__, 'auto_upgrade'),
        clusters=pulumi.get(__ret__, 'clusters'),
        cni=pulumi.get(__ret__, 'cni'),
        control_plane_ca=pulumi.get(__ret__, 'control_plane_ca'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        enable_kube_proxy=pulumi.get(__ret__, 'enable_kube_proxy'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        exoscale_ccm=pulumi.get(__ret__, 'exoscale_ccm'),
        exoscale_csi=pulumi.get(__ret__, 'exoscale_csi'),
        id=pulumi.get(__ret__, 'id'),
        kubelet_ca=pulumi.get(__ret__, 'kubelet_ca'),
        labels=pulumi.get(__ret__, 'labels'),
        metrics_server=pulumi.get(__ret__, 'metrics_server'),
        name=pulumi.get(__ret__, 'name'),
        service_level=pulumi.get(__ret__, 'service_level'),
        state=pulumi.get(__ret__, 'state'),
        version=pulumi.get(__ret__, 'version'),
        zone=pulumi.get(__ret__, 'zone'))
def get_sks_cluster_list_output(aggregation_ca: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                auto_upgrade: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                cni: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                control_plane_ca: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                created_at: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                description: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                enable_kube_proxy: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                endpoint: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                exoscale_ccm: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                exoscale_csi: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                kubelet_ca: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                labels: Optional[pulumi.Input[Optional[Mapping[str, builtins.str]]]] = None,
                                metrics_server: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                service_level: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                state: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                version: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                zone: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSksClusterListResult]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str aggregation_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.bool auto_upgrade: Match against this bool
    :param builtins.str cni: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str control_plane_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str created_at: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str description: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.bool enable_kube_proxy: Match against this bool
    :param builtins.str endpoint: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.bool exoscale_ccm: Match against this bool
    :param builtins.bool exoscale_csi: Match against this bool
    :param builtins.str id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str kubelet_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param Mapping[str, builtins.str] labels: Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
    :param builtins.bool metrics_server: Match against this bool
    :param builtins.str name: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str service_level: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str state: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str version: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param builtins.str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['aggregationCa'] = aggregation_ca
    __args__['autoUpgrade'] = auto_upgrade
    __args__['cni'] = cni
    __args__['controlPlaneCa'] = control_plane_ca
    __args__['createdAt'] = created_at
    __args__['description'] = description
    __args__['enableKubeProxy'] = enable_kube_proxy
    __args__['endpoint'] = endpoint
    __args__['exoscaleCcm'] = exoscale_ccm
    __args__['exoscaleCsi'] = exoscale_csi
    __args__['id'] = id
    __args__['kubeletCa'] = kubelet_ca
    __args__['labels'] = labels
    __args__['metricsServer'] = metrics_server
    __args__['name'] = name
    __args__['serviceLevel'] = service_level
    __args__['state'] = state
    __args__['version'] = version
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('exoscale:index/getSksClusterList:getSksClusterList', __args__, opts=opts, typ=GetSksClusterListResult)
    return __ret__.apply(lambda __response__: GetSksClusterListResult(
        aggregation_ca=pulumi.get(__response__, 'aggregation_ca'),
        auto_upgrade=pulumi.get(__response__, 'auto_upgrade'),
        clusters=pulumi.get(__response__, 'clusters'),
        cni=pulumi.get(__response__, 'cni'),
        control_plane_ca=pulumi.get(__response__, 'control_plane_ca'),
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        enable_kube_proxy=pulumi.get(__response__, 'enable_kube_proxy'),
        endpoint=pulumi.get(__response__, 'endpoint'),
        exoscale_ccm=pulumi.get(__response__, 'exoscale_ccm'),
        exoscale_csi=pulumi.get(__response__, 'exoscale_csi'),
        id=pulumi.get(__response__, 'id'),
        kubelet_ca=pulumi.get(__response__, 'kubelet_ca'),
        labels=pulumi.get(__response__, 'labels'),
        metrics_server=pulumi.get(__response__, 'metrics_server'),
        name=pulumi.get(__response__, 'name'),
        service_level=pulumi.get(__response__, 'service_level'),
        state=pulumi.get(__response__, 'state'),
        version=pulumi.get(__response__, 'version'),
        zone=pulumi.get(__response__, 'zone')))
