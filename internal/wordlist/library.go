// Package wordlist indexes SecLists and the small built-in fallbacks.
package wordlist

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

// Entry is one dictionary file under the SecLists root.
type Entry struct {
	Path     string `json:"path"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Size     int64  `json:"size"`
	Lines    int    `json:"lines"`
}

// Library is a read-only index of every wordlist on disk.
type Library struct {
	Root string
	mu   sync.RWMutex
	all  []Entry
}

var skipExt = map[string]struct{}{
	".md": {}, ".git": {}, ".png": {}, ".jpg": {}, ".jpeg": {}, ".gif": {},
	".svg": {}, ".webp": {}, ".ico": {}, ".pdf": {}, ".zip": {}, ".gz": {},
	".tgz": {}, ".bz2": {}, ".xz": {}, ".7z": {}, ".rar": {}, ".exe": {},
	".dll": {}, ".so": {}, ".dylib": {}, ".pcap": {}, ".cap": {}, ".mp4": {},
	".html": {}, ".htm": {}, ".css": {}, ".js": {}, ".json": {}, ".xml": {},
	".yml": {}, ".yaml": {}, ".toml": {}, ".go": {}, ".py": {}, ".sh": {},
	".php": {}, ".asp": {}, ".aspx": {}, ".jsp": {}, ".war": {}, ".jar": {},
	".class": {}, ".bin": {}, ".dat": {}, ".db": {}, ".sqlite": {},
}

var skipNames = map[string]struct{}{
	"license": {}, "license.txt": {}, "readme": {}, "readme.md": {},
	"contributing.md": {}, "changelog.md": {}, ".gitignore": {},
	".gitattributes": {}, "makefile": {}, "dockerfile": {},
}

// Open indexes every dictionary under root. Missing root yields an empty library.
func Open(root string) *Library {
	lib := &Library{Root: root}
	if root == "" {
		return lib
	}
	st, err := os.Stat(root)
	if err != nil || !st.IsDir() {
		return lib
	}
	var out []Entry
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		name := d.Name()
		if d.IsDir() {
			if name == ".git" || name == ".github" {
				return fs.SkipDir
			}
			return nil
		}
		low := strings.ToLower(name)
		if _, skip := skipNames[low]; skip {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(name))
		if _, skip := skipExt[ext]; skip {
			return nil
		}
		info, err := d.Info()
		if err != nil || info.Size() == 0 {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)
		cat, _, _ := strings.Cut(rel, "/")
		if cat == rel {
			cat = "root"
		}
		e := Entry{
			Path:     rel,
			Name:     name,
			Category: cat,
			Size:     info.Size(),
			Lines:    -1,
		}
		out = append(out, e)
		return nil
	})
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	lib.all = out
	return lib
}

// Catalog returns a copy of the index.
func (l *Library) Catalog() []Entry {
	if l == nil {
		return nil
	}
	l.mu.RLock()
	defer l.mu.RUnlock()
	out := make([]Entry, len(l.all))
	copy(out, l.all)
	return out
}

// Len is the number of indexed dictionaries.
func (l *Library) Len() int {
	if l == nil {
		return 0
	}
	return len(l.all)
}

// Resolve maps a relative catalog path to an absolute file inside the root.
func (l *Library) Resolve(rel string) (string, error) {
	if l == nil || l.Root == "" {
		return "", fmt.Errorf("seclists is not installed")
	}
	rel = filepath.ToSlash(strings.TrimSpace(rel))
	rel = strings.TrimPrefix(rel, "/")
	if rel == "" || strings.Contains(rel, "..") {
		return "", fmt.Errorf("invalid wordlist path")
	}
	abs := filepath.Join(l.Root, filepath.FromSlash(rel))
	abs, err := filepath.Abs(abs)
	if err != nil {
		return "", err
	}
	root, err := filepath.Abs(l.Root)
	if err != nil {
		return "", err
	}
	if abs != root && !strings.HasPrefix(abs, root+string(os.PathSeparator)) {
		return "", fmt.Errorf("wordlist path escapes seclists root")
	}
	st, err := os.Stat(abs)
	if err != nil {
		return "", fmt.Errorf("wordlist not found: %s", rel)
	}
	if st.IsDir() {
		return "", fmt.Errorf("not a file: %s", rel)
	}
	return abs, nil
}

// Load reads every non-empty, non-comment line from a catalog path.
func (l *Library) Load(rel string) ([]string, error) {
	abs, err := l.Resolve(rel)
	if err != nil {
		return nil, err
	}
	return ReadFile(abs)
}

// LoadFirst returns the first existing list among rel paths, or fallback.
func (l *Library) LoadFirst(fallback []string, rels ...string) []string {
	if l != nil {
		for _, rel := range rels {
			words, err := l.Load(rel)
			if err == nil && len(words) > 0 {
				return words
			}
		}
	}
	return fallback
}

// Preview returns up to n lines from a list.
func (l *Library) Preview(rel string, n int) ([]string, error) {
	words, err := l.Load(rel)
	if err != nil {
		return nil, err
	}
	if n > 0 && len(words) > n {
		return words[:n], nil
	}
	return words, nil
}

// ReadFile loads a wordlist from an absolute path.
func ReadFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	var out []string
	seen := map[string]struct{}{}
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if _, ok := seen[line]; ok {
			continue
		}
		seen[line] = struct{}{}
		out = append(out, line)
	}
	if err := sc.Err(); err != nil {
		return out, err
	}
	return out, nil
}

// Defaults are preferred SecLists files for each job. Missing files fall back
// to the tiny built-in lists.
var (
	DefaultDirs = []string{
		"Discovery/Web-Content/common.txt",
		"Discovery/Web-Content/raft-small-directories.txt",
	}
	DefaultParams = []string{
		"Discovery/Web-Content/burp-parameter-names.txt",
	}
	DefaultDNS = []string{
		"Discovery/DNS/subdomains-top1million-5000.txt",
		"Discovery/DNS/namelist.txt",
	}
)
