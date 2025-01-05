package main

// Sum int items
func Sum(items []int) int {
	if len(items) == 0 {
		return 0
	}

	res := 0
	for _, item := range items {
		res += item
	}
	return res
}
