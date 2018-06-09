package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type CustomersTestSuite struct {
	BaseTestSuite
}

func (s *CustomersTestSuite) TestCreateReturnsHttpError() {
	s.mux.HandleFunc(customersResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	customer := &Customer{Username: "admin"}

	_, res, err := s.client.Customers.Create(customer)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CustomersTestSuite) TestCreate() {
	s.mux.HandleFunc(customersResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusCreated)

		file, _ := s.LoadFixture("fixtures/customers_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Customer{
		Username: "admin",
		CustomId: "1",
	}

	customer, res, err := s.client.Customers.Create(payload)

	s.assert.IsType(&Customer{}, customer)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(customer.Id)
	s.assert.NotEmpty(customer.CreatedAt)
}

func (s *CustomersTestSuite) TestListReturnsHttpError() {
	s.mux.HandleFunc(customersResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	client, _ := New(nil, s.server.URL)
	_, res, err := client.Customers.List(nil)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CustomersTestSuite) TestList() {
	s.mux.HandleFunc(customersResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		file, _ := s.LoadFixture("fixtures/customers_list.json")

		io.Copy(w, file)

		defer file.Close()
	})

	customers, res, err := s.client.Customers.List(nil)

	s.assert.IsType(&Customer{}, customers[0])
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(customers)
	s.assert.NotZero(customers[0].CreatedAt.Unix())
	s.assert.NotEmpty(customers[0].Id)
	s.assert.NotEmpty(customers[0].CustomId)
	s.assert.NotEmpty(customers[0].Username)
}

func (s *CustomersTestSuite) TestListWithOptions() {
	offset := "WyIzMzllZDk0YS03ZmJjLTQ1MTMtOGExMS03ZjEwMmYwOGVhMDUiXQ"
	options := &ListCustomersOptions{
		Id:       "ec2778a3-fdf5-4901-9f76-f93a1ac1828a",
		Username: "admin",
		CustomId: "1",
		Size:     1,
		Offset:   offset,
	}

	s.mux.HandleFunc(customersResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)
		s.assert.Equal("ec2778a3-fdf5-4901-9f76-f93a1ac1828a", r.URL.Query().Get("id"))
		s.assert.Equal("admin", r.URL.Query().Get("username"))
		s.assert.Equal("1", r.URL.Query().Get("custom_id"))
		s.assert.Equal("1", r.URL.Query().Get("size"))
		s.assert.Equal(offset, r.URL.Query().Get("offset"))

		file, _ := s.LoadFixture("fixtures/customers_list.json")

		io.Copy(w, file)

		defer file.Close()
	})

	customers, res, err := s.client.Customers.List(options)

	s.assert.NotZero(customers)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func (s *CustomersTestSuite) TestGetReturnsHttpError() {
	s.mux.HandleFunc(customersResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	_, res, err := s.client.Customers.Get("test-example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CustomersTestSuite) TestGet() {
	s.mux.HandleFunc(customersResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/customers_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	customer, res, err := s.client.Customers.Get("2962c1d6b0e6")

	s.assert.IsType(&Customer{}, customer)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(customer.CreatedAt.Unix())
	s.assert.NotEmpty(customer.Id)
	s.assert.NotEmpty(customer.Username)
	s.assert.NotEmpty(customer.CustomId)
}

func (s *CustomersTestSuite) TestUpdateReturnsHttpError() {
	s.mux.HandleFunc(customersResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	customer := &Customer{Username: "admin"}

	_, res, err := s.client.Customers.Update("example", customer)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CustomersTestSuite) TestUpdate() {
	s.mux.HandleFunc(customersResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/customers_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Customer{
		Username: "admin",
	}

	customer, res, err := s.client.Customers.Update("2962c1d6b0e6", payload)

	s.assert.IsType(&Customer{}, customer)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(customer.Id)
	s.assert.NotEmpty(customer.CreatedAt)
}

func (s *CustomersTestSuite) TestDeleteReturnsHttpError() {
	s.mux.HandleFunc(customersResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	res, err := s.client.Customers.Delete("example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *CustomersTestSuite) TestDelete() {
	s.mux.HandleFunc(customersResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNoContent)
	})

	res, err := s.client.Customers.Delete("2962c1d6b0e6")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func TestCustomersTestSuite(t *testing.T) {
	suite.Run(t, new(CustomersTestSuite))
}
