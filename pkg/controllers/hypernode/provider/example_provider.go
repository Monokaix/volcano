/*
Copyright 2025 The Volcano Authors.

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

package provider

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"volcano.sh/apis/pkg/apis/topology/v1alpha1"
	topologyinformerv1alpha1 "volcano.sh/apis/pkg/client/informers/externalversions/topology/v1alpha1"
)

// ExampleProvider is an example provider of hyperNodes.
type ExampleProvider struct {
	hyperNodeInformer topologyinformerv1alpha1.HyperNodeInformer
}

// Name returns the name of the vendor.
func (e *ExampleProvider) Name() string {
	return "ExampleProvider"
}

// Start starts the vendor provider.
func (e *ExampleProvider) Start(eventCh chan<- Event, controlCh <-chan Reply, informer topologyinformerv1alpha1.HyperNodeInformer) error {
	e.hyperNodeInformer = informer
	go func() {
		for {
			select {
			default:
				event := Event{
					Type: EventAdd,
					HyperNode: v1alpha1.HyperNode{
						Spec: v1alpha1.HyperNodeSpec{
							Tier: 1,
							Members: []v1alpha1.MemberSpec{
								{
									Type: v1alpha1.MemberTypeNode,
									Selector: v1alpha1.MemberSelector{
										ExactMatch: &v1alpha1.ExactMatch{
											Name: "node-1",
										},
									},
								},
							},
						},
						Status: v1alpha1.HyperNodeStatus{
							NodeCount: 1,
							Conditions: []metav1.Condition{
								{
									Type:    "Ready",
									Status:  metav1.ConditionTrue,
									Reason:  "NodeAdded",
									Message: "Node added successfully",
								},
							},
						},
					},
				}
				eventCh <- event
				time.Sleep(5 * time.Second)
			}
		}
	}()

	return nil
}

// Stop stops the provider.
func (e *ExampleProvider) Stop() error {
	fmt.Println("ExampleProvider stopped")
	return nil
}
