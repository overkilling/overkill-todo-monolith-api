// +build integration

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationRouter(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/health", nil)

	router().ServeHTTP(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "ok", string(content))
}
