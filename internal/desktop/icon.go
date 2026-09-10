package desktop

import (
	"os"
	"path/filepath"
	"strings"
)

func findIconFile() string {
	names := []string{"aura-icon.png", "aura-icon.svg"}
	var roots []string
	if v := strings.TrimSpace(os.Getenv("AURA_ROOT")); v != "" {
		roots = append(roots, v)
	}
	if v := strings.TrimSpace(os.Getenv("MEB_ROOT")); v != "" {
		roots = append(roots, v)
	}
	if exe, err := os.Executable(); err == nil {
		roots = append(roots, filepath.Dir(exe))
		if real, err := filepath.EvalSymlinks(exe); err == nil {
			roots = append(roots, filepath.Dir(real))
		}
	}
	if wd, err := os.Getwd(); err == nil {
		roots = append(roots, wd)
	}
	seen := map[string]bool{}
	for _, root := range roots {
		abs, err := filepath.Abs(root)
		if err != nil || seen[abs] {
			continue
		}
		seen[abs] = true
		for _, name := range names {
			p := filepath.Join(abs, "web", "static", name)
			if st, err := os.Stat(p); err == nil && !st.IsDir() {
				return p
			}
		}
	}
	return ""
}
