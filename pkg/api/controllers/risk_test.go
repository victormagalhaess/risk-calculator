package controllers_test

import (
	"io/ioutil"
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
		body := ioutil.NopCloser(strings.NewReader("{unparsable: 'data'}"))
		w := mock.ResponseWriter{}
		r := http.Request{Body: body}
		controllers.Risk(&w, &r)
		if w.Status != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Status)
		}
	})

	t.Run("Body is parsable", func(t *testing.T) {
		input := "{\r\n  \"age\": 35,\r\n  \"dependents\": 2,\r\n  \"house\": {\"ownership_status\": \"owned\"},\r\n  \"income\": 0,\r\n  \"marital_status\": \"married\",\r\n  \"risk_questions\": [0, 1, 0],\r\n  \"vehicle\": {\"year\": 2018}\r\n}"
		output := "{\"auto\":\"regular\",\"disability\":\"ineligible\",\"home\":\"economic\",\"life\":\"regular\"}"

		body := ioutil.NopCloser(strings.NewReader(input))
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
