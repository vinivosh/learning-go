package main

func Sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lenNums := len(numbersToSum)
	results := make([]int, lenNums)

	for idx, numbers := range numbersToSum {
		results[idx] = Sum(numbers)
	}

	return results
}
