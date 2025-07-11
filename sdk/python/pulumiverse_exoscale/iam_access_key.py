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

__all__ = ['IamAccessKeyArgs', 'IamAccessKey']

@pulumi.input_type
class IamAccessKeyArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 operations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a IamAccessKey resource.
        :param pulumi.Input[builtins.str] name: ❗ The IAM access key name.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] operations: ❗ A list of API operations to restrict the key to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] resources: ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: ❗ A list of tags to restrict the key to.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operations is not None:
            pulumi.set(__self__, "operations", operations)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ❗ The IAM access key name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def operations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        ❗ A list of API operations to restrict the key to.
        """
        return pulumi.get(self, "operations")

    @operations.setter
    def operations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "operations", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        ❗ A list of tags to restrict the key to.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _IamAccessKeyState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 operations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 secret: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags_operations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering IamAccessKey resources.
        :param pulumi.Input[builtins.str] key: The IAM access key (identifier).
        :param pulumi.Input[builtins.str] name: ❗ The IAM access key name.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] operations: ❗ A list of API operations to restrict the key to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] resources: ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
        :param pulumi.Input[builtins.str] secret: The key secret.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: ❗ A list of tags to restrict the key to.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operations is not None:
            pulumi.set(__self__, "operations", operations)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_operations is not None:
            pulumi.set(__self__, "tags_operations", tags_operations)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IAM access key (identifier).
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ❗ The IAM access key name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def operations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        ❗ A list of API operations to restrict the key to.
        """
        return pulumi.get(self, "operations")

    @operations.setter
    def operations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "operations", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The key secret.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        ❗ A list of tags to restrict the key to.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsOperations")
    def tags_operations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags_operations")

    @tags_operations.setter
    def tags_operations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags_operations", value)


@pulumi.type_token("exoscale:index/iamAccessKey:IamAccessKey")
class IamAccessKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 operations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_sos_access_key = exoscale.IamAccessKey("mySosAccessKey",
            operations=[
                "get-sos-object",
                "list-sos-bucket",
            ],
            resources=["sos/bucket:my-bucket"])
        my_sks_access_key = exoscale.IamAccessKey("mySksAccessKey", tags=["sks"])
        ```

        Please refer to the examples
        directory for complete configuration examples.

        > **NOTE:** You can retrieve the list of available operations and tags using the [Exoscale CLI](https://github.com/exoscale/cli/): `exo iam access-key list-operations`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: ❗ The IAM access key name.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] operations: ❗ A list of API operations to restrict the key to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] resources: ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: ❗ A list of tags to restrict the key to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IamAccessKeyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_sos_access_key = exoscale.IamAccessKey("mySosAccessKey",
            operations=[
                "get-sos-object",
                "list-sos-bucket",
            ],
            resources=["sos/bucket:my-bucket"])
        my_sks_access_key = exoscale.IamAccessKey("mySksAccessKey", tags=["sks"])
        ```

        Please refer to the examples
        directory for complete configuration examples.

        > **NOTE:** You can retrieve the list of available operations and tags using the [Exoscale CLI](https://github.com/exoscale/cli/): `exo iam access-key list-operations`.

        :param str resource_name: The name of the resource.
        :param IamAccessKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamAccessKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 operations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamAccessKeyArgs.__new__(IamAccessKeyArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["operations"] = operations
            __props__.__dict__["resources"] = resources
            __props__.__dict__["tags"] = tags
            __props__.__dict__["key"] = None
            __props__.__dict__["secret"] = None
            __props__.__dict__["tags_operations"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["key", "secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(IamAccessKey, __self__).__init__(
            'exoscale:index/iamAccessKey:IamAccessKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            operations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            secret: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            tags_operations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'IamAccessKey':
        """
        Get an existing IamAccessKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] key: The IAM access key (identifier).
        :param pulumi.Input[builtins.str] name: ❗ The IAM access key name.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] operations: ❗ A list of API operations to restrict the key to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] resources: ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
        :param pulumi.Input[builtins.str] secret: The key secret.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: ❗ A list of tags to restrict the key to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamAccessKeyState.__new__(_IamAccessKeyState)

        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["operations"] = operations
        __props__.__dict__["resources"] = resources
        __props__.__dict__["secret"] = secret
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_operations"] = tags_operations
        return IamAccessKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[builtins.str]:
        """
        The IAM access key (identifier).
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        ❗ The IAM access key name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def operations(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        ❗ A list of API operations to restrict the key to.
        """
        return pulumi.get(self, "operations")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[builtins.str]:
        """
        The key secret.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        ❗ A list of tags to restrict the key to.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsOperations")
    def tags_operations(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "tags_operations")

