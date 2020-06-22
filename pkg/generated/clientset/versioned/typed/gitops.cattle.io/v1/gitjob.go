/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/rancher/gitjobs/pkg/apis/gitops.cattle.io/v1"
	scheme "github.com/rancher/gitjobs/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GitJobsGetter has a method to return a GitJobInterface.
// A group's client should implement this interface.
type GitJobsGetter interface {
	GitJobs(namespace string) GitJobInterface
}

// GitJobInterface has methods to work with GitJob resources.
type GitJobInterface interface {
	Create(ctx context.Context, gitJob *v1.GitJob, opts metav1.CreateOptions) (*v1.GitJob, error)
	Update(ctx context.Context, gitJob *v1.GitJob, opts metav1.UpdateOptions) (*v1.GitJob, error)
	UpdateStatus(ctx context.Context, gitJob *v1.GitJob, opts metav1.UpdateOptions) (*v1.GitJob, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GitJob, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GitJobList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GitJob, err error)
	GitJobExpansion
}

// gitJobs implements GitJobInterface
type gitJobs struct {
	client rest.Interface
	ns     string
}

// newGitJobs returns a GitJobs
func newGitJobs(c *GitopsV1Client, namespace string) *gitJobs {
	return &gitJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gitJob, and returns the corresponding gitJob object, and an error if there is any.
func (c *gitJobs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GitJob, err error) {
	result = &v1.GitJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gitjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GitJobs that match those selectors.
func (c *gitJobs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GitJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GitJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gitjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gitJobs.
func (c *gitJobs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gitjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gitJob and creates it.  Returns the server's representation of the gitJob, and an error, if there is any.
func (c *gitJobs) Create(ctx context.Context, gitJob *v1.GitJob, opts metav1.CreateOptions) (result *v1.GitJob, err error) {
	result = &v1.GitJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gitjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gitJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gitJob and updates it. Returns the server's representation of the gitJob, and an error, if there is any.
func (c *gitJobs) Update(ctx context.Context, gitJob *v1.GitJob, opts metav1.UpdateOptions) (result *v1.GitJob, err error) {
	result = &v1.GitJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gitjobs").
		Name(gitJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gitJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gitJobs) UpdateStatus(ctx context.Context, gitJob *v1.GitJob, opts metav1.UpdateOptions) (result *v1.GitJob, err error) {
	result = &v1.GitJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gitjobs").
		Name(gitJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gitJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gitJob and deletes it. Returns an error if one occurs.
func (c *gitJobs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gitjobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gitJobs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gitjobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gitJob.
func (c *gitJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GitJob, err error) {
	result = &v1.GitJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gitjobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
