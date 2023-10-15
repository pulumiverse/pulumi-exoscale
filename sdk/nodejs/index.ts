// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { AffinityArgs, AffinityState } from "./affinity";
export type Affinity = import("./affinity").Affinity;
export const Affinity: typeof import("./affinity").Affinity = null as any;
utilities.lazyLoad(exports, ["Affinity"], () => require("./affinity"));

export { AntiAffinityGroupArgs, AntiAffinityGroupState } from "./antiAffinityGroup";
export type AntiAffinityGroup = import("./antiAffinityGroup").AntiAffinityGroup;
export const AntiAffinityGroup: typeof import("./antiAffinityGroup").AntiAffinityGroup = null as any;
utilities.lazyLoad(exports, ["AntiAffinityGroup"], () => require("./antiAffinityGroup"));

export { ComputeArgs, ComputeState } from "./compute";
export type Compute = import("./compute").Compute;
export const Compute: typeof import("./compute").Compute = null as any;
utilities.lazyLoad(exports, ["Compute"], () => require("./compute"));

export { ComputeInstanceArgs, ComputeInstanceState } from "./computeInstance";
export type ComputeInstance = import("./computeInstance").ComputeInstance;
export const ComputeInstance: typeof import("./computeInstance").ComputeInstance = null as any;
utilities.lazyLoad(exports, ["ComputeInstance"], () => require("./computeInstance"));

export { DatabaseArgs, DatabaseState } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { DomainArgs, DomainState } from "./domain";
export type Domain = import("./domain").Domain;
export const Domain: typeof import("./domain").Domain = null as any;
utilities.lazyLoad(exports, ["Domain"], () => require("./domain"));

export { DomainRecordArgs, DomainRecordState } from "./domainRecord";
export type DomainRecord = import("./domainRecord").DomainRecord;
export const DomainRecord: typeof import("./domainRecord").DomainRecord = null as any;
utilities.lazyLoad(exports, ["DomainRecord"], () => require("./domainRecord"));

export { ElasticIPArgs, ElasticIPState } from "./elasticIP";
export type ElasticIP = import("./elasticIP").ElasticIP;
export const ElasticIP: typeof import("./elasticIP").ElasticIP = null as any;
utilities.lazyLoad(exports, ["ElasticIP"], () => require("./elasticIP"));

export { GetAffinityArgs, GetAffinityResult, GetAffinityOutputArgs } from "./getAffinity";
export const getAffinity: typeof import("./getAffinity").getAffinity = null as any;
export const getAffinityOutput: typeof import("./getAffinity").getAffinityOutput = null as any;
utilities.lazyLoad(exports, ["getAffinity","getAffinityOutput"], () => require("./getAffinity"));

export { GetAntiAffinityGroupArgs, GetAntiAffinityGroupResult, GetAntiAffinityGroupOutputArgs } from "./getAntiAffinityGroup";
export const getAntiAffinityGroup: typeof import("./getAntiAffinityGroup").getAntiAffinityGroup = null as any;
export const getAntiAffinityGroupOutput: typeof import("./getAntiAffinityGroup").getAntiAffinityGroupOutput = null as any;
utilities.lazyLoad(exports, ["getAntiAffinityGroup","getAntiAffinityGroupOutput"], () => require("./getAntiAffinityGroup"));

export { GetComputeArgs, GetComputeResult, GetComputeOutputArgs } from "./getCompute";
export const getCompute: typeof import("./getCompute").getCompute = null as any;
export const getComputeOutput: typeof import("./getCompute").getComputeOutput = null as any;
utilities.lazyLoad(exports, ["getCompute","getComputeOutput"], () => require("./getCompute"));

export { GetComputeIPAddressArgs, GetComputeIPAddressResult, GetComputeIPAddressOutputArgs } from "./getComputeIPAddress";
export const getComputeIPAddress: typeof import("./getComputeIPAddress").getComputeIPAddress = null as any;
export const getComputeIPAddressOutput: typeof import("./getComputeIPAddress").getComputeIPAddressOutput = null as any;
utilities.lazyLoad(exports, ["getComputeIPAddress","getComputeIPAddressOutput"], () => require("./getComputeIPAddress"));

export { GetComputeInstanceArgs, GetComputeInstanceResult, GetComputeInstanceOutputArgs } from "./getComputeInstance";
export const getComputeInstance: typeof import("./getComputeInstance").getComputeInstance = null as any;
export const getComputeInstanceOutput: typeof import("./getComputeInstance").getComputeInstanceOutput = null as any;
utilities.lazyLoad(exports, ["getComputeInstance","getComputeInstanceOutput"], () => require("./getComputeInstance"));

export { GetComputeInstanceListArgs, GetComputeInstanceListResult, GetComputeInstanceListOutputArgs } from "./getComputeInstanceList";
export const getComputeInstanceList: typeof import("./getComputeInstanceList").getComputeInstanceList = null as any;
export const getComputeInstanceListOutput: typeof import("./getComputeInstanceList").getComputeInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getComputeInstanceList","getComputeInstanceListOutput"], () => require("./getComputeInstanceList"));

export { GetComputeTemplateArgs, GetComputeTemplateResult, GetComputeTemplateOutputArgs } from "./getComputeTemplate";
export const getComputeTemplate: typeof import("./getComputeTemplate").getComputeTemplate = null as any;
export const getComputeTemplateOutput: typeof import("./getComputeTemplate").getComputeTemplateOutput = null as any;
utilities.lazyLoad(exports, ["getComputeTemplate","getComputeTemplateOutput"], () => require("./getComputeTemplate"));

export { GetDatabaseURIArgs, GetDatabaseURIResult, GetDatabaseURIOutputArgs } from "./getDatabaseURI";
export const getDatabaseURI: typeof import("./getDatabaseURI").getDatabaseURI = null as any;
export const getDatabaseURIOutput: typeof import("./getDatabaseURI").getDatabaseURIOutput = null as any;
utilities.lazyLoad(exports, ["getDatabaseURI","getDatabaseURIOutput"], () => require("./getDatabaseURI"));

export { GetDomainArgs, GetDomainResult, GetDomainOutputArgs } from "./getDomain";
export const getDomain: typeof import("./getDomain").getDomain = null as any;
export const getDomainOutput: typeof import("./getDomain").getDomainOutput = null as any;
utilities.lazyLoad(exports, ["getDomain","getDomainOutput"], () => require("./getDomain"));

export { GetDomainRecordArgs, GetDomainRecordResult, GetDomainRecordOutputArgs } from "./getDomainRecord";
export const getDomainRecord: typeof import("./getDomainRecord").getDomainRecord = null as any;
export const getDomainRecordOutput: typeof import("./getDomainRecord").getDomainRecordOutput = null as any;
utilities.lazyLoad(exports, ["getDomainRecord","getDomainRecordOutput"], () => require("./getDomainRecord"));

export { GetElasticIPArgs, GetElasticIPResult, GetElasticIPOutputArgs } from "./getElasticIP";
export const getElasticIP: typeof import("./getElasticIP").getElasticIP = null as any;
export const getElasticIPOutput: typeof import("./getElasticIP").getElasticIPOutput = null as any;
utilities.lazyLoad(exports, ["getElasticIP","getElasticIPOutput"], () => require("./getElasticIP"));

export { GetIAMAPIKeyArgs, GetIAMAPIKeyResult, GetIAMAPIKeyOutputArgs } from "./getIAMAPIKey";
export const getIAMAPIKey: typeof import("./getIAMAPIKey").getIAMAPIKey = null as any;
export const getIAMAPIKeyOutput: typeof import("./getIAMAPIKey").getIAMAPIKeyOutput = null as any;
utilities.lazyLoad(exports, ["getIAMAPIKey","getIAMAPIKeyOutput"], () => require("./getIAMAPIKey"));

export { GetIAMOrgPolicyArgs, GetIAMOrgPolicyResult, GetIAMOrgPolicyOutputArgs } from "./getIAMOrgPolicy";
export const getIAMOrgPolicy: typeof import("./getIAMOrgPolicy").getIAMOrgPolicy = null as any;
export const getIAMOrgPolicyOutput: typeof import("./getIAMOrgPolicy").getIAMOrgPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getIAMOrgPolicy","getIAMOrgPolicyOutput"], () => require("./getIAMOrgPolicy"));

export { GetIAMRoleArgs, GetIAMRoleResult, GetIAMRoleOutputArgs } from "./getIAMRole";
export const getIAMRole: typeof import("./getIAMRole").getIAMRole = null as any;
export const getIAMRoleOutput: typeof import("./getIAMRole").getIAMRoleOutput = null as any;
utilities.lazyLoad(exports, ["getIAMRole","getIAMRoleOutput"], () => require("./getIAMRole"));

export { GetInstancePoolArgs, GetInstancePoolResult, GetInstancePoolOutputArgs } from "./getInstancePool";
export const getInstancePool: typeof import("./getInstancePool").getInstancePool = null as any;
export const getInstancePoolOutput: typeof import("./getInstancePool").getInstancePoolOutput = null as any;
utilities.lazyLoad(exports, ["getInstancePool","getInstancePoolOutput"], () => require("./getInstancePool"));

export { GetInstancePoolListArgs, GetInstancePoolListResult, GetInstancePoolListOutputArgs } from "./getInstancePoolList";
export const getInstancePoolList: typeof import("./getInstancePoolList").getInstancePoolList = null as any;
export const getInstancePoolListOutput: typeof import("./getInstancePoolList").getInstancePoolListOutput = null as any;
utilities.lazyLoad(exports, ["getInstancePoolList","getInstancePoolListOutput"], () => require("./getInstancePoolList"));

export { GetNLBArgs, GetNLBResult, GetNLBOutputArgs } from "./getNLB";
export const getNLB: typeof import("./getNLB").getNLB = null as any;
export const getNLBOutput: typeof import("./getNLB").getNLBOutput = null as any;
utilities.lazyLoad(exports, ["getNLB","getNLBOutput"], () => require("./getNLB"));

export { GetNLBServiceListArgs, GetNLBServiceListResult, GetNLBServiceListOutputArgs } from "./getNLBServiceList";
export const getNLBServiceList: typeof import("./getNLBServiceList").getNLBServiceList = null as any;
export const getNLBServiceListOutput: typeof import("./getNLBServiceList").getNLBServiceListOutput = null as any;
utilities.lazyLoad(exports, ["getNLBServiceList","getNLBServiceListOutput"], () => require("./getNLBServiceList"));

export { GetNetworkArgs, GetNetworkResult, GetNetworkOutputArgs } from "./getNetwork";
export const getNetwork: typeof import("./getNetwork").getNetwork = null as any;
export const getNetworkOutput: typeof import("./getNetwork").getNetworkOutput = null as any;
utilities.lazyLoad(exports, ["getNetwork","getNetworkOutput"], () => require("./getNetwork"));

export { GetPrivateNetworkArgs, GetPrivateNetworkResult, GetPrivateNetworkOutputArgs } from "./getPrivateNetwork";
export const getPrivateNetwork: typeof import("./getPrivateNetwork").getPrivateNetwork = null as any;
export const getPrivateNetworkOutput: typeof import("./getPrivateNetwork").getPrivateNetworkOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateNetwork","getPrivateNetworkOutput"], () => require("./getPrivateNetwork"));

export { GetSKSClusterArgs, GetSKSClusterResult, GetSKSClusterOutputArgs } from "./getSKSCluster";
export const getSKSCluster: typeof import("./getSKSCluster").getSKSCluster = null as any;
export const getSKSClusterOutput: typeof import("./getSKSCluster").getSKSClusterOutput = null as any;
utilities.lazyLoad(exports, ["getSKSCluster","getSKSClusterOutput"], () => require("./getSKSCluster"));

export { GetSKSClusterListArgs, GetSKSClusterListResult, GetSKSClusterListOutputArgs } from "./getSKSClusterList";
export const getSKSClusterList: typeof import("./getSKSClusterList").getSKSClusterList = null as any;
export const getSKSClusterListOutput: typeof import("./getSKSClusterList").getSKSClusterListOutput = null as any;
utilities.lazyLoad(exports, ["getSKSClusterList","getSKSClusterListOutput"], () => require("./getSKSClusterList"));

export { GetSKSNodepoolArgs, GetSKSNodepoolResult, GetSKSNodepoolOutputArgs } from "./getSKSNodepool";
export const getSKSNodepool: typeof import("./getSKSNodepool").getSKSNodepool = null as any;
export const getSKSNodepoolOutput: typeof import("./getSKSNodepool").getSKSNodepoolOutput = null as any;
utilities.lazyLoad(exports, ["getSKSNodepool","getSKSNodepoolOutput"], () => require("./getSKSNodepool"));

export { GetSKSNodepoolListArgs, GetSKSNodepoolListResult, GetSKSNodepoolListOutputArgs } from "./getSKSNodepoolList";
export const getSKSNodepoolList: typeof import("./getSKSNodepoolList").getSKSNodepoolList = null as any;
export const getSKSNodepoolListOutput: typeof import("./getSKSNodepoolList").getSKSNodepoolListOutput = null as any;
utilities.lazyLoad(exports, ["getSKSNodepoolList","getSKSNodepoolListOutput"], () => require("./getSKSNodepoolList"));

export { GetSecurityGroupArgs, GetSecurityGroupResult, GetSecurityGroupOutputArgs } from "./getSecurityGroup";
export const getSecurityGroup: typeof import("./getSecurityGroup").getSecurityGroup = null as any;
export const getSecurityGroupOutput: typeof import("./getSecurityGroup").getSecurityGroupOutput = null as any;
utilities.lazyLoad(exports, ["getSecurityGroup","getSecurityGroupOutput"], () => require("./getSecurityGroup"));

export { GetTemplateArgs, GetTemplateResult, GetTemplateOutputArgs } from "./getTemplate";
export const getTemplate: typeof import("./getTemplate").getTemplate = null as any;
export const getTemplateOutput: typeof import("./getTemplate").getTemplateOutput = null as any;
utilities.lazyLoad(exports, ["getTemplate","getTemplateOutput"], () => require("./getTemplate"));

export { GetZonesResult } from "./getZones";
export const getZones: typeof import("./getZones").getZones = null as any;
export const getZonesOutput: typeof import("./getZones").getZonesOutput = null as any;
utilities.lazyLoad(exports, ["getZones","getZonesOutput"], () => require("./getZones"));

export { IAMAccessKeyArgs, IAMAccessKeyState } from "./iamaccessKey";
export type IAMAccessKey = import("./iamaccessKey").IAMAccessKey;
export const IAMAccessKey: typeof import("./iamaccessKey").IAMAccessKey = null as any;
utilities.lazyLoad(exports, ["IAMAccessKey"], () => require("./iamaccessKey"));

export { IAMAPIKeyArgs, IAMAPIKeyState } from "./iamapikey";
export type IAMAPIKey = import("./iamapikey").IAMAPIKey;
export const IAMAPIKey: typeof import("./iamapikey").IAMAPIKey = null as any;
utilities.lazyLoad(exports, ["IAMAPIKey"], () => require("./iamapikey"));

export { IAMOrgPolicyArgs, IAMOrgPolicyState } from "./iamorgPolicy";
export type IAMOrgPolicy = import("./iamorgPolicy").IAMOrgPolicy;
export const IAMOrgPolicy: typeof import("./iamorgPolicy").IAMOrgPolicy = null as any;
utilities.lazyLoad(exports, ["IAMOrgPolicy"], () => require("./iamorgPolicy"));

export { IAMRoleArgs, IAMRoleState } from "./iamrole";
export type IAMRole = import("./iamrole").IAMRole;
export const IAMRole: typeof import("./iamrole").IAMRole = null as any;
utilities.lazyLoad(exports, ["IAMRole"], () => require("./iamrole"));

export { InstancePoolArgs, InstancePoolState } from "./instancePool";
export type InstancePool = import("./instancePool").InstancePool;
export const InstancePool: typeof import("./instancePool").InstancePool = null as any;
utilities.lazyLoad(exports, ["InstancePool"], () => require("./instancePool"));

export { IPAddressArgs, IPAddressState } from "./ipaddress";
export type IPAddress = import("./ipaddress").IPAddress;
export const IPAddress: typeof import("./ipaddress").IPAddress = null as any;
utilities.lazyLoad(exports, ["IPAddress"], () => require("./ipaddress"));

export { NetworkArgs, NetworkState } from "./network";
export type Network = import("./network").Network;
export const Network: typeof import("./network").Network = null as any;
utilities.lazyLoad(exports, ["Network"], () => require("./network"));

export { NICArgs, NICState } from "./nic";
export type NIC = import("./nic").NIC;
export const NIC: typeof import("./nic").NIC = null as any;
utilities.lazyLoad(exports, ["NIC"], () => require("./nic"));

export { NLBArgs, NLBState } from "./nlb";
export type NLB = import("./nlb").NLB;
export const NLB: typeof import("./nlb").NLB = null as any;
utilities.lazyLoad(exports, ["NLB"], () => require("./nlb"));

export { NLBServiceArgs, NLBServiceState } from "./nlbservice";
export type NLBService = import("./nlbservice").NLBService;
export const NLBService: typeof import("./nlbservice").NLBService = null as any;
utilities.lazyLoad(exports, ["NLBService"], () => require("./nlbservice"));

export { PrivateNetworkArgs, PrivateNetworkState } from "./privateNetwork";
export type PrivateNetwork = import("./privateNetwork").PrivateNetwork;
export const PrivateNetwork: typeof import("./privateNetwork").PrivateNetwork = null as any;
utilities.lazyLoad(exports, ["PrivateNetwork"], () => require("./privateNetwork"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { SecondaryIPAddressArgs, SecondaryIPAddressState } from "./secondaryIPAddress";
export type SecondaryIPAddress = import("./secondaryIPAddress").SecondaryIPAddress;
export const SecondaryIPAddress: typeof import("./secondaryIPAddress").SecondaryIPAddress = null as any;
utilities.lazyLoad(exports, ["SecondaryIPAddress"], () => require("./secondaryIPAddress"));

export { SecurityGroupArgs, SecurityGroupState } from "./securityGroup";
export type SecurityGroup = import("./securityGroup").SecurityGroup;
export const SecurityGroup: typeof import("./securityGroup").SecurityGroup = null as any;
utilities.lazyLoad(exports, ["SecurityGroup"], () => require("./securityGroup"));

export { SecurityGroupRuleArgs, SecurityGroupRuleState } from "./securityGroupRule";
export type SecurityGroupRule = import("./securityGroupRule").SecurityGroupRule;
export const SecurityGroupRule: typeof import("./securityGroupRule").SecurityGroupRule = null as any;
utilities.lazyLoad(exports, ["SecurityGroupRule"], () => require("./securityGroupRule"));

export { SecurityGroupRulesArgs, SecurityGroupRulesState } from "./securityGroupRules";
export type SecurityGroupRules = import("./securityGroupRules").SecurityGroupRules;
export const SecurityGroupRules: typeof import("./securityGroupRules").SecurityGroupRules = null as any;
utilities.lazyLoad(exports, ["SecurityGroupRules"], () => require("./securityGroupRules"));

export { SKSClusterArgs, SKSClusterState } from "./skscluster";
export type SKSCluster = import("./skscluster").SKSCluster;
export const SKSCluster: typeof import("./skscluster").SKSCluster = null as any;
utilities.lazyLoad(exports, ["SKSCluster"], () => require("./skscluster"));

export { SKSKubeconfigArgs, SKSKubeconfigState } from "./skskubeconfig";
export type SKSKubeconfig = import("./skskubeconfig").SKSKubeconfig;
export const SKSKubeconfig: typeof import("./skskubeconfig").SKSKubeconfig = null as any;
utilities.lazyLoad(exports, ["SKSKubeconfig"], () => require("./skskubeconfig"));

export { SKSNodepoolArgs, SKSNodepoolState } from "./sksnodepool";
export type SKSNodepool = import("./sksnodepool").SKSNodepool;
export const SKSNodepool: typeof import("./sksnodepool").SKSNodepool = null as any;
utilities.lazyLoad(exports, ["SKSNodepool"], () => require("./sksnodepool"));

export { SSHKeyArgs, SSHKeyState } from "./sshkey";
export type SSHKey = import("./sshkey").SSHKey;
export const SSHKey: typeof import("./sshkey").SSHKey = null as any;
utilities.lazyLoad(exports, ["SSHKey"], () => require("./sshkey"));

export { SSHKeypairArgs, SSHKeypairState } from "./sshkeypair";
export type SSHKeypair = import("./sshkeypair").SSHKeypair;
export const SSHKeypair: typeof import("./sshkeypair").SSHKeypair = null as any;
utilities.lazyLoad(exports, ["SSHKeypair"], () => require("./sshkeypair"));


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
            case "exoscale:index/affinity:Affinity":
                return new Affinity(name, <any>undefined, { urn })
            case "exoscale:index/antiAffinityGroup:AntiAffinityGroup":
                return new AntiAffinityGroup(name, <any>undefined, { urn })
            case "exoscale:index/compute:Compute":
                return new Compute(name, <any>undefined, { urn })
            case "exoscale:index/computeInstance:ComputeInstance":
                return new ComputeInstance(name, <any>undefined, { urn })
            case "exoscale:index/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "exoscale:index/domain:Domain":
                return new Domain(name, <any>undefined, { urn })
            case "exoscale:index/domainRecord:DomainRecord":
                return new DomainRecord(name, <any>undefined, { urn })
            case "exoscale:index/elasticIP:ElasticIP":
                return new ElasticIP(name, <any>undefined, { urn })
            case "exoscale:index/iAMAPIKey:IAMAPIKey":
                return new IAMAPIKey(name, <any>undefined, { urn })
            case "exoscale:index/iAMAccessKey:IAMAccessKey":
                return new IAMAccessKey(name, <any>undefined, { urn })
            case "exoscale:index/iAMOrgPolicy:IAMOrgPolicy":
                return new IAMOrgPolicy(name, <any>undefined, { urn })
            case "exoscale:index/iAMRole:IAMRole":
                return new IAMRole(name, <any>undefined, { urn })
            case "exoscale:index/iPAddress:IPAddress":
                return new IPAddress(name, <any>undefined, { urn })
            case "exoscale:index/instancePool:InstancePool":
                return new InstancePool(name, <any>undefined, { urn })
            case "exoscale:index/nIC:NIC":
                return new NIC(name, <any>undefined, { urn })
            case "exoscale:index/nLB:NLB":
                return new NLB(name, <any>undefined, { urn })
            case "exoscale:index/nLBService:NLBService":
                return new NLBService(name, <any>undefined, { urn })
            case "exoscale:index/network:Network":
                return new Network(name, <any>undefined, { urn })
            case "exoscale:index/privateNetwork:PrivateNetwork":
                return new PrivateNetwork(name, <any>undefined, { urn })
            case "exoscale:index/sKSCluster:SKSCluster":
                return new SKSCluster(name, <any>undefined, { urn })
            case "exoscale:index/sKSKubeconfig:SKSKubeconfig":
                return new SKSKubeconfig(name, <any>undefined, { urn })
            case "exoscale:index/sKSNodepool:SKSNodepool":
                return new SKSNodepool(name, <any>undefined, { urn })
            case "exoscale:index/sSHKey:SSHKey":
                return new SSHKey(name, <any>undefined, { urn })
            case "exoscale:index/sSHKeypair:SSHKeypair":
                return new SSHKeypair(name, <any>undefined, { urn })
            case "exoscale:index/secondaryIPAddress:SecondaryIPAddress":
                return new SecondaryIPAddress(name, <any>undefined, { urn })
            case "exoscale:index/securityGroup:SecurityGroup":
                return new SecurityGroup(name, <any>undefined, { urn })
            case "exoscale:index/securityGroupRule:SecurityGroupRule":
                return new SecurityGroupRule(name, <any>undefined, { urn })
            case "exoscale:index/securityGroupRules:SecurityGroupRules":
                return new SecurityGroupRules(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("exoscale", "index/affinity", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/antiAffinityGroup", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/compute", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/computeInstance", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/database", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/domain", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/domainRecord", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/elasticIP", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/iAMAPIKey", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/iAMAccessKey", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/iAMOrgPolicy", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/iAMRole", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/iPAddress", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/instancePool", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/nIC", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/nLB", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/nLBService", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/network", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/privateNetwork", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/sKSCluster", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/sKSKubeconfig", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/sKSNodepool", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/sSHKey", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/sSHKeypair", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/secondaryIPAddress", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/securityGroup", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/securityGroupRule", _module)
pulumi.runtime.registerResourceModule("exoscale", "index/securityGroupRules", _module)
pulumi.runtime.registerResourcePackage("exoscale", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:exoscale") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
