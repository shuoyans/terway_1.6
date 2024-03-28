/*
Copyright 2021 Terway Authors.

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
	"context"

	v1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodENIs implements PodENIInterface
type FakePodENIs struct {
	Fake *FakeNetworkV1beta1
	ns   string
}

var podenisResource = schema.GroupVersionResource{Group: "network.alibabacloud.com", Version: "v1beta1", Resource: "podenis"}

var podenisKind = schema.GroupVersionKind{Group: "network.alibabacloud.com", Version: "v1beta1", Kind: "PodENI"}

// Get takes name of the podENI, and returns the corresponding podENI object, and an error if there is any.
func (c *FakePodENIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PodENI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podenisResource, c.ns, name), &v1beta1.PodENI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodENI), err
}

// List takes label and field selectors, and returns the list of PodENIs that match those selectors.
func (c *FakePodENIs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PodENIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podenisResource, podenisKind, c.ns, opts), &v1beta1.PodENIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PodENIList{ListMeta: obj.(*v1beta1.PodENIList).ListMeta}
	for _, item := range obj.(*v1beta1.PodENIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podENIs.
func (c *FakePodENIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podenisResource, c.ns, opts))

}

// Create takes the representation of a podENI and creates it.  Returns the server's representation of the podENI, and an error, if there is any.
func (c *FakePodENIs) Create(ctx context.Context, podENI *v1beta1.PodENI, opts v1.CreateOptions) (result *v1beta1.PodENI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podenisResource, c.ns, podENI), &v1beta1.PodENI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodENI), err
}

// Update takes the representation of a podENI and updates it. Returns the server's representation of the podENI, and an error, if there is any.
func (c *FakePodENIs) Update(ctx context.Context, podENI *v1beta1.PodENI, opts v1.UpdateOptions) (result *v1beta1.PodENI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podenisResource, c.ns, podENI), &v1beta1.PodENI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodENI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodENIs) UpdateStatus(ctx context.Context, podENI *v1beta1.PodENI, opts v1.UpdateOptions) (*v1beta1.PodENI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podenisResource, "status", c.ns, podENI), &v1beta1.PodENI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodENI), err
}

// Delete takes name of the podENI and deletes it. Returns an error if one occurs.
func (c *FakePodENIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(podenisResource, c.ns, name, opts), &v1beta1.PodENI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodENIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podenisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PodENIList{})
	return err
}

// Patch applies the patch and returns the patched podENI.
func (c *FakePodENIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PodENI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podenisResource, c.ns, name, pt, data, subresources...), &v1beta1.PodENI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodENI), err
}