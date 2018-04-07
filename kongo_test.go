package kongo

import (
	"context"
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
}

func (s *KongoTestSuite) TestCreateRequestWithInvalidResourcePath() {
	_, err := s.client.NewRequest(context.TODO(), http.MethodGet, "/%1status")

	s.assert.Error(err)
}

func (s *KongoTestSuite) TestCreateRequestWithInvalidMethod() {
	_, err := s.client.NewRequest(context.TODO(), "bad method", "/status")

	s.assert.Error(err)
}

func (s *KongoTestSuite) TestCreateRequest() {
	ctx := context.TODO()
	req, _ := s.client.NewRequest(ctx, http.MethodGet, "/status")
	url, _ := url.Parse("/status")

	s.assert.NotNil(req)
	s.assert.Equal(http.MethodGet, req.Method)
	s.assert.Equal(url.Path, req.URL.Path)
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

	req, _ := s.client.NewRequest(context.TODO(), http.MethodGet, "/status")
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

	req, _ := s.client.NewRequest(context.TODO(), http.MethodGet, "/status")
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

	req, _ := s.client.NewRequest(context.TODO(), http.MethodPost, "/status")
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

	req, _ := s.client.NewRequest(context.TODO(), http.MethodHead, "/s")
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
	req, _ := client.NewRequest(context.TODO(), http.MethodGet, "/status")
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

	req, _ := s.client.NewRequest(context.TODO(), http.MethodGet, "/status")
	res, err := s.client.Do(req, &v)

	s.assert.NotNil(res)
	s.assert.Nil(err)
	s.assert.True(v.Database.Reachable)
}

func TestKongoTestSuite(t *testing.T) {
	suite.Run(t, new(KongoTestSuite))
}
