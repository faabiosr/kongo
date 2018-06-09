package kongo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

type (
	BaseTestSuite struct {
		suite.Suite

		assert  *assert.Assertions
		client  *Kongo
		baseUrl string

		mux    *http.ServeMux
		server *httptest.Server
	}

	KongoTestSuite struct {
		BaseTestSuite
	}

	MockData struct {
		Created Time `json:"created"`
	}
)

func (s *BaseTestSuite) SetupTest() {
	s.mux = http.NewServeMux()
	s.server = httptest.NewServer(s.mux)

	client, _ := New(nil, s.server.URL)
	s.client = client

	s.assert = assert.New(s.T())
}

func (s *BaseTestSuite) TearDownTest() {
	s.server.Close()
}

func (s *BaseTestSuite) LoadFixture(filePath string) (io.ReadCloser, error) {
	filename, err := filepath.Abs(filePath)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *KongoTestSuite) TestFactoryClientWithEmptyURL() {
	_, err := NewClient(nil, nil)

	s.assert.EqualError(err, "Empty URL is not allowed")
}

func (s *KongoTestSuite) TestFactoryWithEmptyURL() {
	_, err := New(nil, "")

	s.assert.EqualError(err, "Empty URL is not allowed")
}

func (s *KongoTestSuite) TestFactoryWithInvalidURL() {
	_, err := New(nil, "http://192.168.1.%1/")

	s.assert.Error(err)
}

func (s *KongoTestSuite) TestInstance() {
	s.assert.IsType(new(Kongo), s.client)
	s.assert.Equal(userAgent, s.client.UserAgent)
	s.assert.Implements(new(Node), s.client.Node)
	s.assert.Implements(new(Services), s.client.Services)
	s.assert.Implements(new(Routes), s.client.Routes)
	s.assert.Implements(new(Customers), s.client.Customers)
}

func (s *KongoTestSuite) TestCreateRequestWithInvalidMethod() {
	resource, _ := url.Parse("/status")
	_, err := s.client.NewRequest(context.TODO(), "bad method", resource, nil)

	s.assert.Error(err)
}

func (s *KongoTestSuite) TestCreateRequest() {
	ctx := context.TODO()
	resource, _ := url.Parse("/status")
	req, _ := s.client.NewRequest(ctx, http.MethodGet, resource, nil)

	s.assert.NotNil(req)
	s.assert.Equal(http.MethodGet, req.Method)
	s.assert.Equal(resource.Path, req.URL.Path)
	s.assert.Equal(ctx, req.Context())
	s.assert.Equal(mediaType, req.Header.Get("Accept"))
	s.assert.Equal(mediaType, req.Header.Get("Content-Type"))
	s.assert.Equal(userAgent, req.Header.Get("User-Agent"))
}

func (s *KongoTestSuite) TestCallApiWithoutRequestUrl() {
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	res, err := s.client.Do(req, nil)

	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *KongoTestSuite) TestCallApiResourceWithoutValueReturns() {
	s.mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		fmt.Fprint(w, `{"status": true}`)
	})

	resource, _ := url.Parse("/status")
	req, _ := s.client.NewRequest(context.TODO(), http.MethodGet, resource, nil)
	res, err := s.client.Do(req, nil)

	s.assert.NotNil(res)
	s.assert.Nil(err)
}

func (s *KongoTestSuite) TestCallApiResourceWithWrongValueJsonStruct() {
	s.mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		fmt.Fprint(w, `{"database": {"reachable": true}}`)
	})

	v := struct {
		Database struct {
			Reachable string `json:"reachable"`
		} `json:"database"`
	}{}

	resource, _ := url.Parse("/status")
	req, _ := s.client.NewRequest(context.TODO(), http.MethodGet, resource, nil)
	res, err := s.client.Do(req, &v)

	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *KongoTestSuite) TestCallApiWhenReturnsHttpErrors() {
	s.mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusMethodNotAllowed)

		fmt.Fprint(w, `{"message": "Method not allowed"}`)
	})

	resource, _ := url.Parse("/status")
	req, _ := s.client.NewRequest(context.TODO(), http.MethodPost, resource, nil)
	res, err := s.client.Do(req, nil)

	s.assert.NotNil(res)
	s.assert.EqualError(err, "405 Method not allowed")
}

func (s *KongoTestSuite) TestCallApiWhenReturnsHttpErrorWithEmptyBody() {
	s.mux.HandleFunc("/s", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodHead, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, ``)
	})

	resource, _ := url.Parse("/s")
	req, _ := s.client.NewRequest(context.TODO(), http.MethodHead, resource, nil)
	res, err := s.client.Do(req, nil)

	s.assert.NotNil(res)
	s.assert.EqualError(err, "404 Request error")
}

func (s *KongoTestSuite) TestCallApiWhenReturnsHttpErrorWithNonJsonBody() {
	s.mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, `Something wrong`)
	})

	client, _ := New(nil, s.server.URL)
	resource, _ := url.Parse("/status")
	req, _ := client.NewRequest(context.TODO(), http.MethodGet, resource, nil)
	res, err := client.Do(req, nil)

	s.assert.NotNil(res)
	s.assert.EqualError(err, "400 Something wrong")
}

func (s *KongoTestSuite) TestCallApiResource() {
	s.mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		fmt.Fprint(w, `{"database": {"reachable": true}}`)
	})

	v := struct {
		Database struct {
			Reachable bool `json:"reachable"`
		} `json:"database"`
	}{}

	resource, _ := url.Parse("/status")
	req, _ := s.client.NewRequest(context.TODO(), http.MethodGet, resource, nil)
	res, err := s.client.Do(req, &v)

	s.assert.NotNil(res)
	s.assert.Nil(err)
	s.assert.True(v.Database.Reachable)
}

func (s *KongoTestSuite) TestCallApiResourceWithBody() {
	s.mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		var body map[string]string

		json.NewDecoder(r.Body).Decode(&body)

		s.assert.Equal("test", body["name"])

		fmt.Fprint(w, `{"status": true}`)
	})

	v := struct {
		Status bool `json:"status"`
	}{}

	resource, _ := url.Parse("/status")
	body := map[string]string{"name": "test"}

	client, _ := New(nil, s.server.URL)
	req, _ := client.NewRequest(context.TODO(), http.MethodGet, resource, body)
	res, err := s.client.Do(req, &v)

	s.assert.NotNil(res)
	s.assert.Nil(err)
	s.assert.True(v.Status)
}

func (s *KongoTestSuite) TestJSONTimeParsingWithEmptyValue() {
	var data MockData

	json.Unmarshal(
		[]byte(`{"created": ""}`),
		&data,
	)

	s.assert.Equal("0001-01-01", data.Created.Format("2006-01-02"))
}

func (s *KongoTestSuite) TestJSONTimeParsingWithWrongData() {
	var data MockData

	err := json.Unmarshal(
		[]byte(`{"created": {}}`),
		&data,
	)

	s.assert.Error(err)
}

func (s *KongoTestSuite) TestJSONTimeParsing() {
	var data MockData

	json.Unmarshal(
		[]byte(`{"created": "1522832400"}`),
		&data,
	)

	s.assert.Equal("2018-04-04", data.Created.Format("2006-01-02"))
}

func TestKongoTestSuite(t *testing.T) {
	suite.Run(t, new(KongoTestSuite))
}
