package week1

func main66() {
	plusOne([]int{1, 2, 3})
	plusOne([]int{1, 2, 9})
	plusOne([]int{9, 9})
}

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for i >= 0 {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}

		digits[i] = 0
		i--
	}

	if i < 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
