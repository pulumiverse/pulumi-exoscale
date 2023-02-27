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
    'GetComputeInstanceListResult',
    'AwaitableGetComputeInstanceListResult',
    'get_compute_instance_list',
    'get_compute_instance_list_output',
]

@pulumi.output_type
class GetComputeInstanceListResult:
    """
    A collection of values returned by getComputeInstanceList.
    """
    def __init__(__self__, created_at=None, deploy_target_id=None, disk_size=None, id=None, instances=None, ipv6=None, ipv6_address=None, labels=None, manager_id=None, manager_type=None, name=None, public_ip_address=None, reverse_dns=None, ssh_key=None, state=None, template_id=None, type=None, user_data=None, zone=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if deploy_target_id and not isinstance(deploy_target_id, str):
            raise TypeError("Expected argument 'deploy_target_id' to be a str")
        pulumi.set(__self__, "deploy_target_id", deploy_target_id)
        if disk_size and not isinstance(disk_size, int):
            raise TypeError("Expected argument 'disk_size' to be a int")
        pulumi.set(__self__, "disk_size", disk_size)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if ipv6 and not isinstance(ipv6, bool):
            raise TypeError("Expected argument 'ipv6' to be a bool")
        pulumi.set(__self__, "ipv6", ipv6)
        if ipv6_address and not isinstance(ipv6_address, str):
            raise TypeError("Expected argument 'ipv6_address' to be a str")
        pulumi.set(__self__, "ipv6_address", ipv6_address)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if manager_id and not isinstance(manager_id, str):
            raise TypeError("Expected argument 'manager_id' to be a str")
        pulumi.set(__self__, "manager_id", manager_id)
        if manager_type and not isinstance(manager_type, str):
            raise TypeError("Expected argument 'manager_type' to be a str")
        pulumi.set(__self__, "manager_type", manager_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if public_ip_address and not isinstance(public_ip_address, str):
            raise TypeError("Expected argument 'public_ip_address' to be a str")
        pulumi.set(__self__, "public_ip_address", public_ip_address)
        if reverse_dns and not isinstance(reverse_dns, str):
            raise TypeError("Expected argument 'reverse_dns' to be a str")
        pulumi.set(__self__, "reverse_dns", reverse_dns)
        if ssh_key and not isinstance(ssh_key, str):
            raise TypeError("Expected argument 'ssh_key' to be a str")
        pulumi.set(__self__, "ssh_key", ssh_key)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if template_id and not isinstance(template_id, str):
            raise TypeError("Expected argument 'template_id' to be a str")
        pulumi.set(__self__, "template_id", template_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        pulumi.set(__self__, "user_data", user_data)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deployTargetId")
    def deploy_target_id(self) -> Optional[str]:
        return pulumi.get(self, "deploy_target_id")

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> Optional[int]:
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetComputeInstanceListInstanceResult']:
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[bool]:
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="ipv6Address")
    def ipv6_address(self) -> Optional[str]:
        return pulumi.get(self, "ipv6_address")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="managerId")
    def manager_id(self) -> Optional[str]:
        return pulumi.get(self, "manager_id")

    @property
    @pulumi.getter(name="managerType")
    def manager_type(self) -> Optional[str]:
        return pulumi.get(self, "manager_type")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> Optional[str]:
        return pulumi.get(self, "public_ip_address")

    @property
    @pulumi.getter(name="reverseDns")
    def reverse_dns(self) -> Optional[str]:
        return pulumi.get(self, "reverse_dns")

    @property
    @pulumi.getter(name="sshKey")
    def ssh_key(self) -> Optional[str]:
        return pulumi.get(self, "ssh_key")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[str]:
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> Optional[str]:
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetComputeInstanceListResult(GetComputeInstanceListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeInstanceListResult(
            created_at=self.created_at,
            deploy_target_id=self.deploy_target_id,
            disk_size=self.disk_size,
            id=self.id,
            instances=self.instances,
            ipv6=self.ipv6,
            ipv6_address=self.ipv6_address,
            labels=self.labels,
            manager_id=self.manager_id,
            manager_type=self.manager_type,
            name=self.name,
            public_ip_address=self.public_ip_address,
            reverse_dns=self.reverse_dns,
            ssh_key=self.ssh_key,
            state=self.state,
            template_id=self.template_id,
            type=self.type,
            user_data=self.user_data,
            zone=self.zone)


def get_compute_instance_list(created_at: Optional[str] = None,
                              deploy_target_id: Optional[str] = None,
                              disk_size: Optional[int] = None,
                              id: Optional[str] = None,
                              ipv6: Optional[bool] = None,
                              ipv6_address: Optional[str] = None,
                              labels: Optional[Mapping[str, str]] = None,
                              manager_id: Optional[str] = None,
                              manager_type: Optional[str] = None,
                              name: Optional[str] = None,
                              public_ip_address: Optional[str] = None,
                              reverse_dns: Optional[str] = None,
                              ssh_key: Optional[str] = None,
                              state: Optional[str] = None,
                              template_id: Optional[str] = None,
                              type: Optional[str] = None,
                              user_data: Optional[str] = None,
                              zone: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeInstanceListResult:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The Exoscale [Zone][zone] name.
    """
    __args__ = dict()
    __args__['createdAt'] = created_at
    __args__['deployTargetId'] = deploy_target_id
    __args__['diskSize'] = disk_size
    __args__['id'] = id
    __args__['ipv6'] = ipv6
    __args__['ipv6Address'] = ipv6_address
    __args__['labels'] = labels
    __args__['managerId'] = manager_id
    __args__['managerType'] = manager_type
    __args__['name'] = name
    __args__['publicIpAddress'] = public_ip_address
    __args__['reverseDns'] = reverse_dns
    __args__['sshKey'] = ssh_key
    __args__['state'] = state
    __args__['templateId'] = template_id
    __args__['type'] = type
    __args__['userData'] = user_data
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getComputeInstanceList:getComputeInstanceList', __args__, opts=opts, typ=GetComputeInstanceListResult).value

    return AwaitableGetComputeInstanceListResult(
        created_at=__ret__.created_at,
        deploy_target_id=__ret__.deploy_target_id,
        disk_size=__ret__.disk_size,
        id=__ret__.id,
        instances=__ret__.instances,
        ipv6=__ret__.ipv6,
        ipv6_address=__ret__.ipv6_address,
        labels=__ret__.labels,
        manager_id=__ret__.manager_id,
        manager_type=__ret__.manager_type,
        name=__ret__.name,
        public_ip_address=__ret__.public_ip_address,
        reverse_dns=__ret__.reverse_dns,
        ssh_key=__ret__.ssh_key,
        state=__ret__.state,
        template_id=__ret__.template_id,
        type=__ret__.type,
        user_data=__ret__.user_data,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_compute_instance_list)
def get_compute_instance_list_output(created_at: Optional[pulumi.Input[Optional[str]]] = None,
                                     deploy_target_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     disk_size: Optional[pulumi.Input[Optional[int]]] = None,
                                     id: Optional[pulumi.Input[Optional[str]]] = None,
                                     ipv6: Optional[pulumi.Input[Optional[bool]]] = None,
                                     ipv6_address: Optional[pulumi.Input[Optional[str]]] = None,
                                     labels: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                                     manager_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     manager_type: Optional[pulumi.Input[Optional[str]]] = None,
                                     name: Optional[pulumi.Input[Optional[str]]] = None,
                                     public_ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                                     reverse_dns: Optional[pulumi.Input[Optional[str]]] = None,
                                     ssh_key: Optional[pulumi.Input[Optional[str]]] = None,
                                     state: Optional[pulumi.Input[Optional[str]]] = None,
                                     template_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     type: Optional[pulumi.Input[Optional[str]]] = None,
                                     user_data: Optional[pulumi.Input[Optional[str]]] = None,
                                     zone: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetComputeInstanceListResult]:
    """
    Use this data source to access information about an existing resource.

    :param str zone: The Exoscale [Zone][zone] name.
    """
    ...
