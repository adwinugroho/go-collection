package count_duplicate_element

// ex: list [hello, how, are, today, you, fine]
// result gonna be like hello: 1, how: 1, are: 1 today: 1 and you: 2
func CountDuplicateElement(list []string) map[string]int {
	DuplicateElement := make(map[string]int)
	// looping list
	for _, item := range list {
		// check if keys found in map
		_, exist := DuplicateElement[item]
		if exist {
			DuplicateElement[item] += 1 // iteration element
		} else {
			DuplicateElement[item] = 1
		}
	}
	return DuplicateElement
}
