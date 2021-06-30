/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudmemorystore

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	redisv1pb "google.golang.org/genproto/googleapis/cloud/redis/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/crossplane/crossplane-runtime/pkg/meta"

	"github.com/crossplane/provider-gcp/apis/cache/v1beta1"
)

const (
	region        = "us-cool1"
	project       = "coolProject"
	instanceName  = "claimnamespace-claimname-342sd"
	fullName      = "projects/coolProject/locations/us-cool1/instances/claimnamespace-claimname-342sd"
	qualifiedName = "projects/" + project + "/locations/" + region + "/instances/" + instanceName
	parent        = "projects/" + project + "/locations/" + region

	memorySizeGB = 1
)

var (
	authorizedNetwork = "default"

	redisConfigs = map[string]string{"cool": "socool"}
)

func TestInstanceID(t *testing.T) {
	cases := []struct {
		name       string
		project    string
		i          *v1beta1.CloudMemorystoreInstance
		want       InstanceID
		wantName   string
		wantParent string
	}{
		{
			name:    "Success",
			project: project,
			i: &v1beta1.CloudMemorystoreInstance{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						meta.AnnotationKeyExternalName: instanceName,
					},
				},
				Spec: v1beta1.CloudMemorystoreInstanceSpec{
					ForProvider: v1beta1.CloudMemorystoreInstanceParameters{
						Region: region,
					},
				},
			},
			want: InstanceID{
				Project:  project,
				Region:   region,
				Instance: instanceName,
			},
			wantName:   qualifiedName,
			wantParent: parent,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewInstanceID(tc.project, tc.i)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("NewInstanceID(...): -want, +got:\n%s", diff)
			}

			gotName := got.Name()
			if gotName != tc.wantName {
				t.Errorf("got.Name(): want: %s got: %s", tc.wantName, gotName)
			}

			gotParent := got.Parent()
			if gotParent != tc.wantParent {
				t.Errorf("got.Parent(): want: %s got: %s", tc.wantParent, gotParent)
			}
		})
	}
}

func TestIsUpToDate(t *testing.T) {
	randString := "wat"
	type want struct {
		upToDate bool
		isErr    bool
	}
	cases := []struct {
		name string
		id   InstanceID
		kube *v1beta1.CloudMemorystoreInstance
		gcp  *redisv1pb.Instance
		want want
	}{
		{
			name: "NeedsLessMemory",
			id:   InstanceID{Project: project, Region: region, Instance: instanceName},
			kube: &v1beta1.CloudMemorystoreInstance{
				Spec: v1beta1.CloudMemorystoreInstanceSpec{
					ForProvider: v1beta1.CloudMemorystoreInstanceParameters{
						RedisConfigs: redisConfigs,
						MemorySizeGB: memorySizeGB,
					},
				},
			},
			gcp: &redisv1pb.Instance{
				Name:         fullName,
				MemorySizeGb: memorySizeGB + 1,
			},
			want: want{upToDate: false, isErr: false},
		},
		{
			name: "NeedsNewRedisConfigs",
			id:   InstanceID{Project: project, Region: region, Instance: instanceName},
			kube: &v1beta1.CloudMemorystoreInstance{
				Spec: v1beta1.CloudMemorystoreInstanceSpec{
					ForProvider: v1beta1.CloudMemorystoreInstanceParameters{
						RedisConfigs: redisConfigs,
					},
				},
			},
			gcp: &redisv1pb.Instance{
				Name:         fullName,
				RedisConfigs: map[string]string{"super": "cool"},
			},
			want: want{upToDate: false, isErr: false},
		},
		{
			name: "NeedsNoUpdate",
			id:   InstanceID{Project: project, Region: region, Instance: instanceName},
			kube: &v1beta1.CloudMemorystoreInstance{
				Spec: v1beta1.CloudMemorystoreInstanceSpec{
					ForProvider: v1beta1.CloudMemorystoreInstanceParameters{
						RedisConfigs: redisConfigs,
						MemorySizeGB: memorySizeGB,
					},
				},
			},
			gcp: &redisv1pb.Instance{
				Name:         fullName,
				RedisConfigs: redisConfigs,
				MemorySizeGb: memorySizeGB,
			},
			want: want{upToDate: true, isErr: false},
		},
		{
			name: "NeedsNoUpdateWithOutputFields",
			id:   InstanceID{Project: project, Region: region, Instance: instanceName},
			kube: &v1beta1.CloudMemorystoreInstance{
				Spec: v1beta1.CloudMemorystoreInstanceSpec{
					ForProvider: v1beta1.CloudMemorystoreInstanceParameters{
						RedisConfigs: redisConfigs,
						MemorySizeGB: memorySizeGB,
					},
				},
			},
			gcp: &redisv1pb.Instance{
				Name:          fullName,
				RedisConfigs:  redisConfigs,
				MemorySizeGb:  memorySizeGB,
				StatusMessage: "definitely not in spec",
			},
			want: want{upToDate: true, isErr: false},
		},
		{
			name: "CannotUpdateField",
			id:   InstanceID{Project: project, Region: region, Instance: instanceName},
			kube: &v1beta1.CloudMemorystoreInstance{
				Spec: v1beta1.CloudMemorystoreInstanceSpec{
					ForProvider: v1beta1.CloudMemorystoreInstanceParameters{
						MemorySizeGB: memorySizeGB,

						// We can't change this field without destroying and recreating
						// the instance so it does not register as needing an update.
						AuthorizedNetwork: &randString,
					},
				},
			},
			gcp: &redisv1pb.Instance{
				Name:              fullName,
				MemorySizeGb:      memorySizeGB,
				AuthorizedNetwork: authorizedNetwork,
			},
			want: want{upToDate: true, isErr: false},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := IsUpToDate(tc.id, &tc.kube.Spec.ForProvider, tc.gcp)
			if err != nil && !tc.want.isErr {
				t.Error("IsUpToDate(...) unexpected error")
			}
			if got != tc.want.upToDate {
				t.Errorf("IsUpToDate(...): want: %t got: %t", tc.want, got)
			}
		})
	}
}
