package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"meb/internal/storage"
)

// OrganizerStore is the subset of storage.Store used by Organizer endpoints.
type OrganizerStore interface {
	AddOrganizer(in storage.AddOrganizerInput) (storage.OrganizerItem, error)
	ListOrganizer(collection string) ([]storage.OrganizerItem, error)
	ListOrganizerCollections() ([]string, error)
	UpdateOrganizer(id, notes, status, highlight, collection string) error
	DeleteOrganizer(id string) error
}

func (s *Server) organizerList(w http.ResponseWriter, r *http.Request) {
	if s.Organizer == nil {
		writeErr(w, 503, "organizer not configured")
		return
	}
	coll := r.URL.Query().Get("collection")
	items, err := s.Organizer.ListOrganizer(coll)
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	if items == nil {
		items = []storage.OrganizerItem{}
	}
	colls, _ := s.Organizer.ListOrganizerCollections()
	writeJSON(w, 200, map[string]any{"items": items, "collections": colls})
}

func (s *Server) organizerAdd(w http.ResponseWriter, r *http.Request) {
	if s.Organizer == nil {
		writeErr(w, 503, "organizer not configured")
		return
	}
	var body struct {
		RequestRaw  string `json:"request_raw"`
		ResponseRaw string `json:"response_raw"`
		Tool        string `json:"tool"`
		Collection  string `json:"collection"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "invalid JSON: "+err.Error())
		return
	}
	if body.RequestRaw == "" {
		writeErr(w, 400, "request_raw is required")
		return
	}
	method, host, path, query := parseRequestMeta(body.RequestRaw)
	in := storage.AddOrganizerInput{
		ID:          uuid.NewString(),
		Tool:        body.Tool,
		Collection:  body.Collection,
		Method:      method,
		Host:        host,
		Path:        path,
		Query:       query,
		RequestRaw:  body.RequestRaw,
		ResponseRaw: body.ResponseRaw,
	}
	if body.ResponseRaw != "" {
		in.StatusCode, in.Length = parseResponseMeta(body.ResponseRaw)
	}
	item, err := s.Organizer.AddOrganizer(in)
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	writeJSON(w, 200, item)
}

func (s *Server) organizerUpdate(w http.ResponseWriter, r *http.Request) {
	if s.Organizer == nil {
		writeErr(w, 503, "organizer not configured")
		return
	}
	id := r.PathValue("id")
	if id == "" {
		writeErr(w, 400, "id is required")
		return
	}
	var body struct {
		Notes      string `json:"notes"`
		Status     string `json:"status"`
		Highlight  string `json:"highlight"`
		Collection string `json:"collection"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "invalid JSON: "+err.Error())
		return
	}
	if err := s.Organizer.UpdateOrganizer(id, body.Notes, body.Status, body.Highlight, body.Collection); err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true})
}

func (s *Server) organizerDelete(w http.ResponseWriter, r *http.Request) {
	if s.Organizer == nil {
		writeErr(w, 503, "organizer not configured")
		return
	}
	id := r.PathValue("id")
	if id == "" {
		writeErr(w, 400, "id is required")
		return
	}
	if err := s.Organizer.DeleteOrganizer(id); err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true})
}

// parseRequestMeta extracts method, host, path, query from a raw HTTP request.
func parseRequestMeta(raw string) (method, host, path, query string) {
	lines := strings.SplitN(raw, "\n", 2)
	parts := strings.Fields(lines[0])
	if len(parts) >= 1 {
		method = parts[0]
	}
	if len(parts) >= 2 {
		path, query = splitPathQuery(parts[1])
	}
	// Host from Host header
	for _, line := range strings.Split(raw, "\n") {
		l := strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(l), "host:") {
			host = strings.TrimSpace(l[len("host:"):])
			break
		}
	}
	return method, host, path, query
}

// parseResponseMeta extracts status code and body length from a raw HTTP response.
func parseResponseMeta(raw string) (status, length int) {
	lines := strings.SplitN(raw, "\n", 2)
	parts := strings.Fields(lines[0])
	if len(parts) >= 2 {
		if code, err := strconv.Atoi(parts[1]); err == nil {
			status = code
		}
	}
	if i := strings.Index(raw, "\r\n\r\n"); i >= 0 {
		length = len(raw) - (i + 4)
	} else if i := strings.Index(raw, "\n\n"); i >= 0 {
		length = len(raw) - (i + 2)
	}
	return status, length
}

// splitPathQuery splits a URL path into path and query.
func splitPathQuery(u string) (string, string) {
	if i := strings.IndexByte(u, '?'); i >= 0 {
		return u[:i], u[i+1:]
	}
	return u, ""
}
