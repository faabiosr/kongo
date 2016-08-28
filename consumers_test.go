package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ConsumersTestSuite struct {
	KongoTestSuite
}

func (s *ConsumersTestSuite) TestListShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	consumers := &ConsumersServiceOp{client}

	list, res, err := consumers.List()

	s.assert.Nil(list)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ConsumersTestSuite) TestListShouldRetrieveErrorWhenRequest() {
	list, res, err := s.client.Consumers.List()

	s.assert.Nil(list)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ConsumersTestSuite) TestList() {
	s.mux.HandleFunc("/consumers", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("GET", r.Method)

		response := `
		{
            "data": [
                {
                    "created_at": 1462745022000,
                    "custom_id": "9ao2",
                    "id": "9e2bcc03-ce96-4e7f-9aad-3bb5d019088d",
                    "username": "fabiorphp"
                }
            ],
            "total": 1
		}`

		fmt.Fprint(w, response)
	})

	list, res, err := s.client.Consumers.List()

	s.assert.IsType(&ConsumersList{}, list)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Len(list.Consumers, 1)
	s.assert.Equal(1462745022000, list.Consumers[0].CreatedAt)
	s.assert.Equal("9ao2", list.Consumers[0].CustomId)
	s.assert.Equal("9e2bcc03-ce96-4e7f-9aad-3bb5d019088d", list.Consumers[0].Id)
	s.assert.Equal("fabiorphp", list.Consumers[0].Username)
	s.assert.Equal(1, list.Total)
}

func (s *ConsumersTestSuite) TestGetShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	consumers := &ConsumersServiceOp{client}

	consumer, res, err := consumers.Get("9a")

	s.assert.Nil(consumer)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ConsumersTestSuite) TestGetShouldRetrieveErrorWhenRequest() {
	consumer, res, err := s.client.Consumers.Get("9b")

	s.assert.Nil(consumer)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ConsumersTestSuite) TestGet() {
	s.mux.HandleFunc("/consumers/9e2bcc03-ce96-4e7f-9aad-3bb5d019088d", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("GET", r.Method)

		response := `
		    {
	            "created_at": 1462745022000,
	            "custom_id": "9ao2",
	            "id": "9e2bcc03-ce96-4e7f-9aad-3bb5d019088d",
	            "username": "fabiorphp"
			}`

		fmt.Fprint(w, response)
	})

	consumer, res, err := s.client.Consumers.Get("9e2bcc03-ce96-4e7f-9aad-3bb5d019088d")

	s.assert.IsType(&Consumer{}, consumer)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Equal(1462745022000, consumer.CreatedAt)
	s.assert.Equal("9ao2", consumer.CustomId)
	s.assert.Equal("9e2bcc03-ce96-4e7f-9aad-3bb5d019088d", consumer.Id)
	s.assert.Equal("fabiorphp", consumer.Username)
}

func (s *ConsumersTestSuite) TestDeleteShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	consumers := &ConsumersServiceOp{client}

	res, err := consumers.Delete("9a")

	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ConsumersTestSuite) TestDelete() {
	s.mux.HandleFunc("/consumers/9e2bcc03-ce96-4e7f-9aad-3bb5d019088d", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("DELETE", r.Method)

		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, "")
	})

	res, err := s.client.Consumers.Delete("9e2bcc03-ce96-4e7f-9aad-3bb5d019088d")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func (s *ConsumersTestSuite) TestCreateShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	consumers := &ConsumersServiceOp{client}

	consumer := &Consumer{CustomId: "9ao2"}
	_, res, err := consumers.Create(consumer)

	s.assert.Empty(consumer.Id)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ConsumersTestSuite) TestCreateShouldRetrieveErrorWhenRequest() {
	consumer := &Consumer{CustomId: "9ao2"}
	_, res, err := s.client.Consumers.Create(consumer)

	s.assert.Empty(consumer.Id)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *ConsumersTestSuite) TestCreate() {
	s.mux.HandleFunc("/consumers", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("POST", r.Method)

		response := `
		    {
	            "created_at": 1462745022000,
	            "custom_id": "9ao2",
	            "id": "9e2bcc03-ce96-4e7f-9aad-3bb5d019088d"
			}`

		fmt.Fprint(w, response)
	})

	consumer := &Consumer{CustomId: "9ao2"}

	consumerResponse, res, err := s.client.Consumers.Create(consumer)

	s.assert.IsType(&Consumer{}, consumerResponse)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Equal(1462745022000, consumer.CreatedAt)
	s.assert.Equal("9ao2", consumer.CustomId)
	s.assert.Equal("9e2bcc03-ce96-4e7f-9aad-3bb5d019088d", consumer.Id)
}

func TestConsumersTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumersTestSuite))
}
