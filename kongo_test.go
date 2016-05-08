package kongo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type KongoTestSuite struct {
	suite.Suite

	mux    *http.ServeMux
	server *httptest.Server
	client *Kongo
	assert *assert.Assertions
}

func (s *KongoTestSuite) SetupTest() {
	s.mux = http.NewServeMux()
	s.server = httptest.NewServer(s.mux)

	client, _ := New(s.server.URL)
	s.client = client

	s.assert = assert.New(s.T())
}

func (s *KongoTestSuite) TearDownTest() {
	s.server.Close()
}

func (s *KongoTestSuite) TestFactoryWithEmptyUrl() {
	_, err := New("")

	s.assert.Error(err)
}

func (s *KongoTestSuite) TestInstance() {
	s.assert.IsType(new(Kongo), s.client)
	s.assert.Implements(new(NodeService), s.client.Node)
	s.assert.Implements(new(ClusterService), s.client.Cluster)
}

func (s *KongoTestSuite) TestCallApiWithoutRequestUrl() {
	req, _ := http.NewRequest("GET", "", nil)
	res, err := s.client.Do(req, nil)

	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *KongoTestSuite) TestCallApiWithoutValueInterface() {
	req, _ := s.client.NewRequest("GET", "/t")
	res, err := s.client.Do(req, nil)

	s.assert.Nil(res)
	s.assert.Error(err)
}

func TestKongoTestSuite(t *testing.T) {
	suite.Run(t, new(KongoTestSuite))
}
