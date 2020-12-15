package try_metahphone

import (
	"fmt"
	"strings"

	"github.com/dotcypress/phonetics"
)

type interfaceSvcMetaphone struct{}

func GetSvcMetaphone() *interfaceSvcMetaphone {
	return &interfaceSvcMetaphone{}
}

func (service *interfaceSvcMetaphone) TryMetaphone() {
	//initial nama dokter
	namaDokter := "dr. Agus Darajat, SpA, M.Kes"
	var initMetaphone, resMetaphone string
	var initArray, splitNamaDokter []string
	var checkDot, checkComma bool
	//initial untuk cek titik sama koma
	checkDot = strings.Contains(namaDokter, ".")
	checkComma = strings.Contains(namaDokter, ",")
	fmt.Printf("check dot in name: %+v\n", checkDot)
	fmt.Printf("check comma in name: %+v\n", checkComma)
	//cek titik
	if checkDot {
		namaDokter = strings.Replace(namaDokter, ".", " ", -1)
		fmt.Printf("check namaDokter after checkDot true: %+v\n", namaDokter)
	}
	//cek koma
	if checkComma {
		namaDokter = strings.Replace(namaDokter, ",", " ", -1)
		fmt.Printf("check namaDokter after checkComma true: %+v\n", namaDokter)
	}
	//split by space
	splitNamaDokter = strings.Split(namaDokter, " ")
	fmt.Printf("check splitNamaDokter: %+v\n", splitNamaDokter)
	//loop slice (nama dokter)
	for _, sliceInNama := range splitNamaDokter {
		initMetaphone = phonetics.EncodeMetaphone(string(sliceInNama))
		initArray = append(initArray, initMetaphone)
	}
	//join array yang di initial diatas dengan ""
	resMetaphone = strings.Join(initArray, "")
	//tes := phonetics.EncodeMetaphone()
	fmt.Println("---- hasil ----")
	fmt.Printf("check resMetaphone: %+v\n", resMetaphone)
}
