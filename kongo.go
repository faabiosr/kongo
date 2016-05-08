package kongo

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Kongo struct {
	client  *http.Client
	baseUrl string

	Node    NodeService
	Cluster ClusterService
}

func New(url string) (*Kongo, error) {
	if url == "" {
		return nil, errors.New("Empty url is not allowed")
	}

	k := &Kongo{client: http.DefaultClient, baseUrl: url}
	k.Node = &NodeServiceOp{client: k}
	k.Cluster = &ClusterServiceOp{client: k}

	return k, nil
}

func (k *Kongo) NewRequest(method string, resource string) (*http.Request, error) {
	url := k.baseUrl + resource

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (k *Kongo) Do(req *http.Request, value interface{}) (*http.Response, error) {
	res, err := k.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if value == nil {
		return nil, errors.New("Value parameter is required")
	}

	err = json.NewDecoder(res.Body).Decode(value)

	if err != nil {
		return nil, err
	}

	return res, nil
}
