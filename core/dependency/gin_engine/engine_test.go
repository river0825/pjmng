package gin_engine_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"river0825/cleanarchitecture/core/dependency/gin_engine"
	"river0825/cleanarchitecture/core/dependency/settings"
)

func TestCORSPolicy(t *testing.T) {
	const expectedOrigin = "https://service.beiqilqi.com"

	dp := &settings.Config{
		App: settings.AppConfig{
			Port:                     "8080",
			AccessControlAllowOrigin: expectedOrigin,
		},
	}
	// Initialize your GinEngine
	engine := gin_engine.NewGinEngine(
		dp,
		nil,
		nil,
		nil,
		nil,
	)

	// Create a request to pass to our handler.
	req, err := http.NewRequest(http.MethodGet, "/v2/livez", nil)
	req.Header.Add("Origin", expectedOrigin)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Create a router and attach the handler
	_ = engine.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the CORS headers
	assert.Equal(t, "https://service.beiqilqi.com", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET,POST,PUT,PATCH,DELETE,OPTIONS", rr.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "authorization, origin, content-type, accept", rr.Header().Get("Access-Control-Allow-Headers"))
}
