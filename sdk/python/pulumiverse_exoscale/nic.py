# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NICArgs', 'NIC']

@pulumi.input_type
class NICArgs:
    def __init__(__self__, *,
                 compute_id: pulumi.Input[str],
                 network_id: pulumi.Input[str],
                 ip_address: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NIC resource.
        :param pulumi.Input[str] compute_id: ❗ The compute instance ID.
        :param pulumi.Input[str] network_id: ❗ The private network ID.
        :param pulumi.Input[str] ip_address: The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        """
        pulumi.set(__self__, "compute_id", compute_id)
        pulumi.set(__self__, "network_id", network_id)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)

    @property
    @pulumi.getter(name="computeId")
    def compute_id(self) -> pulumi.Input[str]:
        """
        ❗ The compute instance ID.
        """
        return pulumi.get(self, "compute_id")

    @compute_id.setter
    def compute_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "compute_id", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        """
        ❗ The private network ID.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)


@pulumi.input_type
class _NICState:
    def __init__(__self__, *,
                 compute_id: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 mac_address: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NIC resources.
        :param pulumi.Input[str] compute_id: ❗ The compute instance ID.
        :param pulumi.Input[str] ip_address: The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        :param pulumi.Input[str] mac_address: The NIC MAC address.
        :param pulumi.Input[str] network_id: ❗ The private network ID.
        """
        if compute_id is not None:
            pulumi.set(__self__, "compute_id", compute_id)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if mac_address is not None:
            pulumi.set(__self__, "mac_address", mac_address)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)

    @property
    @pulumi.getter(name="computeId")
    def compute_id(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The compute instance ID.
        """
        return pulumi.get(self, "compute_id")

    @compute_id.setter
    def compute_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_id", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[pulumi.Input[str]]:
        """
        The NIC MAC address.
        """
        return pulumi.get(self, "mac_address")

    @mac_address.setter
    def mac_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_address", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The private network ID.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)


class NIC(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use ComputeInstance `network_interface` block instead.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_id: ❗ The compute instance ID.
        :param pulumi.Input[str] ip_address: The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        :param pulumi.Input[str] network_id: ❗ The private network ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NICArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use ComputeInstance `network_interface` block instead.

        :param str resource_name: The name of the resource.
        :param NICArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NICArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NICArgs.__new__(NICArgs)

            if compute_id is None and not opts.urn:
                raise TypeError("Missing required property 'compute_id'")
            __props__.__dict__["compute_id"] = compute_id
            __props__.__dict__["ip_address"] = ip_address
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["gateway"] = None
            __props__.__dict__["mac_address"] = None
            __props__.__dict__["netmask"] = None
        super(NIC, __self__).__init__(
            'exoscale:index/nIC:NIC',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compute_id: Optional[pulumi.Input[str]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            mac_address: Optional[pulumi.Input[str]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None) -> 'NIC':
        """
        Get an existing NIC resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_id: ❗ The compute instance ID.
        :param pulumi.Input[str] ip_address: The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        :param pulumi.Input[str] mac_address: The NIC MAC address.
        :param pulumi.Input[str] network_id: ❗ The private network ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NICState.__new__(_NICState)

        __props__.__dict__["compute_id"] = compute_id
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["mac_address"] = mac_address
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["network_id"] = network_id
        return NIC(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="computeId")
    def compute_id(self) -> pulumi.Output[str]:
        """
        ❗ The compute instance ID.
        """
        return pulumi.get(self, "compute_id")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The IPv4 address to request as static DHCP lease if the NIC is attached to a *managed* private network.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> pulumi.Output[str]:
        """
        The NIC MAC address.
        """
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Output[str]:
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        ❗ The private network ID.
        """
        return pulumi.get(self, "network_id")

