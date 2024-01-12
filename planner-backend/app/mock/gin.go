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

func NewTestContextBuilder(w *httptest.ResponseRecorder) *TestContextBuilder {
	/**
	* Creates a new TestContextBuilder with an empty HTTP header and URL.
	* http.Header is a map[string][]string, which is used to represent the header of an HTTP request or response.
	* url.URL is a struct that represents a parsed URL.
	**/
	t := &TestContextBuilder{
		header: make(http.Header),
		url:    &url.URL{},
	}

	t.w = w

	return t
}

func (b *TestContextBuilder) WithMethod(method string) *TestContextBuilder {
	/**
	 * Sets the method of the TestContextBuilder.
	 * The method is the HTTP method of the request (GET, POST, PUT, etc.)
	 **/
	b.method = method
	return b
}

func (b *TestContextBuilder) WithParams(params gin.Params) *TestContextBuilder {
	/**
	 * Sets the parameters of the TestContextBuilder.
	 * gin.Params is a struct that represents the parameters of a Gin request.
	 **/
	b.params = params
	return b
}

func (b *TestContextBuilder) WithBody(body interface{}) *TestContextBuilder {
	/**
	 * Sets the body of the TestContextBuilder.
	 * The body is the body of the HTTP request.
	 **/
	b.body = body
	return b
}

func (b *TestContextBuilder) WithHeader(key, value string) *TestContextBuilder {
	/**
	* Sets the header of the TestContextBuilder.
	* The header is the HTTP header of the request.
	**/
	b.header.Set(key, value)
	return b
}

func (b *TestContextBuilder) WithQueries(queries map[string]string) *TestContextBuilder {
	/**
	* Sets the queries of the TestContextBuilder.
	* The queries are the query strings of the request (e.g. ?key=value).
	**/
	q := b.url.Query()
	for key, value := range queries {
		q.Set(key, value)
	}
	b.url.RawQuery = q.Encode()

	return b
}

func (b *TestContextBuilder) WithQuery(key, value string) *TestContextBuilder {
	/**
	* Sets the query of the TestContextBuilder.
	* The query is the query string of the request (e.g. ?key=value).
	**/
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
