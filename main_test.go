package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDB struct {
	alive bool
}

func (db mockDB) Alive() bool {
	return db.alive
}

func TestUnitHealthcheckHandler(t *testing.T) {
	testCases := []struct {
		alive    bool
		expected string
	}{
		{alive: true, expected: "ok"},
		{alive: false, expected: "db down"},
	}
	for _, tC := range testCases {
		t.Run(strconv.FormatBool(tC.alive), func(t *testing.T) {
			db := mockDB{alive: tC.alive}
			res := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "http://localhost:3000/health", nil)

			healthcheckHandler(db.Alive)(res, req)

			content, _ := ioutil.ReadAll(res.Body)
			assert.Equal(t, tC.expected, string(content))
		})
	}
}
