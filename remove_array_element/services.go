package remove_array_element

import "fmt"

// this for implement observer pattern
type interfaceSvcRemoveElementArr struct{}

func GetSvcRemoveElementArr() *interfaceSvcRemoveElementArr {
	return &interfaceSvcRemoveElementArr{}
}

// initial reciever to get func
func (service *interfaceSvcRemoveElementArr) RemoveArrayElement() {
	dataReq := "Lab Update" //data yang di request
	dataHist := []string{"Lab Dua", "Lab Satu", "Test"}
	dataOld := []string{"Lab Satu", "Lab Dua"}
	dataOldOne := dataOld[0] //data yang di update
	fmt.Printf("Data Hist Before:%s\n", dataHist)
	if dataOldOne != dataReq {
		for i, dataInHist := range dataHist {
			if dataInHist == dataOldOne {
				dataHist = append(dataHist[:i], dataHist[i+1:]...) //hapus data yang lama
				dataHist = append(dataHist, dataReq)               // append data yang baru
			}
		}
	}
	fmt.Printf("Data Hist After:%s\n", dataHist)
}
