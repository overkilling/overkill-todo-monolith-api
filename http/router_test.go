package http_test

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"

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
	err := waitForServer(3000)
	assert.NoError(t, err, "failed to wait for server to start")

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
			assert.NoError(t, err, "failed to create request")

			content, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "failed to read response's body")
			assert.Equal(t, tc.expected, string(content))

		})
	}
}

func testGetHandler(response string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, response)
	}
}

func waitForServer(port int) error {
	address := net.JoinHostPort("localhost", strconv.Itoa(port))

	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return err
	}
	if conn == nil {
		return errors.New("Could not connect")
	}
	return nil
}
