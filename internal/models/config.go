package models

// ProjectConfig is the persisted workspace settings for the local proxy UI.
type ProjectConfig struct {
	Name             string   `json:"name" gorm:"column:name" db:"name"`
	ProxyPort        int      `json:"proxy_port" gorm:"column:proxy_port" db:"proxy_port"`
	InterceptEnabled bool     `json:"intercept_enabled" gorm:"column:intercept_enabled" db:"intercept_enabled"`
	ScopeRules       []string `json:"scope_rules" gorm:"serializer:json" db:"scope_rules"`
}
