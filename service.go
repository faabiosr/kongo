package kongo

import (
	"context"
	"github.com/google/go-querystring/query"
	"github.com/liip/sheriff"
	"net/http"
	"net/url"
	"path"
)

const (
	servicesResourcePath = "/services"
)

type (
	// Services manages the Kong upstream services.
	Services interface {
		// Create creates a new service.
		Create(svc *Service) (*Service, *http.Response, error)

		// CreateWithContext creates a new service.
		CreateWithContext(ctx context.Context, svc *Service) (*Service, *http.Response, error)

		// CreateByURL creates a new service by URL.
		CreateByURL(svc *Service) (*Service, *http.Response, error)

		// CreateByURLWithContext creates a new service by URL.
		CreateByURLWithContext(ctx context.Context, svc *Service) (*Service, *http.Response, error)

		// Delete deletes registered service by ID or Name.
		Delete(idOrName string) (*http.Response, error)

		// DeleteWithContext deletes registered service by ID or Name.
		DeleteWithContext(ctx context.Context, idOrName string) (*http.Response, error)

		// Get retrieves registered service by ID or Name.
		Get(idOrName string) (*Service, *http.Response, error)

		// GetWithContext retrieves registered service by ID or Name.
		GetWithContext(ctx context.Context, idOrName string) (*Service, *http.Response, error)

		// List retrieves the registered services.
		List(options *ListServicesOptions) ([]*Service, *http.Response, error)

		// ListWithContext retrieves the registered services.
		ListWithContext(ctx context.Context, options *ListServicesOptions) ([]*Service, *http.Response, error)

		// Update updates a service registered by ID or Name.
		Update(idOrName string, svc *Service) (*Service, *http.Response, error)

		// UpdateWithContext updates a service registered by ID or Name.
		UpdateWithContext(ctx context.Context, idOrName string, svc *Service) (*Service, *http.Response, error)

		// UpdateByURL updates a service registered ID or Name passing the URL.
		UpdateByURL(idOrName string, svc *Service) (*Service, *http.Response, error)

		// UpdateByURLWithContext updates a service registered by ID or Name passing the URL.
		UpdateByURLWithContext(ctx context.Context, idOrName string, svc *Service) (*Service, *http.Response, error)
	}

	// ServicesService it's a concrete instance of service.
	ServicesService struct {
		// Kongo client manages communication throught API.
		client *Kongo
	}

	// Service it's a structure of API result.
	Service struct {
		// The timeout in milliseconds for establishing a connection to the upstream server. Defaults to 60000.
		ConnectTimeout int64 `json:"connect_timeout,omitempty" groups:"create,update"`

		// The date when the service was registered.
		CreatedAt Time `json:"created_at"`

		// The host of the upstream server.
		Host string `json:"host" groups:"create,update"`

		// The identification of service registered.
		ID string `json:"id"`

		// The service name.
		Name string `json:"name" groups:"create,create_url,update,update_url"`

		// The path to be used in requests to the upstream server. Empty by default.
		Path string `json:"path,omitempty" groups:"create,update"`

		// The upstream server port. Defaults to 80.
		Port int `json:"port,omitempty" groups:"create,update"`

		// The protocol used to communicate with the upstream. It can be one of http (default) or https.
		Protocol string `json:"protocol" groups:"create,update"`

		// The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server. Defaults to 60000.
		ReadTimeout int `json:"read_timeout,omitempty" groups:"create,update"`

		// The number of retries to execute upon failure to proxy. The default is 5.
		Retries int `json:"retries,omitempty" groups:"create,update"`

		// The date when the service was updated
		UpdatedAt Time `json:"updated_at"`

		// Shorthand attribute to set protocol, host, port and path at once. This attribute is write-only (the Admin API never "returns" the url).
		URL string `json:"url,omitempty" groups:"create_url,update_url"`

		// The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server. Defaults to 60000.
		WriteTimeout int `json:"write_timeout,omitempty" groups:"create,update"`
	}

	// ServicesRoot it's a structure of API result list
	ServicesRoot struct {
		Services []*Service `json:"data"`
	}

	// ListServicesOptions stores the options you can set for requesting the service list
	ListServicesOptions struct {
		// A cursor used for pagination. offset is an object identifier that defines a place in the list.
		Offset string `url:"offset, omitempty"`

		// A limit on the number of objects to be returned per page. Defaults is 100 and max is 100.
		Size int `url:"size, omitempty"`
	}
)

// create creates a new service
func (s *ServicesService) create(ctx context.Context, svc *Service, groupName string) (*Service, *http.Response, error) {
	resource, _ := url.Parse(servicesResourcePath)

	opts := &sheriff.Options{Groups: []string{groupName}}
	body, err := sheriff.Marshal(opts, svc)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, resource, body)

	if err != nil {
		return nil, nil, err
	}

	root := new(Service)

	res, err := s.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// CreateWithContext creates a new service.
func (s *ServicesService) CreateWithContext(ctx context.Context, svc *Service) (*Service, *http.Response, error) {
	return s.create(ctx, svc, "create")
}

// Create creates a new service.
func (s *ServicesService) Create(svc *Service) (*Service, *http.Response, error) {
	return s.CreateWithContext(context.TODO(), svc)
}

// CreateByURLWithContext creates a new service by URL.
func (s *ServicesService) CreateByURLWithContext(ctx context.Context, svc *Service) (*Service, *http.Response, error) {
	return s.create(ctx, svc, "create_url")
}

// CreateByURL creates a new service by URL.
func (s *ServicesService) CreateByURL(svc *Service) (*Service, *http.Response, error) {
	return s.CreateByURLWithContext(context.TODO(), svc)
}

// DeleteWithContext deletes registered service by ID or Name.
func (s *ServicesService) DeleteWithContext(ctx context.Context, idOrName string) (*http.Response, error) {
	resource, _ := url.Parse(servicesResourcePath)
	resource.Path = path.Join(resource.Path, idOrName)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, resource, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Delete deletes registered service by ID or Name.
func (s *ServicesService) Delete(idOrName string) (*http.Response, error) {
	return s.DeleteWithContext(context.TODO(), idOrName)
}

// GetWithContext retrieves registered service by ID or Name.
func (s *ServicesService) GetWithContext(ctx context.Context, idOrName string) (*Service, *http.Response, error) {
	resource, _ := url.Parse(servicesResourcePath)
	resource.Path = path.Join(resource.Path, idOrName)

	req, err := s.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	svc := new(Service)

	res, err := s.client.Do(req, svc)

	if err != nil {
		return nil, res, err
	}

	return svc, res, nil
}

// Get retrieves registered service by ID or Name.
func (s *ServicesService) Get(idOrName string) (*Service, *http.Response, error) {
	return s.GetWithContext(context.TODO(), idOrName)
}

// ListWithContext retrieves the registered services.
func (s *ServicesService) ListWithContext(ctx context.Context, options *ListServicesOptions) ([]*Service, *http.Response, error) {
	opts, _ := query.Values(options)
	resource, _ := url.Parse(servicesResourcePath)
	resource.RawQuery = opts.Encode()

	req, err := s.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	root := new(ServicesRoot)

	res, err := s.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root.Services, res, nil
}

// List retrieves the registered services.
func (s *ServicesService) List(options *ListServicesOptions) ([]*Service, *http.Response, error) {
	return s.ListWithContext(context.TODO(), options)
}

func (s *ServicesService) update(ctx context.Context, idOrName string, svc *Service, groupName string) (*Service, *http.Response, error) {
	resource, _ := url.Parse(servicesResourcePath)
	resource.Path = path.Join(resource.Path, idOrName)

	opts := &sheriff.Options{Groups: []string{groupName}}
	body, err := sheriff.Marshal(opts, svc)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPatch, resource, body)

	if err != nil {
		return nil, nil, err
	}

	root := new(Service)

	res, err := s.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// UpdateWithContext updates a service registered by ID or Name.
func (s *ServicesService) UpdateWithContext(ctx context.Context, idOrName string, svc *Service) (*Service, *http.Response, error) {
	return s.update(ctx, idOrName, svc, "update")
}

// Update updates a service registered by ID or Name.
func (s *ServicesService) Update(idOrName string, svc *Service) (*Service, *http.Response, error) {
	return s.UpdateWithContext(context.TODO(), idOrName, svc)
}

// UpdateByURLWithContext updates a service registered by ID or Name passing the URL.
func (s *ServicesService) UpdateByURLWithContext(ctx context.Context, idOrName string, svc *Service) (*Service, *http.Response, error) {
	return s.update(ctx, idOrName, svc, "update_url")
}

// UpdateByURL updates a service registered ID or Name passing the URL.
func (s *ServicesService) UpdateByURL(idOrName string, svc *Service) (*Service, *http.Response, error) {
	return s.UpdateByURLWithContext(context.TODO(), idOrName, svc)
}
