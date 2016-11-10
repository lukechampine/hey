hey
---

[![GoDoc](https://godoc.org/github.com/lukechampine/hey?status.svg)](https://godoc.org/github.com/lukechampine/hey)
[![Go Report Card](http://goreportcard.com/badge/github.com/lukechampine/hey)](https://goreportcard.com/report/github.com/lukechampine/hey)

hey is a tiny library for displaying notifications on Linux.

```go
hey.Push(hey.Notification{
	Title:    "foo",
	Body:     "bar",
	AppName:  "computer",
	Duration: hey.DefaultDuration,
})
```

Run `go test` to see an example notification.
