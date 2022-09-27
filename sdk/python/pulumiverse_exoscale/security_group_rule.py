# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SecurityGroupRuleArgs', 'SecurityGroupRule']

@pulumi.input_type
class SecurityGroupRuleArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_port: Optional[pulumi.Input[int]] = None,
                 icmp_code: Optional[pulumi.Input[int]] = None,
                 icmp_type: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 start_port: Optional[pulumi.Input[int]] = None,
                 user_security_group: Optional[pulumi.Input[str]] = None,
                 user_security_group_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecurityGroupRule resource.
        :param pulumi.Input[str] type: The traffic direction to match (`INGRESS` or `EGRESS`).
        :param pulumi.Input[str] cidr: An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation][cidr]) to match (conflicts with `user_security_group`/`user_security_group_id`).
               * `start_port`/`end_port` - A `TCP`/`UDP` port range to match.
               * `icmp_type`/`icmp_code` - An ICMP/ICMPv6 [type/code][icmp] to match.
        :param pulumi.Input[str] description: A free-form text describing the security group rule.
        :param pulumi.Input[str] protocol: The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        :param pulumi.Input[str] security_group: The parent security group name. Please use the `security_group_id` argument along the SecurityGroup data source instead.
        :param pulumi.Input[str] security_group_id: The parent SecurityGroup ID.
        :param pulumi.Input[str] user_security_group: An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`user_security_group_id`). Please use the `user_security_group_id` argument along the SecurityGroup data source instead.
        :param pulumi.Input[str] user_security_group_id: An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`user_security_group)`).
        """
        pulumi.set(__self__, "type", type)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_port is not None:
            pulumi.set(__self__, "end_port", end_port)
        if icmp_code is not None:
            pulumi.set(__self__, "icmp_code", icmp_code)
        if icmp_type is not None:
            pulumi.set(__self__, "icmp_type", icmp_type)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if start_port is not None:
            pulumi.set(__self__, "start_port", start_port)
        if user_security_group is not None:
            pulumi.set(__self__, "user_security_group", user_security_group)
        if user_security_group_id is not None:
            pulumi.set(__self__, "user_security_group_id", user_security_group_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The traffic direction to match (`INGRESS` or `EGRESS`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation][cidr]) to match (conflicts with `user_security_group`/`user_security_group_id`).
        * `start_port`/`end_port` - A `TCP`/`UDP` port range to match.
        * `icmp_type`/`icmp_code` - An ICMP/ICMPv6 [type/code][icmp] to match.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the security group rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end_port")

    @end_port.setter
    def end_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_port", value)

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "icmp_code")

    @icmp_code.setter
    def icmp_code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_code", value)

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "icmp_type")

    @icmp_type.setter
    def icmp_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_type", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[str]]:
        """
        The parent security group name. Please use the `security_group_id` argument along the SecurityGroup data source instead.
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The parent SecurityGroup ID.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start_port")

    @start_port.setter
    def start_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_port", value)

    @property
    @pulumi.getter(name="userSecurityGroup")
    def user_security_group(self) -> Optional[pulumi.Input[str]]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`user_security_group_id`). Please use the `user_security_group_id` argument along the SecurityGroup data source instead.
        """
        return pulumi.get(self, "user_security_group")

    @user_security_group.setter
    def user_security_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_security_group", value)

    @property
    @pulumi.getter(name="userSecurityGroupId")
    def user_security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`user_security_group)`).
        """
        return pulumi.get(self, "user_security_group_id")

    @user_security_group_id.setter
    def user_security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_security_group_id", value)


@pulumi.input_type
class _SecurityGroupRuleState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_port: Optional[pulumi.Input[int]] = None,
                 icmp_code: Optional[pulumi.Input[int]] = None,
                 icmp_type: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 start_port: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_security_group: Optional[pulumi.Input[str]] = None,
                 user_security_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityGroupRule resources.
        :param pulumi.Input[str] cidr: An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation][cidr]) to match (conflicts with `user_security_group`/`user_security_group_id`).
               * `start_port`/`end_port` - A `TCP`/`UDP` port range to match.
               * `icmp_type`/`icmp_code` - An ICMP/ICMPv6 [type/code][icmp] to match.
        :param pulumi.Input[str] description: A free-form text describing the security group rule.
        :param pulumi.Input[str] protocol: The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        :param pulumi.Input[str] security_group: The parent security group name. Please use the `security_group_id` argument along the SecurityGroup data source instead.
        :param pulumi.Input[str] security_group_id: The parent SecurityGroup ID.
        :param pulumi.Input[str] type: The traffic direction to match (`INGRESS` or `EGRESS`).
        :param pulumi.Input[str] user_security_group: An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`user_security_group_id`). Please use the `user_security_group_id` argument along the SecurityGroup data source instead.
        :param pulumi.Input[str] user_security_group_id: An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`user_security_group)`).
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_port is not None:
            pulumi.set(__self__, "end_port", end_port)
        if icmp_code is not None:
            pulumi.set(__self__, "icmp_code", icmp_code)
        if icmp_type is not None:
            pulumi.set(__self__, "icmp_type", icmp_type)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if start_port is not None:
            pulumi.set(__self__, "start_port", start_port)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_security_group is not None:
            pulumi.set(__self__, "user_security_group", user_security_group)
        if user_security_group_id is not None:
            pulumi.set(__self__, "user_security_group_id", user_security_group_id)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation][cidr]) to match (conflicts with `user_security_group`/`user_security_group_id`).
        * `start_port`/`end_port` - A `TCP`/`UDP` port range to match.
        * `icmp_type`/`icmp_code` - An ICMP/ICMPv6 [type/code][icmp] to match.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the security group rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end_port")

    @end_port.setter
    def end_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_port", value)

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "icmp_code")

    @icmp_code.setter
    def icmp_code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_code", value)

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "icmp_type")

    @icmp_type.setter
    def icmp_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_type", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[str]]:
        """
        The parent security group name. Please use the `security_group_id` argument along the SecurityGroup data source instead.
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The parent SecurityGroup ID.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start_port")

    @start_port.setter
    def start_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_port", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The traffic direction to match (`INGRESS` or `EGRESS`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userSecurityGroup")
    def user_security_group(self) -> Optional[pulumi.Input[str]]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`user_security_group_id`). Please use the `user_security_group_id` argument along the SecurityGroup data source instead.
        """
        return pulumi.get(self, "user_security_group")

    @user_security_group.setter
    def user_security_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_security_group", value)

    @property
    @pulumi.getter(name="userSecurityGroupId")
    def user_security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`user_security_group)`).
        """
        return pulumi.get(self, "user_security_group_id")

    @user_security_group_id.setter
    def user_security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_security_group_id", value)


class SecurityGroupRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_port: Optional[pulumi.Input[int]] = None,
                 icmp_code: Optional[pulumi.Input[int]] = None,
                 icmp_type: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 start_port: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_security_group: Optional[pulumi.Input[str]] = None,
                 user_security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage Exoscale [Security Group](https://community.exoscale.com/documentation/compute/security-groups/) Rules.

        ## Usage

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_security_group = exoscale.SecurityGroup("mySecurityGroup")
        my_security_group_rule = exoscale.SecurityGroupRule("mySecurityGroupRule",
            security_group_id=my_security_group.id,
            type="INGRESS",
            protocol="TCP",
            cidr="0.0.0.0/0",
            start_port=80,
            end_port=80)
        ```

        ## Import

        An existing security group rule may be imported by `<security-group-ID>/<security-group-rule-ID>`console

        ```sh
         $ pulumi import exoscale:index/securityGroupRule:SecurityGroupRule \\
        ```

         exoscale_security_group_rule.my_security_group_rule \\

         f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation][cidr]) to match (conflicts with `user_security_group`/`user_security_group_id`).
               * `start_port`/`end_port` - A `TCP`/`UDP` port range to match.
               * `icmp_type`/`icmp_code` - An ICMP/ICMPv6 [type/code][icmp] to match.
        :param pulumi.Input[str] description: A free-form text describing the security group rule.
        :param pulumi.Input[str] protocol: The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        :param pulumi.Input[str] security_group: The parent security group name. Please use the `security_group_id` argument along the SecurityGroup data source instead.
        :param pulumi.Input[str] security_group_id: The parent SecurityGroup ID.
        :param pulumi.Input[str] type: The traffic direction to match (`INGRESS` or `EGRESS`).
        :param pulumi.Input[str] user_security_group: An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`user_security_group_id`). Please use the `user_security_group_id` argument along the SecurityGroup data source instead.
        :param pulumi.Input[str] user_security_group_id: An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`user_security_group)`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityGroupRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage Exoscale [Security Group](https://community.exoscale.com/documentation/compute/security-groups/) Rules.

        ## Usage

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_security_group = exoscale.SecurityGroup("mySecurityGroup")
        my_security_group_rule = exoscale.SecurityGroupRule("mySecurityGroupRule",
            security_group_id=my_security_group.id,
            type="INGRESS",
            protocol="TCP",
            cidr="0.0.0.0/0",
            start_port=80,
            end_port=80)
        ```

        ## Import

        An existing security group rule may be imported by `<security-group-ID>/<security-group-rule-ID>`console

        ```sh
         $ pulumi import exoscale:index/securityGroupRule:SecurityGroupRule \\
        ```

         exoscale_security_group_rule.my_security_group_rule \\

         f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524

        :param str resource_name: The name of the resource.
        :param SecurityGroupRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityGroupRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_port: Optional[pulumi.Input[int]] = None,
                 icmp_code: Optional[pulumi.Input[int]] = None,
                 icmp_type: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 start_port: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_security_group: Optional[pulumi.Input[str]] = None,
                 user_security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityGroupRuleArgs.__new__(SecurityGroupRuleArgs)

            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["description"] = description
            __props__.__dict__["end_port"] = end_port
            __props__.__dict__["icmp_code"] = icmp_code
            __props__.__dict__["icmp_type"] = icmp_type
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["security_group"] = security_group
            __props__.__dict__["security_group_id"] = security_group_id
            __props__.__dict__["start_port"] = start_port
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["user_security_group"] = user_security_group
            __props__.__dict__["user_security_group_id"] = user_security_group_id
        super(SecurityGroupRule, __self__).__init__(
            'exoscale:index/securityGroupRule:SecurityGroupRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            end_port: Optional[pulumi.Input[int]] = None,
            icmp_code: Optional[pulumi.Input[int]] = None,
            icmp_type: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            security_group: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            start_port: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user_security_group: Optional[pulumi.Input[str]] = None,
            user_security_group_id: Optional[pulumi.Input[str]] = None) -> 'SecurityGroupRule':
        """
        Get an existing SecurityGroupRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation][cidr]) to match (conflicts with `user_security_group`/`user_security_group_id`).
               * `start_port`/`end_port` - A `TCP`/`UDP` port range to match.
               * `icmp_type`/`icmp_code` - An ICMP/ICMPv6 [type/code][icmp] to match.
        :param pulumi.Input[str] description: A free-form text describing the security group rule.
        :param pulumi.Input[str] protocol: The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        :param pulumi.Input[str] security_group: The parent security group name. Please use the `security_group_id` argument along the SecurityGroup data source instead.
        :param pulumi.Input[str] security_group_id: The parent SecurityGroup ID.
        :param pulumi.Input[str] type: The traffic direction to match (`INGRESS` or `EGRESS`).
        :param pulumi.Input[str] user_security_group: An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`user_security_group_id`). Please use the `user_security_group_id` argument along the SecurityGroup data source instead.
        :param pulumi.Input[str] user_security_group_id: An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`user_security_group)`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityGroupRuleState.__new__(_SecurityGroupRuleState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["description"] = description
        __props__.__dict__["end_port"] = end_port
        __props__.__dict__["icmp_code"] = icmp_code
        __props__.__dict__["icmp_type"] = icmp_type
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["security_group"] = security_group
        __props__.__dict__["security_group_id"] = security_group_id
        __props__.__dict__["start_port"] = start_port
        __props__.__dict__["type"] = type
        __props__.__dict__["user_security_group"] = user_security_group
        __props__.__dict__["user_security_group_id"] = user_security_group_id
        return SecurityGroupRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[Optional[str]]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination IP subnet (in [CIDR notation][cidr]) to match (conflicts with `user_security_group`/`user_security_group_id`).
        * `start_port`/`end_port` - A `TCP`/`UDP` port range to match.
        * `icmp_type`/`icmp_code` - An ICMP/ICMPv6 [type/code][icmp] to match.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A free-form text describing the security group rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "end_port")

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "icmp_code")

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "icmp_type")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        """
        The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`)
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> pulumi.Output[str]:
        """
        The parent security group name. Please use the `security_group_id` argument along the SecurityGroup data source instead.
        """
        return pulumi.get(self, "security_group")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The parent SecurityGroup ID.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "start_port")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The traffic direction to match (`INGRESS` or `EGRESS`).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userSecurityGroup")
    def user_security_group(self) -> pulumi.Output[str]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination security group name to match (conflicts with `cidr`/`user_security_group_id`). Please use the `user_security_group_id` argument along the SecurityGroup data source instead.
        """
        return pulumi.get(self, "user_security_group")

    @property
    @pulumi.getter(name="userSecurityGroupId")
    def user_security_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        An (`INGRESS`) source / (`EGRESS`) destination security group ID to match (conflicts with `cidr`/`user_security_group)`).
        """
        return pulumi.get(self, "user_security_group_id")

