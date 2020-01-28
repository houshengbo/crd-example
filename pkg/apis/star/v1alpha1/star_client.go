package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func (c *StarV1Alpha1Client) Stars(namespace string) StarInterface {
	return &starclient{
		client: c.restClient,
		ns:     namespace,
	}
}

type StarV1Alpha1Client struct {
	restClient rest.Interface
}

type StarInterface interface {
	Create(obj *Star) (*Star, error)
	Update(obj *Star) (*Star, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	Get(name string) (*Star, error)
}

type starclient struct {
	client rest.Interface
	ns     string
}

func (c *starclient) Create(obj *Star) (*Star, error) {
	result := &Star{}
	err := c.client.Post().
		Namespace(c.ns).Resource("stars").
		Body(obj).Do().Into(result)
	return result, err
}

func (c *starclient) Update(obj *Star) (*Star, error) {
	result := &Star{}
	err := c.client.Put().
		Namespace(c.ns).Resource("stars").
		Body(obj).Do().Into(result)
	return result, err
}

func (c *starclient) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).Resource("stars").
		Name(name).Body(options).Do().
		Error()
}

func (c *starclient) Get(name string) (*Star, error) {
	result := &Star{}
	err := c.client.Get().
		Namespace(c.ns).Resource("stars").
		Name(name).Do().Into(result)
	return result, err
}
