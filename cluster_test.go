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

func (s *ClusterTestSuite) TestDeleteShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	cluster := &ClusterServiceOp{client}

	res, err := cluster.Delete("9a")

	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ClusterTestSuite) TestDelete() {
	s.mux.HandleFunc("/cluster/7cad3ed8f8bc_0.0.0.0:7946_8d7cfba508144a5c9c471cf3b5e4ecc3", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("DELETE", r.Method)

		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, "")
	})

	res, err := s.client.Cluster.Delete("7cad3ed8f8bc_0.0.0.0:7946_8d7cfba508144a5c9c471cf3b5e4ecc3")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func TestClusterTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterTestSuite))
}
