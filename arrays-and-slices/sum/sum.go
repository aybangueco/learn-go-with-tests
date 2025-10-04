package sum

// Returns sum of array of numbers.
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
