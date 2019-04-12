package endpoints

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestStatusInit(t *testing.T) {

	// Setup
	e := echo.New()

	httpRequest := httptest.NewRequest("GET", "/test", bytes.NewBuffer([]byte("status test")))

	httpRequest.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpRequest, rec)

	statusAPI := &StatusEP{}

	//Test 0: Validate endpoint initialization
	assert.Nil(t, statusAPI.InitEndpoint(e, "/test_path"))

	//Test 1: Validate endpoint response
	handler := statusAPI.GetHandlerFunc()

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, c.Response().Status)
	}

}
