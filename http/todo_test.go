package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	todo "github.com/overkilling/overkill-todo-monolith-api"
	todoHttp "github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

func validateSchema(t *testing.T, document []byte) {
	schemaLoader := gojsonschema.NewReferenceLoader("file://./schemas/todos.schema.json")
	documentLoader := gojsonschema.NewBytesLoader(document)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	assert.NoError(t, err, "failed to parse")

	if !result.Valid() {
		var errors strings.Builder
		for _, error := range result.Errors() {
			errors.WriteString(fmt.Sprintf("%s\n", error.String()))
		}
		assert.Fail(t, "failed schema validation: %s", errors.String())
	}

}

func TestGetTodosHandler(t *testing.T) {
	fecher := func() ([]todo.Todo, error) {
		return []todo.Todo{
			{Todo: "Some todo"},
			{Todo: "Another todo"},
			{Todo: "Yet another todo"},
		}, nil
	}

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/todos", nil)

	todoHttp.NewTodosHandler(fecher)(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.JSONEq(t, `
		[
			{"todo": "Some todo"},
			{"todo": "Another todo"},
			{"todo": "Yet another todo"}
		]
	`, string(content))
	validateSchema(t, content)
}
