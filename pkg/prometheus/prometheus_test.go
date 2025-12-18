package prometheus

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestMetricsMiddleware_RecordsMetrics(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(MetricsMiddleware)
	r.GET("/test-path", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest("GET", "/test-path", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", w.Code)
	}

	// Check that httpRequestsTotal metric was incremented for /test-path
	metricVal := testutil.ToFloat64(httpRequestsTotal.WithLabelValues("/test-path"))
	if metricVal != 1 {
		t.Errorf("Expected counter for /test-path to be 1, got %f", metricVal)
	}
}

func TestMetricsMiddleware_ActiveConnectionsGauge(t *testing.T) {
	// This test simply ensures that the Gauge is being incremented and decremented.
	initial := testutil.ToFloat64(activeConnections)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(MetricsMiddleware)
	r.GET("/conn", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest("GET", "/conn", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	final := testutil.ToFloat64(activeConnections)
	if final != initial {
		t.Errorf("Expected activeConnections to equal initial value after request, got initial=%v final=%v", initial, final)
	}
}
