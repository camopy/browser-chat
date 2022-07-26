package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/camopy/browser-chat/app/infra/websocket/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	expRespBody    = "{\"message\":\"Hello World!\"}"
	expContentType = "application/json;charset=utf8"
)

func TestContentTypeJson(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	middleware.ContentTypeJson(http.HandlerFunc(sampleHandlerFunc())).ServeHTTP(rr, r)

	response := rr.Result()

	assert := assert.New(t)

	assert.Equal(expRespBody, rr.Body.String(), "Wrong response body")
	assert.Equal(http.StatusOK, response.StatusCode, "Wrong status code")
	assert.Equal(expContentType, response.Header.Get("Content-Type"), "Wrong content type")
}

func sampleHandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expRespBody)
	}
}
