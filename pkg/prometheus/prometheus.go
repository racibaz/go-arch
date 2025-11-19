package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Prometheus metrics
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	activeConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Number of active connections",
		},
	)
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration, activeConnections)
}

// MetricsMiddleware Middleware to track Prometheus metrics
func MetricsMiddleware(c *gin.Context) {
	path := c.Request.URL.Path

	// begin timer to measure the requests duration
	timer := prometheus.NewTimer(httpRequestDuration.WithLabelValues(path))

	// increment total request counter
	httpRequestsTotal.WithLabelValues(path).Inc()

	// increment number of active connections
	activeConnections.Inc()

	// complete processing request
	c.Next()

	// record request duration (post processing)
	timer.ObserveDuration()

	// decrement total number of active connections (post processing)
	activeConnections.Dec()
}
