package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/name", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	person := Person{}
	err := json.NewDecoder(rr.Body).Decode(&person)
	assert.Nil(t, err)
	assert.Equal(t, "Developer", person.Name)
}
