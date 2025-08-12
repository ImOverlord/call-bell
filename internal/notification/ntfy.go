package notification

import (
	"ImOverlord/call-bell/internal/metrics"
	"context"
	"fmt"
	"net/http"
	"strings"
)

type Ntfy struct {
	URL string
}

func NewNtfyNotification(url string) Notification {
	return &Ntfy{
		URL: url,
	}
}

func (ntfy *Ntfy) Notify(msg string) error {
	payload := strings.NewReader(msg)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ntfy.URL, payload)
	if err != nil {
		metrics.NotificationErrorCount.Inc()
		return err
	}

	req.Header.Add("Title", "Call Bell")
	req.Header.Add("Priority", "urgent")
	req.Header.Add("Tags", "warning")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		metrics.NotificationErrorCount.Inc()
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		metrics.NotificationErrorCount.Inc()
		return fmt.Errorf("Ntfy Error %d: %s", res.StatusCode, res.Status)
	}

	metrics.NotificationSuccessCount.Inc()
	return nil
}
