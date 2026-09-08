package intel

import (
	"net/url"
	"sort"
	"strings"
)

// BuildMap groups artifacts into a host → path tree.
func BuildMap(t Target, arts []Artifact) AppMap {
	type hostAcc struct {
		live  bool
		tech  map[string]struct{}
		ports map[string]struct{}
		paths map[string]*PathNode
	}
	hosts := map[string]*hostAcc{}
	ensure := func(h string) *hostAcc {
		h = strings.ToLower(strings.TrimSpace(h))
		if h == "" {
			h = t.Domain
		}
		if acc, ok := hosts[h]; ok {
			return acc
		}
		acc := &hostAcc{
			tech:  map[string]struct{}{},
			ports: map[string]struct{}{},
			paths: map[string]*PathNode{},
		}
		hosts[h] = acc
		return acc
	}
	ensure(t.Domain)
	stats := MapStats{}

	for _, a := range arts {
		switch a.Kind {
		case KindSubdomain:
			ensure(a.Value)
			stats.Subdomains++
		case KindHost:
			acc := ensure(a.Value)
			acc.live = true
		case KindPort:
			h := a.Extra["host"]
			if h == "" {
				h, _, _ = strings.Cut(a.Value, ":")
			}
			ensure(h).ports[a.Value] = struct{}{}
		case KindTech:
			h := a.Extra["host"]
			ensure(h).tech[a.Value] = struct{}{}
			stats.Tech++
		case KindPath:
			h := a.Extra["host"]
			acc := ensure(h)
			p := a.Value
			if p == "" {
				p = "/"
			}
			pn := acc.paths[p]
			if pn == nil {
				pn = &PathNode{Path: p}
				acc.paths[p] = pn
			}
			if st := a.Extra["status"]; st != "" {
				pn.Status = st
			}
		case KindParam:
			stats.Params++
			acc := ensure(t.Domain)
			pn := acc.paths["/"]
			if pn == nil {
				pn = &PathNode{Path: "/"}
				acc.paths["/"] = pn
			}
			if !contains(pn.Params, a.Value) {
				pn.Params = append(pn.Params, a.Value)
			}
		case KindJS:
			stats.JS++
			u, err := url.Parse(a.Value)
			h := t.Domain
			if err == nil && u.Host != "" {
				h = u.Hostname()
			}
			acc := ensure(h)
			p := "/"
			if err == nil && u.Path != "" {
				p = u.Path
			}
			pn := acc.paths[p]
			if pn == nil {
				pn = &PathNode{Path: p}
				acc.paths[p] = pn
			}
			if !contains(pn.JS, a.Value) {
				pn.JS = append(pn.JS, a.Value)
			}
		case KindURL:
			stats.URLs++
			u, err := url.Parse(a.Value)
			if err != nil {
				continue
			}
			acc := ensure(u.Hostname())
			p := u.Path
			if p == "" {
				p = "/"
			}
			pn := acc.paths[p]
			if pn == nil {
				pn = &PathNode{Path: p}
				acc.paths[p] = pn
			}
			if !contains(pn.URLs, a.Value) {
				pn.URLs = append(pn.URLs, a.Value)
			}
			for key := range u.Query() {
				if !contains(pn.Params, key) {
					pn.Params = append(pn.Params, key)
					stats.Params++
				}
			}
		}
	}

	var nodes []HostNode
	for h, acc := range hosts {
		if acc.live {
			stats.LiveHosts++
		}
		n := HostNode{Host: h, Live: acc.live}
		for tech := range acc.tech {
			n.Tech = append(n.Tech, tech)
		}
		sort.Strings(n.Tech)
		for p := range acc.ports {
			n.Ports = append(n.Ports, p)
		}
		sort.Strings(n.Ports)
		for _, pn := range acc.paths {
			sort.Strings(pn.Params)
			sort.Strings(pn.JS)
			sort.Strings(pn.URLs)
			n.Paths = append(n.Paths, *pn)
			stats.Paths++
		}
		sort.Slice(n.Paths, func(i, j int) bool { return n.Paths[i].Path < n.Paths[j].Path })
		nodes = append(nodes, n)
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Host == t.Domain {
			return true
		}
		if nodes[j].Host == t.Domain {
			return false
		}
		return nodes[i].Host < nodes[j].Host
	})
	return AppMap{Domain: t.Domain, Base: t.BaseURL, Hosts: nodes, Stats: stats}
}

func contains(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}
