package main

import (
	"fmt"

	"github.com/adwinugroho/go-collection/remove_array_element"
)

func main() {
	fmt.Println("this function main")
	fmt.Println("===========================")
	getSvcRemoveElement := remove_array_element.GetSvcRemoveElementArr()
	getSvcRemoveElement.RemoveArrayElement()

}
