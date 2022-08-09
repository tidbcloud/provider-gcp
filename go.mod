module github.com/crossplane/provider-gcp

go 1.13

require (
	cloud.google.com/go v0.81.0
	cloud.google.com/go/pubsub v1.3.1
	cloud.google.com/go/storage v1.14.0
	github.com/crossplane/crossplane-runtime v0.15.1-0.20220315141414-988c9ba9c255
	github.com/crossplane/crossplane-tools v0.0.0-20220310165030-1f43fc12793e
	github.com/google/go-cmp v0.5.6
	github.com/googleapis/gax-go v1.0.3
	github.com/googleapis/gax-go/v2 v2.0.5
	github.com/imdario/mergo v0.3.12
	github.com/mitchellh/copystructure v1.0.0
	github.com/onsi/gomega v1.18.1
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a
	google.golang.org/api v0.47.0
	google.golang.org/genproto v0.0.0-20210831024726-fe130286e0e2
	google.golang.org/grpc v1.41.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/controller-tools v0.8.0
)
