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

package v1

import (
	"context"
	"time"

	v1 "github.com/istio-ecosystem/admiral/admiral/pkg/apis/admiral/v1"
	scheme "github.com/istio-ecosystem/admiral/admiral/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DependencyProxiesGetter has a method to return a DependencyProxyInterface.
// A group's client should implement this interface.
type DependencyProxiesGetter interface {
	DependencyProxies(namespace string) DependencyProxyInterface
}

// DependencyProxyInterface has methods to work with DependencyProxy resources.
type DependencyProxyInterface interface {
	Create(ctx context.Context, dependencyProxy *v1.DependencyProxy, opts metav1.CreateOptions) (*v1.DependencyProxy, error)
	Update(ctx context.Context, dependencyProxy *v1.DependencyProxy, opts metav1.UpdateOptions) (*v1.DependencyProxy, error)
	UpdateStatus(ctx context.Context, dependencyProxy *v1.DependencyProxy, opts metav1.UpdateOptions) (*v1.DependencyProxy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DependencyProxy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DependencyProxyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DependencyProxy, err error)
	DependencyProxyExpansion
}

// dependencyProxies implements DependencyProxyInterface
type dependencyProxies struct {
	client rest.Interface
	ns     string
}

// newDependencyProxies returns a DependencyProxies
func newDependencyProxies(c *AdmiralV1Client, namespace string) *dependencyProxies {
	return &dependencyProxies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dependencyProxy, and returns the corresponding dependencyProxy object, and an error if there is any.
func (c *dependencyProxies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DependencyProxy, err error) {
	result = &v1.DependencyProxy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dependencyproxies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DependencyProxies that match those selectors.
func (c *dependencyProxies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DependencyProxyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DependencyProxyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dependencyproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dependencyProxies.
func (c *dependencyProxies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dependencyproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dependencyProxy and creates it.  Returns the server's representation of the dependencyProxy, and an error, if there is any.
func (c *dependencyProxies) Create(ctx context.Context, dependencyProxy *v1.DependencyProxy, opts metav1.CreateOptions) (result *v1.DependencyProxy, err error) {
	result = &v1.DependencyProxy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dependencyproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dependencyProxy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dependencyProxy and updates it. Returns the server's representation of the dependencyProxy, and an error, if there is any.
func (c *dependencyProxies) Update(ctx context.Context, dependencyProxy *v1.DependencyProxy, opts metav1.UpdateOptions) (result *v1.DependencyProxy, err error) {
	result = &v1.DependencyProxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dependencyproxies").
		Name(dependencyProxy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dependencyProxy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dependencyProxies) UpdateStatus(ctx context.Context, dependencyProxy *v1.DependencyProxy, opts metav1.UpdateOptions) (result *v1.DependencyProxy, err error) {
	result = &v1.DependencyProxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dependencyproxies").
		Name(dependencyProxy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dependencyProxy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dependencyProxy and deletes it. Returns an error if one occurs.
func (c *dependencyProxies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dependencyproxies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dependencyProxies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dependencyproxies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dependencyProxy.
func (c *dependencyProxies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DependencyProxy, err error) {
	result = &v1.DependencyProxy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dependencyproxies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
