package kongo

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Kongo struct {
	client  *http.Client
	baseUrl string

	Node      NodeService
	Cluster   ClusterService
	Consumers ConsumersService
	Apis      ApisService
}

func New(url string) (*Kongo, error) {
	if url == "" {
		return nil, errors.New("Empty url is not allowed")
	}

	k := &Kongo{client: http.DefaultClient, baseUrl: url}
	k.Node = &NodeServiceOp{client: k}
	k.Cluster = &ClusterServiceOp{client: k}
	k.Consumers = &ConsumersServiceOp{client: k}
	k.Apis = &ApisServiceOp{client: k}

	return k, nil
}

func (k *Kongo) NewRequest(method string, resource string, body interface{}) (*http.Request, error) {
	url := k.baseUrl + resource

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (k *Kongo) Do(req *http.Request, value interface{}) (*http.Response, error) {
	res, err := k.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if value == nil {
		return res, nil
	}

	err = json.NewDecoder(res.Body).Decode(value)

	if err != nil {
		return nil, err
	}

	return res, nil
}
