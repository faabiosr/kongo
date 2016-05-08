package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ClusterTestSuite struct {
	KongoTestSuite
}

func (s *ClusterTestSuite) TestStatusShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	cluster := &ClusterServiceOp{client}

	status, res, err := cluster.Status()

	s.assert.Nil(status)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ClusterTestSuite) TestStatusShouldRetrieveErrorWhenRequest() {
	status, res, err := s.client.Cluster.Status()

	s.assert.Nil(status)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ClusterTestSuite) TestStatus() {
	s.mux.HandleFunc("/cluster", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("GET", r.Method)

		response := `
		{
            "data": [
                {
                    "address": "172.17.0.3:7946",
                    "name": "dd90b6072768_0.0.0.0:7946",
                    "status": "alive"
                }
            ],
            "total": 1
		}`

		fmt.Fprint(w, response)
	})

	status, res, err := s.client.Cluster.Status()

	s.assert.IsType(&ClusterStatus{}, status)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Len(status.Nodes, 1)
	s.assert.Equal("172.17.0.3:7946", status.Nodes[0].Address)
	s.assert.Equal("dd90b6072768_0.0.0.0:7946", status.Nodes[0].Name)
	s.assert.Equal("alive", status.Nodes[0].Status)
	s.assert.Equal(1, status.Total)
}

func TestClusterTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterTestSuite))
}
