package main

import (
	"fmt"

	"github.com/adwinugroho/go-collection/generate_random_number"
	"github.com/adwinugroho/go-collection/generate_random_string"
)

func main() {
	fmt.Println("this function main")
	fmt.Println("===========================")
	fmt.Println("exampe call package")
	stringRandom := generate_random_string.GenerateRandomString(5)
	intRandom := generate_random_number.GenerateRandomNumber(6)
	fmt.Println(stringRandom)
	fmt.Println(intRandom)
}
