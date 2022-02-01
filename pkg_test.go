package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", fetchTodos).Methods("GET")
	router.HandleFunc("/todo/{title}", searchTodo).Methods("GET")
	router.HandleFunc("/todo/", createTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	return router
}

//get test
func TestFetchTodos(t *testing.T) {

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "ok response is expected")
}

//search test
func TestSearchTodo(t *testing.T) {
	request, _ := http.NewRequest("GET", "/todo/team", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "ok response is expected")
}

//create test
func TestCreateTodo(t *testing.T) {
	var json = []byte(`{"title" : "meet 8:meeting time 8 o'clock "}`)

	request, _ := http.NewRequest("POST", "/todo/", bytes.NewBuffer(json))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "ok response is expected")
}

//update test
func TestUpdateTodo(t *testing.T) {

	var json = []byte(`{"title" : "meet time updated1:meet time update through testing1","completed":true}`)
	request, _ := http.NewRequest("PUT", "/todo/61f61ea3a4c4d144f18a2d97", bytes.NewBuffer(json))
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "ok response is expected")
}

//delete test
func TestDeleteTodo(t *testing.T) {
	var json = []byte(`{}`)
	request, _ := http.NewRequest("DELETE", "/todo/61f61ea3a4c4d144f18a2d97", bytes.NewBuffer(json))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "ok response is expected")
}
