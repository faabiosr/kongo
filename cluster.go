package kongo

import (
	"net/http"
)

type ClusterService interface {
	Status() (*ClusterStatus, *http.Response, error)
}

type ClusterServiceOp struct {
	client *Kongo
}

type ClusterStatus struct {
	Nodes []*ClusterNode `json:"data, omitempty"`
	Total int            `json:"total, omitempty"`
}

type ClusterNode struct {
	Address string `json:"address, omitempty"`
	Name    string `json:"name, omitempty"`
	Status  string `json:"status, omitempty"`
}

func (n *ClusterServiceOp) Status() (*ClusterStatus, *http.Response, error) {
	resource := "/cluster"

	req, err := n.client.NewRequest("GET", resource)

	if err != nil {
		return nil, nil, err
	}

	clusterStatus := new(ClusterStatus)

	res, err := n.client.Do(req, clusterStatus)

	if err != nil {
		return nil, res, err
	}

	return clusterStatus, res, nil
}
