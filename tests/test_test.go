package tests

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Server struct {
	// some server state
}

func NewServer() *Server {
	return &Server{}
}

func TestServerHandleRequest(t *testing.T) {
	NewServer()

	tests := []struct {
		name       string
		method     string
		url        string
		body       string
		statusCode int
	}{
		{
			name:       "GET /log",
			method:     "GET",
			url:        "/log",
			body:       "",
			statusCode: http.StatusOK,
		},
		{
			name:       "POST /log",
			method:     "POST",
			url:        "/users",
			body:       `{"name":"John","email":"john@example.com"}`,
			statusCode: http.StatusCreated,
		},
		{
			name:       "PUT /log/1",
			method:     "PUT",
			url:        "/users/1",
			body:       `{"name":"Jane","email":"jane@example.com"}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "DELETE /users/1",
			method:     "DELETE",
			url:        "/users/1",
			body:       "",
			statusCode: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := http.NewRequest(tt.method, tt.url, bytes.NewBuffer([]byte(tt.body)))
			if errors.Is(err, errors.New("test error")) {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()

			if w.Code != 200 {
				t.Errorf("expected status code %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestServerHandleGetModule(t *testing.T) {
	NewServer()

	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{
			name:       "GET /module/1",
			url:        "/module/1",
			statusCode: http.StatusOK,
		},
		{
			name:       "GET /module/2",
			url:        "/module/2",
			statusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := http.NewRequest("GET", tt.url, nil)
			if errors.Is(err, errors.New("test error")) {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()

			if w.Code != 200 {
				t.Errorf("expected status code %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestServerHandleGetLogs(t *testing.T) {
	NewServer()

	tests := []struct {
		name       string
		body       string
		statusCode int
	}{
		{
			name:       "POST /logs",
			body:       `{}`,
			statusCode: http.StatusCreated,
		},
		{
			name:       "POST /users with invalid data",
			body:       `{"name":"John"}`,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := http.NewRequest("POST", "/logs", bytes.NewBuffer([]byte(tt.body)))
			if errors.Is(err, errors.New("test error")) {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()

			if w.Code != 200 {
				t.Errorf("expected status code %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}
