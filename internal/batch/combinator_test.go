package batch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateCombinations(t *testing.T) {
	t.Parallel()

	threeByThree := [][]string{
		{"a", "b", "c"},
		{"1", "2", "3"},
		{"x", "y", "z"},
	}
	var twentySeven [][]string
	for _, a := range threeByThree[0] {
		for _, b := range threeByThree[1] {
			for _, c := range threeByThree[2] {
				twentySeven = append(twentySeven, []string{a, b, c})
			}
		}
	}

	tests := []struct {
		name string
		in   [][]string
		want [][]string
	}{
		{
			name: "two groups of two",
			in:   [][]string{{"a", "b"}, {"1", "2"}},
			want: [][]string{{"a", "1"}, {"a", "2"}, {"b", "1"}, {"b", "2"}},
		},
		{
			name: "three groups of three",
			in:   threeByThree,
			want: twentySeven,
		},
		{
			name: "one group of five",
			in:   [][]string{{"a", "b", "c", "d", "e"}},
			want: [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}},
		},
		{
			name: "empty input",
			in:   [][]string{},
			want: [][]string{},
		},
		{
			name: "empty inner slice",
			in:   [][]string{{"a", "b"}, {}, {"1"}},
			want: [][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := GenerateCombinations(tt.in)
			require.Equal(t, tt.want, got)
			if tt.name == "three groups of three" {
				require.Len(t, got, 27)
			}
		})
	}
}

func TestAttackCombinationsCapsCombo(t *testing.T) {
	t.Parallel()
	set := make([]string, 200)
	for i := range set {
		set[i] = "x"
	}
	_, err := attackCombinations("combo", 2, [][]string{set, set})
	require.Error(t, err)
	require.Contains(t, err.Error(), "too many combinations")
}
