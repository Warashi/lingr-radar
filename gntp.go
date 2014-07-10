package main

import (
	"github.com/mattn/go-gntp"
)

var gc *gntp.Client

func init() {
	gc = gntp.NewClient()
	gc.AppName = "Lingr Radar"
	gc.Register([]gntp.Notification{{
		Event: "message",
	}})
}

func notify(title, message, icon string) error {
	return gc.Notify(&gntp.Message{
		Event: "message",
		Title: title,
		Text:  message,
		Icon:  icon,
	})
}
