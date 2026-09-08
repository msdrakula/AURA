package batch

// GenerateCombinations returns the cartesian product of the input sets.
// Each combination has the same length as input. An empty input, or any
// empty inner slice, yields an empty result.
func GenerateCombinations(input [][]string) [][]string {
	if len(input) == 0 {
		return [][]string{}
	}
	total := 1
	for _, set := range input {
		if len(set) == 0 {
			return [][]string{}
		}
		next := total * len(set)
		if len(set) > 0 && next/len(set) != total {
			return [][]string{}
		}
		total = next
	}

	out := make([][]string, 0, total)
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
