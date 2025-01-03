package iterations

func Repeat(item string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += item
	}
	return repeated
}
