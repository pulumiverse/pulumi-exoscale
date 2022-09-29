"""A Python Pulumi program"""

import pulumi
import pulumiverse_exoscale as exoscale

cluster = exoscale.SKSCluster("cluster",
                              zone="ch-gva-2",
                              name="my-sks-cluster",
                              )

pulumi.export("endpoint", cluster.endpoint)
