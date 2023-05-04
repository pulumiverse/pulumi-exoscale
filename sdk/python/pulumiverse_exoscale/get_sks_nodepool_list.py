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
    'GetSKSNodepoolListResult',
    'AwaitableGetSKSNodepoolListResult',
    'get_sks_nodepool_list',
    'get_sks_nodepool_list_output',
]

@pulumi.output_type
class GetSKSNodepoolListResult:
    """
    A collection of values returned by getSKSNodepoolList.
    """
    def __init__(__self__, cluster_id=None, created_at=None, deploy_target_id=None, description=None, disk_size=None, id=None, instance_pool_id=None, instance_prefix=None, instance_type=None, labels=None, name=None, nodepools=None, size=None, state=None, taints=None, template_id=None, version=None, zone=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if deploy_target_id and not isinstance(deploy_target_id, str):
            raise TypeError("Expected argument 'deploy_target_id' to be a str")
        pulumi.set(__self__, "deploy_target_id", deploy_target_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size and not isinstance(disk_size, int):
            raise TypeError("Expected argument 'disk_size' to be a int")
        pulumi.set(__self__, "disk_size", disk_size)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_pool_id and not isinstance(instance_pool_id, str):
            raise TypeError("Expected argument 'instance_pool_id' to be a str")
        pulumi.set(__self__, "instance_pool_id", instance_pool_id)
        if instance_prefix and not isinstance(instance_prefix, str):
            raise TypeError("Expected argument 'instance_prefix' to be a str")
        pulumi.set(__self__, "instance_prefix", instance_prefix)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nodepools and not isinstance(nodepools, list):
            raise TypeError("Expected argument 'nodepools' to be a list")
        pulumi.set(__self__, "nodepools", nodepools)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if taints and not isinstance(taints, dict):
            raise TypeError("Expected argument 'taints' to be a dict")
        pulumi.set(__self__, "taints", taints)
        if template_id and not isinstance(template_id, str):
            raise TypeError("Expected argument 'template_id' to be a str")
        pulumi.set(__self__, "template_id", template_id)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deployTargetId")
    def deploy_target_id(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "deploy_target_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> Optional[int]:
        """
        Match against this int
        """
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instancePoolId")
    def instance_pool_id(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "instance_pool_id")

    @property
    @pulumi.getter(name="instancePrefix")
    def instance_prefix(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "instance_prefix")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nodepools(self) -> Sequence['outputs.GetSKSNodepoolListNodepoolResult']:
        return pulumi.get(self, "nodepools")

    @property
    @pulumi.getter
    def size(self) -> Optional[int]:
        """
        Match against this int
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def taints(self) -> Optional[Mapping[str, str]]:
        """
        Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
        """
        return pulumi.get(self, "taints")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
        """
        return pulumi.get(self, "zone")


class AwaitableGetSKSNodepoolListResult(GetSKSNodepoolListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSKSNodepoolListResult(
            cluster_id=self.cluster_id,
            created_at=self.created_at,
            deploy_target_id=self.deploy_target_id,
            description=self.description,
            disk_size=self.disk_size,
            id=self.id,
            instance_pool_id=self.instance_pool_id,
            instance_prefix=self.instance_prefix,
            instance_type=self.instance_type,
            labels=self.labels,
            name=self.name,
            nodepools=self.nodepools,
            size=self.size,
            state=self.state,
            taints=self.taints,
            template_id=self.template_id,
            version=self.version,
            zone=self.zone)


def get_sks_nodepool_list(cluster_id: Optional[str] = None,
                          created_at: Optional[str] = None,
                          deploy_target_id: Optional[str] = None,
                          description: Optional[str] = None,
                          disk_size: Optional[int] = None,
                          id: Optional[str] = None,
                          instance_pool_id: Optional[str] = None,
                          instance_prefix: Optional[str] = None,
                          instance_type: Optional[str] = None,
                          labels: Optional[Mapping[str, str]] = None,
                          name: Optional[str] = None,
                          size: Optional[int] = None,
                          state: Optional[str] = None,
                          taints: Optional[Mapping[str, str]] = None,
                          template_id: Optional[str] = None,
                          version: Optional[str] = None,
                          zone: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSKSNodepoolListResult:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str created_at: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str deploy_target_id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str description: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param int disk_size: Match against this int
    :param str id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str instance_pool_id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str instance_prefix: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str instance_type: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param Mapping[str, str] labels: Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
    :param str name: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param int size: Match against this int
    :param str state: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param Mapping[str, str] taints: Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
    :param str template_id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str version: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str zone: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['createdAt'] = created_at
    __args__['deployTargetId'] = deploy_target_id
    __args__['description'] = description
    __args__['diskSize'] = disk_size
    __args__['id'] = id
    __args__['instancePoolId'] = instance_pool_id
    __args__['instancePrefix'] = instance_prefix
    __args__['instanceType'] = instance_type
    __args__['labels'] = labels
    __args__['name'] = name
    __args__['size'] = size
    __args__['state'] = state
    __args__['taints'] = taints
    __args__['templateId'] = template_id
    __args__['version'] = version
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getSKSNodepoolList:getSKSNodepoolList', __args__, opts=opts, typ=GetSKSNodepoolListResult).value

    return AwaitableGetSKSNodepoolListResult(
        cluster_id=__ret__.cluster_id,
        created_at=__ret__.created_at,
        deploy_target_id=__ret__.deploy_target_id,
        description=__ret__.description,
        disk_size=__ret__.disk_size,
        id=__ret__.id,
        instance_pool_id=__ret__.instance_pool_id,
        instance_prefix=__ret__.instance_prefix,
        instance_type=__ret__.instance_type,
        labels=__ret__.labels,
        name=__ret__.name,
        nodepools=__ret__.nodepools,
        size=__ret__.size,
        state=__ret__.state,
        taints=__ret__.taints,
        template_id=__ret__.template_id,
        version=__ret__.version,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_sks_nodepool_list)
def get_sks_nodepool_list_output(cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 created_at: Optional[pulumi.Input[Optional[str]]] = None,
                                 deploy_target_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 description: Optional[pulumi.Input[Optional[str]]] = None,
                                 disk_size: Optional[pulumi.Input[Optional[int]]] = None,
                                 id: Optional[pulumi.Input[Optional[str]]] = None,
                                 instance_pool_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 instance_prefix: Optional[pulumi.Input[Optional[str]]] = None,
                                 instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                                 labels: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                                 name: Optional[pulumi.Input[Optional[str]]] = None,
                                 size: Optional[pulumi.Input[Optional[int]]] = None,
                                 state: Optional[pulumi.Input[Optional[str]]] = None,
                                 taints: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                                 template_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 version: Optional[pulumi.Input[Optional[str]]] = None,
                                 zone: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSKSNodepoolListResult]:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str created_at: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str deploy_target_id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str description: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param int disk_size: Match against this int
    :param str id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str instance_pool_id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str instance_prefix: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str instance_type: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param Mapping[str, str] labels: Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
    :param str name: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param int size: Match against this int
    :param str state: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param Mapping[str, str] taints: Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
    :param str template_id: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str version: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    :param str zone: Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
    """
    ...
