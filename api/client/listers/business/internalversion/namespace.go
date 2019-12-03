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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	business "tkestack.io/tke/api/business"
)

// NamespaceLister helps list Namespaces.
type NamespaceLister interface {
	// List lists all Namespaces in the indexer.
	List(selector labels.Selector) (ret []*business.Namespace, err error)
	// Namespaces returns an object that can list and get Namespaces.
	Namespaces(namespace string) NamespaceNamespaceLister
	NamespaceListerExpansion
}

// namespaceLister implements the NamespaceLister interface.
type namespaceLister struct {
	indexer cache.Indexer
}

// NewNamespaceLister returns a new NamespaceLister.
func NewNamespaceLister(indexer cache.Indexer) NamespaceLister {
	return &namespaceLister{indexer: indexer}
}

// List lists all Namespaces in the indexer.
func (s *namespaceLister) List(selector labels.Selector) (ret []*business.Namespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*business.Namespace))
	})
	return ret, err
}

// Namespaces returns an object that can list and get Namespaces.
func (s *namespaceLister) Namespaces(namespace string) NamespaceNamespaceLister {
	return namespaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NamespaceNamespaceLister helps list and get Namespaces.
type NamespaceNamespaceLister interface {
	// List lists all Namespaces in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*business.Namespace, err error)
	// Get retrieves the Namespace from the indexer for a given namespace and name.
	Get(name string) (*business.Namespace, error)
	NamespaceNamespaceListerExpansion
}

// namespaceNamespaceLister implements the NamespaceNamespaceLister
// interface.
type namespaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Namespaces in the indexer for a given namespace.
func (s namespaceNamespaceLister) List(selector labels.Selector) (ret []*business.Namespace, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*business.Namespace))
	})
	return ret, err
}

// Get retrieves the Namespace from the indexer for a given namespace and name.
func (s namespaceNamespaceLister) Get(name string) (*business.Namespace, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(business.Resource("namespace"), name)
	}
	return obj.(*business.Namespace), nil
}
