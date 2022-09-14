module github.com/crossplane/provider-gcp

go 1.13

require (
	cloud.google.com/go v0.81.0
	cloud.google.com/go/pubsub v1.3.1
	cloud.google.com/go/storage v1.10.0
	github.com/crossplane/crossplane-runtime v0.15.1-0.20210930095326-d5661210733b
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1 v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.6
	github.com/googleapis/gax-go v1.0.3
	github.com/googleapis/gax-go/v2 v2.0.5
	github.com/imdario/mergo v0.3.12
	github.com/mitchellh/copystructure v1.0.0
	github.com/onsi/gomega v1.18.1
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20210503060351-7fd8e65b6420
	google.golang.org/api v0.47.0
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.38.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)

replace github.com/crossplane/provider-gcp/apis/vpcpeering/v1beta1 => ./apis/vpcpeering/v1beta1
