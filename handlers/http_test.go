package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHandler(t *testing.T) {
	assertHTTP(t, "GET", "/asdf", http.StatusNotFound, "404 page not found\n")
	assertHTTP(t, "POST", "/add?a=3&b=4", http.StatusMethodNotAllowed, "Method Not Allowed\n")
	assertHTTP(t, "GET", "/add?a=NaN&b=4", http.StatusUnprocessableEntity, "a must be an integer\n")
	assertHTTP(t, "GET", "/add?a=1&b=NaN", http.StatusUnprocessableEntity, "b must be an integer\n")
	assertHTTP(t, "GET", "/add?a=3&b=4", http.StatusOK, "7")
	assertHTTP(t, "GET", "/sub?a=4&b=3", http.StatusOK, "1")
	assertHTTP(t, "GET", "/mul?a=4&b=3", http.StatusOK, "12")
	assertHTTP(t, "GET", "/div?a=12&b=3", http.StatusOK, "4")
}

func assertHTTP(t *testing.T, method, target string, statusCode int, responseBody string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		router := NewHTTPRouter()
		response := httptest.NewRecorder()
		request := httptest.NewRequest(method, target, nil)

		router.ServeHTTP(response, request)

		assertEqual(t, response.Code, statusCode)
		assertEqual(t, response.Header().Get("Content-Type"), "text/plain; charset=utf-8")
		assertEqual(t, response.Body.String(), responseBody)
	})
}
