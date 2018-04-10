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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kong/kubernetes-ingress-controller/internal/apis/credential/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KongCredentialLister helps list KongCredentials.
type KongCredentialLister interface {
	// List lists all KongCredentials in the indexer.
	List(selector labels.Selector) (ret []*v1.KongCredential, err error)
	// KongCredentials returns an object that can list and get KongCredentials.
	KongCredentials(namespace string) KongCredentialNamespaceLister
	KongCredentialListerExpansion
}

// kongCredentialLister implements the KongCredentialLister interface.
type kongCredentialLister struct {
	indexer cache.Indexer
}

// NewKongCredentialLister returns a new KongCredentialLister.
func NewKongCredentialLister(indexer cache.Indexer) KongCredentialLister {
	return &kongCredentialLister{indexer: indexer}
}

// List lists all KongCredentials in the indexer.
func (s *kongCredentialLister) List(selector labels.Selector) (ret []*v1.KongCredential, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KongCredential))
	})
	return ret, err
}

// KongCredentials returns an object that can list and get KongCredentials.
func (s *kongCredentialLister) KongCredentials(namespace string) KongCredentialNamespaceLister {
	return kongCredentialNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KongCredentialNamespaceLister helps list and get KongCredentials.
type KongCredentialNamespaceLister interface {
	// List lists all KongCredentials in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.KongCredential, err error)
	// Get retrieves the KongCredential from the indexer for a given namespace and name.
	Get(name string) (*v1.KongCredential, error)
	KongCredentialNamespaceListerExpansion
}

// kongCredentialNamespaceLister implements the KongCredentialNamespaceLister
// interface.
type kongCredentialNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KongCredentials in the indexer for a given namespace.
func (s kongCredentialNamespaceLister) List(selector labels.Selector) (ret []*v1.KongCredential, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KongCredential))
	})
	return ret, err
}

// Get retrieves the KongCredential from the indexer for a given namespace and name.
func (s kongCredentialNamespaceLister) Get(name string) (*v1.KongCredential, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kongcredential"), name)
	}
	return obj.(*v1.KongCredential), nil
}
