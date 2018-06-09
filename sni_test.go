package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type SNIsTestSuite struct {
	BaseTestSuite
}

func (s *SNIsTestSuite) TestCreateReturnsHttpError() {
	s.mux.HandleFunc(snisResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	sni := &SNI{
		Name:             "example.com",
		SSLCertificateId: "ce9832aa-5d26-41c4-af86-197c7732df1c",
	}

	_, res, err := s.client.SNIs.Create(sni)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *SNIsTestSuite) TestCreate() {
	s.mux.HandleFunc(snisResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusCreated)

		file, _ := s.LoadFixture("fixtures/snis_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &SNI{
		Name:             "example.com",
		SSLCertificateId: "ce9832aa-5d26-41c4-af86-197c7732df1c",
	}

	sni, res, err := s.client.SNIs.Create(payload)

	s.assert.IsType(&SNI{}, sni)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(sni.Name)
	s.assert.NotEmpty(sni.CreatedAt)
	s.assert.NotEmpty(sni.SSLCertificateId)
}

func (s *SNIsTestSuite) TestListReturnsHttpError() {
	s.mux.HandleFunc(snisResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	client, _ := New(nil, s.server.URL)
	_, res, err := client.SNIs.List()

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *SNIsTestSuite) TestList() {
	s.mux.HandleFunc(snisResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		file, _ := s.LoadFixture("fixtures/snis_list.json")

		io.Copy(w, file)

		defer file.Close()
	})

	snis, res, err := s.client.SNIs.List()

	s.assert.IsType(&SNI{}, snis[0])
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(snis)
	s.assert.NotZero(snis[0].CreatedAt.Unix())
	s.assert.NotEmpty(snis[0].Name)
	s.assert.NotEmpty(snis[0].SSLCertificateId)
}

func (s *SNIsTestSuite) TestGetReturnsHttpError() {
	s.mux.HandleFunc(snisResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	_, res, err := s.client.SNIs.Get("test-example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *SNIsTestSuite) TestGet() {
	s.mux.HandleFunc(snisResourcePath+"/example.com", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/snis_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	sni, res, err := s.client.SNIs.Get("example.com")

	s.assert.IsType(&SNI{}, sni)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(sni.CreatedAt.Unix())
	s.assert.NotEmpty(sni.Name)
	s.assert.NotEmpty(sni.SSLCertificateId)
}

func (s *SNIsTestSuite) TestUpdateReturnsHttpError() {
	s.mux.HandleFunc(snisResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	sni := &SNI{
		Name:             "example.com",
		SSLCertificateId: "ce9832aa-5d26-41c4-af86-197c7732df1c",
	}

	_, res, err := s.client.SNIs.Update("example", sni)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *SNIsTestSuite) TestUpdate() {
	s.mux.HandleFunc(snisResourcePath+"/example.com", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/snis_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &SNI{
		Name:             "example.com",
		SSLCertificateId: "ce9832aa-5d26-41c4-af86-197c7732df1c",
	}

	sni, res, err := s.client.SNIs.Update("example.com", payload)

	s.assert.IsType(&SNI{}, sni)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(sni.Name)
	s.assert.NotEmpty(sni.CreatedAt)
}

func (s *SNIsTestSuite) TestDeleteReturnsHttpError() {
	s.mux.HandleFunc(snisResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	res, err := s.client.SNIs.Delete("example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *SNIsTestSuite) TestDelete() {
	s.mux.HandleFunc(snisResourcePath+"/example.com", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNoContent)
	})

	res, err := s.client.SNIs.Delete("example.com")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func TestSNIsTestSuite(t *testing.T) {
	suite.Run(t, new(SNIsTestSuite))
}
