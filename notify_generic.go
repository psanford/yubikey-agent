//go:build !linux
// +build !linux

package main

var notificationDelay = 5 * time.Second

func showNotification(message string) func() {
	clearFunc := func() {
	}

	switch runtime.GOOS {
	case "darwin":
		message = strings.ReplaceAll(message, `\`, `\\`)
		message = strings.ReplaceAll(message, `"`, `\"`)
		appleScript := `display notification "%s" with title "yubikey-agent"`
		exec.Command("osascript", "-e", fmt.Sprintf(appleScript, message)).Run()
	}

	return clearFunc
}
