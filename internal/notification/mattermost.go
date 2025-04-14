package notification

import (
	"ImOverlord/call-bell/internal/metrics"
	"fmt"
	"net/http"
	"strings"
)

type Mattermost struct {
	Url string
}

func NewMattermostNotification(url string) Notification {
	return &Mattermost{
		Url: url,
	}
}

func (mattermost *Mattermost) Notify(msg string) error {
	text := fmt.Sprintf("{\n  \"text\": \"@channel %s\",\n  \"success\": \"true\"\n}", msg)
	payload := strings.NewReader(text)
	req, err := http.NewRequest("POST", mattermost.Url, payload)
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
		return fmt.Errorf("Mattermost Error %d: %s", res.StatusCode, res.Status)
	}

	metrics.NotificationSuccessCount.Inc()
	return nil
}
