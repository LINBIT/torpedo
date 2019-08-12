/*
Copyright 2018 Openstorage.org

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

package v1alpha1

import (
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	scheme "github.com/libopenstorage/stork/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationBackupSchedulesGetter has a method to return a ApplicationBackupScheduleInterface.
// A group's client should implement this interface.
type ApplicationBackupSchedulesGetter interface {
	ApplicationBackupSchedules(namespace string) ApplicationBackupScheduleInterface
}

// ApplicationBackupScheduleInterface has methods to work with ApplicationBackupSchedule resources.
type ApplicationBackupScheduleInterface interface {
	Create(*v1alpha1.ApplicationBackupSchedule) (*v1alpha1.ApplicationBackupSchedule, error)
	Update(*v1alpha1.ApplicationBackupSchedule) (*v1alpha1.ApplicationBackupSchedule, error)
	UpdateStatus(*v1alpha1.ApplicationBackupSchedule) (*v1alpha1.ApplicationBackupSchedule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApplicationBackupSchedule, error)
	List(opts v1.ListOptions) (*v1alpha1.ApplicationBackupScheduleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationBackupSchedule, err error)
	ApplicationBackupScheduleExpansion
}

// applicationBackupSchedules implements ApplicationBackupScheduleInterface
type applicationBackupSchedules struct {
	client rest.Interface
	ns     string
}

// newApplicationBackupSchedules returns a ApplicationBackupSchedules
func newApplicationBackupSchedules(c *StorkV1alpha1Client, namespace string) *applicationBackupSchedules {
	return &applicationBackupSchedules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationBackupSchedule, and returns the corresponding applicationBackupSchedule object, and an error if there is any.
func (c *applicationBackupSchedules) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	result = &v1alpha1.ApplicationBackupSchedule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationBackupSchedules that match those selectors.
func (c *applicationBackupSchedules) List(opts v1.ListOptions) (result *v1alpha1.ApplicationBackupScheduleList, err error) {
	result = &v1alpha1.ApplicationBackupScheduleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationBackupSchedules.
func (c *applicationBackupSchedules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a applicationBackupSchedule and creates it.  Returns the server's representation of the applicationBackupSchedule, and an error, if there is any.
func (c *applicationBackupSchedules) Create(applicationBackupSchedule *v1alpha1.ApplicationBackupSchedule) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	result = &v1alpha1.ApplicationBackupSchedule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		Body(applicationBackupSchedule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a applicationBackupSchedule and updates it. Returns the server's representation of the applicationBackupSchedule, and an error, if there is any.
func (c *applicationBackupSchedules) Update(applicationBackupSchedule *v1alpha1.ApplicationBackupSchedule) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	result = &v1alpha1.ApplicationBackupSchedule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		Name(applicationBackupSchedule.Name).
		Body(applicationBackupSchedule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *applicationBackupSchedules) UpdateStatus(applicationBackupSchedule *v1alpha1.ApplicationBackupSchedule) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	result = &v1alpha1.ApplicationBackupSchedule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		Name(applicationBackupSchedule.Name).
		SubResource("status").
		Body(applicationBackupSchedule).
		Do().
		Into(result)
	return
}

// Delete takes name of the applicationBackupSchedule and deletes it. Returns an error if one occurs.
func (c *applicationBackupSchedules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationBackupSchedules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched applicationBackupSchedule.
func (c *applicationBackupSchedules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	result = &v1alpha1.ApplicationBackupSchedule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationbackupschedules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
