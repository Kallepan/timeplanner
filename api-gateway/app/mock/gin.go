package mock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

/*
* TestContextBuilder is a builder for gin context
NewTestContextBuilder initializes a new TestContextBuilder with an empty HTTP header and URL.
WithResponseRecorder, WithMethod, WithParams, WithBody, WithHeader, and WithQuery are chainable methods that set the respective fields of the TestContextBuilder.
Build finalizes the building process. It sets the Gin mode to test, creates a new test context, sets the "Content-Type" header to "application/json", marshals the body into JSON, creates a new HTTP request with the specified method, headers, URL, and body, and assigns this request and the parameters to the test context.
*/
type TestContextBuilder struct {
	w       *httptest.ResponseRecorder
	method  string
	params  gin.Params
	body    interface{}
	header  http.Header
	request *http.Request
	url     *url.URL
}

func NewTestContextBuilder() *TestContextBuilder {
	return &TestContextBuilder{
		header: make(http.Header),
		url:    &url.URL{},
	}
}

func (b *TestContextBuilder) WithResponseRecorder(w *httptest.ResponseRecorder) *TestContextBuilder {
	b.w = w
	return b
}

func (b *TestContextBuilder) WithMethod(method string) *TestContextBuilder {
	b.method = method
	return b
}

func (b *TestContextBuilder) WithParams(params gin.Params) *TestContextBuilder {
	b.params = params
	return b
}

func (b *TestContextBuilder) WithBody(body interface{}) *TestContextBuilder {
	b.body = body
	return b
}

func (b *TestContextBuilder) WithHeader(key, value string) *TestContextBuilder {
	b.header.Set(key, value)
	return b
}

func (b *TestContextBuilder) WithQuery(key, value string) *TestContextBuilder {
	q := b.url.Query()
	q.Set(key, value)
	b.url.RawQuery = q.Encode()

	return b
}

func (b *TestContextBuilder) Build() (*gin.Context, error) {
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(b.w)

	b.header.Set("Content-Type", "application/json")

	bodyBytes, err := json.Marshal(b.body)
	if err != nil {
		return nil, err
	}

	b.request = &http.Request{
		Header: b.header,
		URL:    b.url,
		Method: b.method,
		Body:   io.NopCloser(bytes.NewBuffer(bodyBytes)),
	}

	c.Request = b.request
	c.Params = b.params

	return c, nil
}
