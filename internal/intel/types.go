// Package intel runs guided recon stages and builds an application map.
package intel

import "time"

const (
	KindSubdomain = "subdomain"
	KindHost      = "host"
	KindPort      = "port"
	KindURL       = "url"
	KindPath      = "path"
	KindParam     = "param"
	KindJS        = "js"
	KindTech      = "tech"
	KindNote      = "note"
)

const (
	ModePassive = "passive"
	ModeActive  = "active"
)

const (
	StatusIdle    = "idle"
	StatusRunning = "running"
	StatusDone    = "done"
	StatusSkipped = "skipped"
	StatusError   = "error"
)

// Target is one investigation workspace.
type Target struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Domain     string    `json:"domain"`
	BaseURL    string    `json:"base_url"`
	Authorized bool      `json:"authorized"`
	Created    time.Time `json:"created"`
}

// Artifact is one collected fact about the target.
type Artifact struct {
	ID       string            `json:"id"`
	TargetID string            `json:"target_id"`
	Kind     string            `json:"kind"`
	Value    string            `json:"value"`
	Extra    map[string]string `json:"extra,omitempty"`
	Source   string            `json:"source"`
	Created  time.Time         `json:"created"`
}

// StageState is progress for one mindmap step.
type StageState struct {
	ID      string    `json:"id"`
	Status  string    `json:"status"`
	Summary string    `json:"summary"`
	Updated time.Time `json:"updated"`
}

// CatalogStage describes a step shown in the UI.
type CatalogStage struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Hint  string `json:"hint"`
	Mode  string `json:"mode"`
}

// HostNode is one host in the application map.
type HostNode struct {
	Host  string     `json:"host"`
	Live  bool       `json:"live"`
	Tech  []string   `json:"tech"`
	Ports []string   `json:"ports"`
	Paths []PathNode `json:"paths"`
}

// PathNode is a URL path with related params and scripts.
type PathNode struct {
	Path   string   `json:"path"`
	Status string   `json:"status,omitempty"`
	Params []string `json:"params,omitempty"`
	JS     []string `json:"js,omitempty"`
	URLs   []string `json:"urls,omitempty"`
}

// AppMap is the assembled picture of the investigated app.
type AppMap struct {
	Domain string     `json:"domain"`
	Base   string     `json:"base_url"`
	Hosts  []HostNode `json:"hosts"`
	Stats  MapStats   `json:"stats"`
}

// MapStats is a compact count of collected facts.
type MapStats struct {
	Subdomains int `json:"subdomains"`
	LiveHosts  int `json:"live_hosts"`
	Paths      int `json:"paths"`
	Params     int `json:"params"`
	JS         int `json:"js"`
	Tech       int `json:"tech"`
	URLs       int `json:"urls"`
}

// Store persists targets, stages, and artifacts.
type Store interface {
	CreateTarget(t Target) error
	GetTarget(id string) (Target, error)
	ListTargets() ([]Target, error)
	UpdateTarget(t Target) error
	UpsertStage(targetID string, st StageState) error
	ListStages(targetID string) ([]StageState, error)
	AddArtifact(a Artifact) error
	ListArtifacts(targetID string) ([]Artifact, error)
}

// RunResult is what a stage returns to the API.
type RunResult struct {
	Stage     StageState `json:"stage"`
	Added     int        `json:"added"`
	Artifacts []Artifact `json:"artifacts"`
}
