// Copyright 2019 IBM Corp
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v2alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	kinds = []struct {
		groupVersion schema.GroupVersion
		obj          runtime.Object
		resource     string
		apiPath      string
	}{
		{v1.SchemeGroupVersion, &v1.ReplicationController{}, "replicationcontrollers", "/api"},
		{v1.SchemeGroupVersion, &v1.Pod{}, "pods", "/api"},

		{appsv1.SchemeGroupVersion, &appsv1.Deployment{}, "deployments", "/apis"},
		{appsv1.SchemeGroupVersion, &appsv1.DaemonSet{}, "daemonsets", "/apis"},
		{appsv1.SchemeGroupVersion, &appsv1.ReplicaSet{}, "replicasets", "/apis"},

		{batchv1.SchemeGroupVersion, &batchv1.Job{}, "jobs", "/apis"},
		{v2alpha1.SchemeGroupVersion, &v2alpha1.CronJob{}, "cronjobs", "/apis"},

		{appsv1.SchemeGroupVersion, &appsv1.StatefulSet{}, "statefulsets", "/apis"},

		{v1.SchemeGroupVersion, &v1.List{}, "lists", "/apis"},
	}
	scheme = runtime.NewScheme()
)

func init() {
	for _, kind := range kinds {
		scheme.AddKnownTypes(kind.groupVersion, kind.obj)
		scheme.AddUnversionedTypes(kind.groupVersion, kind.obj)
	}
}
