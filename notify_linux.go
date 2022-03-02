//go:build linux
// +build linux

package main

import (
	"time"

	dbus "github.com/godbus/dbus/v5"
)

var notificationDelay = 1 * time.Millisecond

func showNotification(message string) func() {
	emptyClear := func() {}

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return emptyClear
	}

	emptyClear = func() {
		conn.Close()
	}

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "yubikey-agent", uint32(0),
		"dialog-password", "yubikey-agent summary", message, []string{},
		map[string]dbus.Variant{}, int32(0))
	if call.Err != nil {
		return emptyClear
	}

	if len(call.Body) < 0 {
		return emptyClear
	}

	msgID, ok := call.Body[0].(uint32)
	if !ok {
		return emptyClear
	}

	return func() {
		obj.Call("org.freedesktop.Notifications.CloseNotification", 0, msgID)

		conn.Close()
	}
}
