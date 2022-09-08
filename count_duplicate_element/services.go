package count_duplicate_element

// ex: list [hello, how, are, you, today, you, fine]
// result gonna be like hello: 1, how: 1, are: 1 today: 1 and you: 2
func CountDuplicateElement(list []string) map[string]int {
	duplicateElement := make(map[string]int)
	// looping list
	for _, item := range list {
		// check if keys found in map
		_, exist := duplicateElement[item]
		if exist {
			duplicateElement[item] += 1 // iteration element
		} else {
			duplicateElement[item] = 1
		}
	}
	return duplicateElement
}
