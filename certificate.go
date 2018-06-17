package kongo

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

const (
	certificatesResourcePath = "/certificates"
)

type (
	// Certificates manages the Kong public certificate/private key pair for an SSL certificates.
	Certificates interface {
		// Create creates a new certificate.
		Create(certificate *Certificate) (*Certificate, *http.Response, error)

		// CreateWithContext creates a new certificate.
		CreateWithContext(ctx context.Context, certificate *Certificate) (*Certificate, *http.Response, error)

		// Delete deletes registered certificate by ID or SNI.
		Delete(idOrSNI string) (*http.Response, error)

		// DeleteWithContext deletes registered certificate by ID or SNI.
		DeleteWithContext(ctx context.Context, idOrSNI string) (*http.Response, error)

		// Get retrieves registered certificate by ID or SNI.
		Get(idOrSNI string) (*Certificate, *http.Response, error)

		// GetWithContext retrieves registered certificate by ID or SNI.
		GetWithContext(ctx context.Context, idOrSNI string) (*Certificate, *http.Response, error)

		// List retrieves the registered certificates.
		List() ([]*Certificate, *http.Response, error)

		// ListWithContext retrieves the registered certificates.
		ListWithContext(ctx context.Context) ([]*Certificate, *http.Response, error)

		// Update updates a certificate registered by ID or SNI.
		Update(idOrSNI string, certificate *Certificate) (*Certificate, *http.Response, error)

		// UpdateWithContext updates a certificate registered by ID or SNI.
		UpdateWithContext(ctx context.Context, idOrSNI string, certificate *Certificate) (*Certificate, *http.Response, error)
	}

	// CertificatesService it's a concrete instance of certificates.
	CertificatesService struct {
		// Kongo client manages communication throught API.
		client *Kongo
	}

	// Certificate it's a structure of API result.
	Certificate struct {
		// PEM-encoded public certificate of the SSL key pair.
		Cert string `json:"cert"`

		// The date when the certificate was registered.
		CreatedAt Time `json:"created_at"`

		// The identification of certificate registered.
		ID string `json:"id"`

		// PEM-encoded private key of the SSL key pair.
		Key string `json:"key"`

		// One or more hostnames to associate with this certificate as an SNI.
		SNIs []string `json:"snis,omitempty"`
	}

	// CertificatesRoot it's a structure of API result list.
	CertificatesRoot struct {
		// List of certificates.
		Certificates []*Certificate `json:"data"`
	}
)

// CreateWithContext creates a new certificate.
func (c *CertificatesService) CreateWithContext(ctx context.Context, certificate *Certificate) (*Certificate, *http.Response, error) {
	resource, _ := url.Parse(certificatesResourcePath)

	req, err := c.client.NewRequest(ctx, http.MethodPost, resource, certificate)

	if err != nil {
		return nil, nil, err
	}

	root := new(Certificate)

	res, err := c.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Create creates a new certificate.
func (c *CertificatesService) Create(certificate *Certificate) (*Certificate, *http.Response, error) {
	return c.CreateWithContext(context.TODO(), certificate)
}

// DeleteWithContext retrieves registered certificate by ID or SNI.
func (c *CertificatesService) DeleteWithContext(ctx context.Context, idOrSNI string) (*http.Response, error) {
	resource, _ := url.Parse(certificatesResourcePath)
	resource.Path = path.Join(resource.Path, idOrSNI)

	req, err := c.client.NewRequest(ctx, http.MethodDelete, resource, nil)

	if err != nil {
		return nil, err
	}

	return c.client.Do(req, nil)
}

// Delete retrieves registered certificate by ID or SNI.
func (c *CertificatesService) Delete(idOrSNI string) (*http.Response, error) {
	return c.DeleteWithContext(context.TODO(), idOrSNI)
}

// GetWithContext retrieves registered certificate by ID or SNI.
func (c *CertificatesService) GetWithContext(ctx context.Context, idOrSNI string) (*Certificate, *http.Response, error) {
	resource, _ := url.Parse(certificatesResourcePath)
	resource.Path = path.Join(resource.Path, idOrSNI)

	req, err := c.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	certificate := new(Certificate)

	res, err := c.client.Do(req, certificate)

	if err != nil {
		return nil, res, err
	}

	return certificate, res, nil
}

// Get retrieves registered certificate by ID.
func (c *CertificatesService) Get(idOrSNI string) (*Certificate, *http.Response, error) {
	return c.GetWithContext(context.TODO(), idOrSNI)
}

// ListWithContext retrieves the registered certificates.
func (c *CertificatesService) ListWithContext(ctx context.Context) ([]*Certificate, *http.Response, error) {
	resource, _ := url.Parse(certificatesResourcePath)

	req, err := c.client.NewRequest(ctx, http.MethodGet, resource, nil)

	if err != nil {
		return nil, nil, err
	}

	root := new(CertificatesRoot)

	res, err := c.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root.Certificates, res, nil
}

// List retrieves the registered certificates.
func (c *CertificatesService) List() ([]*Certificate, *http.Response, error) {
	return c.ListWithContext(context.TODO())
}

// UpdateWithContext updates a certificate registered by ID or SNI.
func (c *CertificatesService) UpdateWithContext(ctx context.Context, idOrSNI string, certificate *Certificate) (*Certificate, *http.Response, error) {
	resource, _ := url.Parse(certificatesResourcePath)
	resource.Path = path.Join(resource.Path, idOrSNI)

	req, err := c.client.NewRequest(ctx, http.MethodPatch, resource, certificate)

	if err != nil {
		return nil, nil, err
	}

	root := new(Certificate)

	res, err := c.client.Do(req, root)

	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Update updates a certificate registered by ID or SNI.
func (c *CertificatesService) Update(idOrSNI string, certificate *Certificate) (*Certificate, *http.Response, error) {
	return c.UpdateWithContext(context.TODO(), idOrSNI, certificate)
}
