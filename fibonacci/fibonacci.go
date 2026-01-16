package fibonacci

// Calculate returns the first n numbers of the Fibonacci series.
// If n <= 0, it returns an empty slice.
// If n == 1, it returns [0].
func Calculate(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	series := make([]int, n)
	series[0] = 0
	series[1] = 1
	for i := 2; i < n; i++ {
		series[i] = series[i-1] + series[i-2]
	}
	return series
}
