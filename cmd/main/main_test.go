package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manisharigala/Foo-Datastore/api"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Get the server
	srv := api.NewServer().Router

	// Test homepage response
	request, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	srv.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)

	// Test POST API to create new entry
	message := []byte(`{"Name": "Jake"}`)
	request, err = http.NewRequest("POST", "/foo", bytes.NewBuffer(message))
	assert.Nil(t, err)
	response = httptest.NewRecorder()
	srv.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
	var foo api.Foo
	json.NewDecoder(response.Body).Decode(&foo)
	assert.NotNil(t, foo.ID)

	// Test GET API to test if the foo object created is found
	request, err = http.NewRequest("GET", "/foo/"+foo.ID, nil)
	assert.Nil(t, err)
	response = httptest.NewRecorder()
	srv.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
	id := foo.ID
	json.NewDecoder(response.Body).Decode(&foo)
	assert.NotNil(t, foo.ID)
	assert.Equal(t, foo.ID, id)

	// Test DELETE API to test deletion of foo record from memory
	request, err = http.NewRequest("DELETE", "/foo/"+id, nil)
	assert.Nil(t, err)
	response = httptest.NewRecorder()
	srv.ServeHTTP(response, request)
	assert.Equal(t, 204, response.Code)

	// Test if deleted object is still present in memory
	request, err = http.NewRequest("GET", "/foo/"+id, nil)
	assert.Nil(t, err)
	response = httptest.NewRecorder()
	srv.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code)

	// Test deleting record that is already deleted
	request, err = http.NewRequest("DELETE", "/foo/"+id, nil)
	assert.Nil(t, err)
	response = httptest.NewRecorder()
	srv.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code)
}
