package hey

import "testing"

func TestPush(t *testing.T) {
	Push(Notification{
		Title:    "foo",
		Body:     "bar",
		IconPath: "computer",
		Duration: DefaultDuration,
	})
}
