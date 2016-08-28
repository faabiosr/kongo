package kongo

import (
	"net/http"
)

type ConsumersService interface {
	List() (*ConsumersList, *http.Response, error)
	Get(id string) (*Consumer, *http.Response, error)
	Delete(id string) (*http.Response, error)
	Create(consumer *Consumer) (*Consumer, *http.Response, error)
}

type ConsumersServiceOp struct {
	client *Kongo
}

type ConsumersList struct {
	Consumers []Consumer `json:"data,omitempty"`
	Total     int        `json:"total,omitempty"`
	Next      string     `json:"next,omitempty"`
}

type Consumer struct {
	CreatedAt int    `json:"created_at,omitempty"`
	CustomId  string `json:"custom_id,omitempty"`
	Id        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
}

func (c *ConsumersServiceOp) List() (*ConsumersList, *http.Response, error) {
	resource := "/consumers"

	req, err := c.client.NewRequest("GET", resource, nil)

	if err != nil {
		return nil, nil, err
	}

	consumersList := new(ConsumersList)

	res, err := c.client.Do(req, consumersList)

	if err != nil {
		return nil, res, err
	}

	return consumersList, res, nil
}

func (c *ConsumersServiceOp) Get(id string) (*Consumer, *http.Response, error) {
	resource := "/consumers/" + id

	req, err := c.client.NewRequest("GET", resource, nil)

	if err != nil {
		return nil, nil, err
	}

	consumer := new(Consumer)

	res, err := c.client.Do(req, consumer)

	if err != nil {
		return nil, res, err
	}

	return consumer, res, nil
}

func (c *ConsumersServiceOp) Delete(id string) (*http.Response, error) {
	resource := "/consumers/" + id

	req, err := c.client.NewRequest("DELETE", resource, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req, nil)

	return res, err
}

func (c *ConsumersServiceOp) Create(consumer *Consumer) (*Consumer, *http.Response, error) {
	resource := "/consumers"

	req, err := c.client.NewRequest("POST", resource, consumer)

	if err != nil {
		return nil, nil, err
	}

	res, err := c.client.Do(req, consumer)

	if err != nil {
		return nil, res, err
	}

	return consumer, res, nil
}
