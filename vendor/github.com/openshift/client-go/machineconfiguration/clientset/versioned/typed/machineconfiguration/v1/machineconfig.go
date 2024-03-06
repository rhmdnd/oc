// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/api/machineconfiguration/v1"
	machineconfigurationv1 "github.com/openshift/client-go/machineconfiguration/applyconfigurations/machineconfiguration/v1"
	scheme "github.com/openshift/client-go/machineconfiguration/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineConfigsGetter has a method to return a MachineConfigInterface.
// A group's client should implement this interface.
type MachineConfigsGetter interface {
	MachineConfigs() MachineConfigInterface
}

// MachineConfigInterface has methods to work with MachineConfig resources.
type MachineConfigInterface interface {
	Create(ctx context.Context, machineConfig *v1.MachineConfig, opts metav1.CreateOptions) (*v1.MachineConfig, error)
	Update(ctx context.Context, machineConfig *v1.MachineConfig, opts metav1.UpdateOptions) (*v1.MachineConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.MachineConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MachineConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MachineConfig, err error)
	Apply(ctx context.Context, machineConfig *machineconfigurationv1.MachineConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.MachineConfig, err error)
	MachineConfigExpansion
}

// machineConfigs implements MachineConfigInterface
type machineConfigs struct {
	client rest.Interface
}

// newMachineConfigs returns a MachineConfigs
func newMachineConfigs(c *MachineconfigurationV1Client) *machineConfigs {
	return &machineConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the machineConfig, and returns the corresponding machineConfig object, and an error if there is any.
func (c *machineConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MachineConfig, err error) {
	result = &v1.MachineConfig{}
	err = c.client.Get().
		Resource("machineconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineConfigs that match those selectors.
func (c *machineConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MachineConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MachineConfigList{}
	err = c.client.Get().
		Resource("machineconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineConfigs.
func (c *machineConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("machineconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a machineConfig and creates it.  Returns the server's representation of the machineConfig, and an error, if there is any.
func (c *machineConfigs) Create(ctx context.Context, machineConfig *v1.MachineConfig, opts metav1.CreateOptions) (result *v1.MachineConfig, err error) {
	result = &v1.MachineConfig{}
	err = c.client.Post().
		Resource("machineconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a machineConfig and updates it. Returns the server's representation of the machineConfig, and an error, if there is any.
func (c *machineConfigs) Update(ctx context.Context, machineConfig *v1.MachineConfig, opts metav1.UpdateOptions) (result *v1.MachineConfig, err error) {
	result = &v1.MachineConfig{}
	err = c.client.Put().
		Resource("machineconfigs").
		Name(machineConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the machineConfig and deletes it. Returns an error if one occurs.
func (c *machineConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("machineconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("machineconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched machineConfig.
func (c *machineConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MachineConfig, err error) {
	result = &v1.MachineConfig{}
	err = c.client.Patch(pt).
		Resource("machineconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied machineConfig.
func (c *machineConfigs) Apply(ctx context.Context, machineConfig *machineconfigurationv1.MachineConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.MachineConfig, err error) {
	if machineConfig == nil {
		return nil, fmt.Errorf("machineConfig provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(machineConfig)
	if err != nil {
		return nil, err
	}
	name := machineConfig.Name
	if name == nil {
		return nil, fmt.Errorf("machineConfig.Name must be provided to Apply")
	}
	result = &v1.MachineConfig{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("machineconfigs").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
