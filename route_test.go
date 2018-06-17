package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type RoutesTestSuite struct {
	BaseTestSuite
}

func (s *RoutesTestSuite) TestCreateReturnsHttpError() {
	s.mux.HandleFunc(routesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	route := &Route{Protocols: []string{"https"}}

	_, res, err := s.client.Routes.Create(route)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *RoutesTestSuite) TestCreate() {
	s.mux.HandleFunc(routesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPost, r.Method)

		w.WriteHeader(http.StatusCreated)

		file, _ := s.LoadFixture("fixtures/routes_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Route{
		Protocols: []string{"https"},
		Hosts:     []string{"foo.org"},
		Service:   RouteService{ID: "0daad537-6699-4765-baa1-dbe74a95d541"},
	}

	route, res, err := s.client.Routes.Create(payload)

	s.assert.IsType(&Route{}, route)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(route.ID)
	s.assert.NotEmpty(route.CreatedAt)
	s.assert.NotEmpty(route.UpdatedAt)
	s.assert.Equal(route.Service.ID, payload.Service.ID)
}

func (s *RoutesTestSuite) TestListReturnsHttpError() {
	s.mux.HandleFunc(routesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	client, _ := New(nil, s.server.URL)
	_, res, err := client.Routes.List(nil)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *RoutesTestSuite) TestList() {
	s.mux.HandleFunc(routesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		file, _ := s.LoadFixture("fixtures/routes_list.json")

		io.Copy(w, file)

		defer file.Close()
	})

	routes, res, err := s.client.Routes.List(nil)

	s.assert.IsType(&Route{}, routes[0])
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(routes)
	s.assert.NotZero(routes[0].CreatedAt.Unix())
	s.assert.NotZero(routes[0].Hosts)
	s.assert.NotEmpty(routes[0].ID)
	s.assert.NotZero(routes[0].Methods)
	s.assert.NotZero(routes[0].Paths)
	s.assert.False(routes[0].PreserveHost)
	s.assert.NotZero(routes[0].Protocols)
	s.assert.NotEmpty(routes[0].Service.ID)
	s.assert.True(routes[0].StripPath)
	s.assert.NotZero(routes[0].UpdatedAt.Unix())
}

func (s *RoutesTestSuite) TestListWithOptions() {
	offset := "WyIzMzllZDk0YS03ZmJjLTQ1MTMtOGExMS03ZjEwMmYwOGVhMDUiXQ"
	options := &ListRoutesOptions{Size: 1, Offset: offset}

	s.mux.HandleFunc(routesResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)
		s.assert.Equal("1", r.URL.Query().Get("size"))
		s.assert.Equal(offset, r.URL.Query().Get("offset"))

		file, _ := s.LoadFixture("fixtures/routes_list.json")

		io.Copy(w, file)

		defer file.Close()
	})

	routes, res, err := s.client.Routes.List(options)

	s.assert.NotZero(routes)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func (s *RoutesTestSuite) TestGetReturnsHttpError() {
	s.mux.HandleFunc(routesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	_, res, err := s.client.Routes.Get("test-example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *RoutesTestSuite) TestGet() {
	s.mux.HandleFunc(routesResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/routes_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	route, res, err := s.client.Routes.Get("2962c1d6b0e6")

	s.assert.IsType(&Route{}, route)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(route.CreatedAt.Unix())
	s.assert.NotZero(route.Hosts)
	s.assert.NotEmpty(route.ID)
	s.assert.NotZero(route.Methods)
	s.assert.NotZero(route.Paths)
	s.assert.False(route.PreserveHost)
	s.assert.NotZero(route.Protocols)
	s.assert.NotEmpty(route.Service.ID)
	s.assert.True(route.StripPath)
	s.assert.NotZero(route.UpdatedAt.Unix())
}

func (s *RoutesTestSuite) TestUpdateReturnsHttpError() {
	s.mux.HandleFunc(routesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	route := &Route{Protocols: []string{"http"}}

	_, res, err := s.client.Routes.Update("example", route)

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *RoutesTestSuite) TestUpdate() {
	s.mux.HandleFunc(routesResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodPatch, r.Method)

		w.WriteHeader(http.StatusOK)

		file, _ := s.LoadFixture("fixtures/routes_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	payload := &Route{
		Paths: []string{"/api"},
	}

	route, res, err := s.client.Routes.Update("2962c1d6b0e6", payload)

	s.assert.IsType(&Route{}, route)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotEmpty(route.ID)
	s.assert.NotEmpty(route.CreatedAt)
	s.assert.NotEmpty(route.UpdatedAt)
}

func (s *RoutesTestSuite) TestDeleteReturnsHttpError() {
	s.mux.HandleFunc(routesResourcePath+"/example", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "")
	})

	res, err := s.client.Routes.Delete("example")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *RoutesTestSuite) TestDelete() {
	s.mux.HandleFunc(routesResourcePath+"/2962c1d6b0e6", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNoContent)
	})

	res, err := s.client.Routes.Delete("2962c1d6b0e6")

	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)
}

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}
