package devslog

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"testing"
)

func TestDevslog(t *testing.T) {
	h := NewHandler(os.Stdout, &Options{
		HandlerOptions: &slog.HandlerOptions{
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if !slices.Contains(groups, "request") {
					fmt.Println("not in request group", groups, a.Key)
					return a
				}
				switch a.Key {
				case "host":
					fmt.Println("dropping host")
					return slog.Attr{}
				default:
					fmt.Println("keeping", a.Key)
					return a
				}
			},
		},
	})
	slogger := slog.New(h)

	slogger.WithGroup("request").Info("Incoming request",
		slog.String("host", "example.com"),
		slog.String("path", "/"),
	)
}
