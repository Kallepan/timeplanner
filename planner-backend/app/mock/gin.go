package mock

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetGinTestContext(w *httptest.ResponseRecorder, methodName string, params gin.Params) *gin.Context {
	/* GinTestContext is a function to create gin context for testing
	 * @param w is httptest.ResponseRecorder
	 * @return *gin.Context
	 */
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(w)

	header := make(http.Header)
	header.Set("Content-Type", "application/json")

	c.Request = &http.Request{
		Header: header,
		URL:    &url.URL{},
		Method: methodName,
	}

	c.Params = params

	return c
}

func GetGinTestContextWithBody(w *httptest.ResponseRecorder, methodName string, params gin.Params, body interface{}) *gin.Context {
	/* GinTestContextWithBody is a function to create gin context for testing with body
	 * @param w is httptest.ResponseRecorder
	 * @param body is string
	 * @return *gin.Context
	 */
	gin.SetMode(gin.TestMode)

	// context
	c, _ := gin.CreateTestContext(w)

	// header
	header := make(http.Header)
	header.Set("Content-Type", "application/json")

	// body
	b, err := json.Marshal(body)
	if err != nil {
		slog.Error("Error when marshalling body", "error", err)
		panic(err)
	}

	// create request
	c.Request = &http.Request{
		Header: header,
		URL:    &url.URL{},
		Method: methodName,
		Body:   io.NopCloser(bytes.NewBuffer(b)),
	}

	c.Params = params

	return c
}

func ParseJSONResponse(w *bytes.Buffer, v interface{}) {
	/* A simple helper function to parse json response
	 * @param w is bytes.Buffer
	 * @param v is interface{}
	 * @return void
	 */
	if err := json.Unmarshal(w.Bytes(), v); err != nil {
		slog.Error("Error when parsing json response", "error", err)
		panic(err)
	}
}
