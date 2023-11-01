//test to the Status REST API
package status_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/status"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/mock"
)

func TestBadRequest(t *testing.T) {
	w := &mock.ResponseWriter{}
	err := errors.New("error")
	status.BadRequest(w, err)
	if w.Status != http.StatusBadRequest {
		t.Errorf("Expected Status %d, got %d", http.StatusBadRequest, w.Status)
	}
	if string(w.Data) != "error" {
		t.Errorf("Expected Data %s, got %s", "error", w.Data)
	}

}

func TestServerError(t *testing.T) {
	w := &mock.ResponseWriter{}
	err := errors.New("error")
	status.ServerError(w, err)
	if w.Status != http.StatusInternalServerError {
		t.Errorf("Expected Status %d, got %d", http.StatusInternalServerError, w.Status)
	}
	if string(w.Data) != "error" {
		t.Errorf("Expected Data %s, got %s", "error", w.Data)
	}
}

func TestSuccess(t *testing.T) {
	w := &mock.ResponseWriter{}
	Data := []byte("Data")
	status.Success(w, Data)
	if w.Status != http.StatusOK {
		t.Errorf("Expected Status %d, got %d", http.StatusOK, w.Status)
	}
	if string(w.Data) != "Data" {
		t.Errorf("Expected Data %s, got %s", "Data", w.Data)
	}
}
