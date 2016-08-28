package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ApisTestSuite struct {
	KongoTestSuite
}

func (s *ApisTestSuite) TestListShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	apis := &ApisServiceOp{client}

	list, res, err := apis.List()

	s.assert.Nil(list)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ApisTestSuite) TestListShouldRetrieveErrorWhenRequest() {
	list, res, err := s.client.Apis.List()

	s.assert.Nil(list)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ApisTestSuite) TestList() {
	s.mux.HandleFunc("/apis", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("GET", r.Method)

		response := `
		{
            "data": [
                {
                    "id": "4d924084-1adb-40a5-c042-63b19db421d1",
                    "name": "Mockbin",
                    "request_host": "mockbin.com",
                    "upstream_url": "https://mockbin.com",
                    "preserve_host": false,
                    "created_at": 1422386534
                }
            ],
            "total": 1
		}`

		fmt.Fprint(w, response)
	})

	list, res, err := s.client.Apis.List()

	s.assert.IsType(&ApisList{}, list)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Len(list.Apis, 1)
	s.assert.Equal("4d924084-1adb-40a5-c042-63b19db421d1", list.Apis[0].Id)
	s.assert.Equal("Mockbin", list.Apis[0].Name)
	s.assert.Equal("mockbin.com", list.Apis[0].RequestHost)
	s.assert.Equal("https://mockbin.com", list.Apis[0].UpstreamUrl)
	s.assert.False(list.Apis[0].PreserveHost)
	s.assert.Equal(1422386534, list.Apis[0].CreatedAt)
	s.assert.Equal(1, list.Total)
}

func (s *ApisTestSuite) TestGetShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	apis := &ApisServiceOp{client}

	api, res, err := apis.Get("9a")

	s.assert.Nil(api)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ApisTestSuite) TestGetShouldRetrieveErrorWhenRequest() {
	api, res, err := s.client.Apis.Get("9b")

	s.assert.Nil(api)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ApisTestSuite) TestGet() {
	s.mux.HandleFunc("/apis/4d924084-1adb-40a5-c042-63b19db421d1", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("GET", r.Method)

		response := `
		    {
                "id": "4d924084-1adb-40a5-c042-63b19db421d1",
                "name": "Mockbin",
                "request_host": "mockbin.com",
                "upstream_url": "https://mockbin.com",
                "preserve_host": false,
                "created_at": 1422386534
			}`

		fmt.Fprint(w, response)
	})

	api, res, err := s.client.Apis.Get("4d924084-1adb-40a5-c042-63b19db421d1")

	s.assert.IsType(&Api{}, api)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Equal("4d924084-1adb-40a5-c042-63b19db421d1", api.Id)
	s.assert.Equal("Mockbin", api.Name)
	s.assert.Equal("mockbin.com", api.RequestHost)
	s.assert.Equal("https://mockbin.com", api.UpstreamUrl)
	s.assert.False(api.PreserveHost)
	s.assert.Equal(1422386534, api.CreatedAt)
}

func (s *ApisTestSuite) TestDeleteShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	apis := &ApisServiceOp{client}

	res, err := apis.Delete("9a")

	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ApisTestSuite) TestDelete() {
	s.mux.HandleFunc("/apis/4d924084-1adb-40a5-c042-63b19db421d1", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("DELETE", r.Method)

		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, "")
	})

	res, err := s.client.Apis.Delete("4d924084-1adb-40a5-c042-63b19db421d1")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func (s *ApisTestSuite) TestCreateShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	apis := &ApisServiceOp{client}

	api := &Api{Name: "mockbin", UpstreamUrl: "http://mockbin.com/", RequestHost: "mockbin.com"}
	_, res, err := apis.Create(api)

	s.assert.Empty(api.Id)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ApisTestSuite) TestCreateShouldRetrieveErrorWhenRequest() {
	api := &Api{Name: "mockbin", UpstreamUrl: "http://mockbin.com/", RequestHost: "mockbin.com"}
	_, res, err := s.client.Apis.Create(api)

	s.assert.Empty(api.Id)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ApisTestSuite) TestCreate() {
	s.mux.HandleFunc("/apis", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("POST", r.Method)

		response := `
		    {
                "id": "4d924084-1adb-40a5-c042-63b19db421d1",
                "name": "Mockbin",
                "request_host": "mockbin.com",
                "upstream_url": "https://mockbin.com",
                "created_at": 1422386534
			}`

		fmt.Fprint(w, response)
	})

	api := &Api{Name: "mockbin", UpstreamUrl: "http://mockbin.com/", RequestHost: "mockbin.com"}

	apiResponse, res, err := s.client.Apis.Create(api)

	s.assert.IsType(&Api{}, apiResponse)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Equal("4d924084-1adb-40a5-c042-63b19db421d1", api.Id)
	s.assert.Equal("Mockbin", api.Name)
	s.assert.Equal("mockbin.com", api.RequestHost)
	s.assert.Equal("https://mockbin.com", api.UpstreamUrl)
	s.assert.Equal(1422386534, api.CreatedAt)
}

func TestApisTestSuite(t *testing.T) {
	suite.Run(t, new(ApisTestSuite))
}
