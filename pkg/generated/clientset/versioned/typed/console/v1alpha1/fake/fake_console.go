// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/console-operator/pkg/apis/console/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConsoles implements ConsoleInterface
type FakeConsoles struct {
	Fake *FakeConsoleV1alpha1
	ns   string
}

var consolesResource = schema.GroupVersionResource{Group: "console.openshift.io", Version: "v1alpha1", Resource: "consoles"}

var consolesKind = schema.GroupVersionKind{Group: "console.openshift.io", Version: "v1alpha1", Kind: "Console"}

// Get takes name of the console, and returns the corresponding console object, and an error if there is any.
func (c *FakeConsoles) Get(name string, options v1.GetOptions) (result *v1alpha1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(consolesResource, c.ns, name), &v1alpha1.Console{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Console), err
}

// List takes label and field selectors, and returns the list of Consoles that match those selectors.
func (c *FakeConsoles) List(opts v1.ListOptions) (result *v1alpha1.ConsoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(consolesResource, consolesKind, c.ns, opts), &v1alpha1.ConsoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConsoleList{ListMeta: obj.(*v1alpha1.ConsoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConsoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consoles.
func (c *FakeConsoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(consolesResource, c.ns, opts))

}

// Create takes the representation of a console and creates it.  Returns the server's representation of the console, and an error, if there is any.
func (c *FakeConsoles) Create(console *v1alpha1.Console) (result *v1alpha1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(consolesResource, c.ns, console), &v1alpha1.Console{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Console), err
}

// Update takes the representation of a console and updates it. Returns the server's representation of the console, and an error, if there is any.
func (c *FakeConsoles) Update(console *v1alpha1.Console) (result *v1alpha1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(consolesResource, c.ns, console), &v1alpha1.Console{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Console), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConsoles) UpdateStatus(console *v1alpha1.Console) (*v1alpha1.Console, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(consolesResource, "status", c.ns, console), &v1alpha1.Console{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Console), err
}

// Delete takes name of the console and deletes it. Returns an error if one occurs.
func (c *FakeConsoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(consolesResource, c.ns, name), &v1alpha1.Console{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(consolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConsoleList{})
	return err
}

// Patch applies the patch and returns the patched console.
func (c *FakeConsoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(consolesResource, c.ns, name, data, subresources...), &v1alpha1.Console{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Console), err
}
