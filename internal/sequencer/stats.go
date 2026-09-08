package sequencer

import (
	"fmt"
	"math"
)

func normCDF(x float64) float64 {
	return 0.5 * (1 + math.Erf(x/math.Sqrt2))
}

// chiSquareCDF returns the cumulative distribution function of the
// chi-square distribution with k degrees of freedom, using the regularized
// lower incomplete gamma function.
func chiSquareCDF(x float64, k int) float64 {
	if x <= 0 {
		return 0
	}
	return lowerRegularizedGamma(float64(k)/2, x/2)
}

// lowerRegularizedGamma is P(a, x) = gamma(a, x) / Gamma(a).
// Implemented via the series and continued-fraction expansions from
// Numerical Recipes.
func lowerRegularizedGamma(a, x float64) float64 {
	if x < 0 || a <= 0 {
		return 0
	}
	if x < a+1 {
		return seriesGamma(a, x)
	}
	return 1 - continuedFractionGamma(a, x)
}

func seriesGamma(a, x float64) float64 {
	const maxIter = 200
	const eps = 1e-14
	ap := a
	sum := 1.0 / a
	term := sum
	for n := 0; n < maxIter; n++ {
		ap++
		term *= x / ap
		sum += term
		if math.Abs(term) < math.Abs(sum)*eps {
			break
		}
	}
	// sum * exp(-x + a*ln(x) - lgamma(a))
	gln, _ := math.Lgamma(a)
	return sum * math.Exp(-x+a*math.Log(x)-gln)
}

func continuedFractionGamma(a, x float64) float64 {
	const maxIter = 300
	const eps = 1e-14
	const fpmin = 1e-300
	b := x + 1 - a
	c := 1.0 / fpmin
	d := 1.0 / b
	h := d
	for i := 1; i <= maxIter; i++ {
		an := -float64(i) * (float64(i) - a)
		b += 2
		d = an*d + b
		if math.Abs(d) < fpmin {
			d = fpmin
		}
		c = b + an/c
		if math.Abs(c) < fpmin {
			c = fpmin
		}
		d = 1 / d
		delta := d * c
		h *= delta
		if math.Abs(delta-1) < eps {
			break
		}
	}
	gln, _ := math.Lgamma(a)
	return math.Exp(-x+a*math.Log(x)-gln) * h
}

// format helpers used by the FIPS tests.
func formatF(layout string, args ...any) string {
	return fmt.Sprintf(layout, args...)
}

func formatCount(name string, val, lo, hi int) string {
	return fmt.Sprintf("%s=%d (expected %d..%d)", name, val, lo, hi)
}
