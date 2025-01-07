/*
Copyright 2017 The Kubernetes Authors.

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

package api

import (
	"k8s.io/apimachinery/pkg/util/sets"

	topologyv1alpha1 "volcano.sh/apis/pkg/apis/topology/v1alpha1"
)

type HyperNodesInfo struct {
	HyperNodes           map[string]*topologyv1alpha1.HyperNode
	HyperNodesListByTier map[int][]string
	HyperNodesSet        map[string]sets.Set[string]

	ProcessedNodes sets.Set[string]
}

func NewHyperNodesInfo() *HyperNodesInfo {
	hni := &HyperNodesInfo{
		HyperNodes:           make(map[string]*topologyv1alpha1.HyperNode),
		HyperNodesListByTier: make(map[int][]string),
		HyperNodesSet:        make(map[string]sets.Set[string]),
		ProcessedNodes:       sets.Set[string]{},
	}
	return hni
}
