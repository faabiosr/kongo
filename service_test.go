package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type ServicesTestSuite struct {
	BaseTestSuite
}

func (s *ServicesTestSuite) TestCreateReturnsHttpError() {
	s.mux.HandleFunc(servicesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	svc := &Service{Name: "test"}

	_, res, err := s.client.Services.Create(svc)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *ServicesTestSuite) TestCreate() {
	s.mux.HandleFunc(servicesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusCreated)

		file, _ := s.LoadFixture("fixtures/services_payload_foo.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Service{
		Name:     "foo-service",
		Protocol: "https",
		Host:     "foo.org",
		Path:     "/api",
	}

	svc, res, err := s.client.Services.Create(payload)

	s.assert.IsType(&Service{}, svc)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(svc.ID)
	s.assert.NotEmpty(svc.CreatedAt)
	s.assert.NotEmpty(svc.UpdatedAt)
}

func (s *ServicesTestSuite) TestCreateByURL() {
	s.mux.HandleFunc(servicesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusCreated)

		file, _ := s.LoadFixture("fixtures/services_payload_bar.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Service{
		Name: "bar-s",
		URL:  "https://bar-s.gov:9988/api",
	}

	svc, res, err := s.client.Services.CreateByURL(payload)

	s.assert.IsType(&Service{}, svc)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(svc.ID)
	s.assert.NotEmpty(svc.CreatedAt)
	s.assert.NotEmpty(svc.UpdatedAt)
	s.assert.Equal("https", svc.Protocol)
	s.assert.Equal("bar-s.gov", svc.Host)
	s.assert.Equal(9988, svc.Port)
	s.assert.Equal("/api", svc.Path)
}

func (s *ServicesTestSuite) TestListReturnsHttpError() {
	s.mux.HandleFunc(servicesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	client, _ := New(nil, s.server.URL)
	_, res, err := client.Services.List(nil)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *ServicesTestSuite) TestList() {
	s.mux.HandleFunc(servicesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		file, _ := s.LoadFixture("fixtures/services_list.json")

		io.Copy(w, file)

		defer file.Close()
	})

	services, res, err := s.client.Services.List(nil)

	s.assert.IsType(&Service{}, services[0])
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(services)
	s.assert.NotZero(services[0].ConnectTimeout)
	s.assert.NotZero(services[0].CreatedAt.Unix())
	s.assert.NotEmpty(services[0].Host)
	s.assert.NotEmpty(services[0].ID)
	s.assert.NotEmpty(services[0].Name)
	s.assert.NotEmpty(services[0].Path)
	s.assert.NotZero(services[0].Port)
	s.assert.NotEmpty(services[0].Protocol)
	s.assert.NotZero(services[0].ReadTimeout)
	s.assert.NotZero(services[0].Retries)
	s.assert.NotZero(services[0].UpdatedAt.Unix())
	s.assert.NotZero(services[0].WriteTimeout)
}

func (s *ServicesTestSuite) TestListWithOptions() {
	offset := "WyIzMzllZDk0YS03ZmJjLTQ1MTMtOGExMS03ZjEwMmYwOGVhMDUiXQ"
	options := &ListServicesOptions{Size: 1, Offset: offset}

	s.mux.HandleFunc(servicesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)
		s.assert.Equal("1", r.URL.Query().Get("size"))
		s.assert.Equal(offset, r.URL.Query().Get("offset"))

		file, _ := s.LoadFixture("fixtures/services_list.json")

		io.Copy(w, file)

		defer file.Close()
	})

	services, res, err := s.client.Services.List(options)

	s.assert.NotZero(services)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func (s *ServicesTestSuite) TestGetReturnsHttpError() {
	s.mux.HandleFunc(servicesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	_, res, err := s.client.Services.Get("test-example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *ServicesTestSuite) TestGet() {
	s.mux.HandleFunc(servicesResourcePath+"/bar-s", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/services_payload_bar.json")

		io.Copy(w, file)

		defer file.Close()
	})

	svc, res, err := s.client.Services.Get("bar-s")

	s.assert.IsType(&Service{}, svc)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(svc.ID)
	s.assert.NotEmpty(svc.CreatedAt)
	s.assert.NotEmpty(svc.UpdatedAt)
	s.assert.Equal("https", svc.Protocol)
	s.assert.Equal("bar-s.gov", svc.Host)
	s.assert.Equal(9988, svc.Port)
	s.assert.Equal("/api", svc.Path)
}

func (s *ServicesTestSuite) TestUpdateReturnsHttpError() {
	s.mux.HandleFunc(servicesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	svc := &Service{Name: "test"}

	_, res, err := s.client.Services.Update("example", svc)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *ServicesTestSuite) TestUpdate() {
	s.mux.HandleFunc(servicesResourcePath+"/bar-s", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/services_payload_foo.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Service{
		Protocol: "https",
		Host:     "foo.org",
		Path:     "/api",
	}

	svc, res, err := s.client.Services.Update("bar-s", payload)

	s.assert.IsType(&Service{}, svc)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(svc.ID)
	s.assert.NotEmpty(svc.CreatedAt)
	s.assert.NotEmpty(svc.UpdatedAt)
}

func (s *ServicesTestSuite) TestUpdateByURL() {
	s.mux.HandleFunc(servicesResourcePath+"/bar-s", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/services_payload_bar.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Service{
		Name: "bar-s",
		URL:  "https://bar-s.gov:9988/api",
	}

	svc, res, err := s.client.Services.UpdateByURL("bar-s", payload)

	s.assert.IsType(&Service{}, svc)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(svc.ID)
	s.assert.NotEmpty(svc.CreatedAt)
	s.assert.NotEmpty(svc.UpdatedAt)
	s.assert.Equal("https", svc.Protocol)
	s.assert.Equal("bar-s.gov", svc.Host)
	s.assert.Equal(9988, svc.Port)
	s.assert.Equal("/api", svc.Path)
}

func (s *ServicesTestSuite) TestDeleteReturnsHttpError() {
	s.mux.HandleFunc(servicesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	res, err := s.client.Services.Delete("example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *ServicesTestSuite) TestDelete() {
	s.mux.HandleFunc(servicesResourcePath+"/bar-s", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNoContent)
	})

	res, err := s.client.Services.Delete("bar-s")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func TestServicesTestSuite(t *testing.T) {
	suite.Run(t, new(ServicesTestSuite))
}
