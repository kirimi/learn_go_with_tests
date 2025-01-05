package main

func SumAll(slices ...[]int) []int {
	result := []int{}

	for _, slice := range slices {
		sum := Sum(slice)
		result = append(result, sum)
	}

	return result
}
