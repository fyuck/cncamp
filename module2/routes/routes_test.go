package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	HealthCheckHandler(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok\n", w.Body.String())
}

func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name string
		testHeader string
		testEnv	string
		testResponseBody string
	}{
		{"empty version", "test1", "", "Request headers are written in response headers.\n"},
		{"version v1", "test2", "v1", "Request headers are written in response headers.\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			r.Header.Add("testHeaderKey", tt.testHeader)
			err := os.Setenv("VERSION", tt.testEnv)
			require.NoError(t, err)
			if tt.testEnv == "" {
				tt.testEnv = "emptyVersion"
			}
			HomeHandler(w, r)
			
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, tt.testHeader, w.Header().Get("testHeaderKey"))
			assert.Equal(t, tt.testEnv, w.Header().Get("VERSION"))
			assert.Equal(t, tt.testResponseBody, w.Body.String())

			t.Cleanup(func() {os.Unsetenv("VERSION")})
		})
	}
}
