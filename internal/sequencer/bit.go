package sequencer

import "math"

func analyzeBitLevel(bits []int) BitLevelReport {
	rep := BitLevelReport{TotalBits: len(bits)}
	if len(bits) < 20000 {
		rep.Anomalies = append(rep.Anomalies, "FIPS tests require >= 20000 bits; results are indicative only")
	}
	rep.FIPSTests = append(rep.FIPSTests, monobitTest(bits))
	rep.FIPSTests = append(rep.FIPSTests, pokerTest(bits))
	rep.FIPSTests = append(rep.FIPSTests, runsTest(bits))
	rep.FIPSTests = append(rep.FIPSTests, longRunTest(bits))
	return rep
}

func monobitTest(bits []int) FIPSTest {
	ones := 0
	for _, b := range bits {
		ones += b
	}
	n := len(bits)
	mean := float64(n) / 2
	dev := math.Sqrt(float64(n)) / 2
	lo := mean - 1.96*dev
	hi := mean + 1.96*dev
	pass := ones >= int(lo) && ones <= int(hi)
	z := math.Abs(float64(ones)-mean) / dev
	p := 2 * (1 - normCDF(z))
	return FIPSTest{
		Name:        "Monobit",
		Pass:        pass,
		PValue:      p,
		Description: "Counts the number of 1s. A random stream has ~n/2 ones.",
		Detail:      formatCount("ones", ones, int(lo), int(hi)),
	}
}

func pokerTest(bits []int) FIPSTest {
	if len(bits) < 4 {
		return FIPSTest{Name: "Poker", Pass: false, Description: "too few bits"}
	}
	counts := make([]int, 16)
	for i := 0; i+4 <= len(bits); i += 4 {
		v := bits[i]<<3 | bits[i+1]<<2 | bits[i+2]<<1 | bits[i+3]
		counts[v]++
	}
	m := len(bits) / 4
	if m < 1 {
		return FIPSTest{Name: "Poker", Pass: false, Description: "too few 4-bit blocks"}
	}
	sum := 0
	for _, c := range counts {
		sum += c * c
	}
	x := float64(16)*float64(sum)/float64(m) - float64(m)
	p := 1 - chiSquareCDF(x, 15)
	pass := x > 2.16 && x < 46.17
	if m != 5000 {
		pass = p > 0.01
	}
	return FIPSTest{
		Name:        "Poker",
		Pass:        pass,
		PValue:      p,
		Description: "Chi-square over the 16 possible 4-bit patterns.",
		Detail:      formatF("X=%.2f (FIPS 20000-bit range 2.16..46.17)", x),
	}
}

func runsTest(bits []int) FIPSTest {
	if len(bits) == 0 {
		return FIPSTest{Name: "Runs", Pass: false, Description: "no bits"}
	}
	counts := [2][6]int{}
	cur := bits[0]
	run := 1
	for i := 1; i < len(bits); i++ {
		if bits[i] == cur {
			run++
		} else {
			idx := run
			if idx > 6 {
				idx = 6
			}
			counts[cur][idx-1]++
			cur = bits[i]
			run = 1
		}
	}
	idx := run
	if idx > 6 {
		idx = 6
	}
	counts[cur][idx-1]++

	ranges := [6]struct{ lo, hi int }{
		{2315, 2685}, {1114, 1386}, {527, 723}, {240, 384}, {103, 209}, {103, 209},
	}
	// FIPS ranges are calibrated for 20000 bits; scale for other sample sizes.
	scale := 1.0
	if n := len(bits); n > 0 {
		scale = float64(n) / 20000.0
	}
	pass := true
	detail := ""
	for b := 0; b < 2; b++ {
		for l := 0; l < 6; l++ {
			c := counts[b][l]
			lo := int(math.Round(float64(ranges[l].lo) * scale))
			hi := int(math.Round(float64(ranges[l].hi) * scale))
			ok := c >= lo && c <= hi
			if !ok && len(bits) >= 20000 {
				pass = false
			}
			detail += formatF("bit%d len%d=%d(%d..%d) ", b, l+1, c, lo, hi)
		}
	}
	return FIPSTest{
		Name:        "Runs",
		Pass:        pass,
		PValue:      0.0,
		Description: "Counts runs of consecutive identical bits by length.",
		Detail:      detail,
	}
}

func longRunTest(bits []int) FIPSTest {
	longest := 0
	cur := 1
	for i := 1; i < len(bits); i++ {
		if bits[i] == bits[i-1] {
			cur++
			if cur > longest {
				longest = cur
			}
		} else {
			cur = 1
		}
	}
	pass := longest < 34
	return FIPSTest{
		Name:        "Long run",
		Pass:        pass,
		PValue:      0.0,
		Description: "No run of identical bits may be 34 or longer.",
		Detail:      formatF("longest run=%d (max allowed 33)", longest),
	}
}
