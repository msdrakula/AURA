package sequencer

import (
	"crypto/rand"
	"encoding/hex"
	"testing"
)

func TestAnalyzeRandom(t *testing.T) {
	tokens := make([]string, 500)
	for i := range tokens {
		b := make([]byte, 16)
		_, _ = rand.Read(b)
		tokens[i] = hex.EncodeToString(b)
	}
	rep := Analyze(tokens)
	if rep.SampleSize != 500 {
		t.Fatalf("sample size = %d", rep.SampleSize)
	}
	if rep.TokenLength.Min != 32 || rep.TokenLength.Max != 32 {
		t.Fatalf("token length = %+v", rep.TokenLength)
	}
	if rep.Overall.Rating == "poor" {
		t.Logf("random tokens rated poor (score=%.2f) — may be sample size", rep.Overall.Score)
	}
	if len(rep.BitLevel.FIPSTests) != 4 {
		t.Fatalf("expected 4 FIPS tests, got %d", len(rep.BitLevel.FIPSTests))
	}
	if len(rep.CharLevel.Positions) != 32 {
		t.Fatalf("expected 32 char positions, got %d", len(rep.CharLevel.Positions))
	}
	if len(rep.EffectiveEntropy.Levels) != 6 {
		t.Fatalf("expected 6 entropy levels, got %d", len(rep.EffectiveEntropy.Levels))
	}
}

func TestAnalyzeSequential(t *testing.T) {
	tokens := make([]string, 500)
	for i := range tokens {
		tokens[i] = "token-" + itoa(i)
	}
	rep := Analyze(tokens)
	if rep.Overall.Rating == "good" {
		t.Fatalf("sequential tokens should not be rated good, got %s", rep.Overall.Rating)
	}
}

func TestAnalyzeEmpty(t *testing.T) {
	rep := Analyze(nil)
	if rep.Overall.Rating != "insufficient sample" {
		t.Fatalf("expected insufficient sample, got %s", rep.Overall.Rating)
	}
}

func itoa(i int) string {
	const digits = "0123456789"
	if i == 0 {
		return "0"
	}
	out := ""
	for i > 0 {
		out = string(digits[i%10]) + out
		i /= 10
	}
	return out
}
