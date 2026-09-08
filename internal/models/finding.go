package models

// Severity is a coarse finding level shown in the UI.
type Severity string

const (
	SeverityInfo   Severity = "info"
	SeverityLow    Severity = "low"
	SeverityMedium Severity = "medium"
	SeverityHigh   Severity = "high"
)

// AnalysisFinding is one rule hit tied to a stored transaction.
type AnalysisFinding struct {
	ID            string   `json:"id" gorm:"column:id;primaryKey" db:"id"`
	TransactionID string   `json:"transaction_id" gorm:"column:transaction_id;index" db:"transaction_id"`
	RuleName      string   `json:"rule_name" gorm:"column:rule_name" db:"rule_name"`
	Severity      Severity `json:"severity" gorm:"column:severity" db:"severity"`
	Title         string   `json:"title" gorm:"column:title" db:"title"`
	Description   string   `json:"description" gorm:"column:description" db:"description"`
	Evidence      string   `json:"evidence" gorm:"column:evidence" db:"evidence"`
}
