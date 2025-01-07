/*
Copyright 2024 The Volcano Authors.

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

package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"

	"github.com/stretchr/testify/assert"

	topologyv1alpha1 "volcano.sh/apis/pkg/apis/topology/v1alpha1"
)

func TestNewFakeBinder(t *testing.T) {
	tests := []struct {
		fkBinder   *FakeBinder
		cap, lenth int
	}{
		{
			fkBinder: NewFakeBinder(0),
			cap:      0, lenth: 0,
		},
		{
			fkBinder: NewFakeBinder(10),
			cap:      10, lenth: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.cap, cap(test.fkBinder.Channel))
		assert.Equal(t, test.lenth, test.fkBinder.Length())
		assert.Equal(t, test.lenth, len(test.fkBinder.Binds()))
	}
}

func TestNewFakeEvictor(t *testing.T) {
	tests := []struct {
		fkEvictor  *FakeEvictor
		cap, lenth int
	}{
		{
			fkEvictor: NewFakeEvictor(0),
			cap:       0, lenth: 0,
		},
		{
			fkEvictor: NewFakeEvictor(10),
			cap:       10, lenth: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.cap, cap(test.fkEvictor.Channel))
		assert.Equal(t, test.cap, cap(test.fkEvictor.evicts))
		assert.Equal(t, test.lenth, test.fkEvictor.Length())
		assert.Equal(t, test.lenth, len(test.fkEvictor.Evicts()))
	}
}

func TestBuildHyperNode(t *testing.T) {
	tests := []struct {
		name          string
		hyperNodeName string
		tier          string
		memberType    topologyv1alpha1.MemberType
		members       []string
		want          *topologyv1alpha1.HyperNode
	}{
		{
			name:          "build leaf hyperNode",
			hyperNodeName: "s0",
			tier:          "1",
			memberType:    topologyv1alpha1.MemberTypeNode,
			members:       []string{"node-1", "node-2"},
			want: &topologyv1alpha1.HyperNode{
				ObjectMeta: metav1.ObjectMeta{
					Name: "s0",
				},
				Spec: topologyv1alpha1.HyperNodeSpec{
					Tier: "1",
					Members: []topologyv1alpha1.MemberSpec{
						{Type: topologyv1alpha1.MemberTypeNode, Selector: topologyv1alpha1.MemberSelector{ExactMatch: &topologyv1alpha1.ExactMatch{Name: "node-1"}}},
						{Type: topologyv1alpha1.MemberTypeNode, Selector: topologyv1alpha1.MemberSelector{ExactMatch: &topologyv1alpha1.ExactMatch{Name: "node-2"}}},
					},
				},
			},
		},
		{
			name:          "build non-leaf hyperNode",
			hyperNodeName: "s4",
			tier:          "2",
			memberType:    topologyv1alpha1.MemberTypeHyperNode,
			members:       []string{"s0", "s1"},
			want: &topologyv1alpha1.HyperNode{
				ObjectMeta: metav1.ObjectMeta{
					Name: "s4",
				},
				Spec: topologyv1alpha1.HyperNodeSpec{
					Tier: "2",
					Members: []topologyv1alpha1.MemberSpec{
						{Type: topologyv1alpha1.MemberTypeHyperNode, Selector: topologyv1alpha1.MemberSelector{ExactMatch: &topologyv1alpha1.ExactMatch{Name: "s0"}}},
						{Type: topologyv1alpha1.MemberTypeHyperNode, Selector: topologyv1alpha1.MemberSelector{ExactMatch: &topologyv1alpha1.ExactMatch{Name: "s1"}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, BuildHyperNode(tt.hyperNodeName, tt.tier, tt.memberType, tt.members), "BuildHyperNode(%v, %v, %v, %v)", tt.hyperNodeName, tt.tier, tt.memberType, tt.members)
		})
	}
}
