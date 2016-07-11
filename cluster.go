package kongo

import (
	"net/http"
)

type ClusterService interface {
	Status() (*ClusterStatus, *http.Response, error)
	Delete(name string) (*http.Response, error)
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

func (c *ClusterServiceOp) Status() (*ClusterStatus, *http.Response, error) {
	resource := "/cluster"

	req, err := c.client.NewRequest("GET", resource, nil)

	if err != nil {
		return nil, nil, err
	}

	clusterStatus := new(ClusterStatus)

	res, err := c.client.Do(req, clusterStatus)

	if err != nil {
		return nil, res, err
	}

	return clusterStatus, res, nil
}

func (c *ClusterServiceOp) Delete(name string) (*http.Response, error) {
	resource := "/cluster/" + name

	req, err := c.client.NewRequest("DELETE", resource, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req, nil)

	return res, err
}
