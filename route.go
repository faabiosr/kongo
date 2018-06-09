package kongo

import (
	"context"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
	"path"
)

const (
	routesResourcePath = "/routes"
)

type (
	// Routes manages the Kong route rules.
	Routes interface {
		// Create creates a new route.
		Create(route *Route) (*Route, *http.Response, error)

		// CreateWithContext creates a new route.
		CreateWithContext(ctx context.Context, route *Route) (*Route, *http.Response, error)

		// Delete deletes registered route by ID or Name.
		Delete(id string) (*http.Response, error)

		// DeleteWithContext deletes registered route by ID or Name.
		DeleteWithContext(ctx context.Context, id string) (*http.Response, error)

		// Get retrieves registered route by ID.
		Get(id string) (*Route, *http.Response, error)

		// GetWithContext retrieves registered route by ID.
		GetWithContext(ctx context.Context, id string) (*Route, *http.Response, error)

		// List retrieves a list of registered routes.
		List(options *ListRoutesOptions) ([]*Route, *http.Response, error)

		// ListWithContext retrieves a list of registered routes.
		ListWithContext(ctx context.Context, options *ListRoutesOptions) ([]*Route, *http.Response, error)

		// Update updates a route registered by ID.
		Update(id string, route *Route) (*Route, *http.Response, error)

		// UpdateWithContext updates a route registered by ID.
		UpdateWithContext(ctx context.Context, id string, route *Route) (*Route, *http.Response, error)
	}

	// RoutesService it's a concrete instance of route.
	RoutesService struct {
		// Kongo client manages communication by API.
		client *Kongo
	}

	// Route it's a structure of API result.
	Route struct {
		// The date when the route was registered.
		CreatedAt Time `json:"created_at"`

		// A list of domain names that match this Route. At least one of hosts, paths, or methods must be set.
		Hosts []string `json:"hosts,omitempty"`

		// The identification of route registered.
		Id string `json:"id"`

		// A list of HTTP methods that match this Route. At least one of hosts, paths, or methods must be set.
		Methods []string `json:"methods,omitempty"`

		// A list of paths that match this Route. At least one of hosts, paths, or methods must be set.
		Paths []string `json:"paths,omitempty"`

		// When matching a Route via one of the hosts domain names, use the request Host header in the upstream request headers.
		PreserveHost bool `json:"preserve_host,omitempty"`

		// A list of the protocols this Route should allow. By default it is ["http", "https"].
		Protocols []string `json:"protocols"`

		// The Service this Route is associated to. This is where the Route proxies traffic to.
		Service RouteService `json:"service"`

		// When matching a Route via one of the paths, strip the matching prefix from the upstream request URL.
		StripPath bool `json:"strip_path,omitempty"`

		// The date when the route was updated.
		UpdatedAt Time `json:"updated_at"`
	}

	// RouteService it's a structure of API result.
	RouteService struct {
		// Service id associated.
		Id string `json:"id"`
	}

	// RoutesRoot it's a structure of API result list.
	RoutesRoot struct {
		// List of routes.
		Routes []*Route `json:"data"`
	}

	// ListRoutesOptions stores the options you can set for requesting the route list.
	ListRoutesOptions struct {
		// A cursor used for pagination. offset is an object identifier that defines a place in the list.
		Offset string `url:"offset, omitempty"`

		// A limit on the number of objects to be returned per page. Defaults is 100 and max is 100.
		Size int `url:"size, omitempty"`
	}
)

// CreateWithContext creates a new route.
func (r *RoutesService) CreateWithContext(ctx context.Context, route *Route) (*Route, *http.Response, error) {
	resource, _ := url.Parse(routesResourcePath)

	req, err := r.client.NewRequest(ctx, http.MethodPost, resource, route)

	if err != nil {
		return nil, nil, err
	}

	root := new(Route)

	res, err := r.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Create creates a new route.
func (r *RoutesService) Create(route *Route) (*Route, *http.Response, error) {
	return r.CreateWithContext(context.TODO(), route)
}

// DeleteWithContext retrieves registered route by ID.
func (r *RoutesService) DeleteWithContext(ctx context.Context, id string) (*http.Response, error) {
	resource, _ := url.Parse(routesResourcePath)
	resource.Path = path.Join(resource.Path, id)

	req, err := r.client.NewRequest(ctx, http.MethodDelete, resource, nil)

	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}

// Delete retrieves registered route by ID or Name.
func (r *RoutesService) Delete(id string) (*http.Response, error) {
	return r.DeleteWithContext(context.TODO(), id)
}

// GetWithContext retrieves registered route by ID.
func (r *RoutesService) GetWithContext(ctx context.Context, id string) (*Route, *http.Response, error) {
	resource, _ := url.Parse(routesResourcePath)
	resource.Path = path.Join(resource.Path, id)

	req, err := r.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	route := new(Route)

	res, err := r.client.Do(req, route)

	if err != nil {
		return nil, res, err
	}

	return route, res, nil
}

// Get retrieves registered route by ID.
func (r *RoutesService) Get(id string) (*Route, *http.Response, error) {
	return r.GetWithContext(context.TODO(), id)
}

// ListWithContext retrieves a list of registered routes.
func (r *RoutesService) ListWithContext(ctx context.Context, options *ListRoutesOptions) ([]*Route, *http.Response, error) {
	opts, _ := query.Values(options)
	resource, _ := url.Parse(routesResourcePath)
	resource.RawQuery = opts.Encode()

	req, err := r.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	root := new(RoutesRoot)

	res, err := r.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root.Routes, res, nil
}

// List retrieves a list of registered routes.
func (r *RoutesService) List(options *ListRoutesOptions) ([]*Route, *http.Response, error) {
	return r.ListWithContext(context.TODO(), options)
}

// UpdateWithContext updates a route.
func (r *RoutesService) UpdateWithContext(ctx context.Context, id string, route *Route) (*Route, *http.Response, error) {
	resource, _ := url.Parse(routesResourcePath)
	resource.Path = path.Join(resource.Path, id)

	req, err := r.client.NewRequest(ctx, http.MethodPatch, resource, route)

	if err != nil {
		return nil, nil, err
	}

	root := new(Route)

	res, err := r.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Update updates a route.
func (r *RoutesService) Update(id string, route *Route) (*Route, *http.Response, error) {
	return r.UpdateWithContext(context.TODO(), id, route)
}
