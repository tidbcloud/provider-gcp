module github.com/crossplane/provider-gcp

go 1.13

require (
	cloud.google.com/go/pubsub v1.3.1
	cloud.google.com/go/redis v1.4.0
	cloud.google.com/go/storage v1.22.1
	github.com/crossplane/crossplane-runtime v0.19.3
	github.com/crossplane/crossplane-tools v0.0.0-20220310165030-1f43fc12793e
	github.com/crossplane/provider-gcp/apis/dns/v1alpha1 v0.0.0-00010101000000-000000000000
	github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1 v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.9
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/googleapis/gax-go v1.0.3
	github.com/googleapis/gax-go/v2 v2.4.0
	github.com/imdario/mergo v0.3.12
	github.com/mitchellh/copystructure v1.0.0
	github.com/onsi/gomega v1.24.2
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.4.0
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1
	google.golang.org/api v0.84.0
	google.golang.org/genproto v0.0.0-20220616135557-88e70c0c3a90
	google.golang.org/grpc v1.49.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	k8s.io/api v0.26.1
	k8s.io/apimachinery v0.26.1
	k8s.io/client-go v0.26.1
	sigs.k8s.io/controller-runtime v0.14.1
	sigs.k8s.io/controller-tools v0.11.1
)

replace github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1 => ./apis/vpcpeering/v1beta1

replace github.com/crossplane/provider-gcp/apis/dns/v1alpha1 => ./apis/dns/v1alpha1
