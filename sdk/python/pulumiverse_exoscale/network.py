# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NetworkArgs', 'Network']

@pulumi.input_type
class NetworkArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 display_text: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network_offering: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Network resource.
        :param pulumi.Input[str] zone: ❗ The Exoscale Zone name.
        :param pulumi.Input[str] display_text: A free-form text describing the network.
        :param pulumi.Input[str] end_ip: The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        :param pulumi.Input[str] name: The private network name.
        :param pulumi.Input[str] netmask: The network mask defining the IP network allowed for static leases (see `NIC` resource). Required for *managed* private networks.
        :param pulumi.Input[str] start_ip: The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags (key/value). To remove all tags, set `tags = {}`.
        """
        pulumi.set(__self__, "zone", zone)
        if display_text is not None:
            pulumi.set(__self__, "display_text", display_text)
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if network_offering is not None:
            warnings.warn("""This attribute is deprecated, please remove it from your configuration.""", DeprecationWarning)
            pulumi.log.warn("""network_offering is deprecated: This attribute is deprecated, please remove it from your configuration.""")
        if network_offering is not None:
            pulumi.set(__self__, "network_offering", network_offering)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        ❗ The Exoscale Zone name.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="displayText")
    def display_text(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the network.
        """
        return pulumi.get(self, "display_text")

    @display_text.setter
    def display_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_text", value)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The private network name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        The network mask defining the IP network allowed for static leases (see `NIC` resource). Required for *managed* private networks.
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter(name="networkOffering")
    def network_offering(self) -> Optional[pulumi.Input[str]]:
        warnings.warn("""This attribute is deprecated, please remove it from your configuration.""", DeprecationWarning)
        pulumi.log.warn("""network_offering is deprecated: This attribute is deprecated, please remove it from your configuration.""")

        return pulumi.get(self, "network_offering")

    @network_offering.setter
    def network_offering(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_offering", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags (key/value). To remove all tags, set `tags = {}`.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _NetworkState:
    def __init__(__self__, *,
                 display_text: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network_offering: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Network resources.
        :param pulumi.Input[str] display_text: A free-form text describing the network.
        :param pulumi.Input[str] end_ip: The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        :param pulumi.Input[str] name: The private network name.
        :param pulumi.Input[str] netmask: The network mask defining the IP network allowed for static leases (see `NIC` resource). Required for *managed* private networks.
        :param pulumi.Input[str] start_ip: The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags (key/value). To remove all tags, set `tags = {}`.
        :param pulumi.Input[str] zone: ❗ The Exoscale Zone name.
        """
        if display_text is not None:
            pulumi.set(__self__, "display_text", display_text)
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if network_offering is not None:
            warnings.warn("""This attribute is deprecated, please remove it from your configuration.""", DeprecationWarning)
            pulumi.log.warn("""network_offering is deprecated: This attribute is deprecated, please remove it from your configuration.""")
        if network_offering is not None:
            pulumi.set(__self__, "network_offering", network_offering)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="displayText")
    def display_text(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the network.
        """
        return pulumi.get(self, "display_text")

    @display_text.setter
    def display_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_text", value)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The private network name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        The network mask defining the IP network allowed for static leases (see `NIC` resource). Required for *managed* private networks.
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter(name="networkOffering")
    def network_offering(self) -> Optional[pulumi.Input[str]]:
        warnings.warn("""This attribute is deprecated, please remove it from your configuration.""", DeprecationWarning)
        pulumi.log.warn("""network_offering is deprecated: This attribute is deprecated, please remove it from your configuration.""")

        return pulumi.get(self, "network_offering")

    @network_offering.setter
    def network_offering(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_offering", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags (key/value). To remove all tags, set `tags = {}`.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The Exoscale Zone name.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class Network(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_text: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network_offering: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use PrivateNetwork instead.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_text: A free-form text describing the network.
        :param pulumi.Input[str] end_ip: The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        :param pulumi.Input[str] name: The private network name.
        :param pulumi.Input[str] netmask: The network mask defining the IP network allowed for static leases (see `NIC` resource). Required for *managed* private networks.
        :param pulumi.Input[str] start_ip: The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags (key/value). To remove all tags, set `tags = {}`.
        :param pulumi.Input[str] zone: ❗ The Exoscale Zone name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use PrivateNetwork instead.

        :param str resource_name: The name of the resource.
        :param NetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_text: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network_offering: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkArgs.__new__(NetworkArgs)

            __props__.__dict__["display_text"] = display_text
            __props__.__dict__["end_ip"] = end_ip
            __props__.__dict__["name"] = name
            __props__.__dict__["netmask"] = netmask
            __props__.__dict__["network_offering"] = network_offering
            __props__.__dict__["start_ip"] = start_ip
            __props__.__dict__["tags"] = tags
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(Network, __self__).__init__(
            'exoscale:index/network:Network',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_text: Optional[pulumi.Input[str]] = None,
            end_ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            network_offering: Optional[pulumi.Input[str]] = None,
            start_ip: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Network':
        """
        Get an existing Network resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_text: A free-form text describing the network.
        :param pulumi.Input[str] end_ip: The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        :param pulumi.Input[str] name: The private network name.
        :param pulumi.Input[str] netmask: The network mask defining the IP network allowed for static leases (see `NIC` resource). Required for *managed* private networks.
        :param pulumi.Input[str] start_ip: The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags (key/value). To remove all tags, set `tags = {}`.
        :param pulumi.Input[str] zone: ❗ The Exoscale Zone name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkState.__new__(_NetworkState)

        __props__.__dict__["display_text"] = display_text
        __props__.__dict__["end_ip"] = end_ip
        __props__.__dict__["name"] = name
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["network_offering"] = network_offering
        __props__.__dict__["start_ip"] = start_ip
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zone"] = zone
        return Network(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayText")
    def display_text(self) -> pulumi.Output[str]:
        """
        A free-form text describing the network.
        """
        return pulumi.get(self, "display_text")

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The private network name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Output[Optional[str]]:
        """
        The network mask defining the IP network allowed for static leases (see `NIC` resource). Required for *managed* private networks.
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter(name="networkOffering")
    def network_offering(self) -> pulumi.Output[Optional[str]]:
        warnings.warn("""This attribute is deprecated, please remove it from your configuration.""", DeprecationWarning)
        pulumi.log.warn("""network_offering is deprecated: This attribute is deprecated, please remove it from your configuration.""")

        return pulumi.get(self, "network_offering")

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The first/last IP addresses used by the DHCP service for dynamic leases. Required for *managed* private networks.
        """
        return pulumi.get(self, "start_ip")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of tags (key/value). To remove all tags, set `tags = {}`.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        ❗ The Exoscale Zone name.
        """
        return pulumi.get(self, "zone")

