package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mdw-smarty/calc-lib2"
)

func NewHTTPRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("GET /add", NewHTTPHandler(&calc.Addition{}))
	router.Handle("GET /sub", NewHTTPHandler(&calc.Subtraction{}))
	router.Handle("GET /mul", NewHTTPHandler(&calc.Multiplication{}))
	router.Handle("GET /div", NewHTTPHandler(&calc.Division{}))
	return router
}

type HTTPHandler struct {
	calculator Calculator
}

func NewHTTPHandler(calculator Calculator) http.Handler {
	return &HTTPHandler{calculator: calculator}
}

func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	a, err := strconv.Atoi(query.Get("a"))
	if err != nil {
		http.Error(response, "a must be an integer", http.StatusUnprocessableEntity)
		return
	}
	b, err := strconv.Atoi(query.Get("b"))
	if err != nil {
		http.Error(response, "b must be an integer", http.StatusUnprocessableEntity)
		return
	}
	c := this.calculator.Calculate(a, b)
	response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(response, c)
	if err != nil {
		log.Println("response write err:", err)
	}
}
