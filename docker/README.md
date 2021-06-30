provider-gcp does not support higher googleapi version util https://github.com/crossplane/provider-gcp/pull/308

But we need to use higher google api to support workload identity, so we have to upgrade google api version to 0.47.0

It is caused that the `gke` and `pubsub` do not works well, so we have to build image by our docker file.

# Build
```shell
 GOOS=linux go build -o docker/crossplane-gcp-provider cmd/provider/main.go
```