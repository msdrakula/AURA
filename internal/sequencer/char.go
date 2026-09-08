package sequencer

import "math"

func analyzeCharLevel(tokens [][]rune, commonLen int) CharLevelReport {
	rep := CharLevelReport{}
	if commonLen == 0 {
		return rep
	}
	for pos := 0; pos < commonLen; pos++ {
		freq := map[string]int{}
		for _, t := range tokens {
			if pos < len(t) {
				freq[string(t[pos])]++
			}
		}
		k := len(freq)
		if k < 2 {
			rep.Positions = append(rep.Positions, PositionAnalysis{
				Index: pos, CharSetSize: k, MaxEntropyBits: 0,
				Frequency: freq, Pass: k > 1,
			})
			continue
		}
		n := len(tokens)
		expected := float64(n) / float64(k)
		chi := 0.0
		for _, c := range freq {
			d := float64(c) - expected
			chi += d * d / expected
		}
		df := k - 1
		p := 1 - chiSquareCDF(chi, df)
		pass := p > 0.01
		maxBits := math.Log2(float64(k))
		rep.Positions = append(rep.Positions, PositionAnalysis{
			Index:          pos,
			CharSetSize:    k,
			MaxEntropyBits: maxBits,
			Frequency:      freq,
			ChiSquare:      chi,
			PValue:         p,
			Pass:           pass,
		})
	}
	levels := []string{"0.001%", "0.01%", "0.1%", "1%", "5%", "10%"}
	pvals := []float64{0.00001, 0.0001, 0.001, 0.01, 0.05, 0.10}
	for i, lvl := range levels {
		bits := 0
		for _, p := range rep.Positions {
			if p.Pass && p.PValue > pvals[i] {
				bits += int(math.Floor(p.MaxEntropyBits))
			}
		}
		rep.Significance = append(rep.Significance, EntropyPoint{Level: lvl, BitsPassing: bits, PValue: pvals[i]})
	}
	return rep
}

func effectiveEntropy(ch CharLevelReport, bl BitLevelReport) EntropyChart {
	out := EntropyChart{}
	levels := []string{"0.001%", "0.01%", "0.1%", "1%", "5%", "10%"}
	pvals := []float64{0.00001, 0.0001, 0.001, 0.01, 0.05, 0.10}
	for i, lvl := range levels {
		bits := 0
		for _, p := range ch.Positions {
			if p.Pass && p.PValue > pvals[i] {
				bits += int(math.Floor(p.MaxEntropyBits))
			}
		}
		for _, t := range bl.FIPSTests {
			if t.Pass {
				bits++
			}
		}
		out.Levels = append(out.Levels, EntropyPoint{Level: lvl, BitsPassing: bits, PValue: pvals[i]})
	}
	return out
}
