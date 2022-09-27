import * as pulumi from "@pulumi/pulumi";
import * as exoscale from "@pulumiverse/exoscale"


let cluster = new exoscale.SKSCluster("cluster", {
  zone: "ch-gva-2",
  name: "my-sks-cluster"
})

export const endpoint = cluster.endpoint
