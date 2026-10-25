package main

func Sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var results []int

	for _, numbers := range numbersToSum {
		results = append(results, Sum(numbers))
	}

	return results
}

func SumAllTails(numbersToSum ...[]int) []int {
	var results []int

	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			results = append(results, Sum(numbers[1:]))
		} else {
			results = append(results, Sum(numbers))
		}
	}

	return results
}
