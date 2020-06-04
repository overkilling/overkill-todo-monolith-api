package http_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
	todoHttp "github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	config := todoHttp.Endpoints{
		Healthcheck: testGetHandler("healthcheck"),
		Todos:       testGetHandler("todos"),
	}
	go todoHttp.NewRouter(config).ServeOn(3000)

	tt := []struct {
		endpoint string
		expected string
	}{
		{endpoint: "/health", expected: "healthcheck"},
		{endpoint: "/todos", expected: "todos"},
	}
	for _, tc := range tt {
		t.Run(tc.endpoint, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("http://localhost:3000%s", tc.endpoint))
			if err != nil {
				panic(err)
			}

			content, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			assert.Equal(t, tc.expected, string(content))

		})
	}
}

func testGetHandler(response string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, response)
	}
}
