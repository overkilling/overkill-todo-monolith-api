package todo

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

func TestHealthcheckHandler(t *testing.T) {
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

			healthcheckHandler(db.Alive)(res, req)

			content, _ := ioutil.ReadAll(res.Body)
			assert.Equal(t, tC.expected, string(content))
		})
	}
}

func TestGetTodosHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/todos", nil)

	getTodosHandler()(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "[{\"todo\":\"Some task\"}]", string(content))
}
