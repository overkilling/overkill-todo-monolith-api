package http_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	todoHttp "github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/stretchr/testify/assert"
)

type mockDB struct {
	alive bool
}

func (db mockDB) Alive() bool {
	return db.alive
}

func TestNewHealthcheckHandler(t *testing.T) {
	testCases := []struct {
		alive    bool
		expected string
	}{
		{alive: true, expected: "{\"status\":\"ok\"}"},
		{alive: false, expected: "{\"status\":\"fail\"}"},
	}
	for _, tC := range testCases {
		t.Run(strconv.FormatBool(tC.alive), func(t *testing.T) {
			db := mockDB{alive: tC.alive}
			res := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "http://localhost:3000/health", nil)

			todoHttp.NewHealthcheckHandler(db.Alive)(res, req)

			content, _ := ioutil.ReadAll(res.Body)
			assert.Equal(t, tC.expected, string(content))
		})
	}
}
