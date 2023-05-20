package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/h2non/filetype"
)

func main() {
	// "https://gorm.io/sponsors-imgs/bytebase.png"
	// "https://bpsdm.pu.go.id/center/pelatihan/uploads/edok/2019/02/bf105_MDL_Sistem_Informasi_dan_Data_SDA.docx"
	// "https://www.stat.ipb.ac.id/agusms/wp-content/uploads/2022/02/Mater-4-Visualisasi-Data-1.pdf"
	fileBytes, err := GetFileFromURL("https://bpsdm.pu.go.id/center/pelatihan/uploads/edok/2019/02/bf105_MDL_Sistem_Informasi_dan_Data_SDA.docx")
	if err != nil {
		panic(err)
	}
	// MIME Type -> image
	// MIME SubType -> extension/document type
	// MIME Value -> content type
	if fileBytes == nil {
		log.Println("data not found!")
		return
	}
	typeFile, _ := filetype.Get(fileBytes)
	filename := fmt.Sprintf("./download/download.%s", typeFile.Extension)
	fmt.Println(typeFile.MIME.Subtype)
	err = os.WriteFile(filename, fileBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func GetFileFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error message: %+v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error message: %+v\n", err)
		return nil, err
	}
	// check if status not OK
	if resp.StatusCode != 200 {
		log.Printf("error message: %+v\n", err)
		return nil, nil
	}
	return bodyBytes, nil
}
