package models

import "time"

// ReconItemKind is the category of a collected recon artifact.
type ReconItemKind string

const (
	KindSubdomain ReconItemKind = "subdomain"
	KindPort      ReconItemKind = "port"
	KindPath      ReconItemKind = "path"
	KindParam     ReconItemKind = "param"
	KindTech      ReconItemKind = "tech"
	KindJS        ReconItemKind = "js"
	KindVuln      ReconItemKind = "vuln"
	KindDoc       ReconItemKind = "doc"
)

// ReconTarget is one researched host. All collected artifacts hang off it.
type ReconTarget struct {
	ID        string    `json:"id"`
	Host      string    `json:"host"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	Notes     string    `json:"notes"`
}

// ReconItem is a single collected artifact tied to a target.
// Meta is a free-form JSON blob carrying kind-specific details
// (status code, length, service, severity, evidence, ...).
type ReconItem struct {
	ID        string         `json:"id"`
	TargetID  string         `json:"target_id"`
	Stage     string         `json:"stage"` // mindmap stage that produced it
	Kind      ReconItemKind  `json:"kind"`
	Value     string         `json:"value"`
	Meta      map[string]any `json:"meta"`
	Created   time.Time      `json:"created"`
}
