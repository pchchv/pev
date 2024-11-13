package pev

import "math"

func logX(base, n float64) float64 {
	if base == 0 {
		return 0
	}

	// change of base formulae
	return math.Log2(n) / math.Log2(base)
}

// logPow calculates log_base(x^y)
// without leaving logspace for each multiplication step,
// allowing it to take up less memory space.
func logPow(expBase float64, pow int, logBase float64) (total float64) {
	// logb (MN) = logb M + logb N
	for i := 0; i < pow; i++ {
		total += logX(logBase, expBase)
	}

	return
}

func getEntropy(password string) float64 {
	base := getBase(password)
	length := getLength(password)

	// calculate log2(base^length)
	return logPow(float64(base), length, 2)
}
