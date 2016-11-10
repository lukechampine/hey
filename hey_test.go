package hey

import (
	"testing"
	"time"
)

func TestPush(t *testing.T) {
	id, err := Push(Notification{
		Title:    "foo",
		Body:     "bar",
		AppName:  "computer",
		Duration: DefaultDuration,
	})
	if err != nil {
		t.Fatal(err)
	}

	// sleep for a second, then replace the notification
	time.Sleep(time.Second)

	_, err = Push(Notification{
		Title:      "replaced!",
		Body:       "bar",
		AppName:    "computer",
		Duration:   DefaultDuration,
		ReplacesID: id,
	})
	if err != nil {
		t.Fatal(err)
	}
}
