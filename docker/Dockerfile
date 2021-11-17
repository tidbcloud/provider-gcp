FROM alpine:3.13.7 as base

# this fils is generate by go build
COPY crossplane-gcp-provider  /crossplane-gcp-provider
ENTRYPOINT ["/crossplane-gcp-provider"]