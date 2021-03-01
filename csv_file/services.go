package csv_file

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"strconv"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func CreateNewCsvFile() error {

	users = append(users, User{ID: "1", Username: "Pangeran", Email: "bangpange@gmail.com"})
	users = append(users, User{ID: "2", Username: "Sihaan", Email: "bangsihaan@gmail.com"})

	var dataCSV [][]string
	var empData []string
	var initialHeader []string

	initialHeader = append(initialHeader, "No", "ID", "Username", "E-Mail")
	dataCSV = append(dataCSV, initialHeader)
	var count int
	for _, item := range users {
		count++
		empData = []string{strconv.Itoa(count), item.ID, item.Username, item.Email}
		dataCSV = append(dataCSV, empData)
	}

	dir := "upload"
	os.Mkdir(dir, 0666)
	fileName := path.Join(dir, "file.csv")
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Printf("failed creating file: %s", err)
		return err
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	for _, empRow := range dataCSV {
		_ = csvwriter.Write(empRow)
	}

	return nil
}
