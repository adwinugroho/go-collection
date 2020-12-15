package main

import (
	"fmt"

	csvFile "github.com/adwinugroho/go-collection/csv_file"
	"github.com/adwinugroho/go-collection/generate_random_number"
	"github.com/adwinugroho/go-collection/generate_random_string"
	"github.com/adwinugroho/go-collection/remove_array_element"
)

func main() {
	fmt.Println("this function main")
	fmt.Println("===========================")
	getSvcRemoveElement := remove_array_element.GetSvcRemoveElementArr()
	getSvcRemoveElement.RemoveArrayElement()
	err := csvFile.GetSvcCsvFile().CreateNewCsvFile()
	if err != nil {
		panic(err)
	}
	stringRandom := generate_random_string.GenerateRandomString(5)
	intRandom := generate_random_number.GenerateRandomNumber(6)
	fmt.Println(stringRandom)
	fmt.Println(intRandom)
}
