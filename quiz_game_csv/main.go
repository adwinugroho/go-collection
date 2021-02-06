package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readCSV() ([]string, error) {
	// load csv file
	var itemInRecord []string
	f, _ := os.Open("problem.csv")

	// create a new reader
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error while reading file, cause:%+v\n", err)
		}
		itemInRecord = append(itemInRecord, record...)

	}
	return itemInRecord, nil
}

func oddNumber(number int) bool {
	return number%2 == 1
}

func answerQuestion() []string {
	var newItem []string
	itemCSV, err := readCSV()
	if err != nil {
		panic(err)
	}
	for indeks, itemInCSVFile := range itemCSV {
		checkIndeks := oddNumber(indeks)
		if checkIndeks {
			newItem = append(newItem, itemInCSVFile)
		}
	}
	return newItem
}

func questionProblem() []string {
	var newItem []string
	itemCSV, err := readCSV()
	if err != nil {
		panic(err)
	}
	for indeks, itemInCSVFile := range itemCSV {
		checkIndeks := oddNumber(indeks)
		if !checkIndeks {
			newItem = append(newItem, itemInCSVFile)
		}
	}
	return newItem
}

func CSVToMap(reader io.Reader) []map[string]string {
	r := csv.NewReader(reader)
	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	return rows
}

func main() {
	var correct int = 0
	var value string
	question := questionProblem()
	answer := answerQuestion()
	for i, item := range question {
		fmt.Printf("question number %d. %s?, ", i+1, item)
		fmt.Print("your answer= ")
		fmt.Scanln(&value)
		if value == answer[i] {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(question))
}
