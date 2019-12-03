/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	notify "tkestack.io/tke/api/notify"
)

// FakeMessageRequests implements MessageRequestInterface
type FakeMessageRequests struct {
	Fake *FakeNotify
	ns   string
}

var messagerequestsResource = schema.GroupVersionResource{Group: "notify.tkestack.io", Version: "", Resource: "messagerequests"}

var messagerequestsKind = schema.GroupVersionKind{Group: "notify.tkestack.io", Version: "", Kind: "MessageRequest"}

// Get takes name of the messageRequest, and returns the corresponding messageRequest object, and an error if there is any.
func (c *FakeMessageRequests) Get(name string, options v1.GetOptions) (result *notify.MessageRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(messagerequestsResource, c.ns, name), &notify.MessageRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*notify.MessageRequest), err
}

// List takes label and field selectors, and returns the list of MessageRequests that match those selectors.
func (c *FakeMessageRequests) List(opts v1.ListOptions) (result *notify.MessageRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(messagerequestsResource, messagerequestsKind, c.ns, opts), &notify.MessageRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &notify.MessageRequestList{ListMeta: obj.(*notify.MessageRequestList).ListMeta}
	for _, item := range obj.(*notify.MessageRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested messageRequests.
func (c *FakeMessageRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(messagerequestsResource, c.ns, opts))

}

// Create takes the representation of a messageRequest and creates it.  Returns the server's representation of the messageRequest, and an error, if there is any.
func (c *FakeMessageRequests) Create(messageRequest *notify.MessageRequest) (result *notify.MessageRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(messagerequestsResource, c.ns, messageRequest), &notify.MessageRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*notify.MessageRequest), err
}

// Update takes the representation of a messageRequest and updates it. Returns the server's representation of the messageRequest, and an error, if there is any.
func (c *FakeMessageRequests) Update(messageRequest *notify.MessageRequest) (result *notify.MessageRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(messagerequestsResource, c.ns, messageRequest), &notify.MessageRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*notify.MessageRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMessageRequests) UpdateStatus(messageRequest *notify.MessageRequest) (*notify.MessageRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(messagerequestsResource, "status", c.ns, messageRequest), &notify.MessageRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*notify.MessageRequest), err
}

// Delete takes name of the messageRequest and deletes it. Returns an error if one occurs.
func (c *FakeMessageRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(messagerequestsResource, c.ns, name), &notify.MessageRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMessageRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(messagerequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &notify.MessageRequestList{})
	return err
}

// Patch applies the patch and returns the patched messageRequest.
func (c *FakeMessageRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.MessageRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(messagerequestsResource, c.ns, name, pt, data, subresources...), &notify.MessageRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*notify.MessageRequest), err
}
