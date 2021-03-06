/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	auth "tkestack.io/tke/api/auth"
)

// FakeIdentityProviders implements IdentityProviderInterface
type FakeIdentityProviders struct {
	Fake *FakeAuth
}

var identityprovidersResource = schema.GroupVersionResource{Group: "auth.tkestack.io", Version: "", Resource: "identityproviders"}

var identityprovidersKind = schema.GroupVersionKind{Group: "auth.tkestack.io", Version: "", Kind: "IdentityProvider"}

// Get takes name of the identityProvider, and returns the corresponding identityProvider object, and an error if there is any.
func (c *FakeIdentityProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *auth.IdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(identityprovidersResource, name), &auth.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.IdentityProvider), err
}

// List takes label and field selectors, and returns the list of IdentityProviders that match those selectors.
func (c *FakeIdentityProviders) List(ctx context.Context, opts v1.ListOptions) (result *auth.IdentityProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(identityprovidersResource, identityprovidersKind, opts), &auth.IdentityProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &auth.IdentityProviderList{ListMeta: obj.(*auth.IdentityProviderList).ListMeta}
	for _, item := range obj.(*auth.IdentityProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityProviders.
func (c *FakeIdentityProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(identityprovidersResource, opts))
}

// Create takes the representation of a identityProvider and creates it.  Returns the server's representation of the identityProvider, and an error, if there is any.
func (c *FakeIdentityProviders) Create(ctx context.Context, identityProvider *auth.IdentityProvider, opts v1.CreateOptions) (result *auth.IdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(identityprovidersResource, identityProvider), &auth.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.IdentityProvider), err
}

// Update takes the representation of a identityProvider and updates it. Returns the server's representation of the identityProvider, and an error, if there is any.
func (c *FakeIdentityProviders) Update(ctx context.Context, identityProvider *auth.IdentityProvider, opts v1.UpdateOptions) (result *auth.IdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(identityprovidersResource, identityProvider), &auth.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.IdentityProvider), err
}

// Delete takes name of the identityProvider and deletes it. Returns an error if one occurs.
func (c *FakeIdentityProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(identityprovidersResource, name), &auth.IdentityProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(identityprovidersResource, listOpts)

	_, err := c.Fake.Invokes(action, &auth.IdentityProviderList{})
	return err
}

// Patch applies the patch and returns the patched identityProvider.
func (c *FakeIdentityProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *auth.IdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(identityprovidersResource, name, pt, data, subresources...), &auth.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.IdentityProvider), err
}
