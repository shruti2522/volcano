/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	batchv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/batch/v1alpha1"
	busv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/bus/v1alpha1"
	schedulingv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	BatchV1alpha1() batchv1alpha1.BatchV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Batch() batchv1alpha1.BatchV1alpha1Interface
	BusV1alpha1() busv1alpha1.BusV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Bus() busv1alpha1.BusV1alpha1Interface
	SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Scheduling() schedulingv1alpha1.SchedulingV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	batchV1alpha1      *batchv1alpha1.BatchV1alpha1Client
	busV1alpha1        *busv1alpha1.BusV1alpha1Client
	schedulingV1alpha1 *schedulingv1alpha1.SchedulingV1alpha1Client
}

// BatchV1alpha1 retrieves the BatchV1alpha1Client
func (c *Clientset) BatchV1alpha1() batchv1alpha1.BatchV1alpha1Interface {
	return c.batchV1alpha1
}

// Deprecated: Batch retrieves the default version of BatchClient.
// Please explicitly pick a version.
func (c *Clientset) Batch() batchv1alpha1.BatchV1alpha1Interface {
	return c.batchV1alpha1
}

// BusV1alpha1 retrieves the BusV1alpha1Client
func (c *Clientset) BusV1alpha1() busv1alpha1.BusV1alpha1Interface {
	return c.busV1alpha1
}

// Deprecated: Bus retrieves the default version of BusClient.
// Please explicitly pick a version.
func (c *Clientset) Bus() busv1alpha1.BusV1alpha1Interface {
	return c.busV1alpha1
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

// Deprecated: Scheduling retrieves the default version of SchedulingClient.
// Please explicitly pick a version.
func (c *Clientset) Scheduling() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.batchV1alpha1, err = batchv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.busV1alpha1, err = busv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.batchV1alpha1 = batchv1alpha1.NewForConfigOrDie(c)
	cs.busV1alpha1 = busv1alpha1.NewForConfigOrDie(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.batchV1alpha1 = batchv1alpha1.New(c)
	cs.busV1alpha1 = busv1alpha1.New(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
