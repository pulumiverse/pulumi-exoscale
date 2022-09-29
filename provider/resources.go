// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exoscale

import (
	"fmt"
	"path/filepath"

	"github.com/exoscale/terraform-provider-exoscale/exoscale"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumiverse/pulumi-exoscale/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "exoscale"
	// modules:
	mainMod = "index" // the exoscale module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(exoscale.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "exoscale",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Exoscale",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://avatars3.githubusercontent.com/exoscale",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/pulumiverse/pulumi-exoscale",
		Description:       "A Pulumi package for creating and managing exoscale cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "exoscale", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/pulumiverse/pulumi-exoscale",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "exoscale",
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
			"config": {
				CSharpName: "CloudStackIniConfig",
			},
			"timeout": {
				Type: "integer",
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"exoscale_anti_affinity_group": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AntiAffinityGroup")},
			"exoscale_compute_instance":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ComputeInstance")},
			"exoscale_database":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Database")},
			"exoscale_domain":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Domain")},
			"exoscale_domain_record":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DomainRecord")},
			"exoscale_elastic_ip":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ElasticIP")},
			"exoscale_iam_access_key":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IAMAccessKey")},
			"exoscale_instance_pool":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstancePool")},
			"exoscale_nlb":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NLB")},
			"exoscale_nlb_service":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NLBService")},
			"exoscale_private_network":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PrivateNetwork")},
			"exoscale_security_group":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SecurityGroup")},
			"exoscale_security_group_rule": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SecurityGroupRule")},
			"exoscale_sks_cluster":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SKSCluster")},
			"exoscale_sks_kubeconfig":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SKSKubeconfig")},
			"exoscale_sks_nodepool":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SKSNodepool")},
			"exoscale_ssh_key":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SSHKey")},
			// Deprecated Ressources, will be removed with next major version
			"exoscale_affinity":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Affinity")},
			"exoscale_compute":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Compute")},
			"exoscale_ipaddress":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPAddress")},
			"exoscale_network":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Network")},
			"exoscale_nic":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NIC")},
			"exoscale_secondary_ipaddress":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SecondaryIPAddress")},
			"exoscale_security_group_rules": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SecurityGroupRules")},
			"exoscale_ssh_keypair":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SSHKeypair")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"exoscale_anti_affinity_group":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAntiAffinityGroup")},
			"exoscale_compute_instance":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getComputeInstance")},
			"exoscale_compute_instance_list": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getComputeInstanceList")},
			"exoscale_compute_template":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getComputeTemplate")},
			"exoscale_domain":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDomain")},
			"exoscale_domain_record":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDomainRecord")},
			"exoscale_elastic_ip":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getElasticIP")},
			"exoscale_instance_pool":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstancePool")},
			"exoscale_instance_pool_list":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstancePoolList")},
			"exoscale_nlb":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNLB")},
			"exoscale_private_network":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPrivateNetwork")},
			"exoscale_security_group":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSecurityGroup")},
			// Deprecated DataSources, will be removed with next major version
			"exoscale_affinity":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAffinity")},
			"exoscale_compute":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCompute")},
			"exoscale_compute_ipaddress": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getComputeIPAddress")},
			"exoscale_network":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetwork")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/exoscale",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_exoscale",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumiverse",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
