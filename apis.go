package kongo

import (
	"net/http"
)

type ApisService interface {
	List() (*ApisList, *http.Response, error)
	Get(id string) (*Api, *http.Response, error)
	Delete(id string) (*http.Response, error)
	Create(api *Api) (*Api, *http.Response, error)
}

type ApisServiceOp struct {
	client *Kongo
}

type ApisList struct {
	Apis  []Api  `json:"data, omitempty"`
	Total int    `json:"total, omitempty"`
	Next  string `json:"next, omitempty"`
}

type Api struct {
	Id           string `json:"id, omitempty"`
	Name         string `json:"name, omitempty"`
	RequestHost  string `json:"request_host, omitempty"`
	UpstreamUrl  string `json:"upstream_url, omitempty"`
	PreserveHost bool   `json:"preserve_host, omitempty"`
	CreatedAt    int    `json:"created_at, omitempty"`
}

func (c *ApisServiceOp) List() (*ApisList, *http.Response, error) {
	resource := "/apis"

	req, err := c.client.NewRequest("GET", resource, nil)

	if err != nil {
		return nil, nil, err
	}

	apisList := new(ApisList)

	res, err := c.client.Do(req, apisList)

	if err != nil {
		return nil, res, err
	}

	return apisList, res, nil
}

func (c *ApisServiceOp) Get(id string) (*Api, *http.Response, error) {
	resource := "/apis/" + id

	req, err := c.client.NewRequest("GET", resource, nil)

	if err != nil {
		return nil, nil, err
	}

	api := new(Api)

	res, err := c.client.Do(req, api)

	if err != nil {
		return nil, res, err
	}

	return api, res, nil
}

func (c *ApisServiceOp) Delete(id string) (*http.Response, error) {
	resource := "/apis/" + id

	req, err := c.client.NewRequest("DELETE", resource, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req, nil)

	return res, err
}

func (c *ApisServiceOp) Create(api *Api) (*Api, *http.Response, error) {
	resource := "/apis"

	req, err := c.client.NewRequest("POST", resource, api)

	if err != nil {
		return nil, nil, err
	}

	res, err := c.client.Do(req, api)

	if err != nil {
		return nil, res, err
	}

	return api, res, nil
}
