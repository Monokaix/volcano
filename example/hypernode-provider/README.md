
# Overview

As stared in the [network topology aware scheduling](https://github.com/volcano-sh/volcano/blob/master/docs/design/Network%20Topology%20Aware%20Scheduling.md#network-topology-generation-and-update), datacenter's network topology details are different, besides create and update hyperNodes CRD manually, Volcano should also support a method that auto-discovery the network topology and update the hyperNodes CRD automatically, this need the cooperation with under-layer hardware vendors, Volcano controller supports the basic hyperNodes reconcile framework, and expose an interface which can interactive with the hardware vendors, the vendors behaves as a hyeprNodes providers and reconcile hyperNodes such as creating/updating/reporting healthy status, through which Volcano can adapt any hardware vendors with auto-discovery network topology tools supported, and vendors just need to focus on the auto-discovery mechanism while Volcnao supports a basic framework and integrate them with an extensible way. This is similar to the [cloud provider mechanism](https://github.com/kubernetes/cloud-provider) in kubernetes.

# How to use

## Write your provider
Write your codes locally and implement the `Plugin` interface in file `pkg/controllers/hypernode/provider/interface.go`, there are two critical params that you need to concern:
`eventCh chan<- Event`: Vendors should send the hyperNode create/update/delete event to this channel, and Volcano controller will communicate to API Server to store them.
`replyCh <-chan Reply`: Volcano will reply errors to vendor providers through this channel when an unexpected error occurs when communicating with the API Server and the retry does not succeed, providers should be aware of that and should resend the event or perform fault-tolerant processing.

There in an example in `example/hypernode-provider/example_provider.go` demonstrated how to write a provider.

## Build the provider with volcano controller
