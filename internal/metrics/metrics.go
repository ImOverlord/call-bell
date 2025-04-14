package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//nolint:gochecknoglobals // Prometheus metrics are global
var NotificationErrorCount = promauto.NewCounter(prometheus.CounterOpts{
	Name: "notification_error_total",
	Help: "Count of Notification errors",
})

//nolint:gochecknoglobals // Prometheus metrics are global
var NotificationSuccessCount = promauto.NewCounter(prometheus.CounterOpts{
	Name: "notification_success_total",
	Help: "Count of successfully sent Notification",
})
