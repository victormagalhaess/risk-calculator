package controllers_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/controllers"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/mock"
)

func TestRisk(t *testing.T) {
	t.Run("Body is empty", func(t *testing.T) {
		w := mock.ResponseWriter{}
		r := http.Request{Body: http.NoBody}
		controllers.Risk(&w, &r)
		if w.Status != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Status)
		}
	})

	t.Run("Body is unparsable", func(t *testing.T) {
		body := io.NopCloser(strings.NewReader("{unparsable: 'data'}"))
		w := mock.ResponseWriter{}
		r := http.Request{Body: body}
		controllers.Risk(&w, &r)
		if w.Status != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, w.Status)
		}
	})

	t.Run("Body is parsable", func(t *testing.T) {
		input := "{\r\n  \"age\": null,\r\n  \"dependents\": null,\r\n  \"house\": {\"ownership_status\": null},\r\n  \"income\": null,\r\n  \"marital_status\": null,\r\n  \"risk_questions\": null,\r\n  \"vehicle\": {\"year\": null}\r\n}"

		body := io.NopCloser(strings.NewReader(input))
		w := mock.ResponseWriter{}
		r := http.Request{Body: body}
		controllers.Risk(&w, &r)
		if w.Status != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Status)
		}

	})

	t.Run("Produces response", func(t *testing.T) {
		input := "{\r\n  \"age\": 35,\r\n  \"dependents\": 2,\r\n  \"house\": {\"ownership_status\": \"owned\"},\r\n  \"income\": 0,\r\n  \"marital_status\": \"married\",\r\n  \"risk_questions\": [0, 1, 0],\r\n  \"vehicle\": {\"year\": 2018}\r\n}"
		output := "{\"auto\":\"economic\",\"disability\":\"ineligible\",\"home\":\"economic\",\"life\":\"regular\"}"

		body := io.NopCloser(strings.NewReader(input))
		w := mock.ResponseWriter{}
		r := http.Request{Body: body}
		controllers.Risk(&w, &r)
		if w.Status != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Status)
		}
		if string(w.Data) != output {
			t.Errorf("Expected data %s, got %s", output, w.Data)
		}
	})

}
