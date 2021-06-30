# provider-gcp

## Overview

This `provider-gcp` repository is forked from [crossplane](https://github.com/crossplane/provider-gcp).
We just start some controllers to avoid to apply all CRD.
Currently, [infra-cd](https://github.com/tidbcloud/infra-cd) apply following crds:
```antlrv4
compute.gcp.crossplane.io_networks.yaml
gcp.crossplane.io_providerconfigs.yaml
gcp.crossplane.io_providerconfigusages.yaml
iam.gcp.crossplane.io_serviceaccountpolicies.yaml
iam.gcp.crossplane.io_serviceaccounts.yaml
storage.gcp.crossplane.io_bucketpolicies.yaml
storage.gcp.crossplane.io_bucketpolicymembers.yaml
storage.gcp.crossplane.io_buckets.yaml
dns.gcp.crossplane.io_managedzones.yaml
dns.gcp.crossplane.io_resourcerecordsets.yaml
vpcpeering.gcp.crossplane.io_peerings.yaml
```
*Note*

For supporting gcp impersonation auth, we have to upgrade google api version. The high google api version is not incompatible with low version, it impact gke cluster controller.
Currently , we just remove changed field to make lint successful. 

# Build Docker Image
## Build binary
```shell
GOOS=linux go build -o docker/crossplane-gcp-provider cmd/provider/main.go
```
## Build image
```shell
cd docker
docker build -t $REGISTRY/provider-gcp:v0.19.2-dev .
```