// Package sequencer analyzes the randomness quality of a sample of tokens
// using public FIPS-140-2 style bit tests, character-level frequency tests,
// and an effective-entropy estimate.
package sequencer

// Report is the result of analyzing a sample of tokens.
type Report struct {
	SampleSize       int             `json:"sample_size"`
	TokenLength      LengthRange     `json:"token_length"`
	Padded           int             `json:"padded"`
	Overall          OverallResult   `json:"overall"`
	EffectiveEntropy EntropyChart    `json:"effective_entropy"`
	BitLevel         BitLevelReport  `json:"bit_level"`
	CharLevel        CharLevelReport `json:"char_level"`
	Reliability      string          `json:"reliability"`
}

type LengthRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type OverallResult struct {
	Rating string  `json:"rating"`
	Score  float64 `json:"score"`
}

type EntropyChart struct {
	Levels []EntropyPoint `json:"levels"`
}

type EntropyPoint struct {
	Level       string  `json:"level"`
	BitsPassing int     `json:"bits_passing"`
	PValue      float64 `json:"p_value"`
}

type BitLevelReport struct {
	TotalBits int        `json:"total_bits"`
	FIPSTests []FIPSTest `json:"fips_tests"`
	Anomalies []string   `json:"anomalies"`
}

type FIPSTest struct {
	Name        string  `json:"name"`
	Pass        bool    `json:"pass"`
	PValue      float64 `json:"p_value"`
	Description string  `json:"description"`
	Detail      string  `json:"detail"`
}

type CharLevelReport struct {
	Positions    []PositionAnalysis `json:"positions"`
	Significance []EntropyPoint     `json:"significance"`
}

type PositionAnalysis struct {
	Index          int            `json:"index"`
	CharSetSize    int            `json:"charset_size"`
	MaxEntropyBits float64        `json:"max_entropy_bits"`
	Frequency      map[string]int `json:"frequency"`
	ChiSquare      float64        `json:"chi_square"`
	PValue         float64        `json:"p_value"`
	Pass           bool           `json:"pass"`
}
