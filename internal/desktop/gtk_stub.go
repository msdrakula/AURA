//go:build !linux || !cgo

package desktop

import (
	"context"
	"fmt"
)

func nativeAvailable() bool { return false }

// Native reports whether this build can open a GTK/WebKit window.
func Native() bool { return nativeAvailable() }

func RunNative(ctx context.Context, title, uiURL string) error {
	return fmt.Errorf("native GTK/WebKit window requires Linux and CGO")
}
