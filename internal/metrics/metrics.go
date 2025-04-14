package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var NotificationErrorCount = promauto.NewCounter(prometheus.CounterOpts{
	Name: "notification_error_count",
	Help: "Count of Notification errors",
})

var NotificationSuccessCount = promauto.NewCounter(prometheus.CounterOpts{
	Name: "notification_success_count",
	Help: "Count of successfully sent Notification",
})
