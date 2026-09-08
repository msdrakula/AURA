package wordlist

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenAndLoad(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "Discovery", "Web-Content")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := "# comment\nadmin\n\nlogin\nadmin\n"
	if err := os.WriteFile(filepath.Join(dir, "common.txt"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("no"), 0o644); err != nil {
		t.Fatal(err)
	}
	lib := Open(root)
	if lib.Len() != 1 {
		t.Fatalf("len=%d catalog=%v", lib.Len(), lib.Catalog())
	}
	words, err := lib.Load("Discovery/Web-Content/common.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(words) != 2 || words[0] != "admin" || words[1] != "login" {
		t.Fatalf("words=%v", words)
	}
	if _, err := lib.Load("../etc/passwd"); err == nil {
		t.Fatal("expected path escape to fail")
	}
	got := lib.LoadFirst(nil, "missing.txt", "Discovery/Web-Content/common.txt")
	if len(got) != 2 {
		t.Fatalf("LoadFirst=%v", got)
	}
}
