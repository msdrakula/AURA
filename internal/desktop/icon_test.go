package desktop

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindIconFile(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	root := wd
	found := false
	for i := 0; i < 6; i++ {
		if _, err := os.Stat(filepath.Join(root, "web", "static", "aura-icon.svg")); err == nil {
			found = true
			break
		}
		root = filepath.Dir(root)
	}
	if !found {
		t.Fatal("could not find repo root with web/static/aura-icon.svg")
	}
	t.Setenv("AURA_ROOT", root)
	got := findIconFile()
	if got == "" {
		t.Fatal("expected icon path")
	}
	if _, err := os.Stat(got); err != nil {
		t.Fatal(err)
	}
}
