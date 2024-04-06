package api_test

import (
	"github.com/gkampitakis/go-snaps/snaps"
	"net/http"
	"net/http/httptest"
	"server-template/pkg/api"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetRoute(t *testing.T) {
	r := gin.Default()

	r.GET("/get/:name", api.Get)
	r.GET("/g/:name", api.Get)

	t.Run("if keys is not set, get returns 0", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/get/test?index=0", nil)
		r.ServeHTTP(w, req)

		snaps.MatchSnapshot(t, w.Body.String())
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
