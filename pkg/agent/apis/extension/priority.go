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
	v1 "k8s.io/api/core/v1"
)

type PriorityClass string

const (
	// PriorityClassProduction Priority value range [900000, 999999].
	PriorityClassProduction PriorityClass = "volcano-production"
	// PriorityClassMid Priority value range [800000, 899999].
	PriorityClassMid PriorityClass = "volcano-mid"
	// PriorityClassDefault Priority value=0.

	// PriorityClassBatch Priority value range [-9999, -9000].
	PriorityClassBatch PriorityClass = "volcano-batch"
	// PriorityClassFree Priority value range [-99999, -90000].
	PriorityClassFree PriorityClass = "volcano-free"

	PriorityClassProductionMaxValue = 999999
)

var QosLevelMap = map[QosLevel]func(pod *v1.Pod) int32{
	QosLevelGuaranteed: qosLevelOfGuaranteed,
	QosLevelBurstable:  qosLevelOfBurstable,
	QosLevelBestEffort: qosLevelOfBestEffort,
}

func qosLevelOfGuaranteed(pod *v1.Pod) int32 {
	if pod.Spec.Priority == nil {
		return 2
	}

	if pod.Spec.PriorityClassName == string(PriorityClassProduction) || *pod.Spec.Priority >= PriorityClassProductionMaxValue {
		return 2
	}

	if pod.Spec.PriorityClassName == string(PriorityClassMid) {
		return 1
	}

	// default to 2 means the highest priority.
	return 2
}

func qosLevelOfBurstable(pod *v1.Pod) int32 {
	return 0
}

func qosLevelOfBestEffort(pod *v1.Pod) int32 {
	if pod.Spec.Priority == nil {
		return -2
	}

	if pod.Spec.PriorityClassName == string(PriorityClassFree) {
		return -2
	}

	if pod.Spec.PriorityClassName == string(PriorityClassBatch) {
		return -1
	}

	// default to -2 means the lowest priority.
	return -2
}
