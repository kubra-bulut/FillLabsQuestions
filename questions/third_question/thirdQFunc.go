package third_question

// MostRepeated function finds the most repeated string in a slice.
func MostRepeated(data []string) string {
	// Map to store frequency of each element
	frequency := make(map[string]int)

	// Count occurrences of each element
	for _, item := range data {
		frequency[item]++
	}

	// Find the most repeated element
	maxCount := 0
	mostRepeatedItem := ""
	for item, count := range frequency {
		if count > maxCount {
			maxCount = count
			mostRepeatedItem = item
		}
	}

	return mostRepeatedItem
}
