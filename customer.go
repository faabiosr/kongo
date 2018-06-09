package kongo

import (
	"context"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
	"path"
)

const (
	customersResourcePath = "/customers"
)

type (
	// Customers manages the Kong customer rules.
	Customers interface {
		// Create creates a new customer.
		Create(customer *Customer) (*Customer, *http.Response, error)

		// CreateWithContext creates a new customer.
		CreateWithContext(ctx context.Context, customer *Customer) (*Customer, *http.Response, error)

		// Delete deletes registered customer by ID or Username.
		Delete(idOrUsername string) (*http.Response, error)

		// DeleteWithContext deletes registered customer by ID or Username.
		DeleteWithContext(ctx context.Context, idOrUsername string) (*http.Response, error)

		// Get retrieves registered customer by ID or username.
		Get(idOrUsername string) (*Customer, *http.Response, error)

		// GetWithContext retrieves registered customer by ID or username.
		GetWithContext(ctx context.Context, idOrUsername string) (*Customer, *http.Response, error)

		// List retrieves a list of registered customers.
		List(options *ListCustomersOptions) ([]*Customer, *http.Response, error)

		// ListWithContext retrieves a list of registered customers.
		ListWithContext(ctx context.Context, options *ListCustomersOptions) ([]*Customer, *http.Response, error)

		// Update updates a customer registered by ID or Username.
		Update(idOrUsername string, customer *Customer) (*Customer, *http.Response, error)

		// UpdateWithContext updates a customer registered by ID or Username.
		UpdateWithContext(ctx context.Context, idOrUsername string, customer *Customer) (*Customer, *http.Response, error)
	}

	// CustomersService it's a concrete instance of customers.
	CustomersService struct {
		// Kongo client manages communication by API.
		client *Kongo
	}

	// Customer it's a structure of API result.
	Customer struct {
		// The date when the customer was registered.
		CreatedAt Time `json:"created_at"`

		// Field for storing an existing unique ID for the consumer. You must send either this field or username with the request.
		CustomId string `json:"custom_id,omitempty"`

		// The identification of customer registered.
		Id string `json:"id"`

		// The unique username of the consumer. You must send either this field or custom_id with the request.
		Username string `json:"username,omitempty"`
	}

	// CustomersRoot it's a structure of API result list.
	CustomersRoot struct {
		// List of customers.
		Customers []*Customer `json:"data"`
	}

	// ListCustomersOptions stores the options you can set for requesting the customer list.
	ListCustomersOptions struct {
		// A filter on the list based on the consumer custom_id field.
		CustomId string `url:"custom_id, omitempty"`

		// A filter on the list based on the consumer id field.
		Id string `url:"id, omitempty"`

		// A cursor used for pagination. offset is an object identifier that defines a place in the list.
		Offset string `url:"offset, omitempty"`

		// A limit on the number of objects to be returned per page. Defaults is 100 and max is 100.
		Size int `url:"size, omitempty"`

		// A filter on the list based on the consumer username field.
		Username string `url:"username, omitempty"`
	}
)

// CreateWithContext creates a new customer.
func (c *CustomersService) CreateWithContext(ctx context.Context, customer *Customer) (*Customer, *http.Response, error) {
	resource, _ := url.Parse(customersResourcePath)

	req, err := c.client.NewRequest(ctx, http.MethodPost, resource, customer)

	if err != nil {
		return nil, nil, err
	}

	root := new(Customer)

	res, err := c.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Create creates a new customer.
func (c *CustomersService) Create(customer *Customer) (*Customer, *http.Response, error) {
	return c.CreateWithContext(context.TODO(), customer)
}

// DeleteWithContext retrieves registered customer by ID or Username.
func (c *CustomersService) DeleteWithContext(ctx context.Context, idOrUsername string) (*http.Response, error) {
	resource, _ := url.Parse(customersResourcePath)
	resource.Path = path.Join(resource.Path, idOrUsername)

	req, err := c.client.NewRequest(ctx, http.MethodDelete, resource, nil)

	if err != nil {
		return nil, err
	}

	return c.client.Do(req, nil)
}

// Delete retrieves registered customer by ID or Username.
func (c *CustomersService) Delete(idOrUsername string) (*http.Response, error) {
	return c.DeleteWithContext(context.TODO(), idOrUsername)
}

// GetWithContext retrieves registered customer by ID or Username.
func (c *CustomersService) GetWithContext(ctx context.Context, idOrUsername string) (*Customer, *http.Response, error) {
	resource, _ := url.Parse(customersResourcePath)
	resource.Path = path.Join(resource.Path, idOrUsername)

	req, err := c.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	customer := new(Customer)

	res, err := c.client.Do(req, customer)

	if err != nil {
		return nil, res, err
	}

	return customer, res, nil
}

// Get retrieves registered customer by ID or Username.
func (c *CustomersService) Get(idOrUsername string) (*Customer, *http.Response, error) {
	return c.GetWithContext(context.TODO(), idOrUsername)
}

// ListWithContext retrieves a list of registered customers.
func (c *CustomersService) ListWithContext(ctx context.Context, options *ListCustomersOptions) ([]*Customer, *http.Response, error) {
	opts, _ := query.Values(options)
	resource, _ := url.Parse(customersResourcePath)
	resource.RawQuery = opts.Encode()

	req, err := c.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	root := new(CustomersRoot)

	res, err := c.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root.Customers, res, nil
}

// List retrieves a list of registered customers.
func (c *CustomersService) List(options *ListCustomersOptions) ([]*Customer, *http.Response, error) {
	return c.ListWithContext(context.TODO(), options)
}

// UpdateWithContext updates a customer registered by ID or Username.
func (c *CustomersService) UpdateWithContext(ctx context.Context, idOrUsername string, customer *Customer) (*Customer, *http.Response, error) {
	resource, _ := url.Parse(customersResourcePath)
	resource.Path = path.Join(resource.Path, idOrUsername)

	req, err := c.client.NewRequest(ctx, http.MethodPatch, resource, customer)

	if err != nil {
		return nil, nil, err
	}

	root := new(Customer)

	res, err := c.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Update updates a customer registered by ID or Username.
func (c *CustomersService) Update(idOrUsername string, customer *Customer) (*Customer, *http.Response, error) {
	return c.UpdateWithContext(context.TODO(), idOrUsername, customer)
}
