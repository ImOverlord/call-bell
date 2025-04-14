package notification

import (
	"ImOverlord/call-bell/internal/metrics"
	"fmt"
	"net/http"
	"strings"
)

type Slack struct {
	Url string
}

func NewSlackNotification(url string) Notification {
	return &Slack{
		Url: url,
	}
}

func (slack *Slack) Notify(msg string) error {
	payload := strings.NewReader("{\n  \"message\": \"Attention has been requested\",\n  \"success\": \"true\"\n}")
	req, err := http.NewRequest("POST", slack.Url, payload)
	if err != nil {
		metrics.NotificationErrorCount.Inc()
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		metrics.NotificationErrorCount.Inc()
		return err
	}

	if res.StatusCode != http.StatusOK {
		metrics.NotificationErrorCount.Inc()
		return fmt.Errorf("Slack Error %d: %s", res.StatusCode, res.Status)
	}

	metrics.NotificationSuccessCount.Inc()
	return nil
}
