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

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	monitor "tkestack.io/tke/api/monitor"
)

// PrometheusesGetter has a method to return a PrometheusInterface.
// A group's client should implement this interface.
type PrometheusesGetter interface {
	Prometheuses() PrometheusInterface
}

// PrometheusInterface has methods to work with Prometheus resources.
type PrometheusInterface interface {
	Create(ctx context.Context, prometheus *monitor.Prometheus, opts v1.CreateOptions) (*monitor.Prometheus, error)
	Update(ctx context.Context, prometheus *monitor.Prometheus, opts v1.UpdateOptions) (*monitor.Prometheus, error)
	UpdateStatus(ctx context.Context, prometheus *monitor.Prometheus, opts v1.UpdateOptions) (*monitor.Prometheus, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*monitor.Prometheus, error)
	List(ctx context.Context, opts v1.ListOptions) (*monitor.PrometheusList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *monitor.Prometheus, err error)
	PrometheusExpansion
}

// prometheuses implements PrometheusInterface
type prometheuses struct {
	client rest.Interface
}

// newPrometheuses returns a Prometheuses
func newPrometheuses(c *MonitorClient) *prometheuses {
	return &prometheuses{
		client: c.RESTClient(),
	}
}

// Get takes name of the prometheus, and returns the corresponding prometheus object, and an error if there is any.
func (c *prometheuses) Get(ctx context.Context, name string, options v1.GetOptions) (result *monitor.Prometheus, err error) {
	result = &monitor.Prometheus{}
	err = c.client.Get().
		Resource("prometheuses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Prometheuses that match those selectors.
func (c *prometheuses) List(ctx context.Context, opts v1.ListOptions) (result *monitor.PrometheusList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &monitor.PrometheusList{}
	err = c.client.Get().
		Resource("prometheuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested prometheuses.
func (c *prometheuses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("prometheuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a prometheus and creates it.  Returns the server's representation of the prometheus, and an error, if there is any.
func (c *prometheuses) Create(ctx context.Context, prometheus *monitor.Prometheus, opts v1.CreateOptions) (result *monitor.Prometheus, err error) {
	result = &monitor.Prometheus{}
	err = c.client.Post().
		Resource("prometheuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prometheus).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a prometheus and updates it. Returns the server's representation of the prometheus, and an error, if there is any.
func (c *prometheuses) Update(ctx context.Context, prometheus *monitor.Prometheus, opts v1.UpdateOptions) (result *monitor.Prometheus, err error) {
	result = &monitor.Prometheus{}
	err = c.client.Put().
		Resource("prometheuses").
		Name(prometheus.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prometheus).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *prometheuses) UpdateStatus(ctx context.Context, prometheus *monitor.Prometheus, opts v1.UpdateOptions) (result *monitor.Prometheus, err error) {
	result = &monitor.Prometheus{}
	err = c.client.Put().
		Resource("prometheuses").
		Name(prometheus.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prometheus).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the prometheus and deletes it. Returns an error if one occurs.
func (c *prometheuses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("prometheuses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched prometheus.
func (c *prometheuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *monitor.Prometheus, err error) {
	result = &monitor.Prometheus{}
	err = c.client.Patch(pt).
		Resource("prometheuses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
