package middleware

import (
	"api-gateway/app/mock"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

type compareOriginTest struct {
	origin         string
	allowedOrigins string
	expected       string
}

var compareOriginTests = []compareOriginTest{
	{"http://localhost:3000", "http://localhost:3000", "http://localhost:3000"},
	{"http://localhost:3000", "http://localhost:3000,http://localhost:3001", "http://localhost:3000"},
	{"http://localhost:3000", "http://localhost:3001", ""},
	{"http://localhost:3000", "", "*"},
	{"", "http://localhost:3000", ""},
	{"", "", "*"},
}

func TestCompareOrigin(t *testing.T) {
	for _, testStep := range compareOriginTests {
		// prepare the necessary variables
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContext(w, "GET", gin.Params{}, nil)

		// set the origin
		ctx.Request.Header.Set("Origin", testStep.origin)

		// set the allowed origins
		os.Setenv("GATEWAY_ALLOWED_ORIGINS", testStep.allowedOrigins)

		// compare the origin
		origin := compareOrigin(ctx)
		if origin != testStep.expected {
			t.Errorf("Origin should be %s, got %s", testStep.expected, origin)
		}
	}
}

type corsTest struct {
	method         string
	origin         string
	allowedOrigins string
	expected       string
}

var corsTests = []corsTest{
	{"GET", "http://localhost:3000", "http://localhost:3000", "http://localhost:3000"},
	{"GET", "http://localhost:3000", "http://localhost:3000,http://localhost:3001", "http://localhost:3000"},
	{"GET", "http://localhost:3000", "http://localhost:3001", ""},
	{"GET", "http://localhost:3000", "", "*"},

	{"OPTIONS", "http://localhost:3000", "http://localhost:3000", "http://localhost:3000"},
	{"OPTIONS", "http://localhost:3000", "http://localhost:3000,http://localhost:3001", "http://localhost:3000"},
	{"OPTIONS", "http://localhost:3000", "http://localhost:3001", ""},
	{"OPTIONS", "http://localhost:3000", "", "*"},
}

func TestCors(t *testing.T) {
	for _, testStep := range corsTests {
		// prepare the necessary variables
		w := httptest.NewRecorder()
		ctx := mock.GetGinTestContext(w, "GET", gin.Params{}, nil)

		// set the method
		ctx.Request.Method = testStep.method

		// set the origin
		ctx.Request.Header.Set("Origin", testStep.origin)

		// set the allowed origins
		os.Setenv("GATEWAY_ALLOWED_ORIGINS", testStep.allowedOrigins)

		// call the cors middleware
		Cors()(ctx)

		// check the response
		origin := ctx.Writer.Header().Get("Access-Control-Allow-Origin")
		if origin != testStep.expected {
			t.Errorf("Origin should be %s, got %s", testStep.expected, origin)
		}
	}
}
