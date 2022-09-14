module github.com/crossplane/provider-gcp

go 1.13

require (
	cloud.google.com/go/pubsub v1.3.1
	cloud.google.com/go/redis v1.6.0
	cloud.google.com/go/storage v1.22.1
	github.com/crossplane/crossplane-runtime v0.17.0
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1 v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.8
	github.com/googleapis/gax-go v1.0.3
	github.com/googleapis/gax-go/v2 v2.4.0
	github.com/imdario/mergo v0.3.12
	github.com/mitchellh/copystructure v1.0.0
	github.com/onsi/gomega v1.19.0
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	google.golang.org/api v0.85.0
	google.golang.org/genproto v0.0.0-20220617124728-180714bec0ad
	google.golang.org/grpc v1.47.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	k8s.io/api v0.25.0
	k8s.io/apimachinery v0.25.0
	k8s.io/client-go v0.25.0
	sigs.k8s.io/controller-runtime v0.13.0
	sigs.k8s.io/controller-tools v0.8.0
)

replace github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1 => ./apis/vpcpeering/v1beta1
