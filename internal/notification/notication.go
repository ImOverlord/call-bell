package notification

type Notification interface {
	Notify(msg string) error
}
