package sequencer

import (
	"math"
	"sort"
	"unicode/utf8"
)

// Analyze runs token entropy and frequency tests over a sample.
func Analyze(tokens []string) Report {
	clean := make([]string, 0, len(tokens))
	for _, t := range tokens {
		if t != "" {
			clean = append(clean, t)
		}
	}
	rep := Report{SampleSize: len(clean)}
	if len(clean) == 0 {
		rep.Overall = OverallResult{Rating: "insufficient sample"}
		rep.Reliability = "no data"
		return rep
	}

	minL, maxL := utf8.RuneCountInString(clean[0]), utf8.RuneCountInString(clean[0])
	for _, t := range clean[1:] {
		n := utf8.RuneCountInString(t)
		if n < minL {
			minL = n
		}
		if n > maxL {
			maxL = n
		}
	}
	rep.TokenLength = LengthRange{Min: minL, Max: maxL}

	padded := make([][]rune, len(clean))
	paddedCount := 0
	for i, t := range clean {
		rs := []rune(t)
		if len(rs) < maxL {
			diff := maxL - len(rs)
			rs = append(rs, make([]rune, diff)...)
			paddedCount++
		}
		padded[i] = rs
	}
	rep.Padded = paddedCount

	rep.CharLevel = analyzeCharLevel(padded, minL)
	bits := tokensToBits(clean)
	rep.BitLevel = analyzeBitLevel(bits)
	rep.EffectiveEntropy = effectiveEntropy(rep.CharLevel, rep.BitLevel)

	passed, total := 0, 0
	for _, t := range rep.BitLevel.FIPSTests {
		total++
		if t.Pass {
			passed++
		}
	}
	for _, p := range rep.CharLevel.Positions {
		total++
		if p.Pass {
			passed++
		}
	}
	score := 0.0
	if total > 0 {
		score = float64(passed) / float64(total)
	}
	rep.Overall.Score = score
	switch {
	case len(clean) < 100:
		rep.Overall.Rating = "insufficient sample"
		rep.Reliability = "low - sample < 100 tokens"
	case score >= 0.95:
		rep.Overall.Rating = "good"
		rep.Reliability = "high"
	case score >= 0.8:
		rep.Overall.Rating = "fair"
		rep.Reliability = "medium"
	default:
		rep.Overall.Rating = "poor"
		rep.Reliability = "low - tokens appear non-random"
	}
	return rep
}

func tokensToBits(tokens []string) []int {
	if len(tokens) == 0 {
		return nil
	}
	// Determine common length (min rune length) and per-position charset.
	maxLen := 0
	for _, t := range tokens {
		if n := len([]rune(t)); n > maxLen {
			maxLen = n
		}
	}
	// Build per-position ordered charset.
	charsets := make([]map[rune]int, maxLen)
	for i := range charsets {
		charsets[i] = map[rune]int{}
	}
	for _, t := range tokens {
		rs := []rune(t)
		for i, r := range rs {
			if i < maxLen {
				charsets[i][r] = 0
			}
		}
	}
	// Assign indices in rune-sorted order for determinism.
	for i := range charsets {
		keys := make([]rune, 0, len(charsets[i]))
		for r := range charsets[i] {
			keys = append(keys, r)
		}
		sort.Slice(keys, func(a, b int) bool { return keys[a] < keys[b] })
		for idx, r := range keys {
			charsets[i][r] = idx
		}
	}
	// Bits per position = floor(log2(roundedDownPowerOf2(k))).
	bitsPerPos := make([]int, maxLen)
	for i, cs := range charsets {
		k := len(cs)
		if k < 2 {
			bitsPerPos[i] = 0
			continue
		}
		p := 2
		for p*2 <= k {
			p *= 2
		}
		bitsPerPos[i] = int(math.Log2(float64(p)))
	}
	// Emit bits per token per position.
	bits := make([]int, 0)
	for _, t := range tokens {
		rs := []rune(t)
		for i, r := range rs {
			if i >= maxLen {
				break
			}
			nbits := bitsPerPos[i]
			if nbits == 0 {
				continue
			}
			idx := charsets[i][r]
			// mask idx to nbits bits (drop high bits if charset not power of 2)
			masked := idx & ((1 << uint(nbits)) - 1)
			for b := nbits - 1; b >= 0; b-- {
				bits = append(bits, (masked>>uint(b))&1)
			}
		}
	}
	return bits
}
