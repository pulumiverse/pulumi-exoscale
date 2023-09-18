// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { DatabaseArgs, DatabaseState } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { GetDatabaseURIArgs, GetDatabaseURIResult, GetDatabaseURIOutputArgs } from "./getDatabaseURI";
export const getDatabaseURI: typeof import("./getDatabaseURI").getDatabaseURI = null as any;
export const getDatabaseURIOutput: typeof import("./getDatabaseURI").getDatabaseURIOutput = null as any;
utilities.lazyLoad(exports, ["getDatabaseURI","getDatabaseURIOutput"], () => require("./getDatabaseURI"));

export { GetNLBServiceListArgs, GetNLBServiceListResult, GetNLBServiceListOutputArgs } from "./getNLBServiceList";
export const getNLBServiceList: typeof import("./getNLBServiceList").getNLBServiceList = null as any;
export const getNLBServiceListOutput: typeof import("./getNLBServiceList").getNLBServiceListOutput = null as any;
utilities.lazyLoad(exports, ["getNLBServiceList","getNLBServiceListOutput"], () => require("./getNLBServiceList"));

export { GetZonesResult } from "./getZones";
export const getZones: typeof import("./getZones").getZones = null as any;
export const getZonesOutput: typeof import("./getZones").getZonesOutput = null as any;
utilities.lazyLoad(exports, ["getZones","getZonesOutput"], () => require("./getZones"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "exoscale:index/database:Database":
                return new Database(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("exoscale", "index/database", _module)
pulumi.runtime.registerResourcePackage("exoscale", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:exoscale") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
