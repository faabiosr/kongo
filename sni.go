package kongo

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

const (
	snisResourcePath = "/snis"
)

type (
	// SNIs manages the Kong SNIs.
	SNIs interface {
		// Create creates a new sni.
		Create(sni *SNI) (*SNI, *http.Response, error)

		// CreateWithContext creates a new sni.
		CreateWithContext(ctx context.Context, sni *SNI) (*SNI, *http.Response, error)

		// Delete deletes registered sni by name.
		Delete(name string) (*http.Response, error)

		// DeleteWithContext deletes registered sni by name.
		DeleteWithContext(ctx context.Context, name string) (*http.Response, error)

		// Get retrieves registered sni by name.
		Get(name string) (*SNI, *http.Response, error)

		// GetWithContext retrieves registered sni by name.
		GetWithContext(ctx context.Context, name string) (*SNI, *http.Response, error)

		// List retrieves the registered snis.
		List() ([]*SNI, *http.Response, error)

		// ListWithContext retrieves the registered snis.
		ListWithContext(ctx context.Context) ([]*SNI, *http.Response, error)

		// Update updates a sni registered by name.
		Update(name string, sni *SNI) (*SNI, *http.Response, error)

		// UpdateWithContext updates a sni registered by name.
		UpdateWithContext(ctx context.Context, name string, sni *SNI) (*SNI, *http.Response, error)
	}

	// SNIsService it's a concrete instance of SNIs.
	SNIsService struct {
		// Kongo client manages communication throught API.
		client *Kongo
	}

	// SNI it's a structure of API result.
	SNI struct {
		// The date when the sni was registered.
		CreatedAt Time `json:"created_at"`

		// The SNI name to associate with the given certificate.
		Name string `json:"name"`

		// The "id" (a UUID) of the certificate with which to associate the SNI hostname.
		SSLCertificateID string `json:"ssl_certificate_id"`
	}

	// SNIsRoot it's a structure of API result list.
	SNIsRoot struct {
		// List of snis.
		SNIs []*SNI `json:"data"`
	}
)

// CreateWithContext creates a new sni.
func (s *SNIsService) CreateWithContext(ctx context.Context, sni *SNI) (*SNI, *http.Response, error) {
	resource, _ := url.Parse(snisResourcePath)

	req, err := s.client.NewRequest(ctx, http.MethodPost, resource, sni)

	if err != nil {
		return nil, nil, err
	}

	root := new(SNI)

	res, err := s.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Create creates a new sni.
func (s *SNIsService) Create(sni *SNI) (*SNI, *http.Response, error) {
	return s.CreateWithContext(context.TODO(), sni)
}

// DeleteWithContext retrieves registered sni by ID or SNI.
func (s *SNIsService) DeleteWithContext(ctx context.Context, idOrSNI string) (*http.Response, error) {
	resource, _ := url.Parse(snisResourcePath)
	resource.Path = path.Join(resource.Path, idOrSNI)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, resource, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Delete retrieves registered sni by ID or SNI.
func (s *SNIsService) Delete(idOrSNI string) (*http.Response, error) {
	return s.DeleteWithContext(context.TODO(), idOrSNI)
}

// GetWithContext retrieves registered sni by ID or SNI.
func (s *SNIsService) GetWithContext(ctx context.Context, idOrSNI string) (*SNI, *http.Response, error) {
	resource, _ := url.Parse(snisResourcePath)
	resource.Path = path.Join(resource.Path, idOrSNI)

	req, err := s.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	sni := new(SNI)

	res, err := s.client.Do(req, sni)

	if err != nil {
		return nil, res, err
	}

	return sni, res, nil
}

// Get retrieves registered sni by ID.
func (s *SNIsService) Get(idOrSNI string) (*SNI, *http.Response, error) {
	return s.GetWithContext(context.TODO(), idOrSNI)
}

// ListWithContext retrieves the registered snis.
func (s *SNIsService) ListWithContext(ctx context.Context) ([]*SNI, *http.Response, error) {
	resource, _ := url.Parse(snisResourcePath)

	req, err := s.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	root := new(SNIsRoot)

	res, err := s.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root.SNIs, res, nil
}

// List retrieves the registered snis.
func (s *SNIsService) List() ([]*SNI, *http.Response, error) {
	return s.ListWithContext(context.TODO())
}

// UpdateWithContext updates a sni registered by ID or SNI.
func (s *SNIsService) UpdateWithContext(ctx context.Context, idOrSNI string, sni *SNI) (*SNI, *http.Response, error) {
	resource, _ := url.Parse(snisResourcePath)
	resource.Path = path.Join(resource.Path, idOrSNI)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, resource, sni)

	if err != nil {
		return nil, nil, err
	}

	root := new(SNI)

	res, err := s.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Update updates a sni registered by ID or SNI.
func (s *SNIsService) Update(idOrSNI string, sni *SNI) (*SNI, *http.Response, error) {
	return s.UpdateWithContext(context.TODO(), idOrSNI, sni)
}
