package notification

import (
	"ImOverlord/call-bell/internal/metrics"
	"context"
	"fmt"
	"net/http"
	"strings"
)

type Slack struct {
	URL string
}

func NewSlackNotification(url string) Notification {
	return &Slack{
		URL: url,
	}
}

func (slack *Slack) Notify(msg string) error {
	text := fmt.Sprintf("{\n  \"text\": \"@channel %s\",\n  \"success\": \"true\"\n}", msg)
	payload := strings.NewReader(text)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, slack.URL, payload)
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
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		metrics.NotificationErrorCount.Inc()
		return fmt.Errorf("Slack Error %d: %s", res.StatusCode, res.Status)
	}

	metrics.NotificationSuccessCount.Inc()
	return nil
}
