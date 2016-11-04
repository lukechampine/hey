package hey

import (
	"errors"
	"time"

	"github.com/godbus/dbus"
)

// DefaultDuration is the OS's default notification duration.
const DefaultDuration = -1 * time.Millisecond

// A Notification represents a notification to be shown to the user.
type Notification struct {
	Title    string
	Body     string
	IconPath string
	Duration time.Duration // 0 means notification must be dismissed manually
}

// Push displays a notification.
func Push(n Notification) error {
	if n.Title == "" {
		return errors.New("notifications must have a title")
	}
	actions := []string(nil)
	hints := map[string]dbus.Variant(nil)
	expire := int32(n.Duration.Nanoseconds() / 1e6)

	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	defer conn.Close()

	call := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications").Call(
		"org.freedesktop.Notifications.Notify", 0,
		n.IconPath, // STRING app_icon
		uint32(0),  // UINT32 replaces_id
		"",         // STRING app_name
		n.Title,    // STRING summary
		n.Body,     // STRING body
		actions,    // ARRAY  actions
		hints,      // DICT   hints
		expire,     // INT32  expire_timeout
	)
	return call.Err
}
