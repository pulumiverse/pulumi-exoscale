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
from ._inputs import *

__all__ = ['ElasticIPArgs', 'ElasticIP']

@pulumi.input_type
class ElasticIPArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input['ElasticIPHealthcheckArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_dns: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ElasticIP resource.
        :param pulumi.Input[str] zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        :param pulumi.Input[str] address_family: The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        :param pulumi.Input[str] description: A free-form text describing the Elastic IP (EIP).
        :param pulumi.Input['ElasticIPHealthcheckArgs'] healthcheck: Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] reverse_dns: Domain name for reverse DNS record.
        """
        pulumi.set(__self__, "zone", zone)
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if healthcheck is not None:
            pulumi.set(__self__, "healthcheck", healthcheck)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if reverse_dns is not None:
            pulumi.set(__self__, "reverse_dns", reverse_dns)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the Elastic IP (EIP).
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def healthcheck(self) -> Optional[pulumi.Input['ElasticIPHealthcheckArgs']]:
        """
        Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        """
        return pulumi.get(self, "healthcheck")

    @healthcheck.setter
    def healthcheck(self, value: Optional[pulumi.Input['ElasticIPHealthcheckArgs']]):
        pulumi.set(self, "healthcheck", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="reverseDns")
    def reverse_dns(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name for reverse DNS record.
        """
        return pulumi.get(self, "reverse_dns")

    @reverse_dns.setter
    def reverse_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reverse_dns", value)


@pulumi.input_type
class _ElasticIPState:
    def __init__(__self__, *,
                 address_family: Optional[pulumi.Input[str]] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input['ElasticIPHealthcheckArgs']] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_dns: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ElasticIP resources.
        :param pulumi.Input[str] address_family: The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        :param pulumi.Input[str] cidr: The Elastic IP (EIP) CIDR.
        :param pulumi.Input[str] description: A free-form text describing the Elastic IP (EIP).
        :param pulumi.Input['ElasticIPHealthcheckArgs'] healthcheck: Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        :param pulumi.Input[str] ip_address: The Elastic IP (EIP) IPv4 or IPv6 address.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] reverse_dns: Domain name for reverse DNS record.
        :param pulumi.Input[str] zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if healthcheck is not None:
            pulumi.set(__self__, "healthcheck", healthcheck)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if reverse_dns is not None:
            pulumi.set(__self__, "reverse_dns", reverse_dns)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The Elastic IP (EIP) CIDR.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the Elastic IP (EIP).
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def healthcheck(self) -> Optional[pulumi.Input['ElasticIPHealthcheckArgs']]:
        """
        Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        """
        return pulumi.get(self, "healthcheck")

    @healthcheck.setter
    def healthcheck(self, value: Optional[pulumi.Input['ElasticIPHealthcheckArgs']]):
        pulumi.set(self, "healthcheck", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The Elastic IP (EIP) IPv4 or IPv6 address.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="reverseDns")
    def reverse_dns(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name for reverse DNS record.
        """
        return pulumi.get(self, "reverse_dns")

    @reverse_dns.setter
    def reverse_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reverse_dns", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class ElasticIP(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input[pulumi.InputType['ElasticIPHealthcheckArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_dns: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        An existing Elastic IP (EIP) may be imported by `<ID>@<zone>`

        ```sh
         $ pulumi import exoscale:index/elasticIP:ElasticIP \\
        ```

         exoscale_elastic_ip.my_elastic_ip \\

         f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        :param pulumi.Input[str] description: A free-form text describing the Elastic IP (EIP).
        :param pulumi.Input[pulumi.InputType['ElasticIPHealthcheckArgs']] healthcheck: Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] reverse_dns: Domain name for reverse DNS record.
        :param pulumi.Input[str] zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ElasticIPArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        An existing Elastic IP (EIP) may be imported by `<ID>@<zone>`

        ```sh
         $ pulumi import exoscale:index/elasticIP:ElasticIP \\
        ```

         exoscale_elastic_ip.my_elastic_ip \\

         f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2

        :param str resource_name: The name of the resource.
        :param ElasticIPArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ElasticIPArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input[pulumi.InputType['ElasticIPHealthcheckArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_dns: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ElasticIPArgs.__new__(ElasticIPArgs)

            __props__.__dict__["address_family"] = address_family
            __props__.__dict__["description"] = description
            __props__.__dict__["healthcheck"] = healthcheck
            __props__.__dict__["labels"] = labels
            __props__.__dict__["reverse_dns"] = reverse_dns
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["cidr"] = None
            __props__.__dict__["ip_address"] = None
        super(ElasticIP, __self__).__init__(
            'exoscale:index/elasticIP:ElasticIP',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_family: Optional[pulumi.Input[str]] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            healthcheck: Optional[pulumi.Input[pulumi.InputType['ElasticIPHealthcheckArgs']]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            reverse_dns: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'ElasticIP':
        """
        Get an existing ElasticIP resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        :param pulumi.Input[str] cidr: The Elastic IP (EIP) CIDR.
        :param pulumi.Input[str] description: A free-form text describing the Elastic IP (EIP).
        :param pulumi.Input[pulumi.InputType['ElasticIPHealthcheckArgs']] healthcheck: Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        :param pulumi.Input[str] ip_address: The Elastic IP (EIP) IPv4 or IPv6 address.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] reverse_dns: Domain name for reverse DNS record.
        :param pulumi.Input[str] zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ElasticIPState.__new__(_ElasticIPState)

        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["description"] = description
        __props__.__dict__["healthcheck"] = healthcheck
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["labels"] = labels
        __props__.__dict__["reverse_dns"] = reverse_dns
        __props__.__dict__["zone"] = zone
        return ElasticIP(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Output[str]:
        """
        The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        The Elastic IP (EIP) CIDR.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A free-form text describing the Elastic IP (EIP).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def healthcheck(self) -> pulumi.Output['outputs.ElasticIPHealthcheck']:
        """
        Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        """
        return pulumi.get(self, "healthcheck")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The Elastic IP (EIP) IPv4 or IPv6 address.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="reverseDns")
    def reverse_dns(self) -> pulumi.Output[Optional[str]]:
        """
        Domain name for reverse DNS record.
        """
        return pulumi.get(self, "reverse_dns")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

