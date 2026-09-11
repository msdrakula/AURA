package batch

const MaxCombinations = 20000

// GenerateCombinations returns the cartesian product of the input sets.
// Each combination has the same length as input. An empty input, or any
// empty inner slice, yields an empty result. Overflow returns empty
// (attackCombinations rejects oversized products before calling this).
func GenerateCombinations(input [][]string) [][]string {
	n, overflow := combinationCount(input)
	if overflow || n == 0 {
		return [][]string{}
	}
	out := make([][]string, 0, n)
	idx := make([]int, len(input))
	for {
		combo := make([]string, len(input))
		for i, set := range input {
			combo[i] = set[idx[i]]
		}
		out = append(out, combo)

		pos := len(input) - 1
		for pos >= 0 {
			idx[pos]++
			if idx[pos] < len(input[pos]) {
				break
			}
			idx[pos] = 0
			pos--
		}
		if pos < 0 {
			break
		}
	}
	return out
}

func combinationCount(sets [][]string) (n int, overflow bool) {
	if len(sets) == 0 {
		return 0, false
	}
	n = 1
	for _, set := range sets {
		if len(set) == 0 {
			return 0, false
		}
		next := n * len(set)
		if next/len(set) != n {
			return 0, true
		}
		n = next
	}
	return n, false
}
