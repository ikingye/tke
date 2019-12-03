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
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakePrometheuses implements PrometheusInterface
type FakePrometheuses struct {
	Fake *FakePlatformV1
}

var prometheusesResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "prometheuses"}

var prometheusesKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "Prometheus"}

// Get takes name of the prometheus, and returns the corresponding prometheus object, and an error if there is any.
func (c *FakePrometheuses) Get(name string, options v1.GetOptions) (result *platformv1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(prometheusesResource, name), &platformv1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Prometheus), err
}

// List takes label and field selectors, and returns the list of Prometheuses that match those selectors.
func (c *FakePrometheuses) List(opts v1.ListOptions) (result *platformv1.PrometheusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(prometheusesResource, prometheusesKind, opts), &platformv1.PrometheusList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.PrometheusList{ListMeta: obj.(*platformv1.PrometheusList).ListMeta}
	for _, item := range obj.(*platformv1.PrometheusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prometheuses.
func (c *FakePrometheuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(prometheusesResource, opts))
}

// Create takes the representation of a prometheus and creates it.  Returns the server's representation of the prometheus, and an error, if there is any.
func (c *FakePrometheuses) Create(prometheus *platformv1.Prometheus) (result *platformv1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(prometheusesResource, prometheus), &platformv1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Prometheus), err
}

// Update takes the representation of a prometheus and updates it. Returns the server's representation of the prometheus, and an error, if there is any.
func (c *FakePrometheuses) Update(prometheus *platformv1.Prometheus) (result *platformv1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(prometheusesResource, prometheus), &platformv1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Prometheus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrometheuses) UpdateStatus(prometheus *platformv1.Prometheus) (*platformv1.Prometheus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(prometheusesResource, "status", prometheus), &platformv1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Prometheus), err
}

// Delete takes name of the prometheus and deletes it. Returns an error if one occurs.
func (c *FakePrometheuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(prometheusesResource, name), &platformv1.Prometheus{})
	return err
}

// Patch applies the patch and returns the patched prometheus.
func (c *FakePrometheuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platformv1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(prometheusesResource, name, pt, data, subresources...), &platformv1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Prometheus), err
}
