//go:build !linux || !cgo

package desktop

import "fmt"

func nativeSupported() bool { return false }

func runNative(uiURL, title string, width, height int) error {
	return fmt.Errorf("native GTK/WebKit window is only built on Linux with CGO")
}

func quitNative() {}
