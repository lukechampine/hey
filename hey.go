package hey

import (
	"errors"
	"time"

	"github.com/godbus/dbus"
)

// DefaultDuration is the OS's default notification duration.
const DefaultDuration = -1 * time.Millisecond

// NotificationID is returned by Push and may be used to replace an existing notification
type NotificationID uint32

// A Notification represents a notification to be shown to the user.
type Notification struct {
	AppName    string // optional name of the application sending the notification
	Title      string
	Body       string
	IconPath   string
	Duration   time.Duration  // 0 means notification must be dismissed manually
	ReplacesID NotificationID // optional notification ID that this notification replaces
}

// Push displays a notification.
func Push(n Notification) (NotificationID, error) {
	if n.Title == "" {
		return 0, errors.New("notifications must have a title")
	}
	actions := []string(nil)
	hints := map[string]dbus.Variant(nil)
	expire := int32(n.Duration.Nanoseconds() / 1e6)

	conn, err := dbus.SessionBus()
	if err != nil {
		return 0, err
	}

	call := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications").Call(
		"org.freedesktop.Notifications.Notify", 0,
		n.AppName,            // STRING app_name
		uint32(n.ReplacesID), // UINT32 replaces_id
		n.IconPath,           // STRING app_icon
		n.Title,              // STRING summary
		n.Body,               // STRING body
		actions,              // ARRAY  actions
		hints,                // DICT   hints
		expire,               // INT32  expire_timeout
	)
	var notID NotificationID
	if len(call.Body) == 1 {
		if u, ok := call.Body[0].(uint32); ok {
			notID = NotificationID(u)
		}
	}
	return notID, call.Err
}
