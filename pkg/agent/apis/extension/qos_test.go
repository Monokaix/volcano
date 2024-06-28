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

package extension

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetQosLevel(t *testing.T) {
	tests := []struct {
		name string
		pod  *corev1.Pod
		want int32
	}{
		{
			name: "no qos level specified, default to 0",
			pod:  makePod("", "", 0),
			want: 0,
		},
		// Guaranteed
		{
			name: "qos Guaranteed, no priority class",
			pod:  makePod("Guaranteed", "", 0),
			want: 2,
		},
		{
			name: "qos Guaranteed, production priority class",
			pod:  makePod("Guaranteed", "volcano-production", 999999),
			want: 2,
		},
		{
			name: "qos Guaranteed, mid priority class",
			pod:  makePod("Guaranteed", "volcano-mid", 899999),
			want: 1,
		},
		{
			name: "qos Guaranteed, higher priority value",
			pod:  makePod("Guaranteed", "", 1000000),
			want: 2,
		},
		// Burstable
		{
			name: "qos Burstable, no priority class",
			pod:  makePod("Burstable", "", 0),
			want: 0,
		},
		{
			name: "qos Burstable, other priority class",
			pod:  makePod("Burstable", "other", 1111),
			want: 0,
		},
		// BestEffort
		{
			name: "qos BestEffort, no priority class",
			pod:  makePod("BestEffort", "", 0),
			want: -2,
		},
		{
			name: "qos BestEffort, free priority class",
			pod:  makePod("BestEffort", "volcano-free", -90000),
			want: -2,
		},
		{
			name: "qos BestEffort, batch priority class",
			pod:  makePod("BestEffort", "volcano-batch", -9000),
			want: -1,
		},
		{
			name: "qos BestEffort, other priority value",
			pod:  makePod("BestEffort", "", -1000),
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQosLevel(tt.pod); got != tt.want {
				t.Errorf("GetQosLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makePod(qosLevel string, pc string, priority int32) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"volcano.sh/qos-level": qosLevel,
			},
		},
		Spec: corev1.PodSpec{
			PriorityClassName: pc,
			Priority:          &priority,
		},
	}
}
