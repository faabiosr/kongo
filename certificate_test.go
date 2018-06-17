package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type CertificatesTestSuite struct {
	BaseTestSuite
}

func (s *CertificatesTestSuite) TestCreateReturnsHttpError() {
	s.mux.HandleFunc(certificatesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	certificate := &Certificate{
		Cert: "-----BEGIN CERTIFICATE-----...",
		Key:  "-----BEGIN RSA PRIVATE KEY-----...",
	}

	_, res, err := s.client.Certificates.Create(certificate)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CertificatesTestSuite) TestCreate() {
	s.mux.HandleFunc(certificatesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusCreated)

		file, _ := s.LoadFixture("fixtures/certificates_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Certificate{
		Cert: "-----BEGIN CERTIFICATE-----...",
		Key:  "-----BEGIN RSA PRIVATE KEY-----...",
		SNIs: []string{"example.com"},
	}

	certificate, res, err := s.client.Certificates.Create(payload)

	s.assert.IsType(&Certificate{}, certificate)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(certificate.ID)
	s.assert.NotEmpty(certificate.CreatedAt)
	s.assert.NotEmpty(certificate.Cert)
	s.assert.NotEmpty(certificate.Key)
	s.assert.NotZero(certificate.SNIs)
}

func (s *CertificatesTestSuite) TestListReturnsHttpError() {
	s.mux.HandleFunc(certificatesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	client, _ := New(nil, s.server.URL)
	_, res, err := client.Certificates.List()

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CertificatesTestSuite) TestList() {
	s.mux.HandleFunc(certificatesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		file, _ := s.LoadFixture("fixtures/certificates_list.json")

		io.Copy(w, file)

		defer file.Close()
	})

	certificates, res, err := s.client.Certificates.List()

	s.assert.IsType(&Certificate{}, certificates[0])
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(certificates)
	s.assert.NotZero(certificates[0].CreatedAt.Unix())
	s.assert.NotEmpty(certificates[0].Cert)
	s.assert.NotEmpty(certificates[0].ID)
	s.assert.NotZero(certificates[0].SNIs)
}

func (s *CertificatesTestSuite) TestGetReturnsHttpError() {
	s.mux.HandleFunc(certificatesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	_, res, err := s.client.Certificates.Get("test-example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CertificatesTestSuite) TestGet() {
	s.mux.HandleFunc(certificatesResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/certificates_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	certificate, res, err := s.client.Certificates.Get("2962c1d6b0e6")

	s.assert.IsType(&Certificate{}, certificate)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(certificate.CreatedAt.Unix())
	s.assert.NotEmpty(certificate.ID)
	s.assert.NotEmpty(certificate.Cert)
	s.assert.NotEmpty(certificate.Key)
	s.assert.NotZero(certificate.SNIs)
}

func (s *CertificatesTestSuite) TestUpdateReturnsHttpError() {
	s.mux.HandleFunc(certificatesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	certificate := &Certificate{
		Cert: "-----BEGIN CERTIFICATE-----...",
		Key:  "-----BEGIN RSA PRIVATE KEY-----...",
	}

	_, res, err := s.client.Certificates.Update("example", certificate)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CertificatesTestSuite) TestUpdate() {
	s.mux.HandleFunc(certificatesResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/certificates_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Certificate{
		Cert: "-----BEGIN CERTIFICATE-----...",
		Key:  "-----BEGIN RSA PRIVATE KEY-----...",
	}

	certificate, res, err := s.client.Certificates.Update("2962c1d6b0e6", payload)

	s.assert.IsType(&Certificate{}, certificate)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(certificate.ID)
	s.assert.NotEmpty(certificate.CreatedAt)
}

func (s *CertificatesTestSuite) TestDeleteReturnsHttpError() {
	s.mux.HandleFunc(certificatesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	res, err := s.client.Certificates.Delete("example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CertificatesTestSuite) TestDelete() {
	s.mux.HandleFunc(certificatesResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNoContent)
	})

	res, err := s.client.Certificates.Delete("2962c1d6b0e6")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func TestCertificatesTestSuite(t *testing.T) {
	suite.Run(t, new(CertificatesTestSuite))
}
