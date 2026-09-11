// Package wordlist holds small built-in dictionaries for content and parameter discovery.
package wordlist

// Dirs is a compact path list for directory/file discovery (gobuster dir analog).
func Dirs() []string {
	return append([]string{}, dirs...)
}

// Params is a compact query-parameter name list (Arjun/ffuf analog).
func Params() []string {
	return append([]string{}, params...)
}

// HostPrefixes is a short DNS prefix list for common hostnames.
func HostPrefixes() []string {
	return append([]string{}, hostPrefixes...)
}

var dirs = []string{
	"robots.txt", "sitemap.xml", "favicon.ico", "security.txt", ".well-known/security.txt",
	"admin", "login", "dashboard", "panel", "console", "account", "user", "users", "auth",
	"api", "api/v1", "api/v2", "api/docs", "graphql", "swagger", "swagger.json", "swagger-ui",
	"openapi.json", "docs", "redoc", "health", "healthz", "ready", "status", "metrics",
	"actuator", "actuator/health", "debug", "test", "internal", "private",
	"static", "assets", "js", "css", "images", "uploads", "files", "media",
	"backup", "backups", "config", "configs", "env", ".env", ".git", ".git/HEAD",
	"wp-admin", "wp-login.php", "phpmyadmin", "server-status",
	"search", "graphql/console", "v1", "v2", "rest", "oauth", "token",
	"manifest.json", "package.json",
}

var params = []string{
	"id", "page", "q", "query", "search", "user", "username", "email", "token",
	"redirect", "url", "next", "return", "callback", "file", "path", "format",
	"lang", "debug", "sort", "filter", "limit", "offset", "key", "api_key",
	"apikey", "ref", "source", "view", "action", "include", "template",
}

var hostPrefixes = []string{
	"www", "api", "app", "admin", "dev", "test", "staging", "stage", "uat",
	"mail", "cdn", "static", "img", "media", "docs", "blog", "shop", "portal",
	"auth", "sso", "vpn", "git", "ci", "monitor", "status", "beta",
	"www2", "api-v1", "api-v2", "backend", "front", "frontend", "internal",
	"intranet", "old", "new", "demo", "sandbox", "qa", "preprod", "prod",
	"m", "mobile", "account", "accounts", "login", "id", "ids", "oauth",
	"graphql", "ws", "socket", "assets", "files", "upload", "uploads",
	"db", "sql", "mysql", "postgres", "redis", "elastic", "kibana",
	"grafana", "prometheus", "jenkins", "gitlab", "github", "jira",
	"confluence", "wiki", "help", "support", "cdn1", "cdn2", "img1",
	"ns1", "ns2", "mx", "smtp", "imap", "pop", "ftp", "sftp", "ssh",
	"remote", "office", "webmail", "cpanel", "whm", "panel", "dashboard",
	"cms", "crm", "erp", "pay", "payments", "billing", "invoice", "shop",
	"store", "cart", "checkout", "search", "news", "forum", "chat",
	"dev1", "dev2", "test1", "test2", "stg", "preview", "edge", "origin",
}
