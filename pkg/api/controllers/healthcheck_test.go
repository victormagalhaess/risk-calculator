package controllers_test

//commented out integration tests

// import (
// 	"net/http"
// 	"testing"

// 	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/controllers"
// 	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/mock"
// )

// func TestHealthcheck(t *testing.T) {
// 	w := &mock.ResponseWriter{}
// 	controllers.Healthcheck(w, nil)
// 	if w.Status != http.StatusOK {
// 		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Status)
// 	}
// 	if string(w.Data) != "OK" {
// 		t.Errorf("Expected data %s, got %s", "OK", w.Data)
// 	}
// }
