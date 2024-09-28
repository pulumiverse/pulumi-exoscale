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

__all__ = [
    'GetDatabaseUriResult',
    'AwaitableGetDatabaseUriResult',
    'get_database_uri',
    'get_database_uri_output',
]

@pulumi.output_type
class GetDatabaseUriResult:
    """
    A collection of values returned by getDatabaseUri.
    """
    def __init__(__self__, db_name=None, host=None, id=None, name=None, password=None, port=None, schema=None, timeouts=None, type=None, uri=None, username=None, zone=None):
        if db_name and not isinstance(db_name, str):
            raise TypeError("Expected argument 'db_name' to be a str")
        pulumi.set(__self__, "db_name", db_name)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if schema and not isinstance(schema, str):
            raise TypeError("Expected argument 'schema' to be a str")
        pulumi.set(__self__, "schema", schema)
        if timeouts and not isinstance(timeouts, dict):
            raise TypeError("Expected argument 'timeouts' to be a dict")
        pulumi.set(__self__, "timeouts", timeouts)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uri and not isinstance(uri, str):
            raise TypeError("Expected argument 'uri' to be a str")
        pulumi.set(__self__, "uri", uri)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> str:
        """
        Default database name
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        Database service hostname
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of database service to match.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Admin user password
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Database service port
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def schema(self) -> str:
        """
        Database service connection schema
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter
    def timeouts(self) -> Optional['outputs.GetDatabaseUriTimeoutsResult']:
        return pulumi.get(self, "timeouts")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        Database service connection URI.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        Admin user username
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The Exoscale Zone name.
        """
        return pulumi.get(self, "zone")


class AwaitableGetDatabaseUriResult(GetDatabaseUriResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseUriResult(
            db_name=self.db_name,
            host=self.host,
            id=self.id,
            name=self.name,
            password=self.password,
            port=self.port,
            schema=self.schema,
            timeouts=self.timeouts,
            type=self.type,
            uri=self.uri,
            username=self.username,
            zone=self.zone)


def get_database_uri(name: Optional[str] = None,
                     timeouts: Optional[Union['GetDatabaseUriTimeoutsArgs', 'GetDatabaseUriTimeoutsArgsDict']] = None,
                     type: Optional[str] = None,
                     zone: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseUriResult:
    """
    ## Example Usage


    :param str name: Name of database service to match.
    :param str type: The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
    :param str zone: The Exoscale Zone name.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['timeouts'] = timeouts
    __args__['type'] = type
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getDatabaseUri:getDatabaseUri', __args__, opts=opts, typ=GetDatabaseUriResult).value

    return AwaitableGetDatabaseUriResult(
        db_name=pulumi.get(__ret__, 'db_name'),
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        password=pulumi.get(__ret__, 'password'),
        port=pulumi.get(__ret__, 'port'),
        schema=pulumi.get(__ret__, 'schema'),
        timeouts=pulumi.get(__ret__, 'timeouts'),
        type=pulumi.get(__ret__, 'type'),
        uri=pulumi.get(__ret__, 'uri'),
        username=pulumi.get(__ret__, 'username'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_database_uri)
def get_database_uri_output(name: Optional[pulumi.Input[str]] = None,
                            timeouts: Optional[pulumi.Input[Optional[Union['GetDatabaseUriTimeoutsArgs', 'GetDatabaseUriTimeoutsArgsDict']]]] = None,
                            type: Optional[pulumi.Input[str]] = None,
                            zone: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseUriResult]:
    """
    ## Example Usage


    :param str name: Name of database service to match.
    :param str type: The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
    :param str zone: The Exoscale Zone name.
    """
    ...
