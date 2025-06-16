package metrics

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics tracks various metrics for the application
type Metrics struct {
	requestDuration *prometheus.HistogramVec
	requestTotal    *prometheus.CounterVec
	errorTotal      *prometheus.CounterVec
	activeRequests  *prometheus.GaugeVec
	workerPoolSize  prometheus.Gauge
	mu              sync.RWMutex
}

// New creates a new Metrics instance
func New() *Metrics {
	return &Metrics{
		requestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "Duration of HTTP requests in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"path", "method", "status"},
		),
		requestTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"path", "method", "status"},
		),
		errorTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_errors_total",
				Help: "Total number of HTTP errors",
			},
			[]string{"path", "method", "error_type"},
		),
		activeRequests: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "http_active_requests",
				Help: "Number of active HTTP requests",
			},
			[]string{"path"},
		),
		workerPoolSize: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "worker_pool_size",
				Help: "Current size of the worker pool",
			},
		),
	}
}

// ObserveRequest records metrics for a completed request
func (m *Metrics) ObserveRequest(path, method string, status int, duration time.Duration) {
	statusStr := strconv.Itoa(status)
	m.requestDuration.WithLabelValues(path, method, statusStr).Observe(duration.Seconds())
	m.requestTotal.WithLabelValues(path, method, statusStr).Inc()
}

// ObserveError records metrics for a request error
func (m *Metrics) ObserveError(path, method, errorType string) {
	m.errorTotal.WithLabelValues(path, method, errorType).Inc()
}

// SetActiveRequests updates the number of active requests
func (m *Metrics) SetActiveRequests(path string, count int) {
	m.activeRequests.WithLabelValues(path).Set(float64(count))
}

// SetWorkerPoolSize updates the worker pool size metric
func (m *Metrics) SetWorkerPoolSize(size int) {
	m.workerPoolSize.Set(float64(size))
}

// Handler returns an http.Handler for the metrics endpoint
func (m *Metrics) Handler() http.Handler {
	return promhttp.Handler()
}
