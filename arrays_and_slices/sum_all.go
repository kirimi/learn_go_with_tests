package main

func SumAll(slices ...[]int) []int {
	result := make([]int, 0, len(slices))

	for _, slice := range slices {
		sum := Sum(slice)
		result = append(result, sum)
	}

	return result
}
