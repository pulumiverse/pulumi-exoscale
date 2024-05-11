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
    def __init__(__self__, aggregation_ca=None, auto_upgrade=None, clusters=None, cni=None, control_plane_ca=None, created_at=None, description=None, endpoint=None, exoscale_ccm=None, exoscale_csi=None, id=None, kubelet_ca=None, labels=None, metrics_server=None, name=None, service_level=None, state=None, version=None, zone=None):
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
    def aggregation_ca(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "aggregation_ca")

    @property
    @pulumi.getter(name="autoUpgrade")
    def auto_upgrade(self) -> Optional[bool]:
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
    def cni(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "cni")

    @property
    @pulumi.getter(name="controlPlaneCa")
    def control_plane_ca(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "control_plane_ca")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="exoscaleCcm")
    def exoscale_ccm(self) -> Optional[bool]:
        """
        Match against this bool
        """
        return pulumi.get(self, "exoscale_ccm")

    @property
    @pulumi.getter(name="exoscaleCsi")
    def exoscale_csi(self) -> Optional[bool]:
        """
        Match against this bool
        """
        return pulumi.get(self, "exoscale_csi")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kubeletCa")
    def kubelet_ca(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "kubelet_ca")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="metricsServer")
    def metrics_server(self) -> Optional[bool]:
        """
        Match against this bool
        """
        return pulumi.get(self, "metrics_server")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "service_level")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zone(self) -> str:
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


def get_sks_cluster_list(aggregation_ca: Optional[str] = None,
                         auto_upgrade: Optional[bool] = None,
                         cni: Optional[str] = None,
                         control_plane_ca: Optional[str] = None,
                         created_at: Optional[str] = None,
                         description: Optional[str] = None,
                         endpoint: Optional[str] = None,
                         exoscale_ccm: Optional[bool] = None,
                         exoscale_csi: Optional[bool] = None,
                         id: Optional[str] = None,
                         kubelet_ca: Optional[str] = None,
                         labels: Optional[Mapping[str, str]] = None,
                         metrics_server: Optional[bool] = None,
                         name: Optional[str] = None,
                         service_level: Optional[str] = None,
                         state: Optional[str] = None,
                         version: Optional[str] = None,
                         zone: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSksClusterListResult:
    """
    Use this data source to access information about an existing resource.

    :param str aggregation_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param bool auto_upgrade: Match against this bool
    :param str cni: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str control_plane_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str created_at: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str description: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str endpoint: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param bool exoscale_ccm: Match against this bool
    :param bool exoscale_csi: Match against this bool
    :param str id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str kubelet_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param Mapping[str, str] labels: Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
    :param bool metrics_server: Match against this bool
    :param str name: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str service_level: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str state: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str version: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['aggregationCa'] = aggregation_ca
    __args__['autoUpgrade'] = auto_upgrade
    __args__['cni'] = cni
    __args__['controlPlaneCa'] = control_plane_ca
    __args__['createdAt'] = created_at
    __args__['description'] = description
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


@_utilities.lift_output_func(get_sks_cluster_list)
def get_sks_cluster_list_output(aggregation_ca: Optional[pulumi.Input[Optional[str]]] = None,
                                auto_upgrade: Optional[pulumi.Input[Optional[bool]]] = None,
                                cni: Optional[pulumi.Input[Optional[str]]] = None,
                                control_plane_ca: Optional[pulumi.Input[Optional[str]]] = None,
                                created_at: Optional[pulumi.Input[Optional[str]]] = None,
                                description: Optional[pulumi.Input[Optional[str]]] = None,
                                endpoint: Optional[pulumi.Input[Optional[str]]] = None,
                                exoscale_ccm: Optional[pulumi.Input[Optional[bool]]] = None,
                                exoscale_csi: Optional[pulumi.Input[Optional[bool]]] = None,
                                id: Optional[pulumi.Input[Optional[str]]] = None,
                                kubelet_ca: Optional[pulumi.Input[Optional[str]]] = None,
                                labels: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                                metrics_server: Optional[pulumi.Input[Optional[bool]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                service_level: Optional[pulumi.Input[Optional[str]]] = None,
                                state: Optional[pulumi.Input[Optional[str]]] = None,
                                version: Optional[pulumi.Input[Optional[str]]] = None,
                                zone: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSksClusterListResult]:
    """
    Use this data source to access information about an existing resource.

    :param str aggregation_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param bool auto_upgrade: Match against this bool
    :param str cni: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str control_plane_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str created_at: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str description: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str endpoint: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param bool exoscale_ccm: Match against this bool
    :param bool exoscale_csi: Match against this bool
    :param str id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str kubelet_ca: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param Mapping[str, str] labels: Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
    :param bool metrics_server: Match against this bool
    :param str name: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str service_level: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str state: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str version: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    ...
