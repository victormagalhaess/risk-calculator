package api_test

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api"
)

func TestApiInitializeRouter(t *testing.T) {
	a := &api.Application{}
	a.InitializeRouter()
	count := 0
	routes := []string{"/api/v1/healthcheck", "/api/v1/risk"}
	a.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		if err != nil {
			t.Errorf("Error getting path template: %v", err)
		}
		if template != routes[count] {
			t.Errorf("Expected route %v, got %v", routes[count], template)
		}
		count++
		return nil
	})
	if count != len(routes) {
		t.Errorf("Expected %v routes, got %v", len(routes), count)
	}

}

func TestApiInitializeSwagger(t *testing.T) {
	a := &api.Application{}
	a.InitializeRouter()
	a.InitializeSwagger()
	count := 0
	routes := []string{"/api/v1/healthcheck", "/api/v1/risk", "/swagger/"}
	a.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		if err != nil {
			t.Errorf("Error getting path template: %v", err)
		}
		if template != routes[count] {
			t.Errorf("Expected route %v, got %v", routes[count], template)
		}
		count++
		return nil
	})
	if count != len(routes) {
		t.Errorf("Expected %v routes, got %v", len(routes), count)
	}
}
