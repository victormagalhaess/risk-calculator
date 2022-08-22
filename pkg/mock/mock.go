package mock

import "net/http"

type ResponseWriter struct {
	Status int
	Data   []byte
}

type Header map[string][]string

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.Status = statusCode
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Data = data
	return len(data), nil
}

func (w *ResponseWriter) Header() http.Header {
	return http.Header{}
}
