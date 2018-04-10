/*
Copyright 2018 The Kong Authors.

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

package fake

import (
	configuration_v1 "github.com/kong/kubernetes-ingress-controller/internal/apis/configuration/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKongIngressUpstreams implements KongIngressUpstreamInterface
type FakeKongIngressUpstreams struct {
	Fake *FakeConfigurationV1
	ns   string
}

var kongingressupstreamsResource = schema.GroupVersionResource{Group: "configuration.konghq.com", Version: "v1", Resource: "kongingressupstreams"}

var kongingressupstreamsKind = schema.GroupVersionKind{Group: "configuration.konghq.com", Version: "v1", Kind: "KongIngressUpstream"}

// Get takes name of the kongIngressUpstream, and returns the corresponding kongIngressUpstream object, and an error if there is any.
func (c *FakeKongIngressUpstreams) Get(name string, options v1.GetOptions) (result *configuration_v1.KongIngressUpstream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kongingressupstreamsResource, c.ns, name), &configuration_v1.KongIngressUpstream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configuration_v1.KongIngressUpstream), err
}

// List takes label and field selectors, and returns the list of KongIngressUpstreams that match those selectors.
func (c *FakeKongIngressUpstreams) List(opts v1.ListOptions) (result *configuration_v1.KongIngressUpstreamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kongingressupstreamsResource, kongingressupstreamsKind, c.ns, opts), &configuration_v1.KongIngressUpstreamList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configuration_v1.KongIngressUpstreamList{}
	for _, item := range obj.(*configuration_v1.KongIngressUpstreamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kongIngressUpstreams.
func (c *FakeKongIngressUpstreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kongingressupstreamsResource, c.ns, opts))

}

// Create takes the representation of a kongIngressUpstream and creates it.  Returns the server's representation of the kongIngressUpstream, and an error, if there is any.
func (c *FakeKongIngressUpstreams) Create(kongIngressUpstream *configuration_v1.KongIngressUpstream) (result *configuration_v1.KongIngressUpstream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kongingressupstreamsResource, c.ns, kongIngressUpstream), &configuration_v1.KongIngressUpstream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configuration_v1.KongIngressUpstream), err
}

// Update takes the representation of a kongIngressUpstream and updates it. Returns the server's representation of the kongIngressUpstream, and an error, if there is any.
func (c *FakeKongIngressUpstreams) Update(kongIngressUpstream *configuration_v1.KongIngressUpstream) (result *configuration_v1.KongIngressUpstream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kongingressupstreamsResource, c.ns, kongIngressUpstream), &configuration_v1.KongIngressUpstream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configuration_v1.KongIngressUpstream), err
}

// Delete takes name of the kongIngressUpstream and deletes it. Returns an error if one occurs.
func (c *FakeKongIngressUpstreams) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kongingressupstreamsResource, c.ns, name), &configuration_v1.KongIngressUpstream{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKongIngressUpstreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kongingressupstreamsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &configuration_v1.KongIngressUpstreamList{})
	return err
}

// Patch applies the patch and returns the patched kongIngressUpstream.
func (c *FakeKongIngressUpstreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *configuration_v1.KongIngressUpstream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kongingressupstreamsResource, c.ns, name, data, subresources...), &configuration_v1.KongIngressUpstream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configuration_v1.KongIngressUpstream), err
}
