package extractor

import "regexp"

var jsFinders = []*regexp.Regexp{
	regexp.MustCompile(`(?i)\bfetch\s*\(\s*'([^']+)'`),
	regexp.MustCompile(`(?i)\bfetch\s*\(\s*"([^"]+)"`),
	regexp.MustCompile("(?i)\\bfetch\\s*\\(\\s*`([^`$]+)`"),
	regexp.MustCompile(`(?i)\baxios(?:\s*\.\s*(?:get|post|put|patch|delete|head|request))?\s*\(\s*'([^']+)'`),
	regexp.MustCompile(`(?i)\baxios(?:\s*\.\s*(?:get|post|put|patch|delete|head|request))?\s*\(\s*"([^"]+)"`),
	regexp.MustCompile(`(?i)\.open\s*\(\s*'[A-Z]+'\s*,\s*'([^']+)'`),
	regexp.MustCompile(`(?i)\.open\s*\(\s*"[A-Z]+"\s*,\s*"([^"]+)"`),
	regexp.MustCompile(`(?i)'(https?://[^']+)'`),
	regexp.MustCompile(`(?i)"(https?://[^"]+)"`),
	regexp.MustCompile("`https?://[^`$]+`"),
	regexp.MustCompile(`'(/[A-Za-z0-9._~!$&()*+,;=:@%/?#-]+)'`),
	regexp.MustCompile(`"(/[A-Za-z0-9._~!$&()*+,;=:@%/?#-]+)"`),
	regexp.MustCompile("`(/[A-Za-z0-9._~!$&()*+,;=:@%/?#-]+)`"),
}

func findJSURLs(src string) []string {
	if src == "" {
		return nil
	}
	var out []string
	for _, re := range jsFinders {
		for _, m := range re.FindAllStringSubmatch(src, -1) {
			switch {
			case len(m) > 1 && m[1] != "":
				out = append(out, m[1])
			case len(m) == 1 && len(m[0]) >= 2:
				// backtick abs URL: whole match includes quotes
				out = append(out, trimJSQuotes(m[0]))
			}
		}
	}
	return out
}

func trimJSQuotes(s string) string {
	if len(s) >= 2 {
		switch s[0] {
		case '\'', '"', '`':
			return s[1 : len(s)-1]
		}
	}
	return s
}
