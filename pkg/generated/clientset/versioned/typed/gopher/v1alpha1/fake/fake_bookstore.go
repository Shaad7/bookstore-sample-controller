// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/Shaad7/bookstore-sample-controller/pkg/apis/gopher/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBookstores implements BookstoreInterface
type FakeBookstores struct {
	Fake *FakeGopherV1alpha1
	ns   string
}

var bookstoresResource = schema.GroupVersionResource{Group: "gopher.com", Version: "v1alpha1", Resource: "bookstores"}

var bookstoresKind = schema.GroupVersionKind{Group: "gopher.com", Version: "v1alpha1", Kind: "Bookstore"}

// Get takes name of the bookstore, and returns the corresponding bookstore object, and an error if there is any.
func (c *FakeBookstores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Bookstore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bookstoresResource, c.ns, name), &v1alpha1.Bookstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bookstore), err
}

// List takes label and field selectors, and returns the list of Bookstores that match those selectors.
func (c *FakeBookstores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BookstoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bookstoresResource, bookstoresKind, c.ns, opts), &v1alpha1.BookstoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BookstoreList{ListMeta: obj.(*v1alpha1.BookstoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.BookstoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bookstores.
func (c *FakeBookstores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bookstoresResource, c.ns, opts))

}

// Create takes the representation of a bookstore and creates it.  Returns the server's representation of the bookstore, and an error, if there is any.
func (c *FakeBookstores) Create(ctx context.Context, bookstore *v1alpha1.Bookstore, opts v1.CreateOptions) (result *v1alpha1.Bookstore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bookstoresResource, c.ns, bookstore), &v1alpha1.Bookstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bookstore), err
}

// Update takes the representation of a bookstore and updates it. Returns the server's representation of the bookstore, and an error, if there is any.
func (c *FakeBookstores) Update(ctx context.Context, bookstore *v1alpha1.Bookstore, opts v1.UpdateOptions) (result *v1alpha1.Bookstore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bookstoresResource, c.ns, bookstore), &v1alpha1.Bookstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bookstore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBookstores) UpdateStatus(ctx context.Context, bookstore *v1alpha1.Bookstore, opts v1.UpdateOptions) (*v1alpha1.Bookstore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bookstoresResource, "status", c.ns, bookstore), &v1alpha1.Bookstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bookstore), err
}

// Delete takes name of the bookstore and deletes it. Returns an error if one occurs.
func (c *FakeBookstores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(bookstoresResource, c.ns, name, opts), &v1alpha1.Bookstore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBookstores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bookstoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BookstoreList{})
	return err
}

// Patch applies the patch and returns the patched bookstore.
func (c *FakeBookstores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Bookstore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bookstoresResource, c.ns, name, pt, data, subresources...), &v1alpha1.Bookstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bookstore), err
}