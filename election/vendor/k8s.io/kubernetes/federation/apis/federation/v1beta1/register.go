/*
Copyright 2016 The Kubernetes Authors.

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

package v1beta1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/runtime"
	versionedwatch "k8s.io/kubernetes/pkg/watch/versioned"
)

// GroupName is the group name use in this package
const GroupName = "federation"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1beta1"}

// Adds the list of known types to api.Scheme.
func AddToScheme(scheme *runtime.Scheme) {
	addKnownTypes(scheme)
	addDefaultingFuncs(scheme)
	addConversionFuncs(scheme)
}

func addKnownTypes(scheme *runtime.Scheme) {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Cluster{},
		&ClusterList{},
		&v1.ListOptions{},
		&v1.DeleteOptions{},
	)
	versionedwatch.AddToGroupVersion(scheme, SchemeGroupVersion)
}

func (obj *Cluster) GetObjectKind() unversioned.ObjectKind     { return &obj.TypeMeta }
func (obj *ClusterList) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
